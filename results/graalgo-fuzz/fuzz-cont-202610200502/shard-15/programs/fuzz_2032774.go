package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0(p0 int32, p1 uint32) int16 {
    var v4 uint = 2
    v5 := []int{-9, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []int{-2147483648, 908}
    var v2 string = (("a" + ("" + ("foo bar" + ("hello" + "")))) + "")
    var v3 uint64 = 42
    v2 = "a"
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint8 = 128
        v3 -= v3
        _ = fi7
        _ = v8
    }
    v6 = 834
    if true {
        fmt.Println(v3)
    } else {
        v9 := GenS1{F0: 18446744073709551615, F1: 311, F2: 1}
        _ = v9
    }
    if true {
        switch 127 {
        case 0:
            switch "" {
            case "foo bar":
                var v10 bool = (false && (-1 == -1))
                _ = v10
            }
        case 4294967295:
            var v11 uint64 = v3
            _ = v11
        default:
            if false {
                v1 = append(v1, 1)
            }
        }
    }
    var v12 float32 = 0.1
    v12 += (v12 - float32(1.5))
    v2 = "hello"
    v3 += v3
    var v13 float32 = v12
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: (v6 * 2), F1: 65535, F2: -881}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: v6, F1: 721, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
