package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int8 }
func genFn0(p0 int) int16 {
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 = v6
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                p0 = 127
                _ = fi8
            }
            v6 -= (v6 + (^(v6)))
        } else {
            var v9 string = "foo bar"
            v10 := []int{p0, p0, 990, p0}
            var v11 int16 = -162
            _ = v10
            _ = v11
            _ = v9
        }
        switch p0 {
        case -868:
            var v12 bool = true
            _ = v12
        }
        _ = fi7
    }
    switch p0 {
    case -29:
        v13 := GenS1{F0: 18446744073709551615, F1: false}
        _ = v13
    case 973:
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 string = "foo bar"
            _ = fi14
            _ = v15
        }
    case 1:
        var v16 string = ("foo bar" + "hello")
        _ = v16
    default:
        var v17 uint8 = 218
        _ = v17
    }
    var v18 uint = 2
    _ = p0
    _ = v18
    _ = v6
    return 32767
}
func genFn1() uint16 {
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            if true {
                fmt.Println(fi19)
                v21 := GenS1{F0: 924, F1: true}
                var v22 int32 = -1
                _ = v21
                _ = v22
            } else {
                var v23 int64 = 99
                if false {
                    var v24 string = ("a" + ("\n\t" + ""))
                    v23 = v23
                    v23 -= (v23 & (v23 * (v23 >> ((1) & 63))))
                    _ = v24
                }
                if false {
                    v25 := GenS1{F0: 2, F1: false}
                    v26 := []int{(-(fi20)), -905, fi19}
                    v23 -= (v23 << ((769) & 63))
                    if false {
                        var v27 int16 = 1000
                        _ = v27
                    }
                    var v28 uint = 11
                    _ = v25
                    _ = v26
                    _ = v28
                }
                _ = v23
            }
            var v29 int = 42
            _ = fi20
            _ = v29
        }
        v30 := []int{1, 2147483647, fi19}
        _ = fi19
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if true {
            if ((fi31 + 128) <= 390) {
                if true {
                    var v32 float32 = (1.5 + 0.1)
                    v32 = v32
                    v33 := GenS2{F0: 0, F1: 783, F2: 66}
                    _ = v32
                    _ = v33
                }
                v34 := []int{-290}
                _ = v34
            }
        }
        _ = fi31
    }
    return 509
}

func main() {
    var v1 uint16 = 620
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    var v5 int8 = 0
    if (!false) {
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 uint16 = (v1 | (v1 >> ((228) & 63)))
            _ = fi35
            _ = v36
        }
    }
    v37 := GenS3{F0: -9223372036854775808, F1: 4294967295}
    if true {
        for fi38 := 0; fi38 < 2; fi38++ {
            v5 += v5
            _ = fi38
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            fmt.Println(v5, v1, v5)
            _ = fi39
        }
    }
    v1 = (v1 + (v1 - 0))
    if false {
        delete(v4, 2147483647)
        v40 := GenS2{F0: v3, F1: 2, F2: -51}
        _ = v40
    }
    if true {
        var v41 int32 = -106
        _ = v41
    } else {
        if true {
            var v42 int32 = -1
            _ = v42
        } else {
            v43 := GenS4{F0: "", F1: 381}
            _ = v43
        }
        var v44 int64 = 0
        var v45 float64 = 0.0
        _ = v44
        _ = v45
    }
    var v46 int8 = -128
    fmt.Println(v3, v3, v1)
    v47 := GenS5{F0: 2, F1: 127, F2: (float32((1.5 / ((0.0 + (3.14159 / 3.14159)) - 3.14159))) != 3.14159)}
    v48 := GenS6{F0: 2147483648}
    var v49 uint = 667
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 int = fi50
        _ = fi50
        _ = v51
    }
    delete(v2, 424)
    var v52 float32 = 3.14159
    v53 := GenS7{F0: true, F1: (v5 - 42)}
    v54 := []int{-1, -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
