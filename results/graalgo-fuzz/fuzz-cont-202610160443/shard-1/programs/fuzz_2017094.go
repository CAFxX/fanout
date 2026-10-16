package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint) string {
    var v7 bool = (!false)
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return (("" + "hello") + "")
}

func main() {
    var v1 bool = false
    v2 := []string{"a", "", "\n\t"}
    var v3 int8 = 100
    var v4 bool = true
    var v5 float64 = 3.14159
    v6 := []string{"\n\t", "a", "foo bar"}
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: -9223372036854775808}
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int64 = -1
            fmt.Println(v1, v3)
            var v12 int = fi8
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if (-1 >= 9223372036854775807) {
        v1 = (2147483647 >= -146)
        v13 := GenS1{F0: -9223372036854775808}
        _ = v13
    }
    switch "" {
    case "a":
        var v14 int = 102
        _ = v14
    }
    _ = len(v6)
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    v1 = v4
    if len(v6) > 0 {
        v6[0] = (("" + "hello") + ("hello" + ("hello" + ("hello" + (((("foo bar" + ("\n\t" + "")) + "") + "") + "hello")))))
    }
    fmt.Println(v3, v5)
    v5 = 1e10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
