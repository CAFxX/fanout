package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 int64 }
func genFn0() int16 {
    v5 := []int{4294967295, -2147483648, 206, -1}
    v5 = append(v5, 128)
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "\n\t" {
        case "\n\t":
            for fi7 := 0; fi7 < 17; fi7++ {
                if (!false) {
                    if false {
                        var v8 bool = ((255 > 194) || (!false))
                        var v9 uint64 = 1
                        _ = v8
                        _ = v9
                    }
                } else {
                    v10 := GenS1{F0: 18446744073709551615, F1: 65535}
                    var v11 uint = 0
                    switch 42 {
                    case -140:
                        if true {
                            var v12 string = ("foo bar" + "")
                            var v13 bool = false
                            fmt.Println(v11, v11)
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint8 = 251
                            v15 := GenS1{F0: (v11 | 65535), F1: 2147483648}
                            var v16 float32 = (3.14159 * (((1e10 / 0.1) - (3.14159 + float32((1e10 / ((1.5 - 1e10) + 3.14159))))) * float32(3.14159)))
                            v16 -= float32(0.0)
                            v5 = append(v5, fi7)
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = v10
                    _ = v11
                }
                v17 := GenS2{F0: 128}
                _ = fi7
                _ = v17
            }
        case "a":
            fmt.Println(fi6)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v18 float64 = 1e10
        if false {
            if false {
                v19 := GenS3{F0: fi6, F1: (fi6 % ((fi6) | 1)), F2: 1}
                if false {
                    switch -1 {
                    case 32768:
                        v18 = v18
                    case 32767:
                        v20 := GenS3{F0: -224, F1: fi6, F2: 466}
                        _ = v20
                    case -495:
                        v21 := GenS4{F0: 32767}
                        _ = v21
                    }
                } else {
                    v22 := GenS5{F0: -7, F1: 429}
                    v23 := GenS4{F0: 255}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
            v18 = v18
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS1{F0: 65535, F1: 412}
                v26 := GenS6{F0: 100, F1: 2147483647, F2: 42}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = fi6
        _ = v18
    }
    _ = v5
    return 32767
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -1
    v3 := make(map[string]string)
    var v4 int = 32768
    v2 = (^(v2))
    v1 = 4294967295
    switch (("a" + (("\n\t" + "") + "a")) + "\n\t") {
    case "\n\t":
        var v27 int8 = -128
        _ = v27
    default:
        v28 := GenS3{F0: v4, F1: v4, F2: v1}
        _ = v28
    }
    v1 = 526
    var v29 int32 = (v2 / ((0) | 1))
    v29 -= (v2 / (((v29 * v2)) | 1))
    v1 += (v1 + (v1 / (((v1 | 123)) | 1)))
    v30 := GenS4{F0: 0}
    for fi31 := 0; fi31 < 8; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS3{F0: (fi31 * (fi31 - 32767)), F1: (v4 - (v4 % ((v4) | 1))), F2: 737}
            v1 = (v1 >> ((441) & 63))
            _ = fi32
            _ = v33
        }
        var v34 uint32 = 1
        _ = fi31
        _ = v34
    }
    if (65535 < 590) {
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS7{F0: (2 > 4294967295), F1: (v2 / (((v29 << ((2) & 63))) | 1)), F2: 9223372036854775807}
            var v37 bool = false
            _ = fi35
            _ = v36
            _ = v37
        }
        v38 := GenS3{F0: (v4 ^ 1), F1: v4, F2: 0}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
