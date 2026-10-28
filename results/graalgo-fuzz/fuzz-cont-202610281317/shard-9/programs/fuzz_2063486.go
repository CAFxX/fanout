package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint; F2 int16 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := []int{-2147483648, 4294967295, -9223372036854775808}
    var v4 uint64 = 719
    v5 := make(map[uint32]int)
    var v6 int8 = 69
    switch 650 {
    case 2147483647:
        v7 := GenS1{F0: 1, F1: false, F2: (2147483647 > -283)}
        _ = v7
    case -462:
        for fi8 := 0; fi8 < 19; fi8++ {
            v6 = 42
            _ = fi8
        }
    case 42:
        fmt.Println(v6, v4)
    default:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 bool = false
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v3) > 0 {
            v3[0] = -322
        }
        _ = fi11
    }
    v12 := GenS1{F0: 53, F1: false, F2: false}
    if (!false) {
        switch "foo bar" {
        case "":
            if true {
                var v13 bool = (1000 > 0)
                if (!true) {
                    var v14 int8 = v6
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        _ = len(v3)
                        var v16 int8 = (v6 % ((v6) | 1))
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if v13 {
                                v6 -= (^(v16))
                                var v18 uint8 = 106
                                _ = v18
                            } else {
                                fmt.Println(fi17)
                                if false {
                                    v19 := GenS1{F0: 255, F1: (0.0 < 1.5), F2: v13}
                                    v4 -= (v4 / ((980) | 1))
                                    _ = v19
                                } else {
                                    _ = len(v3)
                                }
                                v20 := GenS1{F0: 203, F1: (32767 < 151), F2: false}
                                _ = v20
                            }
                            v13 = true
                            _ = fi17
                        }
                        _ = fi15
                        _ = v16
                    }
                    v6 = (v6 - v6)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            if false {
                                v6 = v6
                            }
                            _ = fi22
                        }
                        v23 := GenS2{F0: (v13 && ((1 < 42) && false))}
                        var v24 uint16 = 255
                        _ = fi21
                        _ = v23
                        _ = v24
                    }
                }
                _ = v13
            }
        case "a":
            if (0 < 65535) {
                var v25 uint16 = 456
                _ = v25
            } else {
                var v26 uint = 260
                _ = v26
            }
        case "hello":
            var v27 float64 = 0.1
            _ = v27
        default:
            if (668 == 4294967295) {
                if ((1000 == 32767) && (!false)) {
                    var v28 float32 = 0.0
                    fmt.Println(v28, v28)
                    var v29 int = 32767
                    _ = v28
                    _ = v29
                } else {
                    delete(v5, 2147483647)
                    v30 := GenS3{F0: 0}
                    _ = v30
                }
                var v31 int = -2147483648
                if false {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v33 := GenS3{F0: 408}
                        if (v6 <= 127) {
                            v4 += 2
                            v34 := GenS3{F0: 18446744073709551615}
                            _ = v34
                        }
                        if true {
                            v4 -= v4
                            v35 := GenS4{F0: 255, F1: 255, F2: 42}
                            v6 += v6
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                    if (1000 <= 242) {
                        fmt.Println(v31)
                        if true {
                            v36 := GenS5{F0: -9223372036854775808, F1: v31}
                            _ = v36
                        }
                    } else {
                        var v37 int16 = 32767
                        v38 := GenS3{F0: 4294967295}
                        v37 = -32768
                        _ = v37
                        _ = v38
                    }
                }
                _ = v31
            }
        }
        var v39 uint16 = 2
        var v40 int32 = -957
        _ = v39
        _ = v40
    }
    v41 := GenS6{F0: 426}
    switch -1 {
    case 32768:
        v3 = append(v3, -358)
    }
    var v42 int16 = 1000
    v42 += 0
    for fi43 := 0; fi43 < 10; fi43++ {
        if len(v3) > 0 {
            v3[0] = 578
        }
        v44 := GenS4{F0: 42, F1: 0, F2: v42}
        _ = fi43
        _ = v44
    }
    var v45 float32 = -2.25
    v3 = append(v3, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
