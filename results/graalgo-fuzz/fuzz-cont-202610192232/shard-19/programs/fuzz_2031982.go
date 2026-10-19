package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool, p1 bool, p2 uint32) int8 {
    p1 = p0
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 1}
        fmt.Println(p2)
        p2 -= p2
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint8, p1 uint8, p2 uint32) uint8 {
    var v7 string = "hello"
    var v8 int = 2147483647
    var v9 bool = (!false)
    var v10 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p1 >> ((842) & 63))
}

func main() {
    v1 := []int{4294967295, 1, -2147483648, 2147483647}
    v2 := make(map[int]bool)
    var v3 int16 = 1
    var v4 bool = false
    v3 += 769
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: (v3 & (v3 | int16(fi11)))}
        v3 = 42
        switch "\n\t" {
        case "":
            if (v4 && (!v4)) {
                _ = len(v1)
                if (1 <= -108) {
                    var v13 uint64 = 2
                    switch -397 {
                    case -783:
                        switch "" {
                        case "a":
                            v3 = 32767
                        case "hello":
                            var v14 bool = (v13 > v13)
                            _ = v14
                        default:
                            var v15 int64 = int64(v3)
                            _ = v15
                        }
                    case 127:
                        _ = len(v2)
                    case -7:
                        if v4 {
                            delete(v2, (fi11 - 405))
                            if v4 {
                                var v16 int64 = -1
                                _ = v16
                            } else {
                                v17 := GenS1{F0: v3}
                                _ = v17
                            }
                        }
                    default:
                        v3 = (v3 % (((v3 + v3)) | 1))
                    }
                    _ = v13
                } else {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v19 := v2[(fi18 + (fi11 / ((fi11) | 1)))]
                        var v20 float64 = 0.1
                        v20 -= v20
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                var v21 uint8 = 127
                _ = v21
            }
        }
        _ = fi11
        _ = v12
    }
    v4 = ("" == ("\n\t" + ""))
    if true {
        var v22 int32 = -529
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
    } else {
        v23 := GenS1{F0: 32767}
        v24 := GenS1{F0: 42}
        _ = v23
        _ = v24
    }
    var v25 int64 = -544
    fmt.Println(v25, v25, v3)
    var v26 string = "foo bar"
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS1{F0: int16(fi27)}
        v29 := GenS1{F0: (v3 | (v3 << ((785) & 63)))}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int16 = 1000
    fmt.Println(v26, v3)
    v25 -= (v25 + (-(v25)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
