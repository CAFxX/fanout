package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 uint8, p1 int16) uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch 1 {
            case 2147483647:
                var v9 int8 = 42
                _ = v9
            case -1:
                p0 += (^(p0))
                var v10 int8 = 0
                _ = v10
            case 42:
                var v11 int8 = 127
                _ = v11
            }
            _ = fi8
        }
        var v12 uint8 = 127
        _ = fi7
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(p1, p1, fi13)
        _ = fi13
    }
    if ((p0 * (p0 / (((p0 + (p0 + p0))) | 1))) < 128) {
        var v14 uint16 = 65535
        _ = v14
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint16 = 42
            var v17 uint32 = 2147483647
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    if false {
        v18 := []int{-865}
        _ = v18
    } else {
        p1 -= p1
        v19 := []int{272, -984, 9223372036854775807}
        v20 := GenS1{F0: 65535, F1: 301, F2: 42}
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    return uint(p0)
}
func genFn1(p0 int64, p1 uint8) int16 {
    if true {
        v21 := []int{42, 2147483647}
        var v22 float64 = -2.25
        _ = v21
        _ = v22
    }
    p0 = (p0 + p0)
    fmt.Println(p0, p1)
    v23 := []int{2147483647}
    _ = p0
    _ = p1
    _ = v23
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 330
    var v3 uint32 = 2
    v4 := make(map[uint32]int)
    var v5 uint = (v2 >> ((57) & 63))
    v6 := []string{"a", "a", ("" + ("" + "hello")), ("hello" + ((("foo bar" + "hello") + "foo bar") + ""))}
    var v24 int8 = -100
    v25 := GenS1{F0: v3, F1: 0, F2: 9223372036854775807}
    v26 := GenS1{F0: v3, F1: 2, F2: 219}
    switch "a" {
    case "":
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v5, v5)
            _ = fi27
        }
    case "foo bar":
        fmt.Println(v24, v3)
    case "\n\t":
        fmt.Println(v2, v24)
    }
    v28 := v4[1]
    if false {
        fmt.Println(v28, v24)
    } else {
        v28 -= 0
        if (((v3 % ((4294967295) | 1)) < (v3 % ((42) | 1))) && (1 > 32767)) {
            var v29 int64 = 1
            _ = v29
        } else {
            v3 = v3
        }
    }
    switch (v28 >> ((v2) & 63)) {
    case 32767:
        if true {
            v28 -= 0
            switch int(v24) {
            case -9223372036854775808:
                var v30 string = "foo bar"
                _ = v30
            case 32767:
                var v31 uint8 = (v1 & v1)
                _ = v31
            case 9223372036854775807:
                v5 += 463
            default:
                if true {
                    var v32 int32 = -508
                    _ = v32
                } else {
                    if true {
                        fmt.Println(v1, v28, v5)
                    } else {
                        v3 -= (v3 - (v3 * v3))
                        var v33 string = (("hello" + "a") + ("a" + "\n\t"))
                        v34 := v4[(v3 & (^(v3)))]
                        _ = v33
                        _ = v34
                    }
                }
            }
        }
    default:
        for fi35 := 0; fi35 < 12; fi35++ {
            fmt.Println(v3, v1)
            _ = len(v6)
            _ = fi35
        }
    }
    if ((1 == -9223372036854775808) && false) {
        if true {
            var v36 uint64 = 18446744073709551615
            if true {
                for fi37 := 0; fi37 < 12; fi37++ {
                    v38 := GenS2{F0: (v3 % ((v3) | 1)), F1: v2}
                    v6 = append(v6, "hello")
                    if false {
                        fmt.Println(v3, fi37, v28)
                    } else {
                        v28 = (v28 * fi37)
                        _ = len(v6)
                    }
                    _ = fi37
                    _ = v38
                }
                var v39 string = ("\n\t" + "foo bar")
                _ = v39
            } else {
                v2 = 255
                switch (v28 << (((^(v5))) & 63)) {
                case 127:
                    for fi40 := 0; fi40 < 18; fi40++ {
                        v41 := GenS3{F0: ("hello" + ("\n\t" + ("a" + ("" + "foo bar")))), F1: (false || ((true && (!(!(-1 <= -1)))) || false))}
                        _ = fi40
                        _ = v41
                    }
                case 128:
                    for fi42 := 0; fi42 < 16; fi42++ {
                        var v43 uint64 = (^(v36))
                        fmt.Println(v1, v5)
                        _ = fi42
                        _ = v43
                    }
                case -968:
                    if len(v6) > 0 {
                        v6[0] = (("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + ((("foo bar" + "a") + ("\n\t" + "foo bar")) + "a"))
                    }
                }
            }
            fmt.Println(v3, v28)
            _ = v36
        } else {
            fmt.Println(v24, v3)
        }
        _ = len(v4)
    }
    v4[(v3 ^ 1)] = (v28 | 898)
    var v44 uint64 = 1
    fmt.Println(v28, v44)
    if ((v3 * 970) == 4294967295) {
        v45 := GenS2{F0: (v3 >> ((v2) & 63)), F1: genFn0(v1, 1)}
        v46 := GenS3{F0: "hello", F1: (true || (true && false))}
        for fi47 := 0; fi47 < 10; fi47++ {
            v24 -= (v24 - -1)
            _ = fi47
        }
        _ = v45
        _ = v46
    }
    var v48 int8 = int8(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v5
    _ = v6
}
