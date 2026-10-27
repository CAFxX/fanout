package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 115
    var v3 int32 = 225
    switch 9223372036854775807 {
    case 32768:
        v4 := GenS1{F0: 0}
        _ = v4
    case -600:
        var v5 uint8 = 128
        _ = v5
    case 1:
        v3 += v3
    default:
        v6 := GenS1{F0: v3}
        _ = v6
    }
    switch ("\n\t" + (("foo bar" + "") + "")) {
    case "hello":
        var v7 uint16 = 65535
        _ = v7
    case "a":
        fmt.Println(v2)
    }
    v8 := GenS1{F0: -2147483648}
    var v9 float32 = 3.14159
    if true {
        fmt.Println(v2, v2)
    }
    if (!(0 > 32767)) {
        for fi10 := 0; fi10 < 13; fi10++ {
            switch ("hello" + "hello") {
            case "foo bar":
                v3 = -48
            case "a":
                v11 := GenS1{F0: v3}
                _ = v11
            default:
                _ = len(v1)
            }
            _ = fi10
        }
    } else {
        switch 700 {
        case -2147483648:
            var v12 uint64 = 891
            _ = v12
        }
    }
    var v13 int64 = 0
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            switch "a" {
            case "":
                v15 := GenS1{F0: (v3 << ((4294967295) & 63))}
                _ = v15
            case "hello":
                v9 = v9
            default:
                var v16 int32 = v3
                _ = v16
            }
            var v17 int8 = 1
            _ = fi14
            _ = v17
        }
        v1[127] = -2147483648
        var v18 int = -820
        _ = v18
    } else {
        var v19 string = ("" + "foo bar")
        _ = v19
    }
    v20 := GenS2{F0: v3, F1: 1}
    if (v2 != v2) {
        v21 := GenS3{F0: v2, F1: 186, F2: 42}
        _ = v21
    }
    v3 = (-(v3))
    for fi22 := 0; fi22 < 8; fi22++ {
        v2 += (v2 % ((v2) | 1))
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
