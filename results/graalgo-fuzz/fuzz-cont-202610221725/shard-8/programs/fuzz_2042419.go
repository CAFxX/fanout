package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0(p0 int, p1 uint64, p2 bool) bool {
    fmt.Println(p1)
    v5 := GenS1{F0: 32767, F1: (!(-1 != 0)), F2: uint(p1)}
    v6 := GenS1{F0: 478, F1: p2, F2: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 uint16) int16 {
    if (!false) {
        v7 := GenS2{F0: 9223372036854775807, F1: 1, F2: 18446744073709551615}
        _ = v7
    }
    p0 += p0
    var v8 string = ("" + "foo bar")
    v9 := GenS1{F0: p0, F1: ((255 > 230) && (!true)), F2: 0}
    _ = p0
    _ = v8
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint8 = 2
    v4 := []string{}
    v4 = append(v4, "a")
    fmt.Println(v3)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (fi10 | fi10) {
        case 42:
            v2[(fi10 | fi10)] = (fi10 % ((fi10) | 1))
        case 9223372036854775807:
            var v11 int = 128
            _ = v11
        default:
            switch "a" {
            case "a":
                _ = len(v4)
            }
        }
        _ = fi10
    }
    if false {
        v12 := GenS2{F0: 9223372036854775807, F1: -9223372036854775808, F2: 18446744073709551615}
        fmt.Println(v3)
        _ = v12
    }
    fmt.Println(v3)
    v13 := GenS2{F0: 1, F1: -705, F2: 42}
    v3 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint32 = 4294967295
    for fi15 := 0; fi15 < 13; fi15++ {
        v1[fi15] = true
        v4 = append(v4, "")
        _ = fi15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v3 += v3
        var v17 uint8 = (v3 * v3)
        _ = fi16
        _ = v17
    }
    if (-1 != 2147483647) {
        v14 += (v14 | 2)
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int32 = 2147483647
            v21 := GenS3{F0: v3, F1: -764, F2: v3}
            if (fi18 > -713) {
                var v22 uint8 = v3
                fmt.Println(v3, v22, v22)
                var v23 uint64 = 882
                v14 += 1
                v20 = 840
                _ = v22
                _ = v23
            }
            var v24 bool = (!(-520 > 0))
            _ = fi19
            _ = v20
            _ = v21
            _ = v24
        }
        if false {
            v14 = 2
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        }
        _ = fi18
    }
    var v25 bool = false
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 string = ((("foo bar" + ("foo bar" + "hello")) + "foo bar") + (("\n\t" + ("hello" + ("foo bar" + ("\n\t" + "foo bar")))) + ("foo bar" + "\n\t")))
        var v28 uint8 = v3
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 int32 = 42
            if ((fi29 << ((0) & 63)) >= fi29) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v31 int16 = genFn1(1)
                fmt.Println(v30)
                _ = v31
            }
            _ = fi29
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v4 = append(v4, "")
    v25 = ((("\n\t" + ("\n\t" + ((("" + ("" + "a")) + "a") + "hello"))) + ("foo bar" + "foo bar")) < ("hello" + (("hello" + "\n\t") + (("hello" + "foo bar") + "a"))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
