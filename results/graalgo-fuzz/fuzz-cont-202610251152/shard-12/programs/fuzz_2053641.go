package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 uint8, p1 int16) string {
    v4 := []int{}
    switch "" {
    case "":
        v5 := GenS1{F0: 42, F1: "foo bar", F2: 0}
        _ = v5
    case "foo bar":
        if (18446744073709551615 <= 18446744073709551615) {
            fmt.Println(p1, p1)
            var v6 int64 = 42
            var v7 int16 = (-(p1))
            _ = v6
            _ = v7
        }
    default:
        switch (("foo bar" + ("" + "\n\t")) + (("a" + "hello") + "a")) {
        case "a":
            switch -1 {
            case 4294967295:
                v8 := GenS1{F0: 0, F1: (("" + "\n\t") + "a"), F2: 2}
                _ = v8
            case 127:
                p1 = p1
            case 370:
                p0 += p0
            default:
                for fi9 := 0; fi9 < 12; fi9++ {
                    switch "a" {
                    case "\n\t":
                        v10 := GenS1{F0: 51, F1: "hello", F2: 2}
                        _ = v10
                    case "a":
                        var v11 float64 = 3.14159
                        _ = v11
                    default:
                        p1 += (p1 - p1)
                    }
                    var v12 string = ("a" + (("a" + (("" + "hello") + ("\n\t" + "\n\t"))) + ("hello" + "foo bar")))
                    _ = fi9
                    _ = v12
                }
            }
        default:
            var v13 uint16 = 174
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return "\n\t"
}
func genFn1(p0 uint16) uint8 {
    v14 := []int{0, 127}
    v15 := GenS2{F0: -2147483648, F1: 295, F2: 4294967295}
    var v16 uint64 = 4294967295
    v14 = append(v14, 1)
    var v17 int32 = 894
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 183
    v3 := []uint32{65535, 627, 2147483647, 106}
    fmt.Println(v2)
    if false {
        if (!(32767 >= 42)) {
            for fi18 := 0; fi18 < 9; fi18++ {
                v2 = (v2 & v2)
                _ = fi18
            }
            var v19 float32 = (1e10 * 1.5)
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 uint8 = 0
                v2 = (v2 ^ v2)
                switch "\n\t" {
                case "\n\t":
                    v2 = 0
                case "":
                    fmt.Println(fi20)
                }
                _ = fi21
                _ = v22
            }
            _ = len(v3)
            _ = fi20
        }
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v2 -= (v2 | v2)
                switch ((("foo bar" + ("foo bar" + ("a" + "a"))) + (("hello" + "hello") + "foo bar")) + ("a" + "\n\t")) {
                case "a":
                    var v25 int32 = v2
                    _ = v25
                case "":
                    fmt.Println(fi24, v2, fi24)
                default:
                    v2 -= (v2 << ((0) & 63))
                }
                _ = fi24
            }
            _ = fi23
        }
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint16 = 248
            _ = fi26
            _ = v27
        }
        fmt.Println(v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int = 882
    fmt.Println(v2, v28)
    if true {
        var v29 int64 = 1
        if ((v28 + 32768) < (v28 >> ((42) & 63))) {
            var v30 uint16 = 42
            _ = len(v1)
            var v31 uint64 = 18446744073709551615
            _ = v30
            _ = v31
        } else {
            v32 := v1[2147483647]
            var v33 bool = false
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            _ = len(v1)
            _ = fi34
        }
        _ = v29
    }
    v1[345] = ("foo bar" + "\n\t")
    v2 -= 42
    var v35 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
}
