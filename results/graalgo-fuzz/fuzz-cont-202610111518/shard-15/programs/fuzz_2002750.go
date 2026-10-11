package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
func genFn0() int16 {
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: 405, F1: 2, F2: 128}
                fmt.Println(v5, fi7)
                _ = fi7
                _ = v8
            }
        }
        if true {
            v9 := []int{fi6, (fi6 ^ fi6), fi6, 1}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{fi6}
                v5 = (v5 | -1)
                _ = fi10
                _ = v11
            }
            v5 = 1
            for fi12 := 0; fi12 < 19; fi12++ {
                v5 = -32768
                if true {
                    switch "" {
                    case "hello":
                        var v13 uint16 = 65535
                        _ = v13
                    case "foo bar":
                        fmt.Println(v5, fi6, fi6)
                    default:
                        fmt.Println(fi12, fi12)
                    }
                    v5 = (v5 & v5)
                    if false {
                        if (!true) {
                            v14 := []int{}
                            _ = v14
                        }
                        v5 += (v5 + (v5 & (v5 >> ((18446744073709551615) & 63))))
                    }
                    var v15 string = "\n\t"
                    _ = v15
                }
                var v16 float64 = 3.14159
                v5 = v5
                _ = fi12
                _ = v16
            }
        }
        if false {
            for fi17 := 0; fi17 < 9; fi17++ {
                switch fi17 {
                case 4294967295:
                    var v18 uint8 = 128
                    _ = v18
                default:
                    v19 := []int{0, fi6, 9223372036854775807}
                    _ = v19
                }
                _ = fi17
            }
            switch ("hello" + ("\n\t" + ("hello" + ""))) {
            case "a":
                v5 = v5
            default:
                var v20 int16 = (-(v5))
                _ = v20
            }
            v5 = 0
        } else {
            v5 += v5
        }
        _ = fi6
    }
    v5 += -453
    var v21 string = ("\n\t" + "hello")
    _ = v21
    _ = v5
    return (v5 / ((v5) | 1))
}
func genFn1(p0 uint64, p1 bool) string {
    var v22 int = 305
    switch "a" {
    case "a":
        var v23 bool = p1
        _ = v23
    case "hello":
        var v24 int = (v22 % (((v22 << ((815) & 63))) | 1))
        _ = v24
    default:
        var v25 int8 = 119
        _ = v25
    }
    fmt.Println(p1, v22)
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint8 = 82
        fmt.Println(p0, v22, v22)
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v22
    return (("" + "foo bar") + "\n\t")
}

func main() {
    v1 := []uint32{}
    var v2 float64 = (((1.5 + ((((-2.25 + 3.14159) / 1e10) / 3.14159) / -2.25)) * (0.0 / 1.5)) / 1.5)
    v3 := make(map[uint32]string)
    var v4 int64 = 42
    v4 -= (v4 | 0)
    var v28 int = 32768
    v29 := GenS2{F0: 0}
    if (true || (3 != 10)) {
        v30 := GenS1{F0: v4, F1: 122, F2: 134}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v31 int32 = -842
            v2 = v2
            var v32 int8 = 127
            _ = v31
            _ = v32
        } else {
            var v33 int16 = 42
            v34 := GenS2{F0: (v28 / ((128) | 1))}
            if (!(((("\n\t" + "a") + "\n\t") + "a") <= "")) {
                switch (v28 << ((1) & 63)) {
                case 1:
                    v4 = (v4 % (((v4 + 0)) | 1))
                }
                if (!((v33 + v33) <= 32767)) {
                    v28 = v28
                    switch "foo bar" {
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "":
                        v2 = v2
                    }
                }
            }
            _ = v33
            _ = v34
        }
        _ = v30
    } else {
        var v35 uint64 = 114
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 int8 = -41
            _ = fi36
            _ = v37
        }
        var v38 float32 = -2.25
        _ = v35
        _ = v38
    }
    v2 += 3.14159
    if true {
        var v39 int64 = v4
        v1 = append(v1, 2)
        if (false && (722 == 4294967295)) {
            v40 := GenS3{F0: (v4 / ((0) | 1))}
            if false {
                v41 := GenS4{F0: 255}
                _ = v41
            }
            _ = v40
        }
        _ = v39
    } else {
        var v42 uint8 = 198
        var v43 int64 = (v4 % ((v4) | 1))
        _ = v42
        _ = v43
    }
    v28 -= int(v4)
    v28 += v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
