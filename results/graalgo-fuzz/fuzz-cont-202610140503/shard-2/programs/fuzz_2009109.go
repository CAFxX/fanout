package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() int8 {
    v4 := GenS1{F0: -456}
    var v5 bool = true
    switch -589 {
    case 128:
        if ((false || (3.14159 > float32(0.1))) || v5) {
            for fi6 := 0; fi6 < 15; fi6++ {
                if (127 == 2) {
                    for fi7 := 0; fi7 < 8; fi7++ {
                        v8 := []int{}
                        v9 := GenS1{F0: 0}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                }
                v10 := []int{fi6, -1, (fi6 * (fi6 & 325))}
                _ = fi6
                _ = v10
            }
        } else {
            v11 := GenS1{F0: 949}
            _ = v11
        }
    case 980:
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 = v5
            _ = fi12
        }
    case -617:
        var v13 uint64 = 42
        _ = v13
    default:
        var v14 uint = 2
        _ = v14
    }
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 bool, p2 int) int {
    var v15 uint8 = 2
    v16 := []int{(^(p2)), (-(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return (p2 ^ (p2 & p2))
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]bool)
    var v3 int64 = 42
    fmt.Println(v1)
    switch "\n\t" {
    case "hello":
        switch genFn1(985, (0.0 == -2.25), (v1 << ((2) & 63))) {
        case 128:
            for fi17 := 0; fi17 < 14; fi17++ {
                switch "" {
                case "hello":
                    v3 += -943
                default:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v1 += fi18
                        var v19 uint = 255
                        v20 := []int{(^(fi18)), v1, (fi17 ^ (v1 & (fi18 & -845))), fi17}
                        _ = len(v20)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi17
            }
        case -518:
            fmt.Println(v1, v1)
        }
    case "a":
        var v21 uint16 = 0
        _ = v21
    case "\n\t":
        if (false || (!true)) {
            var v22 string = "foo bar"
            _ = v22
        } else {
            for fi23 := 0; fi23 < 2; fi23++ {
                v1 = (v1 >> ((0) & 63))
                _ = fi23
            }
            if (0.0 >= 1.5) {
                var v24 uint8 = 255
                _ = v24
            }
            var v25 uint64 = 211
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        switch "a" {
        case "":
            var v27 uint32 = 4294967295
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 16; fi28++ {
                if (v1 >= fi26) {
                    v29 := GenS1{F0: 9223372036854775807}
                    var v30 int8 = -128
                    switch fi28 {
                    case 32768:
                        switch fi26 {
                        case 32768:
                            v3 += (v3 & v3)
                        case 1:
                            v31 := []int{(-(fi26)), v1}
                            _ = v31
                        case 0:
                            v32 := GenS2{F0: false, F1: 1}
                            _ = v32
                        }
                    case 9223372036854775807:
                        v30 = (^(v30))
                    case -1:
                        var v33 uint16 = 702
                        _ = v33
                    }
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
        }
        if ((!false) && false) {
            v34 := v2[""]
            v1 = (v1 & -381)
            v35 := []int{(fi26 >> ((65535) & 63)), (fi26 + 858), (fi26 / (((-(v1))) | 1))}
            _ = v34
            _ = v35
        }
        _ = fi26
    }
    var v36 uint32 = 608
    var v37 string = ""
    fmt.Println(v1)
    fmt.Println(v37, v3, v3)
    for fi38 := 0; fi38 < 19; fi38++ {
        v1 += (fi38 - fi38)
        switch "a" {
        case "hello":
            if false {
                var v39 int64 = v3
                _ = v39
            } else {
                var v40 uint16 = 32767
                _ = v40
            }
        case "":
            var v41 uint8 = 42
            _ = v41
        case "foo bar":
            for fi42 := 0; fi42 < 5; fi42++ {
                v37 = (v37 + v37)
                var v43 int16 = 0
                _ = fi42
                _ = v43
            }
        }
        fmt.Println(v37, v36)
        _ = fi38
    }
    fmt.Println(v36)
    var v44 int32 = -893
    var v45 int16 = 434
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v44
    _ = v45
}
