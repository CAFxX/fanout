package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    fmt.Println(v3, v3)
    switch (("\n\t" + "a") + "a") {
    case "foo bar":
        v1 = v1
    case "":
        v3 = 42
    case "hello":
        fmt.Println(v3)
    default:
        v4 := GenS1{F0: (v3 ^ v3)}
        _ = v4
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 = (-1 == 32767)
        fmt.Println(v1)
        _ = fi5
    }
    var v6 float64 = ((1e10 + float64(0.1)) - 1.5)
    v6 = (float64(1e10) / v6)
    if (!v1) {
        v7 := v2[2]
        _ = v7
    } else {
        var v8 int = 42
        _ = v8
    }
    fmt.Println(v1, v1)
    if (!v1) {
        if (255 < 255) {
            v9 := GenS1{F0: (v3 >> ((255) & 63))}
            v1 = v1
            _ = v9
        }
        switch (("foo bar" + "\n\t") + ((("a" + "hello") + "a") + "\n\t")) {
        case "\n\t":
            if true {
                fmt.Println(v3, v6, v1)
                for fi10 := 0; fi10 < 3; fi10++ {
                    _ = len(v2)
                    v11 := GenS1{F0: v3}
                    _ = fi10
                    _ = v11
                }
            } else {
                if (false || true) {
                    delete(v2, 42)
                    if v1 {
                        fmt.Println(v6, v3)
                        v6 = (v6 + (v6 * float64(1e10)))
                        var v12 bool = (false || v1)
                        _ = v12
                    }
                    var v13 bool = (2 <= 127)
                    _ = v13
                }
                v3 += (v3 << ((65535) & 63))
                v1 = true
            }
        case "foo bar":
            if v1 {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14)
                    _ = fi14
                }
                v2[2] = 1
            } else {
                v1 = true
                if ((-871 == 1000) && false) {
                    v15 := GenS1{F0: v3}
                    if true {
                        v2[2] = 9223372036854775807
                        var v16 int = -185
                        _ = v16
                    }
                    _ = v15
                } else {
                    if (true || v1) {
                        var v17 int = -1
                        if v1 {
                            var v18 int32 = 0
                            _ = v18
                        } else {
                            v19 := GenS1{F0: (v3 / ((-9223372036854775808) | 1))}
                            for fi20 := 0; fi20 < 9; fi20++ {
                                v21 := []int{}
                                var v22 uint8 = 37
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    if false {
                                        v3 = (v3 >> ((255) & 63))
                                        v17 -= (fi23 | (v17 / (((fi20 * fi20)) | 1)))
                                        v24 := GenS2{F0: 42, F1: 549, F2: ("\n\t" + "a")}
                                        var v25 bool = v1
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 int8 = int8(v3)
                                    _ = fi23
                                    _ = v26
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        }
                        _ = v17
                    }
                }
            }
        case "":
            switch "hello" {
            case "foo bar":
                for fi27 := 0; fi27 < 11; fi27++ {
                    v3 = v3
                    fmt.Println(v6, fi27, fi27)
                    var v28 int8 = -128
                    _ = fi27
                    _ = v28
                }
            case "hello":
                v1 = true
            case "\n\t":
                switch 289 {
                case -9223372036854775808:
                    v6 -= 0.1
                case -866:
                    v1 = v1
                case 1:
                    for fi29 := 0; fi29 < 1; fi29++ {
                        for fi30 := 0; fi30 < 15; fi30++ {
                            if v1 {
                                v3 = v3
                            }
                            var v31 int = (fi30 & fi30)
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                }
            }
        default:
            v2[689] = -9223372036854775808
        }
    }
    var v32 string = ("\n\t" + "")
    v33 := []int{-898}
    for fi34 := 0; fi34 < 2; fi34++ {
        v3 = -37
        v35 := GenS1{F0: v3}
        for fi36 := 0; fi36 < 17; fi36++ {
            v33 = append(v33, 4294967295)
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    v37 := GenS3{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v6
}
