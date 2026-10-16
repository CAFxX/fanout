package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 bool }
func genFn0(p0 bool, p1 int64, p2 uint) uint8 {
    fmt.Println(p2, p2, p2)
    switch "hello" {
    case "foo bar":
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: 1, F1: 100, F2: 0}
                p2 = p2
                if p0 {
                    var v10 uint32 = 2147483647
                    p1 = (p1 << ((2) & 63))
                    _ = v10
                } else {
                    v11 := []int{42, (^(fi7)), (fi8 - int(p1)), fi7}
                    if p0 {
                        var v12 uint32 = 42
                        v13 := GenS1{F0: 2, F1: -128, F2: -1}
                        var v14 uint64 = 633
                        fmt.Println(v14, v12, p1)
                        var v15 int32 = -2147483648
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 float32 = 0.1
                    _ = v11
                    _ = v16
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "\n\t":
        p1 = p1
    case "hello":
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 227
}
func genFn1(p0 int) bool {
    if (0.1 < -2.25) {
        p0 += p0
    }
    var v17 int32 = 0
    v18 := GenS1{F0: 127, F1: 127, F2: 42}
    v19 := []int{}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return true
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"", ("" + "a")}
    v3 := []string{"\n\t", "\n\t", "\n\t", "\n\t"}
    var v4 uint16 = 1
    v5 := []uint32{65535}
    var v6 int8 = 42
    v4 = 1
    v4 -= v4
    var v20 int16 = -790
    v21 := GenS2{F0: 18446744073709551615, F1: 0, F2: false}
    v1[1] = (((("hello" + ("" + ("\n\t" + ("hello" + "a")))) + "") + ("" + "\n\t")) + "")
    switch 128 {
    case 0:
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 int8 = -1
            _ = fi22
            _ = v23
        }
    case 4294967295:
        v1[-116] = ("" + "foo bar")
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v24 uint64 = 1
    if false {
        v4 = (v4 << ((1) & 63))
        if (163 > 1) {
            delete(v1, -9223372036854775808)
            var v25 int = -213
            _ = v25
        } else {
            v26 := GenS2{F0: 18446744073709551615, F1: -387, F2: ((v20 >> ((152) & 63)) >= -32768)}
            _ = v26
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            switch fi27 {
            case 32767:
                v24 -= 2
            case -124:
                var v28 uint64 = v24
                _ = v28
            default:
                _ = len(v1)
            }
            v4 = 42
            _ = fi27
        }
    } else {
        var v29 float64 = ((((((-2.25 + 0.0) * 0.1) / 1e10) + 3.14159) + ((1e10 / 1.0) / (0.1 / 1e10))) / 0.1)
        if true {
            var v30 float64 = 0.1
            v31 := GenS2{F0: 42, F1: 1, F2: true}
            var v32 int = int(v20)
            _ = v30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v33 int8 = v6
        _ = v33
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            _ = len(v1)
            fmt.Println(v4, fi34, v20)
            v20 -= v20
            _ = fi34
        }
    }
    v35 := GenS1{F0: 255, F1: v6, F2: (v20 & (v20 - (v20 - (v20 - (v20 & (^(v20)))))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
