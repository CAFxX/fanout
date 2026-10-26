package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 int }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int }

func main() {
    var v1 int32 = 235
    var v2 uint8 = 1
    var v3 uint8 = v2
    v4 := make(map[uint32]string)
    var v5 int16 = -1
    var v6 uint32 = 2147483648
    v2 = v2
    var v7 bool = (-432 <= 127)
    v7 = (251 >= (v5 & (v5 / ((int16(v6)) | 1))))
    v2 -= (v3 ^ 128)
    switch 1 {
    case -2147483648:
        _ = len(v4)
    case -111:
        var v8 uint8 = (v2 + (v2 | v3))
        _ = v8
    case 310:
        var v9 uint8 = (v3 * v2)
        _ = v9
    default:
        if false {
            switch 9223372036854775807 {
            case -186:
                var v10 int = 474
                _ = v10
            case -413:
                var v11 uint64 = 4294967295
                _ = v11
            case 175:
                for fi12 := 0; fi12 < 6; fi12++ {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        if true {
                            v14 := []int{}
                            var v15 uint16 = 921
                            _ = v14
                            _ = v15
                        }
                        v16 := []int{(fi13 ^ 1), fi13}
                        if v7 {
                            switch "hello" {
                            case "hello":
                                var v17 string = "foo bar"
                                _ = v17
                            case "\n\t":
                                switch "" {
                                case "foo bar":
                                    var v18 int32 = -2147483648
                                    var v19 uint8 = 2
                                    _ = len(v4)
                                    _ = v18
                                    _ = v19
                                case "":
                                    v20 := GenS1{F0: ("a" + "")}
                                    _ = v20
                                case "hello":
                                    var v21 int16 = v5
                                    _ = v21
                                }
                            }
                            var v22 int64 = 0
                            _ = v22
                        }
                        v3 = v3
                        v5 = v5
                        _ = fi13
                        _ = v16
                    }
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v1 -= v1
                        _ = fi23
                    }
                    _ = fi12
                }
            }
            switch "" {
            case "hello":
                v1 -= int32(v6)
            case "":
                var v24 bool = (!v7)
                _ = v24
            default:
                if ((v1 * 42) > (^(v1))) {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v26 := GenS2{F0: 2, F1: -1}
                        v1 = 42
                        v27 := GenS2{F0: 666, F1: -1}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v28 := []int{127, 1, -2147483648}
                    _ = v28
                } else {
                    switch -500 {
                    case -2147483648:
                        if (false && v7) {
                            if false {
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    for fi30 := 0; fi30 < 15; fi30++ {
                                        v31 := GenS3{F0: "foo bar", F1: v2, F2: fi29}
                                        switch fi30 {
                                        case -928:
                                            if true {
                                                switch (fi30 / ((fi30) | 1)) {
                                                case -661:
                                                    if v7 {
                                                        delete(v4, 2)
                                                    }
                                                case -797:
                                                    v2 = (^(v3))
                                                    v5 -= (v5 | (v5 % ((338) | 1)))
                                                }
                                                v3 += v2
                                                v6 -= v6
                                            } else {
                                                v32 := []int{fi29, fi30, (fi30 + 1), fi30}
                                                _ = v32
                                            }
                                        case 32767:
                                            v1 += v1
                                        default:
                                            v6 = (v6 + (^(v6)))
                                        }
                                        var v33 uint8 = 74
                                        _ = fi30
                                        _ = v31
                                        _ = v33
                                    }
                                    v34 := []int{4294967295, (fi29 / ((fi29) | 1)), (fi29 >> ((uint(v6)) & 63)), fi29}
                                    var v35 int8 = 100
                                    _ = fi29
                                    _ = v34
                                    _ = v35
                                }
                            } else {
                                if v7 {
                                    var v36 uint = 588
                                    _ = v36
                                } else {
                                    var v37 uint8 = (v2 % ((v3) | 1))
                                    v7 = (v7 && (v2 > (v37 & (^(v37)))))
                                    _ = v37
                                }
                                v1 = 326
                            }
                            v38 := v4[0]
                            _ = v38
                        }
                    default:
                        v39 := []int{-528, -596, 32768, 32768}
                        _ = v39
                    }
                    fmt.Println(v3)
                }
            }
            if (!true) {
                var v40 float64 = 1e10
                var v41 int16 = int16(v3)
                _ = v40
                _ = v41
            }
        }
    }
    if (!false) {
        v42 := []int{300, 960, 0}
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := GenS4{F0: ("foo bar" + ("hello" + "\n\t")), F1: 0}
            v5 += (v5 - v5)
            fmt.Println(v5)
            _ = fi43
            _ = v44
        }
        _ = v42
    }
    v7 = false
    fmt.Println(v7, v3, v7)
    v4[v6] = "a"
    switch 2147483647 {
    case 42:
        fmt.Println(v3, v2, v6)
    case 128:
        if ((v5 + int16(v6)) < (v5 >> ((uint(v3)) & 63))) {
            var v45 string = "\n\t"
            v3 = v2
            _ = v45
        } else {
            v7 = v7
            v46 := GenS5{F0: 4294967295, F1: -1}
            _ = v46
        }
    case 112:
        fmt.Println(v3)
    }
    switch -9223372036854775808 {
    case 697:
        if v7 {
            v2 = 1
        }
    }
    v47 := GenS1{F0: ("foo bar" + ("hello" + "\n\t"))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
