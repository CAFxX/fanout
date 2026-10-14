package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint, p1 int16, p2 int64) uint16 {
    p1 -= int16(p2)
    var v4 bool = false
    v5 := GenS1{F0: 828}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: -1}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    var v3 int64 = 1
    v2 = float64(0.1)
    var v8 uint = 4294967295
    v8 = v8
    v3 += v3
    switch "" {
    case "hello":
        if (!true) {
            v3 -= (v3 ^ 0)
        } else {
            v1[42] = (("\n\t" + "hello") + "hello")
            var v9 uint8 = 2
            _ = v9
        }
    default:
        var v10 int8 = -1
        _ = v10
    }
    var v11 int32 = -91
    v8 -= v8
    v12 := []int{2147483647}
    var v13 int = 799
    switch ("foo bar" + ("" + "\n\t")) {
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            if (!false) {
                var v15 float64 = 1.5
                v13 = -1
                var v16 string = "foo bar"
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: 1}
                v3 = v3
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    v3 -= -9223372036854775808
                    var v19 int64 = (v3 >> ((v8) & 63))
                    var v20 int16 = -508
                    _ = v19
                    _ = v20
                } else {
                    if (v8 != 65535) {
                        if true {
                            fmt.Println(fi14, v2)
                        }
                        if len(v12) > 0 {
                            v12[0] = int(v11)
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v21 uint8 = 1
                        _ = v21
                    }
                    v2 -= (v2 + float64(0.1))
                    var v22 bool = (float64(1.5) >= 1.5)
                    switch (v13 & fi18) {
                    case 42:
                        if (1e10 > 1.5) {
                            fmt.Println(v3)
                            if true {
                                v13 = (^(fi14))
                            }
                        }
                        switch "a" {
                        case "":
                            var v23 uint8 = 127
                            _ = v23
                        }
                        if false {
                            v24 := GenS1{F0: -2147483648}
                            fmt.Println(v8)
                            _ = v24
                        }
                    case 127:
                        v25 := GenS1{F0: v11}
                        _ = v25
                    case 0:
                        v8 = (v8 << ((255) & 63))
                        v3 += 389
                        if v22 {
                            v26 := GenS2{F0: 255}
                            var v27 int64 = v3
                            _ = v26
                            _ = v27
                        }
                    default:
                        var v28 float32 = 1e10
                        _ = v28
                    }
                    if true {
                        switch "a" {
                        case "foo bar":
                            v13 = (fi14 >> (((^(v8))) & 63))
                        }
                    } else {
                        switch (fi14 >> ((65535) & 63)) {
                        case 32768:
                            if ((252 >= 42) || v22) {
                                if len(v12) > 0 {
                                    v12[0] = fi14
                                }
                            }
                        case 9223372036854775807:
                            var v29 uint32 = 42
                            _ = v29
                        case 574:
                            v30 := GenS2{F0: 2}
                            _ = v30
                        }
                        v31 := GenS3{F0: 1, F1: 42}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v31
                    }
                    _ = v22
                }
                var v32 bool = false
                var v33 uint64 = 2
                _ = fi18
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                if len(v12) > 0 {
                    v12[0] = 523
                }
                _ = fi34
            }
            _ = fi14
        }
    }
    var v35 uint32 = 370
    v12 = append(v12, v13)
    fmt.Println(v35)
    v13 -= v13
    var v36 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v8
}
