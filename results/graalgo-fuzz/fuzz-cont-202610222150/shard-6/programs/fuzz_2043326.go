package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    var v1 string = "a"
    var v2 int32 = -2147483648
    v3 := make(map[string]uint32)
    var v4 int64 = -25
    v5 := make(map[int]uint32)
    v1 = ((v1 + ((((v1 + v1) + v1) + ("hello" + ((("a" + "\n\t") + (("" + (v1 + ("hello" + "hello"))) + ("" + "\n\t"))) + (v1 + v1)))) + v1)) + (((v1 + "a") + (((v1 + "foo bar") + ((v1 + ((v1 + "hello") + "\n\t")) + (("\n\t" + v1) + v1))) + v1)) + ("\n\t" + v1)))
    fmt.Println(v2)
    v6 := GenS1{F0: 255, F1: 522, F2: 36}
    v4 += 0
    _ = len(v5)
    switch 496 {
    case -360:
        switch v1 {
        case "a":
            if true {
                for fi7 := 0; fi7 < 17; fi7++ {
                    if true {
                        switch fi7 {
                        case 9223372036854775807:
                            if true {
                                for fi8 := 0; fi8 < 19; fi8++ {
                                    delete(v5, fi8)
                                    _ = fi8
                                }
                                var v9 int32 = v2
                                _ = v9
                            }
                        case -249:
                            v10 := GenS2{F0: (v4 + (v4 - v4))}
                            _ = v10
                        case -9223372036854775808:
                            var v11 int16 = -541
                            _ = v11
                        }
                        var v12 bool = (!(2147483647 > 1))
                        v13 := v5[-2147483648]
                        _ = v12
                        _ = v13
                    }
                    var v14 int = (fi7 + (fi7 % ((fi7) | 1)))
                    if (false && (((v4 ^ int64(v14)) > (-(v4))) || (!true))) {
                        if ((^(v4)) >= v4) {
                            v15 := []int{fi7, v14, fi7, (^(v14))}
                            _ = v15
                        }
                        v4 = (v4 | (v4 - (-(v4))))
                        v16 := v5[fi7]
                        _ = v16
                    } else {
                        v4 = v4
                    }
                    _ = fi7
                    _ = v14
                }
                if (((v2 | v2) > 1) || (1 > 4294967295)) {
                    var v17 uint64 = 42
                    if true {
                        v18 := GenS3{F0: ("foo bar" + v1)}
                        v17 += v17
                        v1 = (v1 + v1)
                        _ = v18
                    } else {
                        if (1 != -964) {
                            v2 -= -238
                        }
                    }
                    _ = v17
                }
            }
        case "foo bar":
            v19 := []int{}
            _ = v19
        case "hello":
            fmt.Println(v2, v2)
        default:
            switch (((v1 + "hello") + v1) + "\n\t") {
            case "\n\t":
                v4 = (v4 << ((674) & 63))
            case "":
                var v20 uint8 = 2
                _ = v20
            case "a":
                v4 = (v4 & v4)
            }
        }
    case 128:
        v2 += (^(v2))
    }
    v21 := v5[128]
    if (2147483647 < -2147483648) {
        for fi22 := 0; fi22 < 18; fi22++ {
            fmt.Println(v1, v4)
            var v23 float32 = 1.5
            var v24 float32 = ((v23 + v23) * float32((1e10 - (1.5 / (((3.14159 / 1.5) / 1.0) / ((3.14159 - ((float64(v23) / 1e10) / (1.5 - (((-2.25 - (float64(float32(float64(v23))) - (1.5 / 3.14159))) * (1e10 + -2.25)) - -2.25)))) * -2.25))))))
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS4{F0: v2, F1: 343}
        v26 := GenS2{F0: v4}
        _ = v25
        _ = v26
    } else {
        fmt.Println(v2, v21, v21)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
