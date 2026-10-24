package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int; F2 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint, p2 uint32) int16 {
    v5 := []int{127, -2147483648}
    v6 := GenS1{F0: 1, F1: 4294967295, F2: 616}
    if len(v5) > 0 {
        v5[0] = 32768
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int16 = 361
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1e10
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float64 = 1e10
        if (100 >= 42) {
            v9 := GenS2{F0: 9223372036854775807}
            _ = v9
        }
        switch "\n\t" {
        case "":
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: v1, F1: 1, F2: 35}
                if (32768 < -513) {
                    switch "foo bar" {
                    case "foo bar":
                        if (("" + "foo bar") != "a") {
                            v4 = 1.5
                        } else {
                            v12 := GenS2{F0: -1}
                            _ = v12
                        }
                    case "\n\t":
                        var v13 int32 = -2147483648
                        _ = v13
                    default:
                        var v14 int16 = (^(v1))
                        _ = v14
                    }
                    v2[1] = (-(fi7))
                    v15 := GenS3{F0: 255}
                    v1 += (v1 >> ((42) & 63))
                    var v16 int = -345
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS4{F0: 109, F1: 127, F2: 926}
                    v18 := v2[(^(fi10))]
                    v19 := []int{(fi10 + (v18 * (fi10 << ((4294967295) & 63)))), (fi7 ^ 4294967295), (v18 % (((fi10 ^ (^(fi10)))) | 1)), fi10}
                    var v20 int8 = 103
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi7
        _ = v8
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v1 += (v1 | genFn0(100, 894, 42))
        v4 = v4
        v1 += genFn0(-128, 4294967295, 2147483648)
        _ = fi21
    }
    switch -1 {
    case 127:
        for fi22 := 0; fi22 < 17; fi22++ {
            v1 -= (v1 / (((v1 & v1)) | 1))
            _ = fi22
        }
    default:
        var v23 uint32 = 493
        _ = v23
    }
    v1 += v1
    var v24 uint32 = 0
    for fi25 := 0; fi25 < 4; fi25++ {
        switch "" {
        case "hello":
            var v26 string = "\n\t"
            _ = v26
        case "foo bar":
            switch fi25 {
            case 42:
                if (float64((3.14159 + 0.0)) >= v4) {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if (false || (3.14159 > (float64((1e10 * 3.14159)) + v4))) {
                            if (65535 > 65535) {
                                v28 := GenS2{F0: 0}
                                switch (fi27 | (^(fi27))) {
                                case 42:
                                    var v29 uint = 1
                                    _ = v29
                                case -1:
                                    var v30 uint32 = v24
                                    _ = v30
                                }
                                if true {
                                    if (v1 < 1) {
                                        v24 = v24
                                    } else {
                                        v31 := GenS5{F0: 65535, F1: (fi27 << ((58) & 63)), F2: 42}
                                        v1 += 0
                                        _ = v31
                                    }
                                }
                                _ = v28
                            } else {
                                if (852 < 1) {
                                    v32 := GenS6{F0: (v1 & -324)}
                                    _ = v32
                                }
                            }
                            v24 -= 42
                            v4 = (((v4 - (float64(1e10) + v4)) + 1e10) / v4)
                        } else {
                            _ = len(v3)
                        }
                        var v33 uint32 = 695
                        _ = fi27
                        _ = v33
                    }
                    var v34 float32 = float32(float64(1.5))
                    switch fi25 {
                    case 127:
                        v35 := GenS3{F0: 0}
                        _ = v35
                    }
                    _ = v34
                }
            case 9223372036854775807:
                if (!(42 < 42)) {
                    v36 := v3[(v24 + (^(v24)))]
                    var v37 string = (("a" + "a") + "\n\t")
                    if true {
                        var v38 string = "foo bar"
                        v39 := GenS1{F0: genFn0(-1, 2, v24), F1: 2, F2: 810}
                        for fi40 := 0; fi40 < 16; fi40++ {
                            var v41 string = "a"
                            _ = fi40
                            _ = v41
                        }
                        _ = v38
                        _ = v39
                    } else {
                        if true {
                            _ = len(v3)
                            var v42 int64 = 254
                            _ = v42
                        }
                    }
                    _ = v36
                    _ = v37
                } else {
                    _ = len(v2)
                    for fi43 := 0; fi43 < 6; fi43++ {
                        v4 += (float64(float32(float64(1.5))) + 1.5)
                        var v44 int32 = 0
                        _ = fi43
                        _ = v44
                    }
                    v24 -= (^(v24))
                }
            }
        case "a":
            v45 := GenS7{F0: 1}
            _ = v45
        }
        if false {
            v46 := v2[(fi25 & (fi25 % ((-724) | 1)))]
            delete(v3, 2147483647)
            _ = v46
        }
        _ = fi25
    }
    switch "hello" {
    case "\n\t":
        var v47 string = "a"
        _ = v47
    case "hello":
        for fi48 := 0; fi48 < 5; fi48++ {
            switch 1 {
            case -586:
                v4 = v4
            case -2147483648:
                var v49 uint64 = 1
                _ = v49
            case -1:
                if (42 > int64(fi48)) {
                    v50 := GenS6{F0: (v1 / (((v1 & (-(v1)))) | 1))}
                    for fi51 := 0; fi51 < 15; fi51++ {
                        if ((129 <= 471) || false) {
                            var v52 uint64 = 0
                            _ = v52
                        } else {
                            v53 := GenS5{F0: (^(v24)), F1: fi48, F2: 2147483647}
                            _ = v53
                        }
                        fmt.Println(v1)
                        v4 -= v4
                        _ = fi51
                    }
                    _ = v50
                }
            }
            _ = fi48
        }
    }
    v1 = (v1 / ((genFn0(-1, 160, v24)) | 1))
    v4 += (1e10 - ((((float64((-2.25 - float32(0.1))) - (v4 - 3.14159)) * (((float64(0.0) * v4) - float64(-2.25)) - (((v4 - (v4 + (v4 * -2.25))) - v4) / v4))) + (((3.14159 + v4) * v4) / ((3.14159 - v4) / 3.14159))) * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
