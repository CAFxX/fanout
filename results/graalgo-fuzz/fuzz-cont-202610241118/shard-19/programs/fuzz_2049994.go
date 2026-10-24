package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 string; F1 int16; F2 int8 }
type GenS6 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0(p0 uint, p1 uint16, p2 uint8) uint32 {
    var v5 int = -871
    fmt.Println(p0)
    v6 := []int{(v5 % (((v5 - v5)) | 1)), 0, v5}
    if (((1.5 != 1e10) && false) || true) {
        v5 += v5
    } else {
        if (!false) {
            v5 = int(p2)
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(p2, p0)
            v6 = append(v6, fi7)
            p0 = (p0 >> ((p0) & 63))
            _ = fi7
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                v5 = -830
                p2 = (^(p2))
            }
            _ = fi8
        }
    }
    v9 := GenS1{F0: 2, F1: false, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v9
    return 108
}
func genFn1(p0 uint16, p1 bool) int16 {
    for fi10 := 0; fi10 < 11; fi10++ {
        if (true || true) {
            var v11 string = ""
            p0 = p0
            v11 = v11
            _ = v11
        } else {
            if (75 == 145) {
                v12 := GenS2{F0: false, F1: 0}
                var v13 string = ""
                var v14 int = fi10
                _ = v12
                _ = v13
                _ = v14
            }
            if p1 {
                var v15 uint = uint(p0)
                if (-975 <= fi10) {
                    var v16 uint16 = (p0 / (((p0 / ((p0) | 1))) | 1))
                    fmt.Println(v16, p0)
                    _ = v16
                } else {
                    v15 -= 2
                }
                _ = v15
            }
            var v17 int32 = -2147483648
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            p0 += p0
            v19 := GenS2{F0: ((-(fi10)) != 127), F1: 32767}
            _ = fi18
            _ = v19
        }
        v20 := GenS2{F0: (21 != int8(fi10)), F1: 1000}
        _ = fi10
        _ = v20
    }
    var v21 bool = (3.14159 != 1.5)
    _ = p0
    _ = p1
    _ = v21
    return 681
}

