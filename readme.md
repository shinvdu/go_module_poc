## An example to explain how to use go module.

pay attention to 2 points:

#### 1. How to import own source package

import own package with module as prefix. 

example: 

- wrong: 
```
import (
	"fmt"
	"./config"
)
```

- correct: 
```
import (
	"fmt"
	"github.com/shinvdu/go_module_poc/config"
)

```

<github.com/shinvdu/go_module_poc> is the module name, also it can be a string without github.com as prefix. 

this is specified by go mod init 

$: go mod init github.com/shinvdu/go_module_poc

#### 2. A folder is a package, and it can have several files in it with same package name

in this example, a package config have two file

config/
├── bao.go
└── config.go


after I imported as below
```
import (
	"fmt"
	"github.com/shinvdu/go_module_poc/config"
)
```

I can call both defined struct in two file.

like

```
a := config.GetConfig()
c := config.GetDaoConfig()
var b config.DaoConfig
b.Username = "BCD"
```


