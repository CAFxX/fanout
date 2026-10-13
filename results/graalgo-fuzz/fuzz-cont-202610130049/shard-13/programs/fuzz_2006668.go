package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 int16) int {
    v4 := GenS1{F0: -2147483648, F1: 533}
    var v5 uint16 = 255
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2
    var v3 uint32 = 2
    v6 := v1[2147483647]
    var v7 uint8 = 174
    v8 := GenS1{F0: -16, F1: v2}
    var v9 int16 = -78
    switch -552 {
    case -1:
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (("foo bar" + ("foo bar" + "")) + "hello") {
            case "\n\t":
                v11 := GenS1{F0: -2147483648, F1: v2}
                _ = v11
            }
            _ = fi10
        }
    default:
        var v12 int32 = 0
        _ = v12
    }
    v3 += (^(v2))
    v6 = v6
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (("\n\t" + "a") + (("" + "\n\t") + "a")) {
        case "\n\t":
            if true {
                var v14 float32 = float32(float64(0.0))
                _ = v14
            }
        case "":
            fmt.Println(v6, fi13)
        case "a":
            v9 = (v9 ^ (v9 + (v9 ^ (v9 + (-(v9))))))
        default:
            switch (fi13 & (v6 & (fi13 >> ((65535) & 63)))) {
            case 0:
                v2 = v2
            case 678:
                switch (fi13 << ((404) & 63)) {
                case 293:
                    fmt.Println(fi13)
                default:
                    if (!(79 <= 65535)) {
                        v15 := []int{-1, 32768, genFn0((v9 + -504))}
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v7 = v7
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        v2 -= (v3 % (((^(v3))) | 1))
                        v6 = 32768
                        if (297 == 590) {
                            fmt.Println(v2, v6)
                            var v17 float32 = ((float32(float64(0.1)) + 0.1) + 3.14159)
                            switch "" {
                            case "foo bar":
                                v18 := GenS1{F0: 42, F1: (v2 - 2)}
                                _ = v18
                            case "a":
                                if false {
                                    v19 := GenS1{F0: 888, F1: v2}
                                    _ = v19
                                } else {
                                    fmt.Println(v9, v3)
                                    if true {
                                        fmt.Println(v6)
                                    }
                                }
                            case "hello":
                                var v20 float32 = float32((1e10 + (1e10 * -2.25)))
                                _ = v20
                            default:
                                v21 := GenS1{F0: 2147483647, F1: 981}
                                _ = v21
                            }
                            _ = v17
                        } else {
                            v6 = (fi13 - v6)
                            switch "\n\t" {
                            case "\n\t":
                                v22 := v1[(v2 / ((42) | 1))]
                                _ = v22
                            default:
                                switch "hello" {
                                case "foo bar":
                                    v7 += v7
                                case "\n\t":
                                    v1[4294967295] = v6
                                }
                            }
                            v2 = 0
                        }
                    }
                }
            case 42:
                var v23 uint = 1
                _ = v23
            }
        }
        fmt.Println(v3)
        v24 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807}
        _ = fi13
        _ = v24
    }
    _ = len(v1)
    v25 := GenS3{F0: v2, F1: 42}
    fmt.Println(v6, v7, v3)
    var v26 uint16 = 42
    if ((v26 * (v26 - v26)) != v26) {
        var v27 int64 = 9223372036854775807
        var v28 int64 = 42
        _ = v27
        _ = v28
    }
    var v29 int8 = 100
    if (true && false) {
        if false {
            for fi30 := 0; fi30 < 6; fi30++ {
                v7 += v7
                for fi31 := 0; fi31 < 12; fi31++ {
                    fmt.Println(v2, fi30)
                    _ = fi31
                }
                _ = fi30
            }
        } else {
            switch genFn0(v9) {
            case -9223372036854775808:
                v9 = v9
            case 32768:
                var v32 uint32 = (v2 << ((82) & 63))
                _ = v32
            case 9223372036854775807:
                for fi33 := 0; fi33 < 4; fi33++ {
                    var v34 uint64 = uint64(v7)
                    v7 += (^(v7))
                    var v35 int = (fi33 - 32767)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
        }
        v36 := GenS3{F0: 2, F1: 42}
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
