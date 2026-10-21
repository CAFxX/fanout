package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 int) int16 {
    p0 += 42
    var v6 int32 = 2147483647
    v7 := GenS1{F0: 32767, F1: 533, F2: 42}
    p2 += p2
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 int8, p1 int16) uint64 {
    p1 += genFn0(0, 619, 9223372036854775807)
    var v9 int64 = 0
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int32 = -2147483648
            fmt.Println(fi10, p1, p0)
            _ = fi11
            _ = v12
        }
        p1 = 0
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 198
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 329
    v3 := make(map[int]int)
    var v4 int = int(v2)
    v5 := []int{(v4 << ((670) & 63)), v4, v4, v4}
    if (true && false) {
        switch 32767 {
        case -1:
            v13 := GenS1{F0: 65535, F1: 175, F2: 0}
            _ = v13
        case -805:
            v14 := GenS1{F0: 0, F1: 255, F2: 2}
            _ = v14
        case 4294967295:
            _ = len(v3)
        }
    }
    v15 := GenS1{F0: 0, F1: 0, F2: 502}
    fmt.Println(v2, v2, v1)
    _ = len(v5)
    var v16 int8 = 23
    var v17 int8 = (v16 << ((662) & 63))
    if (42 >= 204) {
        v1 -= v1
    } else {
        var v18 uint32 = 1
        for fi19 := 0; fi19 < 9; fi19++ {
            switch (fi19 - 1) {
            case 149:
                _ = len(v3)
            default:
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v16 -= 1
                } else {
                    var v20 int32 = 0
                    _ = v20
                }
            }
            _ = fi19
        }
        _ = v18
    }
    switch v4 {
    case -208:
        v1 -= (v1 | (v1 + (v1 ^ v1)))
    }
    v21 := GenS2{F0: (v1 | 2), F1: v2}
    if false {
        var v22 float64 = (0.0 / 3.14159)
        _ = v22
    } else {
        var v23 uint16 = 377
        var v24 uint64 = 238
        if (!false) {
            v3[int(v2)] = (v4 & (v4 % (((v4 % ((v4) | 1))) | 1)))
            switch (v4 + v4) {
            case 127:
                for fi25 := 0; fi25 < 6; fi25++ {
                    fmt.Println(v17, v23, v1)
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v16, v23)
                        _ = len(v5)
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 uint16 = (v23 / ((v23) | 1))
                        v1 = (v1 << ((42) & 63))
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            case 32768:
                switch "a" {
                case "foo bar":
                    if (0 < 757) {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v3[42] = 32767
                            _ = fi28
                        }
                        var v29 int32 = 912
                        _ = v29
                    } else {
                        var v30 string = "\n\t"
                        var v31 int32 = 611
                        var v32 uint8 = 255
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                case "":
                    v17 = 1
                }
            }
            var v33 int8 = v16
            _ = v33
        }
        _ = v23
        _ = v24
    }
    switch "" {
    case "foo bar":
        v3[v4] = (v4 << ((65535) & 63))
    case "\n\t":
        for fi34 := 0; fi34 < 1; fi34++ {
            v2 -= 0
            delete(v3, fi34)
            _ = fi34
        }
    case "":
        var v35 uint64 = 18446744073709551615
        _ = v35
    default:
        v36 := GenS1{F0: 42, F1: 841, F2: 255}
        _ = v36
    }
    var v37 uint = 571
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
