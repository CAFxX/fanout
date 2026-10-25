package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS4 struct { F0 string; F1 int32; F2 int }
type GenS5 struct { F0 uint; F1 int8; F2 string }
type GenS6 struct { F0 uint64; F1 int32; F2 int64 }

func main() {
    var v1 uint32 = 157
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 255
    for fi4 := 0; fi4 < 16; fi4++ {
        switch (("foo bar" + ("\n\t" + "a")) + "") {
        case "\n\t":
            for fi5 := 0; fi5 < 2; fi5++ {
                var v6 int64 = int64(fi5)
                switch fi4 {
                case 935:
                    if true {
                        if false {
                            var v7 uint16 = v3
                            _ = v7
                        }
                        var v8 uint64 = 4294967295
                        _ = v8
                    }
                case 371:
                    var v9 bool = (true || false)
                    _ = v9
                }
                _ = fi5
                _ = v6
            }
        }
        var v10 string = "a"
        if false {
            v2 -= int64(v1)
            v11 := GenS1{F0: 173, F1: -1}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := []int{4294967295, fi4, (fi12 & fi12), fi12}
                if (!(v1 != (v1 % ((509) | 1)))) {
                    var v14 int8 = 0
                    var v15 bool = ((v14 != (v14 << ((1) & 63))) || (0 == 1))
                    _ = v14
                    _ = v15
                } else {
                    if (!true) {
                        var v16 uint = 2
                        v17 := GenS1{F0: 0, F1: 961}
                        _ = v16
                        _ = v17
                    }
                }
                if (true && true) {
                    var v18 uint64 = 907
                    v3 = 42
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
            switch fi4 {
            case -1:
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS1{F0: 2147483647, F1: 42}
                    v21 := []int{2147483647}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case -177:
                var v22 float32 = 3.14159
                _ = v22
            default:
                for fi23 := 0; fi23 < 4; fi23++ {
                    switch (^(fi4)) {
                    case 42:
                        if false {
                            var v24 uint32 = 391
                            v1 = 4294967295
                            _ = v24
                        } else {
                            v25 := []int{2147483647, (fi23 << ((671) & 63)), (fi23 / (((fi23 & (fi4 >> ((1) & 63)))) | 1)), fi4}
                            _ = v25
                        }
                    case -96:
                        switch fi4 {
                        case -952:
                            if true {
                                v3 += (v3 << ((255) & 63))
                            } else {
                                var v26 int8 = -128
                                fmt.Println(v10)
                                if true {
                                    if false {
                                        var v27 uint16 = (v3 - (v3 ^ v3))
                                        switch v10 {
                                        case "a":
                                            var v28 float32 = 3.14159
                                            _ = v28
                                        default:
                                            switch fi23 {
                                            case -9223372036854775808:
                                                var v29 int16 = -73
                                                _ = v29
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    var v30 uint64 = 4294967295
                                                    _ = v30
                                                default:
                                                    v27 = v27
                                                }
                                            }
                                        }
                                        var v31 uint16 = (v3 + v3)
                                        _ = v27
                                        _ = v31
                                    }
                                    v32 := []int{fi4, 306}
                                    v33 := GenS1{F0: 9, F1: 0}
                                    _ = v32
                                    _ = v33
                                }
                                _ = v26
                            }
                        default:
                            fmt.Println(v10)
                        }
                    default:
                        var v34 int16 = 1000
                        _ = v34
                    }
                    _ = fi23
                }
            }
            switch (^(fi4)) {
            case 683:
                var v35 float64 = 3.14159
                _ = v35
            case -468:
                v36 := []int{}
                _ = v36
            default:
                fmt.Println(fi4)
            }
        }
        _ = fi4
        _ = v10
    }
    if ("\n\t" != "foo bar") {
        fmt.Println(v2)
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := []int{-847, (fi37 - (fi37 & fi37)), (fi37 / ((fi37) | 1))}
            _ = fi37
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 16; fi39++ {
            for fi40 := 0; fi40 < 19; fi40++ {
                v41 := GenS1{F0: 491, F1: 2147483647}
                var v42 int64 = 199
                _ = fi40
                _ = v41
                _ = v42
            }
            v43 := GenS2{F0: 1}
            _ = fi39
            _ = v43
        }
        if false {
            var v44 uint32 = (v1 & (v1 * v1))
            fmt.Println(v1)
            _ = v44
        }
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := GenS3{F0: 0, F1: 1, F2: 1}
            if true {
                v47 := []int{(^(fi45)), fi45}
                v1 = (v1 >> ((18446744073709551615) & 63))
                v48 := GenS4{F0: "", F1: -1, F2: fi45}
                _ = v47
                _ = v48
            }
            _ = fi45
            _ = v46
        }
    }
    v49 := GenS3{F0: 18446744073709551615, F1: 128, F2: 42}
    v1 = 65535
    v50 := GenS5{F0: 5, F1: 0, F2: "foo bar"}
    for fi51 := 0; fi51 < 19; fi51++ {
        for fi52 := 0; fi52 < 14; fi52++ {
            var v53 uint64 = 0
            v2 -= (v2 << ((313) & 63))
            v2 = v2
            _ = fi52
            _ = v53
        }
        if (!true) {
            if (false && (true || (!(73 < 18446744073709551615)))) {
                fmt.Println(v3)
            } else {
                v54 := []int{fi51, 4294967295, -455}
                fmt.Println(fi51, v3)
                _ = v54
            }
        }
        if (!((fi51 & (fi51 - 128)) >= 32768)) {
            v55 := []int{}
            v56 := GenS6{F0: 18446744073709551615, F1: 1, F2: (v2 | v2)}
            _ = v55
            _ = v56
        }
        _ = fi51
    }
    v57 := []int{32768, -213, -9223372036854775808, 0}
    _ = len(v57)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v57:", len(v57), v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v57
}
