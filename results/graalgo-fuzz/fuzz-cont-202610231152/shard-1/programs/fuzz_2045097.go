package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int16 = 42
    v2 := []string{""}
    v3 := make(map[uint32]bool)
    var v4 bool = (2 < 255)
    var v5 int16 = (v1 + v1)
    var v6 int64 = 203
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v5)
        v4 = v4
        delete(v3, 0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            var v9 uint8 = 42
            _ = v9
        } else {
            var v10 uint16 = 0
            v10 -= v10
            _ = v10
        }
        v4 = ((v1 ^ int16(fi8)) < v1)
        _ = fi8
    }
    v4 = (int(v1) < 4294967295)
    if v4 {
        var v11 string = "\n\t"
        var v12 int32 = 649
        switch "a" {
        case "a":
            var v13 float32 = (float32(((-2.25 / 1.5) - (3.14159 / ((0.0 + (3.14159 - ((0.0 + (0.0 + (-2.25 / -2.25))) - 0.1))) * -2.25)))) + float32((1.5 / ((-2.25 / (0.0 - (1e10 * 3.14159))) / 1e10))))
            _ = v13
        default:
            v4 = (((3.14159 + (-2.25 * 0.1)) > ((0.0 + (1.5 * 3.14159)) * 0.0)) || (1e10 >= 0.1))
        }
        _ = v11
        _ = v12
    } else {
        if (0 >= v6) {
            delete(v3, 135)
            if (22 > 2147483648) {
                v14 := GenS1{F0: (v5 / (((v1 / ((v1) | 1))) | 1))}
                _ = v14
            } else {
                switch ((("\n\t" + (("foo bar" + ("foo bar" + ("hello" + "hello"))) + "a")) + "\n\t") + "\n\t") {
                case "hello":
                    v15 := GenS1{F0: int16(v6)}
                    _ = v15
                case "\n\t":
                    var v16 uint64 = 2
                    _ = v16
                }
                v3[42] = ((46 == 32767) || (-784 != -580))
                v17 := GenS1{F0: v5}
                _ = v17
            }
            var v18 float32 = float32(0.1)
            _ = v18
        } else {
            fmt.Println(v5)
        }
    }
    var v19 int8 = 1
    _ = len(v2)
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v4)
        _ = fi20
    }
    var v21 bool = (v6 != v6)
    var v22 int16 = (v5 - v5)
    fmt.Println(v4, v5, v6)
    v2 = append(v2, "a")
    var v23 int16 = (v22 / (((v22 | v1)) | 1))
    fmt.Println(v4, v21, v6)
    var v24 uint8 = 2
    v25 := GenS2{F0: (v24 % (((^(v24))) | 1)), F1: 455}
    v1 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
