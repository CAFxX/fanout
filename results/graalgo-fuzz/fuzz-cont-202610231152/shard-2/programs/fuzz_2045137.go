package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    v1 := []int{932, 32768, 32768, -2147483648}
    var v2 int = 127
    var v3 int8 = -47
    var v4 int8 = v3
    v5 := make(map[string]int)
    v6 := []uint32{0, 4294967295, 65535, 793}
    switch (("" + "\n\t") + ("\n\t" + ("a" + ("" + ("\n\t" + ("" + "\n\t")))))) {
    case "foo bar":
        if false {
            v5[("\n\t" + "\n\t")] = v2
        } else {
            delete(v5, "hello")
        }
    }
    v2 = v2
    var v7 int16 = -489
    _ = len(v1)
    switch 32767 {
    case 2147483647:
        var v8 uint32 = 280
        _ = v8
    }
    var v9 uint32 = 653
    var v10 int16 = v7
    switch "foo bar" {
    case "":
        if len(v6) > 0 {
            v6[0] = v9
        }
    case "hello":
        var v11 uint32 = (^(v9))
        _ = v11
    case "\n\t":
        switch "hello" {
        case "\n\t":
            v10 += -539
        case "a":
            if false {
                var v12 uint32 = (v9 << ((968) & 63))
                if ((^(v2)) >= (v2 - (-(v2)))) {
                    var v13 float32 = (1e10 + 1.5)
                    v4 = 127
                    var v14 uint32 = (v9 & 2147483648)
                    _ = v13
                    _ = v14
                }
                _ = v12
            }
        case "foo bar":
            _ = len(v1)
        default:
            switch "\n\t" {
            case "\n\t":
                for fi15 := 0; fi15 < 19; fi15++ {
                    v10 = int16(v2)
                    _ = fi15
                }
            case "a":
                for fi16 := 0; fi16 < 13; fi16++ {
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 uint = 477
                        v19 := GenS1{F0: 42}
                        v20 := GenS1{F0: 0}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        var v21 uint64 = 42
                        fmt.Println(v9, v10, v2)
                        _ = v21
                    }
                    _ = fi16
                }
            case "foo bar":
                switch (^(v2)) {
                case 9223372036854775807:
                    var v22 uint8 = 121
                    _ = v22
                case 128:
                    var v23 float32 = (-2.25 * 1.5)
                    _ = v23
                default:
                    var v24 bool = (v2 != v2)
                    _ = v24
                }
            default:
                v25 := GenS2{F0: ("" + ("\n\t" + "\n\t")), F1: (v9 - v9)}
                _ = v25
            }
        }
    default:
        v26 := GenS2{F0: "\n\t", F1: (^(v9))}
        _ = v26
    }
    v4 = -63
    var v27 uint64 = 460
    if (true && false) {
        fmt.Println(v7, v4)
        if true {
            switch "hello" {
            case "\n\t":
                if false {
                    var v28 bool = true
                    if false {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v28 = v28
                            var v30 uint16 = 255
                            _ = fi29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 5; fi31++ {
                            if v28 {
                                var v32 int64 = -9223372036854775808
                                v33 := GenS1{F0: (v27 / (((v27 + (^(v27)))) | 1))}
                                var v34 int16 = 162
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            _ = fi31
                        }
                    } else {
                        var v35 uint16 = 255
                        _ = v35
                    }
                    _ = v28
                }
            }
            delete(v5, "a")
            v2 = -774
        }
    } else {
        if false {
            _ = len(v5)
            fmt.Println(v2, v3)
        }
    }
    v2 -= v2
    var v36 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
