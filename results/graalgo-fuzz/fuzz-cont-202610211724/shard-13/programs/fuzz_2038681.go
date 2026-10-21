package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0() uint8 {
    v6 := []int{}
    var v7 int16 = 1000
    _ = v6
    _ = v7
    return 105
}
func genFn1(p0 string, p1 uint8, p2 int64) uint16 {
    p0 = "a"
    v8 := []int{4294967295, 2147483647, 2147483647}
    for fi9 := 0; fi9 < 8; fi9++ {
        switch "a" {
        case "":
            p1 = (p1 | (p1 * p1))
        case "a":
            var v10 bool = false
            _ = v10
        default:
            if (!(!true)) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    p2 = (p2 + -9223372036854775808)
                    v8 = append(v8, (^(fi9)))
                    fmt.Println(fi9, p1, p0)
                    _ = fi11
                }
            } else {
                v12 := GenS1{F0: p2}
                v13 := GenS1{F0: 9223372036854775807}
                _ = v12
                _ = v13
            }
        }
        fmt.Println(fi9, p1, p0)
        _ = fi9
    }
    p1 = (p1 ^ 127)
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 414
    var v3 uint32 = 50
    var v4 bool = true
    var v5 uint32 = (^(v3))
    v14 := []int{-942}
    if false {
        var v15 uint16 = (v2 << ((18446744073709551615) & 63))
        _ = v15
    }
    var v16 int32 = -126
    fmt.Println(v3, v16, v5)
    v5 = 65535
    switch "foo bar" {
    case "hello":
        switch "foo bar" {
        case "hello":
            if len(v14) > 0 {
                v14[0] = 115
            }
        }
    case "a":
        for fi17 := 0; fi17 < 11; fi17++ {
            v4 = (-2.25 > 1e10)
            _ = fi17
        }
    default:
        var v18 int64 = -952
        _ = v18
    }
    fmt.Println(v5)
    fmt.Println(v1, v2, v4)
    v2 = (v2 * 32767)
    v14 = append(v14, -9223372036854775808)
    var v19 uint64 = 2
    var v20 uint64 = (^(v19))
    if v4 {
        v21 := GenS1{F0: 9223372036854775807}
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := GenS2{F0: v1, F1: 0}
            switch "a" {
            case "foo bar":
                v3 += 1
            case "":
                var v24 uint16 = genFn1("", 42, 9223372036854775807)
                _ = v24
            default:
                var v25 uint32 = v3
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v26 uint64 = v20
        _ = v26
    }
    if len(v14) > 0 {
        v14[0] = 678
    }
    if (18 == 127) {
        v4 = (!(true || (!(v4 || (genFn0() < 42)))))
        var v27 int = -1
        _ = v27
    } else {
        for fi28 := 0; fi28 < 16; fi28++ {
            v4 = (!v4)
            if (!v4) {
                v1 += -2
                v14 = append(v14, -587)
            }
            var v29 int = (^(fi28))
            _ = fi28
            _ = v29
        }
        if v4 {
            var v30 uint16 = (v2 + uint16(v20))
            v31 := GenS2{F0: v1, F1: 9223372036854775807}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            v3 = (v5 ^ v3)
            for fi33 := 0; fi33 < 1; fi33++ {
                v1 = (v1 | (^(v1)))
                _ = fi33
            }
            var v34 uint64 = v20
            _ = fi32
            _ = v34
        }
    }
    if (genFn1(("" + "a"), 42, -754) > genFn1("", 127, -909)) {
        if true {
            for fi35 := 0; fi35 < 12; fi35++ {
                _ = len(v14)
                _ = fi35
            }
            switch "foo bar" {
            case "a":
                if v4 {
                    if false {
                        v2 = v2
                    } else {
                        var v36 float32 = (1e10 + 1e10)
                        _ = v36
                    }
                }
            case "foo bar":
                v1 = (-(v1))
            case "hello":
                if true {
                    if true {
                        for fi37 := 0; fi37 < 8; fi37++ {
                            if ((^(v16)) == v16) {
                                var v38 uint16 = 2
                                switch ((("hello" + ("a" + ("a" + "\n\t"))) + ("a" + "\n\t")) + "hello") {
                                case "hello":
                                    fmt.Println(v2)
                                default:
                                    v39 := GenS1{F0: 0}
                                    _ = v39
                                }
                                v1 = 42
                                _ = v38
                            }
                            var v40 int64 = -993
                            v41 := GenS2{F0: (v1 | 0), F1: int64(v2)}
                            _ = fi37
                            _ = v40
                            _ = v41
                        }
                    }
                } else {
                    if (false || (int32(v19) == v16)) {
                        var v42 uint64 = v20
                        _ = v42
                    }
                    v14 = append(v14, 9223372036854775807)
                    if true {
                        var v43 uint32 = 1
                        var v44 int16 = (v1 ^ (v1 % ((v1) | 1)))
                        _ = v43
                        _ = v44
                    }
                }
            }
        }
        var v45 int8 = -128
        _ = v45
    }
    v19 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