func main() {
    var v1 uint8 = 26
    v2 := make(map[string]int)
    var v3 string = ("a" + (("foo bar" + (("hello" + "") + "")) + "hello"))
    v4 := make(map[int]uint32)
    switch v3 {
    case "hello":
        v22 := GenS2{F0: (0 >= 127), F1: 795}
        _ = v22
    case "a":
        if true {
            v1 += 42
            for fi23 := 0; fi23 < 16; fi23++ {
                v1 -= 24
                if true {
                    var v24 int64 = 42
                    _ = v24
                }
                _ = fi23
            }
        }
    case "":
        switch 412 {
        case 4294967295:
            v25 := GenS3{F0: 2, F1: -2147483648}
            _ = v25
        default:
            v26 := GenS2{F0: (false && false), F1: 772}
            _ = v26
        }
    }
    v27 := GenS1{F0: 0, F1: false, F2: 0}
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 uint = 4294967295
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v3)
    if false {
        v30 := v2[(v3 + v3)]
        _ = v30
    } else {
        var v31 bool = true
        for fi32 := 0; fi32 < 4; fi32++ {
            fmt.Println(v1, v1)
            if v31 {
                fmt.Println(fi32, v1)
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 uint8 = 0
                    _ = fi33
                    _ = v34
                }
            } else {
                fmt.Println(fi32, v1, v1)
                v35 := GenS3{F0: v1, F1: -1}
                _ = v35
            }
            _ = fi32
        }
        _ = v31
    }
    if true {
        switch 42 {
        case 32768:
            v36 := GenS4{F0: 2147483647, F1: 129, F2: 42}
            _ = v36
        case -1:
            v3 = ((((v3 + ("foo bar" + "a")) + (v3 + (v3 + ("foo bar" + v3)))) + v3) + "hello")
        case 791:
            v3 = (v3 + "a")
        default:
            v37 := GenS5{F0: v3, F1: 0, F2: 0}
            _ = v37
        }
        for fi38 := 0; fi38 < 18; fi38++ {
            if ((((v3 + ("a" + ((v3 + v3) + ((v3 + (("a" + ("\n\t" + (((v3 + "hello") + v3) + v3))) + v3)) + ((v3 + "hello") + v3))))) + "hello") == ((v3 + "foo bar") + v3)) || (91 < 0)) {
                var v39 uint32 = 335
                v3 = (v3 + "")
                _ = v39
            } else {
                v40 := GenS5{F0: "\n\t", F1: genFn1(255, true), F2: 0}
                switch v3 {
                case "":
                    fmt.Println(fi38)
                case "a":
                    var v41 string = ""
                    _ = v41
                default:
                    v3 = ("foo bar" + ("\n\t" + v3))
                }
                _ = v40
            }
            _ = fi38
        }
        if true {
            if (v3 <= ("\n\t" + v3)) {
                var v42 string = "foo bar"
                v3 = "hello"
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 int32 = -377
                    for fi45 := 0; fi45 < 10; fi45++ {
                        v46 := []int{}
                        v47 := v4[-9223372036854775808]
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    _ = fi43
                    _ = v44
                }
                _ = v42
            } else {
                var v48 uint8 = v1
                _ = v48
            }
            delete(v4, 32767)
            var v49 float32 = (1.5 - ((1.5 / (0.1 / 1e10)) - 0.1))
            _ = v49
        } else {
            fmt.Println(v3)
        }
    } else {
        if true {
            if true {
                if false {
                    var v50 int64 = -9223372036854775808
                    if (769 < 1) {
                        switch "\n\t" {
                        case "hello":
                            var v51 uint8 = (^(v1))
                            _ = v51
                        }
                    }
                    if ((1 <= -1) && true) {
                        var v52 bool = true
                        _ = v52
                    } else {
                        fmt.Println(v1, v1)
                        var v53 int16 = 32767
                        _ = v53
                    }
                    _ = v50
                } else {
                    _ = len(v4)
                    v54 := []int{32767, 1}
                    _ = v54
                }
                v55 := []int{}
                var v56 uint64 = 949
                _ = v55
                _ = v56
            }
            fmt.Println(v1, v3)
            v4[32767] = 65535
        } else {
            if false {
                fmt.Println(v3, v1)
                for fi57 := 0; fi57 < 7; fi57++ {
                    if (!false) {
                        var v58 string = "\n\t"
                        var v59 float64 = float64(((-2.25 * (1e10 + 1.5)) * (3.14159 / (1.5 + 3.14159))))
                        _ = v58
                        _ = v59
                    }
                    switch ("\n\t" + ((v3 + ("foo bar" + "")) + ("foo bar" + ("a" + (((v3 + v3) + ("\n\t" + "foo bar")) + (("\n\t" + (v3 + "")) + "")))))) {
                    case "hello":
                        v60 := GenS5{F0: v3, F1: 42, F2: -1}
                        _ = v60
                    case "foo bar":
                        v61 := GenS5{F0: (v3 + (("foo bar" + (((v3 + v3) + (v3 + (v3 + ((((v3 + v3) + (((v3 + (v3 + v3)) + v3) + v3)) + v3) + "foo bar")))) + v3)) + "\n\t")), F1: 1, F2: int8(v1)}
                        _ = v61
                    }
                    var v62 int32 = 36
                    _ = fi57
                    _ = v62
                }
            }
            var v63 uint32 = 2147483647
            _ = v63
        }
        v3 = v3
        v1 = (v1 - (v1 >> ((4294967295) & 63)))
    }
    v1 = 127
    if false {
        _ = len(v4)
    } else {
        switch v3 {
        case "foo bar":
            v64 := GenS6{F0: 4294967295, F1: -128, F2: 0}
            _ = v64
        case "a":
            _ = len(v4)
        }
    }
    for fi65 := 0; fi65 < 5; fi65++ {
        if (genFn1(42, (((v3 + (("a" + v3) + v3)) != "\n\t") || (-128 >= -128))) != -23) {
            fmt.Println(v1, fi65, v1)
            var v66 string = ""
            v67 := GenS1{F0: 0, F1: (255 >= 300), F2: 18446744073709551615}
            _ = v66
            _ = v67
        } else {
            v1 = 2
        }
        v1 += (v1 << ((0) & 63))
        _ = fi65
    }
    for fi68 := 0; fi68 < 10; fi68++ {
        var v69 int32 = 42
        v69 = v69
        v3 = v3
        _ = fi68
        _ = v69
    }
    delete(v4, -2147483648)
    fmt.Println(v1)
    var v70 uint16 = 844
    v71 := v2[(((v3 + ("\n\t" + v3)) + v3) + v3)]
    var v72 float64 = 0.0
    var v73 uint8 = v1
    for fi74 := 0; fi74 < 6; fi74++ {
        v75 := []int{(-(v71)), (v71 | fi74)}
        _ = fi74
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
