package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }

func main() {
    var v1 int = 72
    v2 := []int{0, v1, (v1 & 2147483647)}
    v3 := make(map[uint32]int)
    var v4 int16 = 32767
    fmt.Println(v4, v4)
    v1 = v1
    v1 -= -1
    v4 += v4
    fmt.Println(v1)
    v2 = append(v2, 128)
    for fi5 := 0; fi5 < 6; fi5++ {
        if (255 <= 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v6 float64 = (0.0 * 1e10)
            _ = v6
        }
        fmt.Println(v1, v4)
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v4, v1, v1)
            switch ("a" + "") {
            case "a":
                fmt.Println(fi7)
            case "\n\t":
                v8 := GenS1{F0: 425}
                _ = v8
            default:
                v1 -= fi7
            }
            v9 := GenS2{F0: (209 <= 146), F1: (0.0 > 1.5), F2: 89}
            _ = fi7
            _ = v9
        }
        _ = fi5
    }
    if (!(0 == 4294967295)) {
        v2 = append(v2, (v1 % ((-9223372036854775808) | 1)))
    }
    v4 = (v4 << ((672) & 63))
    v4 -= v4
    var v10 int = v1
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 = v4
        v10 -= (v10 * 32768)
        _ = fi11
    }
    var v12 uint64 = 2
    v4 -= (v4 | v4)
    v13 := GenS2{F0: (v4 <= v4), F1: true, F2: 2}
    switch ("" + "foo bar") {
    case "":
        var v14 float32 = float32(1e10)
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 19; fi15++ {
            if (-596 == -1) {
                v16 := v3[65535]
                if (1e10 > 3.14159) {
                    v3[1] = -1
                    if (3.14159 >= (0.0 - (0.0 / (0.1 - float32((-2.25 / 1e10)))))) {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            if false {
                                switch fi17 {
                                case 128:
                                    if (18446744073709551615 >= 1) {
                                        v18 := GenS2{F0: false, F1: true, F2: 32}
                                        _ = v18
                                    }
                                }
                                if (!((v12 ^ v12) >= v12)) {
                                    fmt.Println(v10, v12)
                                }
                                if false {
                                    var v19 uint32 = 2
                                    var v20 uint16 = 65535
                                    var v21 uint16 = (v20 | (^(v20)))
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            _ = fi17
                        }
                        var v22 uint = 255
                        _ = v22
                    } else {
                        switch (v16 & 668) {
                        case 264:
                            switch "foo bar" {
                            case "":
                                v23 := GenS3{F0: (v4 * (v4 >> ((42) & 63)))}
                                _ = v23
                            case "hello":
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    if false {
                                        fmt.Println(v4, v1, fi15)
                                    }
                                    var v25 int = fi24
                                    _ = fi24
                                    _ = v25
                                }
                            case "\n\t":
                                switch v1 {
                                case -2147483648:
                                    v26 := GenS4{F0: 32767, F1: 2, F2: 42}
                                    _ = v26
                                }
                            default:
                                switch "a" {
                                case "\n\t":
                                    v2 = append(v2, -9223372036854775808)
                                case "foo bar":
                                    switch (v16 / ((9223372036854775807) | 1)) {
                                    case 117:
                                        fmt.Println(fi15, v1)
                                    case -835:
                                        switch "hello" {
                                        case "\n\t":
                                            v27 := GenS4{F0: 2, F1: (^(v12)), F2: 100}
                                            _ = v27
                                        case "a":
                                            delete(v3, 193)
                                        }
                                    default:
                                        v1 = int(v12)
                                    }
                                default:
                                    v28 := GenS1{F0: 255}
                                    _ = v28
                                }
                            }
                        case 42:
                            if ((!true) || true) {
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v30 := GenS2{F0: (!((1e10 >= (1.5 - 0.1)) && false)), F1: (0.1 != (float64(-2.25) * (float64(-2.25) / (-2.25 / float64(1e10))))), F2: 223}
                                    v1 = fi29
                                    fmt.Println(v10, fi15)
                                    _ = fi29
                                    _ = v30
                                }
                            } else {
                                switch "\n\t" {
                                case "":
                                    v31 := GenS1{F0: 653}
                                    _ = v31
                                case "\n\t":
                                    var v32 int8 = 127
                                    _ = v32
                                default:
                                    v33 := GenS2{F0: true, F1: (1e10 <= (3.14159 - 0.0)), F2: 42}
                                    _ = v33
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    for fi34 := 0; fi34 < 1; fi34++ {
                                        if (56 == 71) {
                                            fmt.Println(v16, fi34)
                                            v10 = (fi34 ^ (v10 + (-(v10))))
                                            var v35 uint16 = 65535
                                            _ = v35
                                        } else {
                                            var v36 int64 = 672
                                            var v37 int16 = 0
                                            var v38 int64 = 1
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                        fmt.Println(fi34, v12, fi15)
                                        _ = fi34
                                    }
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                        case 9223372036854775807:
                            var v39 int = v16
                            _ = v39
                        }
                        v1 -= (v10 % ((1) | 1))
                    }
                }
                v40 := GenS4{F0: 42, F1: 42, F2: -128}
                _ = v16
                _ = v40
            }
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 uint64 = v12
                _ = fi41
                _ = v42
            }
            v12 = v12
            _ = fi15
        }
    case "foo bar":
        _ = len(v2)
    }
    v12 += (v12 >> ((42) & 63))
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
