package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int8 = -46
    var v2 int16 = 618
    var v3 float32 = -2.25
    var v4 bool = (0.0 != 0.0)
    var v5 int8 = (v1 - v1)
    v6 := GenS1{F0: (v5 >> ((255) & 63))}
    v7 := []int{-9223372036854775808, 1}
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 0}
        v11 := GenS2{F0: (119 < 135), F1: (v5 >> ((65535) & 63))}
        fmt.Println(v3)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            _ = len(v7)
            v14 := GenS3{F0: -835}
            if len(v7) > 0 {
                v7[0] = fi13
            }
            _ = fi13
            _ = v14
        }
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 -= (v2 & -857)
                var v16 float64 = (1.5 - float64((v3 * (3.14159 + float32((0.1 + (3.14159 * 0.1)))))))
                _ = fi15
                _ = v16
            }
            v1 = (^(v5))
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: v1}
        fmt.Println(fi17, v2)
        _ = fi17
        _ = v18
    }
    if (v4 || (0 <= 1)) {
        v2 += 1
        if len(v7) > 0 {
            v7[0] = 42
        }
        v19 := GenS4{F0: 0}
        _ = v19
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            if (!(((v4 && true) || (true || (1 > 32767))) && true)) {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v3 = 3.14159
                    v1 -= v5
                    _ = fi21
                }
                switch 127 {
                case 127:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        if ((true || v4) || false) {
                            switch (fi22 - (fi20 ^ (fi22 << ((4294967295) & 63)))) {
                            case 4294967295:
                                v5 = (v1 ^ -8)
                            default:
                                v2 -= 42
                            }
                            var v23 string = (("\n\t" + "foo bar") + "a")
                            _ = v23
                        } else {
                            v7 = append(v7, fi22)
                            switch 4294967295 {
                            case 1:
                                v5 -= 100
                                v24 := GenS1{F0: 42}
                                _ = v24
                            case 9223372036854775807:
                                switch "\n\t" {
                                case "foo bar":
                                    var v25 uint = 0
                                    _ = v25
                                default:
                                    v26 := GenS1{F0: v5}
                                    _ = v26
                                }
                            case 4294967295:
                                v1 = (v5 >> ((4294967295) & 63))
                            }
                            v27 := GenS4{F0: 65535}
                            _ = v27
                        }
                        _ = fi22
                    }
                case 879:
                    var v28 uint = 885
                    _ = v28
                }
                if (v4 || v4) {
                    v8 += 1
                    switch "" {
                    case "foo bar":
                        var v29 uint8 = 255
                        _ = v29
                    }
                }
            } else {
                var v30 int16 = 32767
                var v31 float64 = 3.14159
                _ = v30
                _ = v31
            }
            _ = len(v7)
            _ = fi20
        }
        var v32 float64 = (1.5 * 1e10)
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := GenS4{F0: 99}
            if len(v7) > 0 {
                v7[0] = fi33
            }
            var v36 uint = 4294967295
            _ = fi34
            _ = v35
            _ = v36
        }
        v8 += v8
        _ = fi33
    }
    v4 = true
    for fi37 := 0; fi37 < 18; fi37++ {
        var v38 uint64 = (v8 * v8)
        v5 = -128
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
