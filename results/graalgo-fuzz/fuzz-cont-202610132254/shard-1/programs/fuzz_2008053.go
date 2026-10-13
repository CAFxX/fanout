package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64, p1 int32, p2 int) bool {
    var v6 int16 = 42
    v6 = v6
    var v7 int16 = 1000
    var v8 uint64 = p0
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ((0.0 + 1.5) < ((((1e10 - 0.0) / 0.1) / float64((0.0 + 3.14159))) + 3.14159))
}
func genFn1(p0 uint8, p1 uint16, p2 string) int32 {
    p1 -= (p1 >> ((42) & 63))
    if true {
        switch ("\n\t" + "a") {
        case "\n\t":
            var v9 uint32 = 387
            _ = v9
        case "":
            fmt.Println(p2, p0, p2)
        case "hello":
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(p1, p1, p0)
                p1 += p1
                _ = fi10
            }
        default:
            v11 := GenS1{F0: -1, F1: -1}
            _ = v11
        }
    } else {
        if true {
            fmt.Println(p2, p2)
            v12 := []int{173, 32768, -9223372036854775808, -2147483648}
            v13 := GenS2{F0: 42, F1: 4294967295}
            _ = v12
            _ = v13
        } else {
            fmt.Println(p0)
        }
    }
    p2 = ""
    switch "a" {
    case "foo bar":
        v14 := GenS2{F0: 90, F1: 2}
        _ = v14
    case "\n\t":
        var v15 uint16 = 2
        _ = v15
    case "a":
        switch 42 {
        case 63:
            v16 := GenS2{F0: 9223372036854775807, F1: 2}
            _ = v16
        case 32767:
            for fi17 := 0; fi17 < 3; fi17++ {
                switch p2 {
                case "\n\t":
                    v18 := []int{(fi17 % (((fi17 & (fi17 | (fi17 >> ((726) & 63))))) | 1)), 4294967295, fi17, fi17}
                    _ = v18
                case "hello":
                    v19 := GenS1{F0: 42, F1: -32768}
                    _ = v19
                }
                if false {
                    var v20 int16 = 1000
                    var v21 string = ""
                    _ = v20
                    _ = v21
                } else {
                    var v22 bool = ((!false) && (false && ((0 == uint64(p0)) && true)))
                    _ = v22
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 uint64 = 504
                    var v25 int = (^(fi23))
                    var v26 uint = 493
                    p1 += 1
                    var v27 bool = ((p1 - (p1 >> ((v26) & 63))) <= (p1 | 909))
                    var v28 float64 = 3.14159
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi17
            }
        default:
            fmt.Println(p1, p1, p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 128
    var v3 float64 = ((1e10 * 0.0) - (-2.25 * 0.1))
    v4 := []uint32{285, 481, 4294967295, 65535}
    var v5 int8 = -1
    switch "foo bar" {
    case "a":
        if true {
            _ = len(v4)
        } else {
            var v29 string = (("hello" + "\n\t") + "hello")
            v2 = 155
            v5 = v5
            _ = v29
        }
    case "foo bar":
        v2 += 2
    case "":
        if (false || true) {
            if (!(583 < 41)) {
                v5 = v5
                for fi30 := 0; fi30 < 1; fi30++ {
                    if (true && ((v5 * (v5 & (^(v5)))) > v5)) {
                        var v31 bool = (0 < 489)
                        _ = v31
                    } else {
                        fmt.Println(v2, v3)
                        fmt.Println(v5)
                        _ = len(v4)
                    }
                    fmt.Println(v3, v3, v2)
                    _ = fi30
                }
                v2 = v2
            }
            var v32 int8 = v5
            _ = v32
        }
    default:
        v33 := v1[486]
        _ = v33
    }
    v5 -= (v5 + v5)
    var v34 uint8 = v2
    v35 := GenS2{F0: 0, F1: 2}
    v2 -= (v2 * (v2 | v2))
    var v36 bool = true
    v37 := GenS1{F0: -32768, F1: 42}
    var v38 int32 = 0
    for fi39 := 0; fi39 < 18; fi39++ {
        v5 -= (v5 ^ -1)
        if (!(413 > 2)) {
            v40 := v1[2147483648]
            v41 := GenS3{F0: 1, F1: 1}
            v42 := GenS4{F0: (!(!v36)), F1: 1}
            _ = v40
            _ = v41
            _ = v42
        } else {
            var v43 float32 = -2.25
            _ = v43
        }
        var v44 string = "foo bar"
        _ = fi39
        _ = v44
    }
    v3 += 1.5
    v5 = v5
    v36 = (v36 || v36)
    var v45 int64 = -1
    if false {
        v36 = (!v36)
        v46 := GenS5{F0: (-(v45))}
        _ = v46
    } else {
        var v47 int16 = -764
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
}
