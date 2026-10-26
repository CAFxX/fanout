package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int16 }
func genFn0(p0 int32, p1 uint32, p2 uint) uint16 {
    v4 := GenS1{F0: (p2 / ((p2) | 1)), F1: 98}
    p2 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 186
}
func genFn1(p0 int64, p1 uint16, p2 uint32) int64 {
    p0 = p0
    p0 += p0
    p2 += 2
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = -128
    v3 := make(map[string]uint32)
    v1 -= 4294967295
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 += 18446744073709551615
        var v6 float32 = (0.0 * -2.25)
        v7 := GenS2{F0: 128, F1: (!(!(false && true))), F2: 100}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float32 = 0.0
    var v9 uint16 = uint16(v1)
    var v10 float32 = ((0.1 + (float32(float64(v8)) * ((1e10 * v8) * 0.0))) - (1e10 / float32(1.5)))
    fmt.Println(v10, v2, v8)
    v10 -= (v10 - (1.5 / float32(1e10)))
    v11 := GenS2{F0: 1, F1: false, F2: (v2 / ((v2) | 1))}
    var v12 string = "hello"
    v2 = int8(v9)
    switch 724 {
    case -208:
        fmt.Println(v8, v9)
    case -805:
        for fi13 := 0; fi13 < 20; fi13++ {
            switch fi13 {
            case 32768:
                v8 += (float32(-2.25) * 1e10)
            default:
                v10 = v8
            }
            var v14 int8 = (^(v2))
            v15 := []int{fi13, (fi13 >> ((115) & 63)), 32767, fi13}
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        var v16 int32 = -1
        _ = v16
    }
    if (4294967295 >= -616) {
        v17 := []int{726, 127}
        var v18 bool = (0 > 9223372036854775807)
        _ = v17
        _ = v18
    }
    switch v12 {
    case "a":
        switch 882 {
        case 128:
            var v19 uint8 = 127
            _ = v19
        default:
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v9 += v9
                    v10 = v8
                    fmt.Println(v10)
                    _ = fi21
                }
                v22 := GenS3{F0: 0, F1: 445, F2: 32767}
                if false {
                    v10 = (((3.14159 - v8) * 0.1) / -2.25)
                } else {
                    var v23 int = fi20
                    _ = v23
                }
                _ = fi20
                _ = v22
            }
        }
    case "":
        v8 -= (v10 / 1.0)
    case "foo bar":
        v10 += -2.25
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := v3[""]
        _ = fi24
        _ = v25
    }
    var v26 int = -9223372036854775808
    for fi27 := 0; fi27 < 18; fi27++ {
        switch (("foo bar" + "") + (v12 + (v12 + v12))) {
        case "\n\t":
            for fi28 := 0; fi28 < 8; fi28++ {
                v9 -= (v9 + v9)
                v26 = (^(v26))
                _ = fi28
            }
        case "a":
            for fi29 := 0; fi29 < 1; fi29++ {
                switch "" {
                case "\n\t":
                    fmt.Println(v9, v9, v10)
                case "foo bar":
                    v12 = ((v12 + "\n\t") + v12)
                }
                v12 = ""
                _ = fi29
            }
        case "hello":
            if true {
                var v30 uint16 = v9
                switch v26 {
                case -2147483648:
                    var v31 int16 = 32767
                    _ = v31
                case 32768:
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 int = fi27
                        var v34 int32 = 996
                        var v35 string = "\n\t"
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                case 127:
                    if (false && ((v30 | v30) >= (^(v9)))) {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            v10 = v10
                            _ = fi36
                        }
                    }
                default:
                    var v37 uint16 = 65535
                    _ = v37
                }
                _ = v30
            } else {
                v38 := v3[v12]
                _ = v38
            }
        default:
            switch fi27 {
            case 42:
                for fi39 := 0; fi39 < 1; fi39++ {
                    v40 := GenS4{F0: v9, F1: (v2 | (v2 << ((4294967295) & 63))), F2: 42}
                    _ = fi39
                    _ = v40
                }
            case -2147483648:
                if ((!(true && false)) || (genFn0(2147483647, 2147483647, 0) <= v9)) {
                    var v41 uint = 580
                    v2 = -27
                    for fi42 := 0; fi42 < 11; fi42++ {
                        if false {
                            var v43 int = v26
                            _ = v43
                        } else {
                            fmt.Println(v2, v8)
                        }
                        fmt.Println(v8, fi42, fi42)
                        if true {
                            v44 := GenS4{F0: v9, F1: v2, F2: 1}
                            v8 += 0.1
                            v10 = 0.1
                            var v45 uint = 65535
                            _ = v44
                            _ = v45
                        }
                        _ = fi42
                    }
                    _ = v41
                }
            case 239:
                if ((32767 != 42) && true) {
                    var v46 uint16 = (v9 - genFn0(1, 2, 42))
                    _ = v46
                }
            default:
                var v47 uint32 = 515
                _ = v47
            }
        }
        _ = fi27
    }
    fmt.Println(v10, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
