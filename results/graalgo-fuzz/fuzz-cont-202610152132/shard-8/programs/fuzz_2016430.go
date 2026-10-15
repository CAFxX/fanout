package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int = -283
    var v2 int16 = 42
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 int = (v1 - v1)
    switch ("foo bar" + "") {
    case "hello":
        v1 = -2147483648
    case "a":
        fmt.Println(v3, v3, v3)
    case "":
        var v6 int64 = 9223372036854775807
        _ = v6
    }
    v3 = (-811 >= 42)
    v7 := GenS1{F0: 255}
    if ((v2 & (^(v2))) > v2) {
        var v8 bool = v3
        _ = v8
    } else {
        switch "foo bar" {
        case "hello":
            v9 := GenS1{F0: 605}
            _ = v9
        case "\n\t":
            if v3 {
                v1 -= (v1 / (((v1 & v5)) | 1))
                if (351 > 42) {
                    v2 -= v2
                    for fi10 := 0; fi10 < 15; fi10++ {
                        var v11 uint = 0
                        for fi12 := 0; fi12 < 13; fi12++ {
                            v1 -= (^(v5))
                            v5 = v1
                            v11 -= v11
                            _ = fi12
                        }
                        _ = fi10
                        _ = v11
                    }
                } else {
                    fmt.Println(v1, v2)
                    var v13 string = ""
                    v13 = "\n\t"
                    _ = v13
                }
            } else {
                var v14 uint = 42
                _ = v14
            }
        default:
            v1 += (v1 / (((v1 + int(v2))) | 1))
        }
        _ = len(v4)
        v15 := GenS1{F0: 79}
        _ = v15
    }
    v1 -= v1
    switch "a" {
    case "hello":
        var v16 uint64 = 430
        _ = v16
    }
    v2 += -1
    if true {
        var v17 int32 = 0
        _ = v17
    } else {
        v18 := []int{(v5 % ((v1) | 1)), -758, v1}
        v5 = 38
        if (255 != 43) {
            for fi19 := 0; fi19 < 6; fi19++ {
                delete(v4, 399)
                _ = fi19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(v3)
                _ = fi20
            }
        }
        _ = v18
    }
    if v3 {
        v2 += 1000
        v21 := GenS2{F0: 577, F1: -9223372036854775808, F2: -1}
        if (65535 < 377) {
            v3 = false
            v22 := GenS3{F0: 437, F1: 18}
            if (1 > 663) {
                fmt.Println(v1, v2)
                var v23 bool = true
                _ = v23
            } else {
                var v24 float32 = (float32((((3.14159 / -2.25) + 0.0) / (1e10 - 0.0))) * (0.1 + 0.0))
                if false {
                    v2 += (v2 ^ v2)
                }
                _ = v24
            }
            _ = v22
        } else {
            var v25 uint32 = 4294967295
            if (42 == 18446744073709551615) {
                fmt.Println(v25, v2, v3)
            }
            _ = v25
        }
        _ = v21
    } else {
        switch "a" {
        case "\n\t":
            if (-9223372036854775808 == 535) {
                var v26 string = (("hello" + "foo bar") + "\n\t")
                _ = v26
            }
        }
        v27 := []int{v5, (v1 | 293), v5, (v5 & (v5 % ((32768) | 1)))}
        _ = v27
    }
    var v28 bool = true
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 int8 = -109
        if true {
            v31 := GenS4{F0: 16}
            switch "\n\t" {
            case "":
                switch "hello" {
                case "hello":
                    fmt.Println(v1, v30, v3)
                default:
                    switch (("hello" + "hello") + "") {
                    case "\n\t":
                        switch "" {
                        case "\n\t":
                            v3 = v28
                        case "foo bar":
                            var v32 int64 = 1
                            _ = v32
                        default:
                            v33 := GenS3{F0: 42, F1: (v30 / (((^(v30))) | 1))}
                            _ = v33
                        }
                    default:
                        if true {
                            var v34 uint16 = 1
                            _ = v34
                        } else {
                            v28 = true
                            v3 = v3
                        }
                    }
                }
            default:
                v35 := GenS3{F0: 42, F1: v30}
                _ = v35
            }
            _ = v31
        }
        v2 += v2
        _ = fi29
        _ = v30
    }
    v1 = v5
    var v36 uint8 = 150
    if ((false && v28) && (v3 && true)) {
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v7
}
