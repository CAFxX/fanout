package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int8; F2 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int32, p1 uint16, p2 string) int8 {
    v5 := GenS1{F0: p0, F1: -726}
    p1 = 255
    p0 -= (-(p0))
    switch (p2 + "a") {
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{(fi6 | (fi6 ^ fi6)), fi6, 2147483647}
            var v8 string = "a"
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v10 uint16 = (p1 >> ((65535) & 63))
                    var v11 int32 = p0
                    _ = v10
                    _ = v11
                } else {
                    switch -475 {
                    case 128:
                        fmt.Println(fi6, p1)
                    case -767:
                        if false {
                            if false {
                                if (((127 <= 74) || false) && false) {
                                    if (!true) {
                                        fmt.Println(p2, p1)
                                        fmt.Println(p1)
                                    } else {
                                        var v12 uint32 = 176
                                        _ = v12
                                    }
                                }
                            }
                            v13 := GenS1{F0: (p0 * (-(p0))), F1: -9223372036854775808}
                            v14 := GenS2{F0: false, F1: 205, F2: true}
                            _ = v13
                            _ = v14
                        } else {
                            if false {
                                if true {
                                    if false {
                                        v15 := GenS1{F0: (-(p0)), F1: -1}
                                        switch fi9 {
                                        case 337:
                                            v16 := GenS3{F0: ""}
                                            _ = v16
                                        }
                                        _ = v15
                                    } else {
                                        fmt.Println(p2, fi9)
                                    }
                                }
                            }
                            var v17 uint32 = 2147483647
                            _ = v17
                        }
                    }
                }
                _ = fi9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case "foo bar":
        v18 := GenS1{F0: -200, F1: 9223372036854775807}
        _ = v18
    case "\n\t":
        v19 := []int{32768, 867, 4294967295}
        _ = v19
    default:
        p0 = -2147483648
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -29
}

func main() {
    var v1 uint16 = 875
    v2 := make(map[uint32]bool)
    v3 := []int{4294967295}
    var v4 float32 = -2.25
    fmt.Println(v1)
    v20 := GenS2{F0: true, F1: 0, F2: false}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int64 = -9223372036854775808
        if false {
            _ = len(v3)
            for fi23 := 0; fi23 < 13; fi23++ {
                fmt.Println(v4, v22)
                var v24 uint64 = 0
                v4 = v4
                var v25 int32 = 748
                if true {
                    fmt.Println(v22, v24)
                    _ = len(v2)
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            v22 += -1
        } else {
            v2[0] = true
            var v26 int64 = (-(v22))
            switch "hello" {
            case "foo bar":
                if true {
                    v27 := GenS3{F0: ""}
                    _ = v27
                }
            default:
                var v28 float64 = (1.5 * 1e10)
                _ = v28
            }
            _ = v26
        }
        if (!(-9223372036854775808 >= v22)) {
            var v29 uint8 = 1
            _ = v29
        }
        _ = fi21
        _ = v22
    }
    v30 := GenS2{F0: (v1 == 255), F1: 1, F2: (-9223372036854775808 > -802)}
    for fi31 := 0; fi31 < 19; fi31++ {
        v4 = v4
        _ = fi31
    }
    v1 = (^(v1))
    v4 += (((float32(float64(((float32(float64(float32(((1.5 + (3.14159 * float64((v4 - (-2.25 - float32(1.5)))))) * -2.25)))) * 0.0) + float32(3.14159)))) - float32((3.14159 - 0.1))) * 0.1) * v4)
    fmt.Println(v1)
    v32 := GenS3{F0: "hello"}
    v4 = (v4 + v4)
    switch -56 {
    case 32767:
        v33 := GenS2{F0: false, F1: -2147483648, F2: true}
        _ = v33
    case 565:
        var v34 float64 = 3.14159
        _ = v34
    case -1:
        switch -822 {
        case 32767:
            v4 = -2.25
        case -960:
            for fi35 := 0; fi35 < 13; fi35++ {
                switch (-(fi35)) {
                case 719:
                    var v36 int8 = 127
                    _ = v36
                case 2147483647:
                    var v37 int8 = -117
                    _ = v37
                }
                if true {
                    v38 := GenS2{F0: (3.14159 < (v4 * v4)), F1: -105, F2: true}
                    _ = v38
                }
                _ = fi35
            }
        case -48:
            var v39 float64 = (((-2.25 + 1e10) * 0.1) * 1.5)
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        if (!true) {
            for fi41 := 0; fi41 < 17; fi41++ {
                v3 = append(v3, fi41)
                if true {
                    var v42 int16 = -1
                    v43 := GenS4{F0: 255, F1: 1, F2: true}
                    var v44 int32 = 2147483647
                    v4 = v4
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi41
            }
            fmt.Println(v1, fi40, v1)
            v45 := GenS5{F0: v1, F1: (v1 + (^(v1)))}
            _ = v45
        }
        var v46 uint64 = 1
        _ = fi40
        _ = v46
    }
    fmt.Println(v4, v1)
    fmt.Println(v1, v4)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
