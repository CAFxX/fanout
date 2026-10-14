package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0() uint16 {
    var v5 float32 = 3.14159
    v5 -= 0.0
    _ = v5
    return 65535
}

func main() {
    v1 := []int{}
    var v2 uint = 1
    v3 := make(map[string]uint32)
    var v4 uint32 = 2147483647
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 544
        var v9 uint8 = (^(v6))
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: v4, F1: v2, F2: 876}
            var v12 int64 = 0
            switch (fi10 | -9223372036854775808) {
            case -2147483648:
                v13 := GenS1{F0: uint32(v6), F1: 833, F2: uint16(v8)}
                _ = v13
            case 32767:
                var v14 int = fi10
                v3["\n\t"] = (v4 & (v4 - (v4 | (v4 + v4))))
                v6 = (v6 / ((v9) | 1))
                v3["\n\t"] = v4
                var v15 string = ("\n\t" + "a")
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v6)
    v6 = v6
    v4 = v4
    var v16 uint8 = (v6 << (((v2 >> ((v2) & 63))) & 63))
    v17 := GenS2{F0: 0, F1: 871}
    fmt.Println(v2, v2)
    switch "foo bar" {
    case "":
        var v18 int8 = 0
        _ = v18
    case "\n\t":
        fmt.Println(v4, v6)
    default:
        v16 += (^(v16))
    }
    switch "foo bar" {
    case "\n\t":
        var v19 int8 = int8(v16)
        _ = v19
    default:
        if (!(!((("foo bar" <= "a") || false) || true))) {
            delete(v3, "a")
            var v20 uint64 = 0
            _ = v20
        }
    }
    if (!true) {
        fmt.Println(v6)
    } else {
        v6 -= (^(v16))
        for fi21 := 0; fi21 < 10; fi21++ {
            v6 += (^(v6))
            _ = len(v1)
            v4 += v4
            _ = fi21
        }
    }
    var v22 uint8 = (v6 + 42)
    v22 = uint8(v4)
    switch 42 {
    case -762:
        if len(v1) > 0 {
            v1[0] = 31
        }
    default:
        var v23 int = 0
        _ = v23
    }
    var v24 float32 = 1.5
    if true {
        v4 -= 682
        switch "a" {
        case "hello":
            if (2 <= 2) {
                v1 = append(v1, -706)
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 int32 = 0
                    var v27 float32 = -2.25
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                var v28 int32 = 751
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(v4, v28, v4)
                    if (false || (1 <= 1)) {
                        switch fi29 {
                        case 2147483647:
                            v2 = 2
                        case -1:
                            if (uint32(v2) < v4) {
                                var v30 uint8 = (v16 >> ((65535) & 63))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v30
                            }
                        }
                        fmt.Println(v22)
                        v6 = (^(v16))
                    }
                    _ = fi29
                }
                _ = v28
            }
        }
    } else {
        if true {
            fmt.Println(v16, v4, v4)
        } else {
            v2 += 4294967295
            v31 := v3["a"]
            _ = v31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v6
}
