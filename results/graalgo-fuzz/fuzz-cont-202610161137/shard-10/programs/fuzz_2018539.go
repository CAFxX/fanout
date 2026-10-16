package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 127
    v3 := make(map[string]string)
    var v4 float32 = 0.1
    v5 := []uint32{}
    v6 := GenS1{F0: 2147483648}
    v7 := GenS1{F0: 42}
    fmt.Println(v4)
    if (1 > 32767) {
        v8 := GenS1{F0: 276}
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v2, v4)
        for fi10 := 0; fi10 < 7; fi10++ {
            v2 = (v2 - 255)
            _ = fi10
        }
        delete(v3, ("" + (((("a" + ("foo bar" + (("\n\t" + "") + ("\n\t" + "")))) + "") + ("a" + ("hello" + (("a" + ("" + "foo bar")) + "foo bar")))) + "")))
        _ = fi9
    }
    var v11 float32 = v4
    var v12 int8 = 0
    v13 := GenS2{F0: 0, F1: -1}
    v14 := v3[((("" + "") + "hello") + "\n\t")]
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v11 -= float32(1e10)
    v15 := GenS1{F0: 2}
    var v16 int = 999
    var v17 int16 = 572
    v18 := GenS3{F0: 2, F1: 107}
    if ((!false) || false) {
        v14 = ("" + v14)
        v17 += v17
    }
    if (18446744073709551615 > 255) {
        var v19 int64 = -9223372036854775808
        if ((v16 << ((344) & 63)) == v16) {
            switch 811 {
            case 128:
                for fi20 := 0; fi20 < 15; fi20++ {
                    if true {
                        var v21 uint32 = 975
                        _ = v21
                    }
                    if true {
                        if false {
                            if true {
                                var v22 uint16 = 2
                                _ = v22
                            }
                            v16 += 4294967295
                        } else {
                            fmt.Println(v4)
                        }
                    }
                    _ = len(v1)
                    _ = fi20
                }
            default:
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := GenS3{F0: 2147483648, F1: 42}
                    fmt.Println(v14)
                    _ = fi23
                    _ = v24
                }
            }
        }
        v17 = (v17 * (v17 * (v17 % ((v17) | 1))))
        _ = v19
    } else {
        fmt.Println(v11, v11)
        var v25 uint64 = 4294967295
        for fi26 := 0; fi26 < 4; fi26++ {
            switch ((v14 + ("a" + "a")) + v14) {
            case "":
                var v27 bool = (241 < 42)
                _ = v27
            default:
                for fi28 := 0; fi28 < 11; fi28++ {
                    v14 = v14
                    v25 -= (v25 | v25)
                    _ = fi28
                }
            }
            v25 -= (v25 >> ((18446744073709551615) & 63))
            var v29 int = v16
            _ = fi26
            _ = v29
        }
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
