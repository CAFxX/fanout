package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 string; F2 uint }

func main() {
    var v1 int8 = 127
    var v2 int32 = 2147483647
    var v3 float64 = 1.5
    v4 := []uint32{2, 757}
    v5 := []string{"foo bar", "a"}
    var v6 float64 = (v3 + (float64(((0.1 + 3.14159) * float32((v3 * (3.14159 * 3.14159))))) - ((v3 - float64((1e10 / -2.25))) * (((v3 / ((0.1 * float64(1.5)) * v3)) - ((float64(1.5) * v3) + (v3 + ((1e10 / v3) + 0.1)))) / (0.0 - -2.25)))))
    fmt.Println(v6, v3, v3)
    if (true && false) {
        v2 += v2
    } else {
        v3 = ((v6 - ((float64((0.1 * 1.5)) + v6) / 1.5)) / ((1.5 - (v6 + 3.14159)) * v3))
    }
    switch "a" {
    case "":
        fmt.Println(v6, v3)
    case "foo bar":
        v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: 1}
        _ = v7
    case "hello":
        v3 = v3
    }
    var v8 int = int(v2)
    var v9 int64 = -5
    if (false && false) {
        switch "\n\t" {
        case "\n\t":
            switch "" {
            case "":
                v10 := GenS2{F0: 268}
                _ = v10
            case "foo bar":
                if (463 >= 2) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v8)
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v3 -= (float64(0.1) / 1e10)
                        if ((!(103 == (v2 ^ -106))) || (255 <= 246)) {
                            switch "foo bar" {
                            case "\n\t":
                                switch ("\n\t" + ("hello" + ("" + ((("\n\t" + ("hello" + "")) + "\n\t") + "foo bar")))) {
                                case "a":
                                    for fi12 := 0; fi12 < 7; fi12++ {
                                        v13 := GenS1{F0: 1, F1: 128, F2: 0}
                                        _ = fi12
                                        _ = v13
                                    }
                                }
                            }
                            var v14 int32 = 2147483647
                            _ = v14
                        } else {
                            fmt.Println(v8, v9, v6)
                            var v15 uint16 = 32767
                            fmt.Println(v1)
                            _ = v15
                        }
                        fmt.Println(v2)
                        _ = fi11
                    }
                } else {
                    v16 := GenS2{F0: 845}
                    _ = v16
                }
            case "\n\t":
                var v17 uint16 = 0
                _ = v17
            default:
                var v18 int8 = v1
                _ = v18
            }
        }
        v6 += float64(0.0)
    } else {
        v19 := GenS2{F0: 2}
        _ = v19
    }
    if (245 >= 2147483647) {
        var v20 uint16 = 130
        _ = v20
    } else {
        var v21 int8 = (^(v1))
        _ = v21
    }
    var v22 int8 = v1
    v23 := GenS3{F0: 65535, F1: -1, F2: v2}
    var v24 int32 = (v2 % (((v2 % (((^(v2))) | 1))) | 1))
    if false {
        var v25 int = 2147483647
        var v26 uint = 2
        if true {
            var v27 int8 = (v22 - v22)
            _ = v27
        } else {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint16 = 2
                _ = fi28
                _ = v29
            }
            var v30 uint = v26
            switch (-(v25)) {
            case -577:
                var v31 uint64 = 507
                _ = v31
            case -2147483648:
                for fi32 := 0; fi32 < 12; fi32++ {
                    if (true && (296 <= 117)) {
                        v30 += 316
                        v9 = (v9 / ((-9223372036854775808) | 1))
                    } else {
                        _ = len(v4)
                    }
                    for fi33 := 0; fi33 < 4; fi33++ {
                        fmt.Println(v9, v3)
                        var v34 int32 = int32(v26)
                        var v35 uint64 = 42
                        v26 = 65535
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS3{F0: (v26 >> (((v30 ^ v26)) & 63)), F1: (v9 & -1), F2: v24}
                    _ = fi32
                    _ = v36
                }
            case -603:
                v3 -= (1e10 + v3)
            default:
                v9 = 0
            }
            _ = v30
        }
        _ = v25
        _ = v26
    }
    v37 := GenS4{F0: 0}
    var v38 float32 = 1e10
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if len(v4) > 0 {
        v4[0] = 155
    }
    v39 := GenS5{F0: (v24 - v24), F1: "foo bar", F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
