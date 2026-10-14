package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 uint32) uint16 {
    v5 := GenS1{F0: 659, F1: true, F2: -213}
    if (false || (42 > 1)) {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int8 = 42
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := []int{9223372036854775807}
                var v10 bool = (false && ((1 <= 2) || (229 != 771)))
                var v11 uint16 = 32767
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 int = (fi6 / ((fi6) | 1))
            _ = fi6
            _ = v12
            _ = v7
        }
    } else {
        var v13 uint64 = 2
        _ = v13
    }
    p0 -= 2147483647
    for fi14 := 0; fi14 < 8; fi14++ {
        p0 -= (p0 | p0)
        _ = fi14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if (128 <= 1) {
            p0 -= (p0 / ((p0) | 1))
        } else {
            var v16 int64 = -876
            _ = v16
        }
        p0 += (p0 % (((^(p0))) | 1))
        var v17 int32 = 42
        _ = fi15
        _ = v17
    }
    _ = p0
    _ = v5
    return 32767
}
func genFn1() int16 {
    var v18 uint8 = 3
    var v19 float32 = float32(0.1)
    _ = v18
    _ = v19
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -332
    var v3 string = (((("hello" + "foo bar") + "\n\t") + "a") + ("\n\t" + "foo bar"))
    var v4 int16 = -32768
    var v20 int = 127
    var v21 uint8 = 239
    v1[738] = 600
    v4 -= 1
    v22 := GenS1{F0: 42, F1: true, F2: v2}
    v21 = v21
    var v23 uint64 = 281
    var v24 int64 = v2
    var v25 int8 = 101
    fmt.Println(v25, v23, v4)
    if (v3 > "") {
        for fi26 := 0; fi26 < 6; fi26++ {
            fmt.Println(v23)
            v27 := GenS2{F0: 32767}
            _ = fi26
            _ = v27
        }
        v21 -= 128
        v28 := []int{(v20 << ((18446744073709551615) & 63)), 9223372036854775807, v20, v20}
        _ = v28
    } else {
        v4 = genFn1()
    }
    v29 := GenS1{F0: 891, F1: (true || true), F2: v24}
    v30 := v1[2]
    v31 := GenS1{F0: 255, F1: true, F2: v2}
    if (519 < 300) {
        if (true && (true && true)) {
            v24 -= v24
            var v32 bool = ((true || false) || (-1 != v25))
            var v33 uint = 1
            _ = v32
            _ = v33
        } else {
            fmt.Println(v24, v21, v25)
            if (int(v2) <= (v20 << ((486) & 63))) {
                v20 += -9223372036854775808
            }
            switch v30 {
            case -1:
                for fi34 := 0; fi34 < 14; fi34++ {
                    v35 := GenS1{F0: 0, F1: true, F2: (v24 - v24)}
                    var v36 bool = true
                    v1[4294967295] = (v30 / ((fi34) | 1))
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case -9223372036854775808:
                v20 += (^(v30))
            case -216:
                v37 := GenS1{F0: 42, F1: (v25 <= (v25 * (v25 % ((61) | 1)))), F2: (-(v24))}
                _ = v37
            default:
                switch (v30 + v20) {
                case -389:
                    fmt.Println(v20)
                case 601:
                    var v38 int = (v20 & v30)
                    _ = v38
                default:
                    v39 := GenS3{F0: v25, F1: v4}
                    _ = v39
                }
            }
        }
        if ((v23 / (((v23 << ((0) & 63))) | 1)) > v23) {
            v40 := v1[0]
            switch ("hello" + ((v3 + "\n\t") + v3)) {
            case "a":
                if false {
                    switch v3 {
                    case "hello":
                        fmt.Println(v21)
                    case "":
                        switch v3 {
                        case "hello":
                            v40 += v40
                        }
                    }
                    v1[57] = v20
                    switch v3 {
                    case "":
                        var v41 int16 = int16(v30)
                        _ = v41
                    case "\n\t":
                        var v42 int32 = -1
                        _ = v42
                    case "a":
                        fmt.Println(v25, v4, v3)
                    }
                }
            default:
                delete(v1, 65535)
            }
            var v43 float64 = (float64(float32(1.5)) - (0.0 / 1.0))
            _ = v40
            _ = v43
        } else {
            v25 -= (v25 + -1)
            for fi44 := 0; fi44 < 19; fi44++ {
                delete(v1, 4294967295)
                var v45 int32 = 179
                v45 = (v45 & -267)
                _ = fi44
                _ = v45
            }
        }
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        v47 := v1[1]
        var v48 uint32 = 65535
        if true {
            v49 := GenS2{F0: 1}
            v20 = v20
            v50 := GenS4{F0: v3, F1: (v2 % (((v2 & int64(v48))) | 1))}
            _ = v49
            _ = v50
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    v1[2147483647] = int(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
