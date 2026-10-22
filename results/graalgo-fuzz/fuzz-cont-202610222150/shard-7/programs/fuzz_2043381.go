package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int16; F1 int8; F2 uint16 }
func genFn0() int8 {
    v4 := GenS1{F0: -1, F1: 0}
    v5 := GenS2{F0: 371}
    switch "foo bar" {
    case "a":
        var v6 uint16 = 32767
        _ = v6
    case "hello":
        var v7 string = ("\n\t" + "a")
        _ = v7
    case "":
        var v8 uint8 = 255
        _ = v8
    default:
        var v9 string = "foo bar"
        _ = v9
    }
    _ = v4
    _ = v5
    return -19
}

func main() {
    var v1 float64 = (1.5 / (-2.25 + 0.1))
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v10 string = "a"
    var v11 bool = false
    for fi12 := 0; fi12 < 15; fi12++ {
        v2[4294967295] = (fi12 + 200)
        v13 := GenS3{F0: 299, F1: 2147483647}
        switch ("hello" + v10) {
        case "hello":
            v11 = v11
        case "foo bar":
            v14 := GenS4{F0: 0}
            _ = v14
        case "":
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 float32 = 1.5
                v10 = (v10 + "hello")
                v16 -= float32(((v1 / v1) + (v1 / (-2.25 * 3.14159))))
                _ = fi15
                _ = v16
            }
        }
        _ = fi12
        _ = v13
    }
    var v17 uint8 = 2
    if v11 {
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := []int{}
            v17 = (v17 / (((v17 % (((v17 & v17)) | 1))) | 1))
            _ = fi18
            _ = v19
        }
        v11 = v11
    }
    var v20 bool = v11
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if v20 {
                v23 := GenS5{F0: (fi22 | -750), F1: 860, F2: 118}
                switch "hello" {
                case "\n\t":
                    var v24 uint8 = (v17 - (v17 - (^(v17))))
                    _ = v24
                case "":
                    _ = len(v2)
                default:
                    switch fi22 {
                    case 42:
                        if true {
                            switch fi22 {
                            case 4294967295:
                                fmt.Println(v1, v11)
                            }
                            v2[2147483647] = (fi21 ^ 2147483647)
                        }
                    case 662:
                        v17 = 94
                    case 128:
                        if (!(-2147483648 >= 42)) {
                            switch (fi21 * (fi22 ^ (fi22 - -9223372036854775808))) {
                            case 32768:
                                v11 = (("foo bar" + v10) >= ("foo bar" + v10))
                            case -9223372036854775808:
                                v17 += (v17 << ((497) & 63))
                            case 443:
                                var v25 int64 = 0
                                _ = v25
                            default:
                                v11 = (3.14159 < 0.1)
                            }
                        } else {
                            var v26 uint32 = 2
                            v2[(^(v26))] = (fi22 & -932)
                            _ = v26
                        }
                    default:
                        var v27 uint32 = 492
                        _ = v27
                    }
                }
                _ = v23
            }
            if (32767 > 1000) {
                if (("foo bar" < v10) && ((false || (0.0 >= (0.0 * 1e10))) && v20)) {
                    v28 := GenS4{F0: v17}
                    v10 = ((v10 + v10) + "a")
                    v29 := GenS6{F0: 18446744073709551615, F1: 559}
                    _ = v28
                    _ = v29
                } else {
                    switch (fi21 / (((-(fi22))) | 1)) {
                    case -1:
                        if (true && v20) {
                            switch v10 {
                            case "a":
                                v30 := v2[65535]
                                v31 := GenS5{F0: fi22, F1: 42, F2: -128}
                                v32 := GenS4{F0: v17}
                                _ = v30
                                _ = v31
                                _ = v32
                            default:
                                var v33 float32 = 0.1
                                _ = v33
                            }
                            v17 += (^(v17))
                        }
                    }
                }
                if (42 != 1) {
                    v2[2147483648] = fi21
                    v17 -= (^(v17))
                    v11 = ((119 < v17) || (true && (456 < 42)))
                }
                fmt.Println(v11, fi22)
            }
            v34 := GenS5{F0: 107, F1: 601, F2: 119}
            v11 = v11
            _ = len(v3)
            if v11 {
                var v35 string = v10
                var v36 int16 = 1000
                v35 = v35
                _ = v35
                _ = v36
            } else {
                if (false && false) {
                    v10 = v10
                    v20 = true
                }
            }
            if v20 {
                var v37 uint64 = 0
                _ = v37
            }
            v38 := GenS7{F0: 663, F1: 100, F2: 2}
            _ = fi22
            _ = v34
            _ = v38
        }
        _ = fi21
    }
    if false {
        var v39 int16 = 0
        _ = v39
    } else {
        var v40 uint32 = 2
        _ = v40
    }
    var v41 float64 = v1
    var v42 int16 = -883
    v43 := v3[-1]
    var v44 uint64 = 4294967295
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
