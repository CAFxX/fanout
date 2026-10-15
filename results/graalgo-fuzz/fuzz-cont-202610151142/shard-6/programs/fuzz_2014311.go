package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
func genFn0(p0 uint, p1 int8, p2 uint64) bool {
    v6 := []int{9223372036854775807, 0, 472}
    switch (("a" + ("" + ("\n\t" + "hello"))) + "a") {
    case "\n\t":
        var v7 int64 = 9223372036854775807
        _ = v7
    case "":
        p2 -= (p2 << (((p0 ^ (p0 | (p0 + p0)))) & 63))
    case "foo bar":
        var v8 uint16 = 1
        _ = v8
    }
    if false {
        if len(v6) > 0 {
            v6[0] = 32768
        }
        switch (("" + ("hello" + ("" + (("" + "") + "")))) + "") {
        case "foo bar":
            if (((!true) || true) || ((true && false) || true)) {
                v9 := GenS1{F0: 1, F1: 710, F2: (^(p1))}
                _ = v9
            }
        default:
            switch "a" {
            case "a":
                if true {
                    var v10 int = 959
                    fmt.Println(p1)
                    _ = v10
                } else {
                    p1 = 0
                }
            default:
                switch 830 {
                case -9223372036854775808:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    p2 = 18446744073709551615
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (false || (((2147483647 <= -2147483648) && ((p2 != (p2 / ((p2) | 1))) || (true && ("a" >= ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))))))) && ((42 > 563) && ((875 == 897) || (((0.0 - 1.5) < ((3.14159 / 1.5) * -2.25)) && (true || true))))))
}
func genFn1(p0 bool) uint16 {
    v11 := GenS1{F0: -2147483648, F1: -431, F2: -1}
    var v12 float64 = 1e10
    var v13 float64 = 0.1
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 32767
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 2
    var v3 uint64 = 18446744073709551615
    var v4 int16 = 32767
    v5 := []int{128, -826}
    v2 += 255
    switch 128 {
    case 763:
        var v14 uint8 = 128
        _ = v14
    case 2147483647:
        v3 = 0
    case -850:
        v3 += v3
    default:
        fmt.Println(v4, v2)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        switch ("\n\t" + (("foo bar" + ("\n\t" + "a")) + "a")) {
        case "foo bar":
            v3 = (v3 | v3)
        case "a":
            var v16 int64 = -9223372036854775808
            _ = v16
        default:
            v4 = v4
        }
        fmt.Println(v2, v3, fi15)
        _ = fi15
    }
    v3 -= 4294967295
    for fi17 := 0; fi17 < 6; fi17++ {
        v3 -= v3
        _ = fi17
    }
    v4 = 753
    v3 = 1
    for fi18 := 0; fi18 < 11; fi18++ {
        switch "a" {
        case "":
            if (-2.25 > 0.1) {
                v19 := GenS1{F0: -2147483648, F1: -2147483648, F2: -128}
                v4 = v4
                for fi20 := 0; fi20 < 6; fi20++ {
                    v3 += (^(v3))
                    var v21 int8 = 1
                    switch "\n\t" {
                    case "\n\t":
                        switch ((("" + "") + "\n\t") + ("\n\t" + "hello")) {
                        case "":
                            v22 := GenS1{F0: 897, F1: 0, F2: (-(v21))}
                            _ = v22
                        case "a":
                            var v23 float32 = 0.0
                            _ = v23
                        case "hello":
                            v3 -= v3
                        default:
                            var v24 float32 = (0.1 + 3.14159)
                            _ = v24
                        }
                    case "foo bar":
                        if true {
                            var v25 uint64 = 4294967295
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v25
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        case "a":
            if (false && (int8(v2) != 100)) {
                if false {
                    fmt.Println(fi18, v4)
                } else {
                    var v26 uint32 = 2147483648
                    delete(v1, (v26 % ((v26) | 1)))
                    v5 = append(v5, fi18)
                    _ = v26
                }
            }
        case "foo bar":
            v4 += v4
        }
        _ = fi18
    }
    fmt.Println(v3, v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
