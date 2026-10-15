package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int8 }

func main() {
    var v1 int16 = -820
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 int = 127
    switch "\n\t" {
    case "":
        v5 := []int{v4, v4}
        _ = v5
    case "a":
        fmt.Println(v1)
    }
    switch v4 {
    case 32767:
        fmt.Println(v4, v4)
    case 695:
        fmt.Println(v4, v4)
    case 4294967295:
        var v6 int = 818
        _ = v6
    default:
        switch ((("hello" + ("foo bar" + ("" + ""))) + "a") + "a") {
        case "foo bar":
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(fi7, v1, v4)
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 bool = (2 <= 0)
                    if v9 {
                        _ = len(v2)
                        var v10 uint64 = 519
                        var v11 float32 = float32(0.1)
                        v12 := []int{-9223372036854775808, fi8, (v4 % ((v4) | 1)), fi8}
                        v13 := GenS1{F0: (v1 * (v1 | int16(v10)))}
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
        case "hello":
            v14 := GenS1{F0: v1}
            _ = v14
        default:
            v15 := []int{v4, v4}
            _ = v15
        }
    }
    var v16 int32 = -1
    v16 = (-(v16))
    v17 := v2["a"]
    v4 = 32768
    fmt.Println(v16, v1, v17)
    if (v17 && v17) {
        fmt.Println(v17)
        switch (-(v4)) {
        case 146:
            v4 += (v4 >> ((922) & 63))
        }
        var v18 int64 = 9223372036854775807
        _ = v18
    }
    v3[128] = v17
    v4 = -44
    var v19 float32 = 0.1
    var v20 bool = v17
    fmt.Println(v17, v20, v1)
    var v21 int32 = v16
    var v22 uint8 = 128
    if true {
        for fi23 := 0; fi23 < 4; fi23++ {
            if (-9223372036854775808 != -9223372036854775808) {
                var v24 float32 = v19
                v20 = (!v17)
                _ = v24
            } else {
                v25 := GenS2{F0: 65535, F1: v4}
                _ = v25
            }
            if v20 {
                v26 := GenS2{F0: 442, F1: v4}
                var v27 uint16 = 32767
                if true {
                    v16 = (-(v16))
                    fmt.Println(v21, v16, v27)
                } else {
                    fmt.Println(v16, v21, v19)
                    if true {
                        v21 = v16
                        if (!false) {
                            for fi28 := 0; fi28 < 9; fi28++ {
                                var v29 int16 = v1
                                fmt.Println(v1)
                                v16 = 238
                                var v30 uint8 = (v22 & (v22 % ((v22) | 1)))
                                var v31 int8 = 1
                                var v32 uint16 = v27
                                if true {
                                    v22 = 127
                                    var v33 bool = false
                                    _ = v33
                                } else {
                                    v22 = v22
                                    v34 := []int{}
                                    fmt.Println(v32, v27, v16)
                                    _ = v34
                                }
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            var v35 uint32 = 65535
                            if v17 {
                                for fi36 := 0; fi36 < 15; fi36++ {
                                    v22 = 140
                                    var v37 uint16 = v27
                                    var v38 uint8 = 0
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                var v39 int = fi23
                                _ = v39
                            }
                            _ = v35
                        } else {
                            var v40 string = "\n\t"
                            v41 := GenS1{F0: -1}
                            _ = len(v2)
                            _ = v40
                            _ = v41
                        }
                        if (!v17) {
                            v42 := []int{}
                            _ = v42
                        } else {
                            if false {
                                for fi43 := 0; fi43 < 16; fi43++ {
                                    if (!v17) {
                                        v44 := []int{int(v22), (fi23 % ((fi43) | 1))}
                                        v45 := GenS3{F0: v27, F1: 42, F2: 1}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if (0.0 <= 0.0) {
                                            fmt.Println(fi23, fi23, v27)
                                            v46 := GenS3{F0: v27, F1: (-(v16)), F2: 1}
                                            var v47 float64 = -2.25
                                            _ = v46
                                            _ = v47
                                        } else {
                                            switch ("foo bar" + ("hello" + "")) {
                                            case "\n\t":
                                                var v48 uint8 = v22
                                                _ = v48
                                            case "hello":
                                                v49 := GenS4{F0: 32767, F1: (v27 ^ 255)}
                                                _ = v49
                                            case "":
                                                if (1 >= 354) {
                                                    var v50 uint = 42
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v51 int64 = 1
                                                        _ = v51
                                                    case "hello":
                                                        fmt.Println(v1)
                                                        if (v17 || (255 < v22)) {
                                                            switch "hello" {
                                                            case "hello":
                                                                var v52 int32 = v21
                                                                _ = v52
                                                            case "":
                                                                if (v20 && (v20 && (488 > 42))) {
                                                                    _ = len(v3)
                                                                    v22 += 42
                                                                } else {
                                                                    var v53 int8 = int8(v22)
                                                                    v53 = 127
                                                                    var v54 bool = false
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                            }
                                                            if v20 {
                                                                v55 := []int{fi23, 32768, fi43}
                                                                _ = v55
                                                            }
                                                            v56 := GenS2{F0: v50, F1: (-(fi23))}
                                                            _ = v56
                                                        }
                                                    default:
                                                        v57 := GenS1{F0: (v1 * v1)}
                                                        _ = v57
                                                    }
                                                    v4 = -469
                                                    _ = v50
                                                }
                                            default:
                                                var v58 int = 0
                                                _ = v58
                                            }
                                        }
                                    }
                                    _ = fi43
                                }
                                v2[("foo bar" + (("" + "hello") + ("\n\t" + (("hello" + "foo bar") + "hello"))))] = false
                            }
                            v59 := GenS5{F0: v20}
                            var v60 int = int(v22)
                            _ = v59
                            _ = v60
                        }
                    }
                    v61 := GenS6{F0: (v4 ^ (v4 / ((v4) | 1))), F1: 40}
                    _ = v61
                }
                _ = v26
                _ = v27
            }
            _ = fi23
        }
        for fi62 := 0; fi62 < 17; fi62++ {
            v1 -= (v1 >> ((1) & 63))
            v63 := []int{(fi62 - (fi62 | (v4 + (^(v4))))), (fi62 % (((v4 & (fi62 ^ (v4 * fi62)))) | 1)), (fi62 * 74), (fi62 + 287)}
            _ = fi62
            _ = v63
        }
        if (!v17) {
            v4 += 32767
        } else {
            v17 = false
        }
    } else {
        v16 = (v21 | -380)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
