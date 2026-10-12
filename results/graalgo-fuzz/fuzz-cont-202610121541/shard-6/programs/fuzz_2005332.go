package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint32; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }

func main() {
    var v1 bool = false
    v2 := []uint32{660, 490, 65535, 2147483648}
    v3 := []int{32768, -9223372036854775808, 1, 490}
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    if len(v2) > 0 {
        v2[0] = 218
    }
    if (0 <= -22) {
        v1 = (true && true)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 1, F1: (("a" + "") + "hello"), F2: 29}
        var v8 int32 = 2147483647
        v8 += v8
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v1)
    v1 = (2 <= 42)
    switch 326 {
    case 371:
        v3 = append(v3, 4294967295)
    case 4294967295:
        var v9 float64 = (float64((1e10 - 0.1)) - (-2.25 + ((1.5 + 1.5) + 0.0)))
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = (1e10 >= ((1e10 * 1e10) + (float32(1e10) + (1.5 * (0.1 * 3.14159)))))
        if false {
            v11 = v11
            v1 = (true && ((-2.25 != -2.25) || (fi10 >= fi10)))
        }
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        switch fi12 {
        case 9223372036854775807:
            v13 := GenS1{F0: -966, F1: "hello", F2: 29}
            _ = v13
        case 235:
            _ = len(v3)
        case 42:
            v1 = ((272 >= 2) && true)
        }
        v14 := GenS2{F0: 1, F1: 65535}
        _ = fi12
        _ = v14
    }
    var v15 uint32 = 2
    if false {
        if (1 < -1) {
            v15 += (v15 / (((v15 << ((638) & 63))) | 1))
            if (-32768 == -788) {
                var v16 uint8 = 128
                var v17 uint32 = (v15 % ((v15) | 1))
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: 1}
            _ = v18
        } else {
            if v1 {
                var v19 float32 = 1e10
                v20 := GenS4{F0: 4294967295, F1: (v15 - (v15 << ((42) & 63))), F2: (495 < 1)}
                if false {
                    var v21 int8 = -113
                    switch "a" {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v22 int64 = -1
                        _ = v22
                    }
                    _ = v21
                }
                _ = v19
                _ = v20
            }
            if ("\n\t" >= "a") {
                switch "a" {
                case "a":
                    for fi23 := 0; fi23 < 5; fi23++ {
                        switch fi23 {
                        case 4294967295:
                            v24 := GenS1{F0: -2147483648, F1: "", F2: 42}
                            _ = v24
                        case -227:
                            var v25 uint8 = 128
                            _ = v25
                        case 128:
                            v26 := GenS5{F0: 757}
                            _ = v26
                        default:
                            switch (-(fi23)) {
                            case 127:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    switch fi27 {
                                    case 32767:
                                        v1 = v1
                                    }
                                    if (!(1e10 == (-2.25 + 1.5))) {
                                        var v28 int = fi23
                                        v1 = v1
                                        _ = v28
                                    }
                                    var v29 uint32 = (v15 ^ (v15 / (((v15 + v15)) | 1)))
                                    _ = fi27
                                    _ = v29
                                }
                            case 0:
                                v15 = v15
                            case 380:
                                switch "a" {
                                case "hello":
                                    var v30 uint16 = 760
                                    _ = v30
                                case "":
                                    v1 = true
                                default:
                                    v15 = (v15 % (((v15 | v15)) | 1))
                                }
                            }
                        }
                        _ = fi23
                    }
                case "":
                    v31 := GenS1{F0: 2147483647, F1: "hello", F2: 0}
                    _ = v31
                }
                v32 := GenS6{F0: false}
                fmt.Println(v1, v1)
                _ = v32
            }
        }
    } else {
        var v33 int32 = 931
        _ = v33
    }
    if (!v1) {
        fmt.Println(v1, v1)
        v3 = append(v3, 128)
    } else {
        switch 2147483647 {
        case -1:
            v15 -= 2
        case -561:
            v15 = (v15 & v15)
        default:
            var v34 int8 = -108
            _ = v34
        }
        if false {
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 uint8 = 128
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi35
                _ = v36
            }
        }
        var v37 int32 = 2147483647
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
