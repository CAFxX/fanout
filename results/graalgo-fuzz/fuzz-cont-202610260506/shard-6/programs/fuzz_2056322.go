package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint8, p1 int32) int16 {
    p1 -= p1
    switch "a" {
    case "":
        for fi7 := 0; fi7 < 9; fi7++ {
            p1 -= p1
            if (false || true) {
                if (!false) {
                    v8 := GenS1{F0: (fi7 % (((-(fi7))) | 1)), F1: 946, F2: (42 == 100)}
                    p1 += (p1 >> ((255) & 63))
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 bool = (false && true)
                        var v11 string = ""
                        v12 := []int{fi9, (fi9 ^ fi7), (fi9 + fi7), fi9}
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    _ = v8
                } else {
                    v13 := GenS1{F0: 32768, F1: 1, F2: (42 <= 0)}
                    if false {
                        var v14 int16 = -1
                        _ = v14
                    }
                    _ = v13
                }
                v15 := GenS1{F0: (fi7 << ((4294967295) & 63)), F1: 0, F2: false}
                var v16 bool = true
                _ = v15
                _ = v16
            }
            _ = fi7
        }
    case "a":
        if true {
            var v17 string = (("hello" + "foo bar") + ("hello" + "foo bar"))
            if true {
                var v18 float64 = 0.1
                v19 := GenS1{F0: -351, F1: 65535, F2: false}
                _ = v18
                _ = v19
            } else {
                p0 = p0
            }
            v20 := GenS2{F0: 621, F1: 858}
            _ = v17
            _ = v20
        } else {
            fmt.Println(p1, p1)
            var v21 uint = 111
            v22 := []int{}
            _ = v21
            _ = v22
        }
    }
    var v23 uint32 = 897
    v24 := GenS1{F0: -680, F1: 32767, F2: (((42 < 42) || true) && ((1 != 32767) || true))}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 42
}
func genFn1(p0 uint, p1 int16) uint64 {
    for fi25 := 0; fi25 < 1; fi25++ {
        switch fi25 {
        case 556:
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 uint16 = 42
                var v28 uint16 = (^(v27))
                _ = fi26
                _ = v27
                _ = v28
            }
        case -1:
            if ((!(true && true)) || false) {
                var v29 bool = false
                p0 -= p0
                switch ("foo bar" + "\n\t") {
                case "a":
                    var v30 uint = (p0 % ((p0) | 1))
                    _ = v30
                default:
                    p0 += p0
                }
                _ = v29
            } else {
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(p1, fi25, p1)
                    v32 := GenS1{F0: fi25, F1: 143, F2: (false && ((true && ((!(!false)) || ((((0.0 * 1.5) / (1e10 - 0.1)) - 3.14159) >= 0.1))) || (0 != 2147483647)))}
                    if true {
                        var v33 float32 = -2.25
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
                for fi34 := 0; fi34 < 9; fi34++ {
                    v35 := []int{}
                    p1 = (p1 % ((0) | 1))
                    _ = fi34
                    _ = v35
                }
            }
        case -2147483648:
            v36 := GenS1{F0: fi25, F1: 192, F2: true}
            _ = v36
        }
        if (1000 >= genFn0(0, -950)) {
            switch 669 {
            case 128:
                switch "foo bar" {
                case "":
                    p0 = (p0 | (p0 / ((1) | 1)))
                case "foo bar":
                    switch "\n\t" {
                    case "":
                        p1 -= genFn0(42, -2147483648)
                    case "foo bar":
                        v37 := GenS2{F0: p0, F1: 42}
                        _ = v37
                    default:
                        v38 := []int{}
                        _ = v38
                    }
                default:
                    var v39 uint32 = 295
                    _ = v39
                }
            case 18:
                fmt.Println(fi25, p0, fi25)
            }
            var v40 int32 = 146
            _ = v40
        }
        v41 := []int{fi25, (^(fi25))}
        _ = fi25
        _ = v41
    }
    v42 := []int{0, 127}
    _ = p0
    _ = p1
    _ = v42
    return 220
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = (("hello" + "foo bar") + "a")
    v3 := make(map[int]string)
    var v4 bool = (2 >= 78)
    v5 := make(map[int]bool)
    v6 := make(map[int]int)
    v43 := GenS2{F0: 2, F1: 42}
    var v44 float32 = 1e10
    v2 = v2
    v45 := GenS1{F0: -484, F1: 65535, F2: v4}
    v4 = (v4 && false)
    var v46 uint8 = 2
    v46 -= (v46 << ((0) & 63))
    v3[-1] = "foo bar"
    var v47 uint64 = 985
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
