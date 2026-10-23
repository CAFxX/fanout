package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint16; F1 int8; F2 int16 }
func genFn0(p0 uint16) uint64 {
    if ("foo bar" != "hello") {
        var v5 int16 = 42
        v6 := GenS1{F0: 584, F1: false, F2: 786}
        _ = v5
        _ = v6
    } else {
        switch -922 {
        case 32768:
            v7 := GenS1{F0: 42, F1: true, F2: 660}
            _ = v7
        default:
            p0 += p0
        }
    }
    var v8 float64 = (-2.25 * (1.5 * ((float64(3.14159) * ((1e10 * 1e10) * ((0.0 - 1.5) * (0.1 * float64((1.5 - 3.14159)))))) * -2.25)))
    var v9 string = "a"
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    var v4 uint = 2
    v3 = 1
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int64 = int64(v1)
        v12 := v2[uint32(v4)]
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        if (-93 >= 79) {
            fmt.Println(v3, v4)
        } else {
            var v14 int32 = v3
            var v15 int16 = 279
            v15 = (v15 & v15)
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v1 = (^(v1))
    if false {
        if false {
            fmt.Println(v1, v3)
        }
        fmt.Println(v3, v1)
        for fi16 := 0; fi16 < 13; fi16++ {
            fmt.Println(v3, v3)
            switch "hello" {
            case "foo bar":
                if false {
                    delete(v2, v1)
                } else {
                    var v17 int64 = 117
                    v3 += v3
                    _ = v17
                }
            case "a":
                if true {
                    switch "\n\t" {
                    case "hello":
                        v3 -= (^(v3))
                    default:
                        if ("hello" < ("foo bar" + "")) {
                            if true {
                                var v18 uint32 = v1
                                _ = v18
                            } else {
                                if (1.5 == (3.14159 * (0.0 + (float32(float64(float32(1e10))) / 1.0)))) {
                                    var v19 uint8 = 1
                                    v20 := GenS1{F0: 42, F1: ((v4 ^ v4) < uint(v1)), F2: (^(fi16))}
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint32 = v1
                                    for fi22 := 0; fi22 < 3; fi22++ {
                                        v4 = (v4 % ((v4) | 1))
                                        _ = fi22
                                    }
                                    _ = v21
                                }
                            }
                            var v23 int64 = -1
                            fmt.Println(fi16, v1, fi16)
                            _ = v23
                        } else {
                            v24 := []int{fi16, 32767, fi16, (fi16 | fi16)}
                            var v25 uint = v4
                            var v26 float64 = 3.14159
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 int8 = -86
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint64 = 814
                    _ = v29
                } else {
                    v3 = 0
                }
            case "hello":
                var v30 float32 = float32((1.5 + 1e10))
                _ = v30
            }
            _ = fi16
        }
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 int = fi31
        _ = fi31
        _ = v32
    }
    if false {
        var v33 int = -141
        if (-32768 < 32767) {
            v34 := []int{v33, (v33 | v33)}
            if (false && true) {
                v3 = 0
                v4 -= (v4 - v4)
            }
            _ = v34
        }
        _ = v33
    } else {
        v35 := GenS2{F0: 4294967295, F1: 42}
        if (!true) {
            fmt.Println(v3, v3, v4)
        } else {
            v36 := GenS1{F0: -237, F1: (1 <= 1), F2: 0}
            if true {
                v37 := []int{9223372036854775807, 4294967295}
                fmt.Println(v4)
                _ = v37
            } else {
                v38 := []int{2147483647, -1, 2147483647}
                _ = v38
            }
            for fi39 := 0; fi39 < 12; fi39++ {
                var v40 int = fi39
                if false {
                    v1 = v1
                }
                _ = fi39
                _ = v40
            }
            _ = v36
        }
        _ = v35
    }
    v41 := GenS3{F0: 42, F1: 127, F2: 29}
    v42 := GenS1{F0: 0, F1: (0.0 == 0.1), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
