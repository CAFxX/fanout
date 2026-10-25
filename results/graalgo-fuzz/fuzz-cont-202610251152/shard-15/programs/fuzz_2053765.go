package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 int32, p1 uint) bool {
    if ((p0 - p0) != p0) {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{(fi6 ^ fi6), (-(fi6)), (fi6 + fi6)}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = -2.25
            v10 := GenS1{F0: 206}
            switch 2147483647 {
            case 2147483647:
                var v11 string = ("" + "\n\t")
                _ = v11
            case 42:
                var v12 int32 = 2147483647
                _ = v12
            default:
                v13 := GenS2{F0: 0, F1: 42}
                _ = v13
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v14 := []int{1}
    switch ("a" + "\n\t") {
    case "foo bar":
        v15 := GenS1{F0: 65535}
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        switch (fi16 - (fi16 + 2147483647)) {
        case 1:
            fmt.Println(fi16)
        case -1:
            var v17 uint16 = 42
            _ = v17
        case 32768:
            for fi18 := 0; fi18 < 19; fi18++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                switch "" {
                case "foo bar":
                    fmt.Println(fi16, fi18, p0)
                case "\n\t":
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                case "hello":
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                default:
                    v19 := GenS2{F0: 71, F1: 215}
                    _ = v19
                }
                p1 = 42
                _ = fi18
            }
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v14
    return (1 >= 267)
}
func genFn1(p0 bool, p1 int32, p2 int) int16 {
    if false {
        v20 := []int{p2, (p2 + -478), 42, int(p1)}
        _ = len(v20)
        p2 += (p2 ^ (p2 & (^(p2))))
        _ = v20
    } else {
        p0 = ("hello" != "\n\t")
        if (3.14159 == 1e10) {
            var v21 int32 = 2147483647
            p2 -= p2
            fmt.Println(p0, p0, v21)
            _ = v21
        } else {
            p2 += (p2 - (-(p2)))
            var v22 uint8 = 207
            var v23 int = -1
            _ = v22
            _ = v23
        }
    }
    p2 += p2
    p1 += (p1 + (p1 % (((^(p1))) | 1)))
    switch "" {
    case "\n\t":
        v24 := []int{}
        _ = v24
    case "a":
        var v25 float64 = 0.1
        _ = v25
    default:
        p0 = false
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    v4 := make(map[int]string)
    var v5 uint32 = 0
    v1 -= 876
    var v26 uint16 = v1
    v5 = (v5 / ((215) | 1))
    if (!((("" + "\n\t") <= ("foo bar" + "hello")) || true)) {
        v27 := []int{-2147483648, 0, 127}
        if true {
            v28 := GenS1{F0: (v5 - 967)}
            for fi29 := 0; fi29 < 17; fi29++ {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    for fi30 := 0; fi30 < 19; fi30++ {
                        if ((v3 / v3) <= (1e10 + 1e10)) {
                            if true {
                                v26 -= v26
                                v31 := GenS1{F0: (^(v5))}
                                _ = v31
                            }
                            switch ("" + "foo bar") {
                            case "a":
                                var v32 float32 = -2.25
                                var v33 int16 = -32768
                                _ = v32
                                _ = v33
                            case "":
                                v1 += v1
                            }
                            v3 = (float32(0.0) / (1e10 - v3))
                        } else {
                            var v34 bool = true
                            fmt.Println(fi30, v3)
                            v35 := GenS1{F0: v5}
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                    }
                case "":
                    var v36 int16 = int16(fi29)
                    _ = v36
                }
                _ = fi29
            }
            var v37 uint8 = 42
            _ = v28
            _ = v37
        }
        _ = v27
    }
    switch 2147483647 {
    case -938:
        v38 := v4[4294967295]
        _ = v38
    case 2147483647:
        v39 := []int{}
        _ = v39
    }
    if (!(!(!true))) {
        v26 -= (^(v26))
        var v40 uint32 = v5
        var v41 uint16 = (v26 >> ((4294967295) & 63))
        _ = v40
        _ = v41
    }
    fmt.Println(v3, v26)
    var v42 bool = (0 > -128)
    fmt.Println(v42, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
