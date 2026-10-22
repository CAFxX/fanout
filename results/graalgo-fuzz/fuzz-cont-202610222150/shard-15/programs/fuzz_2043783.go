package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 int8; F2 bool }
func genFn0(p0 int16, p1 uint32) uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42}
    v10 := GenS1{F0: -2147483648}
    v11 := []int{int(p1), -9223372036854775808, 42}
    for fi12 := 0; fi12 < 4; fi12++ {
        switch fi12 {
        case 2147483647:
            p0 = p0
        default:
            v13 := GenS2{F0: 42, F1: 905, F2: 255}
            _ = v13
        }
        v14 := GenS2{F0: p0, F1: 2, F2: 1}
        switch -417 {
        case 9223372036854775807:
            p0 = (^(p0))
        case -176:
            v15 := GenS1{F0: 2147483647}
            _ = v15
        case -601:
            p0 -= (p0 >> ((18446744073709551615) & 63))
        default:
            p1 += 1
        }
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 579
}
func genFn1(p0 int16, p1 int8, p2 uint32) uint16 {
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int = fi17
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if false {
        v19 := []int{-9223372036854775808, 779, 128, -266}
        _ = v19
    }
    if true {
        if true {
            v20 := GenS2{F0: p0, F1: 42, F2: 33}
            switch ("" + ("a" + ((("foo bar" + "foo bar") + "foo bar") + "\n\t"))) {
            case "a":
                v21 := GenS2{F0: 1, F1: 0, F2: 26}
                _ = v21
            }
            v22 := []int{}
            _ = v20
            _ = v22
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                p0 = p0
                _ = fi23
            }
        }
        p2 = 0
        var v24 uint = 460
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 850
    v2 := []uint32{4294967295, 1}
    v3 := []int{}
    var v4 string = "foo bar"
    v5 := make(map[uint32]string)
    var v6 int64 = 1
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS3{F0: 255, F1: v4, F2: int(v1)}
        _ = fi25
        _ = v26
    }
    var v27 int16 = 32767
    if ("\n\t" <= (v4 + v4)) {
        if ((v27 - (^(v27))) != v27) {
            var v28 uint8 = 173
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            v5[1] = v4
            v6 += 42
            _ = fi30
        }
        fmt.Println(v1)
        if (int8(v27) >= int8(v27)) {
            v4 = ("" + ("\n\t" + (v4 + v4)))
            var v31 string = "a"
            v6 = 286
            _ = v31
        } else {
            _ = len(v3)
            if (42 > -1) {
                var v32 uint8 = 128
                fmt.Println(v1, v32)
                switch (fi29 - (fi29 / ((2147483647) | 1))) {
                case 128:
                    v1 += v1
                case 2147483647:
                    if false {
                        for fi33 := 0; fi33 < 18; fi33++ {
                            var v34 bool = false
                            v27 = (v27 / (((v27 ^ 32767)) | 1))
                            _ = fi33
                            _ = v34
                        }
                        if true {
                            v1 -= (^(v1))
                        }
                        if true {
                            v35 := GenS2{F0: v27, F1: 0, F2: (v32 + 0)}
                            fmt.Println(fi29)
                            for fi36 := 0; fi36 < 18; fi36++ {
                                fmt.Println(v1, v6)
                                _ = fi36
                            }
                            _ = v35
                        }
                    }
                }
                _ = v32
            } else {
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 int16 = v27
                    var v39 string = "hello"
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            v6 -= 0
        }
        _ = fi29
    }
    var v40 float32 = 1e10
    if false {
        if true {
            v5[0] = (((v4 + ("foo bar" + v4)) + v4) + (v4 + (v4 + (v4 + v4))))
            fmt.Println(v40, v27, v4)
            if len(v3) > 0 {
                v3[0] = 843
            }
        }
        v4 = "hello"
        if true {
            var v41 uint64 = 435
            var v42 int = 9223372036854775807
            v43 := GenS4{F0: 0, F1: int8(v41), F2: false}
            _ = v41
            _ = v42
            _ = v43
        } else {
            if (162 > 222) {
                v27 = 42
                var v44 int = 461
                _ = v44
            } else {
                v27 -= v27
                v4 = "a"
            }
        }
    } else {
        v27 -= v27
        v45 := GenS3{F0: 42, F1: ("a" + ((v4 + "hello") + v4)), F2: 688}
        _ = v45
    }
    switch ("a" + ((v4 + "\n\t") + v4)) {
    case "\n\t":
        for fi46 := 0; fi46 < 5; fi46++ {
            v27 += (-(v27))
            switch (fi46 / ((2147483647) | 1)) {
            case -1:
                var v47 uint16 = 42
                _ = v47
            case 283:
                switch -9223372036854775808 {
                case 54:
                    fmt.Println(v27, v27, v1)
                case -1:
                    v27 = (v27 ^ 1)
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    for fi48 := 0; fi48 < 11; fi48++ {
                        var v49 int64 = -1
                        v50 := GenS4{F0: 0, F1: -100, F2: (829 <= 666)}
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                default:
                    v1 -= v1
                }
            }
            _ = fi46
        }
    }
    switch 0 {
    case -67:
        fmt.Println(v6, v40)
    case 9223372036854775807:
        var v51 int16 = 1000
        _ = v51
    case 4:
        delete(v5, 605)
    }
    v52 := GenS4{F0: 527, F1: 127, F2: (((-9223372036854775808 <= (-(v1))) || true) && (!false))}
    fmt.Println(v6, v4)
    var v53 uint8 = 2
    var v54 int32 = -1
    for fi55 := 0; fi55 < 5; fi55++ {
        v56 := GenS2{F0: (v27 % ((v27) | 1)), F1: genFn1((v27 >> ((239) & 63)), -128, 2147483648), F2: v53}
        v4 = "a"
        _ = fi55
        _ = v56
    }
    var v57 float64 = 1.5
    v4 = v4
    v2 = append(v2, 925)
    var v58 float32 = (0.1 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v57
    _ = v58
    _ = v6
}
