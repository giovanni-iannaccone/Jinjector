![Logo](https://github.com/user-attachments/assets/f72b05a4-e65d-4e34-ac12-b07e5f9d090e)

Jinjector is a powerful tool written in Go, designed to inject backdoors into Joomla modules effortlessly. Using this tool, you can easily insert a PHP reverse shell into a Joomla module's main file, allowing a connection back to a specified IP and port whenever the module is triggered. This is perfect for penetration testers or researchers aiming to simulate real-world scenarios.

> [!CAUTION]
> This tool is intended for educational and ethical testing purposes only. Unauthorized use of this tool on live systems without permission is illegal and unethical.

## ✨ Features 
- Automatic file discovery: Extracts information from the Joomla XML manifest file to locate the main file.
- Stealthy injection: Injects a PHP reverse shell that connects back to your specified IP and port.
- Ease of use: Specify your IP, port, and the module directory, and let Jinjector handle the rest.

## 📦 Requirements
- Go
- Joomla module directory (with a valid manifest.xml file)

## 🛠️ Installation
1. Clone the repository 
```bash
git clone https://github.com/giovanni-iannaccone/Jinjector
cd Jinjector
```
2. Build 
```bash
cd cmd
go build -o jinjector
```

## 🚀 Usage
Start the program, enter your ip, your port and the module path

## 🌍 How It Works 
**Manifest Extraction**: Jinjector parses the XML file in the given module directory to identify the main PHP file.
**Payload Injection**: Once located, it appends a PHP reverse shell payload to the main file.
**Connection Setup**: Every time the infected Joomla module is used, it attempts to establish a connection to your IP and port.

## ⚡️ Reverse Shell Code
The reverse shell code injected is a simple PHP payload, designed to open a connection to the specified IP and port. You can modify the payload if needed for specific testing purposes (backdoor.php file)

## 🧩 Contributing
We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch ( using <a href="https://medium.com/@abhay.pixolo/naming-conventions-for-git-branches-a-cheatsheet-8549feca2534">this</a> convention).
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

## ⚖ License
This project is licensed under the GPL-3.0 License. See the LICENSE file for details.

## ⚔ Contact
- For any inquiries or support, please contact <a href="mailto:iannacconegiovanni444@gmail.com"> iannacconegiovanni444@gmail.com </a>.
- Visit my site for more informations about me and my work <a href="https://giovanni-iannaccone.github.io" target=”_blank”> https://giovanni-iannaccone.github.io </a>

🐞 Happy Hacking ... 