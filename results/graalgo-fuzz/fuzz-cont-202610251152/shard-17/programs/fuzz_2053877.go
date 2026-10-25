package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = "hello"
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    var v5 int = 9223372036854775807
    v5 += (v5 ^ (^(v5)))
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (true || true) {
            var v7 int16 = -412
            if true {
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int = (fi6 ^ v5)
                    v10 := GenS1{F0: uint16(v1), F1: 127, F2: -1}
                    if false {
                        v1 = (v1 | v1)
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v5 = 4294967295
                v5 += fi6
            } else {
                var v11 int64 = (v3 << ((0) & 63))
                if (false || ("a" != (v2 + (v2 + ((v2 + ((("foo bar" + v2) + ("" + (v2 + v2))) + (("\n\t" + ("\n\t" + v2)) + v2))) + ("" + v2)))))) {
                    switch fi6 {
                    case 127:
                        var v12 uint16 = 42
                        _ = v12
                    }
                    v13 := v4[(fi6 % ((0) | 1))]
                    v14 := GenS1{F0: 32767, F1: 127, F2: 1}
                    _ = v13
                    _ = v14
                } else {
                    if (!false) {
                        var v15 string = "\n\t"
                        var v16 uint32 = 65535
                        var v17 float64 = (1e10 * -2.25)
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
                _ = v11
            }
            _ = v7
        }
        v18 := []int{v5, (v5 / (((v5 / ((fi6) | 1))) | 1)), (fi6 - (fi6 << ((2) & 63))), fi6}
        v19 := GenS1{F0: 255, F1: 127, F2: -183}
        _ = fi6
        _ = v18
        _ = v19
    }
    v2 = v2
    v4[(v5 & 127)] = v2
    switch v2 {
    case "foo bar":
        var v20 uint16 = 1
        _ = v20
    case "":
        var v21 float64 = 3.14159
        _ = v21
    case "a":
        if false {
            v22 := GenS1{F0: 32767, F1: -128, F2: -2147483648}
            _ = v22
        }
    default:
        if true {
            for fi23 := 0; fi23 < 3; fi23++ {
                switch (v2 + v2) {
                case "hello":
                    if false {
                        var v24 int32 = -2147483648
                        v25 := GenS1{F0: 987, F1: 127, F2: (v24 / ((-2147483648) | 1))}
                        var v26 float32 = (((1.5 / -2.25) / 1.5) * 1.5)
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                case "a":
                    var v27 int32 = -2147483648
                    _ = v27
                }
                if (true || (32767 >= 32767)) {
                    var v28 int64 = v3
                    var v29 bool = true
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v5)
                }
                _ = fi23
            }
            if (-128 >= -128) {
                switch v2 {
                case "":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 int = 1
                        var v32 float64 = 1.5
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case "a":
                    var v33 uint64 = (v1 * 4294967295)
                    _ = v33
                case "foo bar":
                    v34 := []int{127, 127, v5}
                    _ = v34
                default:
                    var v35 uint16 = 42
                    _ = v35
                }
                delete(v4, (v5 * int(v3)))
            }
        } else {
            var v36 int = (^(v5))
            v37 := []int{(v5 * v5), 775, (v36 + v5), (v36 | int(v1))}
            _ = v36
            _ = v37
        }
    }
    v3 += v3
    var v38 uint8 = 0
    v39 := GenS2{F0: (v1 * v1), F1: 42, F2: "a"}
    fmt.Println(v1, v3, v3)
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS1{F0: 32767, F1: 1, F2: 0}
        for fi42 := 0; fi42 < 10; fi42++ {
            if (-372 != 462) {
                v5 -= 4294967295
            } else {
                if (0.1 != (1.5 + 1e10)) {
                    v43 := []int{fi42}
                    _ = v43
                } else {
                    var v44 uint64 = (v1 / ((v1) | 1))
                    _ = v44
                }
                var v45 uint32 = 65535
                _ = v45
            }
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    if true {
        v3 -= v3
        v3 = -9223372036854775808
        if true {
            fmt.Println(v3, v38, v3)
            var v46 uint8 = 0
            v47 := GenS1{F0: 65535, F1: 60, F2: 474}
            _ = v46
            _ = v47
        }
    } else {
        var v48 uint16 = 0
        v49 := GenS2{F0: 969, F1: uint(v1), F2: ("\n\t" + v2)}
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
