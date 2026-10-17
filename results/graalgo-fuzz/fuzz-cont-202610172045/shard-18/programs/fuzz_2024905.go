package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 42
    var v3 float64 = 1e10
    var v4 int8 = -128
    v5 := []string{("\n\t" + ("hello" + "a")), ("a" + "a"), ("hello" + "\n\t"), "\n\t"}
    var v6 int32 = 669
    var v7 int = (v2 ^ v2)
    var v8 float64 = (v3 * v3)
    _ = len(v1)
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    var v9 bool = (float32(float64(0.0)) >= (1.5 + 3.14159))
    if (v6 > int32(v7)) {
        v4 -= (-(v4))
    } else {
        _ = len(v1)
    }
    v8 = 1.5
    v2 = -9223372036854775808
    var v10 int64 = 9223372036854775807
    v11 := GenS1{F0: 127, F1: 65535, F2: (^(v7))}
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 610, F1: 1, F2: -15}
        var v14 int16 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    switch v7 {
    case 9223372036854775807:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v8 += float64(float32(-2.25))
    fmt.Println(v3, v2, v8)
    if false {
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                if (384 < 42) {
                    v17 := GenS3{F0: 577, F1: (v6 + -656)}
                    if (v9 && (2 > 706)) {
                        fmt.Println(fi15, v7, v3)
                        var v18 bool = v9
                        _ = v18
                    } else {
                        var v19 uint32 = 175
                        switch 4294967295 {
                        case 32768:
                            var v20 uint64 = 0
                            _ = v20
                        case 128:
                            switch (v2 ^ int(v4)) {
                            case -2147483648:
                                v7 = 0
                            default:
                                if ((("a" + "\n\t") + "a") >= (("foo bar" + "\n\t") + ("foo bar" + "foo bar"))) {
                                    v2 = fi15
                                    if len(v5) > 0 {
                                        v5[0] = ("" + ("hello" + "foo bar"))
                                    }
                                }
                            }
                        }
                        v9 = true
                        _ = v19
                    }
                    _ = v17
                }
                var v21 uint32 = 42
                v3 -= v8
                _ = fi16
                _ = v21
            }
            v22 := GenS4{F0: v9, F1: 127, F2: 1}
            _ = fi15
            _ = v22
        }
    }
    var v23 float64 = float64((3.14159 - (1e10 + 1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
