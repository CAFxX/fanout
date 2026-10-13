package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0() bool {
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float64 = (3.14159 / 1.0)
        _ = fi7
        _ = v8
    }
    v9 := []int{-179, -9223372036854775808, 1, 147}
    switch -2147483648 {
    case -9223372036854775808:
        var v10 uint32 = 2147483648
        _ = v10
    case 127:
        v11 := GenS1{F0: 0}
        _ = v11
    case 128:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        var v12 float64 = 1.5
        _ = v12
    }
    _ = v6
    _ = v9
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 1
    var v3 int8 = 0
    var v4 float32 = 3.14159
    var v5 uint32 = 0
    if (true && (("\n\t" + "foo bar") != (((("a" + "\n\t") + "hello") + "\n\t") + "hello"))) {
        if true {
            v13 := GenS1{F0: 65535}
            fmt.Println(v4, v4, v4)
            _ = v13
        }
        var v14 uint64 = 18446744073709551615
        for fi15 := 0; fi15 < 12; fi15++ {
            v2 -= (v2 ^ (v2 >> ((1) & 63)))
            if ((fi15 <= fi15) || (v3 == 109)) {
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v3, v14, v4)
                    _ = fi16
                }
            } else {
                v17 := GenS1{F0: 2}
                if (42 == -32768) {
                    v2 += v2
                } else {
                    if true {
                        fmt.Println(v14, v5, v3)
                        if false {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                var v19 int64 = int64(v14)
                                if false {
                                    v4 -= (v4 + 0.0)
                                    fmt.Println(v4, v19, v14)
                                    v14 -= v14
                                } else {
                                    fmt.Println(v19, fi15, fi15)
                                }
                                v20 := []int{fi15, fi15}
                                v3 = (^(v3))
                                v2 = (v2 % ((v2) | 1))
                                var v21 uint16 = 32767
                                if true {
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(fi15, v14, v4)
                                        fmt.Println(v5)
                                        fmt.Println(v14)
                                    case "foo bar":
                                        var v22 uint32 = (^(v5))
                                        _ = v22
                                    }
                                }
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            v3 -= (v3 << ((363) & 63))
                            var v23 int16 = 0
                            _ = v23
                        }
                        v2 = v2
                    }
                }
                _ = v17
            }
            _ = fi15
        }
        _ = v14
    }
    v2 -= v2
    v2 += v2
    var v24 uint16 = 164
    _ = len(v1)
    if ((v3 - -100) == (-(v3))) {
        v3 = -99
    } else {
        v4 = v4
    }
    switch "hello" {
    case "":
        v25 := GenS1{F0: v24}
        _ = v25
    }
    var v26 uint32 = (v5 | v5)
    switch "\n\t" {
    case "\n\t":
        var v27 int32 = 0
        _ = v27
    case "a":
        v5 = (v5 - 42)
    }
    var v28 float64 = 3.14159
    for fi29 := 0; fi29 < 7; fi29++ {
        if false {
            var v30 bool = true
            v26 += v5
            v31 := GenS2{F0: 42}
            _ = v30
            _ = v31
        }
        fmt.Println(fi29)
        var v32 float32 = 3.14159
        _ = fi29
        _ = v32
    }
    var v33 float32 = (-2.25 * -2.25)
    var v34 int64 = -112
    var v35 string = ""
    v28 = v28
    fmt.Println(v2, v28, v33)
    for fi36 := 0; fi36 < 4; fi36++ {
        fmt.Println(v3, v35)
        switch 128 {
        case 1:
            v1[(v26 & (v26 ^ v26))] = (v26 | v26)
        case -1:
            fmt.Println(v3, fi36, v34)
        }
        if (!(!(v5 == v26))) {
            v5 = 4294967295
            fmt.Println(v5, fi36)
        } else {
            v4 -= (float32(v28) - v4)
            v2 = (v2 - (v2 | (v2 / (((^(v2))) | 1))))
        }
        _ = fi36
    }
    v2 -= (v2 / ((9) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
