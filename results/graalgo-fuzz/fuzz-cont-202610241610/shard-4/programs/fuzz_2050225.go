package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 bool) uint {
    switch -9223372036854775808 {
    case -1:
        var v6 string = "hello"
        _ = v6
    case 32768:
        var v7 int = -9223372036854775808
        _ = v7
    }
    if (4294967295 > 18446744073709551615) {
        p0 = (p0 || (!false))
        fmt.Println(p0)
        v8 := []int{32767, 588}
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: ("a" + "hello")}
        for fi11 := 0; fi11 < 14; fi11++ {
            p0 = false
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    v5 := make(map[string]string)
    for fi12 := 0; fi12 < 10; fi12++ {
        if (255 == 865) {
            v13 := GenS2{F0: 73}
            _ = v13
        }
        switch fi12 {
        case -2147483648:
            switch ((("\n\t" + ("foo bar" + "hello")) + (((("foo bar" + "\n\t") + "\n\t") + ("" + "foo bar")) + "")) + (("" + ((("hello" + "\n\t") + "hello") + "\n\t")) + "")) {
            case "\n\t":
                delete(v5, "")
            default:
                for fi14 := 0; fi14 < 17; fi14++ {
                    v2 = v2
                    v4 += (((-2.25 - float32(1e10)) - float32((1e10 + 3.14159))) + v4)
                    _ = len(v3)
                    _ = fi14
                }
            }
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v4 += v4
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int = fi16
            v18 := GenS3{F0: 1, F1: 2}
            v17 = (v17 * fi15)
            var v19 uint64 = 42
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    if true {
        var v20 string = "foo bar"
        if false {
            switch v20 {
            case "a":
                if false {
                    v21 := GenS3{F0: 0, F1: 227}
                    _ = v21
                } else {
                    if true {
                        v22 := []int{-1, 32768, -950, -9223372036854775808}
                        if (2 >= 140) {
                            for fi23 := 0; fi23 < 2; fi23++ {
                                var v24 string = ("\n\t" + v20)
                                _ = fi23
                                _ = v24
                            }
                            var v25 int = 2147483647
                            _ = v25
                        }
                        _ = v22
                    } else {
                        var v26 int8 = -128
                        v27 := GenS4{F0: 0, F1: -2147483648, F2: v26}
                        var v28 int16 = 1000
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    if (((1.5 > ((3.14159 - 0.1) * (-2.25 - (float64((v4 / (v4 * -2.25))) - 1e10)))) || (!(!true))) && (false || (32767 != -657))) {
                        fmt.Println(v4)
                        var v29 int64 = (v2 + (v2 * (^(v2))))
                        _ = v29
                    }
                }
            case "\n\t":
                var v30 int64 = (v2 % ((v2) | 1))
                _ = v30
            default:
                var v31 uint32 = 42
                _ = v31
            }
            fmt.Println(v20, v2, v20)
        }
        var v32 int = 4294967295
        _ = v20
        _ = v32
    } else {
        v2 = v2
        v4 = -2.25
        fmt.Println(v2, v4)
    }
    v2 -= v2
    var v33 uint64 = 1
    var v34 bool = (v4 <= v4)
    v34 = (((v4 / v4) != 0.0) || (v33 < v33))
    v35 := GenS3{F0: 825, F1: 32767}
    v34 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
