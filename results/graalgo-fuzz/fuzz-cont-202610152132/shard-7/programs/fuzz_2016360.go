package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 int8, p1 int) uint64 {
    var v7 int64 = -1
    p1 = p1
    var v8 int64 = v7
    v9 := GenS1{F0: 255, F1: (("a" + "") + "hello"), F2: p1}
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            p1 = 885
            _ = fi11
        }
        switch -763 {
        case 128:
            var v12 uint = 65535
            _ = v12
        case -970:
            var v13 int32 = 0
            _ = v13
        default:
            for fi14 := 0; fi14 < 3; fi14++ {
                if false {
                    if (549 > -2147483648) {
                        var v15 int64 = -632
                        if true {
                            var v16 int8 = p0
                            _ = v16
                        }
                        v15 = 9223372036854775807
                        _ = v15
                    } else {
                        v17 := GenS1{F0: 800, F1: "foo bar", F2: 0}
                        _ = v17
                    }
                    var v18 string = "\n\t"
                    _ = v18
                } else {
                    var v19 float64 = (1.5 + (1e10 * float64(3.14159)))
                    var v20 uint16 = 32767
                    if (((((-2.25 * v19) / 1.5) / v19) * v19) == 0.1) {
                        v21 := GenS1{F0: v20, F1: "foo bar", F2: fi14}
                        var v22 int = fi10
                        v19 = v19
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS1{F0: v20, F1: "hello", F2: 32767}
                    _ = v19
                    _ = v20
                    _ = v23
                }
                _ = fi14
            }
        }
        v24 := []int{}
        _ = fi10
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int) bool {
    p0 = p0
    var v25 int64 = 1
    _ = p0
    _ = v25
    return ((-1 < 127) && false)
}

func main() {
    v1 := []string{"\n\t", ("a" + ("\n\t" + (("a" + (("foo bar" + "foo bar") + "hello")) + "hello")))}
    var v2 bool = true
    var v3 uint16 = 255
    var v4 int8 = -128
    v5 := make(map[uint32]int)
    v6 := make(map[string]bool)
    var v26 float32 = 3.14159
    v1 = append(v1, ("a" + "foo bar"))
    for fi27 := 0; fi27 < 4; fi27++ {
        v2 = v2
        if (false || (128 < 36)) {
            v3 += v3
            var v28 uint8 = 0
            var v29 bool = v2
            _ = v28
            _ = v29
        } else {
            v26 = 0.1
        }
        var v30 int16 = -1
        _ = fi27
        _ = v30
    }
    v3 += (^(v3))
    if true {
        _ = len(v1)
        v31 := GenS1{F0: (v3 + (v3 ^ (v3 % ((v3) | 1)))), F1: (((("hello" + "a") + (("foo bar" + "foo bar") + "")) + "\n\t") + ("" + "")), F2: 811}
        _ = v31
    }
    v26 = ((float32(0.0) + ((1.5 * (v26 * (0.0 - -2.25))) * (v26 + v26))) - 0.1)
    for fi32 := 0; fi32 < 15; fi32++ {
        _ = len(v1)
        _ = fi32
    }
    var v33 uint16 = (v3 | 0)
    v2 = (v2 || true)
    if len(v1) > 0 {
        v1[0] = ("" + "")
    }
    switch 0 {
    case 4294967295:
        switch 4294967295 {
        case -2147483648:
            v4 = (v4 + (v4 | v4))
        }
    case 42:
        v34 := GenS2{F0: 807}
        _ = v34
    default:
        v35 := GenS2{F0: 4294967295}
        _ = v35
    }
    if v2 {
        var v36 bool = v2
        fmt.Println(v4, v26)
        var v37 float64 = -2.25
        _ = v36
        _ = v37
    }
    var v38 float64 = (1.5 / 1.0)
    v39 := GenS3{F0: 739, F1: 18446744073709551615}
    v33 += (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
