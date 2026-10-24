package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }

func main() {
    var v1 float64 = (0.0 - (-2.25 / (0.0 - (3.14159 - -2.25))))
    var v2 string = "\n\t"
    var v3 uint64 = 42
    v4 := make(map[int]string)
    var v5 uint64 = 1
    if (true || true) {
        switch 32768 {
        case -1:
            v3 += (v3 >> ((145) & 63))
        case 9223372036854775807:
            v6 := GenS1{F0: 716}
            _ = v6
        case 1:
            var v7 string = ""
            _ = v7
        default:
            if true {
                v8 := GenS2{F0: 42}
                if (-1 < -586) {
                    var v9 int64 = 41
                    if (!(float64(3.14159) == v1)) {
                        v5 = (v3 / ((v3) | 1))
                    }
                    fmt.Println(v3, v1)
                    _ = v9
                } else {
                    v10 := []int{}
                    fmt.Println(v3, v1, v3)
                    switch 47 {
                    case 128:
                        for fi11 := 0; fi11 < 7; fi11++ {
                            var v12 uint16 = 571
                            switch (((v2 + v2) + v2) + v2) {
                            case "":
                                if len(v10) > 0 {
                                    v10[0] = -1
                                }
                            case "foo bar":
                                v1 = v1
                            case "\n\t":
                                switch fi11 {
                                case -134:
                                    switch fi11 {
                                    case 32767:
                                        v1 += v1
                                    case 32768:
                                        v5 = (v5 % ((18446744073709551615) | 1))
                                    default:
                                        fmt.Println(v1, v5, v5)
                                    }
                                case -9223372036854775808:
                                    for fi13 := 0; fi13 < 12; fi13++ {
                                        v1 = 0.0
                                        _ = fi13
                                    }
                                default:
                                    v5 = 2
                                }
                            default:
                                if true {
                                    if true {
                                        for fi14 := 0; fi14 < 15; fi14++ {
                                            fmt.Println(v2, v5)
                                            var v15 uint16 = (v12 % ((65535) | 1))
                                            v15 += (v12 & v15)
                                            _ = fi14
                                            _ = v15
                                        }
                                        if ((fi11 * fi11) == fi11) {
                                            var v16 uint64 = (v5 | (v3 >> ((2) & 63)))
                                            var v17 string = v2
                                            v18 := GenS2{F0: 42}
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        } else {
                                            if true {
                                                v19 := GenS3{F0: true, F1: 0, F2: 1}
                                                var v20 bool = true
                                                _ = v19
                                                _ = v20
                                            }
                                            var v21 uint16 = v12
                                            v5 = (v5 - (v5 % ((v5) | 1)))
                                            _ = v21
                                        }
                                    } else {
                                        fmt.Println(v12, v3)
                                    }
                                }
                            }
                            _ = fi11
                            _ = v12
                        }
                    case 32768:
                        var v22 int8 = 1
                        _ = v22
                    default:
                        var v23 uint32 = 742
                        _ = v23
                    }
                    _ = v10
                }
                v24 := GenS2{F0: 1}
                _ = v24
                _ = v8
            }
        }
        v25 := []int{1, 2147483647, -2147483648, 42}
        if false {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 uint64 = v5
                _ = fi26
                _ = v27
            }
            v28 := GenS1{F0: 86}
            _ = v28
        }
        _ = v25
    }
    var v29 float64 = -2.25
    v4[32768] = v2
    switch -5 {
    case -719:
        v30 := []int{417, 4294967295, 9223372036854775807, 2147483647}
        _ = v30
    case 1:
        fmt.Println(v5)
    default:
        v5 = v3
    }
    v3 = (v5 & (v3 % ((548) | 1)))
    _ = len(v4)
    v29 += (0.0 - (v1 + 0.0))
    var v31 int64 = -111
    fmt.Println(v31, v3)
    v32 := []int{1, 32768, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
