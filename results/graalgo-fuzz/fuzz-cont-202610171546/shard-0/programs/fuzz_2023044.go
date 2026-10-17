package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 int8; F2 string }
type GenS4 struct { F0 uint32; F1 int16; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint64 = 320
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    var v5 uint32 = 953
    if true {
        switch 127 {
        case -380:
            if (!false) {
                var v6 float32 = -2.25
                v7 := GenS1{F0: v5, F1: 2}
                _ = v6
                _ = v7
            }
        case 682:
            v1 = (v1 >> ((4294967295) & 63))
        case 1:
            if false {
                var v8 uint = 4294967295
                switch v2 {
                case "a":
                    v9 := GenS1{F0: (v5 >> ((65535) & 63)), F1: 32767}
                    _ = v9
                case "\n\t":
                    switch 76 {
                    case 1:
                        switch 9223372036854775807 {
                        case 4294967295:
                            for fi10 := 0; fi10 < 15; fi10++ {
                                var v11 int64 = v4
                                for fi12 := 0; fi12 < 20; fi12++ {
                                    v11 = v4
                                    if ((fi10 > fi12) || true) {
                                        v13 := []int{}
                                        v14 := v3[(v5 - (v5 + v5))]
                                        v1 = (v1 * (v1 & (v1 & (v1 % (((v1 * (v1 % ((v1) | 1)))) | 1)))))
                                        v15 := GenS2{F0: 0, F1: (fi12 >> (((v8 + v8)) & 63))}
                                        var v16 int8 = 1
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    } else {
                                        var v17 uint8 = 1
                                        _ = v17
                                    }
                                    _ = fi12
                                }
                                if false {
                                    v18 := GenS3{F0: 1, F1: int8(v11), F2: "a"}
                                    v8 += (^(v8))
                                    _ = v18
                                } else {
                                    v11 = (-(v11))
                                }
                                _ = fi10
                                _ = v11
                            }
                        case 2147483647:
                            var v19 int8 = 73
                            _ = v19
                        default:
                            fmt.Println(v4)
                        }
                    case 873:
                        fmt.Println(v1, v1, v8)
                    }
                case "foo bar":
                    _ = len(v3)
                default:
                    v20 := GenS3{F0: 42, F1: -128, F2: ("" + v2)}
                    _ = v20
                }
                _ = v8
            }
        default:
            v4 = v4
        }
    }
    v21 := []int{0, 844, 440, 0}
    _ = len(v21)
    for fi22 := 0; fi22 < 16; fi22++ {
        if true {
            var v23 int8 = -1
            for fi24 := 0; fi24 < 9; fi24++ {
                v2 = "\n\t"
                var v25 int16 = 0
                _ = fi24
                _ = v25
            }
            if false {
                if len(v21) > 0 {
                    v21[0] = (fi22 >> ((42) & 63))
                }
                v4 = -1
            } else {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v5 = v5
                    v2 = (v2 + "\n\t")
                    var v27 float64 = 3.14159
                    _ = fi26
                    _ = v27
                }
                v1 = v1
            }
            _ = v23
        } else {
            var v28 uint8 = 128
            _ = v28
        }
        v1 -= v1
        v29 := GenS4{F0: v5, F1: 32767, F2: true}
        _ = fi22
        _ = v29
    }
    var v30 int32 = 0
    if true {
        fmt.Println(v1, v1)
        v21 = append(v21, 9223372036854775807)
    }
    v31 := GenS5{F0: 42}
    if true {
        v2 = v2
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS6{F0: int8(fi32)}
            _ = fi32
            _ = v33
        }
        var v34 int = 1
        _ = v34
    } else {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        v2 = v2
    }
    v35 := GenS2{F0: 378, F1: -682}
    var v36 float32 = 3.14159
    var v37 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
