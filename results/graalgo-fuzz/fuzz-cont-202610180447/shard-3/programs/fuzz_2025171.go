package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int8 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 string = (("" + ("\n\t" + "\n\t")) + (("a" + "foo bar") + "foo bar"))
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint = (v1 * v1)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v1 = v1
            var v9 string = v3
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch (("foo bar" + (v3 + v3)) + (v3 + (v3 + (("" + v3) + v3)))) {
    case "":
        v10 := GenS1{F0: 611}
        _ = v10
    case "foo bar":
        v11 := GenS2{F0: 0, F1: 9223372036854775807}
        _ = v11
    case "hello":
        v12 := GenS3{F0: 128, F1: 0}
        _ = v12
    }
    var v13 uint16 = 505
    v4 += v4
    v3 = (v3 + ((("hello" + v3) + (v3 + v3)) + v3))
    for fi14 := 0; fi14 < 4; fi14++ {
        switch fi14 {
        case 32768:
            for fi15 := 0; fi15 < 10; fi15++ {
                v2[2147483648] = v3
                fmt.Println(fi14)
                _ = fi15
            }
        case -1:
            fmt.Println(fi14, v13)
        default:
            var v16 string = (v3 + "foo bar")
            _ = v16
        }
        v4 += v4
        _ = fi14
    }
    var v17 string = v3
    switch 4294967295 {
    case -235:
        delete(v2, 392)
    default:
        for fi18 := 0; fi18 < 8; fi18++ {
            v4 -= 32
            if true {
                v19 := GenS3{F0: fi18, F1: 1}
                v3 = (v17 + v3)
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS4{F0: 2147483647, F1: 42, F2: 100}
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                switch "foo bar" {
                case "":
                    v22 := GenS4{F0: 4294967295, F1: 2147483647, F2: 0}
                    _ = v22
                case "\n\t":
                    switch 32767 {
                    case 1:
                        v23 := GenS1{F0: 862}
                        _ = v23
                    default:
                        switch v3 {
                        case "foo bar":
                            if (42 < -9223372036854775808) {
                                var v24 uint16 = 1
                                v25 := GenS5{F0: 42, F1: 42}
                                _ = v24
                                _ = v25
                            } else {
                                fmt.Println(v17, v1)
                            }
                        case "":
                            v26 := GenS2{F0: 9223372036854775807, F1: fi18}
                            _ = v26
                        case "hello":
                            if false {
                                var v27 int = 128
                                if true {
                                    var v28 float64 = -2.25
                                    var v29 uint16 = v13
                                    v30 := v2[2147483648]
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 2; fi31++ {
                                    v32 := v2[195]
                                    _ = fi31
                                    _ = v32
                                }
                                _ = v27
                            }
                        }
                    }
                default:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 uint8 = (v4 / (((v4 ^ (^(v4)))) | 1))
                        v35 := v2[4294967295]
                        v36 := GenS6{F0: v13}
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
                var v37 uint8 = 128
                if ((1e10 * 3.14159) <= float64((0.1 - (-2.25 * 0.1)))) {
                    v38 := GenS7{F0: 4294967295, F1: 1}
                    v39 := []int{(fi18 * -2147483648), 42}
                    _ = v38
                    _ = v39
                }
                _ = v37
            }
            fmt.Println(v4, fi18, v13)
            _ = fi18
        }
    }
    v40 := GenS8{F0: 0}
    if false {
        v4 = (v4 | v4)
    } else {
        v3 = ""
        v41 := []int{-776, -773, 127, 92}
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        for fi43 := 0; fi43 < 13; fi43++ {
            var v44 float64 = -2.25
            switch v3 {
            case "":
                _ = len(v2)
                v1 -= (v1 | v1)
            }
            _ = fi43
            _ = v44
        }
        var v45 uint8 = v4
        v46 := GenS9{F0: "hello", F1: -1}
        _ = fi42
        _ = v45
        _ = v46
    }
    var v47 int16 = 840
    fmt.Println(v3)
    v48 := []int{int(v4), 42, 32768}
    fmt.Println(v13, v47)
    for fi49 := 0; fi49 < 7; fi49++ {
        v50 := v2[uint32(v1)]
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
}
