package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 string, p1 uint8, p2 uint) int8 {
    var v6 uint64 = 4294967295
    p2 -= (p2 >> ((330) & 63))
    var v7 uint64 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 string = "\n\t"
    var v3 uint16 = 65535
    var v4 string = ("foo bar" + v2)
    v5 := make(map[string]string)
    if false {
        switch "\n\t" {
        case "hello":
            v1 = (v1 - 44)
        case "":
            v3 -= (v3 ^ v3)
        default:
            v8 := GenS1{F0: v2, F1: -922, F2: (v1 & (v1 + (v1 << ((1) & 63))))}
            _ = v8
        }
    }
    var v9 int8 = -128
    v10 := GenS1{F0: "a", F1: 42, F2: v1}
    fmt.Println(v3, v9, v1)
    if ((v3 << ((573) & 63)) > v3) {
        v5[v2] = (("a" + "\n\t") + "a")
        var v11 uint64 = 42
        _ = v11
    } else {
        var v12 string = v4
        fmt.Println(v3, v9, v4)
        v13 := GenS1{F0: "hello", F1: 32767, F2: (^(v1))}
        _ = v12
        _ = v13
    }
    var v14 uint64 = 0
    v1 = v1
    v2 = (("foo bar" + ((v4 + v2) + (v4 + (v4 + (v4 + v2))))) + "a")
    var v15 uint = 451
    var v16 int32 = -2147483648
    var v17 int = 32768
    v18 := GenS1{F0: v4, F1: 909, F2: 2147483647}
    var v19 int = v17
    v20 := []int{v17, v19, int(v9)}
    v15 = (v15 & v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
