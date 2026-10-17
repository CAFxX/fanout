package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint32 = 365
    var v2 int64 = -907
    var v3 bool = (-128 > 36)
    v2 -= v2
    v4 := []int{32768}
    _ = len(v4)
    if ((v2 + (v2 ^ (v2 + v2))) < v2) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2 = v2
    fmt.Println(v1)
    v5 := GenS1{F0: 42, F1: 696, F2: 200}
    v3 = v3
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            _ = len(v4)
            _ = fi7
        }
        _ = fi6
    }
    if v3 {
        v4 = append(v4, -978)
        for fi8 := 0; fi8 < 5; fi8++ {
            if (323 <= -848) {
                var v9 int = fi8
                v2 = 878
                _ = v9
            }
            if v3 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v10 bool = v3
                v3 = (3.14159 < ((0.0 + 1e10) + (1.5 + 0.1)))
                _ = v10
            }
            _ = fi8
        }
        fmt.Println(v2, v1)
    } else {
        if len(v4) > 0 {
            v4[0] = -508
        }
        switch 42 {
        case -406:
            switch 487 {
            case -536:
                switch -695 {
                case 127:
                    v2 = (^(v2))
                case -2147483648:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v1 += v1
                        if (!(591 > 18446744073709551615)) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v12 int = fi11
                            for fi13 := 0; fi13 < 14; fi13++ {
                                switch 4294967295 {
                                case 32768:
                                    v2 += v2
                                case -2147483648:
                                    var v14 uint16 = 65535
                                    _ = v14
                                default:
                                    if (!((false && (127 >= 97)) || false)) {
                                        var v15 uint8 = 211
                                        switch fi13 {
                                        case 9223372036854775807:
                                            v1 += v1
                                        case 32767:
                                            if ((v3 || (true || (18446744073709551615 < 0))) && (42 > -128)) {
                                                var v16 uint64 = 609
                                                _ = v16
                                            } else {
                                                v1 -= v1
                                                var v17 float32 = (0.0 * float32((0.0 * (1.5 + 1.5))))
                                                v1 -= v1
                                                _ = len(v4)
                                                v12 = v12
                                                _ = v17
                                            }
                                        case 0:
                                            if true {
                                                v3 = true
                                                if ("" <= ("hello" + "foo bar")) {
                                                    v18 := GenS1{F0: 1, F1: 4294967295, F2: 96}
                                                    _ = v18
                                                }
                                                _ = len(v4)
                                            } else {
                                                v12 = fi13
                                                var v19 float32 = 0.1
                                                v20 := GenS1{F0: 1, F1: 42, F2: v15}
                                                _ = v19
                                                _ = v20
                                            }
                                        default:
                                            var v21 uint64 = 2
                                            _ = v21
                                        }
                                        _ = v15
                                    } else {
                                        fmt.Println(fi13)
                                        v22 := GenS2{F0: ("hello" + "\n\t")}
                                        _ = v22
                                    }
                                }
                                _ = fi13
                            }
                            _ = v12
                        }
                        var v23 int32 = 0
                        _ = fi11
                        _ = v23
                    }
                }
            }
        case 401:
            v1 -= 711
        default:
            switch "foo bar" {
            case "\n\t":
                v1 = v1
            case "foo bar":
                var v24 uint64 = 456
                _ = v24
            }
        }
    }
    v1 -= 65535
    var v25 int16 = int16(v2)
    var v26 uint16 = 825
    var v27 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
