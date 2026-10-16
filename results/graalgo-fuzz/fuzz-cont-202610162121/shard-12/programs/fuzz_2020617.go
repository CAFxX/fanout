package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 int, p2 uint) int16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        switch "" {
        case "foo bar":
            p2 = (p2 | p2)
        case "a":
            if true {
                v8 := GenS1{F0: 1, F1: 42, F2: 0}
                fmt.Println(p0, p1, p2)
                v9 := GenS1{F0: 9223372036854775807, F1: -126, F2: 2}
                _ = v8
                _ = v9
            }
        default:
            p1 = -647
        }
        _ = fi7
    }
    if true {
        v10 := []int{(p1 << (((p2 >> ((p2) & 63))) & 63)), (p1 + p1), -608, p1}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}
func genFn1(p0 string) int64 {
    p0 = p0
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 42, F1: 100, F2: 4294967295}
        switch (("" + "\n\t") + ((("a" + ((p0 + p0) + p0)) + "foo bar") + p0)) {
        case "foo bar":
            p0 = "a"
        case "\n\t":
            if (!false) {
                v13 := []int{(^(fi11)), fi11, -9223372036854775808, 0}
                v14 := GenS1{F0: -176, F1: -109, F2: 864}
                _ = v13
                _ = v14
            }
        case "a":
            v15 := []int{fi11, 4294967295, (fi11 % (((-(fi11))) | 1))}
            _ = v15
        default:
            p0 = (((((p0 + ((p0 + p0) + (p0 + (p0 + p0)))) + p0) + p0) + ("" + ("hello" + "\n\t"))) + "")
        }
        if false {
            p0 = p0
            fmt.Println(fi11)
        } else {
            v16 := GenS2{F0: 42, F1: 255}
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    switch 32768 {
    case -1:
        for fi17 := 0; fi17 < 20; fi17++ {
            if (4294967295 == 4294967295) {
                var v18 float32 = 0.0
                _ = v18
            }
            var v19 uint16 = 42
            _ = fi17
            _ = v19
        }
    }
    _ = p0
    return 324
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{-450, 638, 128}
    v5 := []string{"\n\t", "foo bar", "hello", "\n\t"}
    v6 := make(map[int]uint32)
    var v20 float64 = v1
    var v21 int64 = -1
    v21 = v21
    var v22 int16 = -32768
    v23 := GenS3{F0: 723, F1: false, F2: 13}
    if false {
        var v24 int8 = -1
        _ = v24
    }
    v25 := GenS4{F0: 246, F1: -263}
    v21 = v21
    for fi26 := 0; fi26 < 18; fi26++ {
        if (v21 > int64(v22)) {
            _ = len(v6)
            fmt.Println(v22, v22, v1)
            if false {
                var v27 int16 = (v22 & (^(v22)))
                for fi28 := 0; fi28 < 16; fi28++ {
                    if (((2 >= 702) || ((((("" + ("" + "a")) + "a") + "foo bar") + "foo bar") >= (("a" + "") + (((("foo bar" + "hello") + "foo bar") + "a") + "")))) || ("\n\t" != "foo bar")) {
                        var v29 float32 = -2.25
                        _ = len(v5)
                        v27 = (v22 >> ((2) & 63))
                        _ = v29
                    } else {
                        var v30 string = "\n\t"
                        var v31 int16 = 32767
                        v4 = append(v4, fi26)
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                }
                _ = v27
            }
        }
        _ = fi26
    }
    var v32 int32 = 181
    var v33 bool = false
    for fi34 := 0; fi34 < 20; fi34++ {
        var v35 uint64 = 681
        var v36 uint64 = 2
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint16 = 32767
    fmt.Println(v37, v33)
    if false {
        var v38 bool = (!(v22 == v22))
        var v39 uint = 18446744073709551615
        v40 := GenS3{F0: 42, F1: (706 <= 0), F2: 2}
        _ = v38
        _ = v39
        _ = v40
    } else {
        var v41 float32 = (1.5 + ((0.0 - 3.14159) - (1e10 + 3.14159)))
        if v33 {
            v42 := GenS4{F0: 255, F1: v32}
            _ = v42
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if v33 {
                fmt.Println(v21, v21, v21)
            }
            var v43 int64 = -1
            _ = v43
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
