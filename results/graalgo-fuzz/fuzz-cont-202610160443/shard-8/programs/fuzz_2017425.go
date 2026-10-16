package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 int32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int64; F1 string; F2 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint) int32 {
    if false {
        v7 := GenS1{F0: 14}
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    v9 := []int{}
    p0 += 2
    _ = p0
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64, p1 uint16, p2 bool) string {
    switch (("" + "a") + "hello") {
    case "\n\t":
        if ((p1 / ((p1) | 1)) >= 44) {
            fmt.Println(p2, p0)
            v10 := []int{32768}
            _ = v10
        }
    case "":
        p0 += p0
    }
    if p2 {
        v11 := GenS1{F0: 255}
        _ = v11
    } else {
        v12 := []int{-2147483648}
        p0 -= (p0 * p0)
        p2 = p2
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + "\n\t")
}

func main() {
    var v1 uint = 42
    v2 := []int{}
    var v3 float32 = 1.5
    var v4 uint32 = 434
    v5 := []uint32{}
    v6 := make(map[uint32]uint32)
    v13 := GenS2{F0: 2147483647, F1: "hello", F2: 42}
    var v14 uint16 = 1
    _ = len(v5)
    if (860 != (^(v1))) {
        v1 += 311
        v1 = 18446744073709551615
    } else {
        delete(v6, (^(v4)))
        if true {
            switch ("" + "foo bar") {
            case "hello":
                var v15 int = 32767
                _ = v15
            default:
                v16 := GenS2{F0: v4, F1: "\n\t", F2: genFn0((^(v1)))}
                _ = v16
            }
            switch 1 {
            case 528:
                v3 -= 3.14159
            default:
                v4 = (v4 >> ((v1) & 63))
            }
            v14 = (v14 - (v14 - (v14 * 153)))
        } else {
            for fi17 := 0; fi17 < 17; fi17++ {
                if true {
                    var v18 int8 = 18
                    _ = v18
                } else {
                    if len(v2) > 0 {
                        v2[0] = fi17
                    }
                    v14 = 0
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int16 = 32767
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch ("\n\t" + (("hello" + "\n\t") + ((("" + "foo bar") + "a") + (("hello" + (("foo bar" + "foo bar") + "")) + "foo bar")))) {
                        case "a":
                            var v21 string = ""
                            _ = v21
                        case "foo bar":
                            v14 = v14
                        default:
                            var v22 uint8 = 128
                            _ = v22
                        }
                        v14 = v14
                        _ = fi19
                        _ = v20
                    }
                }
                var v23 uint16 = v14
                _ = fi17
                _ = v23
            }
        }
    }
    fmt.Println(v3, v1)
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(fi24, v3)
        _ = fi24
    }
    if len(v5) > 0 {
        v5[0] = v4
    }
    if (false || ((float64(1e10) + 0.0) == -2.25)) {
        v3 = v3
        for fi25 := 0; fi25 < 6; fi25++ {
            switch "a" {
            case "hello":
                var v26 string = "foo bar"
                _ = v26
            case "\n\t":
                v3 -= v3
            case "foo bar":
                v27 := GenS3{F0: (fi25 | (-(fi25))), F1: 42}
                _ = v27
            }
            _ = fi25
        }
    } else {
        v1 = 264
        v28 := GenS1{F0: uint16(v1)}
        _ = v28
    }
    v29 := GenS4{F0: -527, F1: "a", F2: 1}
    _ = len(v6)
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v1 -= (v1 + (v1 << (((v1 >> ((uint(v14)) & 63))) & 63)))
            var v32 int = 399
            v33 := GenS5{F0: 81}
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 uint16 = 788
        if len(v2) > 0 {
            v2[0] = fi30
        }
        _ = fi30
        _ = v34
    }
    if (v4 >= (v4 - 524)) {
        v5 = append(v5, v4)
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
