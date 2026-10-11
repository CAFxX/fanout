package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 int32 }
type GenS5 struct { F0 string; F1 uint64; F2 string }
func genFn0(p0 int, p1 string) uint32 {
    if (648 != 575) {
        v4 := []int{}
        if true {
            switch (p0 % (((p0 << ((18446744073709551615) & 63))) | 1)) {
            case 32768:
                var v5 uint32 = 2147483647
                _ = v5
            case 9223372036854775807:
                switch p0 {
                case 1:
                    for fi6 := 0; fi6 < 6; fi6++ {
                        v7 := GenS1{F0: p0}
                        _ = fi6
                        _ = v7
                    }
                case 32767:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        } else {
            v4 = append(v4, p0)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v4
    } else {
        p0 = (-(p0))
    }
    p0 += (p0 ^ p0)
    v8 := GenS2{F0: 2147483647, F1: p0}
    if ((18446744073709551615 > 0) && true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                p1 = p1
                _ = fi10
            }
            p1 = ("" + "")
            var v11 float64 = 1e10
            _ = fi9
            _ = v11
        }
        p1 = (p1 + (p1 + (p1 + p1)))
        if true {
            if false {
                switch (p0 >> ((4294967295) & 63)) {
                case 32768:
                    p1 = p1
                default:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        fmt.Println(p0)
                        v13 := []int{(p0 % ((32767) | 1)), (^(fi12))}
                        v14 := GenS2{F0: -9223372036854775808, F1: p0}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                p0 = -24
                var v15 int = 1
                _ = v15
            }
            var v16 uint32 = 42
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS3{F0: 65535}
            _ = fi17
            _ = v18
        }
        v19 := GenS4{F0: 42, F1: 42, F2: -2147483648}
        if (((!false) || false) || (!false)) {
            for fi20 := 0; fi20 < 2; fi20++ {
                p0 += 824
                switch ("\n\t" + (p1 + "foo bar")) {
                case "hello":
                    if ((!false) || false) {
                        p0 = (fi20 - p0)
                        if false {
                            var v21 uint64 = 0
                            p0 -= fi20
                            fmt.Println(p0, v21, v21)
                            _ = v21
                        }
                    }
                case "a":
                    p0 = p0
                default:
                    var v22 int64 = -9223372036854775808
                    _ = v22
                }
                switch (("foo bar" + "a") + (("" + p1) + p1)) {
                case "":
                    if (620 >= 0) {
                        var v23 int16 = -1
                        _ = v23
                    } else {
                        fmt.Println(p1, p0, p1)
                        fmt.Println(fi20)
                        var v24 uint16 = 42
                        _ = v24
                    }
                }
                _ = fi20
            }
            p1 = p1
        } else {
            v25 := []int{p0, 32767, 32767, (p0 ^ (p0 << ((65535) & 63)))}
            p1 = "foo bar"
            p1 = (p1 + p1)
            _ = v25
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v8
    return 65535
}
func genFn1() int16 {
    if (false || true) {
        var v26 uint16 = 126
        _ = v26
    } else {
        v27 := []int{278, 9223372036854775807}
        _ = v27
    }
    var v28 int64 = 9223372036854775807
    _ = v28
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := GenS2{F0: 32768, F1: fi30}
            v2[((("hello" + "\n\t") + "a") + "")] = "a"
            fmt.Println(fi29)
            _ = fi30
            _ = v31
        }
        v32 := v3[-263]
        _ = fi29
        _ = v32
    }
    var v33 uint = 18446744073709551615
    if true {
        fmt.Println(v33)
    }
    var v34 float32 = 0.0
    v34 = v34
    switch "a" {
    case "a":
        for fi35 := 0; fi35 < 12; fi35++ {
            if (!(v34 == float32((0.0 * float64(float32(1.5)))))) {
                if (int16(fi35) != genFn1()) {
                    fmt.Println(v34)
                } else {
                    fmt.Println(v34, fi35, v33)
                }
            }
            v36 := GenS4{F0: 1, F1: 255, F2: -1}
            delete(v1, 2147483647)
            _ = fi35
            _ = v36
        }
    case "foo bar":
        var v37 uint64 = 4294967295
        _ = v37
    case "":
        switch -38 {
        case 4294967295:
            var v38 int16 = 753
            _ = v38
        case 32767:
            _ = len(v2)
        }
    default:
        if true {
            v33 += (v33 & 229)
        } else {
            v34 -= (v34 - v34)
            var v39 uint8 = 128
            _ = v39
        }
    }
    v34 = (v34 + float32(-2.25))
    var v40 bool = true
    if (1 < 0) {
        if (32767 <= 32767) {
            v40 = v40
        }
    }
    v41 := GenS5{F0: "\n\t", F1: 4294967295, F2: "foo bar"}
    v42 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v40
    _ = v41
    _ = v42
}
