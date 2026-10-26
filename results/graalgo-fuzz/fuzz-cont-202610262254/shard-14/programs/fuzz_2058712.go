package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int, p2 int32) uint8 {
    if ((!true) || true) {
        v7 := GenS1{F0: -128, F1: p1}
        _ = v7
    }
    var v8 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    v1 := []int{704, 32767}
    v2 := make(map[int]string)
    var v3 int = -2147483648
    var v4 bool = (948 == -602)
    var v5 int = (v3 & v3)
    var v6 bool = ((true || v4) || false)
    for fi9 := 0; fi9 < 20; fi9++ {
        v3 = v5
        _ = fi9
    }
    switch 127 {
    case 32767:
        if len(v1) > 0 {
            v1[0] = (v3 / ((v5) | 1))
        }
    case -908:
        if (!v4) {
            var v10 uint64 = 18446744073709551615
            var v11 float32 = (float32(3.14159) - float32(1.5))
            if v4 {
                v4 = v6
            } else {
                var v12 int16 = -1
                for fi13 := 0; fi13 < 14; fi13++ {
                    if ((v5 | (fi13 - 9223372036854775807)) > v5) {
                        fmt.Println(v11, v12)
                        switch ("foo bar" + "a") {
                        case "":
                            var v14 float64 = ((((float64(v11) * ((((float64(v11) + (1e10 * (3.14159 + 1.5))) + (1e10 + ((((1e10 * (0.0 - float64(0.1))) + (0.0 + 1.5)) + float64(v11)) + 0.1))) / float64(v11)) / 1e10)) * -2.25) + 1.5) * 1.5)
                            _ = v14
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if ("a" > "") {
                            var v15 int = 9223372036854775807
                            _ = v15
                        }
                    } else {
                        if false {
                            var v16 int16 = 302
                            v1 = append(v1, (v5 | 316))
                            _ = v16
                        } else {
                            v1 = append(v1, 42)
                            for fi17 := 0; fi17 < 17; fi17++ {
                                v4 = v4
                                var v18 string = (("foo bar" + "a") + ("hello" + (("hello" + "\n\t") + (("\n\t" + "a") + ((("foo bar" + "\n\t") + "foo bar") + ("foo bar" + "foo bar"))))))
                                _ = fi17
                                _ = v18
                            }
                            delete(v2, (fi13 / (((v3 / (((v3 << ((2) & 63))) | 1))) | 1)))
                        }
                        var v19 string = "foo bar"
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(v11)
                        v6 = true
                        _ = fi20
                    }
                    _ = fi13
                }
                v1 = append(v1, v5)
                _ = v12
            }
            _ = v10
            _ = v11
        }
    default:
        var v21 string = ((((("foo bar" + "foo bar") + "a") + "\n\t") + ((("a" + "a") + (("\n\t" + "\n\t") + ("a" + "hello"))) + "\n\t")) + "foo bar")
        _ = v21
    }
    v3 += v5
    fmt.Println(v5, v4, v6)
    v6 = v6
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint16 = 799
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 42, F1: v5}
    if v6 {
        switch "foo bar" {
        case "hello":
            v25 := GenS2{F0: v4, F1: 0}
            _ = v25
        case "\n\t":
            var v26 int = 32767
            _ = v26
        case "foo bar":
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v4, v3)
                _ = fi27
            }
        }
    } else {
        var v28 uint = 0
        _ = v28
    }
    v6 = (((-398 > 1) || (-9223372036854775808 > 536)) && (42 <= 1))
    fmt.Println(v5, v6)
    if (672 > 597) {
        if (1 >= -61) {
            var v29 float32 = 0.1
            _ = v29
        } else {
            switch "hello" {
            case "foo bar":
                v5 = (v5 & v3)
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    v30 := GenS1{F0: 0, F1: (^(v5))}
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 int = 155
                    v33 := GenS2{F0: (!(63 != 240)), F1: 2147483647}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
            fmt.Println(v6)
        }
        switch "a" {
        case "\n\t":
            v34 := GenS2{F0: (v4 || v4), F1: 1}
            _ = v34
        }
    } else {
        v35 := GenS3{F0: 655}
        v36 := GenS2{F0: v6, F1: 1}
        var v37 uint8 = 140
        _ = v35
        _ = v36
        _ = v37
    }
    v38 := GenS1{F0: 42, F1: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
