package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 bool; F1 int8; F2 uint64 }
type GenS6 struct { F0 string; F1 int64; F2 uint8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    v4 := v1[0]
    for fi5 := 0; fi5 < 5; fi5++ {
        if v4 {
            v3 = v3
            var v6 uint32 = 455
            _ = v6
        } else {
            var v7 float32 = (float32(-2.25) / -2.25)
            _ = v7
        }
        v4 = v4
        var v8 int64 = 9223372036854775807
        _ = fi5
        _ = v8
    }
    if v4 {
        switch "hello" {
        case "\n\t":
            v9 := GenS1{F0: -11, F1: 255, F2: 186}
            _ = v9
        case "hello":
            fmt.Println(v3)
        case "":
            for fi10 := 0; fi10 < 7; fi10++ {
                v4 = v4
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(fi10, fi10, fi10)
                    _ = fi11
                }
                var v12 bool = ((fi10 == fi10) || v4)
                _ = fi10
                _ = v12
            }
        default:
            v13 := GenS2{F0: 32768}
            _ = v13
        }
    }
    var v14 int16 = -707
    var v15 uint8 = 128
    if v4 {
        switch "hello" {
        case "hello":
            _ = len(v1)
        case "\n\t":
            v1[2147483647] = true
        }
    }
    v3 -= 887
    v16 := GenS3{F0: 42, F1: 950, F2: "foo bar"}
    v17 := []int{}
    switch int(v15) {
    case 4294967295:
        switch int(v15) {
        case 127:
            for fi18 := 0; fi18 < 19; fi18++ {
                if v4 {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 int16 = (v14 >> (((v3 % ((v3) | 1))) & 63))
                        _ = fi19
                        _ = v20
                    }
                    switch "\n\t" {
                    case "hello":
                        for fi21 := 0; fi21 < 12; fi21++ {
                            fmt.Println(fi18)
                            var v22 uint = (v3 ^ 889)
                            _ = fi21
                            _ = v22
                        }
                    case "\n\t":
                        v23 := GenS4{F0: v15, F1: 4294967295}
                        _ = v23
                    }
                    if (!v4) {
                        if v4 {
                            fmt.Println(v4, v14, v14)
                            fmt.Println(v3, v15)
                            fmt.Println(v4, v3, v4)
                        }
                        switch fi18 {
                        case 0:
                            v24 := GenS3{F0: 0, F1: 18446744073709551615, F2: ("a" + "a")}
                            _ = v24
                        case 828:
                            var v25 uint = v3
                            _ = v25
                        default:
                            switch (("" + "hello") + ("hello" + "\n\t")) {
                            case "\n\t":
                                if v4 {
                                    v26 := GenS5{F0: false, F1: 42, F2: 917}
                                    v15 += (v15 * 127)
                                    v4 = false
                                    _ = v26
                                }
                            default:
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    v3 -= v3
                                    _ = fi27
                                }
                            }
                        }
                    }
                } else {
                    v4 = (false && (0 >= 1))
                    v28 := v2[(fi18 ^ 128)]
                    _ = v28
                }
                _ = fi18
            }
        }
    case -9223372036854775808:
        v14 = (v14 % ((707) | 1))
    case 261:
        var v29 int64 = 783
        _ = v29
    }
    var v30 int64 = 1
    switch 4294967295 {
    case 32767:
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS6{F0: "hello", F1: (^(v30)), F2: 2}
            v17 = append(v17, (fi31 >> (((^(v3))) & 63)))
            _ = fi31
            _ = v32
        }
    default:
        var v33 float64 = 0.0
        _ = v33
    }
    switch "" {
    case "":
        switch 128 {
        case 2147483647:
            var v34 int64 = v30
            _ = v34
        case 469:
            v35 := GenS7{F0: 100, F1: true}
            _ = v35
        case 0:
            v36 := GenS1{F0: -128, F1: (^(v15)), F2: -9223372036854775808}
            _ = v36
        default:
            fmt.Println(v4)
        }
    default:
        if false {
            _ = len(v2)
            switch 4294967295 {
            case -645:
                v37 := GenS8{F0: 255, F1: v4}
                _ = v37
            default:
                var v38 int64 = (v30 - v30)
                _ = v38
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
