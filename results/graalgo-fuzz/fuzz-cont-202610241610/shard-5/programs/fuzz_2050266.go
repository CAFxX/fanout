package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 bool }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 uint8, p2 int64) uint8 {
    v6 := GenS1{F0: (((("foo bar" + "a") + "") + "hello") + ("hello" + (("hello" + "\n\t") + ("" + (("hello" + "\n\t") + "hello"))))), F1: 53}
    v7 := GenS1{F0: ("" + "foo bar"), F1: 0}
    if (!false) {
        fmt.Println(p0)
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 = p0
            p2 += p2
            var v9 uint = 255
            _ = fi8
            _ = v9
        }
        if p0 {
            v10 := []int{}
            v11 := GenS2{F0: p2, F1: 1, F2: -2147483648}
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 float32 = ((3.14159 * ((float32(-2.25) + 1.5) / (3.14159 + 3.14159))) / float32(1e10))
                if p0 {
                    if (1000 >= 96) {
                        p2 += (p2 ^ p2)
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v15 := GenS1{F0: "", F1: 753}
                            p2 += p2
                            v10 = append(v10, (fi12 / (((^(fi12))) | 1)))
                            if (((1e10 * 1.5) - float64(-2.25)) <= 0.1) {
                                if (4294967295 < 2) {
                                    v16 := GenS3{F0: (true || p0), F1: 141, F2: (1e10 < (3.14159 + 3.14159))}
                                    _ = v16
                                }
                                p2 -= (p2 % ((p2) | 1))
                            }
                            _ = fi14
                            _ = v15
                        }
                        switch (fi12 ^ (-(fi12))) {
                        case -1:
                            p2 += (p2 & (-(p2)))
                        case 9223372036854775807:
                            switch ("" + (("foo bar" + "foo bar") + "a")) {
                            case "foo bar":
                                var v17 int8 = 0
                                _ = v17
                            case "\n\t":
                                var v18 float32 = float32(-2.25)
                                _ = v18
                            case "":
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            default:
                                var v19 float32 = (v13 + 0.0)
                                _ = v19
                            }
                        default:
                            p1 += 42
                        }
                    } else {
                        if p0 {
                            v20 := GenS3{F0: (!false), F1: 255, F2: true}
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 18; fi21++ {
                                var v22 bool = true
                                if v22 {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    fmt.Println(fi21)
                                } else {
                                    p2 += 9223372036854775807
                                    v23 := GenS4{F0: 2, F1: p2}
                                    v22 = (42 != p2)
                                    _ = v23
                                }
                                v24 := GenS2{F0: p2, F1: 219, F2: -2147483648}
                                _ = fi21
                                _ = v22
                                _ = v24
                            }
                            p0 = false
                            for fi25 := 0; fi25 < 16; fi25++ {
                                switch "a" {
                                case "\n\t":
                                    if (1 < 42) {
                                        if false {
                                            fmt.Println(p1, p0, p1)
                                        } else {
                                            fmt.Println(v13, v13, v13)
                                            var v26 string = "\n\t"
                                            v27 := GenS2{F0: (p2 % (((p2 / (((p2 ^ p2)) | 1))) | 1)), F1: 20, F2: 42}
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                case "":
                                    v28 := GenS3{F0: ((((0.1 - 0.1) > (-2.25 - 1.5)) || false) || false), F1: 156, F2: (p1 < p1)}
                                    v29 := GenS4{F0: 265, F1: 1}
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi25
                            }
                        }
                    }
                    fmt.Println(fi12, fi12)
                } else {
                    p2 = 157
                    fmt.Println(p2)
                }
                if (-2.25 <= v13) {
                    _ = len(v10)
                } else {
                    switch -1 {
                    case 9223372036854775807:
                        v30 := GenS5{F0: (p1 ^ 161)}
                        _ = v30
                    case 4294967295:
                        p1 = (p1 / (((p1 & p1)) | 1))
                    case 789:
                        var v31 int16 = 1000
                        _ = v31
                    default:
                        p0 = true
                    }
                    fmt.Println(p0, p1)
                    p2 = (p2 & p2)
                }
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 int = 32768
    var v3 int16 = 1000
    v4 := make(map[string]bool)
    var v5 int16 = (v3 % (((^(v3))) | 1))
    var v32 uint16 = 805
    v3 = (v5 | v3)
    if (177 <= 723) {
        if true {
            v4["hello"] = (127 > 138)
            var v33 uint = 2
            if (3.14159 <= (3.14159 / (0.0 + 0.1))) {
                v33 -= 42
                var v34 uint64 = 471
                _ = v34
            } else {
                switch ("\n\t" + "a") {
                case "\n\t":
                    v35 := []int{125}
                    _ = v35
                case "":
                    if false {
                        var v36 int16 = v5
                        var v37 int32 = int32(v2)
                        if (0 >= genFn0((v32 < 255), 128, 4)) {
                            var v38 int64 = 9223372036854775807
                            var v39 string = ("a" + (("\n\t" + ("a" + "\n\t")) + "a"))
                            var v40 uint64 = 4294967295
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            var v41 int32 = v37
                            _ = v41
                        }
                        _ = v36
                        _ = v37
                    } else {
                        for fi42 := 0; fi42 < 3; fi42++ {
                            v43 := []int{(v2 / (((v2 - (v2 << ((v33) & 63)))) | 1)), (^(v2)), fi42, fi42}
                            var v44 uint16 = (v32 / ((v32) | 1))
                            v44 -= v44
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        var v45 uint = (v33 + 0)
                        _ = v45
                    }
                }
                v32 += v32
                for fi46 := 0; fi46 < 12; fi46++ {
                    if ((567 != 0) && ((3.14159 != float64(1e10)) || false)) {
                        var v47 uint64 = 1
                        var v48 int = fi46
                        _ = v47
                        _ = v48
                    }
                    v49 := GenS6{F0: 81, F1: (("a" + "a") + "a"), F2: 65535}
                    for fi50 := 0; fi50 < 13; fi50++ {
                        v51 := []int{fi46, fi46}
                        v5 += (-(v5))
                        _ = fi50
                        _ = v51
                    }
                    _ = fi46
                    _ = v49
                }
            }
            _ = v33
        }
        switch -1 {
        case 4294967295:
            v1 = (v1 & (v1 / ((v1) | 1)))
        default:
            v52 := GenS6{F0: 253, F1: "a", F2: 2}
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        var v54 string = "a"
        switch (v2 | (v2 - fi53)) {
        case 128:
            for fi55 := 0; fi55 < 4; fi55++ {
                v56 := v4[v54]
                fmt.Println(v3, v3, v54)
                v56 = (v5 <= v3)
                _ = fi55
                _ = v56
            }
        case -9223372036854775808:
            fmt.Println(v54, fi53, v3)
        default:
            v57 := []int{(v2 ^ fi53), -115, (^(v2)), v2}
            _ = v57
        }
        var v58 int16 = v5
        _ = fi53
        _ = v54
        _ = v58
    }
    var v59 float32 = 3.14159
    switch v2 {
    case 2147483647:
        for fi60 := 0; fi60 < 4; fi60++ {
            for fi61 := 0; fi61 < 10; fi61++ {
                v62 := v4["a"]
                _ = fi61
                _ = v62
            }
            _ = fi60
        }
    default:
        v63 := []int{v2, v2, -9223372036854775808, v2}
        _ = v63
    }
    v64 := GenS6{F0: 128, F1: "foo bar", F2: v32}
    v65 := v4["foo bar"]
    if (65535 >= 424) {
        v1 = (-(v1))
        v66 := []int{32768, -2147483648, (v2 << ((1) & 63))}
        switch 0 {
        case 1:
            v67 := GenS6{F0: 255, F1: "foo bar", F2: 502}
            _ = v67
        case 32768:
            if true {
                var v68 int32 = -543
                switch "a" {
                case "a":
                    fmt.Println(v32, v2, v65)
                }
                fmt.Println(v3, v32, v3)
                _ = v68
            }
        }
        _ = v66
    } else {
        var v69 bool = (true || (!(85 == 756)))
        v70 := GenS5{F0: 127}
        if v69 {
            var v71 float32 = 1.5
            var v72 float32 = v71
            var v73 uint16 = (v32 & (v32 | v32))
            _ = v71
            _ = v72
            _ = v73
        }
        _ = v69
        _ = v70
    }
    switch v2 {
    case 128:
        v74 := GenS4{F0: 42, F1: 0}
        _ = v74
    case -9223372036854775808:
        var v75 uint64 = 42
        _ = v75
    case 127:
        var v76 string = "hello"
        _ = v76
    }
    v77 := GenS1{F0: "", F1: v32}
    var v78 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v65:", v65)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v59
    _ = v64
    _ = v65
    _ = v77
    _ = v78
}
