package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 bool) uint {
    if false {
        v6 := GenS1{F0: 127, F1: 1, F2: p0}
        _ = v6
    }
    if p0 {
        v7 := []int{127, 1, 128}
        v8 := GenS1{F0: 1, F1: 0, F2: (-625 != 1)}
        _ = v7
        _ = v8
    }
    _ = p0
    return 1
}
func genFn1(p0 uint16) uint16 {
    fmt.Println(p0)
    v9 := []int{128, -963}
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            v11 := GenS1{F0: 100, F1: 32767, F2: (32767 <= 32767)}
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint16 = (p0 >> ((genFn0(true)) & 63))
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int64 = 42
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 692
        _ = fi10
        _ = v16
    }
    var v17 uint8 = 127
    _ = p0
    _ = v17
    _ = v9
    return (p0 >> ((2) & 63))
}

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    var v3 uint32 = 291
    var v4 string = "a"
    var v5 float32 = 0.1
    v18 := GenS2{F0: 100, F1: 255}
    switch "hello" {
    case "foo bar":
        for fi19 := 0; fi19 < 19; fi19++ {
            v2[v4] = v3
            _ = fi19
        }
    default:
        delete(v2, v4)
    }
    v20 := GenS3{F0: 42, F1: 42}
    var v21 float32 = (0.1 / float32(3.14159))
    var v22 int8 = 11
    var v23 string = v4
    v24 := v2[v23]
    fmt.Println(v21, v24)
    for fi25 := 0; fi25 < 18; fi25++ {
        if (902 <= (fi25 * -1)) {
            if (1 <= -1) {
                v26 := GenS2{F0: int8(v3), F1: 255}
                _ = v26
            } else {
                v24 = (v24 / ((v24) | 1))
                switch v4 {
                case "\n\t":
                    for fi27 := 0; fi27 < 13; fi27++ {
                        if (2 == 128) {
                            if (32767 < (fi25 | fi25)) {
                                if true {
                                    v2[v23] = v3
                                    v5 = float32(-2.25)
                                }
                                switch 9223372036854775807 {
                                case -191:
                                    v28 := GenS3{F0: 333, F1: 1}
                                    v29 := GenS1{F0: v22, F1: -32768, F2: ((0.1 / 1.0) == float64(v21))}
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int64 = 305
                                v5 += 0.1
                                if true {
                                    switch v23 {
                                    case "a":
                                        var v31 int8 = v22
                                        _ = v31
                                    default:
                                        var v32 int64 = v30
                                        var v33 int8 = (v22 % ((v22) | 1))
                                        v24 -= 42
                                        _ = v32
                                        _ = v33
                                    }
                                    if true {
                                        if false {
                                            fmt.Println(v30, v21, v21)
                                        }
                                        fmt.Println(fi25)
                                    } else {
                                        fmt.Println(fi27, v5)
                                        if true {
                                            var v34 bool = (!true)
                                            _ = v34
                                        } else {
                                            fmt.Println(v23)
                                            if (v23 != v23) {
                                                v24 = v24
                                                v21 = 3.14159
                                            }
                                        }
                                    }
                                    if false {
                                        var v35 int64 = (v30 % (((v30 - (v30 >> ((42) & 63)))) | 1))
                                        var v36 uint8 = 127
                                        v30 -= (v35 - 9223372036854775807)
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                _ = v30
                            } else {
                                v21 -= (1.5 * 0.0)
                                var v37 uint8 = 1
                                var v38 float64 = float64(float32(-2.25))
                                _ = v37
                                _ = v38
                            }
                        } else {
                            switch v23 {
                            case "a":
                                v21 -= v21
                            }
                            var v39 uint = 672
                            var v40 uint = 255
                            v3 = (v24 - (v24 & (v24 % (((^(v3))) | 1))))
                            _ = v39
                            _ = v40
                        }
                        switch ("hello" + v4) {
                        case "\n\t":
                            v2[((((v4 + "\n\t") + "\n\t") + v4) + (v4 + "a"))] = 347
                            fmt.Println(v5, v5, fi27)
                        }
                        _ = fi27
                    }
                case "hello":
                    delete(v2, "foo bar")
                }
                var v41 uint64 = 601
                _ = v41
            }
            if false {
                v42 := GenS2{F0: v22, F1: 32767}
                _ = v42
            }
            for fi43 := 0; fi43 < 3; fi43++ {
                v44 := GenS2{F0: -127, F1: 424}
                _ = fi43
                _ = v44
            }
        }
        v24 -= v3
        fmt.Println(v21)
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
