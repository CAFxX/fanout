package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
func genFn0(p0 string, p1 int, p2 int64) uint {
    v7 := GenS1{F0: 4294967295}
    var v8 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 706
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 1
    v3 := []int{}
    v4 := []uint32{492, 2147483648, 4294967295}
    var v5 int = 111
    var v6 uint = (v2 & (v2 + v2))
    if ((!(false && false)) || true) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (65535 < (v1 - (v1 ^ 0))) {
            v5 = (v5 & (v5 << ((genFn0("a", 1, -255)) & 63)))
            for fi9 := 0; fi9 < 9; fi9++ {
                v5 += (v5 % ((v5) | 1))
                _ = fi9
            }
        }
    }
    v2 -= (v6 % (((^(v2))) | 1))
    var v10 int16 = 1
    if len(v3) > 0 {
        v3[0] = 165
    }
    _ = len(v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v5 {
    case 32767:
        var v11 int16 = (v10 * (v10 & -1))
        _ = v11
    case 128:
        if (((v6 >> (((v2 % ((genFn0(("foo bar" + (("hello" + "") + "hello")), 32768, int64(v5))) | 1))) & 63)) > 0) && (!(!(!true)))) {
            v5 += v5
            v12 := GenS1{F0: uint(v1)}
            _ = v12
        }
    default:
        fmt.Println(v2)
    }
    switch (v5 / ((0) | 1)) {
    case 0:
        v13 := GenS1{F0: (^(v6))}
        _ = v13
    case -915:
        var v14 uint64 = 2
        _ = v14
    default:
        v5 -= (^(v5))
    }
    var v15 uint16 = (v1 << ((genFn0((("foo bar" + "") + "hello"), v5, 1)) & 63))
    if len(v3) > 0 {
        v3[0] = 127
    }
    v6 = genFn0(("foo bar" + ""), (v5 % ((v5) | 1)), 9223372036854775807)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                v2 += 691
                var v18 int64 = 42
                _ = v18
            }
            if (245 < 42) {
                _ = len(v3)
                fmt.Println(v10, v15)
                v19 := GenS2{F0: 9223372036854775807, F1: 0, F2: true}
                _ = v19
            } else {
                v5 = (fi16 + 1)
                v20 := GenS1{F0: genFn0("hello", (fi17 / ((fi16) | 1)), 9223372036854775807)}
                _ = v20
            }
            _ = fi17
        }
        if len(v3) > 0 {
            v3[0] = fi16
        }
        fmt.Println(fi16)
        _ = fi16
    }
    v10 -= (-(v10))
    v10 += v10
    switch "a" {
    case "":
        for fi21 := 0; fi21 < 7; fi21++ {
            switch "hello" {
            case "foo bar":
                v10 = int16(v2)
            case "\n\t":
                fmt.Println(v15, v15, v5)
            case "a":
                switch ("foo bar" + "hello") {
                case "a":
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 int64 = -9223372036854775808
                        v23 = v23
                        _ = fi22
                        _ = v23
                    }
                case "foo bar":
                    var v24 int32 = 0
                    _ = v24
                case "\n\t":
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 int64 = -9223372036854775808
                        if len(v3) > 0 {
                            v3[0] = -2147483648
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            }
            fmt.Println(v1)
            if (0 < 42) {
                v5 = fi21
                fmt.Println(v10)
                if false {
                    v2 = v6
                    fmt.Println(v5)
                }
            } else {
                fmt.Println(v6)
                v3 = append(v3, (fi21 * fi21))
            }
            _ = fi21
        }
    case "foo bar":
        v6 = v2
    case "hello":
        fmt.Println(v6, v15, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
