package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
func genFn0() uint {
    var v7 float32 = float32((3.14159 * 3.14159))
    switch "hello" {
    case "a":
        v8 := GenS1{F0: 65535}
        _ = v8
    case "foo bar":
        v7 -= v7
    }
    v7 = -2.25
    if true {
        var v9 uint8 = 6
        for fi10 := 0; fi10 < 5; fi10++ {
            switch ("foo bar" + "foo bar") {
            case "":
                v11 := GenS2{F0: 960}
                _ = v11
            case "\n\t":
                var v12 float32 = v7
                _ = v12
            case "hello":
                v13 := []int{fi10, fi10}
                _ = v13
            default:
                var v14 int64 = 818
                _ = v14
            }
            _ = fi10
        }
        _ = v9
    }
    fmt.Println(v7)
    _ = v7
    return 2
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 1e10
    v3 := []int{0, 870, -592}
    var v4 int8 = 42
    var v5 string = ""
    var v6 uint16 = 32767
    v1 = v1
    v4 = v4
    if true {
        v1 = v5
    }
    if false {
        v15 := GenS1{F0: 255}
        v3 = append(v3, 335)
        v3 = append(v3, 4294967295)
        _ = v15
    }
    var v16 uint = 255
    if ((v2 * v2) != v2) {
        v17 := GenS2{F0: v6}
        _ = v17
    }
    if false {
        var v18 float32 = (1.5 + -2.25)
        v19 := GenS1{F0: v16}
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: (v16 | 2)}
        var v21 uint8 = 255
        fmt.Println(v6, v4)
        _ = v20
        _ = v21
    }
    var v22 uint = v16
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
