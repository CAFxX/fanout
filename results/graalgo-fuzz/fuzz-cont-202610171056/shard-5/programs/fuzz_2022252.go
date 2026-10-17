package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
type GenS5 struct { F0 int; F1 uint; F2 int16 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    var v3 float32 = 0.1
    v3 = ((float32((float64(v3) + 1.5)) + (v3 - v3)) * v3)
    if true {
        if true {
            v4 := GenS1{F0: false, F1: 2}
            v1 = (v1 + (v1 + v1))
            _ = v4
        }
    } else {
        switch 4294967295 {
        case -69:
            v5 := v2[4294967295]
            _ = v5
        case 42:
            v6 := GenS1{F0: (false && true), F1: 65535}
            _ = v6
        }
        var v7 bool = (v1 == 88)
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = 1
        _ = fi8
        _ = v9
    }
    if true {
        if ((2 < 0) || false) {
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v3 += (0.0 - float32(-2.25))
                    var v12 int16 = 32767
                    _ = fi11
                    _ = v12
                }
                var v13 int32 = 2147483647
                _ = fi10
                _ = v13
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch fi15 {
                    case 4294967295:
                        v3 += 3.14159
                    case 32768:
                        switch (-(fi14)) {
                        case 922:
                            fmt.Println(v3, v3)
                        default:
                            v16 := GenS2{F0: -653, F1: 75}
                            _ = v16
                        }
                    case 128:
                        v1 = -114
                        v17 := GenS3{F0: ""}
                        var v18 float64 = -2.25
                        _ = v17
                        _ = v18
                    default:
                        v19 := []int{fi14, fi14, 32767}
                        _ = v19
                    }
                    var v20 int64 = -9223372036854775808
                    var v21 int16 = -1
                    v22 := GenS2{F0: (v20 | -490), F1: 1}
                    fmt.Println(v21, v21)
                    _ = fi15
                    _ = v20
                    _ = v21
                    _ = v22
                }
                var v23 uint64 = 518
                var v24 uint16 = 63
                _ = fi14
                _ = v23
                _ = v24
            }
            v3 -= (1e10 - v3)
        } else {
            v1 = v1
            var v25 string = "foo bar"
            _ = v25
        }
    } else {
        var v26 string = ("\n\t" + ("\n\t" + ""))
        delete(v2, 127)
        _ = v26
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v27 uint16 = 32767
        _ = v27
    }
    if true {
        if false {
            switch 42 {
            case -9223372036854775808:
                var v28 float32 = v3
                _ = v28
            case 1:
                var v29 uint64 = 18446744073709551615
                _ = v29
            }
            v30 := GenS2{F0: 1, F1: 153}
            switch 4294967295 {
            case 1:
                v31 := GenS2{F0: 9223372036854775807, F1: 79}
                _ = v31
            case 33:
                v2[42] = "foo bar"
            case -2147483648:
                v32 := GenS4{F0: 86, F1: 42, F2: (2147483648 <= 2)}
                _ = v32
            }
            _ = v30
        }
        fmt.Println(v3, v3)
        fmt.Println(v1, v3)
    } else {
        v33 := GenS5{F0: 0, F1: 921, F2: 255}
        _ = v33
    }
    var v34 bool = (!(v1 <= (v1 | (v1 % ((v1) | 1)))))
    v35 := []int{int(v1)}
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 9; fi37++ {
            if (2 <= 0) {
                v1 += v1
                switch ("foo bar" + (("a" + ("" + "hello")) + ("" + ("\n\t" + ("" + (("\n\t" + (("hello" + "\n\t") + (("" + ("foo bar" + "a")) + "a"))) + ("foo bar" + (("" + "foo bar") + ("\n\t" + ""))))))))) {
                case "\n\t":
                    v1 = v1
                case "hello":
                    fmt.Println(v34)
                case "foo bar":
                    var v38 float64 = -2.25
                    _ = v38
                }
                var v39 int16 = 32767
                _ = v39
            }
            _ = fi37
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            v34 = v34
            var v41 int8 = v1
            _ = fi40
            _ = v41
        }
        v3 = float32(((1.5 + float64((float32(-2.25) / v3))) + (((0.1 + 1e10) + (1.5 - (float64(v3) * (1.5 + 3.14159)))) - -2.25)))
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
