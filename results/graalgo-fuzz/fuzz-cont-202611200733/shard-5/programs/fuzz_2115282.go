package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 int32, p1 string) int32 {
    var v7 int = int(p0)
    var v8 bool = false
    var v9 uint16 = 68
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint = 1
    v2 := []string{"", "hello"}
    var v3 int32 = 736
    v4 := []int{0, 4294967295, 826, 32767}
    var v5 bool = false
    var v6 uint8 = 128
    v1 -= 1
    switch "\n\t" {
    case "foo bar":
        v1 = uint(v6)
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            switch "" {
            case "":
                fmt.Println(v3, v1)
            case "a":
                fmt.Println(fi10)
            default:
                switch (fi10 | (fi10 * fi10)) {
                case -251:
                    v11 := GenS1{F0: 1, F1: -32768, F2: 0}
                    _ = v11
                case 127:
                    var v12 uint64 = 18446744073709551615
                    _ = v12
                case -748:
                    if (("foo bar" >= ("\n\t" + (("a" + (("foo bar" + "foo bar") + "a")) + (("" + (("\n\t" + ("a" + "")) + ("\n\t" + ("\n\t" + ((("hello" + "a") + "hello") + "\n\t"))))) + (((((((("a" + "a") + "a") + "") + "\n\t") + "hello") + "\n\t") + "a") + ""))))) && (!v5)) {
                        var v13 string = ((("\n\t" + ("" + ("a" + "foo bar"))) + ("foo bar" + "foo bar")) + ("foo bar" + "a"))
                        v1 += v1
                        _ = v13
                    } else {
                        var v14 int32 = v3
                        if true {
                            switch "\n\t" {
                            case "":
                                v1 = v1
                            default:
                                v6 += v6
                            }
                            v5 = (!(!v5))
                        } else {
                            fmt.Println(v3, v1, v1)
                        }
                        _ = v14
                    }
                default:
                    v5 = (!true)
                }
            }
            _ = fi10
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            _ = len(v2)
            v5 = (((!(4294967295 > 885)) || v5) || (v1 >= v1))
            switch (fi15 >> ((v1) & 63)) {
            case 9223372036854775807:
                v3 = genFn0((v3 << ((v1) & 63)), "")
            }
            _ = fi16
        }
        v6 -= uint8(v1)
        _ = fi15
    }
    v17 := GenS2{F0: v1, F1: 1, F2: -99}
    if (9 != v6) {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: -32768, F1: -392, F2: v6}
            _ = fi18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS3{F0: (fi20 * (fi20 ^ -846)), F1: 0}
            _ = fi20
            _ = v21
        }
        if len(v4) > 0 {
            v4[0] = 127
        }
        v5 = (float64(1e10) != 1.5)
    }
    fmt.Println(v1)
    v1 = 4294967295
    v22 := GenS4{F0: (("" + (("foo bar" + "hello") + "a")) + "a"), F1: 42}
    if len(v2) > 0 {
        v2[0] = ("foo bar" + (("a" + ("a" + "a")) + ("hello" + "hello")))
    }
    v23 := GenS5{F0: 112}
    if (65535 != 42) {
        v3 = v3
    } else {
        v5 = (255 > 32767)
        v3 += genFn0(-1, "")
    }
    v5 = (18446744073709551615 >= 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
