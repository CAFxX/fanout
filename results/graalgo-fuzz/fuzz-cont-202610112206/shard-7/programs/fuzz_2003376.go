package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint16, p1 uint32, p2 uint) int16 {
    fmt.Println(p1, p0, p2)
    v4 := []int{-1, 0, 32768}
    for fi5 := 0; fi5 < 11; fi5++ {
        switch fi5 {
        case 987:
            v6 := GenS1{F0: 42, F1: (p2 % ((402) | 1)), F2: 4294967295}
            _ = v6
        case 32767:
            var v7 uint64 = 18446744073709551615
            _ = v7
        case 42:
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 string = (("foo bar" + "\n\t") + "a")
                _ = fi8
                _ = v9
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v10 uint64 = 0
        _ = fi5
        _ = v10
    }
    p2 -= p2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int8, p1 uint32, p2 int16) int8 {
    v11 := []int{-9223372036854775808, 42}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v12 := GenS2{F0: 42, F1: genFn0(2, (^(p1)), 2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (^(p0))
}

func main() {
    var v1 bool = false
    var v2 int64 = -594
    var v3 int16 = 1
    var v13 float64 = (-2.25 + 0.1)
    v2 = (-(v2))
    var v14 uint = 42
    var v15 uint16 = 42
    switch 32768 {
    case 32767:
        if (false && true) {
            v16 := GenS2{F0: -1, F1: v3}
            _ = v16
        } else {
            if (42 == -729) {
                v3 = genFn0(42, 0, v14)
                v17 := GenS2{F0: 63, F1: genFn0(v15, 2, (v14 % (((^(v14))) | 1)))}
                _ = v17
            }
            var v18 string = (("foo bar" + ("foo bar" + "\n\t")) + ((("a" + "a") + ("" + "\n\t")) + ("a" + ((("hello" + "hello") + "hello") + "foo bar"))))
            if (2147483647 != 42) {
                v2 = -15
                v13 = (v13 * v13)
                if true {
                    v19 := GenS2{F0: -20, F1: -978}
                    var v20 uint16 = v15
                    v14 = v14
                    _ = v19
                    _ = v20
                }
            } else {
                if (v1 || (1 >= 749)) {
                    v21 := GenS2{F0: 1, F1: v3}
                    v3 -= v3
                    if v1 {
                        v22 := GenS1{F0: 42, F1: 18446744073709551615, F2: (v14 >> ((v14) & 63))}
                        _ = v22
                    }
                    _ = v21
                } else {
                    v23 := GenS2{F0: -11, F1: (v3 >> ((v14) & 63))}
                    fmt.Println(v15)
                    _ = v23
                }
                v1 = (-1 > 2147483647)
            }
            _ = v18
        }
    case -9223372036854775808:
        var v24 string = ("a" + ("foo bar" + "a"))
        _ = v24
    case -487:
        for fi25 := 0; fi25 < 20; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint32 = 4294967295
                _ = fi26
                _ = v27
            }
            var v28 float32 = float32(v13)
            _ = fi25
            _ = v28
        }
    default:
        fmt.Println(v1)
    }
    v14 += v14
    v29 := []int{-2147483648}
    v15 = (^(v15))
    var v30 string = "foo bar"
    var v31 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
