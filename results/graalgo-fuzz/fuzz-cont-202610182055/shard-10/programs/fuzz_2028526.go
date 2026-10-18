package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = ("\n\t" + "hello")
    var v2 int16 = -640
    var v3 int32 = -1
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            if true {
                fmt.Println(fi4)
            } else {
                switch "\n\t" {
                case "a":
                    v6 := []int{743}
                    _ = v6
                case "":
                    if true {
                        var v7 int8 = -1
                        v8 := GenS1{F0: (fi5 << ((610) & 63))}
                        var v9 float64 = (1.5 * (0.0 - (3.14159 + ((1e10 / float64(1.5)) / 1.0))))
                        _ = v7
                        _ = v8
                        _ = v9
                    } else {
                        v10 := []int{}
                        _ = v10
                    }
                default:
                    v11 := GenS2{F0: 65535, F1: false}
                    var v12 uint64 = 826
                    var v13 uint32 = 1
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v1 = v1
            }
            _ = fi5
        }
        _ = fi4
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v3 = (v3 / ((v3) | 1))
        v15 := []int{(fi14 & fi14), (fi14 << ((303) & 63)), fi14, -802}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v1, v1)
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS3{F0: v1, F1: fi16}
        if (0.0 >= 1.5) {
            if true {
                if false {
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v3 += v3
                        var v19 int64 = 42
                        var v20 uint32 = 0
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 float64 = (0.0 - 1e10)
                        if (false && true) {
                            v23 := GenS4{F0: 569}
                            switch (v1 + "hello") {
                            case "hello":
                                var v24 float64 = v22
                                _ = v24
                            case "\n\t":
                                var v25 uint64 = 928
                                _ = v25
                            default:
                                if (!false) {
                                    switch 9223372036854775807 {
                                    case 0:
                                        v1 = v1
                                    case 1:
                                        if (75 >= 0) {
                                            var v26 float32 = 0.1
                                            _ = v26
                                        } else {
                                            v2 -= v2
                                        }
                                    default:
                                        var v27 uint = 65535
                                        _ = v27
                                    }
                                    v1 = ""
                                }
                            }
                            _ = v23
                        } else {
                            v28 := GenS1{F0: fi16}
                            v3 -= (v3 % ((v3) | 1))
                            v29 := GenS5{F0: -9223372036854775808, F1: 32}
                            _ = v28
                            _ = v29
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                v1 = "\n\t"
                if true {
                    var v30 uint64 = 4294967295
                    v31 := []int{-9223372036854775808, (fi16 + (fi16 ^ int(v30))), -9223372036854775808, (fi16 >> ((2) & 63))}
                    _ = v30
                    _ = v31
                }
            }
            v2 = (v2 & (v2 | (v2 | v2)))
        } else {
            if true {
                v32 := GenS1{F0: 32767}
                _ = v32
            }
            var v33 uint = 571
            _ = v33
        }
        fmt.Println(v1, fi16)
        _ = fi16
        _ = v17
    }
    v34 := []int{-248}
    if false {
        v35 := GenS6{F0: v2, F1: (v3 * (v3 % ((v3) | 1)))}
        var v36 uint = 4294967295
        var v37 int = -1
        _ = v35
        _ = v36
        _ = v37
    }
    var v38 float64 = -2.25
    v34 = append(v34, -1)
    v38 = (float64(float32(v38)) + 3.14159)
    var v39 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
}
