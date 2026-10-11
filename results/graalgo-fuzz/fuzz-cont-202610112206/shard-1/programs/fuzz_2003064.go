package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0() uint {
    var v7 string = "hello"
    var v8 uint32 = 2147483647
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := []uint32{222, 4294967295}
    var v4 float64 = ((0.1 / 1.5) * (3.14159 - (0.1 - 1.5)))
    var v5 uint16 = 32767
    var v6 bool = (2 > 127)
    v9 := GenS1{F0: 0, F1: ((((0.1 * ((1.5 - (1e10 / 3.14159)) * 0.1)) + 0.1) - 1e10) >= (1e10 - (((float32((1.5 * (float64((float32((1e10 / 1.5)) - 1e10)) / v4))) + 0.1) / 3.14159) + -2.25)))}
    v1 = v1
    if false {
        if false {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int64 = -105
                var v12 int = (fi10 & fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
            v4 = ((float64(1.5) - v4) - (((v4 * float64(3.14159)) / float64((1e10 - 3.14159))) - (0.0 * ((((((v4 + v4) - v4) + (v4 - ((((float64(3.14159) / v4) + (float64(3.14159) - v4)) * -2.25) * 0.0))) - v4) / (0.0 / v4)) - 1e10))))
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int32 = 42
                if (v14 >= (v14 ^ (v14 - (v14 ^ 0)))) {
                    switch v1 {
                    case "\n\t":
                        v15 := GenS1{F0: 2, F1: v6}
                        _ = v15
                    default:
                        v5 -= 65535
                    }
                    delete(v2, 2)
                }
                fmt.Println(v14, v14, v1)
                _ = fi13
                _ = v14
            }
        }
    } else {
        var v16 uint32 = 1
        _ = v16
    }
    v4 -= 0.1
    v17 := GenS1{F0: 65535, F1: (0 <= 42)}
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            _ = len(v3)
            v1 = (v1 + (v1 + (v1 + "a")))
            v19 := GenS1{F0: 759, F1: false}
            _ = v19
        } else {
            switch (fi18 & fi18) {
            case -9223372036854775808:
                v20 := GenS1{F0: genFn0(), F1: false}
                _ = v20
            case 42:
                fmt.Println(v1, v5)
            default:
                v6 = (v6 && (-991 == -1))
            }
            fmt.Println(v5)
        }
        v5 = (v5 >> ((669) & 63))
        for fi21 := 0; fi21 < 15; fi21++ {
            v3 = append(v3, 42)
            var v22 int16 = -613
            if v6 {
                var v23 uint = 936
                v24 := GenS2{F0: ((!v6) || v6), F1: 1, F2: ((fi21 ^ fi18) <= 214)}
                _ = v23
                _ = v24
            } else {
                if (true && ((fi21 >> ((362) & 63)) > (fi21 >> ((genFn0()) & 63)))) {
                    switch "foo bar" {
                    case "foo bar":
                        switch fi21 {
                        case -221:
                            v6 = (2147483647 <= 1)
                        case 0:
                            v6 = v6
                        case 812:
                            if (int64(v22) != -1) {
                                if len(v3) > 0 {
                                    v3[0] = 963
                                }
                                v25 := GenS3{F0: 2}
                                _ = v25
                            }
                        }
                        var v26 string = "foo bar"
                        _ = v26
                    case "hello":
                        v4 = 1e10
                    }
                }
            }
            _ = fi21
            _ = v22
        }
        _ = fi18
    }
    var v27 bool = v6
    v28 := v2[987]
    if len(v3) > 0 {
        v3[0] = v28
    }
    v3 = append(v3, (v28 ^ v28))
    if true {
        v6 = ((^(v5)) < 42)
    } else {
        fmt.Println(v1, v27, v5)
        if v6 {
            v1 = ((v1 + ("foo bar" + "a")) + "foo bar")
            v29 := GenS3{F0: 333}
            v30 := v2[(^(v28))]
            _ = v29
            _ = v30
        } else {
            var v31 uint64 = 4294967295
            _ = v31
        }
    }
    if v6 {
        v5 = v5
        v1 = (((v1 + ("hello" + ((v1 + (v1 + v1)) + v1))) + "a") + v1)
        var v32 uint16 = v5
        _ = v32
    } else {
        var v33 int64 = -9223372036854775808
        v34 := GenS3{F0: 65535}
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: 163, F1: 18446744073709551615, F2: v6}
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS3{F0: (v28 << ((0) & 63))}
        _ = fi36
        _ = v37
    }
    var v38 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
