package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32) uint8 {
    p0 = (p0 ^ (p0 / (((p0 | 888)) | 1)))
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 -= (p0 % ((4294967295) | 1))
        _ = fi6
    }
    var v7 int64 = 9223372036854775807
    if (916 <= 2) {
        v8 := []int{-1}
        _ = v8
    } else {
        fmt.Println(v7, p0)
    }
    if true {
        p0 += p0
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint8 = 1
            v11 := []int{}
            v10 += (v10 % ((v10) | 1))
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        p0 = (p0 & p0)
        switch (("foo bar" + ("\n\t" + "hello")) + "") {
        case "\n\t":
            for fi12 := 0; fi12 < 15; fi12++ {
                p0 = (p0 | (p0 + (p0 - p0)))
                _ = fi12
            }
        case "a":
            for fi13 := 0; fi13 < 18; fi13++ {
                p0 = (^(p0))
                v7 -= (v7 * v7)
                _ = fi13
            }
        case "":
            switch "a" {
            case "":
                v7 = 0
            }
        }
    }
    _ = p0
    _ = v7
    return 128
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := []string{"\n\t"}
    var v5 int16 = (v1 % (((v1 % (((-(v1))) | 1))) | 1))
    v14 := GenS1{F0: genFn0(0), F1: 0}
    fmt.Println(v1)
    var v15 bool = true
    if false {
        switch "hello" {
        case "\n\t":
            v5 = v1
        case "hello":
            v16 := GenS2{F0: "foo bar", F1: 255}
            _ = v16
        case "a":
            v1 = v5
        default:
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 float64 = (1e10 + 0.0)
                v5 = (v1 | (v5 ^ -1))
                if (!(0 == -1)) {
                    v19 := GenS2{F0: (((("hello" + "") + ("foo bar" + "a")) + ("" + (("foo bar" + ("a" + ("\n\t" + ""))) + ""))) + ""), F1: 224}
                    v20 := GenS2{F0: "hello", F1: 492}
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v5 += (v5 | (^(v1)))
                        if (-2.25 <= (-2.25 / (float32(float64((float32(1.5) - 0.1))) - (float32(-2.25) / 1e10)))) {
                            var v22 float64 = v18
                            _ = v22
                        }
                        if (v5 == (v1 % (((v1 % ((v5) | 1))) | 1))) {
                            var v23 bool = (!(true || v15))
                            v24 := GenS2{F0: "\n\t", F1: 255}
                            var v25 int8 = 127
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            var v26 int8 = int8(v5)
                            _ = len(v4)
                            v26 += 0
                            _ = v26
                        }
                        v1 = (v5 ^ 1)
                        v1 += v5
                        _ = fi21
                    }
                    _ = v19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
        }
        if len(v4) > 0 {
            v4[0] = "a"
        }
    } else {
        var v27 int8 = -34
        for fi28 := 0; fi28 < 7; fi28++ {
            v5 -= 42
            v15 = (!(v15 || v15))
            fmt.Println(v15, v15, v1)
            _ = fi28
        }
        v1 -= (^(v1))
        _ = v27
    }
    var v29 int32 = 2147483647
    for fi30 := 0; fi30 < 8; fi30++ {
        if (100 <= -1) {
            fmt.Println(v5, v29, fi30)
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 string = "foo bar"
            if v15 {
                if v15 {
                    var v33 uint16 = 0
                    _ = v33
                } else {
                    v34 := GenS1{F0: 128, F1: 127}
                    v29 = v29
                    _ = v34
                }
            } else {
                v15 = (321 >= 255)
                var v35 int16 = 1000
                var v36 float64 = 0.1
                _ = v35
                _ = v36
            }
            v32 = v32
            v5 += (v1 | v1)
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v37 := GenS3{F0: (v29 / ((v29) | 1))}
    v1 = v1
    fmt.Println(v15)
    delete(v2, 2147483647)
    var v38 float32 = 1.5
    if (42 < 2) {
        v1 -= v5
        if len(v4) > 0 {
            v4[0] = "a"
        }
        var v39 int16 = (v1 & 0)
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
