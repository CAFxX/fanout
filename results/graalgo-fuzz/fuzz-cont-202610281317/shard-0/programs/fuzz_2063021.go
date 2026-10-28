package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 uint }
func genFn0() int {
    var v4 float32 = 0.0
    v5 := GenS1{F0: 42}
    switch "foo bar" {
    case "":
        var v6 bool = false
        _ = v6
    case "hello":
        var v7 string = "\n\t"
        _ = v7
    case "\n\t":
        var v8 uint8 = 127
        _ = v8
    }
    v9 := GenS1{F0: 4294967295}
    _ = v4
    _ = v5
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"foo bar", "a", (((("a" + "") + "foo bar") + ("\n\t" + ("a" + ""))) + "hello"), "a"}
    v3 := []string{"foo bar", "\n\t"}
    v3 = append(v3, (("\n\t" + ("a" + "a")) + "foo bar"))
    var v10 uint32 = 65535
    v10 = (v10 - (v10 + v10))
    if true {
        v11 := GenS2{F0: (v10 ^ (v10 * (v10 % ((v10) | 1)))), F1: v10}
        var v12 uint32 = (v10 & (v10 | (v10 * v10)))
        switch 4294967295 {
        case 32767:
            v13 := GenS1{F0: 0}
            _ = v13
        case -125:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v10 = 4294967295
        }
        _ = v11
        _ = v12
    }
    v14 := GenS3{F0: v10}
    fmt.Println(v10)
    if (255 <= 255) {
        switch -1 {
        case 2147483647:
            fmt.Println(v10)
        default:
            v1[4294967295] = 4294967295
        }
    } else {
        var v15 string = ("" + (("a" + "") + "a"))
        var v16 int64 = 9223372036854775807
        for fi17 := 0; fi17 < 17; fi17++ {
            _ = len(v1)
            v18 := GenS4{F0: ((4294967295 >= 4294967295) || (2 == 4294967295)), F1: (0 >= 0), F2: 736}
            v15 = (v15 + (((((v15 + (((v15 + v15) + (v15 + "a")) + v15)) + ((v15 + "hello") + v15)) + v15) + v15) + ""))
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    if false {
        for fi19 := 0; fi19 < 5; fi19++ {
            _ = len(v1)
            var v20 float32 = 1.5
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS5{F0: -845, F1: 928, F2: 4294967295}
    v10 += v10
    var v22 uint64 = 42
    fmt.Println(v22, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
