package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }

func main() {
    var v1 int64 = 365
    var v2 float32 = 3.14159
    v3 := []int{2147483647, 32768}
    var v4 uint64 = 1
    v2 = (v2 - (0.1 * v2))
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v2 = (float32(0.0) + 0.0)
            _ = fi6
        }
        v2 = float32(0.1)
        _ = fi5
    }
    v2 -= v2
    if (65535 != 2) {
        fmt.Println(v1)
    }
    v7 := GenS1{F0: "hello", F1: "a", F2: 32767}
    switch (("a" + "\n\t") + ("foo bar" + "a")) {
    case "foo bar":
        var v8 uint8 = 128
        _ = v8
    default:
        fmt.Println(v1, v4, v2)
    }
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 20; fi9++ {
            if (!((fi9 > (fi9 - (-(fi9)))) || false)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if (!true) {
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: "hello", F1: (((("a" + "hello") + ("\n\t" + "\n\t")) + "a") + "foo bar"), F2: fi10}
                    _ = fi10
                    _ = v11
                }
                if (false || false) {
                    if false {
                        v1 = 414
                        v2 -= -2.25
                    } else {
                        fmt.Println(fi9, v1)
                    }
                } else {
                    if (32767 < 957) {
                        switch (fi9 & 1) {
                        case 127:
                            for fi12 := 0; fi12 < 9; fi12++ {
                                v1 = 9223372036854775807
                                v2 += 0.1
                                var v13 int32 = 2147483647
                                _ = fi12
                                _ = v13
                            }
                        default:
                            if (-2147483648 == -1) {
                                var v14 uint64 = v4
                                v2 = -2.25
                                v14 += (v4 ^ 1)
                                _ = v14
                            } else {
                                for fi15 := 0; fi15 < 10; fi15++ {
                                    var v16 uint8 = 1
                                    v2 = v2
                                    v1 = (v1 + 9223372036854775807)
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        }
                        v2 = 0.0
                    } else {
                        var v17 int = fi9
                        switch "hello" {
                        case "":
                            _ = len(v3)
                        case "hello":
                            var v18 int32 = -1
                            _ = v18
                        }
                        if false {
                            var v19 bool = false
                            if len(v3) > 0 {
                                v3[0] = (^(v17))
                            }
                            _ = v19
                        }
                        _ = v17
                    }
                    var v20 uint8 = 42
                    _ = v20
                }
                if true {
                    v3 = append(v3, (fi9 & (fi9 * -2147483648)))
                    _ = len(v3)
                    _ = len(v3)
                }
            } else {
                var v21 int = fi9
                switch 32767 {
                case -241:
                    v22 := GenS2{F0: v4, F1: 556, F2: -32768}
                    _ = v22
                }
                var v23 uint = 666
                _ = v21
                _ = v23
            }
            v4 = v4
            _ = fi9
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            _ = len(v3)
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println(v4, v1)
    if (!true) {
        var v26 uint16 = 255
        _ = v26
    } else {
        v27 := GenS2{F0: 4294967295, F1: 34, F2: -338}
        var v28 int32 = -2147483648
        _ = v27
        _ = v28
    }
    var v29 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v7
}
