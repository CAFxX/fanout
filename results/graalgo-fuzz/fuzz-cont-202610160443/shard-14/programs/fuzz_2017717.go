package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int16; F1 uint }
func genFn0(p0 int32, p1 uint, p2 uint) int16 {
    switch 42 {
    case -9223372036854775808:
        var v4 bool = (1 >= 2)
        _ = v4
    case 840:
        p0 -= p0
    }
    var v5 float64 = 0.0
    if (float32(float64(((1.5 - float32(((3.14159 / v5) * v5))) + (0.0 / 1.5)))) == 1e10) {
        v6 := GenS1{F0: 1}
        _ = v6
    }
    if true {
        switch "\n\t" {
        case "a":
            var v7 string = ("hello" + "a")
            _ = v7
        case "":
            var v8 uint8 = 255
            _ = v8
        case "hello":
            var v9 float32 = 3.14159
            _ = v9
        default:
            switch "a" {
            case "\n\t":
                fmt.Println(v5, p2, v5)
            default:
                switch 128 {
                case -275:
                    p0 += p0
                }
            }
        }
        if (!false) {
            if (18446744073709551615 != 1) {
                if (2147483647 >= 65535) {
                    if false {
                        p2 = (p2 >> ((p2) & 63))
                        p1 -= p1
                    } else {
                        fmt.Println(p1, v5)
                    }
                    v10 := GenS2{F0: (("" + "hello") + ("" + "foo bar"))}
                    _ = v10
                }
                var v11 int64 = 1
                v12 := GenS2{F0: ""}
                _ = v11
                _ = v12
            } else {
                var v13 int = -1
                _ = v13
            }
            var v14 float64 = 3.14159
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS2{F0: "foo bar"}
            v17 := GenS1{F0: 16}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 411
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 bool = (!false)
    v1 -= v1
    var v18 int16 = -1
    var v19 uint8 = 127
    switch "" {
    case "\n\t":
        var v20 int32 = v1
        _ = v20
    case "a":
        var v21 string = ""
        _ = v21
    }
    v22 := GenS3{F0: 42, F1: 980}
    if false {
        var v23 float64 = (0.0 + 0.1)
        var v24 uint64 = 941
        var v25 int = 127
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 int = 781
    v3 = ("" == ((("hello" + ("hello" + "a")) + ("\n\t" + ("\n\t" + "foo bar"))) + (("" + ("hello" + "\n\t")) + "\n\t")))
    var v27 uint8 = v19
    var v28 uint32 = 1
    var v29 int = (v26 ^ v26)
    v3 = true
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS4{F0: (v18 / (((v18 + (^(v18)))) | 1)), F1: 332}
        _ = fi30
        _ = v31
    }
    var v32 uint32 = (v28 | 4294967295)
    for fi33 := 0; fi33 < 10; fi33++ {
        v27 += (v19 >> ((255) & 63))
        var v34 int = (-(fi33))
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
