package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0() string {
    var v6 uint64 = 42
    fmt.Println(v6)
    v7 := GenS1{F0: 65535, F1: 2}
    v6 -= v6
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 42
    var v3 uint16 = 42
    v4 := make(map[uint32]bool)
    var v5 uint32 = 0
    v5 = 0
    var v8 bool = false
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{(fi9 % ((fi9) | 1)), fi9, fi9}
        for fi11 := 0; fi11 < 5; fi11++ {
            if (!v8) {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v2 = v2
            }
            _ = len(v10)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    var v12 uint8 = 0
    v5 = v5
    if (2 <= 255) {
        if (v8 || (v8 || (!v8))) {
            if v8 {
                if (!v8) {
                    v13 := v4[(^(v5))]
                    _ = v13
                }
                fmt.Println(v8)
            } else {
                v14 := []int{(v2 >> ((232) & 63)), -9223372036854775808, 494, v2}
                delete(v4, (v5 << ((65535) & 63)))
                _ = v14
            }
            var v15 string = "foo bar"
            v3 = 15
            _ = v15
        }
        switch v2 {
        case 32768:
            var v16 uint16 = (v3 | v3)
            _ = v16
        case 127:
            if (534 > 418) {
                var v17 uint8 = (v12 & v12)
                switch ("foo bar" + ("hello" + "hello")) {
                case "a":
                    switch "" {
                    case "\n\t":
                        var v18 int8 = -128
                        _ = v18
                    case "foo bar":
                        v19 := GenS1{F0: uint32(v17), F1: (v12 | (v17 << ((65535) & 63)))}
                        _ = v19
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            for fi20 := 0; fi20 < 12; fi20++ {
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    var v22 int8 = 1
                                    var v23 uint8 = 1
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi20
                            }
                        case "hello":
                            fmt.Println(v17)
                        }
                    default:
                        v2 += 32767
                    }
                case "foo bar":
                    v24 := []int{491}
                    _ = v24
                case "\n\t":
                    v25 := []int{(v2 / ((-1) | 1)), -2147483648, (v2 + -1)}
                    _ = v25
                }
                _ = v17
            }
        }
        v26 := GenS1{F0: (^(v5)), F1: v12}
        _ = v26
    } else {
        switch v2 {
        case 127:
            switch "foo bar" {
            case "":
                v8 = v8
            default:
                v12 = v12
            }
        }
        fmt.Println(v3)
        v27 := GenS1{F0: (v5 | (^(v5))), F1: 41}
        _ = v27
    }
    if v8 {
        v2 = (v2 / (((v2 + (v2 & v2))) | 1))
        switch v2 {
        case 42:
            var v28 uint = 2
            _ = v28
        default:
            var v29 int64 = -651
            _ = v29
        }
        fmt.Println(v3, v5)
    }
    v30 := []int{-2147483648, (v2 | (v2 >> ((587) & 63)))}
    v31 := GenS1{F0: (v5 ^ v5), F1: uint8(v3)}
    if (32767 != -32768) {
        if v8 {
            v32 := GenS1{F0: (v5 >> ((2) & 63)), F1: v12}
            _ = v32
        }
        var v33 uint64 = 2
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v8
}
