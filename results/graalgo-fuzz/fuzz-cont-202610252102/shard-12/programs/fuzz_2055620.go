package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int16 = 42
    var v2 uint8 = 255
    var v3 uint16 = 32767
    v4 := []int{-487, 42, 0}
    v5 := GenS1{F0: 42}
    v3 = (v3 >> ((4294967295) & 63))
    fmt.Println(v3)
    switch -186 {
    case -1:
        v6 := GenS1{F0: 18446744073709551615}
        _ = v6
    case 32767:
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(v1)
            if (true || true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch ("hello" + "a") {
                case "hello":
                    for fi8 := 0; fi8 < 8; fi8++ {
                        fmt.Println(fi7, v2)
                        _ = fi8
                    }
                case "a":
                    v2 = (v2 * (v2 + (v2 % ((1) | 1))))
                case "\n\t":
                    for fi9 := 0; fi9 < 15; fi9++ {
                        v10 := GenS1{F0: 4294967295}
                        v1 += (v1 & (v1 * (v1 / ((v1) | 1))))
                        var v11 int64 = int64(v1)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                default:
                    for fi12 := 0; fi12 < 10; fi12++ {
                        switch fi7 {
                        case 196:
                            v3 += v3
                        default:
                            if false {
                                v13 := GenS1{F0: 18446744073709551615}
                                _ = v13
                            } else {
                                var v14 int64 = int64(fi7)
                                var v15 uint = 0
                                if (18446744073709551615 > 42) {
                                    var v16 int32 = -2147483648
                                    v17 := GenS1{F0: v15}
                                    v1 = 1
                                    _ = v16
                                    _ = v17
                                }
                                _ = v14
                                _ = v15
                            }
                        }
                        _ = fi12
                    }
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = int(v3)
                }
            }
            _ = fi7
        }
    }
    v4 = append(v4, int(v3))
    if (255 < 415) {
        v3 = 0
    }
    v18 := GenS2{F0: v1}
    var v19 uint64 = 562
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
