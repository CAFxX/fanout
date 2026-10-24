package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint16) uint32 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = 42
        fmt.Println(p0)
        var v9 int64 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    return 42
}
func genFn1(p0 uint64) uint8 {
    var v10 float32 = (1.5 / 3.14159)
    v11 := []int{32767, 38}
    if (!false) {
        if true {
            if ("a" == "hello") {
                v10 = v10
            }
        }
        if (-9223372036854775808 >= -824) {
            var v12 int32 = 42
            if true {
                v10 = v10
            }
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 194, F1: -74}
        p0 = p0
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = 644
        var v16 int32 = 2147483647
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := []uint32{2147483647, 1}
    var v2 float64 = 0.1
    var v3 float32 = 0.0
    var v4 int16 = -1
    v5 := []int{}
    v6 := make(map[string]string)
    v4 = (v4 << ((42) & 63))
    var v17 int16 = v4
    switch "\n\t" {
    case "hello":
        delete(v6, ((("a" + ("a" + "a")) + (("" + "hello") + ((("a" + "hello") + ("foo bar" + (((("hello" + "hello") + "hello") + ("hello" + "foo bar")) + "hello"))) + "\n\t"))) + ""))
    case "\n\t":
        v6[""] = "foo bar"
    }
    switch 688 {
    case 32767:
        v3 += (0.0 * 1e10)
    case -918:
        fmt.Println(v4, v2, v17)
    default:
        var v18 uint8 = 128
        _ = v18
    }
    var v19 float32 = 1e10
    v17 = v17
    v2 -= v2
    for fi20 := 0; fi20 < 18; fi20++ {
        if true {
            v3 -= v19
            v17 = (v17 << ((255) & 63))
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(v17)
                fmt.Println(fi21, fi20, fi20)
                v22 := GenS1{F0: 65535, F1: -1}
                _ = fi21
                _ = v22
            }
        }
        v23 := GenS1{F0: 2147483647, F1: 864}
        if len(v1) > 0 {
            v1[0] = 990
        }
        _ = fi20
        _ = v23
    }
    v4 -= (v17 ^ v4)
    var v24 int8 = int8(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
