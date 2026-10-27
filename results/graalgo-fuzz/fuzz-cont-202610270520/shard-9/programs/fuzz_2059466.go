package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint64; F1 string; F2 uint16 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint32 }
func genFn0() int {
    var v7 int16 = -1
    if (0 < 511) {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{fi8, 0}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v10 int8 = -128
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = v11
    }
    v7 = -727
    v12 := GenS1{F0: 0, F1: 9223372036854775807, F2: 341}
    _ = v12
    _ = v7
    return 57
}
func genFn1(p0 uint32, p1 int64) uint64 {
    var v13 string = ("a" + "")
    fmt.Println(p0)
    v14 := []int{-1, genFn0(), 0}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 821
}

func main() {
    v1 := []uint32{367, 42, 1, 0}
    var v2 uint16 = 255
    var v3 uint32 = 641
    var v4 uint = 406
    v5 := []uint32{v3}
    v6 := make(map[string]string)
    if ((42 < -128) || false) {
        v3 -= (v3 & 263)
        v3 = v3
    }
    v3 = 4294967295
    v3 = (^(v3))
    if (v4 < (v4 % ((v4) | 1))) {
        if true {
            for fi15 := 0; fi15 < 8; fi15++ {
                if len(v5) > 0 {
                    v5[0] = v3
                }
                _ = fi15
            }
        }
        fmt.Println(v4, v4, v4)
        switch "hello" {
        case "hello":
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 float32 = 3.14159
                v18 := GenS2{F0: v2, F1: -453}
                _ = fi16
                _ = v17
                _ = v18
            }
        case "a":
            v2 += (v2 + (v2 - 32767))
        }
    } else {
        v19 := GenS1{F0: -1, F1: genFn0(), F2: v4}
        v20 := v6["hello"]
        v6[(v20 + v20)] = v20
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if true {
            _ = len(v5)
            _ = len(v6)
            v4 = 255
        } else {
            switch (fi21 >> ((v4) & 63)) {
            case -9223372036854775808:
                var v22 string = ("\n\t" + "\n\t")
                _ = v22
            case -517:
                v23 := GenS1{F0: 0, F1: genFn0(), F2: v4}
                _ = v23
            default:
                v3 = (v3 ^ (v3 << (((^(v4))) & 63)))
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                if (!false) {
                    v3 += 410
                }
                var v25 uint8 = 212
                _ = fi24
                _ = v25
            }
        }
        if false {
            if (true || (0 > (v2 << ((1) & 63)))) {
                v3 += 2
            }
        } else {
            if false {
                fmt.Println(v2, fi21, v3)
            } else {
                var v26 uint8 = 42
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            if (-2.25 == 1.5) {
                v28 := GenS3{F0: -32768, F1: true, F2: v4}
                _ = v28
            } else {
                var v29 float32 = float32(1e10)
                _ = v29
            }
            _ = fi27
        }
        _ = fi21
    }
    var v30 int8 = 34
    for fi31 := 0; fi31 < 15; fi31++ {
        v1 = append(v1, (^(v3)))
        v5 = append(v5, (v3 - (^(v3))))
        _ = fi31
    }
    if false {
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint16 = v2
            v2 = uint16(v3)
            if true {
                if (((1.5 * 0.0) * ((1.5 - float64(1.5)) * (0.0 / 1.0))) > ((1e10 + 3.14159) / 3.14159)) {
                    if (!(!((255 == 128) || false))) {
                        v2 = (^(v33))
                        var v34 uint16 = v2
                        v6["\n\t"] = "a"
                        _ = v34
                    } else {
                        v6["hello"] = "foo bar"
                        v35 := GenS4{F0: fi32, F1: 1, F2: -587}
                        _ = v35
                    }
                    var v36 int32 = 1
                    v37 := GenS5{F0: "a", F1: (^(fi32))}
                    _ = v36
                    _ = v37
                } else {
                    if true {
                        var v38 int64 = 42
                        for fi39 := 0; fi39 < 18; fi39++ {
                            v33 += v33
                            fmt.Println(v38)
                            v40 := GenS6{F0: 992, F1: (("foo bar" + "") + ("hello" + "")), F2: v33}
                            _ = fi39
                            _ = v40
                        }
                        v41 := GenS2{F0: 906, F1: (v38 % ((-9223372036854775808) | 1))}
                        _ = v38
                        _ = v41
                    }
                }
            } else {
                v42 := GenS7{F0: uint8(v4), F1: v30}
                switch "foo bar" {
                case "":
                    fmt.Println(v4)
                case "hello":
                    fmt.Println(fi32)
                case "a":
                    var v43 uint = 255
                    _ = v43
                default:
                    var v44 uint = (v4 & v4)
                    _ = v44
                }
                var v45 float32 = 0.0
                _ = v42
                _ = v45
            }
            _ = fi32
            _ = v33
        }
        var v46 float64 = 0.0
        _ = v46
    } else {
        for fi47 := 0; fi47 < 3; fi47++ {
            v3 += 2
            _ = fi47
        }
        v48 := GenS8{F0: v3}
        _ = v48
    }
    v49 := v6["\n\t"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
