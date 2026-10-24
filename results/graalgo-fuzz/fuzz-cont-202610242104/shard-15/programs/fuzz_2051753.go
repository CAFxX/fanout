package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int32; F2 int8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 uint = 42
    var v2 int = 32768
    var v3 uint32 = 2147483648
    v4 := []int{v2}
    var v5 string = ("" + "hello")
    var v6 uint8 = 128
    v1 -= (^(v1))
    if true {
        v5 = v5
        v4 = append(v4, (^(v2)))
    } else {
        fmt.Println(v1, v6)
        switch (v2 + v2) {
        case 128:
            v7 := GenS1{F0: 42}
            _ = v7
        }
    }
    switch 192 {
    case -382:
        var v8 int16 = 181
        _ = v8
    default:
        v9 := GenS2{F0: 1, F1: v2}
        _ = v9
    }
    v2 += v2
    if len(v4) > 0 {
        v4[0] = v2
    }
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 bool = ((false || (341 <= 2147483647)) && (v1 >= (v1 >> ((v1) & 63))))
                v13 := GenS3{F0: v12, F1: 65535, F2: 593}
                if (!(fi10 >= int(v1))) {
                    var v14 float32 = 1.5
                    var v15 uint8 = v6
                    v6 += 127
                    _ = v14
                    _ = v15
                } else {
                    switch fi11 {
                    case 0:
                        v2 += -894
                    }
                    var v16 uint16 = uint16(v1)
                    _ = v16
                }
                if ((v1 << ((v1) & 63)) > (v1 - v1)) {
                    fmt.Println(fi10)
                }
                if (42 <= 1) {
                    v4 = append(v4, (fi11 << (((v1 | v1)) & 63)))
                    switch v5 {
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    case "":
                        v17 := GenS1{F0: 0}
                        _ = v17
                    case "foo bar":
                        v4 = append(v4, fi10)
                    default:
                        fmt.Println(v1)
                    }
                    if v12 {
                        v18 := GenS4{F0: 1, F1: -970}
                        if (-1 == -100) {
                            var v19 int64 = 0
                            if (65535 < 32767) {
                                var v20 int32 = -938
                                var v21 uint64 = 2
                                v22 := GenS4{F0: -23, F1: 32767}
                                var v23 uint8 = (v6 - uint8(v3))
                                _ = len(v4)
                                _ = v20
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS5{F0: -76}
                                _ = v24
                            }
                            _ = v19
                        }
                        _ = v18
                    }
                } else {
                    switch (fi10 & (v2 % ((fi11) | 1))) {
                    case 381:
                        v1 -= 2
                    default:
                        switch v5 {
                        case "\n\t":
                            var v25 float32 = 1.5
                            _ = v25
                        default:
                            v4 = append(v4, 127)
                        }
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            var v26 bool = true
            v5 = (("\n\t" + v5) + v5)
            _ = fi10
            _ = v26
        }
        switch v2 {
        case 128:
            for fi27 := 0; fi27 < 20; fi27++ {
                v2 += fi27
                _ = fi27
            }
        default:
            switch v5 {
            case "foo bar":
                switch (v2 * (v2 ^ v2)) {
                case 42:
                    v1 -= v1
                }
            case "\n\t":
                fmt.Println(v6, v6, v5)
            case "hello":
                if true {
                    var v28 uint64 = 525
                    var v29 uint32 = 42
                    _ = v28
                    _ = v29
                } else {
                    switch 685 {
                    case 2147483647:
                        for fi30 := 0; fi30 < 20; fi30++ {
                            switch 0 {
                            case -2147483648:
                                _ = len(v4)
                            }
                            if (true && (1 != 0)) {
                                var v31 uint64 = 2
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    var v33 int32 = 0
                                    _ = fi32
                                    _ = v33
                                }
                                fmt.Println(v1)
                                _ = v31
                            } else {
                                if (!true) {
                                    v34 := GenS4{F0: 42, F1: -336}
                                    switch -523 {
                                    case -1:
                                        var v35 int64 = -547
                                        _ = v35
                                    }
                                    v36 := GenS3{F0: (1 == 0), F1: 559, F2: 860}
                                    _ = v34
                                    _ = v36
                                }
                            }
                            _ = fi30
                        }
                    default:
                        fmt.Println(v6)
                    }
                    var v37 uint32 = (v3 & (^(v3)))
                    if len(v4) > 0 {
                        v4[0] = v2
                    }
                    _ = v37
                }
            }
        }
        var v38 int = v2
        _ = v38
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v39 uint16 = 42
        if len(v4) > 0 {
            v4[0] = v2
        }
        _ = v39
    }
    var v40 bool = ((!true) || (1e10 >= -2.25))
    switch v5 {
    case "a":
        if ((0.0 * 0.1) == -2.25) {
            if (!v40) {
                var v41 int32 = -2147483648
                _ = v41
            }
        } else {
            for fi42 := 0; fi42 < 12; fi42++ {
                for fi43 := 0; fi43 < 13; fi43++ {
                    v1 = (v1 | (v1 * v1))
                    v44 := GenS6{F0: 0, F1: 187, F2: 100}
                    var v45 uint = v1
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                var v46 float32 = float32(0.1)
                _ = fi42
                _ = v46
            }
        }
    }
    v4 = append(v4, (-(v2)))
    var v47 uint8 = 42
    fmt.Println(v3, v5, v2)
    if ((v6 | v6) >= v6) {
        v1 = (v1 % (((^(v1))) | 1))
        var v48 uint16 = 255
        _ = v48
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        if v40 {
            fmt.Println(v6, v5)
        } else {
            v50 := GenS7{F0: v3}
            var v51 int64 = 42
            if false {
                if len(v4) > 0 {
                    v4[0] = fi49
                }
                v4 = append(v4, 32768)
                fmt.Println(v51, v47, v40)
            }
            _ = v50
            _ = v51
        }
        _ = fi49
    }
    for fi52 := 0; fi52 < 8; fi52++ {
        v6 -= v47
        v53 := GenS5{F0: 0}
        _ = fi52
        _ = v53
    }
    if (v40 || true) {
        v54 := GenS2{F0: 505, F1: 2147483647}
        v4 = append(v4, 32767)
        _ = v54
    } else {
        for fi55 := 0; fi55 < 14; fi55++ {
            fmt.Println(v47, v5, v6)
            var v56 int16 = -798
            _ = fi55
            _ = v56
        }
    }
    v4 = append(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
