package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 uint64 }
func genFn0(p0 uint8) bool {
    p0 -= 127
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -1, F1: 1}
        var v9 bool = ("foo bar" <= (((("a" + "hello") + "\n\t") + ("foo bar" + ("" + "foo bar"))) + ("a" + "hello")))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = -387
    if false {
        if (1 != 18446744073709551615) {
            fmt.Println(p0, v10)
            v11 := []int{-2147483648, -951, 32768}
            _ = v11
        }
        v12 := []int{-760, -9223372036854775808, 42, -1}
        var v13 uint8 = (^(p0))
        _ = v12
        _ = v13
    }
    v10 = v10
    _ = p0
    _ = v10
    return (!false)
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]int)
    var v3 int = 397
    v4 := []uint32{2, 48, 544, 2147483648}
    v5 := []uint32{4294967295, 1}
    v6 := make(map[uint32]int)
    var v14 uint32 = 65535
    var v15 int32 = 2147483647
    var v16 uint32 = v14
    v17 := GenS2{F0: 18446744073709551615, F1: 42, F2: true}
    switch (v3 ^ (v3 / ((v3) | 1))) {
    case -9223372036854775808:
        for fi18 := 0; fi18 < 5; fi18++ {
            v16 -= v16
            _ = fi18
        }
    case 32767:
        v15 = (v15 | (v15 << ((42) & 63)))
    }
    fmt.Println(v16, v14)
    v19 := GenS3{F0: 334, F1: (v15 % (((-(v15))) | 1))}
    v2["a"] = v3
    if (((1.5 - (3.14159 + -2.25)) > -2.25) && true) {
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS3{F0: 42, F1: -430}
            _ = fi20
            _ = v21
        }
        if false {
            var v22 int = v3
            _ = v22
        }
        v23 := GenS4{F0: 65535, F1: v15, F2: 18446744073709551615}
        _ = v23
    } else {
        if (!true) {
            var v24 int8 = -1
            var v25 int8 = v24
            _ = v24
            _ = v25
        }
    }
    v4 = append(v4, (v16 - (v14 - (v16 << ((2) & 63)))))
    fmt.Println(v3, v1, v16)
    v26 := GenS4{F0: 65535, F1: (v15 ^ v15), F2: 1}
    var v27 uint16 = 2
    _ = len(v2)
    if false {
        v14 = v16
    } else {
        v16 = (v14 + (v14 + uint32(v27)))
    }
    var v28 int64 = 9223372036854775807
    v14 += 2
    if false {
        if len(v4) > 0 {
            v4[0] = (v14 << ((0) & 63))
        }
        switch "hello" {
        case "\n\t":
            var v29 float32 = -2.25
            _ = v29
        case "a":
            v28 = v28
        case "":
            fmt.Println(v14, v27, v3)
        }
    } else {
        switch (v3 >> ((363) & 63)) {
        case -2147483648:
            delete(v2, "\n\t")
        case 9223372036854775807:
            v15 = 463
        default:
            if true {
                _ = len(v5)
                var v30 int = (-(v3))
                v15 -= 455
                _ = v30
            } else {
                v15 -= (v15 + v15)
                v5 = append(v5, 4294967295)
                switch ("\n\t" + "foo bar") {
                case "a":
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 string = "\n\t"
                        for fi33 := 0; fi33 < 2; fi33++ {
                            var v34 uint8 = 255
                            var v35 uint8 = (v34 * 96)
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v32 = v32
                        _ = fi31
                        _ = v32
                    }
                default:
                    v27 = v27
                }
            }
        }
        v14 = (v14 ^ 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
