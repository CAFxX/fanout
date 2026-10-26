package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int; F1 uint; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint32; F1 bool }
func genFn0(p0 int16, p1 uint, p2 uint64) int16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 118, F1: 42, F2: p1}
        _ = fi6
        _ = v7
    }
    p2 = (p2 - (^(p2)))
    fmt.Println(p0, p2)
    v8 := GenS2{F0: -168, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}
func genFn1(p0 uint) uint16 {
    v9 := []int{4294967295, 128, -1, 1}
    var v10 float64 = 3.14159
    if (1.5 > (float32(((v10 / ((float64(3.14159) * v10) - v10)) / v10)) / 1.0)) {
        p0 = 858
        var v11 float64 = (v10 * v10)
        var v12 bool = true
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    var v13 string = "\n\t"
    v13 = v13
    if (!(false || ((2147483647 != 42) && true))) {
        v14 := []int{-2147483648}
        switch (((("" + (((v13 + "foo bar") + (v13 + v13)) + ("\n\t" + v13))) + v13) + "") + (v13 + v13)) {
        case "":
            var v15 float32 = -2.25
            _ = v15
        case "\n\t":
            v16 := GenS3{F0: 128, F1: 682, F2: -2147483648}
            _ = v16
        default:
            if false {
                if (!true) {
                    v2 -= v2
                    v2 = v2
                }
            } else {
                v13 = v13
                fmt.Println(v13)
            }
        }
        v2 = v1
        _ = v14
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint32 = 65535
                delete(v4, v19)
                var v20 int = (fi17 % ((fi17) | 1))
                if false {
                    fmt.Println(v13)
                }
                v21 := GenS4{F0: 1}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v13)
                _ = fi22
            }
            _ = fi17
        }
        v23 := GenS3{F0: 4294967295, F1: 48, F2: 1}
        if false {
            switch -292 {
            case -477:
                fmt.Println(v2, v1)
            }
            fmt.Println(v13, v2, v2)
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS5{F0: 1}
                var v26 uint32 = 4294967295
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            v1 -= v1
            var v27 uint32 = 13
            _ = v27
        }
        _ = v23
    }
    var v28 uint32 = 0
    v29 := GenS5{F0: 202}
    v1 = (v1 * (v2 << ((65535) & 63)))
    for fi30 := 0; fi30 < 6; fi30++ {
        if ((-701 >= -215) && ((v28 / ((v28) | 1)) < v28)) {
            if false {
                v28 -= 2147483647
                switch "hello" {
                case "\n\t":
                    for fi31 := 0; fi31 < 9; fi31++ {
                        fmt.Println(v13)
                        _ = fi31
                    }
                default:
                    v28 -= v28
                }
                v32 := GenS5{F0: 42}
                _ = v32
            } else {
                v33 := []int{-1, 42}
                var v34 uint8 = 226
                switch fi30 {
                case -2147483648:
                    v34 = v34
                default:
                    var v35 int8 = 98
                    _ = v35
                }
                _ = v33
                _ = v34
            }
        } else {
            var v36 string = "\n\t"
            var v37 int = 4294967295
            _ = v36
            _ = v37
        }
        _ = fi30
    }
    delete(v5, v13)
    if (((3.14159 + 1.5) <= -2.25) || (true && (true || true))) {
        v38 := GenS1{F0: v28, F1: 1, F2: 335}
        for fi39 := 0; fi39 < 10; fi39++ {
            v1 = 759
            var v40 int64 = -475
            _ = fi39
            _ = v40
        }
        if ((v28 * v28) > uint32(v2)) {
            v28 = (v28 ^ v28)
        } else {
            v41 := GenS5{F0: 42}
            _ = v41
        }
        _ = v38
    }
    if ((1.5 / 1.5) != (3.14159 / 1.5)) {
        if (0 >= -32768) {
            var v42 uint8 = 65
            v42 = v42
            _ = v42
        } else {
            if (true && true) {
                if false {
                    v28 = 65535
                }
            } else {
                var v43 uint8 = 1
                _ = len(v4)
                v43 = (^(v43))
                _ = v43
            }
            var v44 uint32 = 2147483648
            var v45 uint16 = 42
            _ = v44
            _ = v45
        }
    }
    if ("hello" >= "a") {
        v5[("foo bar" + "")] = (v28 | 89)
    } else {
        if (2147483647 >= v28) {
            switch (v13 + "\n\t") {
            case "a":
                v46 := []int{}
                _ = v46
            default:
                var v47 uint16 = 1
                _ = v47
            }
            if true {
                switch "a" {
                case "a":
                    fmt.Println(v28, v28)
                case "":
                    fmt.Println(v1, v2, v13)
                case "foo bar":
                    fmt.Println(v28, v1, v1)
                default:
                    v48 := GenS6{F0: int8(v28)}
                    _ = v48
                }
                for fi49 := 0; fi49 < 5; fi49++ {
                    var v50 int32 = -477
                    _ = fi49
                    _ = v50
                }
            }
            v51 := GenS7{F0: 173, F1: (v28 * (^(v28)))}
            _ = v51
        } else {
            v28 -= (^(v28))
            if false {
                v13 = "a"
                if ((v1 ^ 1) <= 42) {
                    var v52 int = 42
                    v53 := GenS3{F0: v52, F1: 0, F2: -275}
                    _ = v52
                    _ = v53
                }
                v54 := GenS8{F0: v28, F1: (51 <= 18446744073709551615)}
                _ = v54
            }
            fmt.Println(v28, v1, v1)
        }
        var v55 int8 = 127
        v2 = (v1 & v2)
        _ = v55
    }
    var v56 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
