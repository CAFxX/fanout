package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 int8 }
type GenS3 struct { F0 int32; F1 string }

func main() {
    v1 := []int{4294967295, -894}
    var v2 int16 = -1
    v3 := make(map[string]string)
    var v4 float64 = (float64(0.0) / 1.5)
    var v5 uint32 = 42
    var v6 uint32 = v5
    for fi7 := 0; fi7 < 7; fi7++ {
        if false {
            v8 := GenS1{F0: 255}
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint64 = 534
                if true {
                    if true {
                        v11 := GenS1{F0: 999}
                        var v12 bool = true
                        v10 += (v10 - (^(v10)))
                        _ = v11
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 0}
                        v6 = 4294967295
                        if ((42 > 42) || (-46 < 2147483647)) {
                            if false {
                                v6 -= v6
                                if len(v1) > 0 {
                                    v1[0] = (fi7 * fi7)
                                }
                                v14 := GenS2{F0: 32767, F1: 2, F2: 42}
                                switch (fi7 / (((fi9 - (^(fi9)))) | 1)) {
                                case 127:
                                    v15 := GenS2{F0: (v2 << ((65535) & 63)), F1: 107, F2: 47}
                                    _ = v15
                                case -2147483648:
                                    if (65 >= 42) {
                                        v6 = (^(v5))
                                        v6 -= (v5 - 2147483647)
                                    }
                                }
                                _ = v14
                            }
                        } else {
                            v16 := GenS1{F0: 2}
                            _ = v16
                        }
                        _ = v13
                    }
                }
                v5 = (v5 % (((^(v5))) | 1))
                _ = fi9
                _ = v10
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                if true {
                    v6 = (^(v5))
                    v18 := GenS1{F0: 396}
                    if false {
                        _ = len(v3)
                        v4 = v4
                    }
                    _ = v18
                }
                if (!true) {
                    var v19 uint = 1
                    var v20 int64 = 1
                    _ = v19
                    _ = v20
                }
                var v21 float64 = (float64(3.14159) / 1.5)
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        switch "" {
                        case "a":
                            var v22 int32 = -533
                            _ = v22
                        case "hello":
                            v6 = (v6 >> ((18446744073709551615) & 63))
                        case "":
                            var v23 uint8 = 128
                            _ = v23
                        }
                    }
                    v21 = ((float64(-2.25) + v21) - (float64(0.0) - v4))
                    v24 := GenS3{F0: -2, F1: ""}
                    _ = v24
                } else {
                    fmt.Println(v2)
                    var v25 float64 = (((0.1 + 3.14159) / 0.1) * (((-2.25 - v4) * v21) / ((v21 - float64(0.1)) - v21)))
                    _ = v25
                }
                var v26 uint64 = 0
                _ = fi17
                _ = v21
                _ = v26
            }
            _ = v8
        }
        var v27 uint = 512
        _ = fi7
        _ = v27
    }
    v2 = v2
    var v28 float64 = -2.25
    fmt.Println(v4)
    v29 := v3[""]
    v30 := GenS1{F0: 692}
    v29 = v29
    v4 -= 1.5
    v3[v29] = ("a" + "\n\t")
    v2 -= v2
    v31 := GenS1{F0: 255}
    if (42 == 2) {
        v2 += 230
        v5 = (v5 | v5)
    } else {
        v28 += float64(float32(3.14159))
        var v32 uint16 = 1
        _ = v32
    }
    var v33 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
