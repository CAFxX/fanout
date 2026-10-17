package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int64; F2 int16 }

func main() {
    var v1 bool = (959 < 2147483647)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 float64 = (float64(-2.25) + ((0.1 + (float64(3.14159) - 1.5)) * float64(0.0)))
    v6 := GenS1{F0: 42}
    if v1 {
        if v1 {
            v2[2147483647] = 65535
            var v7 bool = true
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                fmt.Println(v5)
                delete(v2, (fi8 / (((fi8 & -965)) | 1)))
                _ = fi8
            }
        }
        v9 := GenS2{F0: v1, F1: (1 == 411)}
        for fi10 := 0; fi10 < 12; fi10++ {
            if v1 {
                var v11 uint = 2
                var v12 int16 = -1
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS3{F0: "a", F1: 667, F2: 42}
                    v5 = v5
                    _ = fi13
                    _ = v14
                }
                switch (^(fi10)) {
                case 32768:
                    var v15 uint32 = 0
                    _ = v15
                case -2147483648:
                    switch "a" {
                    case "hello":
                        if (669 == 18446744073709551615) {
                            if true {
                                v16 := []int{(fi10 & -600), (fi10 / ((fi10) | 1)), fi10, (fi10 ^ (-(fi10)))}
                                fmt.Println(v1, v1)
                                _ = len(v16)
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 4; fi17++ {
                                    switch -1 {
                                    case -467:
                                        v5 -= 0.1
                                    default:
                                        var v18 float32 = 0.0
                                        if (v1 || ((0.1 < 1.5) && v1)) {
                                            var v19 uint16 = 65535
                                            _ = v19
                                        } else {
                                            var v20 uint32 = 2147483647
                                            _ = v20
                                        }
                                        fmt.Println(v5)
                                        if ((v18 + v18) <= v18) {
                                            switch "hello" {
                                            case "hello":
                                                var v21 int32 = 1
                                                _ = v21
                                            }
                                        }
                                        _ = v18
                                    }
                                    v22 := GenS3{F0: ("a" + ""), F1: 718, F2: 18446744073709551615}
                                    v23 := GenS2{F0: v1, F1: v1}
                                    _ = fi17
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS4{F0: "\n\t"}
                                _ = v24
                            }
                        } else {
                            switch 9223372036854775807 {
                            case 4294967295:
                                var v25 int = fi10
                                _ = v25
                            }
                            if true {
                                var v26 uint = 65535
                                _ = v26
                            }
                        }
                    case "a":
                        v27 := GenS1{F0: -121}
                        _ = v27
                    default:
                        v28 := []int{(-(fi10)), fi10, 32767, 2147483647}
                        _ = v28
                    }
                }
            }
            v29 := GenS5{F0: 1, F1: 42}
            fmt.Println(fi10, v1)
            _ = fi10
            _ = v29
        }
        _ = v9
    }
    v3[790] = v1
    v3[9223372036854775807] = false
    var v30 float32 = -2.25
    fmt.Println(v5, v5, v5)
    var v31 uint16 = 106
    if (2147483647 < -1) {
        var v32 uint8 = 1
        fmt.Println(v31, v5)
        v33 := GenS6{F0: uint32(v31)}
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            fmt.Println(v31, v30, v31)
            var v35 int = -683
            v35 = (fi34 % (((v35 + 0)) | 1))
            _ = fi34
            _ = v35
        }
    }
    var v36 int32 = -2147483648
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 int16 = 861
        _ = fi37
        _ = v38
    }
    v39 := GenS7{F0: 267, F1: -9223372036854775808, F2: 444}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
