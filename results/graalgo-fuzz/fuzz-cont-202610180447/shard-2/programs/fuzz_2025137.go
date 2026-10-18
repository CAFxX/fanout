package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int64 = -1
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    var v4 uint = 561
    v5 := make(map[string]int)
    switch "\n\t" {
    case "a":
        switch 32767 {
        case 32767:
            fmt.Println(v1, v4)
        default:
            v6 := GenS1{F0: 790}
            _ = v6
        }
    default:
        var v7 uint64 = 2
        _ = v7
    }
    v8 := []int{1, 32767}
    fmt.Println(v2, v4, v4)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (v4 == 42) {
        _ = len(v3)
        var v9 uint = 255
        v9 += v4
        _ = v9
    } else {
        if true {
            var v10 int64 = 9223372036854775807
            var v11 string = "a"
            _ = v10
            _ = v11
        } else {
            if (v2 != 127) {
                var v12 int16 = -1
                _ = v12
            } else {
                var v13 bool = (true && (!true))
                for fi14 := 0; fi14 < 17; fi14++ {
                    v2 = v2
                    v2 = (v2 + (^(v2)))
                    v15 := GenS2{F0: 255}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        }
    }
    v3[4294967295] = ((("" + "") + ("\n\t" + (("hello" + "foo bar") + "hello"))) + ("\n\t" + (("foo bar" + "a") + "a")))
    v16 := GenS2{F0: 255}
    var v17 int8 = -1
    v1 += v1
    var v18 float64 = 1.5
    var v19 int64 = v1
    switch 88 {
    case 128:
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                if ((v17 >> (((v4 % ((v4) | 1))) & 63)) < 42) {
                    v17 = (v17 % (((v17 / ((v17) | 1))) | 1))
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                if (597 < 2147483647) {
                    v19 = (v19 * v1)
                    v22 := GenS1{F0: 693}
                    _ = v22
                }
                fmt.Println(v4, v1, fi20)
                _ = fi21
            }
            fmt.Println(fi20)
            for fi23 := 0; fi23 < 14; fi23++ {
                v8 = append(v8, fi20)
                v24 := GenS2{F0: 65535}
                var v25 int = (fi20 & fi20)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi20
        }
    case -2147483648:
        for fi26 := 0; fi26 < 9; fi26++ {
            if false {
                v2 = (v2 & v2)
            } else {
                fmt.Println(v4, v2)
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := v5["\n\t"]
                switch (fi27 ^ fi27) {
                case 9223372036854775807:
                    if false {
                        v29 := GenS1{F0: 65535}
                        var v30 int = -1
                        _ = v29
                        _ = v30
                    }
                case 128:
                    v31 := GenS2{F0: 65535}
                    _ = v31
                case 639:
                    v32 := v3[uint32(v4)]
                    _ = v32
                default:
                    var v33 int16 = -32768
                    _ = v33
                }
                v34 := GenS3{F0: v19, F1: 1, F2: v2}
                v35 := GenS4{F0: (!false), F1: 1, F2: 32767}
                v36 := v3[963]
                _ = fi27
                _ = v28
                _ = v34
                _ = v35
                _ = v36
            }
            v37 := GenS5{F0: (v2 * 127), F1: v2, F2: (v4 << (((v4 + v4)) & 63))}
            _ = fi26
            _ = v37
        }
    case 9223372036854775807:
        v1 += -229
    default:
        v3[1] = "hello"
    }
    fmt.Println(v1, v18)
    if (-9223372036854775808 <= v1) {
        if true {
            if false {
                v4 = v4
                _ = len(v8)
            } else {
                var v38 string = "foo bar"
                if false {
                    v1 = -9223372036854775808
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v38
            }
            v19 += (v1 - (-(v19)))
            for fi39 := 0; fi39 < 17; fi39++ {
                v19 -= 1
                v40 := GenS6{F0: 1}
                _ = fi39
                _ = v40
            }
        }
        var v41 uint8 = v2
        _ = len(v3)
        _ = v41
    }
    v17 = v17
    v1 -= (v19 + 1)
    v42 := v3[2]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v8
}
