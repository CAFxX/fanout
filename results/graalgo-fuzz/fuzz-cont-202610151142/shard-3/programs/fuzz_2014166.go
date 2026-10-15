package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }
type GenS5 struct { F0 bool; F1 int64; F2 uint }
type GenS6 struct { F0 string; F1 int32; F2 int64 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]string)
    var v3 int = -466
    var v4 float32 = 1e10
    switch (v3 / (((v3 / (((v3 >> ((4294967295) & 63))) | 1))) | 1)) {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 20; fi5++ {
            v4 = 0.1
            _ = fi5
        }
    case 160:
        var v6 uint = 18446744073709551615
        _ = v6
    case -1:
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(v1, v3, v4)
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := []int{fi8}
                var v10 float64 = (1.5 * 1.5)
                if (true && (-1 == 65)) {
                    v10 += 3.14159
                }
                v4 -= v4
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    default:
        v11 := v2[v3]
        _ = v11
    }
    v1 = (^(v1))
    var v12 uint8 = 228
    for fi13 := 0; fi13 < 15; fi13++ {
        if true {
            v14 := v2[(fi13 % ((830) | 1))]
            var v15 int = (v3 ^ v3)
            fmt.Println(v4, v3, v12)
            _ = v14
            _ = v15
        } else {
            if true {
                for fi16 := 0; fi16 < 19; fi16++ {
                    if true {
                        v17 := GenS1{F0: 32767, F1: false}
                        if false {
                            v18 := GenS1{F0: 2, F1: (0.0 != 0.0)}
                            v19 := v2[4294967295]
                            v20 := GenS1{F0: 32767, F1: false}
                            fmt.Println(v4, v12, v19)
                            v12 += 128
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        switch v3 {
                        case -9223372036854775808:
                            fmt.Println(fi13, fi13)
                        default:
                            v21 := GenS2{F0: true, F1: -9223372036854775808}
                            _ = v21
                        }
                        _ = v17
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v22 uint = 363
                            _ = v22
                        default:
                            v1 = (v1 * v1)
                        }
                    }
                    _ = fi16
                }
                switch ("" + "") {
                case "a":
                    v23 := GenS3{F0: -1}
                    _ = v23
                case "":
                    v24 := GenS2{F0: false, F1: 402}
                    _ = v24
                }
            } else {
                fmt.Println(v4)
                switch ("" + "hello") {
                case "\n\t":
                    var v25 uint64 = 42
                    _ = v25
                case "a":
                    v12 = (v12 >> ((353) & 63))
                case "foo bar":
                    if false {
                        var v26 float64 = -2.25
                        var v27 int16 = 42
                        _ = v26
                        _ = v27
                    } else {
                        if (true || (1 > 885)) {
                            v28 := GenS3{F0: 9223372036854775807}
                            _ = v28
                        }
                        fmt.Println(v1)
                    }
                default:
                    if (v4 != (0.1 + 0.1)) {
                        v29 := GenS1{F0: 57, F1: (327 > 65535)}
                        v30 := v2[(v3 - fi13)]
                        v4 += (((v4 * ((float32(0.0) + 1.5) + v4)) - v4) / v4)
                        _ = v29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 9; fi31++ {
                    if true {
                        v12 += (v12 / (((v1 + v1)) | 1))
                    }
                    _ = fi31
                }
            }
            if (!(((v4 - v4) >= (3.14159 - ((-2.25 * (((float32(3.14159) / v4) / v4) - 0.1)) - (v4 - (v4 - v4))))) && false)) {
                v32 := GenS4{F0: 230, F1: true, F2: false}
                v33 := GenS3{F0: 0}
                v34 := GenS5{F0: true, F1: 42, F2: 4294967295}
                _ = v32
                _ = v33
                _ = v34
            }
            fmt.Println(fi13, v4, v4)
        }
        _ = fi13
    }
    var v35 int = (v3 / ((v3) | 1))
    if true {
        v1 = (v12 % (((v12 & (^(v12)))) | 1))
        v36 := GenS1{F0: 244, F1: (((v4 / v4) >= float32(0.1)) && false)}
        var v37 int64 = -809
        _ = v36
        _ = v37
    }
    _ = len(v2)
    var v38 int = 2147483647
    v4 = float32((0.0 / 1.5))
    if (0 < 2) {
        var v39 string = "a"
        _ = v39
    } else {
        var v40 int64 = 42
        v41 := GenS2{F0: false, F1: (v40 | v40)}
        v42 := GenS6{F0: "\n\t", F1: int32(v40), F2: v40}
        _ = v40
        _ = v41
        _ = v42
    }
    v38 += 1
    v43 := []int{(v35 - v35)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v43
}
