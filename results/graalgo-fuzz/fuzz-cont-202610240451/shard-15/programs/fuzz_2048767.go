package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 bool) string {
    var v7 int64 = 715
    v7 += v7
    fmt.Println(v7, v7)
    p0 = p0
    _ = p0
    _ = v7
    return ""
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 841
    var v3 int = 32768
    var v4 int64 = 9223372036854775807
    var v5 float64 = 1e10
    var v6 uint = (v1 & 413)
    switch ("\n\t" + "hello") {
    case "\n\t":
        if true {
            var v8 int32 = 1
            fmt.Println(v8, v5)
            _ = v8
        } else {
            var v9 int8 = 42
            var v10 float32 = -2.25
            v11 := GenS1{F0: 32767, F1: v9}
            _ = v10
            _ = v11
            _ = v9
        }
    case "hello":
        var v12 int64 = 0
        _ = v12
    case "":
        var v13 uint16 = 1
        _ = v13
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := []int{}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    var v17 uint8 = 128
    v1 = (v6 << (((v6 / ((v1) | 1))) & 63))
    if true {
        var v18 int32 = -1
        v19 := []int{v3}
        v20 := GenS1{F0: 32767, F1: 100}
        _ = v18
        _ = v19
        _ = v20
    } else {
        if true {
            v4 += v4
            v3 += -9223372036854775808
        }
        var v21 string = ("" + ("" + "a"))
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            switch "" {
            case "\n\t":
                v6 = v6
            case "":
                var v24 float64 = (v5 + v5)
                v24 += -2.25
                _ = v24
            case "foo bar":
                v25 := GenS1{F0: 729, F1: 100}
                _ = v25
            default:
                v26 := []int{(v3 - -1), (fi22 * (-(fi22))), v3}
                _ = v26
            }
            _ = fi23
        }
        v4 += 9223372036854775807
        v27 := []int{v3, (v3 ^ v3)}
        _ = fi22
        _ = v27
    }
    switch v3 {
    case 0:
        switch "a" {
        case "a":
            v28 := []int{v3, (v3 & 9223372036854775807)}
            _ = v28
        default:
            v29 := GenS1{F0: 716, F1: 42}
            _ = v29
        }
    case 127:
        v6 = uint(v17)
    case 292:
        v17 += 128
    }
    v30 := GenS1{F0: 42, F1: -97}
    fmt.Println(v4, v5, v4)
    var v31 int16 = -32768
    fmt.Println(v4, v2)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
