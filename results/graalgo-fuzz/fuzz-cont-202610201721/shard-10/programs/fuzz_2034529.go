package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 uint64, p1 int16, p2 uint16) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint32 = 65535
        if true {
            var v8 float32 = -2.25
            v9 := GenS1{F0: 1, F1: p2, F2: -2147483648}
            _ = v8
            _ = v9
        } else {
            fmt.Println(v7, p1)
        }
        _ = fi6
        _ = v7
    }
    v10 := GenS1{F0: 42, F1: p2, F2: 728}
    p2 = 32767
    var v11 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 float32 = (1e10 / -2.25)
    v2 := make(map[uint32]bool)
    v3 := []int{-1, 32767}
    var v4 string = "foo bar"
    var v5 uint = 1
    var v12 uint = (v5 * (v5 % ((v5) | 1)))
    v1 += (v1 * 3.14159)
    var v13 int32 = 5
    if true {
        fmt.Println(v13)
        v1 = ((((v1 + v1) - v1) * ((1.5 * (0.1 + v1)) + (-2.25 / (0.1 + (v1 / v1))))) - (1.5 / v1))
    }
    v12 -= (v12 * (v5 << ((v5) & 63)))
    v12 = 4294967295
    var v14 float64 = 0.1
    if ((v4 >= v4) || true) {
        for fi15 := 0; fi15 < 20; fi15++ {
            if (fi15 < fi15) {
                if true {
                    if ((fi15 / ((167) | 1)) > (fi15 & (fi15 & (fi15 << ((455) & 63))))) {
                        if (-32768 == 1) {
                            fmt.Println(v14)
                            if false {
                                if (-77 == -1) {
                                    v16 := GenS2{F0: v4, F1: 497}
                                    v5 = 1
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v17 int64 = 42
                                        _ = v17
                                    case "a":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v16
                                } else {
                                    _ = len(v2)
                                    var v18 int16 = 1
                                    var v19 string = ("" + "")
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                        v5 = (^(v12))
                    }
                    _ = len(v2)
                } else {
                    if (42 > 128) {
                        fmt.Println(v4)
                    } else {
                        var v20 int8 = 15
                        _ = v20
                    }
                }
                if (!(float64(v1) == 1.5)) {
                    if true {
                        switch (fi15 & fi15) {
                        case 798:
                            v13 -= 1
                        case 1:
                            delete(v2, 1)
                        case -9223372036854775808:
                            v5 = (^(v5))
                        }
                        v5 -= v12
                        _ = len(v2)
                    } else {
                        v21 := GenS1{F0: 9223372036854775807, F1: 102, F2: (v13 & (v13 + (v13 / ((-2147483648) | 1))))}
                        _ = v21
                    }
                    fmt.Println(v4)
                } else {
                    v4 = v4
                    v5 = v12
                    v3 = append(v3, fi15)
                }
                v22 := GenS2{F0: "a", F1: 4294967295}
                _ = v22
            }
            var v23 uint64 = 42
            var v24 uint = v5
            _ = fi15
            _ = v23
            _ = v24
        }
    } else {
        switch (v4 + v4) {
        case "hello":
            var v25 float32 = (0.0 / v1)
            _ = v25
        }
        var v26 uint64 = 0
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int64 = 35
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 int = fi27
                _ = fi29
                _ = v30
            }
            var v31 int64 = v28
            _ = fi27
            _ = v28
            _ = v31
        }
        _ = v26
    }
    _ = len(v2)
    if (0 >= -1) {
        var v32 uint16 = 20
        var v33 string = v4
        v34 := GenS3{F0: 149, F1: 2, F2: (v32 & (^(v32)))}
        _ = v32
        _ = v33
        _ = v34
    }
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
