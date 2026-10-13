package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int32; F1 uint; F2 int64 }
func genFn0() int64 {
    v6 := []int{-2147483648, 9223372036854775807, 636}
    v7 := GenS1{F0: 0, F1: 42, F2: 32768}
    var v8 uint32 = 1
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint16) uint {
    fmt.Println(p0)
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(p0, p0)
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    var v11 string = (("foo bar" + "a") + "foo bar")
    p0 += p0
    v12 := GenS2{F0: 596, F1: 2, F2: 2147483647}
    _ = p0
    _ = v11
    _ = v12
    return 597
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := []string{"\n\t", ("\n\t" + "\n\t"), ("a" + "foo bar")}
    v4 := make(map[uint32]int)
    var v5 float32 = float32(((float64((-2.25 + (float32(0.1) + -2.25))) + (3.14159 * 0.1)) + 1.5))
    v13 := GenS1{F0: -46, F1: 32767, F2: 876}
    var v14 string = ""
    var v15 float64 = (((3.14159 * 1.5) / (0.1 / 1e10)) / 1e10)
    fmt.Println(v15)
    v14 = ("hello" + "\n\t")
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS2{F0: 65535, F1: 427, F2: (^(fi16))}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v5 -= (v5 * (v5 - v5))
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint32 = 65535
            _ = fi19
            _ = v20
        }
        if (v14 != v14) {
            var v21 int = (fi18 & (fi18 / ((fi18) | 1)))
            var v22 int64 = genFn0()
            v23 := GenS1{F0: 127, F1: 32767, F2: fi18}
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    v5 -= v5
    v4[0] = -1
    v5 -= float32(v15)
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := GenS3{F0: 42, F1: genFn0()}
        var v26 string = v14
        _ = fi24
        _ = v25
        _ = v26
    }
    v15 = 0.0
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            if len(v2) > 0 {
                v2[0] = (((((v14 + "a") + "a") + v14) + ((("a" + v14) + ("foo bar" + "")) + (v14 + v14))) + v14)
            }
            v29 := GenS2{F0: 487, F1: 789, F2: (fi27 % ((fi27) | 1))}
            v15 -= v15
            var v30 uint8 = 180
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 int = (fi27 + fi27)
        v32 := v4[583]
        _ = fi27
        _ = v31
        _ = v32
    }
    v33 := v1[""]
    if len(v2) > 0 {
        v2[0] = v14
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            fmt.Println(v14, fi34)
            v36 := GenS4{F0: -1, F1: 4294967295, F2: -9223372036854775808}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v15 = v15
    switch "\n\t" {
    case "hello":
        for fi37 := 0; fi37 < 1; fi37++ {
            v38 := GenS2{F0: 65535, F1: 2, F2: (^(fi37))}
            v39 := GenS3{F0: 0, F1: 0}
            v40 := GenS1{F0: 42, F1: 42, F2: (fi37 + fi37)}
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    case "":
        if (false || (255 < genFn1(2))) {
            if false {
                if (2 > 18446744073709551615) {
                    v5 -= ((0.1 / -2.25) + (v5 / (float32((1e10 + float64((v5 * float32(v15))))) + 1.5)))
                    switch 2147483647 {
                    case 0:
                        if (!((v14 + "hello") < v33)) {
                            var v41 uint = 4294967295
                            _ = v41
                        }
                    }
                } else {
                    v42 := GenS3{F0: genFn0(), F1: 9223372036854775807}
                    var v43 bool = (-78 < 833)
                    _ = v42
                    _ = v43
                }
                var v44 int64 = -9223372036854775808
                fmt.Println(v15, v44)
                _ = v44
            }
        }
    case "a":
        v45 := GenS1{F0: -128, F1: 245, F2: 9223372036854775807}
        _ = v45
    default:
        var v46 int8 = -1
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
