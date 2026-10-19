package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 uint32, p1 int32) int8 {
    p1 += p1
    v7 := []int{32767, 2147483647}
    fmt.Println(p1)
    _ = len(v7)
    var v8 int32 = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 uint, p1 bool) uint {
    var v9 int64 = 42
    switch "" {
    case "":
        var v10 uint32 = 410
        _ = v10
    case "\n\t":
        var v11 int16 = -1
        _ = v11
    case "foo bar":
        v12 := GenS1{F0: 489}
        _ = v12
    default:
        if p1 {
            if ((v9 | (v9 | v9)) > v9) {
                fmt.Println(p0, p0, p1)
                var v13 int16 = -293
                fmt.Println(v13, v9, p0)
                _ = v13
            } else {
                v14 := GenS2{F0: 128, F1: 42}
                if ((!(!("foo bar" <= "\n\t"))) || (255 != 128)) {
                    v15 := []int{239}
                    var v16 uint8 = 1
                    _ = v15
                    _ = v16
                }
                v17 := []int{9223372036854775807}
                _ = v14
                _ = v17
            }
        }
    }
    p0 = (p0 / ((p0) | 1))
    v18 := []int{127, 32768, 42}
    _ = p0
    _ = p1
    _ = v18
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 185
    v2 := make(map[uint32]uint32)
    var v3 int = 2147483647
    v4 := make(map[string]bool)
    var v5 uint = 42
    var v6 int32 = 0
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            if false {
                var v21 uint64 = 999
                switch fi19 {
                case 9223372036854775807:
                    fmt.Println(fi19, fi20, v21)
                case 128:
                    switch ((("hello" + ("hello" + "\n\t")) + ("hello" + ("" + ""))) + "foo bar") {
                    case "":
                        var v22 int8 = 42
                        _ = v22
                    default:
                        if (int64(fi20) < 42) {
                            var v23 uint = v5
                            _ = len(v4)
                            delete(v2, 2147483647)
                            _ = v23
                        }
                    }
                    var v24 uint16 = 32767
                    v25 := []int{(fi20 - fi19), (v3 % ((fi20) | 1)), fi20, (v3 | (-(fi20)))}
                    _ = v24
                    _ = v25
                case 2147483647:
                    v6 = (^(v6))
                }
                fmt.Println(fi20, fi20)
                _ = v21
            } else {
                fmt.Println(fi20, v3)
                v5 = v5
                switch v3 {
                case 0:
                    v26 := GenS2{F0: uint8(v5), F1: 594}
                    _ = v26
                }
                if false {
                    v1 = v1
                } else {
                    var v27 float64 = 3.14159
                    var v28 uint32 = 197
                    _ = v27
                    _ = v28
                }
                if false {
                    v29 := []int{fi20}
                    _ = v29
                }
            }
            _ = fi20
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := GenS1{F0: 849}
            var v32 int32 = (v6 ^ v6)
            var v33 string = "foo bar"
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        v5 = v5
        _ = fi19
    }
    v34 := []int{}
    v35 := GenS2{F0: 1, F1: 4294967295}
    if false {
        var v36 int16 = 32767
        fmt.Println(v5, v5, v36)
        v6 = (v6 ^ (-(v6)))
        _ = v36
    }
    var v37 int32 = -879
    var v38 int16 = 1000
    v1 -= (v1 - 0)
    var v39 int16 = v38
    v3 -= (-(v3))
    if true {
        v4[("a" + "\n\t")] = ((v1 ^ 563) != (^(v1)))
    } else {
        var v40 float32 = 1e10
        _ = v40
    }
    var v41 float64 = 1e10
    var v42 uint32 = 42
    switch (v3 | v3) {
    case 2147483647:
        v37 = (v6 >> (((v5 + (v5 | (^(v5))))) & 63))
    }
    switch v3 {
    case -2147483648:
        v3 = 1
    default:
        v43 := GenS2{F0: 0, F1: 1}
        _ = v43
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        switch "a" {
        case "\n\t":
            switch "" {
            case "":
                for fi45 := 0; fi45 < 11; fi45++ {
                    var v46 uint = genFn1(121, (true && (151 == v1)))
                    fmt.Println(v39)
                    _ = fi45
                    _ = v46
                }
            case "hello":
                if (42 != 725) {
                    fmt.Println(v38, v38, v6)
                } else {
                    var v47 int16 = -1
                    var v48 uint8 = 202
                    _ = v47
                    _ = v48
                }
            case "a":
                if false {
                    v49 := GenS2{F0: 1, F1: 4294967295}
                    _ = v49
                }
            default:
                v5 -= (v5 << ((2) & 63))
            }
        case "hello":
            v5 += genFn1((v5 << (((v5 & 0)) & 63)), true)
        case "":
            switch (v3 >> (((v5 >> ((v5) & 63))) & 63)) {
            case 128:
                var v50 int64 = 42
                _ = v50
            case 986:
                var v51 int8 = int8(fi44)
                _ = v51
            }
        }
        var v52 int = v3
        _ = fi44
        _ = v52
    }
    v38 += 899
    var v53 uint32 = v42
    v54 := GenS3{F0: v3, F1: (^(v53))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
