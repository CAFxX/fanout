package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 1
    var v4 uint16 = 32767
    var v5 uint64 = 882
    var v6 uint16 = 444
    if (true || false) {
        v5 = v5
        v6 += 2
        var v7 string = ""
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 273, F1: "a", F2: 1}
        _ = fi8
        _ = v9
    }
    if (100 != -1) {
        v10 := GenS2{F0: 768, F1: -2147483648}
        if ((v3 * v3) > int64(v5)) {
            fmt.Println(v5, v3, v6)
            if ((2147483647 < -52) && (!false)) {
                delete(v2, 224)
                fmt.Println(v3, v4)
                var v11 int16 = 1000
                _ = v11
            } else {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 int64 = -9223372036854775808
                    _ = fi12
                    _ = v13
                }
                v3 = v3
            }
        } else {
            if ((!true) || (!true)) {
                if true {
                    var v14 float32 = (((0.0 + 3.14159) + ((float32(-2.25) * -2.25) + float32((-2.25 / (1.5 / (0.0 - 3.14159)))))) / (1.5 / (1.5 / (((0.0 / 1.5) - 1.5) + 0.0))))
                    _ = v14
                }
            } else {
                if (-1 > 19) {
                    _ = len(v2)
                    var v15 uint8 = 58
                    _ = v15
                }
                v3 += (v3 % ((v3) | 1))
            }
            v16 := GenS3{F0: 0, F1: (v3 - v3), F2: 976}
            v4 = v4
            _ = v16
        }
        _ = v10
    }
    var v17 uint64 = 636
    fmt.Println(v4)
    switch "hello" {
    case "a":
        switch "" {
        case "\n\t":
            switch "\n\t" {
            case "hello":
                v3 = 42
            case "foo bar":
                switch ("" + "hello") {
                case "a":
                    if false {
                        v18 := GenS3{F0: 127, F1: (v3 % (((v3 * v3)) | 1)), F2: 1}
                        v19 := []int{-80, 185}
                        _ = v18
                        _ = v19
                    } else {
                        switch 873 {
                        case -126:
                            for fi20 := 0; fi20 < 17; fi20++ {
                                v21 := GenS3{F0: -116, F1: (v3 & v3), F2: 383}
                                _ = fi20
                                _ = v21
                            }
                        case 42:
                            fmt.Println(v4, v5, v4)
                        case 0:
                            v22 := []int{322, 543}
                            _ = v22
                        }
                    }
                default:
                    if false {
                        v5 -= 1
                    } else {
                        v23 := []int{}
                        _ = v23
                    }
                }
            case "a":
                if false {
                    v4 = (v4 / ((v6) | 1))
                } else {
                    v24 := GenS4{F0: v3, F1: 1, F2: 54}
                    v4 = v4
                    _ = v24
                }
            }
        case "foo bar":
            delete(v2, 65535)
        }
    case "":
        var v25 int8 = 70
        _ = v25
    case "hello":
        var v26 bool = false
        _ = v26
    default:
        fmt.Println(v3, v6, v17)
    }
    v5 = v5
    if true {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS4{F0: (v3 << ((0) & 63)), F1: 0, F2: 13}
            switch ("\n\t" + ("hello" + "\n\t")) {
            case "":
                var v29 string = "\n\t"
                _ = v29
            case "\n\t":
                switch (fi27 | (fi27 << ((4294967295) & 63))) {
                case -107:
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v31 := []int{fi30, fi27}
                        _ = fi30
                        _ = v31
                    }
                case -112:
                    for fi32 := 0; fi32 < 20; fi32++ {
                        var v33 int16 = 32767
                        v34 := GenS5{F0: 4294967295, F1: (^(v5)), F2: v17}
                        var v35 float64 = (0.1 - float64(((-2.25 * 0.1) / 1e10)))
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
            case "foo bar":
                var v36 uint16 = v4
                _ = v36
            }
            _ = fi27
            _ = v28
        }
    } else {
        if (!(32767 == 0)) {
            v37 := GenS2{F0: 885, F1: -514}
            _ = v37
        }
        var v38 uint = 4294967295
        v6 = (v6 >> ((v38) & 63))
        _ = v38
    }
    if false {
        v39 := GenS6{F0: 255}
        fmt.Println(v6, v6)
        v40 := GenS4{F0: v3, F1: 128, F2: 55}
        _ = v39
        _ = v40
    }
    fmt.Println(v3, v5, v6)
    if false {
        v41 := GenS6{F0: 2}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 10; fi42++ {
            var v43 float32 = float32(1e10)
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
