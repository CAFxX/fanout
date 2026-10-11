package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 1000
    v3 := []uint32{}
    v4 := make(map[int]int)
    fmt.Println(v2)
    if false {
        v5 := GenS1{F0: 42}
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 42}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint8 = 127
                _ = len(v1)
                v2 = -32768
                fmt.Println(v10)
                _ = fi9
                _ = v10
            }
            _ = fi6
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = 928
            if false {
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS1{F0: 184}
                    if false {
                        fmt.Println(fi11)
                    }
                    if false {
                        var v13 uint64 = 1
                        var v14 int = 42
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            } else {
                if (!(-44 < 2147483647)) {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v4[(fi15 ^ fi15)] = (-(fi15))
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 20; fi16++ {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v2 = (v2 % ((v2) | 1))
                            v18 := GenS1{F0: 32767}
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                    v2 = v2
                }
                switch "\n\t" {
                case "a":
                    v2 += v2
                case "hello":
                    v2 = v2
                }
                v19 := GenS1{F0: 65535}
                _ = v19
            }
        }
        _ = v5
    }
    var v20 string = "hello"
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint = 494
            v24 := GenS1{F0: 806}
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int32 = 2147483647
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(fi21)
            _ = fi26
        }
        _ = fi21
        _ = v25
    }
    v27 := GenS1{F0: 1}
    switch 215 {
    case 42:
        fmt.Println(v2, v20)
    case -2147483648:
        var v28 uint8 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        delete(v4, (fi29 % ((2147483647) | 1)))
        _ = fi29
    }
    fmt.Println(v20, v20)
    fmt.Println(v2)
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := GenS2{F0: 2147483647, F1: fi30}
        _ = fi30
        _ = v31
    }
    v32 := v1[v20]
    v2 -= -1
    for fi33 := 0; fi33 < 4; fi33++ {
        if false {
            switch v32 {
            case 0:
                v2 -= (v2 / (((v2 << ((255) & 63))) | 1))
            case 9223372036854775807:
                fmt.Println(v2, v2)
            default:
                var v34 bool = true
                _ = v34
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                if len(v3) > 0 {
                    v3[0] = 447
                }
                var v36 int = (v32 / (((fi33 + (fi33 * fi35))) | 1))
                _ = fi35
                _ = v36
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi33
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 float32 = (((((float32(0.1) - 0.0) * -2.25) - (1.5 + 1.5)) + 0.1) + 0.0)
        _ = fi37
        _ = v38
    }
    if false {
        fmt.Println(v20, v20)
        for fi39 := 0; fi39 < 14; fi39++ {
            v20 = "a"
            if false {
                v40 := GenS1{F0: 2}
                if true {
                    v20 = ((v20 + v20) + v20)
                    for fi41 := 0; fi41 < 17; fi41++ {
                        var v42 uint32 = 65535
                        _ = fi41
                        _ = v42
                    }
                    fmt.Println(v32)
                } else {
                    v3 = append(v3, 65535)
                    v32 = v32
                    fmt.Println(v2, fi39, v2)
                }
                v2 = -1
                _ = v40
            } else {
                v43 := GenS3{F0: 48, F1: false}
                v2 += v2
                var v44 uint = 1
                _ = v43
                _ = v44
            }
            _ = fi39
        }
        v3 = append(v3, 156)
    }
    v2 -= v2
    v45 := GenS1{F0: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v45
}
