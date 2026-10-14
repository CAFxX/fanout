package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 string }

func main() {
    v1 := []uint32{804, 485, 97}
    var v2 uint32 = 4294967295
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    var v5 float32 = 3.14159
    var v6 int32 = 1
    v6 = v6
    v6 = (v6 | v6)
    v6 += 372
    switch "foo bar" {
    case "hello":
        v6 -= -1
    }
    v6 = (v6 % (((v6 ^ (v6 + v6))) | 1))
    for fi7 := 0; fi7 < 1; fi7++ {
        if (942 == 242) {
            if len(v1) > 0 {
                v1[0] = v2
            }
            var v8 int32 = -853
            v8 += (v8 | v6)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: "foo bar", F1: 255, F2: 42}
                switch fi7 {
                case 9223372036854775807:
                    var v11 uint64 = 1
                    _ = v11
                case -1:
                    switch fi9 {
                    case 127:
                        v6 += (v6 ^ (v6 >> ((2) & 63)))
                    case 554:
                        v1 = append(v1, 65535)
                    default:
                        v12 := GenS2{F0: -9223372036854775808}
                        _ = v12
                    }
                case 128:
                    if false {
                        v2 = v2
                        switch "hello" {
                        case "hello":
                            if (2 == 792) {
                                if (fi9 < (fi7 << ((18446744073709551615) & 63))) {
                                    fmt.Println(v6)
                                } else {
                                    if (128 < 1) {
                                        delete(v4, "foo bar")
                                    } else {
                                        var v13 int64 = -1
                                        _ = v13
                                    }
                                }
                            }
                        }
                    } else {
                        v14 := GenS3{F0: 255, F1: 332, F2: 32767}
                        _ = v14
                    }
                    v6 = v6
                    delete(v4, ("a" + "\n\t"))
                }
                _ = fi9
                _ = v10
            }
            switch ("" + "a") {
            case "foo bar":
                fmt.Println(v2)
            }
            var v15 uint16 = 42
            _ = v15
        }
        _ = fi7
    }
    v16 := GenS4{F0: 42, F1: "hello"}
    switch 128 {
    case 2147483647:
        switch -882 {
        case -1:
            var v17 uint16 = 250
            _ = v17
        case 127:
            v6 -= 0
        default:
            var v18 int16 = -81
            _ = v18
        }
    }
    var v19 uint64 = 1
    var v20 float64 = (-2.25 + 1.5)
    if false {
        _ = len(v1)
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
        v5 = v5
    } else {
        fmt.Println(v5, v5, v19)
    }
    var v23 int16 = 32767
    var v24 int16 = 0
    _ = len(v3)
    _ = len(v3)
    var v25 int16 = (v24 & -881)
    switch 2147483647 {
    case 1:
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS4{F0: -1, F1: "foo bar"}
            for fi28 := 0; fi28 < 18; fi28++ {
                var v29 uint8 = 0
                fmt.Println(fi26, v5)
                v30 := GenS5{F0: ""}
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi26
            _ = v27
        }
    case -2147483648:
        var v31 int16 = 1
        _ = v31
    default:
        v32 := GenS4{F0: 9223372036854775807, F1: ("\n\t" + "foo bar")}
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
