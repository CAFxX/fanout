package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0() bool {
    var v5 uint64 = 18446744073709551615
    var v6 string = (("foo bar" + ("foo bar" + "")) + ("hello" + "a"))
    _ = v5
    _ = v6
    return (-451 > -2147483648)
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    var v3 int8 = -128
    var v4 uint64 = 1
    var v7 int8 = (v3 % ((100) | 1))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    switch ("foo bar" + ("a" + "a")) {
    case "\n\t":
        v7 += (v7 & 127)
    }
    var v10 uint32 = 2147483648
    fmt.Println(v4, v2)
    var v11 float64 = 1.5
    for fi12 := 0; fi12 < 17; fi12++ {
        v3 -= v7
        _ = fi12
    }
    delete(v1, "hello")
    v4 = v4
    fmt.Println(v7, v7, v10)
    if false {
        fmt.Println(v2)
        v3 = (v3 - -128)
    } else {
        v10 = (^(v10))
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 string = "hello"
            v2 = v2
            var v15 int8 = (v7 / (((v3 & -1)) | 1))
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    switch 42 {
    case 128:
        v16 := GenS1{F0: 128, F1: (v2 & 109)}
        _ = v16
    default:
        var v17 uint32 = v10
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
