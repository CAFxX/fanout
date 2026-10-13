package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []int{127, -208, -593}
    var v2 int = 513
    var v3 int64 = -1
    v4 := make(map[uint32]int)
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 15; fi6++ {
        v2 = (fi6 / ((v2) | 1))
        if len(v1) > 0 {
            v1[0] = fi6
        }
        v7 := GenS1{F0: 255}
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    v8 := GenS1{F0: 1}
    v9 := GenS2{F0: (v3 | (v3 / (((v3 - (v3 << ((42) & 63)))) | 1))), F1: 0}
    switch ("\n\t" + ((("hello" + ("a" + ("" + "foo bar"))) + (("hello" + "hello") + "\n\t")) + "foo bar")) {
    case "a":
        fmt.Println(v2, v5)
    case "hello":
        v3 += v3
    }
    v10 := GenS2{F0: 0, F1: 42}
    if len(v1) > 0 {
        v1[0] = v2
    }
    v3 += v3
    switch (v2 >> ((0) & 63)) {
    case -502:
        switch "" {
        case "":
            v5 -= 18446744073709551615
        case "hello":
            if true {
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS3{F0: int32(v3)}
                    var v13 uint64 = (v5 % (((v5 - (v5 | v5))) | 1))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v5 = v5
            } else {
                var v14 int32 = -2147483648
                if (-2.25 <= (0.0 - 1e10)) {
                    if ((float32((((0.0 / 1.5) + 0.1) / 1e10)) + 0.1) != (0.0 + ((-2.25 * -2.25) + 3.14159))) {
                        v15 := GenS4{F0: (v3 & 1)}
                        switch v2 {
                        case 909:
                            v14 += 1
                        case 174:
                            var v16 int = v2
                            _ = v16
                        case 4294967295:
                            v5 = (^(v5))
                        }
                        _ = v15
                    }
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v5 = 0
                        _ = fi17
                    }
                    _ = len(v1)
                }
                _ = v14
            }
        }
    case -2147483648:
        var v18 int16 = 614
        _ = v18
    case -9:
        v19 := GenS3{F0: -1}
        _ = v19
    default:
        var v20 uint8 = 153
        _ = v20
    }
    v2 = v2
    v2 = (v2 & v2)
    if false {
        if (int16(v5) == -324) {
            v2 = 263
            _ = len(v4)
            v5 -= (v5 - 222)
        }
        if false {
            for fi21 := 0; fi21 < 17; fi21++ {
                v2 += 658
                v22 := v4[2147483648]
                _ = fi21
                _ = v22
            }
            switch "hello" {
            case "a":
                if (!(0 == 610)) {
                    v5 = (^(v5))
                    var v23 uint = 0
                    _ = v23
                }
            case "\n\t":
                v1 = append(v1, 128)
            case "":
                fmt.Println(v2, v5, v2)
            default:
                if true {
                    var v24 uint16 = 408
                    v3 = (v3 | (v3 & (v3 | v3)))
                    v25 := GenS4{F0: 204}
                    _ = v24
                    _ = v25
                }
            }
            v3 += (v3 % ((9223372036854775807) | 1))
        }
        v26 := GenS5{F0: 1, F1: -77}
        _ = v26
    } else {
        switch "a" {
        case "foo bar":
            var v27 int8 = 1
            _ = v27
        case "a":
            v28 := GenS2{F0: v3, F1: -622}
            _ = v28
        case "\n\t":
            v5 = 2
        default:
            var v29 float64 = (0.1 * -2.25)
            _ = v29
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
