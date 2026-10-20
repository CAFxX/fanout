package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint32 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 float32 = 1.5
        _ = fi4
        _ = v5
    }
    v6 := []int{}
    if (!false) {
        var v7 int = 0
        var v8 int8 = 0
        _ = v7
        _ = v8
    } else {
        var v9 uint16 = 42
        _ = v9
    }
    _ = v6
    return 42
}
func genFn1(p0 int, p1 uint64) int32 {
    var v10 bool = true
    v10 = v10
    _ = p0
    _ = p1
    _ = v10
    return -627
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 0.0
    var v3 float32 = 1.5
    var v11 uint = 42
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v3 -= v3
            var v14 int64 = -506
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v2 = 3.14159
        _ = fi15
    }
    v16 := []int{-9223372036854775808, -1, -1}
    switch "\n\t" {
    case "hello":
        if (3.14159 > float32(v2)) {
            var v17 uint = (v11 << (((v11 << (((^(v11))) & 63))) & 63))
            var v18 uint8 = v1
            _ = v17
            _ = v18
        } else {
            switch ("hello" + "\n\t") {
            case "a":
                v19 := GenS1{F0: 372, F1: 1}
                _ = v19
            case "\n\t":
                _ = len(v16)
            case "":
                v20 := GenS1{F0: 42, F1: 105}
                _ = v20
            default:
                v3 = 0.0
            }
        }
    case "\n\t":
        var v21 int = 1
        _ = v21
    default:
        var v22 int64 = 9223372036854775807
        _ = v22
    }
    if (65535 < 990) {
        switch "a" {
        case "a":
            if true {
                var v23 uint = (v11 & v11)
                v3 -= v3
                if false {
                    var v24 int8 = 1
                    _ = v24
                }
                _ = v23
            } else {
                v25 := GenS2{F0: 2147483648, F1: -261}
                if ((!(0.0 > float64((v3 * (0.1 - v3))))) && true) {
                    var v26 uint16 = 1
                    v2 = ((3.14159 / v2) - 3.14159)
                    v27 := GenS3{F0: -185, F1: 32768, F2: 20}
                    _ = v26
                    _ = v27
                } else {
                    v3 += (-2.25 / ((((1.5 / float32(v2)) / v3) * ((-2.25 + (v3 * (1.5 * v3))) * float32(v2))) / 1.0))
                }
                for fi28 := 0; fi28 < 8; fi28++ {
                    switch (fi28 - (-(fi28))) {
                    case 32767:
                        switch 127 {
                        case 55:
                            fmt.Println(fi28)
                        }
                    default:
                        v29 := GenS2{F0: 579, F1: 42}
                        _ = v29
                    }
                    v3 += 3.14159
                    _ = fi28
                }
                _ = v25
            }
        case "hello":
            var v30 float32 = (((v3 + 0.0) * float32(-2.25)) * v3)
            _ = v30
        }
        var v31 string = ((("foo bar" + ("" + ("foo bar" + "foo bar"))) + "") + "a")
        fmt.Println(v11)
        _ = v31
    }
    v32 := GenS3{F0: 32767, F1: -2147483648, F2: 2147483648}
    switch 32767 {
    case -264:
        v33 := GenS2{F0: 0, F1: genFn0()}
        _ = v33
    default:
        if (42 > 0) {
            fmt.Println(v1, v11, v2)
        } else {
            if (!false) {
                for fi34 := 0; fi34 < 11; fi34++ {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        if len(v16) > 0 {
                            v16[0] = (^(fi35))
                        }
                        switch -2147483648 {
                        case 876:
                            var v36 int32 = -1
                            _ = v36
                        case 0:
                            if len(v16) > 0 {
                                v16[0] = (fi35 - (-(fi34)))
                            }
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            var v37 int = 0
                            _ = v37
                        default:
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        }
                        v11 += 188
                        v1 = 127
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi35
                    }
                    if ((genFn1(fi34, 0) <= 0) && (genFn0() != -720)) {
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 uint32 = 4294967295
                            fmt.Println(v39, fi38, v3)
                            _ = fi38
                            _ = v39
                        }
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    } else {
                        v16 = append(v16, -1)
                        switch (fi34 % ((fi34) | 1)) {
                        case 0:
                            v11 += v11
                        default:
                            for fi40 := 0; fi40 < 14; fi40++ {
                                if false {
                                    v3 += (0.0 / v3)
                                    var v41 bool = false
                                    if len(v16) > 0 {
                                        v16[0] = fi34
                                    }
                                    _ = v41
                                } else {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    switch ("a" + "a") {
                                    case "foo bar":
                                        var v42 int16 = -964
                                        _ = v42
                                    case "hello":
                                        v11 += (v11 % ((255) | 1))
                                    }
                                }
                                _ = fi40
                            }
                        }
                        var v43 string = ("" + "foo bar")
                        _ = v43
                    }
                    _ = fi34
                }
            }
            fmt.Println(v11, v2)
        }
    }
    fmt.Println(v11, v2)
    var v44 int32 = -1
    var v45 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v44
    _ = v45
}
