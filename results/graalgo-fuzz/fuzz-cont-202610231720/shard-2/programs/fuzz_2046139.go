package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int64; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int64 }
func genFn0() uint64 {
    var v7 int8 = 42
    v7 += (v7 >> ((1) & 63))
    var v8 int8 = (^(v7))
    var v9 int16 = 42
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64, p1 uint64, p2 uint8) uint32 {
    var v10 int = -2147483648
    v11 := []int{int(p2), (v10 & (v10 ^ -212)), int(p0), v10}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 187
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := make(map[int]uint32)
    v1 += v1
    for fi12 := 0; fi12 < 7; fi12++ {
        switch ("\n\t" + ("" + "hello")) {
        case "":
            v13 := []int{(fi12 + (fi12 | (fi12 << ((565) & 63)))), -2147483648}
            _ = v13
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: 0, F1: ""}
                var v16 int32 = 2147483647
                if ((1.5 / (1.5 + 3.14159)) < float32(float64(1e10))) {
                    switch "hello" {
                    case "foo bar":
                        v1 -= (v1 >> ((255) & 63))
                    case "a":
                        v1 -= 0
                    }
                    delete(v4, (fi14 ^ (fi14 >> ((114) & 63))))
                }
                if false {
                    v16 = -648
                    v17 := GenS2{F0: v16, F1: 4294967295}
                    switch (fi14 - (-(fi14))) {
                    case 127:
                        if (65535 >= 920) {
                            if (65535 > 21) {
                                v18 := GenS2{F0: v16, F1: 42}
                                fmt.Println(fi14, fi12)
                                var v19 uint16 = 255
                                _ = v18
                                _ = v19
                            }
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                switch (fi14 / ((-693) | 1)) {
                                case 127:
                                    v20 := []int{(fi14 | -1), fi14, -158, fi14}
                                    _ = v20
                                case 32767:
                                    if true {
                                        switch "" {
                                        case "":
                                            v1 -= 42
                                        case "hello":
                                            var v21 uint8 = 200
                                            _ = v21
                                        case "a":
                                            v22 := GenS1{F0: -1, F1: ""}
                                            _ = v22
                                        }
                                        v16 = (v16 ^ -1)
                                        if true {
                                            var v23 uint16 = 851
                                            _ = v23
                                        } else {
                                            v24 := []int{}
                                            switch ("\n\t" + "foo bar") {
                                            case "hello":
                                                v25 := GenS2{F0: 42, F1: 4294967295}
                                                _ = v25
                                            case "foo bar":
                                                delete(v6, 1)
                                            }
                                            var v26 int16 = 0
                                            _ = v24
                                            _ = v26
                                        }
                                    } else {
                                        if false {
                                            var v27 int64 = 833
                                            fmt.Println(v27, fi14)
                                            switch 4294967295 {
                                            case 32768:
                                                v1 -= v1
                                            case 32767:
                                                var v28 uint64 = 2
                                                _ = v28
                                            }
                                            _ = v27
                                        }
                                    }
                                case -423:
                                    fmt.Println(fi12, v16)
                                default:
                                    v29 := GenS2{F0: v16, F1: 18446744073709551615}
                                    _ = v29
                                }
                            case "\n\t":
                                var v30 float32 = 0.1
                                _ = v30
                            case "a":
                                switch -9223372036854775808 {
                                case 44:
                                    switch ("a" + ("\n\t" + "\n\t")) {
                                    case "a":
                                        v31 := GenS2{F0: v16, F1: genFn0()}
                                        _ = v31
                                    case "":
                                        v32 := GenS2{F0: 1, F1: 0}
                                        _ = v32
                                    default:
                                        if false {
                                            var v33 uint64 = 1
                                            delete(v6, fi14)
                                            v34 := GenS3{F0: false, F1: 4294967295, F2: "foo bar"}
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v16 = v16
                                        }
                                    }
                                case 32767:
                                    if ((fi12 & fi12) <= fi12) {
                                        var v35 int8 = 1
                                        if ((^(v1)) < (v1 << ((65535) & 63))) {
                                            var v36 int64 = -9223372036854775808
                                            _ = v36
                                        } else {
                                            var v37 float32 = -2.25
                                            _ = v37
                                        }
                                        fmt.Println(v35, v1, fi14)
                                        v35 += -1
                                        v16 = (v16 % ((-2147483648) | 1))
                                        _ = v35
                                    }
                                }
                            default:
                                delete(v3, "hello")
                                v1 = v1
                                v1 += 32767
                            }
                            v38 := GenS3{F0: (((!true) || (!(((v16 << ((65535) & 63)) > (v16 >> ((255) & 63))) || (fi14 > fi12)))) || (4294967295 >= 18446744073709551615)), F1: 1, F2: "a"}
                            _ = v38
                        }
                    case -15:
                        v39 := GenS4{F0: 1, F1: v16}
                        _ = v39
                    case -2147483648:
                        v40 := GenS5{F0: v16, F1: 128, F2: 2}
                        _ = v40
                    }
                    _ = v17
                } else {
                    v1 = v1
                    var v41 int16 = v1
                    _ = v41
                }
                v1 = (v1 - v1)
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        if false {
            var v42 uint8 = 79
            v1 = 1
            _ = v42
        }
        _ = fi12
    }
    var v43 uint64 = 2
    v44 := []int{0, 2147483647}
    var v45 int64 = -168
    var v46 int8 = -128
    for fi47 := 0; fi47 < 5; fi47++ {
        switch "foo bar" {
        case "foo bar":
            var v48 string = "a"
            _ = v48
        case "\n\t":
            for fi49 := 0; fi49 < 5; fi49++ {
                if len(v44) > 0 {
                    v44[0] = int(v1)
                }
                var v50 int16 = (v1 ^ v1)
                _ = fi49
                _ = v50
            }
        case "a":
            v51 := GenS4{F0: (v45 / ((v45) | 1)), F1: -172}
            _ = v51
        }
        switch "hello" {
        case "\n\t":
            switch (fi47 + 32767) {
            case 90:
                for fi52 := 0; fi52 < 15; fi52++ {
                    if false {
                        v1 += v1
                    }
                    fmt.Println(v46)
                    _ = fi52
                }
            case 936:
                v43 = v43
            case -2147483648:
                v43 -= 309
            default:
                if (2 > 65535) {
                    v46 = (v46 >> ((1) & 63))
                    v53 := GenS1{F0: (v45 / ((int64(v46)) | 1)), F1: ("\n\t" + "hello")}
                    _ = v53
                }
            }
        case "a":
            for fi54 := 0; fi54 < 3; fi54++ {
                var v55 uint16 = 65535
                _ = fi54
                _ = v55
            }
        }
        _ = fi47
    }
    v56 := GenS3{F0: true, F1: 2, F2: ("a" + "hello")}
    for fi57 := 0; fi57 < 12; fi57++ {
        if (-2.25 < 0.0) {
            v58 := GenS6{F0: v45, F1: false, F2: genFn0()}
            for fi59 := 0; fi59 < 7; fi59++ {
                var v60 int8 = (-(v46))
                switch 9223372036854775807 {
                case 9223372036854775807:
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                default:
                    if false {
                        v61 := GenS5{F0: 42, F1: 217, F2: 65535}
                        _ = v61
                    }
                }
                _ = fi59
                _ = v60
            }
            var v62 uint = 276
            _ = v58
            _ = v62
        }
        _ = len(v4)
        _ = fi57
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        v45 -= v45
        var v64 uint8 = 255
        v45 = (v45 | (v45 >> ((0) & 63)))
        _ = fi63
        _ = v64
    }
    v65 := GenS7{F0: genFn0(), F1: (-(v46))}
    if false {
        if (0 < int(v45)) {
            var v66 int32 = -2147483648
            v43 = 0
            _ = v66
        }
        fmt.Println(v45)
    } else {
        if (true || false) {
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
        } else {
            var v67 int16 = -79
            v68 := GenS3{F0: true, F1: 602, F2: (("hello" + "foo bar") + ("\n\t" + ("a" + ("foo bar" + "a"))))}
            _ = v67
            _ = v68
        }
        switch "a" {
        case "a":
            v69 := GenS8{F0: v45}
            _ = v69
        case "foo bar":
            v2["a"] = -984
        default:
            v1 += (v1 * 1)
        }
        var v70 int64 = (v45 << ((255) & 63))
        _ = v70
    }
    if (42 == 1) {
        _ = len(v6)
        v43 = (^(v43))
    } else {
        v44 = append(v44, 9223372036854775807)
        v43 -= genFn0()
    }
    if (((-2.25 / -2.25) - (1.5 - 3.14159)) > 3.14159) {
        var v71 int64 = v45
        _ = v71
    } else {
        var v72 uint = 18446744073709551615
        _ = len(v5)
        if true {
            v45 = 441
            if ("a" > "a") {
                v1 = v1
            } else {
                var v73 uint64 = v43
                var v74 bool = (2 < uint32(v73))
                _ = v73
                _ = v74
            }
            if false {
                if true {
                    v46 = (v46 - v46)
                    var v75 float32 = 0.1
                    var v76 bool = (false || ((v72 / ((v72) | 1)) > (v72 ^ v72)))
                    _ = v75
                    _ = v76
                } else {
                    var v77 uint64 = 249
                    v78 := GenS6{F0: v45, F1: (false && (true || (((false || (!(434 >= 457))) && (v45 > v45)) || (392 >= -1)))), F2: (v43 & 494)}
                    var v79 int8 = (-(v46))
                    _ = v77
                    _ = v78
                    _ = v79
                }
            } else {
                var v80 uint64 = genFn0()
                _ = v80
            }
        } else {
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
        }
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v56
    _ = v6
    _ = v65
}
