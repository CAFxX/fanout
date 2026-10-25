package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 uint32) uint16 {
    var v4 uint = 255
    switch 127 {
    case 32767:
        p0 += p0
    case 304:
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint64 = 0
            switch ("" + ("" + ("" + ("" + ("a" + "a"))))) {
            case "foo bar":
                v7 := GenS1{F0: 970}
                _ = v7
            case "":
                p0 = (p0 / (((p0 / ((0) | 1))) | 1))
            case "\n\t":
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := GenS1{F0: 42}
                    v4 = v4
                    _ = fi8
                    _ = v9
                }
            default:
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(v4, fi5)
                    fmt.Println(v4, fi10)
                    var v11 int16 = 42
                    if ((0.0 + (0.0 * 0.0)) != 3.14159) {
                        v12 := GenS2{F0: (fi5 | (fi5 ^ 0))}
                        v13 := []int{(fi10 ^ fi5), (fi5 * (fi5 | 32768)), fi10}
                        fmt.Println(v11)
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: int(p0)}
                    _ = fi10
                    _ = v11
                    _ = v14
                }
            }
            _ = fi5
            _ = v6
        }
    case -2147483648:
        p0 += 1
    default:
        v15 := []int{}
        _ = v15
    }
    _ = p0
    _ = v4
    return 86
}

func main() {
    v1 := []int{128}
    v2 := []string{("\n\t" + "\n\t"), "hello"}
    v3 := []string{}
    var v16 bool = true
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 string = "foo bar"
        if (0 > 0) {
            var v19 uint = 758
            var v20 float64 = float64(0.1)
            if v16 {
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS3{F0: true, F1: (v19 | v19)}
                    switch v18 {
                    case "":
                        var v23 uint = 463
                        _ = v23
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = v18
                        }
                    case "\n\t":
                        fmt.Println(fi17, v16)
                    }
                    v24 := GenS3{F0: false, F1: (v19 << ((v19) & 63))}
                    _ = fi21
                    _ = v22
                    _ = v24
                }
            } else {
                v18 = v18
                v3 = append(v3, v18)
            }
            _ = v19
            _ = v20
        } else {
            if len(v1) > 0 {
                v1[0] = (fi17 >> ((780) & 63))
            }
            var v25 int8 = -128
            var v26 int = fi17
            _ = v25
            _ = v26
        }
        _ = fi17
        _ = v18
    }
    if (v16 && false) {
        for fi27 := 0; fi27 < 7; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float64 = 0.0
                _ = fi28
                _ = v29
            }
            var v30 uint16 = 255
            _ = fi27
            _ = v30
        }
        v16 = v16
    }
    var v31 uint8 = 58
    if (!(!v16)) {
        for fi32 := 0; fi32 < 17; fi32++ {
            v16 = (v16 || v16)
            _ = fi32
        }
        var v33 int = 42
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        v35 := GenS4{F0: 9223372036854775807, F1: 42}
        for fi36 := 0; fi36 < 9; fi36++ {
            fmt.Println(fi36, fi36)
            fmt.Println(v16, v31, v31)
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    var v37 int64 = 0
    var v38 uint32 = 79
    var v39 int8 = -128
    for fi40 := 0; fi40 < 10; fi40++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi40
    }
    var v41 int8 = (v39 | (v39 % ((-6) | 1)))
    if v16 {
        v37 = -257
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v41
}
