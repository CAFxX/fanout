package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 bool; F2 int16 }
type GenS5 struct { F0 uint; F1 int64; F2 int8 }

func main() {
    var v1 int8 = -128
    var v2 float64 = (1e10 + 3.14159)
    var v3 uint32 = 2147483648
    v4 := []int{127}
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 1, F1: 255, F2: 1}
        switch (-(fi5)) {
        case 42:
            if (!false) {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v1 += v1
                    v8 := GenS1{F0: 237, F1: 2, F2: 42}
                    v9 := GenS2{F0: 873}
                    switch "\n\t" {
                    case "":
                        fmt.Println(fi7, fi5, v3)
                    case "foo bar":
                        v10 := GenS2{F0: -1}
                        _ = v10
                    default:
                        if true {
                            _ = len(v4)
                        } else {
                            if false {
                                v1 += -1
                            } else {
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    v11 := GenS2{F0: -9223372036854775808}
                                    _ = v11
                                default:
                                    if false {
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = fi5
                                            }
                                            switch fi7 {
                                            case 32768:
                                                var v12 uint8 = 255
                                                _ = v12
                                            case 541:
                                                var v13 int8 = (v1 & v1)
                                                _ = v13
                                            case 395:
                                                v14 := GenS3{F0: v3}
                                                _ = v14
                                            }
                                            v2 -= v2
                                            v15 := GenS3{F0: (^(v3))}
                                            _ = v15
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                v16 := GenS4{F0: 42, F1: (-704 >= -824), F2: 0}
                                                _ = v16
                                            }
                                            v1 += 127
                                        }
                                        var v17 uint = 89
                                        switch 155 {
                                        case -1:
                                            var v18 uint64 = uint64(v3)
                                            _ = v18
                                        case 32767:
                                            v19 := GenS1{F0: v17, F1: 2, F2: 155}
                                            _ = v19
                                        case -9223372036854775808:
                                            var v20 int64 = -9223372036854775808
                                            _ = v20
                                        default:
                                            v21 := GenS1{F0: v17, F1: 612, F2: 2}
                                            _ = v21
                                        }
                                        _ = v17
                                    } else {
                                        fmt.Println(fi7, v2)
                                    }
                                }
                                v3 -= 1
                            }
                            var v22 uint16 = 32767
                            var v23 int8 = (^(v1))
                            _ = v22
                            _ = v23
                        }
                    }
                    switch "hello" {
                    case "foo bar":
                        v3 -= (^(v3))
                    case "a":
                        v3 = v3
                    case "":
                        v24 := GenS1{F0: 605, F1: 255, F2: 995}
                        _ = v24
                    default:
                        if true {
                            v25 := GenS5{F0: 1, F1: -9223372036854775808, F2: -1}
                            var v26 int32 = -1
                            v27 := GenS1{F0: 128, F1: 64, F2: 1}
                            v26 = v26
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v1 -= v1
                for fi28 := 0; fi28 < 1; fi28++ {
                    _ = len(v4)
                    var v29 uint8 = 108
                    _ = fi28
                    _ = v29
                }
            }
        }
        _ = fi5
        _ = v6
    }
    var v30 uint16 = 32767
    var v31 int32 = 42
    v3 = (^(v3))
    v2 = 1e10
    switch 42 {
    case 127:
        v32 := GenS4{F0: 0, F1: (v2 > ((v2 * (v2 - (v2 + (3.14159 + ((((v2 / -2.25) * (0.1 + 0.0)) * v2) + -2.25))))) / v2)), F2: -750}
        _ = v32
    default:
        v31 = v31
    }
    v2 = ((((v2 / 1.0) - v2) * ((1.5 - v2) / 3.14159)) + 0.1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v30 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
