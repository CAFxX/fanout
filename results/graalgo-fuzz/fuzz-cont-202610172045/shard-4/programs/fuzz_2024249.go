package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int64, p1 uint8) int64 {
    if true {
        for fi4 := 0; fi4 < 3; fi4++ {
            fmt.Println(p1, fi4, fi4)
            v5 := GenS1{F0: "\n\t"}
            v6 := []int{(fi4 + fi4), (fi4 & 329), (fi4 + (fi4 & fi4))}
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(fi7 % ((fi7) | 1))}
        if false {
            for fi9 := 0; fi9 < 7; fi9++ {
                switch fi7 {
                case -1:
                    var v10 int64 = (-(p0))
                    _ = v10
                case 32768:
                    if (-2147483648 > 1) {
                        var v11 float32 = (3.14159 + -2.25)
                        _ = v11
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            v12 := GenS1{F0: "hello"}
                            _ = v12
                        case "a":
                            fmt.Println(fi9, fi9)
                        case "hello":
                            var v13 float32 = (1e10 + 3.14159)
                            _ = v13
                        }
                    }
                default:
                    switch fi7 {
                    case -790:
                        p0 += p0
                    case 590:
                        _ = len(v8)
                    }
                }
                var v14 float64 = (0.1 * -2.25)
                var v15 uint8 = p1
                _ = fi9
                _ = v14
                _ = v15
            }
        } else {
            var v16 bool = (!false)
            var v17 int64 = -9223372036854775808
            v18 := GenS1{F0: ("\n\t" + "a")}
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 uint64 = 4294967295
        _ = fi7
        _ = v19
        _ = v8
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS1{F0: "\n\t"}
        _ = fi20
        _ = v21
    }
    v22 := []int{}
    if (!true) {
        switch "" {
        case "\n\t":
            p0 += -946
        default:
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 int8 = -1
                var v25 uint8 = (p1 | 215)
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        var v26 int = 1
        v26 += v26
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v22
    return (p0 / ((p0) | 1))
}
func genFn1(p0 int16, p1 uint8, p2 uint) string {
    v27 := GenS1{F0: ""}
    for fi28 := 0; fi28 < 7; fi28++ {
        p2 = (p2 - (^(p2)))
        fmt.Println(p0)
        _ = fi28
    }
    var v29 uint16 = 0
    p1 -= 2
    v29 = (v29 - (^(v29)))
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v29
    return ""
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{}
    v3 := []int{}
    var v30 uint64 = 337
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 float64 = 3.14159
        v30 -= 0
        _ = fi31
        _ = v32
    }
    var v33 int8 = 1
    var v34 uint = 255
    if false {
        v35 := GenS2{F0: (v1 * (v1 >> ((v34) & 63))), F1: 2147483647, F2: true}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v35
    } else {
        v36 := GenS2{F0: genFn0((v1 & genFn0(v1, 0)), 2), F1: 332, F2: (444 > -115)}
        v37 := GenS1{F0: ("hello" + ("\n\t" + "a"))}
        v33 = v33
        _ = v36
        _ = v37
    }
    v34 -= (v34 ^ v34)
    v33 = v33
    v34 = v34
    v38 := GenS2{F0: 0, F1: 656, F2: true}
    fmt.Println(v1, v34)
    if false {
        var v39 int = 1
        var v40 int8 = (v33 % (((v33 - v33)) | 1))
        _ = v39
        _ = v40
    } else {
        v41 := GenS3{F0: 42, F1: 454}
        v30 += (v30 | v30)
        _ = v41
    }
    var v42 bool = false
    if ((("foo bar" + "") <= ("hello" + "hello")) || ("" <= ("" + "a"))) {
        fmt.Println(v42, v34, v1)
        v43 := GenS1{F0: ((("foo bar" + (("foo bar" + "foo bar") + "hello")) + "") + "")}
        for fi44 := 0; fi44 < 1; fi44++ {
            var v45 int64 = (v1 / ((v1) | 1))
            _ = fi44
            _ = v45
        }
        _ = v43
    } else {
        for fi46 := 0; fi46 < 11; fi46++ {
            v30 = (v30 & (^(v30)))
            var v47 int16 = -1
            _ = fi46
            _ = v47
        }
        v48 := GenS3{F0: 2, F1: uint16(v34)}
        _ = v48
    }
    v33 += (v33 | v33)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v38
    _ = v42
}
