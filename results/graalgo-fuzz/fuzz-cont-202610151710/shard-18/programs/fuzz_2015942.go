package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint8) string {
    var v5 int16 = 1
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint8 = (p0 & p0)
        v7 -= v7
        v8 := GenS1{F0: 699}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = float32(3.14159)
    _ = p0
    _ = v5
    _ = v9
    return (("\n\t" + "\n\t") + ("a" + ("\n\t" + ("" + ((("" + "foo bar") + (("foo bar" + (("" + "a") + (("\n\t" + "a") + ((("\n\t" + "a") + "") + "hello")))) + "foo bar")) + "")))))
}
func genFn1(p0 int) int16 {
    p0 = p0
    var v10 int32 = -2147483648
    switch p0 {
    case -2147483648:
        v11 := GenS2{F0: "foo bar", F1: 4294967295}
        _ = v11
    }
    _ = p0
    _ = v10
    return 32767
}

func main() {
    var v1 float32 = 0.1
    var v2 uint64 = 5
    v3 := make(map[string]string)
    var v4 uint16 = 663
    v1 = 0.1
    v12 := GenS2{F0: "\n\t", F1: 0}
    for fi13 := 0; fi13 < 18; fi13++ {
        switch 4294967295 {
        case -866:
            v14 := GenS2{F0: "foo bar", F1: 0}
            _ = v14
        case 0:
            var v15 uint16 = 255
            _ = v15
        default:
            if false {
                switch 346 {
                case -841:
                    v16 := GenS3{F0: 226, F1: 1}
                    _ = v16
                }
            } else {
                if (-128 == -128) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        v2 = (v2 | 0)
                        _ = fi17
                    }
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v4 = (v4 | v4)
                        _ = fi18
                    }
                    switch (fi13 % ((578) | 1)) {
                    case 32767:
                        if true {
                            v4 -= v4
                            switch "" {
                            case "a":
                                var v19 uint32 = 2147483647
                                _ = v19
                            default:
                                v1 += float32(0.1)
                            }
                        }
                    case -9223372036854775808:
                        v20 := GenS2{F0: "\n\t", F1: 2}
                        _ = v20
                    case 128:
                        var v21 float64 = float64(v1)
                        _ = v21
                    default:
                        v4 -= 0
                    }
                }
            }
        }
        v3["a"] = ("" + "\n\t")
        _ = fi13
    }
    if (!((963 >= -303) || (9223372036854775807 <= 0))) {
        v1 -= v1
        if false {
            switch ("a" + "a") {
            case "\n\t":
                fmt.Println(v2)
            case "":
                _ = len(v3)
            default:
                v4 = v4
            }
            v2 -= 18446744073709551615
            var v22 float32 = v1
            _ = v22
        } else {
            v23 := GenS2{F0: ("" + "\n\t"), F1: 0}
            fmt.Println(v1, v2, v4)
            var v24 uint32 = uint32(v2)
            _ = v23
            _ = v24
        }
    }
    v25 := GenS4{F0: 2}
    v2 = v2
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 float32 = 3.14159
        _ = fi26
        _ = v27
    }
    v4 += v4
    v2 = v2
    fmt.Println(v1, v4)
    v28 := []int{-1, 0}
    v1 += float32(1.5)
    v3["a"] = ("\n\t" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
