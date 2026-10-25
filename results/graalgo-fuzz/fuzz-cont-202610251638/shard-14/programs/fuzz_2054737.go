package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 uint32 = 0
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := []int{}
    v5 := []uint32{(v1 % (((v1 - v1)) | 1)), v1}
    var v6 float32 = 0.0
    var v7 int16 = 1000
    var v8 int = 0
    v3[(v1 | v1)] = "foo bar"
    v6 += ((v6 * (((3.14159 / (v6 / v6)) * (v6 / (v6 - ((v6 / 0.1) * 1e10)))) + 1e10)) * 1e10)
    delete(v3, v1)
    if true {
        var v9 float64 = 0.1
        fmt.Println(v1, v6, v1)
        if (127 < 42) {
            v6 = ((float32(v9) / float32(float64(v6))) / (v6 - v6))
            v10 := GenS1{F0: 70, F1: 255, F2: 129}
            v11 := GenS1{F0: 38, F1: 2, F2: 1}
            _ = v10
            _ = v11
        }
        _ = v9
    }
    v1 += v1
    var v12 bool = false
    v12 = ((!(255 >= 0)) && ("" == ""))
    if (0 > -1) {
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(v1, fi13)
            v3[(^(v1))] = "\n\t"
            _ = fi13
        }
    } else {
        switch "" {
        case "\n\t":
            if (1.5 <= 1.5) {
                v8 -= (^(v8))
                var v14 uint32 = (^(v1))
                var v15 string = "a"
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v4 = append(v4, 711)
                    _ = fi16
                }
                v17 := GenS1{F0: -1, F1: 128, F2: 118}
                _ = v17
            }
        case "foo bar":
            v12 = v12
        case "hello":
            var v18 bool = v12
            _ = v18
        default:
            delete(v3, 65535)
        }
        var v19 uint64 = 4294967295
        _ = v19
    }
    switch "foo bar" {
    case "foo bar":
        var v20 int64 = 42
        _ = v20
    case "\n\t":
        v21 := GenS1{F0: -110, F1: 0, F2: 0}
        _ = v21
    case "a":
        if true {
            if ((2 == 4294967295) && (640 >= 750)) {
                var v22 int64 = 42
                fmt.Println(v22)
                _ = v22
            } else {
                var v23 float32 = -2.25
                _ = v23
            }
            if ((v12 || false) && ((true && true) || v12)) {
                fmt.Println(v1, v1, v1)
            } else {
                var v24 float64 = ((0.1 / 0.1) - 0.1)
                v25 := GenS2{F0: 4294967295, F1: 4294967295, F2: 0}
                var v26 int32 = -855
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 bool = false
            _ = v27
        }
    default:
        if v12 {
            var v28 string = "foo bar"
            _ = v28
        }
    }
    switch "a" {
    case "\n\t":
        if (127 < 0) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            if v12 {
                for fi29 := 0; fi29 < 2; fi29++ {
                    switch (v8 ^ fi29) {
                    case 32768:
                        v30 := GenS3{F0: 176, F1: 32767}
                        _ = v30
                    case 1:
                        var v31 uint16 = 42
                        _ = v31
                    }
                    v1 = (v1 * v1)
                    v32 := GenS1{F0: -128, F1: 1, F2: 2}
                    _ = fi29
                    _ = v32
                }
                fmt.Println(v1)
                v6 -= 1e10
            }
            if (v12 && true) {
                var v33 int16 = (v7 ^ v7)
                fmt.Println(v1)
                switch "foo bar" {
                case "":
                    v33 -= (v7 ^ (v7 << ((323) & 63)))
                case "foo bar":
                    v34 := GenS2{F0: (v8 ^ -374), F1: 32768, F2: 74}
                    _ = v34
                case "hello":
                    v35 := GenS1{F0: -1, F1: 0, F2: 127}
                    _ = v35
                }
                _ = v33
            } else {
                switch (v8 + v8) {
                case 667:
                    var v36 uint64 = 313
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 16; fi37++ {
                        _ = len(v5)
                        var v38 int16 = 1
                        if (2 > 0) {
                            if ((100 == 1) || (-2.25 >= 0.1)) {
                                var v39 uint = 18446744073709551615
                                var v40 bool = false
                                _ = v39
                                _ = v40
                            } else {
                                v41 := GenS4{F0: (^(v7)), F1: -694, F2: v1}
                                if v12 {
                                    v38 += v38
                                } else {
                                    v5 = append(v5, (v1 + (^(v1))))
                                    v42 := GenS3{F0: 47, F1: 2}
                                    var v43 uint16 = 255
                                    _ = v42
                                    _ = v43
                                }
                                _ = v41
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                }
                v2 = append(v2, (("" + "\n\t") + "\n\t"))
            }
        }
    }
    v12 = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
