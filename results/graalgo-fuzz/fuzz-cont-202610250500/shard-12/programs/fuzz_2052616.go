package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 uint64, p1 bool, p2 uint) uint8 {
    var v4 string = ""
    var v5 uint64 = (p0 | 42)
    var v6 uint8 = 255
    p0 = v5
    v7 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return uint8(p0)
}
func genFn1(p0 uint16, p1 int16) uint8 {
    v8 := GenS1{F0: 2, F1: 273}
    var v9 int8 = 127
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                p0 -= (^(p0))
                switch 0 {
                case 1:
                    if true {
                        v12 := []int{}
                        _ = v12
                    }
                case 42:
                    v13 := GenS1{F0: 2, F1: 65535}
                    _ = v13
                default:
                    var v14 int32 = 1
                    _ = v14
                }
                _ = fi11
            }
            var v15 float64 = (0.0 / float64((((float32(3.14159) + 1.5) + (float32((0.0 * (-2.25 + (-2.25 / 1.0)))) + 3.14159)) + ((1.5 * ((((0.1 + ((float32(1.5) / 1e10) + (float32(1e10) - 0.0))) + float32(-2.25)) * 3.14159) / ((float32(((float64(-2.25) / (1.5 - 3.14159)) / 1.0)) + float32(1.5)) * 3.14159))) + 3.14159))))
            _ = fi10
            _ = v15
        }
        if false {
            p1 = p1
            p1 -= (p1 >> ((517) & 63))
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 bool = (-9223372036854775808 == (fi16 & -598))
                _ = fi16
                _ = v17
            }
        }
        v18 := []int{9223372036854775807}
        _ = v18
    }
    var v19 uint64 = 42
    if (42 != 9223372036854775807) {
        v19 += v19
        var v20 string = "a"
        if false {
            switch 4294967295 {
            case -798:
                v21 := []int{-428, 206, 42, 32768}
                _ = v21
            case 487:
                var v22 string = v20
                _ = v22
            default:
                var v23 string = v20
                _ = v23
            }
        } else {
            var v24 int8 = v9
            v9 += v9
            v25 := GenS1{F0: 4294967295, F1: 4294967295}
            _ = v24
            _ = v25
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    v1 = ""
    v1 = (v1 + v1)
    v26 := v3["\n\t"]
    v1 = "foo bar"
    var v27 int16 = 0
    var v28 float64 = (1.5 + 3.14159)
    if (487 == 792) {
        if (716 >= 42) {
            for fi29 := 0; fi29 < 9; fi29++ {
                v1 = v1
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(fi29, v28)
                    if (42 == 9223372036854775807) {
                        fmt.Println(v1)
                        v27 = (v27 ^ (^(v27)))
                        var v31 int32 = -1
                        var v32 int = fi29
                        _ = v31
                        _ = v32
                    } else {
                        fmt.Println(v28, fi30, fi30)
                        v33 := []int{(fi29 & (fi30 >> ((42) & 63))), fi30, (fi30 | fi30)}
                        _ = len(v33)
                        _ = v33
                    }
                    _ = fi30
                }
                _ = fi29
            }
        } else {
            v34 := GenS2{F0: false, F1: 1, F2: 2147483648}
            fmt.Println(v1, v1, v28)
            v35 := GenS1{F0: 707, F1: (v26 ^ 2147483648)}
            _ = v34
            _ = v35
        }
        var v36 int = 127
        _ = v36
    } else {
        v37 := []int{32768}
        _ = v37
    }
    v38 := GenS2{F0: false, F1: -1, F2: 41}
    switch 0 {
    case -829:
        switch 4294967295 {
        case 32767:
            delete(v2, (v26 ^ (^(v26))))
        case 2147483647:
            v28 = ((v28 + v28) / (v28 - 1e10))
        case -868:
            switch ("\n\t" + (v1 + v1)) {
            case "":
                if true {
                    v27 += (v27 % ((-32768) | 1))
                    v39 := GenS3{F0: 2, F1: -459, F2: 551}
                    for fi40 := 0; fi40 < 17; fi40++ {
                        v41 := GenS3{F0: 423, F1: -1, F2: 65535}
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                } else {
                    v28 += ((v28 - (v28 / (-2.25 - 0.0))) / v28)
                }
            case "hello":
                var v42 uint8 = 128
                _ = v42
            default:
                var v43 int64 = 9223372036854775807
                _ = v43
            }
        default:
            var v44 float64 = v28
            _ = v44
        }
    case -16:
        v45 := GenS4{F0: 1, F1: 255, F2: v27}
        _ = v45
    case -2147483648:
        for fi46 := 0; fi46 < 3; fi46++ {
            v1 = v1
            _ = len(v2)
            var v47 int32 = 235
            _ = fi46
            _ = v47
        }
    }
    var v48 int64 = -9223372036854775808
    var v49 uint8 = 252
    if ((2 >= 1) && (-757 >= -2147483648)) {
        switch ((("foo bar" + v1) + v1) + v1) {
        case "foo bar":
            var v50 float32 = float32((-2.25 * (v28 + v28)))
            _ = v50
        default:
            v51 := []int{651}
            _ = v51
        }
        var v52 uint16 = 42
        switch -1 {
        case 32768:
            v53 := []int{}
            _ = v53
        default:
            switch "foo bar" {
            case "a":
                fmt.Println(v52, v48, v49)
            default:
                v1 = "foo bar"
            }
        }
        _ = v52
    } else {
        v54 := v2[(v26 ^ v26)]
        v26 = uint32(v49)
        _ = v54
    }
    v55 := []int{}
    var v56 uint16 = 189
    for fi57 := 0; fi57 < 15; fi57++ {
        fmt.Println(v26, v56, v28)
        _ = len(v55)
        var v58 uint64 = 18446744073709551615
        _ = fi57
        _ = v58
    }
    var v59 uint32 = 4294967295
    var v60 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v48
    _ = v49
    _ = v55
    _ = v56
    _ = v59
    _ = v60
}
