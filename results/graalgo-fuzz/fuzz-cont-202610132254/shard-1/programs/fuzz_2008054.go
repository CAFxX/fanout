package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"hello"}
    var v3 int16 = 42
    var v4 int16 = (v3 | v3)
    fmt.Println(v3, v4)
    var v5 uint32 = 79
    v5 += (v5 >> ((49) & 63))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v2 = append(v2, "hello")
    fmt.Println(v3)
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v5, v5, v4)
        _ = fi6
    }
    v7 := GenS1{F0: 4294967295, F1: (^(v4)), F2: 539}
    v1[-624] = ""
    if (3.14159 >= 0.1) {
        v8 := GenS1{F0: 65535, F1: (v4 + (^(v4))), F2: 18446744073709551615}
        v3 = v3
        fmt.Println(v3, v3)
        _ = v8
    } else {
        _ = len(v2)
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 42, F1: v3, F2: 743}
            v2 = append(v2, "hello")
            for fi11 := 0; fi11 < 10; fi11++ {
                _ = len(v2)
                if (!(((v4 * (^(v3))) == (v3 / ((42) | 1))) && false)) {
                    var v12 uint32 = (v5 | (^(v5)))
                    if false {
                        v12 -= (^(v12))
                        switch "hello" {
                        case "hello":
                            var v13 float32 = 0.0
                            _ = v13
                        case "a":
                            v14 := GenS2{F0: 377, F1: v12, F2: 526}
                            _ = v14
                        case "\n\t":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        delete(v1, 0)
                    } else {
                        var v15 uint64 = 521
                        var v16 uint32 = (v5 / ((v12) | 1))
                        _ = v15
                        _ = v16
                    }
                    _ = v12
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    var v17 float64 = 0.1
    switch 32768 {
    case -925:
        switch -97 {
        case -189:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 1:
            v5 += (v5 >> ((23) & 63))
        case -754:
            for fi18 := 0; fi18 < 5; fi18++ {
                _ = len(v1)
                var v19 string = ""
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 int16 = 32767
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
        }
    case 32768:
        _ = len(v2)
    case -2147483648:
        fmt.Println(v4, v3)
    }
    fmt.Println(v17, v5)
    if false {
        v4 -= (v4 >> ((255) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
