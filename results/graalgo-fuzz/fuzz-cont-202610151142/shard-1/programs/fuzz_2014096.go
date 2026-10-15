package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int) uint32 {
    v7 := []int{p0, (p0 & (p0 % ((p0) | 1))), p0, (p0 + (^(p0)))}
    v7 = append(v7, p0)
    fmt.Println(p0)
    var v8 int = (p0 % ((p0) | 1))
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint {
    var v9 float32 = float32((0.0 + -2.25))
    var v10 int = -2147483648
    _ = v10
    _ = v9
    return 486
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 42
    var v3 uint64 = v2
    v4 := make(map[uint32]int)
    var v5 int16 = 553
    v6 := []uint32{852}
    if true {
        v4[42] = 738
    }
    _ = len(v6)
    if true {
        v3 = (v3 % ((42) | 1))
        v6 = append(v6, 2147483648)
        fmt.Println(v5, v5)
    }
    switch (("" + "foo bar") + "") {
    case "":
        var v11 float32 = (1e10 / -2.25)
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 11; fi12++ {
            if true {
                var v13 int32 = v1
                var v14 bool = true
                _ = v13
                _ = v14
            }
            _ = fi12
        }
    case "\n\t":
        var v15 uint64 = v3
        _ = v15
    default:
        var v16 int32 = v1
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: int8(fi17), F1: "foo bar"}
        _ = fi17
        _ = v18
    }
    v6 = append(v6, 1)
    fmt.Println(v1)
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            v20 := v4[2147483648]
            var v21 int8 = -128
            var v22 int32 = 1
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    v23 := GenS1{F0: 100, F1: "\n\t"}
    v24 := GenS1{F0: 100, F1: (("a" + ("\n\t" + "foo bar")) + "")}
    v25 := GenS2{F0: "\n\t", F1: v5}
    v26 := GenS2{F0: "hello", F1: v5}
    v27 := v4[2]
    var v28 uint8 = 2
    if false {
        if ((1.5 * ((0.1 / (1e10 + (1e10 + (3.14159 / ((3.14159 * ((float64(0.0) - 1.5) + 3.14159)) - -2.25))))) + -2.25)) < float64(0.1)) {
            if ((v3 >> ((209) & 63)) <= (v3 - 0)) {
                if (-1 > (v1 - (v1 ^ int32(v2)))) {
                    v29 := v4[4294967295]
                    var v30 int32 = (v1 + v1)
                    var v31 uint8 = (^(v28))
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    v32 := GenS3{F0: 2, F1: v28}
                    v33 := GenS4{F0: "foo bar"}
                    _ = v32
                    _ = v33
                }
                switch (v27 ^ v27) {
                case -783:
                    fmt.Println(v2, v27)
                case 4294967295:
                    var v34 uint64 = (v3 ^ v2)
                    _ = v34
                case 693:
                    _ = len(v4)
                }
                for fi35 := 0; fi35 < 17; fi35++ {
                    fmt.Println(v3, v1)
                    _ = fi35
                }
            } else {
                fmt.Println(v1)
            }
        } else {
            fmt.Println(v2, v1, v1)
            if true {
                if (0 < v27) {
                    for fi36 := 0; fi36 < 12; fi36++ {
                        var v37 bool = (((1e10 + -2.25) - (0.1 * 1e10)) >= (-2.25 / 1e10))
                        _ = fi36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 14; fi38++ {
                        for fi39 := 0; fi39 < 13; fi39++ {
                            v5 = 0
                            _ = fi39
                        }
                        var v40 int32 = (v1 / (((v1 | v1)) | 1))
                        _ = fi38
                        _ = v40
                    }
                }
                var v41 int32 = -1
                var v42 int16 = int16(v28)
                _ = v41
                _ = v42
            }
            var v43 uint8 = (v28 - 1)
            _ = v43
        }
    } else {
        v27 = v27
        var v44 string = "foo bar"
        v45 := GenS1{F0: 1, F1: (v44 + (v44 + (v44 + (("foo bar" + v44) + "hello"))))}
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
