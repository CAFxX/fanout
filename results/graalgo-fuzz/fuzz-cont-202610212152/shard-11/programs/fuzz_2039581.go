package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0(p0 uint32) string {
    p0 = (p0 >> ((394) & 63))
    fmt.Println(p0)
    _ = p0
    return ("foo bar" + "a")
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int16 = 1000
    var v4 uint32 = 427
    var v5 int8 = 127
    v6 := []uint32{(v4 >> ((0) & 63))}
    if len(v6) > 0 {
        v6[0] = v4
    }
    var v7 uint64 = 18446744073709551615
    v3 -= (v3 + v3)
    if true {
        if true {
            v3 += -32768
            v4 = 0
            fmt.Println(v5, v4, v4)
        } else {
            var v8 int = 32768
            switch "foo bar" {
            case "":
                for fi9 := 0; fi9 < 1; fi9++ {
                    v3 += (v3 | 1)
                    var v10 string = ("a" + ("\n\t" + ((("hello" + "a") + "") + (("" + "hello") + ("a" + "\n\t")))))
                    _ = fi9
                    _ = v10
                }
            case "hello":
                v5 = 100
            default:
                v3 -= (v3 >> ((0) & 63))
            }
            _ = v8
        }
    }
    if (!false) {
        var v11 float32 = 1.5
        _ = v11
    } else {
        v4 = 1
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v5 -= (v5 - v5)
        v13 := GenS1{F0: 0, F1: "\n\t"}
        _ = fi12
        _ = v13
    }
    v14 := v2[851]
    if (!true) {
        var v15 float32 = 0.0
        _ = v15
    }
    v7 += (v7 >> ((1) & 63))
    switch v14 {
    case "\n\t":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint8 = 128
            if false {
                fmt.Println(fi16, v7)
                v17 += v17
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS1{F0: (v5 >> ((18446744073709551615) & 63)), F1: "hello"}
                    fmt.Println(v17)
                    _ = fi18
                    _ = v19
                }
            } else {
                v14 = v14
                v20 := v1[fi16]
                _ = len(v2)
                _ = v20
            }
            v21 := GenS1{F0: 21, F1: ((v14 + ("hello" + ((("a" + "") + v14) + "\n\t"))) + v14)}
            _ = fi16
            _ = v17
            _ = v21
        }
    case "":
        var v22 int32 = -210
        _ = v22
    default:
        if (!true) {
            for fi23 := 0; fi23 < 19; fi23++ {
                v3 += v3
                var v24 int64 = -1
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi23
                _ = v24
            }
            switch v14 {
            case "foo bar":
                for fi25 := 0; fi25 < 19; fi25++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi25
                }
            case "\n\t":
                var v26 int8 = -128
                _ = v26
            case "":
                fmt.Println(v14)
            default:
                v27 := GenS1{F0: v5, F1: v14}
                _ = v27
            }
        }
    }
    v7 += 1
    v28 := GenS2{F0: 66, F1: v5, F2: 255}
    v3 -= (-(v3))
    var v29 int8 = (v5 & -52)
    var v30 int8 = v29
    delete(v2, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
