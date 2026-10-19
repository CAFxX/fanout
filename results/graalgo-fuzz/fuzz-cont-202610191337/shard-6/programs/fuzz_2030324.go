package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{2, 825, 1}
    var v3 float32 = 3.14159
    v4 := []int{}
    var v5 int8 = -128
    switch ("foo bar" + (("hello" + ("hello" + "\n\t")) + "")) {
    case "\n\t":
        for fi6 := 0; fi6 < 10; fi6++ {
            v1 -= (v1 ^ 745)
            _ = fi6
        }
    case "hello":
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int8 = -63
            switch ("\n\t" + "a") {
            case "hello":
                var v9 uint8 = 64
                _ = v9
            case "foo bar":
                fmt.Println(fi7, fi7)
            case "\n\t":
                var v10 int32 = 2147483647
                _ = v10
            default:
                var v11 uint8 = 1
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    case "a":
        v1 = v1
    default:
        var v12 uint16 = 65535
        _ = v12
    }
    var v13 uint8 = 0
    v14 := GenS1{F0: (v1 - v1), F1: "", F2: 259}
    v13 = v13
    for fi15 := 0; fi15 < 12; fi15++ {
        if len(v2) > 0 {
            v2[0] = 550
        }
        _ = fi15
    }
    v13 = 215
    if len(v4) > 0 {
        v4[0] = -364
    }
    _ = len(v4)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int8 = 127
        var v18 uint32 = 2147483647
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int16 = 1000
        _ = fi19
        _ = v20
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        v3 -= 3.14159
    case -775:
        if true {
            var v21 uint32 = 65535
            fmt.Println(v21)
            _ = v21
        } else {
            var v22 float64 = float64((((v3 / v3) - ((-2.25 / v3) * v3)) + 1e10))
            var v23 uint = 255
            _ = v22
            _ = v23
        }
    case 32768:
        var v24 int = 127
        _ = v24
    default:
        var v25 int32 = 1
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS1{F0: (^(v1)), F1: ("" + ""), F2: 1}
        _ = fi26
        _ = v27
    }
    var v28 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
