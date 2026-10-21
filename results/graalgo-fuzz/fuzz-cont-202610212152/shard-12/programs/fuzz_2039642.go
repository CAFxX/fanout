package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0() uint {
    v7 := []int{}
    var v8 float64 = 3.14159
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint {
    var v9 string = "a"
    if false {
        v10 := GenS1{F0: ((v9 + (v9 + ("a" + (v9 + ("a" + v9))))) + (v9 + (((v9 + v9) + "\n\t") + v9))), F1: 255}
        var v11 uint64 = 1
        _ = v10
        _ = v11
    }
    v9 = v9
    _ = v9
    return 65
}

func main() {
    v1 := []int{}
    var v2 int8 = 127
    v3 := []uint32{}
    v4 := []string{"foo bar", "\n\t", "hello"}
    var v5 uint = 65535
    var v6 uint32 = 65535
    fmt.Println(v5)
    switch 32767 {
    case 0:
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    case -9223372036854775808:
        v12 := GenS1{F0: (("foo bar" + "") + "\n\t"), F1: (v5 | 1)}
        _ = v12
    case 127:
        for fi13 := 0; fi13 < 8; fi13++ {
            v2 = (v2 + int8(v6))
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint64 = 1
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    default:
        if (-736 >= -9223372036854775808) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if false {
                    v1 = append(v1, fi16)
                    _ = len(v1)
                    v3 = append(v3, (v6 | (^(v6))))
                } else {
                    if true {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            fmt.Println(v6)
                            fmt.Println(v5, v5)
                            v18 := GenS2{F0: 75, F1: v5}
                            _ = fi17
                            _ = v18
                        }
                    }
                }
                fmt.Println(v6, v6, v6)
                v5 = (v5 & (v5 + v5))
                _ = fi16
            }
            v6 = uint32(v5)
        } else {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            v6 += (v6 * v6)
            switch ("hello" + (((("hello" + "hello") + "hello") + "hello") + (("hello" + "") + "foo bar"))) {
            case "\n\t":
                var v19 int16 = -322
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 9; fi20++ {
                    v6 -= (v6 >> (((v5 - genFn1())) & 63))
                    var v21 int32 = 1
                    _ = fi20
                    _ = v21
                }
            case "":
                v5 = v5
            }
        }
    }
    fmt.Println(v2)
    v6 -= 4294967295
    _ = len(v1)
    if (1.5 == 1e10) {
        switch "hello" {
        case "a":
            for fi22 := 0; fi22 < 15; fi22++ {
                if len(v1) > 0 {
                    v1[0] = fi22
                }
                var v23 float64 = 1.5
                if len(v1) > 0 {
                    v1[0] = 9223372036854775807
                }
                _ = fi22
                _ = v23
            }
        case "hello":
            v24 := GenS2{F0: (v5 / (((v5 << ((v5) & 63))) | 1)), F1: (^(v5))}
            _ = v24
        case "\n\t":
            fmt.Println(v2)
        default:
            fmt.Println(v5)
        }
    }
    if true {
        v25 := GenS1{F0: "", F1: (v5 >> ((v5) & 63))}
        v5 = 78
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v25
    } else {
        v5 = (v5 | (v5 >> ((v5) & 63)))
        var v26 int64 = -1
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 string = ("foo bar" + "\n\t")
            v28 = v28
            switch (((v28 + v28) + ("hello" + "a")) + v28) {
            case "\n\t":
                v5 += v5
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    var v29 uint8 = 128
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS3{F0: int64(v5), F1: genFn0(), F2: 1}
        v2 = v2
        _ = fi30
        _ = v31
    }
    v32 := GenS4{F0: -834}
    for fi33 := 0; fi33 < 14; fi33++ {
        _ = len(v1)
        var v34 uint64 = 1
        if false {
            fmt.Println(v34, fi33, v34)
            v6 -= v6
            switch ("a" + ((((("a" + "\n\t") + "a") + "") + ("foo bar" + "")) + "hello")) {
            case "foo bar":
                var v35 uint = genFn0()
                _ = v35
            case "a":
                switch "\n\t" {
                case "hello":
                    fmt.Println(v2)
                }
            }
        }
        _ = fi33
        _ = v34
    }
    _ = len(v1)
    v2 = (v2 + (v2 / (((v2 | v2)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
