package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int32 = 187
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    var v4 float64 = 3.14159
    var v5 uint64 = 2
    var v6 float32 = 0.1
    fmt.Println(v4, v6)
    v7 := []int{}
    v8 := GenS1{F0: false}
    v5 = (v5 | v5)
    v6 -= v6
    v9 := GenS1{F0: true}
    fmt.Println(v4, v4, v5)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v10 := GenS2{F0: 2}
    var v11 int8 = -128
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v5, v11, v4)
        v1 += v1
        _ = fi12
    }
    if false {
        if false {
            var v13 int8 = 100
            switch 32768 {
            case 4294967295:
                v6 -= (((-2.25 * (0.0 + (v6 + (1e10 * v6)))) * float32(0.0)) * v6)
            default:
                var v14 int8 = (-(v13))
                _ = v14
            }
            _ = v13
        } else {
            v15 := v2["a"]
            fmt.Println(v5, v15, v6)
            switch (v15 >> ((1) & 63)) {
            case -1:
                if (((4294967295 <= 0) && false) || (!true)) {
                    if true {
                        var v16 float32 = v6
                        if false {
                            var v17 uint16 = 913
                            _ = v17
                        } else {
                            switch ("foo bar" + "a") {
                            case "a":
                                fmt.Println(v1)
                            case "foo bar":
                                switch (-(v15)) {
                                case -822:
                                    v4 = (v4 + v4)
                                default:
                                    switch "" {
                                    case "\n\t":
                                        if (true && false) {
                                            v3[""] = (!false)
                                        } else {
                                            for fi18 := 0; fi18 < 9; fi18++ {
                                                if false {
                                                    v15 = 32767
                                                    var v19 uint8 = 255
                                                    var v20 int32 = (v1 * (-(v1)))
                                                    _ = v19
                                                    _ = v20
                                                } else {
                                                    fmt.Println(v5, fi18)
                                                }
                                                var v21 float32 = ((float32(v4) * (-2.25 - float32((v4 / v4)))) - 1e10)
                                                _ = fi18
                                                _ = v21
                                            }
                                            switch 127 {
                                            case 0:
                                                v22 := GenS1{F0: (false || false)}
                                                _ = v22
                                            }
                                            fmt.Println(v1, v1)
                                        }
                                    default:
                                        fmt.Println(v16, v5, v6)
                                    }
                                }
                            default:
                                fmt.Println(v6)
                            }
                        }
                        _ = v16
                    }
                    v15 -= v15
                    v23 := GenS3{F0: (-(v1)), F1: v5, F2: v15}
                    _ = v23
                }
            case 1:
                v6 -= (v6 + v6)
            case 42:
                delete(v3, "a")
            default:
                fmt.Println(v11)
            }
            _ = v15
        }
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        if false {
            if true {
                v25 := GenS3{F0: v1, F1: v5, F2: 1}
                v26 := GenS4{F0: 2147483648}
                _ = v25
                _ = v26
            } else {
                v5 += v5
            }
            var v27 int = 747
            _ = v27
        } else {
            v6 += 1.5
            var v28 uint8 = 128
            if (!(int64(v1) == 808)) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v29 uint64 = (v5 << ((563) & 63))
                var v30 int16 = -1
                _ = v29
                _ = v30
            }
            _ = v28
        }
        _ = fi24
    }
    if false {
        var v31 uint8 = 128
        _ = v31
    } else {
        v32 := GenS5{F0: 1}
        var v33 int8 = v11
        _ = v32
        _ = v33
    }
    if true {
        var v34 int8 = (v11 & int8(v5))
        _ = v34
    }
    _ = len(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
