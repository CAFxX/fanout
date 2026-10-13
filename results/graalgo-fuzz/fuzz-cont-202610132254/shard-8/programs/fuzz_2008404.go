package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
func genFn0(p0 uint32) uint16 {
    p0 = p0
    fmt.Println(p0)
    v5 := []int{0, 32768}
    p0 += p0
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int, p1 bool, p2 bool) uint {
    var v6 int64 = -1
    switch "" {
    case "foo bar":
        if false {
            var v7 uint = 42
            _ = v7
        }
    case "":
        fmt.Println(v6)
    default:
        var v8 int16 = 32767
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint = 286
    v3 := []uint32{4294967295, 1}
    v4 := make(map[string]int)
    v2 -= v2
    fmt.Println(v1, v2)
    var v9 int64 = -1
    if (false && false) {
        for fi10 := 0; fi10 < 19; fi10++ {
            v9 += -9223372036854775808
            _ = fi10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            v4[((("hello" + ("foo bar" + "")) + "") + ("\n\t" + "a"))] = (fi11 & fi11)
            switch fi11 {
            case 2147483647:
                if (!(246 <= 4294967295)) {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        delete(v4, "hello")
                        if len(v3) > 0 {
                            v3[0] = 53
                        }
                        var v13 uint32 = 847
                        _ = fi12
                        _ = v13
                    }
                } else {
                    v14 := GenS1{F0: fi11, F1: 389, F2: (v1 - v1)}
                    v1 += -225
                    _ = v14
                }
            case -462:
                var v15 uint32 = 65535
                _ = v15
            case 9223372036854775807:
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS1{F0: (fi11 & -2147483648), F1: -32768, F2: fi11}
                    _ = len(v4)
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi11
        }
        if false {
            if (("foo bar" + "a") >= "") {
                v1 -= -2147483648
            }
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int8 = 127
        _ = fi18
        _ = v19
    }
    v2 += (v2 * (^(v2)))
    for fi20 := 0; fi20 < 14; fi20++ {
        v2 = genFn1((v1 << ((v2) & 63)), true, false)
        var v21 uint16 = uint16(v2)
        fmt.Println(fi20, v9, v1)
        _ = fi20
        _ = v21
    }
    var v22 string = ((("hello" + "hello") + ("\n\t" + "")) + "foo bar")
    if ((100 == 0) && true) {
        var v23 int64 = v9
        v23 = v23
        _ = v23
    } else {
        v22 = (v22 + "\n\t")
    }
    v1 = (v1 * 0)
    v24 := GenS1{F0: v1, F1: 32767, F2: v1}
    fmt.Println(v9, v22, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
