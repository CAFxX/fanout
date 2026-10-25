package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]string)
    var v3 uint16 = 42
    v4 := GenS1{F0: true, F1: (^(v3))}
    var v5 int16 = -1
    v6 := GenS1{F0: (float32(v1) <= ((0.1 / float32(v1)) * (float32(v1) - 1.5))), F1: 257}
    var v7 uint8 = 1
    v5 = -32768
    switch 0 {
    case 128:
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                for fi9 := 0; fi9 < 10; fi9++ {
                    if false {
                        v10 := []int{(^(fi8))}
                        if true {
                            var v11 uint = 18
                            _ = v11
                        } else {
                            var v12 float64 = 0.1
                            v2[0] = "a"
                            _ = v12
                        }
                        _ = v10
                    } else {
                        v7 = (v7 % ((230) | 1))
                        v7 = (v7 >> ((255) & 63))
                        v3 -= v3
                        var v13 float64 = float64(1e10)
                        v14 := []int{fi8, 128, (fi9 & fi8), (fi9 >> ((26) & 63))}
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                }
            } else {
                v7 -= v7
            }
            v1 += v1
            var v15 int = fi8
            _ = fi8
            _ = v15
        }
    case 9223372036854775807:
        v16 := GenS1{F0: true, F1: v3}
        _ = v16
    }
    switch 0 {
    case -9223372036854775808:
        if ((v3 % ((v3) | 1)) < 65535) {
            for fi17 := 0; fi17 < 8; fi17++ {
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 uint16 = v3
                    fmt.Println(v3, v7)
                    _ = fi18
                    _ = v19
                }
                v1 -= float64(-2.25)
                var v20 int = -540
                _ = fi17
                _ = v20
            }
            v2[4294967295] = ""
            v7 -= (v7 >> ((2) & 63))
        }
    case 127:
        v7 += (v7 ^ 42)
    case 32767:
        if true {
            v3 -= (v3 & (v3 / (((^(v3))) | 1)))
        }
    default:
        v21 := GenS1{F0: ((-626 >= 4294967295) && (true || (false || true))), F1: (v3 >> ((uint(v7)) & 63))}
        _ = v21
    }
    v1 = (3.14159 / 1e10)
    for fi22 := 0; fi22 < 11; fi22++ {
        switch "" {
        case "\n\t":
            fmt.Println(v1, fi22, v5)
        case "foo bar":
            if (0.1 >= 0.0) {
                v23 := GenS1{F0: (false || true), F1: v3}
                _ = v23
            }
        }
        var v24 uint8 = (v7 & (v7 ^ v7))
        v25 := GenS1{F0: true, F1: (v3 + 32767)}
        _ = fi22
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint = uint(v7)
        switch "" {
        case "hello":
            _ = len(v2)
        case "":
            v5 -= v5
        }
        v7 = v7
        _ = fi26
        _ = v27
    }
    var v28 bool = true
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
