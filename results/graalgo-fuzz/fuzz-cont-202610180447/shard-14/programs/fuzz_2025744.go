package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int16; F1 int; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 string; F2 uint }
type GenS9 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 int32) uint16 {
    switch ("a" + ("" + (("hello" + (("foo bar" + "") + "hello")) + ("\n\t" + "hello")))) {
    case "hello":
        v6 := []int{32767, -1, 4294967295}
        _ = v6
    case "a":
        var v7 string = "\n\t"
        _ = v7
    case "":
        switch ("foo bar" + "foo bar") {
        case "a":
            if false {
                switch -1 {
                case 0:
                    p1 = p1
                case 32768:
                    v8 := GenS1{F0: 2147483647}
                    _ = v8
                }
            }
        case "hello":
            var v9 uint = 65535
            _ = v9
        default:
            v10 := []int{int(p1), -815}
            _ = v10
        }
    }
    var v11 int64 = -9223372036854775808
    v12 := GenS2{F0: 255}
    var v13 uint = 42
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []int{32767, 909, -855}
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 float64 = 3.14159
    var v14 uint64 = 53
    if true {
        var v15 int = 388
        var v16 float32 = 3.14159
        for fi17 := 0; fi17 < 1; fi17++ {
            switch 0 {
            case 42:
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v2, v16, v14)
                    _ = fi18
                }
            default:
                v2 = 1
            }
            v2 -= v2
            _ = fi17
        }
        _ = v15
        _ = v16
    } else {
        v5 = v5
        v5 -= v5
    }
    if (!false) {
        if true {
            v5 += v5
            v14 -= v14
        } else {
            v14 = (v14 + 2)
            if (42 != 42) {
                _ = len(v3)
                delete(v4, 9223372036854775807)
                switch "\n\t" {
                case "a":
                    if false {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi19
                        }
                        v20 := GenS1{F0: uint32(v2)}
                        if true {
                            switch 2147483647 {
                            case 841:
                                v21 := GenS1{F0: 2}
                                _ = v21
                            case 9223372036854775807:
                                v22 := GenS3{F0: -1, F1: 18446744073709551615}
                                _ = v22
                            default:
                                fmt.Println(v14, v5, v14)
                            }
                            if (((0.1 - (((-2.25 + v5) - 3.14159) + 3.14159)) - float64(1.5)) > v5) {
                                v23 := v3[0]
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    v14 -= v14
                                    _ = fi24
                                }
                                v25 := GenS3{F0: 9223372036854775807, F1: 42}
                                _ = v23
                                _ = v25
                            }
                        }
                        _ = v20
                    }
                }
            }
        }
        v2 = genFn0(100, 42)
    }
    fmt.Println(v14, v5)
    v26 := GenS4{F0: v14, F1: v14, F2: 42}
    switch 127 {
    case -1:
        v27 := GenS4{F0: (v14 / ((v14) | 1)), F1: v14, F2: 139}
        _ = v27
    case 431:
        var v28 uint16 = v2
        _ = v28
    }
    v29 := GenS5{F0: (("hello" + ((("\n\t" + "") + ("a" + "a")) + "foo bar")) + "\n\t"), F1: 19}
    if (!(!(true || true))) {
        if false {
            fmt.Println(v5, v2, v14)
        }
        var v30 float64 = (((float64((((((1e10 - float32((v5 + v5))) / 1.5) / float32(float64((((1e10 / 1.5) + 0.0) / 3.14159)))) + 1.5) * (((1e10 / 3.14159) + (-2.25 - -2.25)) * 1.5))) / 1e10) - float64(-2.25)) * (float64(0.0) + -2.25))
        _ = v30
    }
    _ = len(v1)
    if (1 <= 128) {
        var v31 uint16 = 42
        _ = v31
    }
    var v32 int64 = -9223372036854775808
    fmt.Println(v32, v2, v14)
    if false {
        v2 = (^(v2))
    } else {
        var v33 float32 = -2.25
        var v34 float32 = float32((v5 / float64(v33)))
        delete(v3, 911)
        _ = v33
        _ = v34
    }
    v35 := GenS6{F0: 42, F1: -2147483648, F2: 2147483647}
    if ((!true) && (false && true)) {
        var v36 int16 = 0
        for fi37 := 0; fi37 < 3; fi37++ {
            var v38 bool = (v2 == v2)
            switch "\n\t" {
            case "a":
                fmt.Println(v32, fi37, v5)
            default:
                for fi39 := 0; fi39 < 19; fi39++ {
                    switch (^(fi39)) {
                    case 0:
                        v32 += v32
                    case -375:
                        fmt.Println(v2, v2, v32)
                        v40 := GenS5{F0: "hello", F1: 842}
                        _ = v40
                    case -23:
                        v14 += (v14 | (v14 & v14))
                    default:
                        switch 4294967295 {
                        case 32767:
                            var v41 uint16 = 669
                            _ = v41
                        case 0:
                            if (!v38) {
                                if v38 {
                                    if v38 {
                                        v14 += v14
                                        if v38 {
                                            if v38 {
                                                switch (fi37 & (fi37 & -1)) {
                                                case 9223372036854775807:
                                                    fmt.Println(v36)
                                                case 128:
                                                    _ = len(v1)
                                                    v5 -= v5
                                                    var v42 float32 = float32(0.1)
                                                    _ = v42
                                                case 32767:
                                                    if v38 {
                                                        var v43 float64 = v5
                                                        fmt.Println(fi37, fi39)
                                                        var v44 int32 = 2147483647
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        var v45 string = ("a" + "")
                                                        v45 = ("hello" + ((("foo bar" + "hello") + ((v45 + v45) + v45)) + "hello"))
                                                        v1 = append(v1, fi39)
                                                        _ = v45
                                                    }
                                                }
                                                var v46 uint16 = 895
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = fi37
                                                    }
                                                    v36 = 32767
                                                }
                                                _ = v46
                                            }
                                            v14 += 107
                                            v47 := GenS5{F0: "", F1: 1}
                                            v48 := GenS6{F0: -1, F1: fi39, F2: fi37}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            fmt.Println(v5, fi37, v32)
                                            var v49 uint8 = 42
                                            v50 := GenS3{F0: (-(v32)), F1: 681}
                                            _ = v49
                                            _ = v50
                                        }
                                        v51 := GenS7{F0: 314}
                                        _ = v51
                                    } else {
                                        var v52 int64 = (v32 >> ((42) & 63))
                                        if v38 {
                                            v36 += (v36 | (v36 & (-(v36))))
                                        } else {
                                            switch "" {
                                            case "\n\t":
                                                delete(v3, 2147483648)
                                            }
                                            fmt.Println(v36, fi39)
                                        }
                                        var v53 int64 = int64(v36)
                                        _ = v52
                                        _ = v53
                                    }
                                    v14 -= v14
                                }
                            } else {
                                v36 += 1
                                fmt.Println(v36, fi39, v14)
                            }
                        }
                    }
                    if (!v38) {
                        var v54 int32 = -349
                        _ = v54
                    } else {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi37, fi37, v32)
                        case "hello":
                            var v55 bool = ("foo bar" <= "a")
                            _ = v55
                        case "a":
                            var v56 bool = v38
                            _ = v56
                        default:
                            switch (fi37 / ((-335) | 1)) {
                            case -532:
                                _ = len(v4)
                            }
                        }
                        if (false || v38) {
                            v57 := GenS8{F0: ((("a" + "") + "") + "foo bar"), F1: ("\n\t" + "foo bar"), F2: uint(v2)}
                            v58 := GenS9{F0: v38, F1: 0}
                            v38 = false
                            _ = v57
                            _ = v58
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi39
                }
            }
            v1 = append(v1, fi37)
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    var v59 float32 = (((0.0 * float32(v5)) - 0.0) * 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v59
}
