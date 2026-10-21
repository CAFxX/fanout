package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int16; F2 string }
type GenS6 struct { F0 string; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint64 }
func genFn0() bool {
    var v7 int8 = -1
    var v8 bool = (!true)
    v7 = v7
    var v9 float32 = 3.14159
    v10 := []int{127, 128}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 1
    var v3 int8 = 100
    var v4 int8 = (-(v3))
    var v5 string = "hello"
    v6 := make(map[uint32]int)
    v11 := GenS1{F0: 42, F1: 598, F2: -1}
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int16 = int16(v3)
        v14 := GenS2{F0: (^(fi12))}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float32 = -2.25
    var v16 int32 = 686
    var v17 uint32 = 911
    var v18 uint16 = 0
    v4 = (v4 | (v3 ^ v3))
    v19 := []int{128, -1, -9223372036854775808}
    if len(v19) > 0 {
        v19[0] = 9223372036854775807
    }
    var v20 int = int(v18)
    switch v20 {
    case 273:
        var v21 bool = false
        _ = v21
    case 595:
        if false {
            var v22 uint64 = 425
            _ = v22
        }
    case 0:
        fmt.Println(v15, v17)
    default:
        v2 += v2
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        switch "hello" {
        case "foo bar":
            var v24 int16 = 0
            _ = v24
        case "hello":
            v2 += v2
        }
        v5 = (v5 + "hello")
        _ = fi23
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        switch (fi25 & -787) {
        case 337:
            v26 := GenS3{F0: fi25}
            _ = v26
        default:
            v19 = append(v19, 127)
        }
        if false {
            v4 = (v3 * (v4 + v3))
            v2 -= (^(v2))
        }
        _ = fi25
    }
    switch (v20 | v20) {
    case 4294967295:
        if (!(v4 != 1)) {
            delete(v6, v17)
            switch "foo bar" {
            case "foo bar":
                var v27 uint64 = 1
                _ = v27
            }
            var v28 int32 = (v16 ^ v16)
            _ = v28
        } else {
            var v29 float64 = 1.5
            for fi30 := 0; fi30 < 14; fi30++ {
                fmt.Println(v16)
                if false {
                    v16 += v16
                    v31 := GenS2{F0: 2147483647}
                    _ = v31
                } else {
                    fmt.Println(v4, v16, v29)
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v33 := GenS4{F0: true}
                        switch fi30 {
                        case -9223372036854775808:
                            _ = len(v1)
                        case 128:
                            var v34 uint64 = 881
                            _ = v34
                        case 9223372036854775807:
                            if len(v19) > 0 {
                                v19[0] = (v20 & int(v17))
                            }
                        }
                        if (0.1 <= (-2.25 + (float64(v15) - v29))) {
                            v20 = (v20 + 128)
                            v4 = -110
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        } else {
                            fmt.Println(v2, v15, v2)
                            v35 := GenS5{F0: 1, F1: -32768, F2: "a"}
                            fmt.Println(v18, v2, fi32)
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                var v36 int16 = 427
                _ = fi30
                _ = v36
            }
            var v37 uint = 18446744073709551615
            _ = v29
            _ = v37
        }
    }
    v5 = ""
    switch 4294967295 {
    case 127:
        v20 = v20
    case 32767:
        delete(v6, (v17 + (v17 + (v17 << ((502) & 63)))))
    case -268:
        if (0 > 312) {
            var v38 int16 = -237
            v39 := GenS1{F0: 1, F1: 2, F2: -235}
            var v40 float32 = v15
            _ = v38
            _ = v39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        if ((v3 << ((18446744073709551615) & 63)) == v3) {
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 int64 = 653
                v44 := GenS6{F0: v5, F1: false, F2: 42}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v45 := GenS1{F0: 2, F1: 0, F2: fi42}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            switch "foo bar" {
            case "hello":
                if false {
                    v16 = v16
                }
            case "":
                switch "hello" {
                case "a":
                    switch v5 {
                    case "\n\t":
                        if false {
                            if (0 >= 673) {
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                var v46 bool = (v20 >= (v20 + (fi41 / (((fi41 - 127)) | 1))))
                                var v47 int16 = 32767
                                _ = v46
                                _ = v47
                            }
                            fmt.Println(v20)
                        } else {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            var v48 uint32 = (v17 / ((v17) | 1))
                            _ = v48
                        }
                    case "foo bar":
                        if (!false) {
                            var v49 uint32 = (v17 * v17)
                            _ = v49
                        } else {
                            v2 += (v2 - v2)
                            if len(v19) > 0 {
                                v19[0] = fi41
                            }
                        }
                    default:
                        var v50 int = 0
                        _ = v50
                    }
                case "\n\t":
                    for fi51 := 0; fi51 < 17; fi51++ {
                        var v52 int16 = 42
                        _ = fi51
                        _ = v52
                    }
                }
            case "foo bar":
                v53 := v6[v17]
                _ = v53
            }
        }
        switch (-(v20)) {
        case 127:
            v20 = (^(fi41))
        case 32768:
            if (0.1 == 0.1) {
                var v54 int64 = -640
                for fi55 := 0; fi55 < 20; fi55++ {
                    v56 := GenS6{F0: ("foo bar" + v5), F1: false, F2: 0}
                    if (v54 > (v2 / ((v2) | 1))) {
                        delete(v1, (v5 + ((v5 + "\n\t") + "foo bar")))
                        v15 = (v15 + ((3.14159 * v15) + (0.0 - float32(float64(v15)))))
                        v57 := GenS7{F0: 0, F1: 2}
                        _ = v57
                    }
                    _ = fi55
                    _ = v56
                }
                _ = v54
            }
        default:
            for fi58 := 0; fi58 < 5; fi58++ {
                var v59 uint = 0
                _ = fi58
                _ = v59
            }
        }
        v19 = append(v19, (^(v20)))
        _ = fi41
    }
    v4 = (v4 + int8(v17))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v15
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
}
