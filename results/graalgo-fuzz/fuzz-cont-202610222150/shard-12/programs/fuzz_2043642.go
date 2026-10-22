package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 int32; F2 bool }
func genFn0() string {
    var v4 string = "\n\t"
    v5 := GenS1{F0: 674}
    if (!true) {
        var v6 uint16 = 42
        for fi7 := 0; fi7 < 16; fi7++ {
            if false {
                var v8 float32 = ((-2.25 / 1.0) - 0.1)
                for fi9 := 0; fi9 < 3; fi9++ {
                    fmt.Println(v6, v6, v6)
                    switch (fi7 << ((2) & 63)) {
                    case 32767:
                        if (!false) {
                            v10 := []int{}
                            v6 = v6
                            _ = v10
                        } else {
                            v11 := []int{(fi7 ^ fi9)}
                            fmt.Println(fi7, v8, fi7)
                            _ = v11
                        }
                    }
                    _ = fi9
                }
                v12 := GenS1{F0: 1}
                _ = v12
                _ = v8
            } else {
                v13 := []int{-9223372036854775808, (fi7 & fi7), fi7, (fi7 ^ (fi7 << ((42) & 63)))}
                v6 += v6
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = v13
            }
            v14 := []int{(fi7 | 42), fi7, fi7, (fi7 | (fi7 - (^(fi7))))}
            _ = fi7
            _ = v14
        }
        _ = v6
    }
    var v15 int32 = -1
    var v16 string = (("a" + (v4 + v4)) + ("" + ""))
    _ = v15
    _ = v16
    _ = v4
    _ = v5
    return v4
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 2
    v3 := []string{"foo bar"}
    v17 := GenS2{F0: 1}
    v2 -= 108
    if false {
        v18 := GenS3{F0: -1, F1: 42, F2: -128}
        fmt.Println(v2)
        _ = v18
    } else {
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    var v20 float64 = 0.1
    v21 := GenS2{F0: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v20 -= 0.1
    var v22 int8 = 42
    v3 = append(v3, ("" + "hello"))
    switch "a" {
    case "":
        v23 := GenS4{F0: (v2 & (v2 & (v2 >> ((65535) & 63)))), F1: 2}
        _ = v23
    case "\n\t":
        if (!(("\n\t" + "hello") != "\n\t")) {
            var v24 uint8 = uint8(v2)
            v20 += v20
            _ = v24
        } else {
            v2 = (v2 << ((42) & 63))
        }
    case "foo bar":
        if false {
            v1[1] = "foo bar"
            var v25 uint16 = (^(v2))
            v26 := GenS5{F0: -117, F1: 1}
            _ = v25
            _ = v26
        }
    default:
        var v27 uint = 907
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        if true {
            switch (-(fi28)) {
            case -2147483648:
                v29 := GenS6{F0: 0, F1: (v2 | 689)}
                _ = v29
            case 552:
                var v30 int32 = 42
                _ = v30
            default:
                for fi31 := 0; fi31 < 16; fi31++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi31
                }
            }
            v32 := GenS7{F0: 316}
            _ = v32
        }
        fmt.Println(v22, v2, v22)
        var v33 bool = (!((3.14159 * 0.1) >= ((((1.5 * float32(float64(3.14159))) - (3.14159 + 0.1)) - (3.14159 / 1.5)) / float32(v20))))
        _ = fi28
        _ = v33
    }
    var v34 float32 = 0.0
    if true {
        _ = len(v1)
        if true {
            var v35 int = -9223372036854775808
            v36 := GenS8{F0: -9223372036854775808, F1: 145, F2: false}
            _ = v35
            _ = v36
        } else {
            v22 -= 111
            v3 = append(v3, "")
            var v37 float64 = v20
            _ = v37
        }
        var v38 uint8 = 0
        _ = v38
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            fmt.Println(v34, fi39, fi39)
            _ = fi39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
}
