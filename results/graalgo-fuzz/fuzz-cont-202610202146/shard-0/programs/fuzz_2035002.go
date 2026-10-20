package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint, p2 int32) int8 {
    fmt.Println(p0, p1, p0)
    for fi4 := 0; fi4 < 11; fi4++ {
        switch -445 {
        case 2147483647:
            for fi5 := 0; fi5 < 14; fi5++ {
                switch ("foo bar" + (("hello" + "") + "")) {
                case "foo bar":
                    v6 := []int{fi4, (fi4 >> ((0) & 63)), (fi4 + fi5), (fi5 - -532)}
                    _ = v6
                case "":
                    v7 := []int{int(p0), (fi4 << ((p1) & 63))}
                    _ = v7
                }
                _ = fi5
            }
        default:
            switch fi4 {
            case 1:
                var v8 int8 = (p0 | 0)
                _ = v8
            case -60:
                v9 := GenS1{F0: 42, F1: 789, F2: "\n\t"}
                _ = v9
            default:
                var v10 uint32 = 42
                _ = v10
            }
        }
        v11 := []int{(fi4 | fi4), (fi4 + (fi4 % ((2147483647) | 1))), (fi4 ^ (fi4 | fi4))}
        _ = fi4
        _ = v11
    }
    v12 := GenS1{F0: 794, F1: -32768, F2: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p0 & 1)
}
func genFn1(p0 uint, p1 int32) uint32 {
    fmt.Println(p0, p1)
    if false {
        var v13 int32 = -2147483648
        _ = v13
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            p1 -= p1
            _ = fi14
        }
        var v15 int8 = genFn0(127, p0, (p1 & 2147483647))
        _ = v15
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v16 float64 = v1
    v3[-2147483648] = "\n\t"
    if (100 <= 100) {
        for fi17 := 0; fi17 < 13; fi17++ {
            if ("hello" <= "hello") {
                _ = len(v3)
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(fi17)
                    switch "\n\t" {
                    case "\n\t":
                        var v19 bool = (330 > 888)
                        v16 -= ((v16 / v16) - v16)
                        v16 += (1e10 * v16)
                        _ = v19
                    case "foo bar":
                        v2[fi18] = ((32767 > 65535) && (((1000 == 0) && (2 < 4294967295)) && false))
                    default:
                        var v20 int16 = -1
                        _ = len(v3)
                        fmt.Println(fi17, fi18, fi17)
                        _ = v20
                    }
                    _ = fi18
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    _ = len(v2)
                    _ = fi21
                }
            } else {
                switch "hello" {
                case "\n\t":
                    var v22 int64 = 42
                    _ = v22
                case "foo bar":
                    v23 := GenS1{F0: 255, F1: -269, F2: ("foo bar" + "")}
                    _ = v23
                }
                fmt.Println(fi17, v16, v16)
                switch fi17 {
                case -339:
                    if (!((((0.0 * (float32(v16) - (1.5 + 0.1))) / 1.5) < 0.1) && false)) {
                        v16 -= (0.1 * (v1 - (float64((1e10 * 3.14159)) * (v1 + 1.5))))
                        var v24 float64 = v1
                        _ = v24
                    }
                case 32768:
                    fmt.Println(v1)
                case 840:
                    if false {
                        v25 := []int{(fi17 - fi17), -9223372036854775808, (^(fi17))}
                        var v26 bool = (1 < 65535)
                        var v27 string = "foo bar"
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
            }
            v28 := GenS1{F0: 32767, F1: 1, F2: "\n\t"}
            _ = fi17
            _ = v28
        }
    } else {
        v1 = v16
        switch "" {
        case "a":
            var v29 int16 = -736
            _ = v29
        case "foo bar":
            var v30 string = ""
            _ = v30
        default:
            switch 1 {
            case -1:
                v31 := GenS1{F0: 0, F1: 1, F2: "a"}
                _ = v31
            case 0:
                if ((4294967295 > genFn1(1, -2147483648)) && true) {
                    var v32 int32 = 0
                    var v33 uint = 42
                    fmt.Println(v32, v33, v16)
                    _ = v32
                    _ = v33
                }
            default:
                for fi34 := 0; fi34 < 14; fi34++ {
                    var v35 uint = 255
                    fmt.Println(fi34)
                    var v36 float64 = v1
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
    }
    _ = len(v2)
    v1 += -2.25
    fmt.Println(v1)
    v37 := GenS2{F0: 2}
    if false {
        var v38 uint32 = 2147483648
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v40 := []int{127, fi39, 4294967295, -1}
        fmt.Println(v16, v1)
        _ = fi39
        _ = v40
    }
    fmt.Println(v16)
    v16 = -2.25
    v16 -= v1
    v41 := []int{}
    delete(v2, 128)
    var v42 uint8 = 255
    v1 = (v16 - v16)
    var v43 uint = 2
    var v44 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
