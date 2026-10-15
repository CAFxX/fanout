package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    if false {
        var v4 uint8 = 255
        v4 -= v4
        var v5 uint16 = 137
        _ = v4
        _ = v5
    } else {
        var v6 int16 = -1
        var v7 string = "a"
        _ = v6
        _ = v7
    }
    var v8 bool = false
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9)
        var v10 int64 = 42
        var v11 int = (fi9 - fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float64 = 1e10
        var v14 int32 = -2147483648
        v14 = v14
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v8)
    var v15 int32 = -2147483648
    if ((v15 % ((v15) | 1)) != 42) {
        fmt.Println(v8, v15)
    } else {
        if v8 {
            for fi16 := 0; fi16 < 4; fi16++ {
                v15 -= (v15 / (((v15 % ((0) | 1))) | 1))
                v17 := GenS1{F0: 1000, F1: int16(fi16), F2: 0}
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                v8 = v8
                fmt.Println(v15, fi18)
                _ = fi18
            }
            v8 = (!v8)
        }
        if (42 != 128) {
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 string = (("" + "foo bar") + "")
                v15 -= (v15 + (v15 | (v15 | 200)))
                if (1 <= 32767) {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v20 = ("\n\t" + v20)
                }
                _ = fi19
                _ = v20
            }
            if true {
                var v21 uint8 = 0
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 float32 = (float32(0.1) * 1e10)
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                for fi24 := 0; fi24 < 11; fi24++ {
                    fmt.Println(fi24, fi24)
                    v8 = v8
                    _ = fi24
                }
            }
            for fi25 := 0; fi25 < 5; fi25++ {
                if (!false) {
                    var v26 int8 = int8(fi25)
                    v15 += (v15 << ((2) & 63))
                    if v8 {
                        if v8 {
                            var v27 int = 1
                            v15 = (v15 << ((42) & 63))
                            v2["a"] = v27
                            _ = v27
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v2)
                    } else {
                        var v28 int16 = 748
                        for fi29 := 0; fi29 < 20; fi29++ {
                            switch 32767 {
                            case 994:
                                if v8 {
                                    v28 += v28
                                } else {
                                    _ = len(v2)
                                    fmt.Println(fi29, fi25, v8)
                                    v30 := GenS2{F0: 415}
                                    _ = v30
                                }
                            case 128:
                                var v31 bool = v8
                                _ = v31
                            case 715:
                                v28 = (v28 ^ v28)
                            default:
                                var v32 int = 128
                                _ = v32
                            }
                            if len(v1) > 0 {
                                v1[0] = 922
                            }
                            switch "a" {
                            case "":
                                fmt.Println(fi29, v26)
                            }
                            _ = fi29
                        }
                        var v33 uint8 = 105
                        _ = v28
                        _ = v33
                    }
                    _ = v26
                } else {
                    v8 = (127 > 1)
                    for fi34 := 0; fi34 < 15; fi34++ {
                        var v35 uint8 = 2
                        if ((255 > 1) || ((float64(-2.25) / 1.0) != (3.14159 * (-2.25 / 1.5)))) {
                            v1 = append(v1, 923)
                        } else {
                            v15 = v15
                        }
                        var v36 int32 = v15
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(v15)
                }
                if false {
                    fmt.Println(fi25, fi25)
                }
                _ = fi25
            }
        }
    }
    var v37 uint64 = 1
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 int16 = -1
        fmt.Println(v15, v15, fi38)
        _ = fi38
        _ = v39
    }
    v37 += (v37 & v37)
    v3["a"] = (139 != 255)
    switch 127 {
    case 390:
        if false {
            v40 := GenS1{F0: 194, F1: 1, F2: 9223372036854775807}
            _ = v40
        }
    default:
        _ = len(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v8
}
