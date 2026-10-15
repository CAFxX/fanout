package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 uint }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint8 }

func main() {
    var v1 float32 = float32(1e10)
    v2 := []int{}
    var v3 int = 424
    var v4 float32 = 0.1
    if (962 > v3) {
        v5 := GenS1{F0: 65535}
        _ = v5
    } else {
        v2 = append(v2, v3)
        v1 -= v4
        v4 -= ((v4 - (-2.25 - (v1 / (v1 * v1)))) + float32((1e10 - 1e10)))
    }
    v6 := GenS1{F0: 1}
    fmt.Println(v3, v1)
    for fi7 := 0; fi7 < 16; fi7++ {
        if (-128 > -1) {
            if false {
                var v8 uint32 = 2147483648
                if (!false) {
                    v9 := GenS2{F0: v8, F1: 255}
                    if len(v2) > 0 {
                        v2[0] = (fi7 | (fi7 ^ fi7))
                    }
                    var v10 float32 = v4
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        var v12 int16 = 32767
                        v13 := GenS3{F0: 1, F1: 720}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v4 = v1
                    var v14 uint16 = 65535
                    _ = v14
                }
                _ = v8
            } else {
                var v15 int = fi7
                _ = v15
            }
        } else {
            fmt.Println(fi7, v4)
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS3{F0: 633, F1: 1}
                v3 = (v3 | (^(fi16)))
                v1 += 3.14159
                _ = fi16
                _ = v17
            }
        }
        var v18 int = (v3 | v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi7
        _ = v18
    }
    v4 = v1
    _ = len(v2)
    if (!true) {
        v1 = v1
    }
    if false {
        var v19 uint64 = 1
        if len(v2) > 0 {
            v2[0] = v3
        }
        v1 = ((0.0 - v1) / v4)
        _ = v19
    } else {
        v3 += (-(v3))
        if (365 > 1) {
            if true {
                v3 -= (v3 & v3)
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v1 -= 3.14159
                        _ = len(v2)
                        _ = fi21
                    }
                    _ = fi20
                }
                var v22 float32 = v1
                _ = v22
            }
        }
    }
    var v23 string = "hello"
    switch (v3 - v3) {
    case 956:
        if len(v2) > 0 {
            v2[0] = v3
        }
    case -1:
        v24 := GenS4{F0: 0, F1: (!false), F2: 2}
        _ = v24
    case 744:
        v25 := GenS5{F0: 814, F1: 996}
        _ = v25
    }
    v3 += v3
    var v26 float32 = v1
    v27 := GenS6{F0: (v23 + (v23 + "foo bar")), F1: 65535}
    v4 += 0.1
    if (false || false) {
        for fi28 := 0; fi28 < 16; fi28++ {
            switch v23 {
            case "foo bar":
                if (2 <= 1) {
                    if (false || true) {
                        fmt.Println(fi28)
                        _ = len(v2)
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v30 := GenS7{F0: 1, F1: 924, F2: 195}
                            v23 = "a"
                            var v31 int32 = -267
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v3 = 495
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v3 -= (v3 << ((2) & 63))
                        switch v23 {
                        case "":
                            var v33 uint64 = 0
                            _ = v33
                        }
                        _ = fi32
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = 32767
            }
            _ = fi28
        }
    }
    v4 += v1
    if len(v2) > 0 {
        v2[0] = (-(v3))
    }
    var v34 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v6
}
