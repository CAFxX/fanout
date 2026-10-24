package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 int16, p2 string) bool {
    v4 := []int{840, -460, 1, -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    p1 = p0
    if (234 < 2) {
        var v5 uint64 = 2
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 int32 = 42
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        _ = v5
    } else {
        _ = len(v4)
        p1 += 1
        var v9 uint8 = 0
        _ = v9
    }
    var v10 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    return false
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"", ("hello" + "hello"), "a"}
    var v3 string = ""
    fmt.Println(v3)
    delete(v1, 1)
    v1[249] = 9223372036854775807
    for fi11 := 0; fi11 < 20; fi11++ {
        v1[65535] = (-(fi11))
        _ = fi11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v3 = "hello"
        _ = fi12
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if (-726 == 0) {
            if true {
                var v13 uint64 = 42
                _ = v13
            } else {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        if (-1 != 1000) {
                            if false {
                                v14 := v1[65535]
                                _ = v14
                            } else {
                                if (!(false && false)) {
                                    var v15 int32 = 0
                                    fmt.Println(v15, v3)
                                    _ = v15
                                }
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    var v17 uint64 = 42
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v18 := GenS1{F0: 32767, F1: 2147483648, F2: 0}
                            _ = v18
                        }
                    }
                }
            }
        }
    } else {
        var v19 uint = 999
        fmt.Println(v3, v19)
        _ = v19
    }
    var v20 int32 = -68
    fmt.Println(v20, v3)
    var v21 float64 = 0.1
    v22 := GenS2{F0: 0}
    _ = len(v2)
    v23 := v1[580]
    switch v23 {
    case 32767:
        fmt.Println(v3, v20)
    case -591:
        v3 = v3
    default:
        if ((false || false) || false) {
            v20 -= v20
            var v24 uint16 = 32767
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
