package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
func genFn0(p0 string, p1 int16) uint64 {
    var v6 uint = 4294967295
    v7 := []int{-914, -393, -2147483648, -2147483648}
    fmt.Println(v6, p1, v6)
    _ = len(v7)
    var v8 uint16 = 65535
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 uint64) uint {
    p1 -= p1
    p1 += p1
    p0 += p0
    v9 := GenS1{F0: ((("\n\t" + ("\n\t" + "foo bar")) + ("" + "foo bar")) + "hello"), F1: -2147483648, F2: p0}
    var v10 uint = uint(p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v10
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!(!("a" <= (("foo bar" + "foo bar") + "hello"))))
    var v3 float32 = (0.1 - (((0.1 * (1e10 / 0.1)) / (((1.5 - 0.0) / 1e10) + 0.0)) * (0.1 - (1.5 / 1.0))))
    v4 := make(map[string]bool)
    var v5 int16 = 42
    switch 9223372036854775807 {
    case 520:
        var v11 int8 = 0
        _ = v11
    case 1:
        for fi12 := 0; fi12 < 18; fi12++ {
            if (v2 && (true || (!v2))) {
                v13 := GenS1{F0: ("hello" + "hello"), F1: 2147483647, F2: 0}
                _ = v13
            }
            _ = fi12
        }
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: "", F1: 1, F2: 2147483647}
            _ = fi14
            _ = v15
        }
    }
    v2 = (!((v5 ^ -32768) == v5))
    if v2 {
        v16 := GenS2{F0: 2, F1: 4294967295, F2: 454}
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint16 = 65535
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        if (18446744073709551615 >= 1) {
            switch 1 {
            case 42:
                var v19 int32 = -873
                _ = v19
            case -9223372036854775808:
                if v2 {
                    delete(v1, 65535)
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v5, v5, v2)
                    }
                    var v20 int16 = v5
                    _ = v20
                }
            default:
                v21 := []int{-446}
                _ = v21
            }
            v22 := GenS1{F0: "\n\t", F1: 0, F2: 875}
            _ = v22
        }
        var v23 uint16 = 65535
        v5 -= (v5 * (v5 | 876))
        _ = v23
    }
    var v24 uint8 = 2
    v2 = v2
    v24 = (v24 - (v24 + (^(v24))))
    v25 := GenS1{F0: "", F1: -777, F2: 65535}
    fmt.Println(v5)
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 uint64 = uint64(v24)
        v2 = (((fi26 >> ((803) & 63)) <= fi26) && (v2 || true))
        _ = fi26
        _ = v27
    }
    v5 = (-(v5))
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 uint64 = genFn0("hello", -32768)
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 int16 = (^(v5))
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    var v32 int8 = 39
    if ((-(v32)) != v32) {
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS1{F0: (((("a" + ("" + "")) + "") + (("hello" + ("" + "foo bar")) + "hello")) + ("\n\t" + "")), F1: -851, F2: 4294967295}
            fmt.Println(v3, v2, v24)
            for fi35 := 0; fi35 < 4; fi35++ {
                if (true && false) {
                    var v36 int32 = -2147483648
                    v32 = v32
                    var v37 uint8 = v24
                    _ = v36
                    _ = v37
                }
                switch "a" {
                case "a":
                    v4[("\n\t" + ("hello" + "a"))] = (v2 && (v24 == v24))
                    v38 := []int{9223372036854775807}
                    _ = len(v38)
                    _ = v38
                default:
                    var v39 int64 = 9223372036854775807
                    _ = v39
                }
                delete(v4, "")
                _ = fi35
            }
            _ = fi33
            _ = v34
        }
        if v2 {
            fmt.Println(v32, v5, v2)
            v3 += v3
            v5 = (v5 ^ (v5 & v5))
        }
        if v2 {
            v5 -= 1000
            var v40 uint = 4294967295
            fmt.Println(v32, v2, v24)
            _ = v40
        }
    }
    var v41 string = "hello"
    v42 := GenS2{F0: 494, F1: 9223372036854775807, F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
