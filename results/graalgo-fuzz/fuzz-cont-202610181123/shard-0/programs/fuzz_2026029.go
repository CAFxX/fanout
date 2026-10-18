package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 uint32; F2 string }

func main() {
    var v1 string = "a"
    var v2 int8 = -55
    v3 := make(map[string]uint32)
    var v4 uint = 4294967295
    var v5 uint = v4
    v2 = 68
    v6 := v3["hello"]
    fmt.Println(v4, v1)
    v7 := GenS1{F0: false, F1: 2, F2: 1}
    var v8 uint32 = v6
    for fi9 := 0; fi9 < 2; fi9++ {
        if (!(-665 >= 1)) {
            v10 := GenS2{F0: 32768, F1: 18446744073709551615, F2: -1}
            _ = v10
        } else {
            if (fi9 != (fi9 % (((fi9 - fi9)) | 1))) {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v6 = v8
                    _ = fi11
                }
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 uint16 = 255
                    if (false || (-2147483648 != -42)) {
                        switch (fi12 ^ -1) {
                        case 683:
                            var v14 float64 = -2.25
                            _ = v14
                        case -1:
                            v8 += v8
                        default:
                            var v15 uint16 = v13
                            _ = v15
                        }
                        v16 := []int{fi9, fi12, (fi9 & 128)}
                        _ = len(v16)
                        if (true || false) {
                            v3[v1] = v6
                            v16 = append(v16, (-(fi9)))
                        }
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
                v4 = 700
                for fi17 := 0; fi17 < 14; fi17++ {
                    if ("hello" < ((v1 + "\n\t") + v1)) {
                        switch v1 {
                        case "a":
                            v6 -= 482
                        case "\n\t":
                            switch v1 {
                            case "":
                                v4 += 4294967295
                            case "a":
                                v18 := GenS3{F0: 65535}
                                var v19 float32 = -2.25
                                _ = v18
                                _ = v19
                            case "hello":
                                fmt.Println(fi9, fi17)
                            default:
                                v3[v1] = v8
                                v6 = 65535
                                switch (fi17 & (fi9 & (fi17 ^ (fi9 / (((-(fi17))) | 1))))) {
                                case 709:
                                    if false {
                                        var v20 float32 = 0.1
                                        switch v1 {
                                        case "a":
                                            v21 := GenS4{F0: 42, F1: v1}
                                            _ = v21
                                        }
                                        v22 := GenS3{F0: 328}
                                        v23 := GenS4{F0: 0, F1: v1}
                                        _ = v20
                                        _ = v22
                                        _ = v23
                                    } else {
                                        fmt.Println(v1, fi17)
                                        v1 = ((v1 + "foo bar") + "")
                                    }
                                case 127:
                                    if false {
                                        v24 := GenS5{F0: -1}
                                        _ = v24
                                    } else {
                                        v3[v1] = 2147483648
                                        v6 += 65535
                                    }
                                case 0:
                                    v25 := []int{(fi17 >> ((v5) & 63)), (^(fi17)), (fi9 / (((fi17 / (((fi17 / ((fi9) | 1))) | 1))) | 1)), fi9}
                                    _ = v25
                                }
                            }
                        case "foo bar":
                            var v26 float64 = 3.14159
                            _ = v26
                        default:
                            fmt.Println(v6)
                        }
                        var v27 uint32 = v8
                        var v28 int64 = -954
                        v29 := []int{(fi17 ^ -151)}
                        if false {
                            if false {
                                v27 = 2147483648
                            }
                            v28 = v28
                        } else {
                            fmt.Println(v28)
                        }
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi17
                }
            }
            v5 -= uint(v6)
        }
        delete(v3, v1)
        _ = fi9
    }
    v30 := GenS6{F0: 2}
    if (true && ((v6 != (v8 + v8)) || (v2 == (v2 - (v2 + (^(v2))))))) {
        for fi31 := 0; fi31 < 14; fi31++ {
            if true {
                delete(v3, ("a" + ""))
                v32 := GenS7{F0: false}
                if true {
                    v3[v1] = (v8 * v6)
                    v2 += 100
                    v2 = (-(v2))
                }
                _ = v32
            } else {
                if (0 >= v5) {
                    var v33 uint32 = (v8 / ((2147483648) | 1))
                    _ = v33
                }
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 uint8 = 1
                    _ = fi34
                    _ = v35
                }
            }
            v36 := []int{fi31, fi31, (fi31 & (fi31 << (((^(v5))) & 63)))}
            _ = fi31
            _ = v36
        }
        var v37 float32 = (0.1 - 1.5)
        _ = v37
    } else {
        if (false && true) {
            v2 = 0
            v38 := GenS6{F0: 60}
            _ = v38
        }
    }
    if true {
        switch -61 {
        case 4294967295:
            v8 += v8
        case 2147483647:
            var v39 int8 = (v2 - v2)
            _ = v39
        case -141:
            var v40 int8 = 119
            _ = v40
        }
        v2 = v2
    }
    switch 42 {
    case 0:
        fmt.Println(v8, v2, v6)
    case -9223372036854775808:
        for fi41 := 0; fi41 < 9; fi41++ {
            v8 = v6
            v2 = (v2 + (^(v2)))
            _ = fi41
        }
    case -95:
        for fi42 := 0; fi42 < 12; fi42++ {
            for fi43 := 0; fi43 < 14; fi43++ {
                switch (fi42 | (fi42 - fi42)) {
                case 348:
                    v2 += (v2 % (((v2 + (^(v2)))) | 1))
                case 0:
                    switch 4294967295 {
                    case 127:
                        if false {
                            switch fi42 {
                            case 127:
                                var v44 float64 = float64(0.1)
                                _ = v44
                            case 32767:
                                var v45 string = ((v1 + (("\n\t" + "hello") + (((v1 + (v1 + v1)) + v1) + "hello"))) + v1)
                                _ = v45
                            case 32768:
                                if true {
                                    var v46 float64 = 0.1
                                    switch (fi42 ^ (fi42 & (fi43 & fi42))) {
                                    case -915:
                                        var v47 string = ("" + ((v1 + (v1 + ((v1 + (v1 + "foo bar")) + ("a" + "hello")))) + ("a" + v1)))
                                        _ = v47
                                    default:
                                        switch -304 {
                                        case 2147483647:
                                            delete(v3, "foo bar")
                                        default:
                                            v48 := GenS4{F0: 42, F1: v1}
                                            var v49 uint32 = v6
                                            var v50 uint64 = 4294967295
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    v8 += v8
                                    _ = v46
                                }
                                if (!true) {
                                    var v51 bool = (!true)
                                    if (!(v51 && (true && true))) {
                                        v52 := GenS8{F0: 504, F1: 4294967295, F2: (("hello" + v1) + v1)}
                                        var v53 uint8 = 42
                                        v54 := GenS1{F0: v51, F1: (v53 % ((v53) | 1)), F2: 65535}
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    v8 += 385
                                    _ = v51
                                } else {
                                    switch fi43 {
                                    case -940:
                                        var v55 float32 = 1.5
                                        _ = v55
                                    default:
                                        delete(v3, v1)
                                    }
                                    var v56 string = v1
                                    if true {
                                        if false {
                                            switch fi43 {
                                            case -552:
                                                v57 := []int{}
                                                _ = v57
                                            }
                                            fmt.Println(fi43, fi42)
                                        }
                                        if false {
                                            var v58 int = int(v4)
                                            v59 := []int{}
                                            var v60 int = fi43
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        } else {
                                            var v61 uint8 = 1
                                            if (945 >= 334) {
                                                var v62 uint32 = v8
                                                _ = v62
                                            }
                                            _ = v61
                                        }
                                    } else {
                                        delete(v3, "a")
                                    }
                                    _ = v56
                                }
                            }
                            var v63 int64 = 9223372036854775807
                            if false {
                                v2 = (v2 << (((v5 - (v5 + (^(v5))))) & 63))
                                v64 := GenS7{F0: false}
                                _ = v64
                            }
                            _ = v63
                        } else {
                            v2 += v2
                        }
                    case 42:
                        if (-1 == (fi43 >> (((v4 & 18446744073709551615)) & 63))) {
                            v6 += (v8 - 0)
                            fmt.Println(v6, v6, v8)
                            switch (fi42 * 32767) {
                            case -2147483648:
                                fmt.Println(fi42, v4)
                            default:
                                var v65 uint8 = 255
                                _ = v65
                            }
                        }
                    case 128:
                        if (2147483647 < -315) {
                            if ((v1 + v1) != (("hello" + "") + v1)) {
                                v1 = "hello"
                                v3[(((v1 + (("\n\t" + v1) + v1)) + "foo bar") + v1)] = v6
                            }
                            v6 = 65535
                            var v66 uint8 = 138
                            _ = v66
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v67 := []int{32767, fi42}
                                _ = v67
                            case "a":
                                var v68 uint8 = 128
                                _ = v68
                            }
                            v69 := []int{fi42, (fi42 - (fi42 ^ fi42)), fi42, (fi42 & (fi43 - -9223372036854775808))}
                            _ = v69
                        }
                    }
                default:
                    fmt.Println(v1, v1, fi42)
                }
                _ = fi43
            }
            v70 := []int{}
            _ = fi42
            _ = v70
        }
    }
    var v71 int32 = -1
    var v72 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v72
    _ = v8
}
