package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 int8, p1 int16, p2 int) uint64 {
    var v4 uint16 = 530
    v5 := []int{0, (^(p2)), p2}
    _ = len(v5)
    v6 := GenS1{F0: 255, F1: "foo bar"}
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, p0, fi7)
        if (42 >= 2147483647) {
            if (0 != 557) {
                var v8 uint32 = 486
                _ = v8
            } else {
                p0 -= -116
                v9 := GenS2{F0: p2, F1: 42}
                v10 := GenS3{F0: -2147483648, F1: 1, F2: -114}
                _ = v10
                _ = v9
            }
            fmt.Println(p1, fi7, p0)
        } else {
            switch "hello" {
            case "foo bar":
                v5 = append(v5, p2)
            default:
                switch "" {
                case "foo bar":
                    var v11 uint32 = 81
                    _ = v11
                case "hello":
                    for fi12 := 0; fi12 < 17; fi12++ {
                        if false {
                            var v13 float64 = (3.14159 - 0.1)
                            if (!(0 == 957)) {
                                fmt.Println(p2, fi7)
                            } else {
                                v4 = 0
                                _ = len(v5)
                                if (916 != fi7) {
                                    switch ("a" + "a") {
                                    case "a":
                                        if (!(((3.14159 - 1.5) - 0.1) != 3.14159)) {
                                            if ((0.0 + 1e10) == float32((v13 - v13))) {
                                                var v14 int32 = 42
                                                _ = v14
                                            } else {
                                                var v15 int32 = 665
                                                v15 -= (v15 >> ((65535) & 63))
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v15
                                            }
                                            switch 32767 {
                                            case 42:
                                                var v16 uint = 18446744073709551615
                                                _ = v16
                                            default:
                                                switch "foo bar" {
                                                case "hello":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    p2 += -2147483648
                                                }
                                            }
                                            p2 += p2
                                        }
                                    }
                                    _ = len(v5)
                                }
                            }
                            _ = v13
                        } else {
                            v17 := GenS1{F0: 802, F1: ""}
                            v4 += v4
                            _ = v17
                        }
                        p1 += p1
                        _ = fi12
                    }
                case "a":
                    v4 += (v4 << ((871) & 63))
                default:
                    var v18 int32 = 1
                    _ = v18
                }
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int32) uint32 {
    if (407 != 42) {
        var v19 uint16 = 1
        v20 := GenS4{F0: -1, F1: false, F2: ""}
        v19 = v19
        _ = v19
        _ = v20
    } else {
        p0 = p0
        var v21 uint = 747
        switch ("" + "a") {
        case "":
            v21 += v21
        default:
            switch 559 {
            case -1:
                fmt.Println(v21, p0)
            default:
                var v22 uint64 = 4294967295
                _ = v22
            }
        }
        _ = v21
    }
    var v23 uint32 = 2147483648
    switch "a" {
    case "a":
        var v24 int = -2147483648
        _ = v24
    case "":
        if false {
            v25 := []int{127, 121}
            _ = v25
        }
    default:
        v26 := GenS3{F0: p0, F1: 42, F2: 127}
        _ = v26
    }
    _ = p0
    _ = v23
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v27 int = 1
    v27 += 2147483647
    v27 = 32768
    switch "a" {
    case "":
        _ = len(v3)
    case "\n\t":
        v28 := v3["foo bar"]
        _ = v28
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := []int{}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        fmt.Println(fi31, fi31)
        var v32 uint8 = 133
        _ = fi31
        _ = v32
    }
    var v33 int64 = -1
    if false {
        if false {
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 int64 = (v33 % (((v33 + (v33 + (v33 << ((302) & 63))))) | 1))
                switch "\n\t" {
                case "hello":
                    var v36 uint16 = 255
                    _ = v36
                }
                v37 := []int{(v27 >> ((1) & 63))}
                _ = fi34
                _ = v35
                _ = v37
            }
        }
        switch "foo bar" {
        case "\n\t":
            v27 = (^(v27))
        default:
            v38 := GenS1{F0: 65535, F1: ("a" + "hello")}
            _ = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
}
