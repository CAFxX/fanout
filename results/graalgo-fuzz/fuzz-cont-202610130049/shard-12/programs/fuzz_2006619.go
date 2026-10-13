package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8) int8 {
    v7 := []int{(p0 / (((p0 * 0)) | 1)), p0, 42, p0}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int32 = -1
        var v10 float64 = float64(1e10)
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (p0 << ((508) & 63)) {
    case -525:
        if (!true) {
            fmt.Println(p1, p1)
            fmt.Println(p0, p1)
        } else {
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(fi11)
        var v12 uint8 = (p1 + 137)
        var v13 float64 = float64(-2.25)
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("foo bar" + ""), ("hello" + "foo bar"), "\n\t"}
    var v3 int32 = 562
    v4 := []int{0, 42, -9223372036854775808, -2147483648}
    v5 := []uint32{2147483648, 4294967295, 4294967295}
    v6 := []int{128}
    var v14 int16 = 1
    var v15 uint32 = 0
    v15 = 1
    var v16 int8 = 100
    fmt.Println(v1, v14, v15)
    v1 = (^(v1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v14)
        v18 := GenS1{F0: v15}
        var v19 uint16 = 255
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v14, v3, v16)
    if true {
        switch 0 {
        case 32768:
            var v20 float64 = 1.5
            _ = v20
        case 908:
            var v21 uint32 = 859
            _ = v21
        default:
            v3 = (v3 & (v3 ^ v3))
        }
        var v22 uint32 = 0
        _ = v22
    }
    var v23 uint16 = 802
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := GenS1{F0: 2}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
