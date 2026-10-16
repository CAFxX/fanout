package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int8, p1 uint) bool {
    v4 := []int{}
    var v5 uint64 = 442
    switch ("hello" + "") {
    case "\n\t":
        if (!(712 != 255)) {
            v6 := GenS1{F0: 4294967295, F1: 2147483648, F2: 818}
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            if (1 != (fi7 - (fi7 * (^(fi7))))) {
                v8 := GenS1{F0: v5, F1: 2147483648, F2: 2147483647}
                _ = v8
            }
            _ = fi7
        }
    }
    if (true && (!(32767 <= 2))) {
        v9 := GenS2{F0: "a", F1: 0, F2: -1}
        _ = len(v4)
        _ = v9
    }
    var v10 int64 = -556
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v5
    return (0.0 == ((((float64((0.1 + 1e10)) - (((3.14159 + (0.0 / 0.1)) + -2.25) - 3.14159)) - -2.25) - 0.1) / (0.1 / (1e10 + ((3.14159 - 0.1) - 1.5)))))
}

func main() {
    var v1 uint16 = 42
    v2 := []string{}
    var v3 int8 = 100
    v1 = (^(v1))
    fmt.Println(v1, v1)
    v1 = v1
    v11 := GenS2{F0: ((("hello" + "\n\t") + "foo bar") + ("" + "")), F1: uint32(v1), F2: (v3 + 0)}
    v12 := GenS3{F0: true}
    switch ("foo bar" + "") {
    case "foo bar":
        switch 9223372036854775807 {
        case 32768:
            v1 += (^(v1))
        case -2147483648:
            for fi13 := 0; fi13 < 6; fi13++ {
                v1 = v1
                v14 := GenS4{F0: 4294967295, F1: 2147483647, F2: 311}
                v3 = (v3 - (v3 ^ v3))
                _ = fi13
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                v3 += (v3 >> ((42) & 63))
                _ = fi15
            }
        }
    case "a":
        if false {
            for fi16 := 0; fi16 < 13; fi16++ {
                v1 = v1
                v3 -= 42
                _ = fi16
            }
            var v17 uint16 = (v1 - 32767)
            _ = v17
        }
    default:
        v18 := GenS5{F0: -1}
        _ = v18
    }
    var v19 int32 = -498
    v19 += 0
    for fi20 := 0; fi20 < 13; fi20++ {
        v3 = 0
        v3 = 100
        _ = fi20
    }
    switch -877 {
    case 128:
        var v21 string = "a"
        _ = v21
    case 832:
        v1 -= 0
    case -9223372036854775808:
        v3 += (^(v3))
    default:
        var v22 uint8 = 255
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
