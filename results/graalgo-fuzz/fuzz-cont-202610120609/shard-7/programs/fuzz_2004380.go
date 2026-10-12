package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int16 }
type GenS5 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 32767
    v3 := []int{}
    var v4 int = -1
    fmt.Println(v4, v1, v1)
    v1 += v1
    for fi5 := 0; fi5 < 8; fi5++ {
        v3 = append(v3, 596)
        fmt.Println(v2)
        switch fi5 {
        case 984:
            var v6 uint16 = v2
            _ = v6
        default:
            v7 := GenS1{F0: 4294967295}
            _ = v7
        }
        _ = fi5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = "hello"
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case -936:
                var v11 bool = false
                v3 = append(v3, fi10)
                v4 = fi8
                _ = v11
            case 1:
                var v12 int = (v4 / (((v4 * (fi8 % ((fi8) | 1)))) | 1))
                _ = v12
            case -9223372036854775808:
                v1 -= (v1 ^ 18446744073709551615)
            default:
                switch (v4 + 272) {
                case 110:
                    switch fi8 {
                    case 1:
                        v13 := GenS2{F0: 821, F1: (fi10 + (fi8 | -2147483648)), F2: 65535}
                        _ = v13
                    case 32:
                        if false {
                            var v14 float64 = float64(float32(0.1))
                            switch fi8 {
                            case -311:
                                fmt.Println(fi8, v2)
                            case 959:
                                v15 := GenS2{F0: -1, F1: fi10, F2: 65535}
                                _ = v15
                            case 32768:
                                fmt.Println(v4, fi8)
                                v16 := GenS3{F0: v2}
                                v14 = v14
                                _ = v16
                            }
                            _ = v14
                        } else {
                            v3 = append(v3, fi10)
                            var v17 uint = 42
                            _ = v17
                        }
                    case -2147483648:
                        var v18 uint = 255
                        _ = v18
                    }
                }
            }
            v3 = append(v3, fi10)
            switch (v9 + (((v9 + "foo bar") + v9) + (v9 + "hello"))) {
            case "hello":
                v2 -= v2
            case "a":
                var v19 int = 363
                _ = v19
            default:
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
            _ = fi10
        }
        v1 -= (v1 / ((1) | 1))
        _ = fi8
        _ = v9
    }
    var v21 float64 = 1e10
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS3{F0: 42}
        v3 = append(v3, (fi22 / (((v4 / (((fi22 >> ((914) & 63))) | 1))) | 1)))
        _ = fi22
        _ = v23
    }
    switch (v4 - 32768) {
    case 390:
        fmt.Println(v21, v4)
    case -690:
        v2 = (v2 ^ (v2 & (v2 << ((1) & 63))))
    default:
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS4{F0: -1, F1: 128, F2: -304}
                var v27 int8 = -1
                _ = fi25
                _ = v26
                _ = v27
            }
            if false {
                _ = len(v3)
            }
            var v28 float32 = (0.1 + 3.14159)
            _ = fi24
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(fi29)
        fmt.Println(v4, v2)
        _ = fi29
    }
    _ = len(v3)
    fmt.Println(v2)
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 uint64 = v1
        if true {
            v3 = append(v3, (^(v4)))
        } else {
            _ = len(v3)
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v21)
    v4 = 911
    if (!(-128 == -1)) {
        var v32 uint16 = 158
        if true {
            if (v4 != (v4 * (v4 - v4))) {
                if ("hello" == "hello") {
                    var v33 uint16 = (v32 << ((950) & 63))
                    _ = v33
                } else {
                    v34 := GenS5{F0: 96, F1: 1}
                    _ = v34
                }
            } else {
                v21 = -2.25
                for fi35 := 0; fi35 < 14; fi35++ {
                    fmt.Println(v32, v1, v1)
                    v2 -= 255
                    v21 += ((float64((-2.25 + float32(v21))) + v21) / (v21 - float64(1e10)))
                    _ = fi35
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch "hello" {
            case "":
                _ = len(v3)
            case "\n\t":
                var v36 float64 = v21
                _ = v36
            }
            if len(v3) > 0 {
                v3[0] = (v4 + (v4 & (^(v4))))
            }
        }
        fmt.Println(v21)
        _ = v32
    }
    var v37 string = ""
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
}
