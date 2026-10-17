package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 bool = false
    var v5 float32 = 0.1
    var v6 bool = v4
    if (107 == 0) {
        fmt.Println(v5)
        var v7 uint = 420
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: 1}
                var v11 uint16 = 255
                var v12 int32 = -2147483648
                v4 = v4
                if (!v4) {
                    var v13 string = ""
                    if (v12 > v12) {
                        var v14 float32 = float32(0.0)
                        var v15 uint16 = v11
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint8 = 1
                        var v17 uint64 = 924
                        v6 = (-105 != int8(v11))
                        _ = v16
                        _ = v17
                    }
                    _ = v13
                } else {
                    if v6 {
                        v18 := GenS2{F0: 42}
                        _ = v18
                    } else {
                        if (42 != (v1 * 1)) {
                            v6 = v4
                        }
                    }
                    if v6 {
                        if (v4 || ((0.1 - 0.1) == (0.1 + ((1e10 + 1e10) * (0.1 / 1.0))))) {
                            v1 += 951
                            switch fi8 {
                            case 32767:
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v6, v12)
                                case "a":
                                    var v19 uint64 = 42
                                    _ = v19
                                }
                            default:
                                v20 := GenS1{F0: 538}
                                _ = v20
                            }
                        } else {
                            switch ("foo bar" + ("\n\t" + "foo bar")) {
                            case "foo bar":
                                v21 := v2[fi9]
                                _ = v21
                            }
                            v6 = v4
                        }
                    }
                }
                v22 := GenS3{F0: 0}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v22
            }
            switch -592 {
            case 127:
                if (0 > 112) {
                    fmt.Println(v1, v4, v6)
                    v7 = v7
                    var v23 int16 = v3
                    _ = v23
                } else {
                    var v24 int16 = -810
                    v25 := GenS1{F0: 48}
                    fmt.Println(v7)
                    _ = v24
                    _ = v25
                }
            case 32768:
                v7 += (^(v7))
            case 954:
                if v4 {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 uint16 = 0
                        _ = fi26
                        _ = v27
                    }
                    var v28 string = ((("\n\t" + "a") + "hello") + "\n\t")
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v28 = ""
                        var v30 float32 = v5
                        fmt.Println(v1)
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            default:
                switch "" {
                case "hello":
                    var v31 int64 = 0
                    _ = v31
                default:
                    v7 = (v7 >> (((^(v7))) & 63))
                }
            }
            _ = fi8
        }
        _ = v7
    }
    v32 := GenS3{F0: 0}
    v1 -= (^(v1))
    var v33 uint32 = 42
    if (v6 || (v5 == 1e10)) {
        v6 = (42 <= -99)
        v3 += (v3 / ((v3) | 1))
        var v34 float64 = (0.1 + (-2.25 / -2.25))
        _ = v34
    }
    var v35 uint16 = 1
    if v6 {
        v35 = v35
        v3 = (v3 & (-(v3)))
    } else {
        var v36 float32 = 1.5
        v36 = 0.0
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint32 = v33
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
