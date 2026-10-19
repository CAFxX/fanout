package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int; F1 int16; F2 int32 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
func genFn0(p0 uint64, p1 uint) uint8 {
    p1 += (^(p1))
    p0 += p0
    if false {
        p1 -= p1
    }
    v7 := []int{189, 127}
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 uint64, p1 string) string {
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 -= p0
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: -9223372036854775808, F1: fi9, F2: 1}
            _ = fi9
            _ = v10
        }
        switch p1 {
        case "":
            p1 = p1
        default:
            if true {
                p0 = (^(p0))
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 string = p1
                    _ = fi11
                    _ = v12
                }
            } else {
                v13 := GenS2{F0: 610, F1: 65535, F2: 65535}
                var v14 int32 = -1
                _ = v13
                _ = v14
            }
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    p1 = p1
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint64 = 332
        _ = fi17
        _ = v18
    }
    v19 := []int{32767, 32768, -959}
    _ = p0
    _ = p1
    _ = v19
    return p1
}

func main() {
    v1 := []uint32{}
    var v2 int = 32767
    v3 := []uint32{2147483648, 777, 188, 393}
    v4 := make(map[int]bool)
    v5 := []string{("\n\t" + ("hello" + (("a" + "\n\t") + ("foo bar" + ("a" + "hello"))))), ("hello" + "foo bar")}
    v6 := make(map[int]bool)
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi21
        }
        _ = fi20
    }
    v2 -= v2
    fmt.Println(v2)
    v2 = (v2 ^ -705)
    if true {
        var v22 uint64 = 2
        v22 += (v22 / ((4294967295) | 1))
        _ = v22
    } else {
        for fi23 := 0; fi23 < 17; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                delete(v4, (v2 ^ fi24))
                v2 = -1
                v25 := GenS1{F0: 1, F1: (-(v2)), F2: 0}
                v2 = v2
                if (false && (false && true)) {
                    v2 += v2
                    var v26 uint = 255
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
            if true {
                v2 += fi23
                switch v2 {
                case -999:
                    v27 := GenS1{F0: 0, F1: v2, F2: 2147483648}
                    _ = v27
                case 32768:
                    v28 := v4[(fi23 ^ fi23)]
                    _ = v28
                case -937:
                    var v29 int64 = 1
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        var v31 int64 = 1
                        fmt.Println(v2, v31, v31)
                        _ = fi30
                        _ = v31
                    }
                }
                v2 += fi23
            }
            _ = fi23
        }
    }
    v32 := GenS3{F0: v2, F1: 473, F2: -727}
    switch "hello" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v33 uint16 = 0
        _ = v33
    case "\n\t":
        var v34 uint64 = 2
        _ = v34
    default:
        var v35 string = ("a" + "\n\t")
        _ = v35
    }
    v2 = v2
    v2 = -964
    v1 = append(v1, 42)
    fmt.Println(v2)
    for fi36 := 0; fi36 < 5; fi36++ {
        if (!true) {
            if (genFn0(1, 65535) >= 127) {
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 string = (("a" + "foo bar") + "\n\t")
                    v38 = v38
                    _ = fi37
                    _ = v38
                }
                switch (v2 & 2147483647) {
                case 32768:
                    var v39 bool = (100 > 128)
                    _ = v39
                case 827:
                    v40 := GenS3{F0: 9223372036854775807, F1: -32768, F2: -1}
                    _ = v40
                case -258:
                    _ = len(v5)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            fmt.Println(fi36)
        } else {
            v2 -= fi36
            switch (v2 * (fi36 % ((v2) | 1))) {
            case -828:
                v41 := GenS4{F0: 1, F1: -9223372036854775808, F2: 9223372036854775807}
                _ = v41
            case -746:
                var v42 float64 = ((float64(-2.25) + 0.1) - ((float64(((3.14159 + 3.14159) * float32((float64(1e10) / 3.14159)))) * 1.5) * 1.5))
                _ = v42
            case 9223372036854775807:
                if len(v3) > 0 {
                    v3[0] = 42
                }
            default:
                v2 = v2
            }
        }
        _ = fi36
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
