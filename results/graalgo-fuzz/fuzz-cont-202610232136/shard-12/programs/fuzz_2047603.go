package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 bool = true
        v8 := []int{fi6, (fi6 - (fi6 + fi6)), fi6, fi6}
        var v9 int16 = 32767
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v5
    return 1
}
func genFn1(p0 string, p1 int32) string {
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int8 = 1
        var v12 string = p0
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int16 = -355
    p0 = (p0 + "foo bar")
    if false {
        v14 := []int{4294967295, 32768, -1, 0}
        p0 = (("" + p0) + ("" + p0))
        _ = v14
    } else {
        var v15 float32 = (float32(-2.25) * 1.5)
        p1 = (-(p1))
        _ = v15
    }
    var v16 int32 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v16
    return (("\n\t" + (p0 + "hello")) + "a")
}

func main() {
    var v1 int8 = 0
    v2 := []int{1, 571}
    var v3 uint8 = 231
    v4 := make(map[int]int)
    switch "hello" {
    case "a":
        for fi17 := 0; fi17 < 3; fi17++ {
            switch (fi17 << ((2) & 63)) {
            case 42:
                fmt.Println(v1, fi17, v1)
            default:
                if len(v2) > 0 {
                    v2[0] = fi17
                }
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                v3 += (^(v3))
                var v19 string = ""
                v19 = ("" + "hello")
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    case "":
        switch 32768 {
        case 604:
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 uint64 = 400
                _ = fi20
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            switch ("hello" + "foo bar") {
            case "\n\t":
                if (9223372036854775807 >= -9223372036854775808) {
                    if (false || (false && (v3 != 91))) {
                        v24 := GenS1{F0: (^(fi23)), F1: -24, F2: (v3 << ((255) & 63))}
                        _ = v24
                    } else {
                        v1 -= v1
                        v3 = v3
                        switch "hello" {
                        case "hello":
                            var v25 float32 = (float32(((0.0 / 0.1) / (-2.25 + 1e10))) * -2.25)
                            _ = v25
                        case "":
                            v26 := GenS1{F0: (fi22 << ((4294967295) & 63)), F1: v1, F2: v3}
                            _ = v26
                        case "foo bar":
                            fmt.Println(fi23, fi22)
                        default:
                            fmt.Println(v1, fi22, v1)
                        }
                        v3 = (v3 & v3)
                        v1 += -1
                    }
                    var v27 string = ""
                    _ = v27
                } else {
                    var v28 uint = 4294967295
                    _ = v28
                }
                var v29 uint16 = 0
                v30 := GenS2{F0: ((-9223372036854775808 == -9223372036854775808) || false), F1: -492, F2: (v29 * 65535)}
                _ = v29
                _ = v30
            case "foo bar":
                v31 := v4[(fi23 - 42)]
                _ = v31
            }
            _ = fi23
        }
        _ = fi22
    }
    var v32 uint = 42
    var v33 uint8 = (v3 + 0)
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 float32 = (1e10 / -2.25)
        _ = fi34
        _ = v35
    }
    var v36 float64 = 1e10
    v37 := GenS3{F0: -876}
    var v38 int64 = 839
    for fi39 := 0; fi39 < 19; fi39++ {
        var v40 uint32 = 4294967295
        v3 = (v3 + (v33 & (^(v33))))
        _ = fi39
        _ = v40
    }
    if (0.1 >= 0.1) {
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 uint32 = 0
            for fi43 := 0; fi43 < 2; fi43++ {
                if (0.1 < 0.1) {
                    var v44 int16 = 457
                    _ = v44
                }
                v45 := GenS1{F0: (fi41 / (((fi41 ^ int(v1))) | 1)), F1: (v1 - (v1 >> ((v32) & 63))), F2: uint8(v42)}
                var v46 uint32 = v42
                _ = fi43
                _ = v45
                _ = v46
            }
            v3 += (^(v33))
            _ = fi41
            _ = v42
        }
        var v47 uint32 = 2147483647
        _ = v47
    } else {
        fmt.Println(v38, v32)
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 uint16 = 32767
            var v50 bool = true
            _ = fi48
            _ = v49
            _ = v50
        }
        if true {
            var v51 int8 = 30
            if ((v3 + v3) >= (v3 >> (((v32 - (v32 - (v32 | 1)))) & 63))) {
                if true {
                    v4[32768] = -2147483648
                    if false {
                        v38 = v38
                    } else {
                        v51 = v1
                        var v52 uint32 = 310
                        _ = v52
                    }
                }
                switch "a" {
                case "foo bar":
                    var v53 uint16 = 1
                    _ = v53
                case "a":
                    fmt.Println(v33)
                case "\n\t":
                    var v54 float64 = float64(3.14159)
                    _ = v54
                default:
                    var v55 uint = 2
                    _ = v55
                }
            }
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
