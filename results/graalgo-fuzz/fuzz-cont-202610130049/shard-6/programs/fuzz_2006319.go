package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint32, p1 uint8) bool {
    var v5 float64 = 3.14159
    v5 += (float64((-2.25 / 3.14159)) / float64(-2.25))
    if (2147483647 >= 1) {
        var v6 int8 = -128
        _ = v6
    }
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return ((false && false) || true)
}

func main() {
    var v1 int8 = 0
    var v2 int = 0
    var v3 string = "hello"
    var v4 int = 9223372036854775807
    var v7 string = v3
    var v8 float64 = float64((0.1 / (-2.25 * (1e10 + 3.14159))))
    v9 := GenS1{F0: 398, F1: v7, F2: (^(v4))}
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!(1 == 220)) {
            if (0 >= 997) {
                if false {
                    v8 = (v8 * v8)
                }
            }
        } else {
            v11 := GenS2{F0: -854}
            var v12 bool = ("" > v3)
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    v4 = v2
    var v13 int16 = -32768
    switch v4 {
    case 32767:
        v1 = v1
    case 127:
        if true {
            fmt.Println(v3, v7)
            switch (v4 ^ -1) {
            case 9223372036854775807:
                v4 = -795
            }
        }
    case -1:
        if true {
            if (1 != 2) {
                if ((42 > 4294967295) || (255 != 255)) {
                    if true {
                        v4 = v2
                    }
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if (((v7 + v7) + "hello") != "\n\t") {
                            for fi15 := 0; fi15 < 18; fi15++ {
                                switch v4 {
                                case -971:
                                    v16 := GenS1{F0: 365, F1: v3, F2: 32767}
                                    var v17 int32 = -2147483648
                                    fmt.Println(v7)
                                    _ = v16
                                    _ = v17
                                case 1:
                                    fmt.Println(v7, v7, fi15)
                                case 14:
                                    var v18 uint8 = 188
                                    _ = v18
                                }
                                _ = fi15
                            }
                        }
                        if false {
                            v19 := []int{fi14, v2, (fi14 / ((fi14) | 1)), fi14}
                            _ = v19
                        }
                        _ = fi14
                    }
                    var v20 string = "foo bar"
                    _ = v20
                }
                fmt.Println(v13, v13, v3)
            } else {
                v4 = (v2 & -539)
            }
            if false {
                var v21 int32 = 1
                _ = v21
            }
        }
    default:
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint = 691
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v2, v1, v7)
    for fi24 := 0; fi24 < 19; fi24++ {
        switch (^(v4)) {
        case -1:
            var v25 string = "\n\t"
            _ = v25
        default:
            v13 -= (v13 / ((102) | 1))
        }
        v26 := []int{}
        _ = fi24
        _ = v26
    }
    v27 := GenS2{F0: (v2 * v4)}
    v28 := GenS2{F0: v2}
    for fi29 := 0; fi29 < 9; fi29++ {
        fmt.Println(v3, v3)
        for fi30 := 0; fi30 < 14; fi30++ {
            v4 = (v4 + v2)
            v1 = 100
            _ = fi30
        }
        _ = fi29
    }
    var v31 uint8 = 3
    switch v7 {
    case "":
        v32 := GenS3{F0: 0, F1: -56}
        _ = v32
    case "hello":
        if (false && ("a" > "")) {
            for fi33 := 0; fi33 < 4; fi33++ {
                for fi34 := 0; fi34 < 20; fi34++ {
                    v35 := []int{-633, (fi34 / (((fi33 % (((v2 ^ v4)) | 1))) | 1)), int(v13), (fi33 * fi33)}
                    v35 = append(v35, int(v13))
                    _ = fi34
                    _ = v35
                }
                switch v7 {
                case "hello":
                    if (!true) {
                        var v36 string = v7
                        fmt.Println(v36, v3)
                        _ = v36
                    } else {
                        v37 := GenS4{F0: (v2 / ((-427) | 1)), F1: v13, F2: 1}
                        _ = v37
                    }
                case "a":
                    var v38 uint64 = 0
                    _ = v38
                }
                _ = fi33
            }
            fmt.Println(v13)
            for fi39 := 0; fi39 < 10; fi39++ {
                for fi40 := 0; fi40 < 2; fi40++ {
                    v41 := []int{428, (fi40 % ((fi40) | 1))}
                    v42 := GenS5{F0: 930}
                    var v43 int64 = -1
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi39
            }
        }
    }
    v44 := GenS3{F0: 1, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v7
    _ = v8
    _ = v9
}
