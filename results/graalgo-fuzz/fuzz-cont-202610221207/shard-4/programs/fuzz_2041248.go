package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int32; F2 string }
type GenS5 struct { F0 uint; F1 uint16; F2 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 float64 = 1e10
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 string = "foo bar"
    v6 := []string{v5, "foo bar", (v5 + v5)}
    var v7 uint32 = 1
    if (0 >= 1) {
        var v8 uint8 = v1
        if (((492 == v7) && (4294967295 != 4294967295)) && (42 > -150)) {
            var v9 int = -9223372036854775808
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch (((v5 + ("\n\t" + (v5 + "a"))) + v5) + v5) {
            case "":
                var v10 uint64 = 1
                _ = v10
            }
            _ = v9
        } else {
            v8 -= 128
            var v11 uint16 = 1
            _ = v11
        }
        _ = v8
    } else {
        v12 := v3[(v7 | (v7 / (((v7 | v7)) | 1)))]
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int16 = 42
        v4[fi13] = (^(v7))
        fmt.Println(v2, v1)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v1 = v1
        fmt.Println(v5)
        v1 -= (v1 - v1)
        _ = fi15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: 255, F1: v7}
        _ = fi16
        _ = v17
    }
    switch v5 {
    case "foo bar":
        if false {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int16 = -765
                if true {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := GenS2{F0: v5, F1: uint16(v1)}
                        delete(v3, v7)
                        v22 := GenS3{F0: v7}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                } else {
                    var v23 float32 = (0.1 + 0.1)
                    if (!(!((3.14159 - v2) < (1e10 / v2)))) {
                        if true {
                            for fi24 := 0; fi24 < 5; fi24++ {
                                var v25 string = v5
                                fmt.Println(v23, v7, v19)
                                var v26 bool = false
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 12; fi27++ {
                                v7 = (v7 / ((v7) | 1))
                                v28 := v3[v7]
                                var v29 float32 = float32(v2)
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v23 -= v23
                                v31 := v4[fi18]
                                _ = fi30
                                _ = v31
                            }
                        }
                        var v32 int32 = 159
                        _ = v32
                    } else {
                        fmt.Println(v1, v7, fi18)
                        v33 := v4[4294967295]
                        _ = v33
                    }
                    _ = v23
                }
                _ = fi18
                _ = v19
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                fmt.Println(v2, v7)
                v4[(fi34 & (fi34 % ((fi34) | 1)))] = (v7 - v7)
                _ = fi34
            }
        } else {
            for fi35 := 0; fi35 < 7; fi35++ {
                var v36 uint = 2
                _ = fi35
                _ = v36
            }
            v1 -= (v1 - (v1 * (v1 / ((v1) | 1))))
        }
    }
    v37 := GenS2{F0: "", F1: 915}
    delete(v3, v7)
    for fi38 := 0; fi38 < 17; fi38++ {
        v7 += v7
        v39 := GenS4{F0: 641, F1: 0, F2: v5}
        _ = fi38
        _ = v39
    }
    switch "\n\t" {
    case "\n\t":
        var v40 int64 = -1
        _ = v40
    case "":
        if false {
            fmt.Println(v1, v5, v7)
            v6 = append(v6, v5)
        }
    case "a":
        if true {
            v41 := GenS5{F0: 1, F1: 65535, F2: 627}
            v42 := GenS4{F0: 941, F1: 42, F2: "a"}
            var v43 int64 = 876
            _ = v41
            _ = v42
            _ = v43
        }
    }
    fmt.Println(v7, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
