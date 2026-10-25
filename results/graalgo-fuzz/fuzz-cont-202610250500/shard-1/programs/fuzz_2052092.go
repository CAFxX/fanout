package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 uint16 }

func main() {
    var v1 int8 = 127
    var v2 uint = 1
    v3 := make(map[int]string)
    var v4 int = 0
    var v5 float64 = (0.1 - 3.14159)
    var v6 uint8 = 255
    switch ("hello" + "hello") {
    case "":
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint32 = 81
            v4 = (v4 - -1)
            _ = fi7
            _ = v8
        }
    case "hello":
        var v9 bool = ((v2 == v2) || (false || false))
        _ = v9
    }
    if (1.5 <= (3.14159 / (3.14159 * (1.5 * (float32(v5) * float32(v5)))))) {
        fmt.Println(v2, v5, v4)
        v4 += (v4 % (((v4 - v4)) | 1))
        v2 += 18446744073709551615
    }
    if (-268 != 1) {
        v2 -= v2
        var v10 uint8 = (v6 ^ (^(v6)))
        v11 := GenS1{F0: v1, F1: (v2 << (((v2 >> (((^(v2))) & 63))) & 63))}
        _ = v10
        _ = v11
    } else {
        var v12 uint16 = 2
        v13 := GenS1{F0: v1, F1: 818}
        if (!true) {
            v4 += 32768
        } else {
            switch "a" {
            case "\n\t":
                var v14 uint8 = 172
                _ = v14
            case "a":
                switch "hello" {
                case "hello":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v17 := GenS1{F0: v1, F1: (v2 ^ 0)}
                            v18 := GenS1{F0: (v1 | (v1 ^ int8(fi15))), F1: v2}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        if (0.0 >= 0.1) {
                            var v19 int64 = 42
                            v20 := GenS1{F0: (^(v1)), F1: (v2 >> ((1) & 63))}
                            _ = v19
                            _ = v20
                        } else {
                            if (!((!false) && true)) {
                                var v21 string = ("hello" + ("foo bar" + ("" + ("foo bar" + "a"))))
                                _ = v21
                            }
                            v12 = (v12 | 566)
                        }
                        var v22 string = ("\n\t" + "a")
                        _ = fi15
                        _ = v22
                    }
                }
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                v5 -= (1.5 * (v5 / 0.1))
                _ = fi23
            }
        }
        _ = v12
        _ = v13
    }
    if ((-1 > 0) || ((1.5 != ((((-2.25 / 1e10) / ((-2.25 - 0.1) + (0.0 + 3.14159))) / 1.5) - ((0.0 * float32(0.1)) * 0.1))) && true)) {
        switch "" {
        case "foo bar":
            v24 := []int{(v4 * 32768), 9223372036854775807}
            _ = v24
        case "a":
            var v25 uint32 = 65535
            _ = v25
        default:
            if ((float32(1e10) - 1e10) <= 0.0) {
                v1 -= 12
                var v26 uint64 = 4294967295
                if false {
                    fmt.Println(v1)
                    var v27 bool = (757 <= 0)
                    var v28 uint32 = 2147483647
                    _ = v27
                    _ = v28
                }
                _ = v26
            }
        }
        v6 = (v6 ^ v6)
        v6 += v6
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS2{F0: "foo bar", F1: "", F2: 255}
            switch fi29 {
            case 32768:
                v31 := []int{4294967295, v4, fi29, (fi29 + (^(v4)))}
                _ = v31
            default:
                if (!true) {
                    v32 := GenS1{F0: v1, F1: v2}
                    v4 = (^(v4))
                    _ = v32
                } else {
                    fmt.Println(v5, v4, v5)
                }
            }
            _ = fi29
            _ = v30
        }
        v6 += v6
        v5 = v5
    }
    var v33 int = 42
    delete(v3, 785)
    v5 = 1.5
    if false {
        switch v4 {
        case -466:
            v2 -= (v2 ^ 0)
        case -9223372036854775808:
            var v34 uint64 = 18446744073709551615
            _ = v34
        }
    } else {
        v4 = v4
        fmt.Println(v4)
        v33 -= (-(v33))
    }
    v35 := GenS1{F0: (v1 % (((v1 / (((v1 & v1)) | 1))) | 1)), F1: v2}
    switch -2147483648 {
    case 9223372036854775807:
        v5 = ((v5 * 1e10) + 1.5)
    case 1:
        for fi36 := 0; fi36 < 14; fi36++ {
            _ = len(v3)
            v4 -= 32767
            _ = fi36
        }
    default:
        var v37 float32 = (1e10 * -2.25)
        _ = v37
    }
    var v38 string = (("a" + "") + "a")
    if true {
        var v39 bool = true
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
