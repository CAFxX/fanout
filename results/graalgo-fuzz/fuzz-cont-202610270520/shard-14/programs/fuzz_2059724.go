package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0() string {
    v6 := GenS1{F0: 2147483647}
    var v7 float32 = 3.14159
    fmt.Println(v7)
    fmt.Println(v7)
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 = float32(float64(0.0))
        _ = fi8
    }
    _ = v6
    _ = v7
    return ("\n\t" + "a")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -804
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    var v5 int16 = v2
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v5, v5)
            v11 := []int{(fi9 >> ((255) & 63)), (-(fi10)), (-(fi10))}
            _ = fi10
            _ = v11
        }
        if (false && (false || (fi9 <= 127))) {
            v5 += (v2 + v5)
            fmt.Println(v5, fi9)
        }
        v12 := GenS2{F0: 152}
        _ = fi9
        _ = v12
    }
    v5 = (v2 | 1)
    v5 -= (v2 * 42)
    v13 := []int{127}
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            v5 -= (v5 / ((-32768) | 1))
            if false {
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := GenS2{F0: 42}
                    v2 += v2
                    if true {
                        if true {
                            fmt.Println(v2, fi14, v5)
                            fmt.Println(fi15, v5)
                            v4 += (((float32(float64((0.1 - (float32(0.0) * v4)))) / 1e10) * v4) * (float32(3.14159) / v4))
                        } else {
                            switch (-(fi14)) {
                            case -520:
                                fmt.Println(v4, v4)
                            default:
                                v2 = v2
                            }
                            fmt.Println(v2, v2, v4)
                            switch fi15 {
                            case 1:
                                fmt.Println(v5)
                            }
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                v17 := GenS3{F0: 2, F1: 971, F2: 255}
                for fi18 := 0; fi18 < 2; fi18++ {
                    if false {
                        v19 := GenS4{F0: -32768}
                        v5 = 868
                        v5 += (v5 & v2)
                        switch 128 {
                        case 127:
                            if false {
                                v5 = -32768
                                var v20 float32 = 0.0
                                v5 += (v5 & 32767)
                                switch "a" {
                                case "a":
                                    var v21 int32 = -1
                                    _ = v21
                                }
                                var v22 int = fi18
                                var v23 uint16 = 502
                                _ = v20
                                _ = v22
                                _ = v23
                            }
                        case 32767:
                            v2 -= (v5 ^ v5)
                        }
                        _ = v19
                    }
                    _ = len(v3)
                    _ = fi18
                }
                _ = v17
            }
            _ = fi14
        }
        v24 := GenS1{F0: 1}
        _ = v24
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println(v2, v2, v4)
    v2 = v5
    switch 422 {
    case 4294967295:
        v25 := GenS2{F0: 1}
        _ = v25
    default:
        var v26 int32 = -907
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if ((v5 < v2) || (("\n\t" + (("" + ("foo bar" + "foo bar")) + "")) <= "a")) {
            v28 := v1[0]
            v29 := GenS5{F0: 1, F1: 2}
            v30 := GenS2{F0: -1}
            _ = v28
            _ = v29
            _ = v30
        } else {
            var v31 string = ""
            var v32 float64 = 0.1
            v33 := GenS2{F0: -876}
            _ = v31
            _ = v32
            _ = v33
        }
        v5 -= v2
        v4 += 0.1
        _ = fi27
    }
    v34 := GenS2{F0: -407}
    v35 := GenS4{F0: 32767}
    for fi36 := 0; fi36 < 8; fi36++ {
        for fi37 := 0; fi37 < 9; fi37++ {
            v5 = (v5 << ((65535) & 63))
            v5 = (v5 << ((644) & 63))
            v5 += (v2 * (v5 ^ 32767))
            _ = fi37
        }
        _ = fi36
    }
    switch ("\n\t" + "hello") {
    case "":
        v5 = (^(v2))
    default:
        var v38 uint32 = 876
        _ = v38
    }
    v4 = (1e10 / ((v4 - (((((0.1 * v4) - ((v4 * 3.14159) * 0.0)) + v4) - 1e10) - v4)) - v4))
    fmt.Println(v2)
    var v39 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
