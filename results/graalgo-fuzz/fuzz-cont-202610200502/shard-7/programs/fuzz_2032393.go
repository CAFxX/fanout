package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint; F1 string; F2 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 uint8 = 255
    var v4 uint = 84
    var v5 int8 = -128
    v5 = v5
    v3 += (v3 / (((^(v3))) | 1))
    var v6 uint64 = 42
    if ("hello" > ("hello" + "a")) {
        v3 = (v3 >> ((v4) & 63))
        switch "hello" {
        case "a":
            if ((true || (false && true)) || true) {
                var v7 uint32 = 2147483647
                if (544 >= -503) {
                    var v8 uint32 = 65535
                    _ = v8
                }
                _ = v7
            }
        case "foo bar":
            if true {
                v6 = 257
                var v9 string = ""
                _ = v9
            } else {
                fmt.Println(v5, v5)
                var v10 uint64 = (v6 / ((18446744073709551615) | 1))
                _ = v10
            }
        }
    }
    var v11 uint = 751
    var v12 int8 = (^(v5))
    v5 += v12
    if false {
        var v13 int8 = (v5 - v5)
        v11 = (v11 & v4)
        if ((true || ((1000 >= 0) && false)) || true) {
            if (1 != v11) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v4 += (v4 - v11)
                    switch (fi14 % (((fi14 | (^(fi14)))) | 1)) {
                    case 32767:
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 string = "hello"
                            var v17 float64 = ((0.1 * 1e10) + 1.5)
                            v13 -= (v12 << (((v4 ^ (v11 * (v4 / ((4294967295) | 1))))) & 63))
                            v18 := GenS1{F0: 42, F1: ((1 >= uint16(v4)) || (186 <= 0)), F2: -32768}
                            v19 := GenS1{F0: 446, F1: ((!(0 <= 0)) && false), F2: -181}
                            if (42 == 0) {
                                fmt.Println(v6)
                            }
                            _ = len(v2)
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    default:
                        var v20 int = 418
                        _ = v20
                    }
                    _ = fi14
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := v1[(fi21 - (fi21 >> ((v11) & 63)))]
                    if false {
                        fmt.Println(v5, v11)
                        var v23 string = "hello"
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            }
            switch "\n\t" {
            case "\n\t":
                var v24 float32 = 1.5
                _ = v24
            case "hello":
                v25 := GenS2{F0: -9223372036854775808, F1: -2147483648}
                _ = v25
            }
        }
        _ = v13
    } else {
        for fi26 := 0; fi26 < 18; fi26++ {
            fmt.Println(v4, v12)
            _ = fi26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int8 = (v12 | v12)
            v29 := GenS2{F0: 185, F1: 2147483647}
            _ = fi27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 13; fi31++ {
                fmt.Println(v4)
                var v32 string = ("a" + "\n\t")
                v33 := GenS3{F0: (^(v11)), F1: 1}
                _ = fi31
                _ = v32
                _ = v33
            }
            if (!(1 >= 791)) {
                if true {
                    v34 := GenS3{F0: v11, F1: 42}
                    _ = v34
                } else {
                    v35 := GenS4{F0: 1, F1: 18446744073709551615}
                    v36 := GenS5{F0: 793, F1: 0}
                    _ = v35
                    _ = v36
                }
            } else {
                fmt.Println(v11, v4)
                if false {
                    switch fi30 {
                    case -9223372036854775808:
                        var v37 int = (fi30 - (fi30 / (((fi30 * (fi30 >> ((2) & 63)))) | 1)))
                        _ = v37
                    case 1:
                        var v38 float64 = float64(0.1)
                        _ = v38
                    default:
                        v1[fi30] = 65535
                    }
                } else {
                    v39 := []int{fi30}
                    fmt.Println(v6, v4)
                    if true {
                        var v40 float64 = 0.1
                        for fi41 := 0; fi41 < 20; fi41++ {
                            v42 := v2[fi30]
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        v43 := v2[fi30]
                        delete(v2, 9223372036854775807)
                        _ = v43
                    }
                    _ = v39
                }
            }
            _ = fi30
        }
    }
    var v44 bool = (!(2147483648 > 2147483648))
    v12 = (v5 / ((-120) | 1))
    v45 := GenS3{F0: 852, F1: 1000}
    v46 := GenS6{F0: 385, F1: "", F2: 42}
    v11 = v4
    for fi47 := 0; fi47 < 19; fi47++ {
        var v48 uint = v4
        switch int(v3) {
        case 128:
            var v49 uint = 0
            _ = v49
        case 1:
            delete(v2, fi47)
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
