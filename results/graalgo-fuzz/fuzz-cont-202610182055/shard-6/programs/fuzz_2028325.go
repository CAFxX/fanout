package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 0
    var v4 float64 = (float64(3.14159) + (1.5 + 0.0))
    var v5 uint64 = 0
    v6 := make(map[int]uint32)
    if false {
        v7 := []int{int(v3), 1, -575, -1}
        _ = v7
    } else {
        v3 = 2
        var v8 uint64 = 42
        v9 := GenS1{F0: 106, F1: -1}
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: 825}
    switch 2147483647 {
    case 4294967295:
        if (false && false) {
            var v11 uint32 = 1
            if ("a" == ("" + "foo bar")) {
                var v12 int16 = -32768
                var v13 bool = ((3.14159 + 1.5) > 3.14159)
                _ = v12
                _ = v13
            }
            fmt.Println(v3, v11)
            _ = v11
        } else {
            v5 += v5
            var v14 float64 = 1.5
            for fi15 := 0; fi15 < 3; fi15++ {
                v4 = (1.5 + 0.0)
                v1[(^(fi15))] = (fi15 % ((fi15) | 1))
                v16 := []int{224, fi15}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    case 2147483647:
        v5 = v5
    default:
        v1[int(v3)] = 127
    }
    v5 = v5
    fmt.Println(v5, v3, v5)
    v4 -= 1e10
    fmt.Println(v5)
    var v17 int64 = 0
    v17 += (-(v17))
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 string = (("a" + "foo bar") + "")
        for fi20 := 0; fi20 < 19; fi20++ {
            if (2147483648 > 2147483648) {
                v4 -= (v4 + (1e10 * float64(1.5)))
                fmt.Println(v17, v17, v3)
                v2[1] = true
            } else {
                var v21 int64 = v17
                v22 := GenS2{F0: 959, F1: 0}
                var v23 int16 = -32768
                v23 = v23
                if (((!true) || ((^(v3)) != 42)) || (-128 > 0)) {
                    var v24 uint16 = 108
                    _ = v24
                } else {
                    var v25 string = ("foo bar" + "")
                    if (v5 == (v5 << ((255) & 63))) {
                        if false {
                            var v26 uint8 = 0
                            v27 := GenS2{F0: -2147483648, F1: 959}
                            v28 := GenS3{F0: 125, F1: 270}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    } else {
                        v5 -= 139
                        switch (v25 + (v19 + ((v25 + ((((v19 + v19) + v19) + (v19 + v25)) + v25)) + ((v25 + v19) + v19)))) {
                        case "foo bar":
                            var v29 uint16 = 255
                            _ = v29
                        case "hello":
                            switch (fi18 - int(v23)) {
                            case 0:
                                v30 := []int{-194}
                                _ = v30
                            case -520:
                                var v31 int32 = -284
                                _ = v31
                            case 852:
                                v32 := GenS3{F0: 417, F1: 787}
                                _ = v32
                            default:
                                v33 := GenS2{F0: 299, F1: 65535}
                                _ = v33
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "":
                                var v34 int64 = int64(fi20)
                                _ = v34
                            case "hello":
                                v35 := GenS4{F0: 9223372036854775807, F1: -1}
                                _ = v35
                            }
                        default:
                            fmt.Println(v25)
                        }
                        if false {
                            if (!(v25 != v19)) {
                                v36 := []int{}
                                var v37 string = v19
                                if ((v3 ^ v3) == 75) {
                                    _ = len(v1)
                                    fmt.Println(v25)
                                }
                                if len(v36) > 0 {
                                    v36[0] = (fi20 % (((fi18 - fi18)) | 1))
                                }
                                v19 = v37
                                v38 := v1[-2147483648]
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                var v39 uint16 = v3
                                var v40 string = "hello"
                                v5 -= 18446744073709551615
                                _ = v39
                                _ = v40
                            }
                            v4 -= v4
                        }
                        var v41 string = v19
                        var v42 uint8 = 128
                        _ = v41
                        _ = v42
                    }
                    _ = v25
                }
                var v43 int32 = 42
                _ = v21
                _ = v22
                _ = v23
                _ = v43
            }
            v4 += float64(1.5)
            switch v19 {
            case "\n\t":
                fmt.Println(fi20)
            case "":
                v4 = ((((v4 - float64(1e10)) * ((v4 + 0.0) + ((((float64(-2.25) - (v4 + v4)) - -2.25) - ((float64((1.5 + 0.0)) * 0.1) / (float64((-2.25 - 1.5)) - v4))) * v4))) - (float64((0.0 / (0.1 * (0.0 + 0.1)))) / 1e10)) - v4)
            case "foo bar":
                v5 = 146
            default:
                var v44 string = v19
                _ = v44
            }
            var v45 string = (v19 + "a")
            v3 = v3
            _ = fi20
            _ = v45
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
