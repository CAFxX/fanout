package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 int32, p1 uint32) int64 {
    if (true || false) {
        var v6 float64 = 3.14159
        fmt.Println(p0, p1)
        p1 = (p1 * p1)
        _ = v6
    } else {
        var v7 int8 = 35
        fmt.Println(v7, v7)
        switch 4294967295 {
        case 9223372036854775807:
            v7 = -94
        default:
            v7 -= v7
        }
        _ = v7
    }
    if (32767 > -524) {
        v8 := []int{288, -542, 128, 4294967295}
        if true {
            v8 = append(v8, 577)
            if true {
                v9 := GenS1{F0: true}
                p1 = p1
                switch ("hello" + ("foo bar" + "")) {
                case "\n\t":
                    v10 := GenS1{F0: false}
                    _ = v10
                case "":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "foo bar":
                    for fi11 := 0; fi11 < 17; fi11++ {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            var v13 int = (fi12 - 42)
                            fmt.Println(p0, p1)
                            p0 = -479
                            _ = fi12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 17; fi14++ {
                            var v15 int64 = -9223372036854775808
                            _ = fi14
                            _ = v15
                        }
                        _ = fi11
                    }
                }
                _ = v9
            } else {
                if false {
                    _ = len(v8)
                    for fi16 := 0; fi16 < 15; fi16++ {
                        fmt.Println(p0, fi16)
                        _ = fi16
                    }
                    var v17 int64 = 49
                    _ = v17
                } else {
                    p0 = p0
                    var v18 int8 = 100
                    _ = v18
                }
            }
        } else {
            p1 += (p1 / (((p1 ^ (^(p1)))) | 1))
            if len(v8) > 0 {
                v8[0] = -2147483648
            }
        }
        fmt.Println(p0, p0)
        _ = v8
    }
    v19 := []int{-583, 962, 0, 2147483647}
    _ = p0
    _ = p1
    _ = v19
    return 104
}
func genFn1() int {
    v20 := GenS1{F0: (532 != 697)}
    if false {
        var v21 float32 = ((0.1 + (0.0 - 1e10)) / 1e10)
        v21 -= float32((-2.25 + ((1.5 * float64(float32(-2.25))) / float64(v21))))
        v22 := []int{-461, -1}
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: (("hello" != ("\n\t" + "\n\t")) && (!(true || false)))}
    _ = v20
    _ = v23
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint64 = 721
    v4 := make(map[uint32]string)
    var v5 uint8 = 1
    v24 := []int{-2147483648, 684}
    for fi25 := 0; fi25 < 10; fi25++ {
        if false {
            v3 = v3
            v26 := GenS2{F0: 1000, F1: 252}
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 int32 = -7
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        var v29 uint64 = (v3 + 42)
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(fi25, fi30)
            var v31 float32 = 0.0
            var v32 float32 = 1.5
            v2[""] = ((("foo bar" + (("a" + "\n\t") + "\n\t")) + "") + "")
            var v33 float64 = 0.0
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi25
        _ = v29
    }
    fmt.Println(v3, v5)
    var v34 int32 = 2147483647
    var v35 int32 = -168
    switch 85 {
    case -429:
        v3 += 2
    case 128:
        var v36 float32 = (float32(0.1) * 0.1)
        _ = v36
    case 32768:
        v37 := v1[645]
        _ = v37
    default:
        var v38 int32 = (v34 >> ((4294967295) & 63))
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        if ((float64(0.1) + (-2.25 - (float64(float32(0.1)) - (1e10 * 0.1)))) <= (-2.25 + float64(-2.25))) {
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
        v24 = append(v24, -9223372036854775808)
        v34 += v35
        _ = fi39
    }
    v35 = (v34 % ((v34) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
