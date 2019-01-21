# Cell Definition Language (CDL)

CDL is a definition language that makes it easy define deployments for a Cell based architecture without considering the runtime.

### Download and Install

#### CDL compiler

Download the binary distribution for Linux and Mac from the release page at https://github.com/Mirage20/cdl/releases

#### VSCode plugin (Optional)

To install the plugin, copy the content inside the [vscode-plugin](/vscode-plugin) directory to `~/.vscode/extensions`

    cp -r ./vscode-plugin/* ~/.vscode/extensions/
    
    
 ![VSCode Sample](/vscode-plugin/mirage20.cdl-0.1.0/vs-code.png)
 
### Running a Sample

Run following commands to download three sample cells.

    curl https://raw.githubusercontent.com/Mirage20/cdl/master/sample/employee.cdl --output employee.cdl
    curl https://raw.githubusercontent.com/Mirage20/cdl/master/sample/hr.cdl --output hr.cdl
    curl https://raw.githubusercontent.com/Mirage20/cdl/master/sample/stock-options.cdl --output stock-options.cdl
    
Run following command to compile the three cells files and deploy it to Kubernetes (requires cell controller)

    cdl hr.cdl employee.cdl stock-options.cdl --runtime=kubernetes | kubectl apply -f -
    
Or to save as a readable format

    cdl hr.cdl employee.cdl stock-options.cdl --runtime=kubernetes -o yaml
    
### CDL Syntax from a Sample

```
# This is a comment

Cell employee {
    Component employee {
        Image:Docker = "docker.io/wso2vick/sampleapp-employee"
        Ports {
            TCP 80->8080
            TCP 443->8081
        }
    }

    Component salary {
        Image:Docker = "docker.io/wso2vick/sampleapp-salary"
        Ports {
             TCP 80->8080
        }
    }

    Ingress:HTTP {
        "/employee" -> employee {
            GET "/"
            POST "/"
        }
    }
}
```
