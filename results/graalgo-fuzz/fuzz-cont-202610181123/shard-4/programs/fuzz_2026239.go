package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int8 }
func genFn0() string {
    v6 := GenS1{F0: 410}
    if false {
        var v7 int16 = 122
        v8 := []int{-770, 32767}
        _ = len(v8)
        _ = v7
        _ = v8
    }
    v9 := []int{0}
    switch 9223372036854775807 {
    case 128:
        v10 := GenS2{F0: 0}
        _ = v10
    }
    _ = v6
    _ = v9
    return "\n\t"
}
func genFn1(p0 int64, p1 uint8) uint {
    for fi11 := 0; fi11 < 7; fi11++ {
        p0 += p0
        _ = fi11
    }
    var v12 bool = (!("" > "foo bar"))
    var v13 int = -9223372036854775808
    var v14 int = 127
    v15 := []int{(v13 & (v13 + v14))}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 4294967295
}

func main() {
    var v1 uint64 = 91
    var v2 float32 = 1.5
    var v3 uint64 = (v1 - (v1 >> ((4294967295) & 63)))
    v4 := make(map[int]uint32)
    v5 := []int{519, 128, -279}
    var v16 int8 = -128
    for fi17 := 0; fi17 < 16; fi17++ {
        switch (fi17 >> ((255) & 63)) {
        case 2147483647:
            v18 := GenS3{F0: (^(v3)), F1: 0}
            _ = v18
        case 4294967295:
            switch -62 {
            case 9223372036854775807:
                if false {
                    fmt.Println(fi17, v16, v1)
                    if false {
                        v16 = v16
                        v19 := GenS1{F0: (v3 >> ((65535) & 63))}
                        _ = len(v4)
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 12; fi20++ {
                            if (true && (true || ("a" != ""))) {
                                v4[fi17] = 1
                            }
                            var v21 uint16 = 635
                            v1 = v1
                            _ = fi20
                            _ = v21
                        }
                        v2 = v2
                    }
                }
            default:
                switch ("a" + ("a" + "")) {
                case "\n\t":
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v16 = v16
                        var v23 uint8 = 255
                        var v24 uint16 = 1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    v1 = v1
                }
            }
        case 1:
            v16 -= (^(v16))
        default:
            v25 := v4[fi17]
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            v3 += v3
            v16 += v16
            var v27 uint32 = 2147483647
            _ = fi26
            _ = v27
        }
        _ = fi17
    }
    var v28 float32 = v2
    if (2 < 2) {
        if (false || ((v3 / ((v1) | 1)) == v3)) {
            v29 := GenS2{F0: (^(v16))}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int16 = -1
            var v33 bool = true
            v34 := v4[(fi30 + (fi30 >> ((18446744073709551615) & 63)))]
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 int8 = (v16 << ((42) & 63))
        _ = fi30
        _ = v35
    }
    v2 = (((v2 / v2) * v28) * 3.14159)
    v28 = v28
    v16 = (v16 + (-(v16)))
    var v36 float64 = 1.5
    v5 = append(v5, 127)
    var v37 bool = true
    fmt.Println(v3, v37, v37)
    v2 -= (float32(1e10) / v2)
    var v38 uint16 = 0
    if v37 {
        v39 := GenS4{F0: 42}
        var v40 uint64 = v3
        _ = v39
        _ = v40
    }
    var v41 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
