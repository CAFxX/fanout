package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 int; F2 uint8 }
func genFn0(p0 bool, p1 string, p2 int) int32 {
    var v4 uint64 = 18446744073709551615
    var v5 int64 = -1
    var v6 float32 = ((3.14159 * (0.1 - -2.25)) + 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 float64 = (-2.25 / ((-2.25 * -2.25) - (0.0 + 3.14159)))
    var v7 float32 = (1.5 + -2.25)
    fmt.Println(v7, v7)
    v8 := GenS1{F0: 4294967295, F1: (2147483648 >= 42)}
    switch "a" {
    case "foo bar":
        var v9 bool = (2 > 258)
        _ = v9
    case "hello":
        if true {
            fmt.Println(v3, v3)
        } else {
            var v10 int8 = 42
            v10 = (v10 % (((v10 & (v10 | 100))) | 1))
            _ = v10
        }
    default:
        v11 := GenS1{F0: 2147483647, F1: (255 > 18446744073709551615)}
        _ = v11
    }
    v12 := GenS1{F0: 379, F1: false}
    v13 := GenS1{F0: 0, F1: true}
    var v14 int32 = genFn0(true, "foo bar", 32768)
    var v15 uint16 = 29
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint8 = 0
        var v18 int = fi16
        for fi19 := 0; fi19 < 5; fi19++ {
            v14 -= v14
            if false {
                v20 := GenS2{F0: 893, F1: 255}
                var v21 uint32 = 4294967295
                _ = v20
                _ = v21
            } else {
                if false {
                    v22 := []int{-159}
                    _ = v22
                } else {
                    v14 += v14
                    v23 := GenS3{F0: v17, F1: 945}
                    v3 = float64((1e10 - v7))
                    _ = v23
                }
                var v24 int16 = 1000
                var v25 int8 = -93
                v3 = v3
                v1["foo bar"] = ("foo bar" + "hello")
                _ = v24
                _ = v25
            }
            v1["a"] = "foo bar"
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v7)
    var v26 int32 = v14
    v27 := GenS2{F0: 1, F1: 1}
    if (true || (!(96 == 34))) {
        if false {
            switch "\n\t" {
            case "":
                v28 := GenS4{F0: "hello", F1: v26}
                _ = v28
            case "a":
                delete(v1, "hello")
            case "hello":
                if (genFn0(false, "", int(v26)) == genFn0(((18446744073709551615 >= 42) || false), "\n\t", 344)) {
                    v29 := GenS5{F0: int8(v26), F1: 0}
                    v7 = float32((v3 - v3))
                    switch 0 {
                    case -2147483648:
                        var v30 uint8 = 1
                        _ = v30
                    case 140:
                        v31 := GenS6{F0: 170}
                        _ = v31
                    case 0:
                        fmt.Println(v26)
                    default:
                        var v32 uint32 = 2
                        _ = v32
                    }
                    _ = v29
                }
            }
            for fi33 := 0; fi33 < 5; fi33++ {
                v15 -= v15
                for fi34 := 0; fi34 < 11; fi34++ {
                    if false {
                        if (1 > 1) {
                            if (true || true) {
                                v35 := GenS7{F0: 18446744073709551615, F1: 101, F2: 127}
                                _ = v35
                            } else {
                                var v36 string = ((("" + "\n\t") + "hello") + "foo bar")
                                delete(v2, v36)
                                var v37 uint16 = (v15 * v15)
                                _ = v36
                                _ = v37
                            }
                        } else {
                            v15 += 937
                            var v38 int16 = 32767
                            if true {
                                var v39 int16 = v38
                                _ = v39
                            }
                            _ = v38
                        }
                        if false {
                            v40 := []int{32767, fi34, (fi33 & fi33)}
                            v2[(("a" + (("hello" + (("hello" + "") + "a")) + ("hello" + (("foo bar" + ((("\n\t" + "hello") + "\n\t") + "\n\t")) + "\n\t")))) + ((("a" + ("\n\t" + ("a" + "\n\t"))) + ("" + "foo bar")) + "hello"))] = (!true)
                            fmt.Println(v14, v26)
                            if false {
                                if true {
                                    v26 -= v26
                                }
                            } else {
                                v41 := GenS2{F0: 65535, F1: 2}
                                _ = v41
                            }
                            v7 -= v7
                            _ = v40
                        } else {
                            var v42 uint16 = (v15 & 42)
                            fmt.Println(v3, v15, v26)
                            _ = v42
                        }
                        v43 := GenS2{F0: 42, F1: 1}
                        _ = v43
                    } else {
                        v2["foo bar"] = false
                        v44 := GenS7{F0: 0, F1: (fi33 << ((42) & 63)), F2: 2}
                        _ = v44
                    }
                    _ = fi34
                }
                _ = fi33
            }
        } else {
            _ = len(v1)
            fmt.Println(v7, v15)
        }
        v45 := GenS3{F0: 1, F1: 779}
        _ = v45
    } else {
        switch -9223372036854775808 {
        case 1:
            var v46 uint32 = 4294967295
            _ = v46
        case 0:
            switch 127 {
            case 32768:
                var v47 uint = 559
                _ = v47
            }
        case 127:
            if (4294967295 < 837) {
                for fi48 := 0; fi48 < 6; fi48++ {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        _ = len(v1)
                        if true {
                            var v50 int32 = 2147483647
                            _ = v50
                        }
                        var v51 uint8 = 131
                        _ = fi49
                        _ = v51
                    }
                    var v52 uint32 = 1
                    _ = fi48
                    _ = v52
                }
            } else {
                v53 := GenS5{F0: -128, F1: -261}
                if false {
                    var v54 int16 = -146
                    switch int(v54) {
                    case 1:
                        for fi55 := 0; fi55 < 7; fi55++ {
                            var v56 uint = 255
                            _ = fi55
                            _ = v56
                        }
                    case -583:
                        for fi57 := 0; fi57 < 12; fi57++ {
                            var v58 uint = 2
                            var v59 int = (fi57 ^ (fi57 / ((fi57) | 1)))
                            v60 := []int{int(v58), (fi57 ^ fi57), (-(fi57)), -497}
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                    case -347:
                        var v61 bool = (!(((v7 * float32(v3)) + 0.1) == v7))
                        _ = v61
                    default:
                        v62 := []int{4294967295}
                        _ = v62
                    }
                    _ = v54
                }
                for fi63 := 0; fi63 < 16; fi63++ {
                    v64 := GenS2{F0: 855, F1: 288}
                    _ = fi63
                    _ = v64
                }
                _ = v53
            }
        default:
            if false {
                for fi65 := 0; fi65 < 10; fi65++ {
                    v15 = 80
                    _ = fi65
                }
            }
        }
    }
    v66 := GenS7{F0: 0, F1: 128, F2: 255}
    if true {
        v15 = (^(v15))
        if (2 > (v15 / ((309) | 1))) {
            v7 -= 0.0
            var v67 int64 = 999
            _ = v67
        }
    } else {
        fmt.Println(v15, v7)
        v7 = 1e10
        if false {
            for fi68 := 0; fi68 < 15; fi68++ {
                var v69 float64 = ((1.5 * ((0.0 - (v3 + (v3 + 0.1))) - (v3 - (float64(3.14159) * 0.0)))) + 0.0)
                fmt.Println(v14)
                _ = fi68
                _ = v69
            }
            var v70 uint8 = 127
            _ = v70
        } else {
            v71 := []int{1, -9223372036854775808, 9223372036854775807, -1}
            _ = v71
        }
    }
    _ = len(v2)
    var v72 int32 = (v14 - -151)
    v72 = (v14 / (((-(v14))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v66
    _ = v7
    _ = v72
    _ = v8
}
