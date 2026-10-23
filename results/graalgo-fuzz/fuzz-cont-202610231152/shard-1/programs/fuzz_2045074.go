package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int8; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint16; F2 int }
type GenS5 struct { F0 int8; F1 int8; F2 string }
type GenS6 struct { F0 uint; F1 int32 }

func main() {
    v1 := []int{2147483647, 9223372036854775807}
    var v2 int = 127
    var v3 uint = 42
    fmt.Println(v3)
    if (1.5 == 1.5) {
        v2 += 32767
        if ((!true) || false) {
            v4 := GenS1{F0: true, F1: 2}
            for fi5 := 0; fi5 < 8; fi5++ {
                _ = len(v1)
                if (0 <= 1) {
                    var v6 uint8 = 83
                    if (v2 <= -1) {
                        v2 = int(v3)
                        v6 += v6
                    }
                    _ = v6
                }
                var v7 bool = true
                _ = fi5
                _ = v7
            }
            var v8 uint8 = 0
            _ = v4
            _ = v8
        } else {
            fmt.Println(v3, v3)
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    if true {
                        var v10 uint16 = 255
                        _ = v10
                    } else {
                        fmt.Println(v2, v2)
                        if len(v1) > 0 {
                            v1[0] = fi9
                        }
                        if false {
                            v3 = v3
                        }
                    }
                    v3 = (v3 / ((2) | 1))
                    _ = fi9
                }
            } else {
                v3 += v3
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v3 = (v3 * v3)
                        switch "a" {
                        case "a":
                            var v12 float64 = ((float64((1.5 / 3.14159)) * 1.5) * (3.14159 + 1e10))
                            _ = v12
                        case "":
                            v2 = fi11
                        }
                        _ = fi11
                    }
                    var v13 float32 = ((0.1 + 1e10) + -2.25)
                    _ = v13
                }
            }
        }
    }
    v3 = v3
    if (!(541 <= -1)) {
        fmt.Println(v2, v3)
        fmt.Println(v3, v2)
        fmt.Println(v2)
    }
    if true {
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                v1 = append(v1, fi14)
                v3 += v3
                if (true || true) {
                    switch "a" {
                    case "hello":
                        var v15 uint = v3
                        _ = v15
                    case "a":
                        var v16 bool = false
                        _ = v16
                    case "":
                        v17 := GenS1{F0: true, F1: 42}
                        _ = v17
                    }
                    switch "a" {
                    case "hello":
                        v2 += (^(v2))
                    case "a":
                        v18 := GenS2{F0: 1, F1: -97, F2: (fi14 << (((v3 & v3)) & 63))}
                        _ = v18
                    case "":
                        if false {
                            v3 = (v3 / ((1) | 1))
                            v3 = 4294967295
                            v19 := GenS3{F0: 4294967295}
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                _ = len(v1)
                                v1 = append(v1, fi14)
                                _ = fi20
                            }
                            v3 += v3
                        }
                    default:
                        v21 := GenS3{F0: 42}
                        _ = v21
                    }
                    if (v3 >= v3) {
                        v1 = append(v1, (v2 & fi14))
                    } else {
                        if true {
                            var v22 float32 = 1.5
                            v22 = 1.5
                            _ = len(v1)
                            _ = v22
                        } else {
                            v23 := GenS4{F0: "a", F1: 2, F2: (v2 % ((fi14) | 1))}
                            for fi24 := 0; fi24 < 10; fi24++ {
                                var v25 float32 = 1e10
                                v2 += (fi14 - fi24)
                                var v26 bool = true
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                        }
                        switch (fi14 | v2) {
                        case 0:
                            for fi27 := 0; fi27 < 14; fi27++ {
                                var v28 int = (fi27 % ((v2) | 1))
                                fmt.Println(v2)
                                var v29 float32 = (3.14159 - (0.1 * 1.5))
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        default:
                            var v30 string = "a"
                            _ = v30
                        }
                        fmt.Println(v3, v2, v3)
                    }
                }
                _ = fi14
            }
            fmt.Println(v3)
        }
        v31 := GenS2{F0: 0, F1: -19, F2: (v2 << ((255) & 63))}
        if false {
            v2 = (v2 % (((-(v2))) | 1))
        }
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := GenS1{F0: true, F1: 1}
        v2 = (^(v2))
        var v34 int16 = 42
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: v2}
    fmt.Println(v2, v2)
    v36 := GenS1{F0: true, F1: 464}
    fmt.Println(v3, v2)
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    v37 := GenS1{F0: true, F1: 98}
    fmt.Println(v3, v3)
    if (!(!false)) {
        v38 := GenS5{F0: -128, F1: -1, F2: "a"}
        if false {
            v3 -= (^(v3))
            v2 = (v2 & (^(v2)))
            _ = len(v1)
        } else {
            var v39 uint16 = 65535
            var v40 int16 = 42
            if true {
                v39 = (v39 % ((65535) | 1))
                v41 := GenS6{F0: 837, F1: -1}
                _ = v41
            }
            _ = v39
            _ = v40
        }
        var v42 int32 = -978
        _ = v38
        _ = v42
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
