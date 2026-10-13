package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint = 0
    v2 := []uint32{2147483647}
    var v3 float64 = (1.5 - (0.1 / 1.0))
    v4 := make(map[string]uint32)
    var v5 int16 = -1
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 16; fi7++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = fi7
    }
    v8 := GenS1{F0: -27, F1: 128, F2: 72}
    var v9 uint16 = 947
    if false {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int8 = 127
            if true {
                fmt.Println(v1, v9, v3)
                switch 2147483647 {
                case 0:
                    fmt.Println(v11)
                case 2147483647:
                    v12 := GenS2{F0: "\n\t"}
                    _ = v12
                }
                var v13 string = ""
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v1, v3, v5)
        if (2147483647 < 450) {
            fmt.Println(v3, v3)
            v9 = 32767
        } else {
            v14 := v4[(("" + "a") + (("hello" + "") + "a"))]
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS3{F0: -202, F1: 2, F2: 930}
        fmt.Println(v9)
        _ = fi15
        _ = v16
    }
    if (-70 <= 0) {
        for fi17 := 0; fi17 < 18; fi17++ {
            if true {
                v18 := GenS2{F0: ""}
                if false {
                    if (!false) {
                        fmt.Println(fi17)
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 uint16 = 0
                            switch "a" {
                            case "":
                                delete(v6, (fi19 | -170))
                            case "\n\t":
                                var v21 int8 = 42
                                _ = v21
                            default:
                                fmt.Println(fi19, v3, v5)
                            }
                            _ = fi19
                            _ = v20
                        }
                        v5 = int16(fi17)
                    }
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS1{F0: 127, F1: 127, F2: 2}
                        _ = fi22
                        _ = v23
                    }
                    if true {
                        v24 := GenS3{F0: 0, F1: 18446744073709551615, F2: 65535}
                        if true {
                            var v25 uint8 = 178
                            _ = v25
                        } else {
                            fmt.Println(v3)
                            for fi26 := 0; fi26 < 3; fi26++ {
                                v27 := GenS3{F0: 30, F1: 2, F2: 2}
                                _ = fi26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 2; fi28++ {
                                fmt.Println(v9)
                                var v29 int8 = 100
                                _ = fi28
                                _ = v29
                            }
                        }
                        v1 += (v1 << (((v1 + 985)) & 63))
                        _ = v24
                    } else {
                        var v30 bool = (2 < 2)
                        var v31 int64 = -9223372036854775808
                        _ = v30
                        _ = v31
                    }
                } else {
                    switch fi17 {
                    case 250:
                        _ = len(v2)
                    case 848:
                        var v32 int64 = -920
                        _ = v32
                    default:
                        var v33 bool = true
                        _ = v33
                    }
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                    fmt.Println(v1)
                }
                var v34 float64 = (v3 * v3)
                _ = v18
                _ = v34
            } else {
                var v35 int16 = v5
                if false {
                    if false {
                        fmt.Println(v35, v9)
                        var v36 uint64 = 224
                        for fi37 := 0; fi37 < 4; fi37++ {
                            var v38 uint16 = (v9 - (^(v9)))
                            var v39 string = "hello"
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = v36
                    }
                    for fi40 := 0; fi40 < 3; fi40++ {
                        switch (^(fi17)) {
                        case 0:
                            v41 := GenS4{F0: v9}
                            var v42 string = ("a" + ("\n\t" + "\n\t"))
                            v1 -= v1
                            _ = v41
                            _ = v42
                        case -962:
                            var v43 int8 = -18
                            _ = v43
                        case 2147483647:
                            v44 := GenS1{F0: int8(fi40), F1: 2, F2: 388}
                            _ = v44
                        }
                        v45 := GenS3{F0: 9223372036854775807, F1: 1, F2: 76}
                        _ = fi40
                        _ = v45
                    }
                    var v46 uint16 = 42
                    _ = v46
                }
                fmt.Println(v5, fi17)
                _ = v35
            }
            v2 = append(v2, 126)
            v5 = (^(v5))
            _ = fi17
        }
    }
    switch 0 {
    case 9223372036854775807:
        v47 := GenS5{F0: v5}
        _ = v47
    default:
        for fi48 := 0; fi48 < 6; fi48++ {
            v1 = (v1 >> (((v1 & (^(v1)))) & 63))
            _ = fi48
        }
    }
    v49 := GenS1{F0: 67, F1: 1, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
