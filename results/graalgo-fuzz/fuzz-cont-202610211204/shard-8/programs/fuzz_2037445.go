package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int; F1 int; F2 uint64 }
func genFn0(p0 int32) uint {
    var v4 uint32 = 42
    v5 := GenS1{F0: (p0 >> ((103) & 63)), F1: 111}
    v4 = (v4 << ((2) & 63))
    fmt.Println(p0, v4)
    v6 := []int{9223372036854775807, -9223372036854775808, 32767, 4294967295}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            var v8 int32 = -1
            _ = v8
        } else {
            v9 := GenS2{F0: 2, F1: fi7}
            _ = v9
        }
        var v10 uint32 = 666
        _ = fi7
        _ = v10
    }
    v11 := GenS3{F0: -125, F1: false}
    var v12 uint8 = 1
    v13 := []int{4294967295, 4294967295}
    v13 = append(v13, 774)
    _ = v11
    _ = v12
    _ = v13
    return 988
}

func main() {
    var v1 bool = (1000 < -886)
    v2 := []string{(("a" + ("\n\t" + "")) + ""), ("foo bar" + ((("foo bar" + "") + (("hello" + "hello") + "foo bar")) + "hello")), ("" + "\n\t"), (((("hello" + "foo bar") + ("\n\t" + (("a" + ("a" + "hello")) + "a"))) + "hello") + "")}
    v3 := make(map[uint32]string)
    var v14 int16 = -166
    fmt.Println(v14)
    fmt.Println(v14, v14)
    fmt.Println(v14)
    if (809 > 1) {
        switch "" {
        case "a":
            v15 := GenS2{F0: 454, F1: -761}
            _ = v15
        default:
            fmt.Println(v1, v1)
        }
        v16 := v3[2147483647]
        _ = v16
    }
    var v17 uint = 632
    if false {
        v17 -= (v17 | (v17 & (v17 ^ v17)))
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            v17 += 255
            _ = fi18
        }
        switch "" {
        case "":
            for fi19 := 0; fi19 < 11; fi19++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v20 uint16 = 255
                v21 := GenS4{F0: 9223372036854775807, F1: 65535, F2: 1}
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            v22 := GenS3{F0: 116, F1: (!(2147483647 >= 821))}
            _ = v22
        }
        switch ("foo bar" + "hello") {
        case "a":
            if v1 {
                switch 999 {
                case -560:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v2 = append(v2, "foo bar")
                        switch ("a" + (("\n\t" + "hello") + ("\n\t" + (("\n\t" + "a") + ("a" + (((("" + "\n\t") + ("\n\t" + ("a" + ("" + ("\n\t" + "hello"))))) + "") + ("a" + "foo bar"))))))) {
                        case "\n\t":
                            var v24 float32 = (-2.25 / 1.5)
                            _ = v24
                        case "a":
                            var v25 float32 = 0.1
                            _ = v25
                        }
                        if false {
                            fmt.Println(v1, v1)
                        }
                        _ = fi23
                    }
                case -177:
                    v26 := GenS5{F0: int32(v14), F1: -1}
                    _ = v26
                default:
                    var v27 uint = (^(v17))
                    _ = v27
                }
            } else {
                var v28 bool = v1
                _ = v28
            }
        case "\n\t":
            fmt.Println(v14, v14)
        default:
            v29 := GenS6{F0: 343, F1: 997, F2: 148}
            _ = v29
        }
    }
    var v30 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v30
}
