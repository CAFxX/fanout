package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 bool, p1 string, p2 int32) int32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 790
        p1 = "\n\t"
        var v8 bool = true
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p2, p1)
    if (!true) {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 42}
            switch fi9 {
            case 9223372036854775807:
                for fi11 := 0; fi11 < 15; fi11++ {
                    p2 = (p2 ^ p2)
                    _ = fi11
                }
            default:
                var v12 string = p1
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        p1 = "a"
    }
    switch ((p1 + "") + p1) {
    case "foo bar":
        v13 := GenS2{F0: 0, F1: 65535, F2: 1}
        _ = v13
    case "hello":
        if (true && (p0 && (p2 != p2))) {
            p1 = ("foo bar" + "")
            v14 := []int{-849, 127}
            if (!true) {
                if false {
                    p1 = (p1 + ((p1 + "\n\t") + ("\n\t" + p1)))
                    var v15 float64 = (3.14159 - (1.5 + (float64(0.1) / 1e10)))
                    v16 := GenS2{F0: 93, F1: 239, F2: 300}
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint = 4294967295
                    var v18 uint = v17
                    _ = len(v14)
                    _ = v17
                    _ = v18
                }
            }
            _ = v14
        }
    case "":
        v19 := GenS1{F0: 255}
        _ = v19
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 int32 = (p2 / ((p2) | 1))
                v23 := []int{0}
                var v24 int = (^(fi20))
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v25 := []int{}
            _ = fi20
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 25
    var v3 uint64 = 458
    var v4 uint = 79
    var v5 uint64 = (v3 | (^(v3)))
    v26 := []int{-769, 32767, 1, 21}
    for fi27 := 0; fi27 < 14; fi27++ {
        if ((0.0 * (((float32((1e10 / 1.0)) - (float32(-2.25) + 0.0)) * 1.5) - -2.25)) > 0.1) {
            v2 = (v2 - (-(v2)))
            switch "" {
            case "":
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            case "foo bar":
                v28 := GenS3{F0: 231}
                _ = v28
            case "\n\t":
                var v29 bool = (1e10 < float32(3.14159))
                _ = v29
            default:
                fmt.Println(fi27)
            }
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := v1[fi30]
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 int32 = 42
        for fi34 := 0; fi34 < 11; fi34++ {
            v3 -= 0
            v35 := GenS4{F0: v3, F1: 1, F2: ("\n\t" + "hello")}
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v2 = v2
    for fi36 := 0; fi36 < 13; fi36++ {
        v37 := v1[fi36]
        _ = fi36
        _ = v37
    }
    var v38 int32 = (v2 - v2)
    var v39 uint = (v4 * v4)
    _ = len(v26)
    v5 = (v3 | (^(v3)))
    if (224 == 165) {
        v40 := GenS2{F0: 127, F1: v4, F2: 65535}
        v41 := GenS3{F0: 18}
        for fi42 := 0; fi42 < 9; fi42++ {
            for fi43 := 0; fi43 < 16; fi43++ {
                switch fi43 {
                case 9223372036854775807:
                    fmt.Println(fi43)
                case 42:
                    if (883 < 476) {
                        if (!false) {
                            v44 := GenS2{F0: 157, F1: 255, F2: 32767}
                            if len(v26) > 0 {
                                v26[0] = fi42
                            }
                            _ = len(v26)
                            _ = v44
                        } else {
                            _ = len(v1)
                        }
                        v45 := GenS5{F0: false, F1: 1000}
                        _ = v45
                    } else {
                        if (!(false || true)) {
                            v46 := GenS6{F0: -983}
                            _ = v46
                        }
                        var v47 int = (^(fi42))
                        if len(v26) > 0 {
                            v26[0] = v47
                        }
                        var v48 uint = (^(v39))
                        _ = v47
                        _ = v48
                    }
                }
                _ = fi43
            }
            if false {
                if false {
                    var v49 bool = true
                    _ = v49
                }
                for fi50 := 0; fi50 < 18; fi50++ {
                    var v51 uint32 = 600
                    _ = fi50
                    _ = v51
                }
                _ = len(v1)
            }
            _ = fi42
        }
        _ = v40
        _ = v41
    }
    v39 -= (v39 + uint(v5))
    if true {
        if ((float64(0.1) + 1.5) < 1.5) {
            v39 += (v39 | v39)
        }
        v5 = v3
    }
    v1[32768] = 2
    v4 -= v4
    var v52 uint64 = v3
    if len(v26) > 0 {
        v26[0] = 9223372036854775807
    }
    v3 = (v3 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v52
}
