package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 uint32, p1 int8) uint16 {
    p1 = 127
    for fi5 := 0; fi5 < 12; fi5++ {
        if (!true) {
            switch ("foo bar" + "\n\t") {
            case "":
                fmt.Println(p0, fi5, p1)
            case "foo bar":
                for fi6 := 0; fi6 < 17; fi6++ {
                    var v7 bool = true
                    _ = fi6
                    _ = v7
                }
            }
        }
        v8 := GenS1{F0: -628, F1: (fi5 | -343), F2: 2147483647}
        var v9 int32 = -845
        _ = fi5
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 890
}

func main() {
    var v1 uint = 42
    var v2 int = 127
    var v3 int8 = 102
    v4 := make(map[string]string)
    fmt.Println(v3, v2, v2)
    var v10 uint16 = 36
    v11 := GenS1{F0: 0, F1: (v2 >> ((v1) & 63)), F2: 2147483647}
    var v12 int = (v2 * (v2 * (v2 >> ((v1) & 63))))
    v2 = 42
    var v13 int8 = -90
    v3 = (v13 & (^(v3)))
    var v14 int8 = (^(v3))
    var v15 int16 = 32767
    v2 = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
