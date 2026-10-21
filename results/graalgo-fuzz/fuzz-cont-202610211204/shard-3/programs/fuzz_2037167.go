package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64, p1 int32, p2 uint16) uint8 {
    if (false || (false || false)) {
        fmt.Println(p0, p0)
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                if ((true && false) || false) {
                    v9 := []int{}
                    _ = v9
                }
                switch (fi7 - int(p2)) {
                case 32767:
                    fmt.Println(p0)
                default:
                    var v10 uint = 65535
                    _ = v10
                }
                _ = fi8
            }
            var v11 int32 = -210
            _ = fi7
            _ = v11
        }
    } else {
        v12 := []int{}
        _ = v12
    }
    if false {
        var v13 string = ("foo bar" + "foo bar")
        _ = v13
    } else {
        if (true || false) {
            var v14 uint16 = (p2 - 662)
            if true {
                switch -9223372036854775808 {
                case 0:
                    if false {
                        var v15 int8 = -128
                        _ = v15
                    }
                case 147:
                    for fi16 := 0; fi16 < 7; fi16++ {
                        p0 -= uint64(p2)
                        v17 := []int{}
                        _ = fi16
                        _ = v17
                    }
                default:
                    p1 = (^(p1))
                }
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v20 := GenS1{F0: 2, F1: -128, F2: (p2 >> ((279) & 63))}
                        var v21 string = "a"
                        var v22 float32 = (-2.25 * 0.0)
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := []int{}
                    var v24 uint8 = 255
                    _ = fi18
                    _ = v23
                    _ = v24
                }
            }
            _ = v14
        }
        fmt.Println(p2)
        var v25 float64 = 0.0
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if false {
            var v27 uint16 = p2
            _ = v27
        }
        p2 -= (^(p2))
        _ = fi26
    }
    var v28 int = -195
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 127
}
func genFn1(p0 uint32, p1 uint64) int64 {
    var v29 uint64 = 18446744073709551615
    switch 2147483647 {
    case 9223372036854775807:
        p1 -= p1
    case -1:
        var v30 int = 4294967295
        _ = v30
    case 127:
        v31 := GenS1{F0: 1, F1: -16, F2: 1}
        _ = v31
    default:
        var v32 int = -2147483648
        _ = v32
    }
    p1 -= 4294967295
    p0 = 0
    var v33 uint16 = 42
    _ = p0
    _ = p1
    _ = v29
    _ = v33
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 1
    var v3 int32 = (v2 + -259)
    var v4 uint16 = 881
    var v5 int32 = (v3 % ((1) | 1))
    var v6 string = "foo bar"
    fmt.Println(v2)
    fmt.Println(v2)
    var v34 int64 = 0
    var v35 int64 = (v34 | (-(v34)))
    v36 := GenS1{F0: 648, F1: 0, F2: 255}
    var v37 uint8 = 255
    v38 := GenS2{F0: 2}
    for fi39 := 0; fi39 < 14; fi39++ {
        switch fi39 {
        case 42:
            for fi40 := 0; fi40 < 10; fi40++ {
                v5 = v5
                v41 := v1[fi39]
                v42 := GenS1{F0: 2147483647, F1: 0, F2: (^(v4))}
                _ = fi40
                _ = v41
                _ = v42
            }
        }
        v43 := GenS2{F0: 292}
        v44 := GenS2{F0: 65535}
        _ = fi39
        _ = v43
        _ = v44
    }
    var v45 uint8 = 128
    v46 := GenS1{F0: 65535, F1: 1, F2: 7}
    var v47 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
