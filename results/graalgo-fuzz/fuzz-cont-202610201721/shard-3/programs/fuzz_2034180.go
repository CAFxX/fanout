package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16, p1 int8) int16 {
    p1 += -1
    p0 += (p0 >> ((105) & 63))
    v4 := []int{32768}
    if false {
        var v5 uint = 42
        v6 := GenS1{F0: "foo bar", F1: 42, F2: 255}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: "foo bar", F1: 32768, F2: 2}
        fmt.Println(p0, p0)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return -889
}

func main() {
    var v1 uint = 881
    var v2 float64 = (-2.25 * 3.14159)
    v3 := []int{2147483647, 42, 670, -146}
    for fi8 := 0; fi8 < 14; fi8++ {
        switch (fi8 >> ((v1) & 63)) {
        case 127:
            if (false && false) {
                var v9 int64 = -9223372036854775808
                v2 = 1e10
                _ = len(v3)
                _ = v9
            } else {
                _ = len(v3)
                v2 = v2
            }
        default:
            v10 := GenS2{F0: 100, F1: 0}
            _ = v10
        }
        v1 += 2
        v11 := GenS2{F0: 127, F1: 42}
        _ = fi8
        _ = v11
    }
    switch "" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = 284
        }
    case "\n\t":
        var v12 int8 = 127
        _ = v12
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS3{F0: 255, F1: (("hello" + "a") + "\n\t"), F2: 42}
            _ = fi13
            _ = v14
        }
    }
    if len(v3) > 0 {
        v3[0] = 162
    }
    var v15 int32 = -843
    if true {
        _ = len(v3)
        var v16 int8 = -128
        switch -964 {
        case -2147483648:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            v17 := GenS2{F0: (-(v16)), F1: 2147483648}
            _ = v17
        default:
            if ("a" >= "hello") {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v2)
            } else {
                var v18 uint8 = 128
                if (!(!(!false))) {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := GenS3{F0: v18, F1: "", F2: v15}
                        if len(v3) > 0 {
                            v3[0] = fi19
                        }
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(v16, v18)
                    if (!false) {
                        v15 -= (v15 + 2147483647)
                        if (32767 >= 37) {
                            for fi21 := 0; fi21 < 18; fi21++ {
                                fmt.Println(v1, v15, v16)
                                var v22 uint32 = 197
                                var v23 int64 = int64(v18)
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                }
                var v24 int32 = (v15 << ((v1) & 63))
                _ = v18
                _ = v24
            }
        }
        _ = v16
    }
    _ = len(v3)
    var v25 int16 = -32768
    fmt.Println(v2, v25, v25)
    v3 = append(v3, -31)
    v26 := GenS4{F0: 753}
    v1 = (^(v1))
    fmt.Println(v25, v2)
    v27 := GenS5{F0: 4294967295}
    var v28 int8 = 100
    var v29 string = "foo bar"
    v25 = v25
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS3{F0: 1, F1: v29, F2: v15}
        v28 = 42
        fmt.Println(v1)
        _ = fi30
        _ = v31
    }
    v32 := GenS6{F0: 42, F1: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
