package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }

func main() {
    var v1 int = 32768
    var v2 uint32 = 0
    var v3 float32 = ((0.1 * 3.14159) + (((-2.25 - 1e10) * (float32(-2.25) * -2.25)) - float32(((float64(((-2.25 + 0.0) * 0.0)) + 0.0) * ((1.5 * float64(-2.25)) * (1.5 - (1.5 + (1.5 - (1e10 / (1.5 / float64(3.14159)))))))))))
    if ((false || (("\n\t" < "foo bar") || true)) || true) {
        switch -979 {
        case 128:
            fmt.Println(v2)
        case 2147483647:
            var v4 int = (^(v1))
            _ = v4
        case 212:
            v1 -= v1
        }
    } else {
        for fi5 := 0; fi5 < 13; fi5++ {
            if (-892 == 609) {
                v3 = (v3 / v3)
                v6 := []int{}
                v3 += ((v3 + float32((3.14159 * float64(0.0)))) - v3)
                _ = v6
            } else {
                v3 -= ((float32(float64(v3)) * v3) / v3)
            }
            var v7 uint = 2
            _ = fi5
            _ = v7
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            v1 = (v1 >> ((uint(v2)) & 63))
            v9 := []int{(fi8 / ((v1) | 1)), fi8, (fi8 * v1), (fi8 * fi8)}
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 42
        _ = v10
    }
    var v11 int8 = 1
    if ((v1 - v1) > (v1 + v1)) {
        v11 -= (^(v11))
        for fi12 := 0; fi12 < 19; fi12++ {
            switch (^(v1)) {
            case 0:
                v13 := []int{494}
                _ = v13
            case 42:
                v14 := []int{}
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint16 = 32767
                v2 = 0
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
        v2 = (^(v2))
    } else {
        v2 += (^(v2))
        if (!(((65535 < 255) || (9223372036854775807 > 9223372036854775807)) && (42 != 18446744073709551615))) {
            if false {
                v17 := []int{v1, 4294967295, v1, v1}
                _ = v17
            }
        }
    }
    if true {
        v18 := []int{(v1 / ((v1) | 1))}
        _ = v18
    } else {
        v3 = (v3 / (v3 - (float32((((1.5 / 1.0) + 0.0) / (1.5 * 1e10))) * -2.25)))
        v3 = v3
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v3, v3, v1)
        if true {
            var v20 uint8 = 180
            for fi21 := 0; fi21 < 10; fi21++ {
                switch 9223372036854775807 {
                case -1:
                    v22 := GenS1{F0: 18446744073709551615}
                    _ = v22
                }
                switch (^(v1)) {
                case -823:
                    v23 := []int{}
                    _ = v23
                case -750:
                    switch fi21 {
                    case 0:
                        fmt.Println(fi21)
                    case 32767:
                        if false {
                            v2 += 65535
                            v24 := []int{9223372036854775807, 128, fi19, 32767}
                            if (-1 != 1000) {
                                _ = len(v24)
                            } else {
                                var v25 uint = 255
                                v2 = uint32(v25)
                                _ = v25
                            }
                            _ = v24
                        }
                    case 4294967295:
                        var v26 int8 = (v11 * 65)
                        _ = v26
                    }
                case 9223372036854775807:
                    var v27 int64 = -1
                    _ = v27
                }
                _ = fi21
            }
            v28 := GenS2{F0: 18446744073709551615, F1: 0, F2: 18446744073709551615}
            _ = v20
            _ = v28
        }
        v2 -= (v2 / (((v2 << ((771) & 63))) | 1))
        _ = fi19
    }
    v29 := []int{}
    var v30 float64 = float64(float32(float64((v3 / v3))))
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v29 = append(v29, (v1 << ((255) & 63)))
    v3 = v3
    fmt.Println(v3, v30)
    fmt.Println(v1, v11)
    v3 = (v3 / v3)
    var v31 float64 = ((3.14159 / (v30 / 3.14159)) - (float64((v3 * 0.1)) + (((v30 - 1e10) * v30) + v30)))
    fmt.Println(v11)
    if len(v29) > 0 {
        v29[0] = (v1 + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
