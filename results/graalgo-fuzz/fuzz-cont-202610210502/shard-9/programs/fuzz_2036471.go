package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS6 struct { F0 string; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 bool) uint8 {
    v7 := GenS1{F0: "a"}
    fmt.Println(p1, p1)
    v8 := GenS1{F0: ""}
    for fi9 := 0; fi9 < 19; fi9++ {
        switch ("a" + (("\n\t" + "\n\t") + (("" + "a") + ("a" + ("a" + "foo bar"))))) {
        case "a":
            var v10 uint32 = 2147483648
            _ = v10
        case "hello":
            if true {
                var v11 string = ""
                v12 := GenS2{F0: 67, F1: 87, F2: (p0 - (p0 >> ((65535) & 63)))}
                _ = v11
                _ = v12
            } else {
                if true {
                    v13 := GenS3{F0: (("\n\t" + (("\n\t" + "hello") + "hello")) + "a"), F1: 128}
                    v14 := []int{-9223372036854775808, (^(fi9)), (fi9 >> ((65535) & 63)), (-(fi9))}
                    _ = v13
                    _ = v14
                }
                if p1 {
                    v15 := []int{127, (fi9 << ((1) & 63)), 1}
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 int = fi16
                        v18 := []int{(fi16 * -1), (fi16 & 4294967295), (fi9 | (fi16 * fi9))}
                        switch (("foo bar" + "hello") + (("\n\t" + "") + "foo bar")) {
                        case "hello":
                            v19 := GenS1{F0: ("\n\t" + (("hello" + ((("hello" + "\n\t") + "") + "foo bar")) + "foo bar"))}
                            _ = v19
                        case "\n\t":
                            var v20 int32 = 325
                            _ = v20
                        default:
                            v21 := GenS2{F0: 100, F1: 18446744073709551615, F2: p0}
                            _ = v21
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    p1 = false
                    var v22 int = 937
                    _ = v22
                }
                if (0 <= 413) {
                    if false {
                        if (65535 > 42) {
                            var v23 uint32 = 42
                            switch "foo bar" {
                            case "foo bar":
                                var v24 int16 = 827
                                _ = v24
                            case "\n\t":
                                fmt.Println(fi9, fi9, v23)
                            case "a":
                                v25 := GenS3{F0: "foo bar", F1: (fi9 | 128)}
                                _ = v25
                            }
                            p0 = 2
                            _ = v23
                        }
                    } else {
                        v26 := []int{fi9, fi9, fi9, fi9}
                        p1 = false
                        _ = v26
                    }
                    v27 := GenS3{F0: "", F1: (-(fi9))}
                    _ = v27
                }
            }
        case "foo bar":
            var v28 float32 = 0.1
            _ = v28
        default:
            fmt.Println(fi9)
        }
        switch "hello" {
        case "":
            v29 := []int{(fi9 * (fi9 + fi9))}
            _ = v29
        case "\n\t":
            if false {
                for fi30 := 0; fi30 < 18; fi30++ {
                    p1 = (!p1)
                    var v31 int16 = 32767
                    fmt.Println(p0, v31, p1)
                    p0 -= (^(p0))
                    _ = fi30
                    _ = v31
                }
                var v32 uint = 0
                v32 = v32
                _ = v32
            } else {
                if p1 {
                    v33 := GenS4{F0: p1, F1: 32767}
                    v34 := GenS4{F0: (p1 && (false && p1)), F1: 255}
                    _ = v33
                    _ = v34
                }
                var v35 uint = 4294967295
                switch "a" {
                case "":
                    if true {
                        fmt.Println(v35)
                    }
                case "foo bar":
                    if true {
                        v36 := GenS3{F0: ((("foo bar" + "hello") + "a") + "foo bar"), F1: fi9}
                        v37 := []int{(^(fi9))}
                        _ = v36
                        _ = v37
                    }
                case "\n\t":
                    var v38 int = fi9
                    _ = v38
                }
                _ = v35
            }
        case "a":
            var v39 uint8 = 56
            _ = v39
        }
        _ = fi9
    }
    if p1 {
        if (1 == 635) {
            var v40 string = "hello"
            switch 1 {
            case 32768:
                v41 := GenS4{F0: (754 < 0), F1: 249}
                _ = v41
            }
            v40 = "a"
            _ = v40
        }
        var v42 int = -537
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 uint32 = 1
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 119
    v6 := make(map[uint32]string)
    v43 := GenS3{F0: "foo bar", F1: -914}
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 int = 1
        v46 := GenS1{F0: ""}
        v47 := GenS5{F0: v3, F1: genFn0((v5 / (((v5 >> ((uint(v5)) & 63))) | 1)), (1e10 < float32(1e10))), F2: uint8(v3)}
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        if false {
            v3 = v3
        }
        v5 -= genFn0((v5 * (^(v5))), (false && ((-1 < -128) && true)))
        _ = fi48
    }
    v49 := GenS1{F0: ""}
    v4 = 4294967295
    v50 := GenS6{F0: "\n\t", F1: 1, F2: 1}
    v5 = uint8(v4)
    v4 += 18446744073709551615
    if (42 == 2) {
        var v51 float32 = (1.5 * -2.25)
        v51 -= v51
        _ = v51
    } else {
        if (!false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v52 string = ""
        switch 128 {
        case -788:
            switch -1 {
            case 9223372036854775807:
                var v53 float32 = (1e10 - 0.1)
                _ = v53
            default:
                for fi54 := 0; fi54 < 8; fi54++ {
                    var v55 uint8 = (v5 | (v5 << ((uint(v5)) & 63)))
                    _ = fi54
                    _ = v55
                }
            }
        }
        _ = v52
    }
    _ = len(v1)
    fmt.Println(v5, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
