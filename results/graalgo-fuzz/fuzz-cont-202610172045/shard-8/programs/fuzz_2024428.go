package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint16, p2 uint16) int16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(p0, p0, fi6)
        _ = fi6
    }
    p1 -= p1
    if (p0 < 605) {
        fmt.Println(p1, p2)
        v7 := GenS1{F0: -70, F1: 42}
        _ = v7
    }
    switch "foo bar" {
    case "":
        var v8 int64 = 1
        _ = v8
    case "hello":
        if ((1e10 / -2.25) == (3.14159 + (3.14159 / (float64(0.1) * (float64((-2.25 * (1.5 - ((-2.25 / 0.1) * (((-2.25 / 1.5) - 0.0) / 1.0))))) + ((((float64(1e10) / 0.1) + -2.25) * 1.5) / 1.0)))))) {
            v9 := []int{993, 701}
            fmt.Println(p0, p0, p0)
            _ = v9
        } else {
            var v10 int16 = -32768
            _ = v10
        }
    default:
        v11 := []int{2147483647, -2147483648}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = 1e10
    var v3 int = 127
    var v4 int32 = 0
    var v5 int = (v3 << ((750) & 63))
    v12 := GenS1{F0: 0, F1: 1}
    switch (^(v3)) {
    case 127:
        v13 := GenS1{F0: 5, F1: 4294967295}
        _ = v13
    case 128:
        var v14 uint16 = 0
        _ = v14
    case -782:
        v2 += v2
    }
    v15 := []int{(-(v5)), v5}
    if true {
        var v16 int32 = (v4 << ((4294967295) & 63))
        v17 := GenS2{F0: v16, F1: 1, F2: 490}
        _ = v16
        _ = v17
    }
    switch v3 {
    case 0:
        fmt.Println(v3, v2, v5)
    }
    v1 = 687
    v15 = append(v15, v5)
    switch (v3 % (((v5 | -9223372036854775808)) | 1)) {
    case 1:
        if (!true) {
            fmt.Println(v1)
            for fi18 := 0; fi18 < 20; fi18++ {
                v4 = v4
                _ = fi18
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint16 = 255
                var v21 int16 = 461
                v22 := GenS2{F0: (v4 << ((273) & 63)), F1: 343, F2: 921}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
    case -9223372036854775808:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    if (1 == 2) {
        if (((v4 % (((^(v4))) | 1)) > (v4 >> ((65535) & 63))) || true) {
            var v23 int = v3
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v24 := GenS2{F0: v4, F1: 2, F2: 4294967295}
            _ = v23
            _ = v24
        } else {
            if false {
                var v25 int64 = 0
                _ = v25
            } else {
                v26 := GenS2{F0: (-(v4)), F1: 65535, F2: 42}
                _ = v26
            }
        }
        var v27 uint8 = 2
        _ = v27
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint16 = 42
            fmt.Println(v29, v1)
            _ = fi28
            _ = v29
        }
        var v30 float32 = 1.5
        switch v5 {
        case -762:
            v31 := GenS3{F0: -128, F1: v4, F2: 128}
            _ = v31
        case 941:
            if false {
                fmt.Println(v3, v30, v1)
            } else {
                v32 := GenS4{F0: 65535, F1: 255}
                if (false || false) {
                    var v33 int8 = -1
                    _ = v33
                } else {
                    var v34 int8 = -128
                    v30 = (v30 + v30)
                    _ = v34
                }
                _ = v32
            }
        case 4294967295:
            v35 := GenS5{F0: 0, F1: 1}
            _ = v35
        default:
            var v36 int32 = (v4 | v4)
            _ = v36
        }
        _ = v30
    }
    var v37 float64 = v2
    v15 = append(v15, -1)
    v15 = append(v15, (-(v5)))
    switch -9223372036854775808 {
    case -973:
        switch -377 {
        case 127:
            var v38 bool = (18446744073709551615 != 0)
            _ = v38
        case -530:
            v15 = append(v15, -1)
        default:
            v37 = 1.5
        }
    case 826:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    case 792:
        if (42 == 18446744073709551615) {
            fmt.Println(v37, v1)
            v39 := GenS4{F0: uint16(v1), F1: 65535}
            var v40 uint = 18446744073709551615
            _ = v39
            _ = v40
        }
    }
    if true {
        v3 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
