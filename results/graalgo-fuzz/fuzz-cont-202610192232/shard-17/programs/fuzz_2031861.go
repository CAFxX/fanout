package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 int64) int8 {
    var v4 uint16 = 1
    p0 = (p0 % ((42) | 1))
    p0 -= p0
    _ = p0
    _ = p1
    _ = v4
    return 30
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{127}
    v3 := make(map[int]int)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -2147483648 {
    case 127:
        var v5 uint64 = 0
        _ = v5
    case 432:
        var v6 int32 = 2147483647
        _ = v6
    case -1:
        v7 := GenS1{F0: false}
        _ = v7
    }
    var v8 float32 = 1.5
    v8 = (0.0 * v8)
    var v9 uint32 = 233
    switch "hello" {
    case "a":
        var v10 uint16 = 773
        _ = v10
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v9)
            _ = fi11
        }
    }
    if len(v2) > 0 {
        v2[0] = 916
    }
    var v12 uint32 = 767
    v8 = (v8 * v8)
    if (971 < 757) {
        v13 := GenS1{F0: false}
        v14 := GenS1{F0: (true || (!(float32((((3.14159 + (((3.14159 / 1e10) - (-2.25 + (0.0 - 0.1))) + 0.0)) / 0.1) / (-2.25 / 1.0))) >= (v8 / v8))))}
        _ = v13
        _ = v14
    }
    v2 = append(v2, 2147483647)
    if (3.14159 > 0.0) {
        if (-2147483648 == 0) {
            delete(v1, "\n\t")
            for fi15 := 0; fi15 < 5; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    if false {
                        if (v8 == float32((((0.1 / 3.14159) / float64(((((3.14159 * 0.1) / 1.0) / (((((v8 / 0.1) / (((v8 * v8) / (v8 * (float32((1.5 + (1.5 / 1.0))) * v8))) - ((v8 * -2.25) + 3.14159))) - 0.0) - v8) + float32(3.14159))) * (float32(3.14159) + (v8 - 0.0))))) - 3.14159))) {
                            var v17 int32 = 595
                            if false {
                                v18 := GenS1{F0: (false && false)}
                                if false {
                                    var v19 bool = false
                                    _ = v19
                                }
                                _ = v18
                            } else {
                                switch ("hello" + "") {
                                case "":
                                    var v20 float32 = float32(0.1)
                                    _ = v20
                                case "a":
                                    v21 := GenS2{F0: 42, F1: 390, F2: 0}
                                    var v22 float64 = float64(v8)
                                    _ = v21
                                    _ = v22
                                }
                            }
                            var v23 int16 = 337
                            v24 := GenS1{F0: true}
                            _ = v17
                            _ = v23
                            _ = v24
                        } else {
                            v9 -= 65535
                        }
                    }
                    _ = fi16
                }
                _ = fi15
            }
            _ = len(v1)
        }
        fmt.Println(v8)
        if false {
            if false {
                v12 = v9
            } else {
                v1[""] = (v12 < v9)
            }
        }
    } else {
        v25 := GenS2{F0: 4294967295, F1: 184, F2: 4294967295}
        if (808 != 831) {
            var v26 uint = 18446744073709551615
            v26 -= (^(v26))
            if (-1 == 1000) {
                var v27 int = -1
                v26 = v26
                if ((v8 * 1.5) != v8) {
                    v26 = (v26 >> (((v26 & v26)) & 63))
                    _ = len(v1)
                    if (!false) {
                        var v28 uint = v26
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch (^(v27)) {
                            case -9223372036854775808:
                                fmt.Println(v8, v27, v28)
                            }
                        } else {
                            v28 -= 319
                            var v29 uint = v26
                            fmt.Println(v29)
                            _ = v29
                        }
                        _ = v28
                    }
                } else {
                    v12 += 0
                    fmt.Println(v9, v12)
                    for fi30 := 0; fi30 < 14; fi30++ {
                        v31 := GenS1{F0: false}
                        var v32 int64 = 42
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                _ = v27
            }
            _ = v26
        }
        _ = v25
    }
    if true {
        v8 = (1.5 * 1.5)
        var v33 int32 = -1
        _ = v33
    } else {
        v8 -= (float32(-2.25) / float32(((-2.25 / -2.25) / 1.5)))
    }
    fmt.Println(v12, v8)
    var v34 int32 = int32(v12)
    v34 = v34
    for fi35 := 0; fi35 < 1; fi35++ {
        v9 = v12
        switch 128 {
        case 32767:
            v8 = v8
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v8
    _ = v9
}
