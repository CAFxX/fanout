package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint8 = 128
        if (!((false && true) && true)) {
            v8 -= v8
        }
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := []int{fi9, fi9, -9223372036854775808, fi9}
        var v11 uint8 = 144
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    var v13 int32 = 2147483647
    _ = v13
    return 255
}
func genFn1(p0 uint8) uint64 {
    p0 = p0
    v14 := GenS1{F0: 2, F1: 31, F2: 1000}
    p0 = p0
    for fi15 := 0; fi15 < 18; fi15++ {
        p0 = p0
        v16 := []int{fi15}
        fmt.Println(p0, fi15)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := []int{fi17}
        p0 = (p0 + p0)
        switch "a" {
        case "foo bar":
            switch 32768 {
            case 1:
                v19 := GenS2{F0: 2}
                _ = v19
            default:
                v20 := GenS2{F0: 0}
                _ = v20
            }
        default:
            v21 := GenS3{F0: 4294967295}
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 42
    var v3 uint16 = 2
    var v4 string = (("" + "a") + "")
    var v5 uint = 183
    var v6 uint64 = 351
    v6 += (v6 >> ((v5) & 63))
    fmt.Println(v4)
    var v22 float64 = 1e10
    switch 32767 {
    case -166:
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v22)
            v22 -= float64(1e10)
            _ = fi23
        }
    case -791:
        for fi24 := 0; fi24 < 18; fi24++ {
            switch v4 {
            case "":
                v25 := []int{fi24, 32768, (fi24 << ((v5) & 63)), fi24}
                _ = v25
            case "foo bar":
                if (v6 == v6) {
                    var v26 int16 = 32767
                    fmt.Println(v26, v2)
                    _ = v26
                }
            default:
                v27 := []int{(^(fi24))}
                _ = v27
            }
            v28 := []int{fi24, 127}
            switch fi24 {
            case 0:
                v28 = append(v28, (fi24 ^ (fi24 * fi24)))
            case 469:
                if (v22 > 1.5) {
                    var v29 int16 = 472
                    _ = v29
                }
            default:
                v1 = (v1 - ((float64(float32(v1)) / v1) + 0.0))
            }
            _ = fi24
            _ = v28
        }
    case 127:
        switch 2147483647 {
        case 0:
            var v30 uint16 = genFn0()
            _ = v30
        case -9223372036854775808:
            for fi31 := 0; fi31 < 1; fi31++ {
                for fi32 := 0; fi32 < 6; fi32++ {
                    v33 := []int{}
                    _ = fi32
                    _ = v33
                }
                _ = fi31
            }
        case -1:
            fmt.Println(v6)
        default:
            var v34 uint64 = genFn1(1)
            _ = v34
        }
    default:
        fmt.Println(v4)
    }
    if false {
        fmt.Println(v2, v6)
        fmt.Println(v6, v4, v5)
        v3 = 255
    }
    v35 := []int{}
    if ((true || (-32768 != 90)) || true) {
        var v36 string = (v4 + v4)
        var v37 string = (v36 + (v4 + v4))
        v38 := GenS4{F0: ((("\n\t" + v37) + ("foo bar" + "")) + v37)}
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 uint16 = v3
    v40 := GenS5{F0: 288}
    v41 := GenS3{F0: (v6 & (^(v6)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
