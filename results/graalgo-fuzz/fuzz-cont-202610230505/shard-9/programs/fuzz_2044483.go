package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint32) uint {
    if true {
        fmt.Println(p1)
        switch 321 {
        case 128:
            p0 -= 1
        }
    } else {
        p1 = p1
    }
    fmt.Println(p1)
    for fi7 := 0; fi7 < 10; fi7++ {
        p1 += p0
        _ = fi7
    }
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int16, p1 uint32) int16 {
    v8 := GenS1{F0: (-1 < 100), F1: 183, F2: -522}
    if true {
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v8
    return (p0 + p0)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (!(65535 <= 2147483648))
    var v3 float64 = (-2.25 + 1e10)
    v4 := []uint32{}
    v5 := make(map[uint32]uint32)
    v6 := []int{32767, 2147483647}
    v2 = v2
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint32 = 42
        v3 -= v3
        _ = fi9
        _ = v10
    }
    switch "" {
    case "hello":
        _ = len(v4)
    }
    var v11 uint8 = 255
    var v12 uint64 = 200
    if v2 {
        if (!true) {
            fmt.Println(v12)
            var v13 uint32 = 0
            v14 := GenS1{F0: (-1 > 575), F1: 127, F2: 1}
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 float64 = v3
        if (v3 < 0.1) {
            var v17 string = "hello"
            if v2 {
                for fi18 := 0; fi18 < 3; fi18++ {
                    if v2 {
                        switch v17 {
                        case "foo bar":
                            v17 = ((v17 + (v17 + v17)) + v17)
                        }
                        if (-32768 >= -733) {
                            v17 = "foo bar"
                        }
                    } else {
                        switch v17 {
                        case "a":
                            var v19 int64 = -547
                            _ = v19
                        case "":
                            if v2 {
                                v17 = (v17 + (v17 + ("\n\t" + (v17 + (("" + v17) + "")))))
                                v20 := GenS1{F0: false, F1: uint8(v12), F2: int32(fi18)}
                                _ = v20
                            } else {
                                if v2 {
                                    v11 -= 128
                                    var v21 uint = 8
                                    v22 := GenS1{F0: v2, F1: (v11 * v11), F2: 2147483647}
                                    delete(v1, 444)
                                    var v23 uint32 = 2
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            }
                            if false {
                                var v24 int32 = -472
                                _ = v24
                            } else {
                                v25 := GenS2{F0: 0}
                                v17 = v17
                                _ = v25
                            }
                            v2 = true
                        }
                        delete(v5, 0)
                        v2 = v2
                        fmt.Println(v17)
                        fmt.Println(v3)
                    }
                    _ = fi18
                }
            }
            var v26 string = v17
            _ = v17
            _ = v26
        } else {
            v12 = v12
            for fi27 := 0; fi27 < 10; fi27++ {
                v28 := GenS2{F0: 355}
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            fmt.Println(v3)
            _ = fi29
        }
        _ = fi15
        _ = v16
    }
    v1[2147483647] = "a"
    v12 -= v12
    switch -60 {
    case 42:
        _ = len(v5)
    case 0:
        fmt.Println(v2, v3, v11)
    default:
        v2 = v2
    }
    fmt.Println(v12, v11, v11)
    var v30 int16 = genFn1(1, 4294967295)
    v31 := GenS1{F0: false, F1: 127, F2: -778}
    v32 := v1[0]
    fmt.Println(v30, v11, v12)
    fmt.Println(v30, v3, v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
