package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 string) int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p0, p1)
        if (-1 > -9223372036854775808) {
            if true {
                var v7 int32 = 0
                _ = v7
            } else {
                v8 := GenS1{F0: ((1 < 1) && false), F1: 42, F2: 2147483647}
                if false {
                    v9 := []int{(fi6 >> ((142) & 63)), (^(fi6)), 42}
                    _ = v9
                }
                p1 = p1
                _ = v8
            }
            v10 := []int{9223372036854775807, fi6}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v10
        }
        p0 = p0
        _ = fi6
    }
    var v11 string = p1
    for fi12 := 0; fi12 < 10; fi12++ {
        v11 = p1
        _ = fi12
    }
    p0 = (p0 % (((p0 % ((18446744073709551615) | 1))) | 1))
    if (32767 != 427) {
        v13 := []int{32768, -1}
        _ = v13
    } else {
        v14 := GenS1{F0: false, F1: p0, F2: 4294967295}
        var v15 uint8 = 154
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    return 32767
}
func genFn1(p0 string, p1 uint32) bool {
    p0 = "a"
    fmt.Println(p1)
    _ = p0
    _ = p1
    return (!((!(42 < p1)) || (((!(true || true)) || ((-895 < -9223372036854775808) || false)) && false)))
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := []string{((("a" + ("foo bar" + ("foo bar" + ""))) + "foo bar") + ("hello" + "hello"))}
    var v5 bool = true
    delete(v2, 2147483648)
    if (2147483648 < 42) {
        v5 = (32768 == -1)
        if false {
            v16 := GenS1{F0: (!false), F1: 42, F2: 2}
            v5 = v5
            for fi17 := 0; fi17 < 20; fi17++ {
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 int32 = -202
                    v3[42] = fi17
                    _ = fi18
                    _ = v19
                }
                switch fi17 {
                case 91:
                    var v20 float64 = (1.5 - 1e10)
                    _ = v20
                case 32767:
                    var v21 uint32 = 500
                    _ = v21
                case 119:
                    v5 = v5
                default:
                    v5 = v5
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int16 = 0
                    v23 -= genFn0(18446744073709551615, "\n\t")
                    if false {
                        var v24 int8 = -6
                        var v25 float64 = ((0.1 * (3.14159 - float64(1.5))) - (1.5 - 3.14159))
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
                _ = fi17
            }
            _ = v16
        }
    } else {
        fmt.Println(v5)
        switch ("hello" + "foo bar") {
        case "a":
            fmt.Println(v5)
        case "\n\t":
            v26 := GenS2{F0: -971}
            _ = v26
        default:
            v5 = (!(-2.25 != 1.5))
        }
    }
    switch "hello" {
    case "\n\t":
        v5 = (2 < 2147483648)
    case "hello":
        switch 648 {
        case 9223372036854775807:
            v2[396] = v5
        case -9223372036854775808:
            v5 = (-128 >= 0)
        }
    case "foo bar":
        fmt.Println(v5)
    default:
        var v27 float32 = (1e10 + -2.25)
        _ = v27
    }
    v5 = (-32768 <= -1)
    v5 = v5
    var v28 bool = true
    for fi29 := 0; fi29 < 4; fi29++ {
        v28 = true
        v28 = (!true)
        _ = fi29
    }
    fmt.Println(v28, v5)
    for fi30 := 0; fi30 < 19; fi30++ {
        v28 = v5
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 float64 = 0.1
            var v33 uint8 = 2
            var v34 float64 = 1.5
            v33 += v33
            var v35 string = ("a" + "foo bar")
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi30
    }
    v36 := GenS2{F0: -32768}
    v28 = (!((!(-1 != 0)) && (v28 && v5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
