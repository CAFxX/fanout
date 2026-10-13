package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
func genFn0(p0 int8) uint8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 = 117
        for fi6 := 0; fi6 < 19; fi6++ {
            p0 -= (-(p0))
            fmt.Println(p0, fi5, p0)
            _ = fi6
        }
        _ = fi5
    }
    fmt.Println(p0)
    var v7 int8 = p0
    switch 457 {
    case -813:
        if ((!true) && false) {
            var v8 float32 = (float32(-2.25) - (-2.25 * (0.1 * (float32((1e10 - float64(3.14159))) * float32(0.0)))))
            var v9 uint16 = 1
            p0 = v7
            _ = v8
            _ = v9
        } else {
            switch "foo bar" {
            case "hello":
                v10 := []int{9223372036854775807, 693, int(p0), 32767}
                _ = v10
            case "\n\t":
                p0 -= (p0 << ((871) & 63))
            }
            v7 = v7
        }
    case 127:
        var v11 float32 = 1e10
        _ = v11
    case 996:
        v12 := GenS1{F0: 1}
        _ = v12
    default:
        p0 = (p0 & 1)
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: 2147483647}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            v7 = -128
            var v17 int = fi16
            _ = fi16
            _ = v17
        }
        if true {
            v7 += v7
            if (811 == 42) {
                v18 := GenS1{F0: 0}
                v19 := GenS2{F0: 4294967295, F1: (("foo bar" + ("a" + ((("" + "foo bar") + "\n\t") + "a"))) + ("\n\t" + "\n\t"))}
                _ = v18
                _ = v19
            } else {
                v7 += 100
                var v20 int = (^(fi13))
                fmt.Println(p0)
                _ = v20
            }
        }
        _ = fi13
    }
    _ = p0
    _ = v7
    return 54
}

func main() {
    var v1 int64 = 42
    var v2 int = 1
    var v3 int8 = 100
    var v4 float64 = 3.14159
    v21 := []int{}
    v2 += v2
    var v22 uint = 18446744073709551615
    if false {
        fmt.Println(v22, v3, v22)
        var v23 uint16 = 1
        _ = v23
    } else {
        var v24 int64 = (v1 ^ v1)
        _ = v24
    }
    var v25 int64 = int64(v3)
    var v26 string = "\n\t"
    fmt.Println(v4)
    v4 += (v4 / v4)
    v4 += -2.25
    if ((0 <= 4294967295) || true) {
        if true {
            var v27 int32 = -265
            switch v2 {
            case -314:
                v28 := GenS3{F0: v22, F1: true, F2: 2}
                _ = v28
            case -22:
                fmt.Println(v2)
            }
            _ = v27
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        var v29 string = v26
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
