package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    var v3 uint8 = 128
    var v4 bool = true
    v5 := make(map[int]uint32)
    var v6 int8 = -83
    fmt.Println(v2, v4, v6)
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v3)
            v9 := GenS1{F0: 255}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: v3}
    if (!v4) {
        fmt.Println(v4, v2)
        if (1 >= 32767) {
            v11 := GenS1{F0: v3}
            v12 := GenS2{F0: 2, F1: 410, F2: false}
            v13 := GenS1{F0: (v3 ^ (v3 >> ((650) & 63)))}
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 float32 = 0.1
        _ = v14
    } else {
        v3 = v3
        v15 := GenS3{F0: v3, F1: (v6 ^ 100)}
        _ = v15
    }
    if v4 {
        fmt.Println(v4, v6, v2)
    }
    var v16 float64 = (1e10 / 1e10)
    var v17 int8 = v6
    v18 := GenS2{F0: 42, F1: 1, F2: (v4 && v4)}
    v19 := GenS4{F0: -693}
    var v20 float32 = 0.0
    v3 += 127
    if v4 {
        if (1 == (^(v3))) {
            var v21 float32 = 3.14159
            for fi22 := 0; fi22 < 20; fi22++ {
                v2 = (v2 & v2)
                v2 = (v2 * v2)
                _ = fi22
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                if v4 {
                    switch fi23 {
                    case 32767:
                        var v24 int16 = (v2 / (((-(v2))) | 1))
                        _ = v24
                    case -9223372036854775808:
                        switch "a" {
                        case "":
                            for fi25 := 0; fi25 < 5; fi25++ {
                                v26 := GenS5{F0: 65535}
                                v3 -= (v3 << ((255) & 63))
                                var v27 int64 = 9223372036854775807
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                    case 1:
                        v3 += (v3 / ((2) | 1))
                    default:
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v21 -= (float32((float64(((-2.25 * ((v20 * float32(float64((0.1 + v21)))) + 0.0)) / 1.0)) / 3.14159)) + 3.14159)
                            _ = fi28
                        }
                    }
                    v29 := []int{fi23, fi23}
                    _ = v29
                } else {
                    v16 = float64(((v20 / v21) - v21))
                    switch "\n\t" {
                    case "\n\t":
                        if (v16 >= float64(-2.25)) {
                            if v4 {
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    var v31 int = (-(fi30))
                                    v32 := GenS1{F0: v3}
                                    switch int(v2) {
                                    case -2147483648:
                                        v33 := v5[(fi23 ^ (v31 - (^(fi23))))]
                                        _ = v33
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                var v34 uint32 = 42
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 5; fi35++ {
                                if (1 == 4294967295) {
                                    switch (fi23 * fi35) {
                                    case 32768:
                                        v1[2] = 0
                                    case 1:
                                        switch "a" {
                                        case "\n\t":
                                            v36 := []int{(fi23 & -998), fi23, 4294967295, 127}
                                            _ = v36
                                        case "hello":
                                            v17 = (^(v17))
                                        }
                                    case -2147483648:
                                        v21 = v21
                                        switch (fi35 << ((42) & 63)) {
                                        case -9223372036854775808:
                                            v6 += v17
                                        case 9223372036854775807:
                                            v21 = ((float32((v16 - v16)) - (v20 / -2.25)) * v21)
                                        }
                                    }
                                }
                                var v37 int32 = -852
                                _ = fi35
                                _ = v37
                            }
                        }
                    default:
                        switch 1 {
                        case 2147483647:
                            v6 = v6
                        case 551:
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    switch (("a" + ((("foo bar" + "foo bar") + ("hello" + "")) + "a")) + "a") {
                                    case "foo bar":
                                        switch (fi23 % (((fi23 & fi23)) | 1)) {
                                        case -1:
                                            fmt.Println(fi23, v17)
                                        case 340:
                                            var v38 uint16 = uint16(v3)
                                            _ = v38
                                        case 32768:
                                            v17 = -1
                                        }
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "":
                                            v39 := GenS5{F0: 18446744073709551615}
                                            _ = v39
                                        }
                                    }
                                case "hello":
                                    var v40 int32 = -1
                                    _ = v40
                                case "foo bar":
                                    var v41 int8 = (v17 << ((uint(v3)) & 63))
                                    _ = v41
                                default:
                                    v42 := GenS5{F0: 18446744073709551615}
                                    _ = v42
                                }
                            }
                        case 32768:
                            for fi43 := 0; fi43 < 10; fi43++ {
                                v44 := v5[42]
                                _ = fi43
                                _ = v44
                            }
                        }
                    }
                }
                _ = fi23
            }
            _ = v21
        } else {
            v45 := GenS2{F0: 4294967295, F1: 2, F2: v4}
            v46 := GenS6{F0: "\n\t", F1: v2}
            var v47 uint32 = 0
            _ = v45
            _ = v46
            _ = v47
        }
    }
    if true {
        v48 := GenS5{F0: uint(v3)}
        if false {
            v17 += (v17 << ((255) & 63))
            v49 := GenS6{F0: "a", F1: 1}
            switch 4294967295 {
            case -553:
                var v50 uint64 = 0
                _ = v50
            }
            _ = v49
        } else {
            var v51 int8 = v17
            for fi52 := 0; fi52 < 1; fi52++ {
                v53 := []int{(fi52 | (fi52 & (fi52 ^ -9223372036854775808))), fi52, -2147483648, (^(fi52))}
                var v54 float32 = v20
                for fi55 := 0; fi55 < 20; fi55++ {
                    var v56 int64 = 42
                    var v57 bool = (true && v4)
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                _ = fi52
                _ = v53
                _ = v54
            }
            v16 = 0.0
            _ = v51
        }
        _ = v48
    } else {
        for fi58 := 0; fi58 < 9; fi58++ {
            if v4 {
                for fi59 := 0; fi59 < 18; fi59++ {
                    var v60 uint32 = 1
                    v61 := GenS1{F0: (v3 >> ((2) & 63))}
                    var v62 int32 = 42
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                v3 = v3
            }
            var v63 uint8 = v3
            _ = fi58
            _ = v63
        }
    }
    var v64 uint8 = 1
    v65 := GenS7{F0: -2147483648, F1: 772}
    v2 = (v2 % (((v2 ^ (v2 & int16(v64)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
