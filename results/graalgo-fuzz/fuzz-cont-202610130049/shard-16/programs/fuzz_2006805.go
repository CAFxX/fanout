package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 int8; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 bool; F2 string }
func genFn0() int32 {
    if false {
        v4 := GenS1{F0: 994}
        var v5 uint8 = 42
        var v6 int16 = 0
        _ = v4
        _ = v5
        _ = v6
    }
    if ("" != ("a" + "hello")) {
        switch -9223372036854775808 {
        case -90:
            var v7 uint16 = 0
            _ = v7
        case 113:
            switch 32768 {
            case 42:
                for fi8 := 0; fi8 < 3; fi8++ {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := []int{127, fi8, fi8, -886}
                        switch (fi8 / ((fi8) | 1)) {
                        case 9223372036854775807:
                            if (((3.14159 * (3.14159 / 0.1)) - 3.14159) > (float32(0.1) * -2.25)) {
                                var v11 int64 = 1
                                v12 := GenS2{F0: 0, F1: 428}
                                _ = v11
                                _ = v12
                            } else {
                                var v13 float32 = 1.5
                                _ = v13
                            }
                        case -9223372036854775808:
                            if false {
                                v14 := GenS3{F0: 4294967295, F1: 2}
                                if ((fi8 | -943) == 32768) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                    var v15 string = (("\n\t" + "hello") + "hello")
                                    _ = v15
                                }
                                var v16 int8 = 0
                                _ = v14
                                _ = v16
                            }
                        }
                        fmt.Println(fi9, fi9)
                        _ = fi9
                        _ = v10
                    }
                    if (("" + "\n\t") == "a") {
                        var v17 uint16 = 32767
                        _ = v17
                    }
                    fmt.Println(fi8)
                    _ = fi8
                }
            case 127:
                if (128 < -154) {
                    var v18 bool = (218 != 2)
                    v19 := GenS1{F0: 42}
                    var v20 bool = (!v18)
                    _ = v18
                    _ = v19
                    _ = v20
                }
            case 176:
                var v21 bool = true
                _ = v21
            }
        default:
            var v22 int = 42
            _ = v22
        }
        if false {
            var v23 int32 = 0
            var v24 uint = 361
            _ = v23
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS4{F0: 42, F1: 0, F2: -1}
            switch fi25 {
            case 321:
                var v27 uint = 65535
                _ = v27
            case -948:
                v28 := []int{fi25, fi25, (fi25 % ((fi25) | 1)), (fi25 * fi25)}
                _ = v28
            default:
                if true {
                    var v29 float32 = 3.14159
                    _ = v29
                }
            }
            fmt.Println(fi25)
            _ = fi25
            _ = v26
        }
    }
    v30 := []int{943, 128, 42}
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int16 = -901
        _ = len(v30)
        if true {
            if true {
                switch ("" + "") {
                case "\n\t":
                    fmt.Println(fi31)
                default:
                    var v33 int16 = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 14; fi34++ {
                    v32 += (-(v32))
                    v32 = (v32 % ((v32) | 1))
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    _ = fi34
                }
            }
            if (-2147483648 != 286) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    fmt.Println(v32)
                case "hello":
                    switch ("foo bar" + "") {
                    case "\n\t":
                        v32 -= v32
                    case "foo bar":
                        var v35 uint = 42
                        _ = v35
                    case "hello":
                        fmt.Println(fi31)
                    default:
                        switch fi31 {
                        case 127:
                            switch "a" {
                            case "foo bar":
                                switch (fi31 | 128) {
                                case 64:
                                    v32 = (-(v32))
                                case 9223372036854775807:
                                    switch (-(fi31)) {
                                    case 42:
                                        v32 = (v32 | v32)
                                    case -565:
                                        var v36 bool = (463 < 4294967295)
                                        _ = v36
                                    default:
                                        v32 += (v32 << ((2) & 63))
                                    }
                                }
                            }
                        case 42:
                            v30 = append(v30, 9223372036854775807)
                        }
                    }
                default:
                    switch (fi31 - fi31) {
                    case 2147483647:
                        var v37 uint = 857
                        _ = v37
                    case 0:
                        var v38 uint = 18446744073709551615
                        _ = v38
                    default:
                        var v39 float32 = 1.5
                        _ = v39
                    }
                }
            } else {
                var v40 uint16 = 832
                for fi41 := 0; fi41 < 20; fi41++ {
                    v42 := GenS3{F0: 0, F1: 18446744073709551615}
                    v40 = v40
                    v43 := GenS5{F0: 128}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            }
        }
        _ = fi31
        _ = v32
    }
    _ = v30
    return 931
}

func main() {
    var v1 uint = 65535
    var v2 float32 = 3.14159
    v3 := make(map[int]string)
    if false {
        fmt.Println(v2, v2)
    } else {
        _ = len(v3)
    }
    v44 := v3[-339]
    v3[-53] = v44
    v45 := GenS6{F0: -70, F1: 5, F2: 0}
    v44 = ((v44 + v44) + "a")
    v46 := GenS7{F0: 42}
    if (3.14159 > 0.1) {
        if (719 >= 837) {
            v44 = v44
            for fi47 := 0; fi47 < 1; fi47++ {
                v48 := GenS1{F0: -241}
                v49 := []int{156}
                v2 += v2
                _ = fi47
                _ = v48
                _ = v49
            }
            v2 = (1e10 - 3.14159)
        }
        v50 := GenS5{F0: -2147483648}
        _ = v50
    }
    var v51 float64 = 1e10
    if (0 <= -1) {
        v2 += ((((v2 + ((3.14159 * (0.0 / ((v2 / float32(((-2.25 * (((v51 * v51) / 1.5) - v51)) + float64(v2)))) / 1e10))) - v2)) - float32(v51)) + v2) - v2)
        v52 := []int{-2147483648, -182, 128}
        fmt.Println(v51, v2)
        _ = v52
    } else {
        switch (v44 + (v44 + v44)) {
        case "hello":
            for fi53 := 0; fi53 < 9; fi53++ {
                v3[(fi53 | (fi53 - fi53))] = v44
                _ = fi53
            }
        default:
            v1 += (v1 + (v1 ^ (v1 | (v1 >> ((18446744073709551615) & 63)))))
        }
        var v54 int8 = 42
        _ = v54
    }
    v55 := GenS8{F0: 51, F1: true, F2: ("a" + v44)}
    for fi56 := 0; fi56 < 14; fi56++ {
        v44 = v44
        if (fi56 >= 4294967295) {
            v57 := []int{(fi56 / ((fi56) | 1)), (fi56 | fi56)}
            delete(v3, fi56)
            _ = v57
        }
        _ = fi56
    }
    v58 := GenS1{F0: -32768}
    fmt.Println(v51)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v51
    _ = v55
    _ = v58
}
