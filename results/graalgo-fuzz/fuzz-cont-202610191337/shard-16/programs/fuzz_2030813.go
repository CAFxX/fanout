package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int16, p2 int16) uint32 {
    switch "hello" {
    case "a":
        var v6 uint8 = 1
        _ = v6
    default:
        p1 = (p2 / (((p1 / (((-(p1))) | 1))) | 1))
    }
    v7 := GenS1{F0: 116}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483648
}
func genFn1() uint16 {
    if (1 >= 42) {
        switch 338 {
        case -216:
            if true {
                var v8 uint = 2
                v9 := []int{42, 128}
                fmt.Println(v8)
                _ = v8
                _ = v9
            }
        default:
            var v10 uint16 = 1
            _ = v10
        }
        var v11 int = 113
        _ = v11
    } else {
        switch 481 {
        case 624:
            v12 := []int{2147483647, -9223372036854775808, -2147483648, -1}
            _ = v12
        }
        v13 := GenS2{F0: 100}
        _ = v13
    }
    var v14 int32 = -1
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    var v3 string = "a"
    v4 := make(map[uint32]bool)
    var v5 float32 = 0.0
    var v15 float32 = float32((float64(((v5 - v5) - 1e10)) * 3.14159))
    fmt.Println(v15, v15)
    var v16 string = v3
    v16 = ("" + (("" + "a") + "hello"))
    fmt.Println(v5)
    var v17 uint64 = 919
    switch 0 {
    case 427:
        switch 127 {
        case 672:
            switch (v16 + ((v16 + ((v16 + "hello") + v16)) + v16)) {
            case "":
                v18 := []int{-1}
                _ = v18
            case "hello":
                if (1 > 2) {
                    if true {
                        var v19 float32 = 1.5
                        _ = v19
                    }
                } else {
                    v20 := GenS1{F0: 218}
                    switch v16 {
                    case "a":
                        v2 = v2
                    case "":
                        if (!true) {
                            v2 += v2
                            fmt.Println(v17, v15, v16)
                            for fi21 := 0; fi21 < 9; fi21++ {
                                switch int(v17) {
                                case 128:
                                    fmt.Println(v16, v5)
                                default:
                                    fmt.Println(v16)
                                }
                                _ = fi21
                            }
                        } else {
                            for fi22 := 0; fi22 < 5; fi22++ {
                                var v23 uint32 = uint32(v17)
                                _ = fi22
                                _ = v23
                            }
                            var v24 string = v16
                            _ = v24
                        }
                    default:
                        var v25 int64 = int64(v17)
                        _ = v25
                    }
                    _ = v20
                }
            }
        case -1:
            delete(v4, 42)
        case 1:
            if true {
                fmt.Println(v5, v3)
                if false {
                    v2 += genFn1()
                    var v26 int16 = -32768
                    _ = v26
                } else {
                    v27 := []int{572, 868, 2147483647}
                    _ = v27
                }
                fmt.Println(v3, v15, v5)
            } else {
                _ = len(v1)
            }
        default:
            var v28 int8 = 127
            _ = v28
        }
    default:
        v2 = 42
    }
    if ((v3 + v3) <= (v16 + "a")) {
        var v29 int8 = 0
        v30 := []int{-1, 32767, -9223372036854775808, 4294967295}
        var v31 float64 = (1.5 * (((0.0 + (0.1 * 1.5)) * float64(v15)) * 3.14159))
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println(v17)
    var v32 uint32 = 2147483647
    var v33 float64 = (0.1 * 0.1)
    if (uint64(v2) >= (v17 * v17)) {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch (fi34 % ((fi34) | 1)) {
            case -2147483648:
                v16 = (("foo bar" + v16) + v3)
            default:
                fmt.Println(v33, v32)
            }
            fmt.Println(v5, v32)
            _ = fi34
        }
        var v35 uint8 = 127
        var v36 int = 0
        _ = v35
        _ = v36
    }
    if false {
        v5 = v5
    } else {
        fmt.Println(v33)
        var v37 int = 1
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
