package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
type GenS5 struct { F0 int16; F1 int; F2 uint32 }
type GenS6 struct { F0 int; F1 uint8; F2 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint16) int64 {
    p0 += p0
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int = fi5
        p0 = (p0 & p0)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 829, F1: 65535, F2: (p0 << ((733) & 63))}
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = -2.25
    var v3 string = ("a" + "")
    v4 := make(map[uint32]uint32)
    switch 700 {
    case 595:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS2{F0: (240 != 255), F1: 951}
            _ = fi8
            _ = v9
        }
    case -1:
        if false {
            var v10 bool = (true && (100 <= 127))
            v11 := v4[65535]
            _ = v10
            _ = v11
        } else {
            var v12 string = (v3 + (v3 + ""))
            _ = v12
        }
    }
    v13 := []int{749, 9223372036854775807, 1}
    delete(v4, 245)
    v13 = append(v13, 32768)
    _ = len(v1)
    if ((!false) || (false || (false && (!false)))) {
        v2 = float64((1e10 + 1.5))
        v3 = v3
    } else {
        v2 += (v2 - v2)
        _ = len(v13)
    }
    v14 := GenS3{F0: 393, F1: -1, F2: -1}
    v3 = "a"
    v3 = (("hello" + "foo bar") + ("a" + ""))
    switch "foo bar" {
    case "foo bar":
        for fi15 := 0; fi15 < 9; fi15++ {
            if true {
                switch -9223372036854775808 {
                case 646:
                    fmt.Println(v3, fi15)
                case 824:
                    var v16 uint8 = 0
                    _ = v16
                default:
                    switch fi15 {
                    case -2147483648:
                        if (true || true) {
                            var v17 int8 = -20
                            var v18 int32 = -1
                            _ = v17
                            _ = v18
                        } else {
                            var v19 float32 = 0.1
                            _ = v19
                        }
                    case 292:
                        v13 = append(v13, (fi15 | fi15))
                    case 42:
                        _ = len(v13)
                    }
                }
                switch (fi15 >> ((42) & 63)) {
                case -9223372036854775808:
                    if (18446744073709551615 == 18446744073709551615) {
                        v20 := GenS4{F0: 376, F1: 65535, F2: 1}
                        _ = len(v4)
                        var v21 int = (fi15 << ((1) & 63))
                        _ = v20
                        _ = v21
                    }
                case 9223372036854775807:
                    if (1e10 <= float32(v2)) {
                        if (("foo bar" + v3) >= (("hello" + (v3 + (v3 + v3))) + "")) {
                            var v22 uint16 = 32767
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v22 = 32767
                                _ = fi23
                            }
                            switch fi15 {
                            case 4294967295:
                                v22 += (v22 - (^(v22)))
                            case -895:
                                var v24 string = ((v3 + v3) + ("\n\t" + (v3 + v3)))
                                _ = v24
                            }
                            _ = v22
                        }
                    }
                default:
                    v25 := GenS5{F0: 1, F1: fi15, F2: 2147483648}
                    _ = v25
                }
            }
            switch (v3 + v3) {
            case "\n\t":
                v26 := GenS6{F0: (-(fi15)), F1: 116, F2: 372}
                _ = v26
            case "":
                v27 := v4[65535]
                _ = v27
            case "foo bar":
                v28 := GenS7{F0: 127}
                _ = v28
            default:
                switch v3 {
                case "\n\t":
                    delete(v1, -130)
                default:
                    var v29 int = fi15
                    _ = v29
                }
            }
            _ = fi15
        }
    case "a":
        switch v3 {
        case "":
            switch 42 {
            case -578:
                if false {
                    v2 += (v2 + ((v2 + float64(-2.25)) - v2))
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v3 = ((v3 + ((v3 + "") + v3)) + "")
                        if len(v13) > 0 {
                            v13[0] = fi30
                        }
                        v3 = ("hello" + "foo bar")
                        _ = fi30
                    }
                }
            default:
                var v31 int64 = genFn0(565)
                _ = v31
            }
        case "foo bar":
            var v32 int16 = -671
            _ = v32
        }
    case "":
        switch v3 {
        case "a":
            if (!false) {
                var v33 uint8 = 138
                v3 = "foo bar"
                _ = v33
            }
        case "hello":
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := GenS8{F0: 0, F1: 129}
                v3 = v3
                _ = fi34
                _ = v35
            }
        case "":
            v1[-551] = 42
        }
    }
    var v36 uint16 = 51
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
