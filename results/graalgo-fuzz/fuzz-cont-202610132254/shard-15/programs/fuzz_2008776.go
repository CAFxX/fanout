package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = float64(3.14159)
    v2 := []string{"foo bar"}
    var v3 int64 = 0
    var v4 uint16 = 32767
    v5 := make(map[int]uint32)
    fmt.Println(v4, v1)
    var v6 uint16 = 104
    v6 -= 1
    v7 := GenS1{F0: 564}
    v6 += (^(v4))
    var v8 int32 = -871
    var v9 uint = 403
    switch 32767 {
    case -1:
        v1 += (v1 / (v1 * 0.1))
    default:
        _ = len(v2)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v9, v6)
    if (true && false) {
        if true {
            v10 := GenS1{F0: v8}
            _ = v10
        }
        var v11 int8 = -128
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            v12 := v5[-9223372036854775808]
            _ = v12
        case "\n\t":
            if false {
                if (42 == 486) {
                    if (-2.25 < (float64(0.1) + v1)) {
                        v5[2147483647] = 42
                        fmt.Println(v9, v9, v3)
                        if (!(v3 == (v3 - (v3 - v3)))) {
                            for fi13 := 0; fi13 < 17; fi13++ {
                                v14 := GenS2{F0: v8, F1: (fi13 + (fi13 ^ 127))}
                                var v15 uint16 = (^(v4))
                                v11 += v11
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        } else {
                            if (uint16(v9) < 77) {
                                var v16 int64 = (v3 ^ 42)
                                _ = v16
                            }
                            v9 = (v9 >> (((^(v9))) & 63))
                            var v17 bool = (1 > 42)
                            _ = v17
                        }
                    } else {
                        if ("\n\t" < "a") {
                            var v18 uint32 = 2
                            fmt.Println(v1)
                            for fi19 := 0; fi19 < 12; fi19++ {
                                v1 -= v1
                                var v20 float64 = v1
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        } else {
                            var v21 uint64 = 802
                            v1 += (1.5 + v1)
                            v4 -= v4
                            _ = v21
                        }
                        v4 = (^(v6))
                        var v22 uint16 = (v4 * (v4 << ((42) & 63)))
                        _ = v22
                    }
                }
                _ = len(v2)
                v23 := GenS3{F0: v11, F1: 175, F2: v6}
                _ = v23
            }
        case "hello":
            v24 := GenS1{F0: (^(v8))}
            _ = v24
        default:
            var v25 uint = v9
            _ = v25
        }
        _ = v11
    } else {
        if true {
            var v26 float32 = ((float32(-2.25) * -2.25) / 1.5)
            _ = v26
        } else {
            if false {
                var v27 uint32 = 42
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS4{F0: ("a" + "")}
                    v30 := GenS4{F0: ("foo bar" + "a")}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            } else {
                fmt.Println(v1, v4)
                var v31 bool = (!false)
                v32 := GenS1{F0: v8}
                _ = v31
                _ = v32
            }
            var v33 uint32 = 2147483647
            _ = v33
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            _ = len(v5)
            _ = fi34
        }
        fmt.Println(v8, v4, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
