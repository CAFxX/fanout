package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("hello" + ""), ("hello" + ("a" + "a")), ("hello" + ("hello" + "a")), ("hello" + "\n\t")}
    var v3 uint8 = 127
    var v4 int = 4294967295
    v5 := make(map[string]uint32)
    switch v4 {
    case 362:
        fmt.Println(v3, v4)
    default:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 float64 = (float64((0.1 * 1.5)) * 0.1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
            _ = v7
        }
    }
    if true {
        var v8 int64 = -23
        switch ("foo bar" + "\n\t") {
        case "a":
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v8, v4)
                fmt.Println(v3)
                var v10 int = fi9
                _ = fi9
                _ = v10
            }
        default:
            var v11 uint32 = 2
            _ = v11
        }
        fmt.Println(v4, v3)
        _ = v8
    } else {
        _ = len(v1)
    }
    v12 := GenS1{F0: "foo bar", F1: "\n\t"}
    var v13 string = "a"
    v14 := v1[2147483647]
    v13 = v13
    v3 = 2
    v15 := GenS1{F0: "\n\t", F1: v13}
    var v16 uint8 = v3
    if true {
        if true {
            fmt.Println(v13, v16, v4)
        }
        var v17 int8 = 1
        _ = v17
    } else {
        fmt.Println(v14)
        switch "\n\t" {
        case "hello":
            v4 += (v4 - v14)
        case "a":
            v18 := GenS2{F0: (v13 + "a"), F1: (v13 + ""), F2: int32(v16)}
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 11; fi19++ {
                v14 += fi19
                v20 := GenS1{F0: ((((("\n\t" + v13) + v13) + "foo bar") + "\n\t") + v13), F1: "hello"}
                _ = fi19
                _ = v20
            }
        }
        if (314 > 65535) {
            v5[(((v13 + ((v13 + v13) + (v13 + v13))) + v13) + v13)] = 958
        } else {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 bool = false
                if ((338 <= 2) || ((v16 % (((v3 % ((v16) | 1))) | 1)) <= (v16 & v3))) {
                    v14 -= (^(v4))
                    v3 += v16
                    v23 := GenS1{F0: v13, F1: ("\n\t" + ((((v13 + "\n\t") + "foo bar") + (v13 + "a")) + v13))}
                    _ = v23
                } else {
                    _ = len(v2)
                    var v24 float64 = -2.25
                    for fi25 := 0; fi25 < 18; fi25++ {
                        if (358 > 2) {
                            v16 += (v16 * (v3 + 1))
                        }
                        v26 := GenS3{F0: v13}
                        v22 = true
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                var v27 bool = (0 != -34)
                _ = fi21
                _ = v22
                _ = v27
            }
        }
    }
    var v28 int32 = -2147483648
    fmt.Println(v28, v3)
    v28 = (v28 ^ v28)
    if (false && false) {
        if (!(false && (false && false))) {
            fmt.Println(v3, v16)
        }
        v29 := GenS2{F0: v13, F1: v13, F2: (v28 << ((42) & 63))}
        v30 := GenS4{F0: -505, F1: 18446744073709551615, F2: v16}
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        switch v14 {
        case 2147483647:
            var v32 int8 = -50
            _ = v32
        case -232:
            var v33 uint8 = v16
            _ = v33
        }
        _ = fi31
    }
    fmt.Println(v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
