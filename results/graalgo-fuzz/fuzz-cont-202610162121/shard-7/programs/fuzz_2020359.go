package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 int8, p1 int64, p2 int64) int64 {
    switch "\n\t" {
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{fi7, (^(fi8)), -779, (fi8 / ((0) | 1))}
                v10 := GenS1{F0: (fi7 % ((fi8) | 1)), F1: fi7}
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := []int{fi7, fi7, (fi7 & fi7), (fi7 & 153)}
            _ = fi7
            _ = v11
        }
    case "\n\t":
        if (("hello" + "hello") >= "\n\t") {
            p1 -= p1
        }
    }
    var v12 uint8 = 228
    var v13 int = -603
    if (!false) {
        v12 = (v12 - (v12 & 1))
    }
    switch "\n\t" {
    case "\n\t":
        var v14 bool = (195 != (v12 >> ((184) & 63)))
        _ = v14
    case "foo bar":
        p1 += p1
    case "":
        var v15 bool = false
        _ = v15
    default:
        p1 = (p2 << ((255) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p2
}

func main() {
    var v1 uint32 = 495
    var v2 uint64 = 0
    var v3 int8 = 127
    var v4 bool = (!false)
    var v5 float64 = 1e10
    var v6 string = (("hello" + (("hello" + ("" + "foo bar")) + "hello")) + "foo bar")
    fmt.Println(v4, v6)
    var v16 bool = (v4 && v4)
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int16 = -16
        if v4 {
            switch -738 {
            case -708:
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 int8 = 0
                    v21 := []int{-1}
                    if (!v16) {
                        _ = len(v21)
                        v21 = append(v21, fi17)
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            default:
                v2 = 42
            }
            if false {
                v22 := GenS1{F0: (fi17 << ((4294967295) & 63)), F1: 32767}
                _ = v22
            }
        } else {
            var v23 uint64 = v2
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: (fi17 ^ (fi17 & (fi24 - fi24))), F1: 1}
            _ = fi24
            _ = v25
        }
        _ = fi17
        _ = v18
    }
    var v26 int32 = -2147483648
    switch ("\n\t" + "hello") {
    case "foo bar":
        switch v6 {
        case "foo bar":
            v27 := GenS1{F0: -702, F1: -132}
            _ = v27
        }
    case "a":
        var v28 uint16 = 32767
        _ = v28
    case "":
        var v29 uint32 = v1
        _ = v29
    }
    v30 := GenS2{F0: 100}
    for fi31 := 0; fi31 < 12; fi31++ {
        for fi32 := 0; fi32 < 12; fi32++ {
            v6 = v6
            _ = fi32
        }
        _ = fi31
    }
    v33 := []int{32768}
    _ = len(v33)
    v34 := GenS3{F0: v2, F1: 32767, F2: 0}
    if v16 {
        for fi35 := 0; fi35 < 17; fi35++ {
            for fi36 := 0; fi36 < 8; fi36++ {
                if (v6 == "foo bar") {
                    if len(v33) > 0 {
                        v33[0] = fi35
                    }
                    switch fi35 {
                    case 4294967295:
                        v6 = "a"
                    case 127:
                        v37 := GenS3{F0: (v2 + v2), F1: (fi36 % ((fi35) | 1)), F2: 128}
                        _ = v37
                    }
                }
                _ = fi36
            }
            _ = fi35
        }
        v38 := GenS4{F0: 0, F1: 2}
        v3 = v3
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
