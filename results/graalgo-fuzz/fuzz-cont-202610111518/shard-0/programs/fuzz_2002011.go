package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int8; F1 int16; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 bool }
func genFn0(p0 uint) int8 {
    p0 = (p0 >> ((1) & 63))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: "hello", F1: -1, F2: 761}
        var v9 uint8 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.0
    v11 := GenS2{F0: 127}
    p0 -= (^(p0))
    _ = p0
    _ = v10
    _ = v11
    return 127
}
func genFn1() int {
    if ("hello" >= "a") {
        var v12 string = ""
        _ = v12
    } else {
        var v13 uint32 = 1
        v14 := []int{}
        _ = v13
        _ = v14
    }
    if ("hello" != "") {
        var v15 uint64 = 463
        for fi16 := 0; fi16 < 11; fi16++ {
            if false {
                v15 = 992
                if false {
                    fmt.Println(v15, v15)
                    v15 -= (v15 & (^(v15)))
                    if true {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v15 = (v15 * 18446744073709551615)
                            v18 := []int{-1}
                            var v19 uint32 = 844
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v15 = v15
                        if (2 < 255) {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                var v21 int = (fi16 | fi20)
                                v22 := GenS1{F0: "\n\t", F1: 0, F2: 1}
                                var v23 int16 = -32768
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            if true {
                                v15 += 42
                                fmt.Println(v15)
                            } else {
                                v15 = (v15 * v15)
                                if (true || ((-(fi16)) < (fi16 >> ((1) & 63)))) {
                                    v15 = (v15 >> ((42) & 63))
                                    v24 := GenS3{F0: 42, F1: 4294967295}
                                    var v25 string = (("\n\t" + "\n\t") + ("hello" + ""))
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint16 = 65535
                                    var v27 int8 = 0
                                    _ = v26
                                    _ = v27
                                }
                            }
                            switch 462 {
                            case 70:
                                v28 := GenS2{F0: 127}
                                _ = v28
                            case 32768:
                                v15 = (v15 << ((51) & 63))
                            case -871:
                                v15 += 0
                            default:
                                v29 := GenS2{F0: 1}
                                _ = v29
                            }
                        } else {
                            v30 := GenS1{F0: "a", F1: 42, F2: 1}
                            _ = v30
                        }
                    } else {
                        v31 := []int{}
                        _ = v31
                    }
                }
                v15 -= v15
            } else {
                v32 := []int{(fi16 ^ fi16)}
                if false {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v34 := GenS3{F0: -1, F1: 929}
                        _ = fi33
                        _ = v34
                    }
                    var v35 uint16 = 2
                    _ = v35
                } else {
                    var v36 uint64 = (^(v15))
                    _ = v36
                }
                fmt.Println(fi16, fi16)
                _ = v32
            }
            switch 9223372036854775807 {
            case 958:
                v37 := GenS1{F0: "a", F1: 0, F2: 328}
                _ = v37
            case 9223372036854775807:
                v38 := GenS4{F0: 127, F1: -32768, F2: false}
                _ = v38
            }
            _ = fi16
        }
        fmt.Println(v15)
        _ = v15
    } else {
        v39 := GenS1{F0: ("foo bar" + ((("hello" + ("hello" + "a")) + "\n\t") + "hello")), F1: -1, F2: -2147483648}
        var v40 uint8 = 138
        _ = v39
        _ = v40
    }
    if (!((1.5 + 0.0) != 1.5)) {
        var v41 uint = 18446744073709551615
        v42 := []int{-2147483648, 9223372036854775807, 4294967295, -745}
        _ = v41
        _ = v42
    }
    return 32768
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 0
    v3 := []uint32{1, 0}
    var v4 float64 = -2.25
    var v5 int = 599
    v6 := make(map[int]bool)
    v1[-2147483648] = (float64(((0.0 - (-2.25 / -2.25)) - 0.1)) > ((1.5 * 1.5) * float64(0.0)))
    v5 = 9223372036854775807
    v43 := GenS5{F0: 0}
    v5 += v5
    v44 := GenS4{F0: 127, F1: 1, F2: (genFn0(0) != 0)}
    v45 := GenS1{F0: "", F1: -543, F2: -2147483648}
    for fi46 := 0; fi46 < 20; fi46++ {
        if true {
            fmt.Println(v5)
            if len(v3) > 0 {
                v3[0] = 889
            }
        } else {
            if (1 > 4294967295) {
                v4 += v4
            }
        }
        var v47 string = "foo bar"
        var v48 uint8 = 128
        _ = fi46
        _ = v47
        _ = v48
    }
    v2 = (^(v2))
    for fi49 := 0; fi49 < 10; fi49++ {
        switch (-(fi49)) {
        case 127:
            delete(v6, (^(v5)))
        case -19:
            for fi50 := 0; fi50 < 15; fi50++ {
                var v51 int8 = 0
                _ = fi50
                _ = v51
            }
        case 128:
            fmt.Println(v5, v5)
        default:
            for fi52 := 0; fi52 < 9; fi52++ {
                var v53 int32 = -2147483648
                _ = fi52
                _ = v53
            }
        }
        var v54 uint32 = 635
        v54 = 765
        _ = fi49
        _ = v54
    }
    v2 = v2
    var v55 uint = 1
    fmt.Println(v5)
    var v56 uint = v55
    for fi57 := 0; fi57 < 7; fi57++ {
        var v58 int32 = 1
        _ = fi57
        _ = v58
    }
    v59 := GenS6{F0: genFn1(), F1: (v55 > v55)}
    var v60 int64 = -9223372036854775808
    if false {
        var v61 int8 = -97
        _ = v61
    } else {
        fmt.Println(v55, v55, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v55
    _ = v56
    _ = v59
    _ = v6
    _ = v60
}
