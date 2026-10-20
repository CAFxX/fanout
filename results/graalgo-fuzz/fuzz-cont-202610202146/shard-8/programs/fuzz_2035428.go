package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 uint32 }
func genFn0(p0 int32, p1 int16) int32 {
    var v7 bool = (197 != -74)
    switch 128 {
    case -9223372036854775808:
        var v8 float64 = ((1e10 - 3.14159) + 1e10)
        _ = v8
    case 4294967295:
        v9 := []int{42}
        _ = v9
    case 32767:
        switch 1 {
        case 253:
            if v7 {
                var v10 int32 = p0
                if (2147483647 >= p0) {
                    v11 := GenS1{F0: (("hello" + "\n\t") + "hello"), F1: 18446744073709551615}
                    _ = v11
                } else {
                    v12 := []int{32768}
                    var v13 uint64 = 2
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: 2147483648}
                _ = v10
                _ = v14
            } else {
                var v15 int16 = (p1 * p1)
                v15 = (v15 << ((2) & 63))
                _ = v15
            }
        case -199:
            v7 = (100 > 1)
        default:
            p0 = (p0 + (p0 << ((42) & 63)))
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v7
    return 222
}

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 65535
    v3 := make(map[string]uint32)
    var v4 bool = true
    var v5 int8 = -128
    var v6 int64 = -1
    if ((3.14159 >= 3.14159) || (0.1 == (3.14159 / (1.5 - ((0.0 / (0.0 - 3.14159)) * 1.5))))) {
        v5 += (^(v5))
        var v16 uint8 = 2
        _ = v16
    } else {
        v17 := GenS1{F0: ("" + "a"), F1: 1}
        _ = v17
    }
    v18 := GenS2{F0: (v2 - (v2 >> (((v1 ^ 65535)) & 63)))}
    var v19 uint32 = (v2 - 279)
    if (1 == -485) {
        var v20 float32 = (0.0 + 0.0)
        v19 += v19
        _ = v20
    }
    fmt.Println(v1, v6, v2)
    delete(v3, (("hello" + "hello") + "a"))
    fmt.Println(v19)
    fmt.Println(v6, v6, v5)
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 uint32 = (v2 ^ 137)
        if ((59 < 127) && v4) {
            var v23 float32 = 3.14159
            var v24 uint64 = 18446744073709551615
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v1, fi25)
                fmt.Println(v2, v6)
                _ = fi25
            }
            if v4 {
                var v26 string = ""
                var v27 int16 = 1
                for fi28 := 0; fi28 < 12; fi28++ {
                    v5 += v5
                    var v29 uint32 = v2
                    if v4 {
                        switch "a" {
                        case "foo bar":
                            v2 += uint32(v1)
                        case "hello":
                            if true {
                                var v30 uint16 = 490
                                var v31 int = (fi28 & (fi28 | (fi21 + fi21)))
                                switch "a" {
                                case "hello":
                                    if (1e10 > 1.5) {
                                        var v32 uint32 = 42
                                        _ = v32
                                    } else {
                                        var v33 int16 = (^(v27))
                                        v34 := v3[v26]
                                        fmt.Println(v30, v19)
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                _ = v30
                                _ = v31
                            } else {
                                if v4 {
                                    switch fi28 {
                                    case 83:
                                        v22 += 2147483647
                                    case 32768:
                                        var v35 bool = (255 <= 438)
                                        _ = v35
                                    default:
                                        v36 := GenS3{F0: (v5 >= int8(v6))}
                                        _ = v36
                                    }
                                    if v4 {
                                        fmt.Println(v1, v22, v5)
                                    }
                                } else {
                                    v26 = v26
                                    var v37 int8 = v5
                                    _ = v37
                                }
                                var v38 uint16 = 0
                                _ = v38
                            }
                        default:
                            v5 += (v5 - (-(v5)))
                        }
                        if v4 {
                            var v39 int32 = genFn0(-2147483648, v27)
                            _ = v39
                        } else {
                            var v40 uint = 42
                            var v41 uint16 = 0
                            v22 -= v2
                            _ = v40
                            _ = v41
                        }
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            } else {
                v42 := []int{}
                _ = v42
            }
        }
        v19 = v2
        _ = fi21
        _ = v22
    }
    switch -9223372036854775808 {
    case -2147483648:
        if v4 {
            v43 := GenS4{F0: 635, F1: 72}
            _ = v43
        } else {
            v44 := GenS5{F0: ("a" + "a"), F1: 18446744073709551615, F2: (v19 | (v19 * (v19 ^ (v2 ^ 2))))}
            _ = v44
        }
    case 261:
        v1 += (v1 << ((65535) & 63))
    default:
        if false {
            delete(v3, ("" + ("hello" + ("foo bar" + "a"))))
            fmt.Println(v5)
        }
    }
    var v45 string = ""
    if true {
        fmt.Println(v5, v5, v6)
        v19 = (v2 << ((42) & 63))
    }
    v46 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
