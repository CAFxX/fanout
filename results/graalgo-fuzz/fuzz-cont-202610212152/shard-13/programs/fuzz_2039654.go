package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0() int {
    v5 := GenS1{F0: 0, F1: -1, F2: ""}
    var v6 string = "foo bar"
    _ = v5
    _ = v6
    return 759
}

func main() {
    var v1 int16 = 414
    var v2 uint8 = 42
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    if false {
        switch "foo bar" {
        case "\n\t":
            v1 += -579
        }
        var v7 uint16 = 268
        _ = v7
    }
    var v8 uint32 = 0
    switch 32768 {
    case 32768:
        var v9 uint8 = 23
        _ = v9
    case 4294967295:
        v10 := GenS1{F0: -202, F1: -9223372036854775808, F2: "a"}
        _ = v10
    case 42:
        v8 = (v8 | v8)
    default:
        if false {
            fmt.Println(v2, v2, v2)
            v8 = 1
            if false {
                if false {
                    v11 := GenS2{F0: ("hello" + "foo bar"), F1: 42, F2: 352}
                    _ = v11
                } else {
                    if true {
                        v8 = (v8 << ((536) & 63))
                        var v12 int32 = 2147483647
                        _ = v12
                    }
                    fmt.Println(v8)
                }
                switch ("foo bar" + "a") {
                case "":
                    fmt.Println(v8)
                case "foo bar":
                    delete(v4, (^(v8)))
                default:
                    v1 += (v1 - -1)
                }
                v13 := GenS2{F0: ((((("a" + "\n\t") + "foo bar") + "foo bar") + "a") + "a"), F1: 1, F2: 532}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 19; fi14++ {
                    if ((v2 / ((255) | 1)) != (v2 + 134)) {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            var v16 int8 = 127
                            _ = fi15
                            _ = v16
                        }
                        fmt.Println(v2, v1, v1)
                    }
                    var v17 uint32 = 2147483647
                    _ = fi14
                    _ = v17
                }
                switch -248 {
                case 689:
                    var v18 float32 = (0.0 + 0.0)
                    _ = v18
                case 1:
                    switch 32767 {
                    case 42:
                        v8 -= (^(v8))
                    }
                }
                var v19 int32 = -1
                _ = v19
            }
        } else {
            v20 := GenS3{F0: 9223372036854775807}
            _ = v20
        }
    }
    var v21 int32 = 0
    v1 = v1
    var v22 uint16 = 32767
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int = genFn0()
        var v25 uint32 = v8
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint = 778
    if (3.14159 > 1.5) {
        v27 := GenS1{F0: int(v21), F1: 128, F2: (("" + "a") + ("foo bar" + ("a" + ("\n\t" + "\n\t"))))}
        switch -76 {
        case 226:
            fmt.Println(v26)
        case 128:
            _ = len(v4)
        case -1:
            for fi28 := 0; fi28 < 8; fi28++ {
                fmt.Println(v26)
                if false {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v3 = append(v3, fi29)
                        _ = fi29
                    }
                }
                _ = fi28
            }
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 uint32 = uint32(v2)
            _ = fi30
            _ = v31
        }
        _ = v27
    }
    var v32 uint64 = 1
    var v33 uint = uint(v2)
    var v34 float64 = float64(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v8
}
