package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 uint) string {
    v7 := GenS1{F0: 1, F1: 7, F2: 1}
    v8 := GenS2{F0: 42, F1: 1}
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            p1 -= (p1 % (((^(p0))) | 1))
            var v10 int32 = -108
            _ = fi9
            _ = v10
        }
        switch ("hello" + ("hello" + ("a" + "foo bar"))) {
        case "a":
            v11 := []int{-528, int(p1), 1, -9223372036854775808}
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 9; fi12++ {
                if true {
                    v13 := GenS2{F0: 1, F1: 2}
                    v14 := GenS2{F0: 0, F1: 0}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int64 = 0
                    _ = v15
                }
                _ = fi12
            }
        case "\n\t":
            fmt.Println(p0, p1)
        default:
            if false {
                switch 728 {
                case -2147483648:
                    v16 := []int{586, 9223372036854775807}
                    _ = v16
                }
            }
        }
    } else {
        var v17 uint8 = 52
        if (!false) {
            if false {
                fmt.Println(p1)
                if (v17 != v17) {
                    var v18 int16 = 0
                    v17 = (v17 % ((1) | 1))
                    _ = v18
                } else {
                    v19 := []int{-2147483648, -2147483648, -624}
                    _ = v19
                }
                var v20 bool = true
                _ = v20
            } else {
                var v21 float32 = -2.25
                v22 := []int{-1, -573, int(v17), 32768}
                v21 = ((v21 - v21) * v21)
                _ = v21
                _ = v22
            }
            p0 -= 2
        } else {
            fmt.Println(v17, p1)
        }
        v23 := []int{}
        _ = v17
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint = (p0 >> ((p0) & 63))
            _ = fi25
            _ = v26
        }
        v27 := GenS2{F0: 0, F1: 2147483648}
        _ = fi24
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []int{32768, 1, 1}
    v3 := make(map[string]bool)
    v4 := []string{"", "", "a", ("" + "\n\t")}
    v5 := []string{("hello" + "\n\t")}
    v6 := []int{1, 127, -9223372036854775808}
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 uint64 = 18446744073709551615
            _ = fi29
            _ = v30
        }
        var v31 uint64 = 837
        switch "" {
        case "a":
            v32 := GenS2{F0: -1, F1: 0}
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := GenS2{F0: -89, F1: 42}
                v35 := GenS3{F0: (!(!(!true)))}
                switch -135 {
                case 42:
                    v31 = 18446744073709551615
                default:
                    if true {
                        v36 := GenS4{F0: 2, F1: 2}
                        fmt.Println(fi28)
                        v31 += v31
                        _ = v36
                    }
                }
                _ = fi33
                _ = v34
                _ = v35
            }
        default:
            v31 -= 1
        }
        _ = fi28
        _ = v31
    }
    v37 := GenS2{F0: -53, F1: 0}
    var v38 uint32 = 4294967295
    var v39 int16 = -1
    if true {
        v40 := GenS1{F0: 19, F1: -128, F2: 164}
        _ = v40
    }
    var v41 bool = true
    v38 += v38
    var v42 float32 = 0.1
    var v43 float32 = ((((v42 - v42) - -2.25) + float32(((0.0 - 1e10) - 1e10))) / 1.0)
    var v44 uint64 = 2
    switch -2147483648 {
    case 4294967295:
        var v45 uint = 2
        _ = v45
    case 128:
        var v46 int16 = (-(v39))
        _ = v46
    case 2147483647:
        v42 += (v42 - (1e10 / (((1.5 / float32((1.5 - (0.0 / 1e10)))) / v42) + -2.25)))
    default:
        v47 := GenS5{F0: 833, F1: ("a" + "\n\t")}
        _ = v47
    }
    v41 = ((true && v41) && (237 < 1))
    switch ("a" + ("a" + "")) {
    case "":
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 string = ""
            _ = fi48
            _ = v49
        }
    }
    var v50 int = -1
    var v51 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
