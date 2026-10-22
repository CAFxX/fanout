package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 bool, p1 uint32) uint16 {
    v4 := []int{-2147483648, -651}
    p1 = p1
    var v5 uint16 = 1
    v6 := GenS1{F0: "hello"}
    var v7 int16 = 823
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v5
}
func genFn1(p0 uint64, p1 string, p2 int64) int64 {
    p2 -= (-(p2))
    switch "" {
    case "foo bar":
        for fi8 := 0; fi8 < 17; fi8++ {
            if (-8 == 103) {
                var v9 float64 = 0.1
                var v10 string = p1
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    case "\n\t":
        for fi11 := 0; fi11 < 5; fi11++ {
            if (-128 == -1) {
                p1 = "hello"
            } else {
                fmt.Println(p0, fi11, p1)
                var v12 bool = false
                v13 := []int{}
                _ = v12
                _ = v13
            }
            _ = fi11
        }
    case "":
        v14 := GenS1{F0: p1}
        _ = v14
    default:
        var v15 int = 9223372036854775807
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v16 uint = 1
    var v17 int8 = 127
    var v18 bool = (!true)
    switch "\n\t" {
    case "foo bar":
        switch 127 {
        case 127:
            if v18 {
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch -1 {
                    case 127:
                        v3 = int16(v16)
                    case 128:
                        if false {
                            v20 := []int{(fi19 * 9223372036854775807), (^(fi19))}
                            v3 = v3
                            _ = v20
                        }
                    case 9223372036854775807:
                        v21 := GenS2{F0: genFn0(v18, 42)}
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 3; fi22++ {
                            v18 = v18
                            var v23 int8 = v17
                            var v24 int = 127
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    v25 := GenS1{F0: "hello"}
                    _ = fi19
                    _ = v25
                }
            } else {
                v3 -= (v3 / (((v3 >> (((v16 - 255)) & 63))) | 1))
                v26 := []int{}
                _ = v26
            }
        case 32768:
            v3 = 0
        default:
            v27 := GenS1{F0: (((((("hello" + ("a" + "hello")) + (((("" + "a") + "") + "a") + "a")) + ((((("" + "a") + ((("" + "\n\t") + "hello") + "a")) + (("\n\t" + "a") + ("" + "\n\t"))) + "foo bar") + "foo bar")) + "foo bar") + "hello") + "hello")}
            _ = v27
        }
    case "hello":
        v28 := GenS3{F0: (v17 % ((v17) | 1)), F1: 128, F2: v1}
        _ = v28
    default:
        v3 += v3
    }
    if ((!(650 < 668)) || v18) {
        v29 := GenS1{F0: "hello"}
        v16 -= (^(v16))
        var v30 int8 = (v17 & v17)
        _ = v29
        _ = v30
    } else {
        v3 = (v3 << (((^(v16))) & 63))
        v18 = v18
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        if ((!true) || (42 < 255)) {
            var v32 float32 = float32((1.5 + (3.14159 / 1e10)))
            v1 = v1
            _ = v32
        } else {
            v33 := GenS2{F0: 886}
            fmt.Println(v1, v3)
            _ = v33
        }
        v1 -= (v1 >> ((18446744073709551615) & 63))
        _ = fi31
    }
    var v34 bool = v18
    for fi35 := 0; fi35 < 10; fi35++ {
        v17 -= (^(v17))
        _ = len(v2)
        v36 := []int{(fi35 - (fi35 * (fi35 ^ fi35))), fi35}
        _ = fi35
        _ = v36
    }
    if true {
        var v37 string = "foo bar"
        v38 := GenS3{F0: (-(v17)), F1: 128, F2: v1}
        for fi39 := 0; fi39 < 14; fi39++ {
            v40 := GenS2{F0: 2}
            for fi41 := 0; fi41 < 2; fi41++ {
                if true {
                    var v42 int8 = int8(fi39)
                    _ = v42
                } else {
                    switch fi39 {
                    case 1:
                        v34 = (!(v34 || v18))
                    }
                    v16 = v16
                }
                v43 := GenS2{F0: 1}
                _ = fi41
                _ = v43
            }
            var v44 string = v37
            _ = fi39
            _ = v40
            _ = v44
        }
        _ = v37
        _ = v38
    }
    var v45 float64 = 3.14159
    v46 := v2[-289]
    var v47 int16 = (v3 & (v3 & (v3 / ((1) | 1))))
    if (v34 || false) {
        fmt.Println(v3, v16, v45)
    }
    v48 := GenS1{F0: ""}
    v49 := GenS4{F0: 65535, F1: v17, F2: ("hello" < "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
