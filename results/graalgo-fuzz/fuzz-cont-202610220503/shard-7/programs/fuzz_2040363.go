package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 1
    var v3 uint = 492
    var v4 int = 610
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    v2 -= (v2 + 406)
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint = 255
                var v10 int8 = -40
                _ = fi8
                _ = v10
                _ = v9
            }
            delete(v1, 177)
            v11 := GenS1{F0: false, F1: 65535, F2: 730}
            _ = fi7
            _ = v11
        }
        var v12 uint8 = 128
        var v13 float64 = 1.5
        _ = v12
        _ = v13
    }
    switch 128 {
    case -863:
        v14 := GenS2{F0: -1}
        _ = v14
    case 1:
        switch v4 {
        case -323:
            var v15 float64 = float64((3.14159 / 1.0))
            _ = v15
        default:
            if true {
                v4 += -467
                var v16 int32 = int32(v3)
                _ = v16
            }
        }
    case -9223372036854775808:
        switch v4 {
        case 2147483647:
            switch "" {
            case "hello":
                v17 := GenS3{F0: (^(v2)), F1: 1}
                _ = v17
            case "foo bar":
                v18 := GenS4{F0: "a", F1: (v2 / ((65535) | 1)), F2: 495}
                _ = v18
            default:
                var v19 int64 = 882
                _ = v19
            }
        case 272:
            v20 := v1[680]
            _ = v20
        case 32767:
            switch (-(v4)) {
            case -2147483648:
                if (32767 <= -1) {
                    _ = len(v6)
                    fmt.Println(v2)
                }
            }
        }
    default:
        v21 := v6[(v4 << ((651) & 63))]
        _ = v21
    }
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 uint32 = 2147483647
        if false {
            v25 := GenS5{F0: "a"}
            fmt.Println(v4)
            _ = v25
        } else {
            var v26 int8 = 42
            switch "a" {
            case "":
                var v27 int32 = (v22 / ((-2147483648) | 1))
                _ = v27
            case "\n\t":
                var v28 uint64 = 18446744073709551615
                _ = v28
            case "hello":
                for fi29 := 0; fi29 < 6; fi29++ {
                    switch (v4 - fi23) {
                    case 32767:
                        switch "a" {
                        case "hello":
                            var v30 float32 = -2.25
                            _ = v30
                        case "foo bar":
                            if true {
                                v2 = 539
                                v31 := []int{804, int(v22)}
                                if (false && false) {
                                    v32 := GenS1{F0: (4294967295 > 4294967295), F1: 255, F2: 18446744073709551615}
                                    v31 = append(v31, (-(fi23)))
                                    _ = v32
                                } else {
                                    var v33 uint16 = 0
                                    delete(v6, (v4 / ((127) | 1)))
                                    _ = v33
                                }
                                _ = v31
                            }
                        case "":
                            var v34 uint16 = 0
                            _ = v34
                        }
                    case 2147483647:
                        var v35 uint32 = v2
                        _ = v35
                    }
                    fmt.Println(v4)
                    var v36 string = "a"
                    _ = fi29
                    _ = v36
                }
            default:
                v37 := GenS4{F0: ((("" + (((("hello" + "hello") + "\n\t") + (("foo bar" + (("foo bar" + ("a" + (("a" + ("hello" + "foo bar")) + "hello"))) + "")) + "\n\t")) + "hello")) + "a") + (("foo bar" + "\n\t") + "\n\t")), F1: 0, F2: 4294967295}
                _ = v37
            }
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 uint8 = 0
        _ = fi38
        _ = v39
    }
    v22 = v22
    v40 := GenS5{F0: (("hello" + "\n\t") + "foo bar")}
    if true {
        fmt.Println(v4, v2)
        v41 := GenS6{F0: 213}
        for fi42 := 0; fi42 < 18; fi42++ {
            var v43 uint64 = 89
            v44 := GenS7{F0: 532, F1: 1, F2: 0}
            if (0 >= 0) {
                v4 -= v4
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    }
    if (v2 <= (v2 * (v2 & (v2 << (((v3 ^ v3)) & 63))))) {
        if (!true) {
            var v45 int64 = 9223372036854775807
            _ = v45
        } else {
            switch (v4 & v4) {
            case 1:
                v46 := GenS8{F0: 255, F1: v22}
                _ = v46
            case 4294967295:
                switch (v4 & -861) {
                case 0:
                    v47 := v1[2147483647]
                    _ = v47
                case 128:
                    switch (v4 << ((v3) & 63)) {
                    case 402:
                        var v48 bool = (!(v2 != (v2 >> ((v3) & 63))))
                        _ = v48
                    case -1:
                        var v49 int = (v4 % ((v4) | 1))
                        _ = v49
                    case 0:
                        if false {
                            var v50 int8 = -45
                            var v51 int64 = 0
                            _ = v50
                            _ = v51
                        }
                    default:
                        delete(v5, (v2 ^ (v2 << (((v3 * 253)) & 63))))
                    }
                }
            }
        }
        v3 = (^(v3))
    }
    v52 := GenS9{F0: 255, F1: (v2 >> ((v3) & 63)), F2: "a"}
    fmt.Println(v2, v22, v22)
    v2 += (v2 % (((v2 - (v2 + (v2 + (v2 % ((2147483647) | 1)))))) | 1))
    if ((v2 & (v2 & (v2 * 1))) <= 0) {
        var v53 string = "a"
        _ = v53
    }
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v6
}
