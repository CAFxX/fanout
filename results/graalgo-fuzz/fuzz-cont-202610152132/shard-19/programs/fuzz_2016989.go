package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int }
func genFn0() uint32 {
    v5 := []int{1, 559, 1}
    var v6 uint32 = 4294967295
    var v7 float32 = 3.14159
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if (127 == 127) {
            v7 = (((((v7 - v7) + v7) - (v7 * (v7 - float32(((float64(3.14159) + -2.25) - 1.5))))) * float32((((float64(v7) / ((-2.25 / (1.5 * 0.1)) * -2.25)) / 1.0) / 0.1))) / 3.14159)
        } else {
            v6 -= (v6 & 42)
            _ = len(v5)
        }
        if (!false) {
            v6 += 65535
        }
        var v8 float64 = 0.1
        _ = v8
    }
    var v9 float32 = v7
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return v6
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = float32((v1 - (3.14159 / 3.14159)))
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    if ((false && ("foo bar" != "a")) || false) {
        if false {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v1)
                v2 += float32(float64(float32((-2.25 * v1))))
                v1 += v1
                _ = fi10
            }
            var v11 uint = 65535
            v12 := GenS1{F0: -32768, F1: false}
            _ = v11
            _ = v12
        } else {
            if false {
                v13 := []int{9223372036854775807}
                var v14 int = 127
                v14 += v14
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 8; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if false {
                            if true {
                                var v17 int64 = 9223372036854775807
                                v18 := GenS2{F0: (1 <= -2147483648), F1: 892, F2: 185}
                                var v19 int8 = -1
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        } else {
                            v1 = v1
                            v20 := []int{}
                            switch fi16 {
                            case 165:
                                fmt.Println(fi16, fi15)
                            case 32768:
                                v21 := GenS3{F0: 0}
                                _ = v21
                            default:
                                if (49 < 1) {
                                    v2 = 1.5
                                }
                            }
                            _ = v20
                        }
                        var v22 string = (("\n\t" + "a") + "foo bar")
                        switch (v22 + ((("hello" + v22) + (v22 + ("foo bar" + ("" + v22)))) + v22)) {
                        case "foo bar":
                            v23 := []int{(fi16 & fi16)}
                            _ = v23
                        }
                        fmt.Println(v22)
                        v1 += v1
                        _ = fi16
                        _ = v22
                    }
                    if true {
                        v24 := GenS3{F0: 9223372036854775807}
                        var v25 int32 = 993
                        _ = v24
                        _ = v25
                    }
                    if true {
                        var v26 int = fi15
                        var v27 float64 = v1
                        _ = v26
                        _ = v27
                    }
                    _ = fi15
                }
                var v28 int8 = -1
                _ = v28
            }
            fmt.Println(v1)
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 float64 = v1
            var v31 uint32 = 348
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 int = 32768
        _ = v32
    }
    if (32767 <= 2) {
        v4[0] = true
        if (false || (1 != 116)) {
            v33 := GenS4{F0: "\n\t", F1: 32768}
            switch 585 {
            case -1:
                var v34 string = (("hello" + "hello") + "foo bar")
                _ = v34
            }
            _ = v33
        }
    } else {
        for fi35 := 0; fi35 < 16; fi35++ {
            if (1000 <= 32767) {
                if (!true) {
                    v1 += (v1 * ((float64(0.1) - 0.1) + (((((v1 * v1) / 1.0) * (v1 * 0.0)) / float64(0.1)) / 0.1)))
                    var v36 int8 = -1
                    var v37 uint64 = 42
                    _ = v36
                    _ = v37
                } else {
                    var v38 int32 = -2147483648
                    _ = v38
                }
            }
            v1 += 3.14159
            switch "foo bar" {
            case "\n\t":
                if true {
                    v39 := []int{fi35, fi35}
                    _ = v39
                } else {
                    var v40 int8 = 127
                    v1 += (-2.25 + (0.0 / v1))
                    _ = v40
                }
            case "foo bar":
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 int = (fi35 << ((2) & 63))
                    fmt.Println(fi41, fi35, fi41)
                    v2 = float32((-2.25 + -2.25))
                    v43 := []int{-1, fi41, v42}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            default:
                v44 := GenS2{F0: false, F1: 0, F2: 0}
                _ = v44
            }
            _ = fi35
        }
    }
    v1 = float64(v2)
    v1 -= 0.0
    switch "\n\t" {
    case "a":
        v1 = v1
    }
    var v45 float32 = v2
    v46 := []int{739}
    if (!false) {
        fmt.Println(v1)
    } else {
        fmt.Println(v2, v1)
        switch 4294967295 {
        case 32767:
            v45 -= (v2 * 3.14159)
        default:
            v45 -= 1.5
        }
    }
    if true {
        for fi47 := 0; fi47 < 5; fi47++ {
            v1 = -2.25
            _ = fi47
        }
        v2 = v2
    }
    v48 := GenS2{F0: true, F1: 41, F2: 1}
    var v49 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v49
}
