package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0() uint64 {
    if false {
        v5 := GenS1{F0: 123, F1: (("hello" + "\n\t") + "hello")}
        var v6 float64 = 3.14159
        v6 -= ((((v6 - (1e10 - ((v6 / v6) - v6))) / (v6 + ((((0.0 + float64(1.5)) - (v6 / float64(1e10))) + (float64(0.1) - (((v6 + v6) * v6) / (float64((1.5 / 1e10)) + 3.14159)))) * 1.5))) + v6) + ((v6 * -2.25) - v6))
        _ = v5
        _ = v6
    }
    var v7 bool = ((1.5 - (float32(1e10) / (1.5 / (-2.25 * 0.1)))) != (1e10 - (0.0 + 0.0)))
    v8 := []int{589, -2147483648, -765}
    if (32767 != 32767) {
        var v9 uint = 1
        _ = v9
    }
    var v10 float32 = float32(1e10)
    _ = v10
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []string{}
    var v2 float64 = (1e10 - 0.1)
    v3 := []uint32{1, 0, 42, 2147483648}
    var v4 int8 = 42
    var v11 uint64 = 4294967295
    if (true || true) {
        var v12 uint64 = 1
        var v13 uint64 = genFn0()
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: ""}
    switch 32768 {
    case 358:
        v11 += (v11 / (((^(v11))) | 1))
    case -1:
        var v15 uint32 = 231
        _ = v15
    case -9223372036854775808:
        v3 = append(v3, 1)
    }
    switch "foo bar" {
    case "":
        v11 -= v11
    default:
        v4 -= v4
    }
    if (128 < 1) {
        if true {
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 = append(v1, "hello")
                _ = fi16
            }
        } else {
            v4 = v4
        }
    }
    var v17 float64 = 0.1
    v3 = append(v3, 931)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if true {
        if true {
            v18 := GenS2{F0: -9223372036854775808, F1: 2, F2: 255}
            _ = v18
        }
        var v19 uint32 = 1
        _ = v19
    }
    v17 -= 3.14159
    v17 += ((0.1 / float64((3.14159 - 1.5))) - float64(0.1))
    var v20 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
