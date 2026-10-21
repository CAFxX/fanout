package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint16) int32 {
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            p0 = p0
            if false {
                for fi8 := 0; fi8 < 6; fi8++ {
                    if false {
                        if (42 >= 18446744073709551615) {
                            switch (fi8 * fi7) {
                            case 1:
                                v9 := []int{}
                                _ = v9
                            case 301:
                                var v10 float64 = (0.1 - (float64(3.14159) + 0.1))
                                _ = v10
                            default:
                                var v11 string = ("\n\t" + "hello")
                                _ = v11
                            }
                        }
                        p1 -= p1
                        var v12 uint16 = (p1 ^ (^(p1)))
                        _ = v12
                    } else {
                        var v13 int16 = 0
                        p1 += 255
                        var v14 uint16 = p1
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{(fi8 / (((fi8 * fi7)) | 1))}
                    _ = fi8
                    _ = v15
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 bool = (127 > 127)
                    _ = fi16
                    _ = v17
                }
            }
            if (32768 != fi7) {
                var v18 string = ""
                _ = v18
            }
            _ = fi7
        }
        v19 := GenS1{F0: 127, F1: 173}
        var v20 int = 127
        _ = v19
        _ = v20
    }
    var v21 int8 = 1
    _ = p0
    _ = p1
    _ = v21
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := []int{-443}
    var v3 uint16 = 542
    var v4 bool = true
    var v5 float64 = (0.0 / 1.5)
    var v6 bool = (!(false && (v5 <= float64((-2.25 - (0.0 - 3.14159))))))
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS1{F0: v1, F1: (v1 - v1)}
        _ = fi22
        _ = v23
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v6)
    v24 := GenS1{F0: (v1 >> ((2) & 63)), F1: v1}
    switch 51 {
    case 42:
        fmt.Println(v4)
    case -2147483648:
        for fi25 := 0; fi25 < 5; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 uint = 4294967295
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(v4, v4, v6)
                _ = fi28
            }
            _ = fi25
        }
    default:
        var v29 int = -285
        _ = v29
    }
    var v30 uint16 = 32767
    fmt.Println(v6)
    fmt.Println(v4)
    if (-1 > -206) {
        if false {
            for fi31 := 0; fi31 < 20; fi31++ {
                for fi32 := 0; fi32 < 13; fi32++ {
                    v30 = (v3 / ((2) | 1))
                    _ = fi32
                }
                v3 -= (^(v3))
                if len(v2) > 0 {
                    v2[0] = (fi31 + (fi31 / ((fi31) | 1)))
                }
                _ = fi31
            }
            fmt.Println(v1)
        } else {
            var v33 uint = 18446744073709551615
            _ = v33
        }
        v30 -= (v30 >> ((0) & 63))
    } else {
        var v34 uint16 = v30
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        if len(v2) > 0 {
            v2[0] = int(v3)
        }
        v36 := GenS2{F0: -64}
        _ = fi35
        _ = v36
    }
    v37 := GenS2{F0: 50}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
