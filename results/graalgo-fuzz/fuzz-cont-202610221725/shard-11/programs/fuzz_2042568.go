package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 int32 }
func genFn0(p0 bool, p1 bool) int {
    var v5 int = -213
    if ((42 > 65535) && true) {
        p0 = p0
    }
    v5 += (-(v5))
    for fi6 := 0; fi6 < 18; fi6++ {
        if (v5 != (fi6 + (fi6 * v5))) {
            var v7 int8 = -1
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: 2, F1: 257, F2: 2}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return v5
}
func genFn1(p0 uint64, p1 uint64, p2 uint) uint32 {
    p1 = (p0 | p1)
    switch -2147483648 {
    case -1:
        v9 := GenS2{F0: "a"}
        _ = v9
    case 42:
        p2 = p2
    case 873:
        fmt.Println(p2, p2, p0)
    }
    if (!(-1 == -103)) {
        fmt.Println(p0, p0, p0)
    } else {
        var v10 bool = (-2.25 < (3.14159 + ((-2.25 - (1e10 - 0.0)) / -2.25)))
        v11 := GenS2{F0: "hello"}
        _ = v10
        _ = v11
    }
    if true {
        v12 := GenS3{F0: 32767, F1: ((-302 < 32767) && false)}
        p1 = p0
        switch -564 {
        case 32768:
            var v13 uint32 = 858
            _ = v13
        case 9223372036854775807:
            v14 := []int{genFn0((-871 <= 0), false), 32767}
            _ = v14
        }
        _ = v12
    } else {
        switch genFn0((0.1 <= 1.5), true) {
        case 0:
            if true {
                var v15 string = (("hello" + "foo bar") + ("\n\t" + "foo bar"))
                v16 := GenS3{F0: -507, F1: (!((!((0.0 / 1e10) == 0.0)) && true))}
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint8 = 1
                    _ = fi17
                    _ = v18
                }
                _ = v15
                _ = v16
            }
        default:
            var v19 float64 = float64((((3.14159 - 0.1) + 1.5) / ((0.1 / (3.14159 - (float32(3.14159) + (3.14159 + float32(1.5))))) * (1.5 - 0.1))))
            _ = v19
        }
        if (((p0 * p1) > p0) && (false || (100 != 125))) {
            v20 := GenS1{F0: 161, F1: (p0 ^ p1), F2: 2}
            if (-216 <= 0) {
                for fi21 := 0; fi21 < 17; fi21++ {
                    fmt.Println(p1, p0, p2)
                    var v22 int32 = -1
                    if false {
                        if (fi21 <= (fi21 % (((fi21 / ((-9223372036854775808) | 1))) | 1))) {
                            v22 -= (v22 / ((1) | 1))
                        }
                    }
                    _ = fi21
                    _ = v22
                }
            }
            _ = v20
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 int = fi23
                v25 := GenS3{F0: -978, F1: false}
                _ = fi23
                _ = v24
                _ = v25
            }
            p2 -= (p2 ^ p2)
        }
    }
    var v26 uint32 = 599
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return (^(v26))
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 int16 = 32767
    var v27 bool = false
    v2 = (v2 ^ v2)
    if v27 {
        v27 = (v2 != v2)
        if v27 {
            v28 := GenS4{F0: 42}
            _ = v28
        } else {
            if v27 {
                var v29 int8 = (v2 ^ 1)
                v30 := GenS5{F0: 2147483647, F1: ("a" + "hello")}
                var v31 float64 = (0.1 * (1.5 / 1e10))
                _ = v29
                _ = v30
                _ = v31
            } else {
                var v32 int64 = 1
                _ = v32
            }
        }
    } else {
        if false {
            var v33 int64 = 985
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        switch ("" + "hello") {
        case "":
            v4 = -206
        case "hello":
            var v35 uint64 = v1
            _ = v35
        case "a":
            fmt.Println(v27, v2)
        }
        var v36 int16 = v4
        if (1e10 <= (((1e10 + (0.1 - float64(-2.25))) / 1.5) - -2.25)) {
            v36 = (v4 | v36)
        }
        _ = fi34
        _ = v36
    }
    fmt.Println(v1, v2)
    if v27 {
        for fi37 := 0; fi37 < 4; fi37++ {
            if true {
                v38 := []int{(fi37 << ((910) & 63)), (fi37 | genFn0(((v1 % ((18446744073709551615) | 1)) == (^(v1))), true))}
                for fi39 := 0; fi39 < 5; fi39++ {
                    var v40 bool = v27
                    if len(v38) > 0 {
                        v38[0] = (fi39 | (fi37 >> ((4294967295) & 63)))
                    }
                    fmt.Println(v4)
                    delete(v3, 2147483647)
                    fmt.Println(fi39, fi39, v27)
                    _ = fi39
                    _ = v40
                }
                _ = v38
            }
            _ = fi37
        }
        v41 := GenS4{F0: genFn1((v1 / ((v1) | 1)), v1, 18446744073709551615)}
        _ = v41
    } else {
        v2 -= (-(v2))
        fmt.Println(v4)
        if v27 {
            fmt.Println(v1, v2)
        } else {
            if v27 {
                switch 127 {
                case 226:
                    var v42 uint16 = 1
                    _ = v42
                case 4294967295:
                    v2 = (v2 & v2)
                case 1:
                    if true {
                        var v43 bool = (v27 && (true || v27))
                        _ = v43
                    } else {
                        v2 -= 1
                        _ = len(v3)
                        v44 := GenS2{F0: "a"}
                        _ = v44
                    }
                default:
                    for fi45 := 0; fi45 < 17; fi45++ {
                        v1 = v1
                        var v46 int32 = 1
                        v27 = true
                        _ = fi45
                        _ = v46
                    }
                }
            } else {
                var v47 uint16 = 590
                v48 := GenS6{F0: 32768, F1: 9223372036854775807}
                _ = v47
                _ = v48
            }
            v49 := GenS7{F0: v1, F1: 42, F2: 1}
            if v27 {
                var v50 bool = ((1.5 == 1e10) || v27)
                switch "\n\t" {
                case "a":
                    v4 = (v4 % ((-1) | 1))
                default:
                    v51 := GenS4{F0: 435}
                    _ = v51
                }
                for fi52 := 0; fi52 < 2; fi52++ {
                    for fi53 := 0; fi53 < 19; fi53++ {
                        v50 = v27
                        v54 := v3[2147483647]
                        _ = fi53
                        _ = v54
                    }
                    if (3.14159 == float64(float32(3.14159))) {
                        v27 = (!(v50 && false))
                        var v55 int = 470
                        v2 = v2
                        _ = v55
                    } else {
                        for fi56 := 0; fi56 < 5; fi56++ {
                            var v57 uint = 869
                            _ = fi56
                            _ = v57
                        }
                        fmt.Println(v4, v4, v50)
                        for fi58 := 0; fi58 < 6; fi58++ {
                            var v59 int16 = (v4 + v4)
                            var v60 bool = (v50 && v27)
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    }
                    var v61 uint = 0
                    _ = fi52
                    _ = v61
                }
                _ = v50
            }
            _ = v49
        }
    }
    if false {
        var v62 uint16 = 42
        _ = v62
    } else {
        for fi63 := 0; fi63 < 19; fi63++ {
            if (405 != 65535) {
                for fi64 := 0; fi64 < 12; fi64++ {
                    v65 := []int{(fi63 | (fi63 / (((fi63 - 2147483647)) | 1))), 605, genFn0((!(65535 > 42)), (2 < 18446744073709551615))}
                    v66 := GenS5{F0: 42, F1: ""}
                    fmt.Println(v1)
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                var v67 uint32 = 42
                _ = v67
            } else {
                v68 := GenS8{F0: (v1 - (v1 & v1)), F1: 112}
                fmt.Println(v1)
                _ = v68
            }
            var v69 bool = v27
            fmt.Println(v4, fi63, fi63)
            _ = fi63
            _ = v69
        }
    }
    v70 := GenS8{F0: v1, F1: 42}
    var v71 string = "foo bar"
    v72 := GenS9{F0: 42}
    v73 := GenS7{F0: (v1 | 2), F1: 9223372036854775807, F2: (v1 ^ (v1 * v1))}
    var v74 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
