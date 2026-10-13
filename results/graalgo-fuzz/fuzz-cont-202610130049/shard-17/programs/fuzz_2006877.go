package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint8) uint {
    var v4 int = 945
    fmt.Println(p0, v4)
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = 42
        _ = fi5
    }
    _ = p0
    _ = v4
    return uint(p0)
}

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    switch "a" {
    case "hello":
        v1 = ((3.14159 / 3.14159) < ((float32(0.1) * 0.0) / float32((0.0 + 0.1))))
    default:
        var v6 int = 98
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint = 556
        v11 := GenS1{F0: 1000, F1: 2147483648}
        fmt.Println(v1, v1)
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 = true
    for fi12 := 0; fi12 < 11; fi12++ {
        switch fi12 {
        case 966:
            fmt.Println(fi12)
        case 32767:
            var v13 int16 = 0
            _ = v13
        case 258:
            if v1 {
                v1 = (195 < 190)
            }
        default:
            if (true && (0 > -32768)) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 uint8 = 0
                    var v16 string = ""
                    v17 := GenS1{F0: -712, F1: 351}
                    if v1 {
                        var v18 float32 = (0.0 - (float32(0.0) / 1.5))
                        v19 := GenS1{F0: -32768, F1: 1}
                        v16 = (v16 + "\n\t")
                        _ = v18
                        _ = v19
                    } else {
                        v1 = false
                    }
                    v2[(fi14 / ((fi12) | 1))] = v1
                    v20 := GenS1{F0: -1, F1: 0}
                    v16 = "foo bar"
                    v21 := GenS2{F0: ("foo bar" + v16), F1: false}
                    switch (-(fi12)) {
                    case 4294967295:
                        v15 = v15
                    case 1:
                        switch v16 {
                        case "foo bar":
                            v22 := GenS2{F0: v16, F1: (v1 || (-1 >= -361))}
                            _ = v22
                        default:
                            switch (fi14 & -9223372036854775808) {
                            case 2147483647:
                                var v23 uint64 = 1
                                _ = v23
                            }
                        }
                    case 2147483647:
                        if v1 {
                            v15 -= (v15 + v15)
                            switch (fi12 / (((fi14 ^ int(v15))) | 1)) {
                            case 229:
                                v24 := GenS3{F0: uint32(v15)}
                                _ = v24
                            default:
                                _ = len(v3)
                            }
                            switch v16 {
                            case "foo bar":
                                if v1 {
                                    switch v16 {
                                    case "":
                                        if true {
                                            var v25 uint32 = 4294967295
                                            switch v16 {
                                            case "foo bar":
                                                var v26 uint8 = (v15 >> ((0) & 63))
                                                _ = v26
                                            case "hello":
                                                v27 := GenS3{F0: v25}
                                                _ = v27
                                            case "":
                                                v28 := []int{}
                                                _ = v28
                                            default:
                                                v15 = 2
                                            }
                                            var v29 uint32 = v25
                                            _ = v25
                                            _ = v29
                                        }
                                    default:
                                        if false {
                                            v16 = (v16 + (v16 + "\n\t"))
                                        }
                                    }
                                    v15 = 127
                                    var v30 float64 = 0.0
                                    _ = v30
                                }
                            case "hello":
                                v31 := GenS1{F0: 0, F1: 144}
                                _ = v31
                            case "":
                                fmt.Println(v16, fi12)
                            }
                        } else {
                            var v32 int8 = 42
                            _ = v32
                        }
                    default:
                        if v1 {
                            v33 := GenS3{F0: 2}
                            var v34 float64 = ((0.0 / 0.1) - 1.5)
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v20
                    _ = v21
                }
                switch "\n\t" {
                case "foo bar":
                    v35 := GenS4{F0: v1}
                    _ = v35
                case "hello":
                    v1 = false
                case "":
                    v36 := []int{}
                    _ = v36
                default:
                    v37 := GenS5{F0: 2}
                    _ = v37
                }
            } else {
                v1 = v1
                if ((1e10 <= 3.14159) && false) {
                    var v38 int = -9223372036854775808
                    v1 = ("a" != "a")
                    _ = v38
                }
            }
        }
        var v39 uint16 = 0
        _ = fi12
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        if (1.5 >= (float32(1.5) * (-2.25 / 0.1))) {
            var v41 bool = true
            v42 := GenS6{F0: false}
            _ = v41
            _ = v42
        } else {
            v1 = true
        }
        _ = fi40
    }
    v1 = (-9223372036854775808 < 0)
    for fi43 := 0; fi43 < 18; fi43++ {
        v44 := GenS5{F0: 18446744073709551615}
        _ = fi43
        _ = v44
    }
    switch 128 {
    case -1:
        var v45 uint = 524
        _ = v45
    case 127:
        v1 = (127 >= 127)
    case 229:
        v46 := GenS1{F0: -863, F1: 819}
        _ = v46
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = v1
    v47 := GenS1{F0: 1, F1: 2}
    var v48 int32 = -2147483648
    var v49 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
}
