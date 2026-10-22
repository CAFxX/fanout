package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 int8; F2 uint64 }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 int8; F1 string; F2 uint32 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 bool; F1 bool }
func genFn0(p0 uint16, p1 uint8) int32 {
    v6 := GenS1{F0: -105, F1: 127, F2: 0}
    p1 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v6
    return -136
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    v4 := []int{1}
    v5 := []int{127, 128}
    v7 := GenS2{F0: 128, F1: -120}
    var v8 uint16 = 255
    v5 = append(v5, 4294967295)
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint8 = 75
                if false {
                    var v12 int64 = 9223372036854775807
                    _ = v12
                }
                v11 -= v11
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        switch 128 {
        case -647:
            var v13 int64 = 0
            _ = v13
        default:
            fmt.Println(v8, v2)
        }
        v14 := GenS1{F0: 0, F1: 0, F2: 18446744073709551615}
        _ = v14
    } else {
        v15 := GenS3{F0: -128, F1: 1}
        _ = v15
    }
    switch -1 {
    case 127:
        if (0 > 1) {
            var v16 string = "hello"
            _ = v16
        } else {
            fmt.Println(v8)
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        delete(v1, fi17)
        var v18 uint16 = 32767
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS2{F0: 1, F1: 88}
        v21 := GenS4{F0: false, F1: 1000}
        v22 := GenS5{F0: v8, F1: 4294967295}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    switch -672 {
    case 128:
        var v23 float64 = 3.14159
        _ = v23
    case 32768:
        v2 += v2
    case 32767:
        if ((^(v8)) > v8) {
            for fi24 := 0; fi24 < 4; fi24++ {
                v8 = (v8 + 951)
                v25 := GenS6{F0: (^(v2)), F1: (fi24 | 42)}
                var v26 int16 = v2
                _ = fi24
                _ = v25
                _ = v26
            }
            if (18446744073709551615 < 4294967295) {
                for fi27 := 0; fi27 < 13; fi27++ {
                    if (!false) {
                        v4 = append(v4, -228)
                    } else {
                        v28 := GenS7{F0: 42}
                        v29 := GenS4{F0: false, F1: 1}
                        if true {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                v31 := GenS8{F0: 1}
                                var v32 uint8 = 255
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = v28
                        _ = v29
                    }
                    var v33 uint64 = 682
                    _ = fi27
                    _ = v33
                }
            }
            var v34 float64 = 1e10
            _ = v34
        } else {
            fmt.Println(v2, v2)
        }
    default:
        var v35 uint64 = 42
        _ = v35
    }
    fmt.Println(v8)
    var v36 int = 32768
    if false {
        var v37 bool = false
        var v38 int = v36
        _ = v37
        _ = v38
    } else {
        fmt.Println(v36, v2)
        var v39 int16 = 821
        _ = v39
    }
    v2 = v2
    _ = len(v3)
    v40 := GenS2{F0: 46, F1: -6}
    if (false || true) {
        fmt.Println(v36, v2)
        var v41 uint8 = 0
        switch (-(v36)) {
        case 2147483647:
            if (uint(v8) <= 4294967295) {
                switch ("a" + "\n\t") {
                case "hello":
                    if (42 <= v41) {
                        v42 := GenS9{F0: 1, F1: 100, F2: 102}
                        v43 := GenS10{F0: 4294967295, F1: v41}
                        _ = v42
                        _ = v43
                    }
                case "":
                    if false {
                        v2 += -1
                        fmt.Println(v41)
                        v44 := GenS11{F0: 127, F1: ("a" + "a"), F2: 989}
                        _ = v44
                    }
                case "foo bar":
                    if ((^(v41)) > (v41 % ((0) | 1))) {
                        fmt.Println(v8)
                    } else {
                        fmt.Println(v8, v8)
                        v45 := GenS6{F0: (^(v2)), F1: 32767}
                        for fi46 := 0; fi46 < 19; fi46++ {
                            _ = len(v1)
                            _ = fi46
                        }
                        _ = v45
                    }
                }
            }
        case 0:
            fmt.Println(v36, v8, v36)
        default:
            _ = len(v3)
        }
        _ = v41
    }
    if true {
        fmt.Println(v36, v8, v36)
        var v47 uint16 = (v8 % ((v8) | 1))
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 int8 = -1
            if true {
                for fi50 := 0; fi50 < 10; fi50++ {
                    v8 -= 655
                    v51 := GenS12{F0: ("\n\t" + ("" + "foo bar")), F1: 1}
                    _ = fi50
                    _ = v51
                }
                var v52 int = (fi48 | 6)
                _ = v52
            }
            v53 := GenS13{F0: (880 >= uint64(v47)), F1: (false && true)}
            _ = fi48
            _ = v49
            _ = v53
        }
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v7
    _ = v8
}
