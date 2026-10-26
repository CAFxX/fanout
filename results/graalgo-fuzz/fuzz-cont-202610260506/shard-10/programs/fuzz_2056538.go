package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 string, p1 uint16, p2 uint8) int32 {
    switch "" {
    case "foo bar":
        fmt.Println(p0)
    case "":
        var v4 int64 = 0
        _ = v4
    default:
        v5 := []int{}
        _ = v5
    }
    var v6 int8 = -125
    var v7 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 uint8, p2 int16) uint {
    fmt.Println(p2)
    var v8 int64 = 42
    switch -586 {
    case 32768:
        v9 := []int{9223372036854775807}
        _ = v9
    case 127:
        v10 := []int{}
        _ = v10
    }
    if ((425 == 4294967295) || (!((false || (!((v8 <= (v8 << ((255) & 63))) && false))) && true))) {
        for fi11 := 0; fi11 < 12; fi11++ {
            p0 = ("foo bar" + "hello")
            p0 = ("" + ((p0 + p0) + p0))
            _ = fi11
        }
    }
    if (450 == 684) {
        p2 = -32768
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 83
}

func main() {
    var v1 int32 = -846
    var v2 uint32 = 0
    var v3 uint32 = 4294967295
    var v12 int = -382
    v3 -= (^(v2))
    v3 = (v2 | (v2 % ((v2) | 1)))
    var v13 uint = 65535
    if true {
        v2 = (v3 & (v3 / (((^(v3))) | 1)))
        var v14 uint64 = 42
        v15 := []int{v12}
        _ = v14
        _ = v15
    }
    var v16 bool = (0 != 4294967295)
    fmt.Println(v2, v1, v12)
    fmt.Println(v12, v3, v2)
    fmt.Println(v13)
    switch v12 {
    case 128:
        var v17 bool = (-758 == -1)
        _ = v17
    case 127:
        fmt.Println(v3, v13, v16)
    }
    if v16 {
        for fi18 := 0; fi18 < 19; fi18++ {
            if ((!((v1 * (v1 ^ v1)) > v1)) || v16) {
                if ((!v16) || true) {
                    v3 -= (v3 & (v2 / (((v3 & 42)) | 1)))
                    v1 += genFn0("\n\t", 653, 197)
                }
            } else {
                var v19 bool = v16
                switch "" {
                case "hello":
                    v13 += (^(v13))
                case "":
                    var v20 float32 = (-2.25 + 0.0)
                    _ = v20
                }
                _ = v19
            }
            _ = fi18
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 float64 = ((1.5 - -2.25) - (1e10 / (1e10 * ((1.5 / float64(0.1)) / ((3.14159 / 1.0) / 0.1)))))
            _ = fi21
            _ = v22
        }
    } else {
        v13 = (v13 >> ((397) & 63))
        var v23 uint = (^(v13))
        v24 := GenS1{F0: 2, F1: 1000}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := []int{fi25}
        v2 = v2
        var v27 uint = (v13 << ((v13) & 63))
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int8 = -128
    var v29 float32 = -2.25
    if (68 < v28) {
        if v16 {
            v30 := GenS1{F0: (v2 - (v2 % ((4294967295) | 1))), F1: 32767}
            v16 = v16
            v31 := GenS1{F0: v2, F1: 287}
            _ = v30
            _ = v31
        } else {
            v32 := GenS1{F0: 65535, F1: 42}
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 5; fi33++ {
            fmt.Println(v12, v12, v16)
            _ = fi33
        }
        v34 := GenS2{F0: v16, F1: ""}
        _ = v34
    }
    v12 = -2147483648
    v35 := GenS2{F0: (v12 < (v12 >> ((v13) & 63))), F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
}
