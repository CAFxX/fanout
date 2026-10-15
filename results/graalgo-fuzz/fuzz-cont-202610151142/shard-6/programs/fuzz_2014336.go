package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := []string{("hello" + ("foo bar" + ("foo bar" + "")))}
    v2 := []uint32{0}
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 int64 = 9223372036854775807
    v6 := make(map[string]uint32)
    v1 = append(v1, "")
    v1 = append(v1, ("\n\t" + (("hello" + "hello") + ("\n\t" + "hello"))))
    v7 := GenS1{F0: 1, F1: -32768, F2: 2147483647}
    switch 658 {
    case 32768:
        switch "hello" {
        case "foo bar":
            v4["foo bar"] = ("a" + (("" + "a") + "\n\t"))
        case "\n\t":
            if (!false) {
                v8 := GenS2{F0: 0}
                _ = v8
            } else {
                var v9 int8 = -1
                switch (("a" + (((("\n\t" + ("\n\t" + "")) + "a") + "") + "")) + "") {
                case "":
                    if (1 == -1) {
                        v9 += v9
                        var v10 int64 = (^(v5))
                        v9 = (v9 ^ v9)
                        _ = v10
                    }
                case "a":
                    v1 = append(v1, "foo bar")
                default:
                    if ((!false) || false) {
                        switch 87 {
                        case 0:
                            if false {
                                var v11 uint32 = 4294967295
                                for fi12 := 0; fi12 < 12; fi12++ {
                                    fmt.Println(v11)
                                    if (!(0 <= 948)) {
                                        var v13 int = fi12
                                        v14 := GenS1{F0: v9, F1: -1, F2: (^(fi12))}
                                        for fi15 := 0; fi15 < 6; fi15++ {
                                            v6["a"] = (^(v11))
                                            var v16 float32 = 0.1
                                            if true {
                                                v16 += (0.0 / (1.5 + v16))
                                                v17 := GenS3{F0: v5, F1: (-(v13))}
                                                if false {
                                                    var v18 uint8 = 1
                                                    _ = v18
                                                }
                                                if (true || false) {
                                                    v9 -= -128
                                                    switch v13 {
                                                    case 307:
                                                        v9 += v9
                                                    }
                                                    v19 := v3["\n\t"]
                                                    _ = v19
                                                }
                                                v9 = v9
                                                _ = v17
                                            }
                                            _ = fi15
                                            _ = v16
                                        }
                                        _ = v13
                                        _ = v14
                                    }
                                    _ = fi12
                                }
                                _ = v11
                            } else {
                                v5 = v5
                                if (1 < -242) {
                                    var v20 int32 = 1
                                    var v21 bool = false
                                    v5 += 42
                                    _ = v20
                                    _ = v21
                                }
                            }
                        }
                    }
                }
                _ = v9
            }
        default:
            v5 -= 42
        }
    case 285:
        v5 = (v5 >> ((0) & 63))
    case 2147483647:
        var v22 bool = true
        _ = v22
    default:
        var v23 int = 0
        _ = v23
    }
    var v24 bool = (true || true)
    v24 = false
    var v25 int32 = -877
    var v26 int8 = 42
    switch -603 {
    case -1:
        fmt.Println(v26, v5)
    case -9223372036854775808:
        for fi27 := 0; fi27 < 18; fi27++ {
            v25 = 2147483647
            _ = fi27
        }
    default:
        switch 2147483647 {
        case -1:
            if len(v2) > 0 {
                v2[0] = 976
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
