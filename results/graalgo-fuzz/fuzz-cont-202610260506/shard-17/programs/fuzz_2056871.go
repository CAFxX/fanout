package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint; F1 int64; F2 int32 }
type GenS5 struct { F0 string; F1 uint16; F2 uint }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 uint; F1 int32; F2 int32 }
type GenS8 struct { F0 bool; F1 int16 }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 float64 = ((1e10 - (1e10 * 1.5)) + ((1.5 - 1.5) + -2.25))
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: (v3 ^ 1)}
    v6 := GenS2{F0: 9223372036854775807}
    if (127 <= 1) {
        v7 := GenS3{F0: v3, F1: ((410 < -701) && (!(false || true)))}
        _ = v7
    } else {
        v2 = (v2 / (v2 * v2))
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 string = "a"
            v9 = v9
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            switch ("foo bar" + "a") {
            case "a":
                if true {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 uint16 = 700
                        _ = fi11
                        _ = v12
                    }
                    v3 = (^(v3))
                } else {
                    v2 = float64(1.5)
                    fmt.Println(v3)
                }
            default:
                if true {
                    var v13 uint = v3
                    _ = v13
                } else {
                    v14 := GenS4{F0: v3, F1: int64(v3), F2: -589}
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(v2, fi10, fi10)
                        var v16 bool = true
                        var v17 uint8 = 1
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
            }
            var v18 float64 = v2
            v2 = float64(-2.25)
            _ = fi10
            _ = v18
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    v19 := GenS3{F0: (^(v3)), F1: (true && true)}
    v3 = (v3 - 255)
    if true {
        _ = len(v4)
        v1 = append(v1, 9223372036854775807)
        var v20 uint16 = 2
        _ = v20
    } else {
        var v21 bool = true
        _ = v21
    }
    var v22 uint32 = 1
    var v23 int = -790
    v24 := GenS1{F0: 42}
    for fi25 := 0; fi25 < 9; fi25++ {
        fmt.Println(v2)
        _ = fi25
    }
    v26 := GenS1{F0: (v3 % ((1) | 1))}
    v3 = (v3 - v3)
    v22 += (v22 + uint32(v3))
    if ("" <= "a") {
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                v29 := GenS5{F0: "\n\t", F1: 1, F2: v3}
                v2 += v2
                var v30 int8 = 0
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        if true {
            if (v3 == v3) {
                if true {
                    if (!false) {
                        v31 := GenS5{F0: "hello", F1: 42, F2: v3}
                        _ = v31
                    } else {
                        v32 := GenS6{F0: 173, F1: 1, F2: (128 == 230)}
                        if true {
                            v23 = 42
                            if (v3 == 630) {
                                v33 := GenS4{F0: v3, F1: 42, F2: -2147483648}
                                switch 9223372036854775807 {
                                case 42:
                                    v22 = v22
                                case -717:
                                    if false {
                                        v23 += (v23 + (-(v23)))
                                        v34 := GenS7{F0: v3, F1: -51, F2: -1}
                                        _ = v34
                                    } else {
                                        v22 = (v22 % (((v22 - v22)) | 1))
                                    }
                                case 128:
                                    if ("" == "") {
                                        v3 -= v3
                                        v3 -= 18446744073709551615
                                        v35 := v4[(v22 << ((uint(v22)) & 63))]
                                        _ = v35
                                    } else {
                                        var v36 bool = false
                                        _ = v36
                                    }
                                default:
                                    var v37 uint16 = 318
                                    _ = v37
                                }
                                _ = v33
                            }
                        }
                        v3 -= (v3 << ((42) & 63))
                        _ = v32
                    }
                    var v38 float32 = 3.14159
                    _ = v38
                } else {
                    switch 143 {
                    case 127:
                        for fi39 := 0; fi39 < 19; fi39++ {
                            for fi40 := 0; fi40 < 15; fi40++ {
                                var v41 int64 = 332
                                _ = fi40
                                _ = v41
                            }
                            var v42 int32 = 1
                            _ = fi39
                            _ = v42
                        }
                    case 111:
                        if false {
                            var v43 int16 = -32768
                            v4[(^(v22))] = -2147483648
                            v3 += v3
                            _ = v43
                        } else {
                            switch (v23 * 2147483647) {
                            case 0:
                                for fi44 := 0; fi44 < 19; fi44++ {
                                    var v45 int8 = -128
                                    _ = fi44
                                    _ = v45
                                }
                            case 4294967295:
                                v22 = v22
                            case 32767:
                                var v46 uint64 = 1
                                _ = v46
                            }
                        }
                    }
                    var v47 uint8 = 0
                    _ = v47
                }
            } else {
                var v48 int16 = 10
                _ = v48
            }
            if true {
                switch ("\n\t" + "") {
                case "a":
                    if true {
                        var v49 uint = v3
                        v23 = v23
                        _ = v49
                    }
                }
            } else {
                switch (v23 | (v23 - -2147483648)) {
                case 0:
                    if (!(!false)) {
                        v3 += v3
                        v50 := GenS6{F0: 2, F1: 42, F2: (9223372036854775807 >= 1)}
                        _ = v50
                    } else {
                        v23 -= 128
                        switch "a" {
                        case "a":
                            for fi51 := 0; fi51 < 7; fi51++ {
                                v2 = (1e10 + ((3.14159 * float64(((0.0 + (0.0 + (0.0 - 1.5))) - float32(((v2 * 1.5) / v2))))) / v2))
                                _ = fi51
                            }
                        case "foo bar":
                            if (112 == 127) {
                                if (v3 > v3) {
                                    v52 := v4[1]
                                    if true {
                                        fmt.Println(v22)
                                        var v53 int = -342
                                        fmt.Println(v52, v2)
                                        _ = v53
                                    } else {
                                        fmt.Println(v2)
                                        v3 += 18446744073709551615
                                    }
                                    v3 -= (^(v3))
                                    _ = v52
                                }
                                v54 := GenS8{F0: true, F1: -32768}
                                _ = v54
                            }
                        default:
                            for fi55 := 0; fi55 < 1; fi55++ {
                                v2 -= ((v2 * ((v2 - -2.25) / 1.0)) + v2)
                                for fi56 := 0; fi56 < 17; fi56++ {
                                    var v57 int32 = 1
                                    _ = len(v1)
                                    _ = fi56
                                    _ = v57
                                }
                                _ = fi55
                            }
                        }
                    }
                case -360:
                    v22 -= uint32(v3)
                case 729:
                    v58 := GenS7{F0: (^(v3)), F1: -1, F2: -1}
                    _ = v58
                }
            }
            v4[v22] = int(v3)
        } else {
            _ = len(v1)
            v2 -= -2.25
            _ = len(v4)
        }
    } else {
        for fi59 := 0; fi59 < 8; fi59++ {
            var v60 int = (v23 * v23)
            _ = fi59
            _ = v60
        }
    }
    var v61 uint = (v3 / ((v3) | 1))
    v61 -= 347
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
