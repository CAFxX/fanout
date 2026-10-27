package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 bool = (float32(1e10) < ((0.1 * 0.0) - 1.5))
    v2 := []string{"a", "\n\t", (("" + "hello") + "foo bar"), ("a" + "")}
    v3 := make(map[string]int)
    var v4 int64 = -1
    v4 -= v4
    v1 = (0.0 != ((0.1 / 1.0) + ((0.0 * ((0.1 / (-2.25 + 1e10)) + 0.0)) / (-2.25 - (1e10 + -2.25)))))
    _ = len(v2)
    if v1 {
        v4 += (v4 % ((9223372036854775807) | 1))
        var v5 uint64 = 18446744073709551615
        _ = v5
    } else {
        fmt.Println(v1, v4)
        var v6 int64 = v4
        _ = v6
    }
    if true {
        if (!(!v1)) {
            var v7 uint32 = 229
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v7 = 909
                    var v10 int8 = 0
                    if ((v1 || ((-(fi9)) >= fi9)) || true) {
                        v7 = (^(v7))
                        var v11 uint8 = 42
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                for fi12 := 0; fi12 < 15; fi12++ {
                    v4 -= 42
                    switch "" {
                    case "hello":
                        v13 := GenS1{F0: 353}
                        _ = v13
                    default:
                        var v14 uint8 = 0
                        _ = v14
                    }
                    _ = fi12
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    _ = len(v2)
                    switch (fi15 << ((65535) & 63)) {
                    case 1:
                        if v1 {
                            _ = len(v2)
                        }
                    case 42:
                        if v1 {
                            _ = len(v3)
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi15
                }
                _ = fi8
            }
            var v16 uint32 = v7
            _ = v16
            _ = v7
        } else {
            v4 -= (v4 * (^(v4)))
        }
        if ((-519 == -120) && false) {
            var v17 uint = 0
            if (v1 && (42 > uint64(v17))) {
                var v18 int8 = 100
                var v19 string = ("hello" + ("" + ""))
                _ = v18
                _ = v19
            } else {
                v17 = 0
                v4 -= v4
            }
            _ = v17
        }
    } else {
        var v20 float64 = (1e10 + 1e10)
        v4 += 787
        _ = v20
    }
    var v21 uint32 = 2
    for fi22 := 0; fi22 < 16; fi22++ {
        v21 += (v21 * 0)
        var v23 int64 = (v4 % (((v4 - (v4 >> ((65535) & 63)))) | 1))
        v24 := GenS1{F0: 0}
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println(v21, v21)
    switch 0 {
    case 425:
        var v25 int = 2147483647
        _ = v25
    case 1:
        if (248 <= 128) {
            _ = len(v2)
            fmt.Println(v4)
            if (839 >= 402) {
                v4 = v4
                fmt.Println(v4, v1, v1)
                v1 = v1
            } else {
                v1 = (693 > 2)
                if false {
                    v21 = 692
                    v26 := v3[(("foo bar" + "\n\t") + "")]
                    _ = v26
                }
                v1 = v1
            }
        }
    case 9223372036854775807:
        switch 32767 {
        case 976:
            for fi27 := 0; fi27 < 19; fi27++ {
                _ = len(v3)
                _ = fi27
            }
        case 806:
            var v28 int32 = 379
            _ = v28
        default:
            v4 = (v4 * v4)
        }
    }
    v29 := GenS1{F0: -2147483648}
    fmt.Println(v4, v1, v21)
    if (true || v1) {
        var v30 uint16 = 605
        v1 = (42 <= -82)
        _ = v30
    } else {
        v31 := GenS2{F0: 688}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
}
