package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 int; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint32) uint32 {
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := []int{fi6, fi7, fi7, 948}
                fmt.Println(p0, fi6)
                fmt.Println(p0)
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        if (true && false) {
            var v9 int64 = 9223372036854775807
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS1{F0: 889}
                var v12 int32 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 9223372036854775807, F1: 255, F2: ("foo bar" + (("\n\t" + "hello") + "hello"))}
            _ = v13
            _ = v9
        } else {
            var v14 string = "foo bar"
            _ = v14
        }
        if true {
            switch "hello" {
            case "a":
                var v15 uint = 715
                _ = v15
            case "foo bar":
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 string = ("foo bar" + "a")
                    switch ("a" + v17) {
                    case "":
                        var v18 int = fi16
                        _ = v18
                    case "foo bar":
                        p0 = p0
                    case "hello":
                        p0 -= (p0 * p0)
                    }
                    _ = fi16
                    _ = v17
                }
            case "":
                v19 := []int{0, -2}
                _ = v19
            default:
                fmt.Println(p0)
            }
        } else {
            var v20 int8 = 127
            p0 -= (^(p0))
            _ = v20
        }
    } else {
        v21 := GenS3{F0: (-1 < -909), F1: 32768, F2: 1}
        v22 := GenS2{F0: 265, F1: 65535, F2: "hello"}
        v23 := []int{42}
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int32 = 1
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(fi25, fi25, p0)
        _ = fi25
    }
    var v26 string = "\n\t"
    v27 := []int{9223372036854775807}
    _ = p0
    _ = v24
    _ = v26
    _ = v27
    return p0
}
func genFn1(p0 int8, p1 int8) bool {
    if (-1 > -120) {
        v28 := []int{32768, -2147483648, -1, -181}
        p1 -= (p0 << ((2) & 63))
        _ = v28
    } else {
        var v29 int = 1
        if (true && false) {
            v30 := []int{(v29 >> ((65535) & 63)), v29}
            for fi31 := 0; fi31 < 18; fi31++ {
                v32 := GenS4{F0: 100}
                v33 := GenS1{F0: 4294967295}
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        } else {
            v34 := []int{v29}
            fmt.Println(p1)
            for fi35 := 0; fi35 < 4; fi35++ {
                v34 = append(v34, 32768)
                _ = fi35
            }
            _ = v34
        }
        if false {
            v36 := GenS5{F0: 0, F1: 1}
            p1 = 34
            _ = v36
        } else {
            v37 := GenS3{F0: (2 > 2), F1: (v29 >> ((0) & 63)), F2: 9223372036854775807}
            v29 = (v29 * v29)
            _ = v37
        }
        _ = v29
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        if (true && true) {
            if true {
                for fi39 := 0; fi39 < 6; fi39++ {
                    p0 = p0
                    fmt.Println(p0, p1, p0)
                    p0 = p1
                    _ = fi39
                }
                var v40 uint16 = 32767
                _ = v40
            } else {
                v41 := []int{fi38}
                var v42 int64 = 1
                _ = v41
                _ = v42
            }
        }
        var v43 uint32 = 0
        v43 -= v43
        _ = fi38
        _ = v43
    }
    _ = p0
    _ = p1
    return (true && (!((!(784 != -59)) || true)))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    var v3 int8 = -128
    v4 := make(map[uint32]int)
    var v5 int8 = v3
    for fi44 := 0; fi44 < 13; fi44++ {
        v2 = (v2 + (v2 + ("" + "")))
        _ = fi44
    }
    fmt.Println(v5, v5, v5)
    var v45 int32 = -319
    v46 := GenS3{F0: ((v2 + (v2 + ("" + ("a" + ((v2 + v2) + ("" + ((v2 + "hello") + ("a" + v2)))))))) >= (("a" + "foo bar") + "hello")), F1: 431, F2: 0}
    switch v2 {
    case "":
        var v47 string = ("" + ("foo bar" + ("\n\t" + v2)))
        _ = v47
    }
    v48 := GenS6{F0: 255}
    v49 := []int{-885, 0, 292}
    for fi50 := 0; fi50 < 5; fi50++ {
        v45 += (-(v45))
        _ = fi50
    }
    fmt.Println(v5, v2)
    var v51 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v51
}
