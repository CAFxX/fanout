package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 603
    var v4 int16 = 42
    v5 := []int{32768, -2147483648, 118, 9223372036854775807}
    switch 2147483647 {
    case -57:
        v6 := GenS1{F0: 101}
        _ = v6
    case -2147483648:
        var v7 int16 = 0
        _ = v7
    default:
        v8 := v1[1]
        _ = v8
    }
    var v9 float32 = float32(-2.25)
    var v10 float64 = 0.1
    var v11 uint64 = 467
    var v12 uint = 65535
    var v13 uint32 = 4294967295
    if true {
        var v14 uint16 = v3
        v4 -= 456
        _ = v14
    } else {
        var v15 uint16 = 514
        fmt.Println(v10, v12)
        var v16 int32 = 897
        _ = v15
        _ = v16
    }
    v13 += v13
    var v17 float32 = (float32(float64((float32(v10) * (v9 + (0.1 / v9))))) - (float32(float64(3.14159)) + v9))
    v18 := GenS2{F0: (v4 | v4), F1: int64(v12)}
    delete(v1, -348)
    if true {
        switch 127 {
        case -1:
            for fi19 := 0; fi19 < 6; fi19++ {
                if true {
                    v11 += (v11 + (^(v11)))
                    if false {
                        if false {
                            var v20 string = "\n\t"
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 uint32 = (^(v13))
                                var v23 uint32 = v22
                                delete(v1, 32767)
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            _ = v20
                        }
                        var v24 int8 = 100
                        _ = v24
                    }
                }
                _ = fi19
            }
        case -2147483648:
            if false {
                if true {
                    fmt.Println(v3)
                }
                v10 -= v10
                for fi25 := 0; fi25 < 16; fi25++ {
                    if (!false) {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            v5 = append(v5, fi26)
                            switch 32767 {
                            case 874:
                                v27 := GenS2{F0: v4, F1: 42}
                                _ = v27
                            case 4294967295:
                                if true {
                                    v5 = append(v5, (-(fi26)))
                                }
                            }
                            _ = fi26
                        }
                    } else {
                        v4 -= (v4 >> ((1) & 63))
                        v13 = (v13 | v13)
                        v28 := v1[128]
                        _ = v28
                    }
                    var v29 float64 = v10
                    fmt.Println(v29)
                    _ = fi25
                    _ = v29
                }
            }
        }
        fmt.Println(v13, v3, v9)
        v9 += 3.14159
    }
    switch "a" {
    case "foo bar":
        v30 := GenS1{F0: 100}
        _ = v30
    case "":
        if (2 < 58) {
            fmt.Println(v4)
        }
    case "a":
        var v31 uint16 = 1
        _ = v31
    default:
        var v32 int64 = -1
        _ = v32
    }
    fmt.Println(v12)
    v33 := v1[-2147483648]
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 18; fi35++ {
            v2[(v13 << ((v12) & 63))] = 4294967295
            _ = fi35
        }
        fmt.Println(fi34, v10)
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
