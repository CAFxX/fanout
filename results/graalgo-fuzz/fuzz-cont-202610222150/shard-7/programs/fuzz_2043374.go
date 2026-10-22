package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 string = "hello"
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32((0.1 - 1.5))
    var v5 uint64 = 979
    var v6 float32 = 0.1
    fmt.Println(v1, v6, v6)
    fmt.Println(v5, v5)
    v1 = v1
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(-(fi7)), 1, fi7, -2147483648}
        _ = fi7
        _ = v8
    }
    switch 127 {
    case 2147483647:
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float64 = (0.0 + 3.14159)
            _ = fi9
            _ = v10
        }
    default:
        v11 := []int{}
        _ = v11
    }
    if ((v2 - (v2 % (((^(v2))) | 1))) >= (^(v2))) {
        var v12 int8 = 100
        var v13 uint = 42
        _ = v12
        _ = v13
    }
    fmt.Println(v6)
    delete(v3, 2)
    var v14 int16 = -918
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := []int{(fi15 - fi15), (fi15 % ((fi15) | 1))}
        v17 := GenS1{F0: 1, F1: int64(v5)}
        v6 = (0.1 / -2.25)
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 0
    v3[142] = 794
    if (!((688 <= -9223372036854775808) || true)) {
        switch "" {
        case "a":
            v19 := v3[696]
            _ = v19
        case "":
            if true {
                var v20 bool = (65535 == 4294967295)
                if (-9223372036854775808 >= 9223372036854775807) {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := []int{fi21}
                        if v20 {
                            for fi23 := 0; fi23 < 2; fi23++ {
                                var v24 float32 = (v4 + (v4 * ((v4 * v4) / ((float32(0.1) - 1.5) - v6))))
                                var v25 float64 = 1.5
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v5 -= 4294967295
                        } else {
                            if v20 {
                                v20 = (2 >= 203)
                                v26 := GenS1{F0: 1, F1: 1}
                                v20 = (v20 && (!(((1e10 * 1.5) / 0.1) == -2.25)))
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 17; fi27++ {
                                switch fi27 {
                                case 1:
                                    v4 = v4
                                case 128:
                                    var v28 bool = false
                                    _ = v28
                                case -9223372036854775808:
                                    v3[653] = 4294967295
                                }
                                v4 += float32(float64((float32(float64((float32(-2.25) * (v6 * v4)))) + v4)))
                                _ = fi27
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    if (310 < 281) {
                        v29 := v3[633]
                        v30 := GenS2{F0: v1, F1: 42}
                        var v31 bool = v20
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
                if true {
                    switch (((v1 + v1) + ("" + (v1 + ("foo bar" + "hello")))) + v1) {
                    case "hello":
                        if v20 {
                            var v32 string = (((((v1 + v1) + v1) + (v1 + v1)) + (v1 + v1)) + v1)
                            for fi33 := 0; fi33 < 3; fi33++ {
                                if v20 {
                                    v14 -= v14
                                } else {
                                    v34 := []int{int(v18), 721, (^(fi33))}
                                    v4 = -2.25
                                    _ = v34
                                }
                                if v20 {
                                    switch fi33 {
                                    case -2147483648:
                                        v6 -= (v4 - (-2.25 / 1.5))
                                    case 2147483647:
                                        var v35 int = 0
                                        _ = v35
                                    default:
                                        if ((0.0 * (float32(1.5) - (v4 * float32(1.5)))) > (v6 * (v4 - v4))) {
                                            var v36 string = "\n\t"
                                            if (!(!(v20 && v20))) {
                                                var v37 float64 = 0.0
                                                v4 = 1e10
                                                v20 = ((-1 < -12) && (!(v20 || v20)))
                                                _ = v37
                                            } else {
                                                var v38 int16 = v14
                                                v32 = v1
                                                _ = v38
                                            }
                                            _ = v36
                                        }
                                    }
                                    var v39 int16 = 1000
                                    v40 := []int{(-(fi33))}
                                    _ = v39
                                    _ = v40
                                } else {
                                    v41 := []int{fi33, 453}
                                    _ = v41
                                }
                                _ = fi33
                            }
                            v42 := []int{127, 377, 4294967295, -527}
                            _ = v32
                            _ = v42
                        } else {
                            v43 := GenS1{F0: -63, F1: -1}
                            var v44 bool = (0 >= 0)
                            _ = v43
                            _ = v44
                        }
                    case "a":
                        switch "hello" {
                        case "":
                            v45 := GenS2{F0: "foo bar", F1: 0}
                            _ = v45
                        case "hello":
                            v46 := GenS3{F0: 42, F1: 128, F2: (1 != v14)}
                            _ = v46
                        case "\n\t":
                            if (1.5 >= float64(v6)) {
                                v1 = v1
                                if v20 {
                                    for fi47 := 0; fi47 < 2; fi47++ {
                                        v18 -= (v18 << ((612) & 63))
                                        var v48 int16 = (^(v14))
                                        v6 -= v6
                                        _ = fi47
                                        _ = v48
                                    }
                                }
                            } else {
                                v14 -= (v14 & v14)
                                var v49 int16 = v14
                                switch 128 {
                                case 516:
                                    if v20 {
                                        fmt.Println(v2)
                                        v50 := GenS4{F0: 2147483647, F1: 2147483648, F2: v18}
                                        _ = v50
                                    } else {
                                        v1 = "hello"
                                    }
                                default:
                                    v51 := GenS5{F0: (v14 & v49)}
                                    _ = v51
                                }
                                _ = v49
                            }
                        }
                    case "":
                        var v52 uint8 = (^(v2))
                        _ = v52
                    }
                }
                _ = v20
            } else {
                v3[2] = 105
                v53 := GenS5{F0: (v14 + v14)}
                for fi54 := 0; fi54 < 4; fi54++ {
                    for fi55 := 0; fi55 < 15; fi55++ {
                        var v56 uint = 65535
                        if false {
                            var v57 uint64 = (v5 & 0)
                            _ = v57
                        } else {
                            v18 -= 700
                            var v58 bool = ("a" < "")
                            if (((-2.25 * 1.5) / 0.1) == float64((v4 - 3.14159))) {
                                if true {
                                    _ = len(v3)
                                } else {
                                    var v59 uint16 = v18
                                    if (4294967295 > v5) {
                                        var v60 int16 = v14
                                        var v61 uint = v56
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v62 := GenS5{F0: 1}
                                        _ = v62
                                    }
                                    var v63 uint16 = v18
                                    _ = v59
                                    _ = v63
                                }
                            } else {
                                if v58 {
                                    v4 = v4
                                }
                                v3[4294967295] = 2
                                if true {
                                    _ = len(v3)
                                    v58 = v58
                                } else {
                                    v18 += (v18 ^ (v18 | v18))
                                    v64 := []int{(fi54 >> (((v56 / ((569) | 1))) & 63)), -449, int(v18), 32767}
                                    _ = v64
                                }
                                v4 = (1.5 / 1.0)
                                if true {
                                    if false {
                                        var v65 uint = (v56 | v56)
                                        _ = v65
                                    } else {
                                        v66 := GenS2{F0: v1, F1: 2147483648}
                                        fmt.Println(v14, v56)
                                        _ = v66
                                    }
                                    var v67 uint16 = (v18 * v18)
                                    v68 := GenS6{F0: 800, F1: (true && v58), F2: (-1 > -99)}
                                    _ = v67
                                    _ = v68
                                }
                            }
                            _ = v58
                        }
                        _ = fi55
                        _ = v56
                    }
                    _ = fi54
                }
                _ = v53
            }
        case "\n\t":
            fmt.Println(v14, v6, v2)
        }
        v1 = (v1 + "\n\t")
    } else {
        v3[0] = 0
        var v69 uint16 = (v18 & v18)
        _ = v69
    }
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
