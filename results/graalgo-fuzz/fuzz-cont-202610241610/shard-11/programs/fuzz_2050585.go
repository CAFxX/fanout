package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{-44, -2147483648, 32767, -712}
    v2 := []uint32{65535, 4294967295, 306, 871}
    v3 := make(map[string]bool)
    var v4 uint8 = 128
    var v5 string = "a"
    for fi6 := 0; fi6 < 19; fi6++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            _ = len(v3)
        case "":
            v5 = v5
        }
        v7 := GenS1{F0: -22}
        _ = fi6
        _ = v7
    }
    var v8 float32 = 1e10
    fmt.Println(v4, v5)
    v9 := GenS2{F0: v4}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint8 = (v4 % (((^(v4))) | 1))
    v11 := GenS3{F0: 18446744073709551615}
    fmt.Println(v5, v5)
    switch (("a" + "hello") + v5) {
    case "":
        var v12 float64 = (0.0 * 3.14159)
        _ = v12
    case "hello":
        fmt.Println(v8)
    }
    if (v8 == v8) {
        var v13 bool = true
        switch ("hello" + v5) {
        case "":
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int32 = 453
                if (!v13) {
                    if false {
                        v16 := GenS3{F0: 144}
                        _ = v16
                    }
                    var v17 uint8 = (v10 >> ((564) & 63))
                    var v18 uint8 = v17
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v10, v15, v4)
                    if (1 != 18446744073709551615) {
                        v8 -= (v8 + v8)
                        v4 += 255
                    } else {
                        v19 := GenS1{F0: 65}
                        var v20 uint64 = 1
                        fmt.Println(v5)
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi14
                _ = v15
            }
        case "a":
            var v21 int32 = 423
            _ = v21
        default:
            if (!false) {
                var v22 int16 = -963
                _ = v22
            } else {
                v8 = (v8 + (-2.25 - float32(((0.1 + 1e10) + 0.0))))
                v3[v5] = false
            }
        }
        v10 += (v4 / (((v4 * (v10 >> ((0) & 63)))) | 1))
        _ = v13
    }
    v23 := GenS1{F0: 0}
    switch 405 {
    case 4294967295:
        v8 += (v8 + ((0.0 - float32((1.5 - (3.14159 - ((1.5 / (3.14159 / 1e10)) - -2.25))))) + v8))
    case 127:
        var v24 int = -443
        _ = v24
    case -2147483648:
        switch 1 {
        case 0:
            v25 := GenS3{F0: 0}
            _ = v25
        case 1:
            v26 := v3[v5]
            _ = v26
        default:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v27 int8 = -40
                _ = v27
            } else {
                var v28 uint8 = (v10 % ((138) | 1))
                if true {
                    if false {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v10 += v10
                        } else {
                            v5 = v5
                            v29 := GenS1{F0: 100}
                            _ = v29
                        }
                    }
                }
                v30 := GenS4{F0: 2}
                _ = v28
                _ = v30
            }
        }
    }
    fmt.Println(v5, v5, v4)
    _ = len(v1)
    v10 += v10
    v31 := v3[v5]
    if (v5 < (v5 + v5)) {
        for fi32 := 0; fi32 < 13; fi32++ {
            v4 = v4
            if v31 {
                _ = len(v1)
            }
            if false {
                fmt.Println(v8)
            } else {
                v5 = (v5 + (v5 + v5))
                fmt.Println(v8)
                fmt.Println(v4, v10)
            }
            _ = fi32
        }
    }
    switch 4294967295 {
    case 42:
        if (1 == 1) {
            var v33 uint = 586
            var v34 uint8 = (v10 - (v10 << ((v33) & 63)))
            _ = v33
            _ = v34
        }
    case -446:
        v8 = (0.0 + ((((v8 * (0.1 + (v8 * v8))) - (v8 * v8)) / 3.14159) - (3.14159 + 3.14159)))
    case 4294967295:
        v5 = v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
