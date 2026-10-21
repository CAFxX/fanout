package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
func genFn0() int16 {
    var v7 int32 = 0
    v8 := []int{-2147483648, 1}
    _ = v7
    _ = v8
    return -93
}

func main() {
    var v1 float32 = (1.5 - 1e10)
    var v2 int16 = 42
    v3 := make(map[string]uint32)
    var v4 uint64 = 4294967295
    var v5 uint8 = 128
    v6 := []string{"\n\t", "\n\t", "hello", "\n\t"}
    _ = len(v3)
    switch 657 {
    case 4294967295:
        var v9 float64 = 1.5
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint16 = 857
            for fi12 := 0; fi12 < 16; fi12++ {
                v11 = 1
                var v13 float32 = 0.0
                if true {
                    v14 := GenS1{F0: 9223372036854775807, F1: 1}
                    v15 := GenS1{F0: 1, F1: 1}
                    _ = v14
                    _ = v15
                }
                if false {
                    var v16 int8 = 1
                    var v17 float32 = v13
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v17 = v13
                    case "a":
                        v18 := GenS1{F0: -747, F1: 524}
                        _ = v18
                    default:
                        v19 := GenS1{F0: 42, F1: 594}
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                } else {
                    var v20 float32 = (float32(1e10) + (v1 * float32(1e10)))
                    _ = v20
                }
                v21 := GenS2{F0: (v2 ^ (v2 | (v2 / (((v2 + -32768)) | 1)))), F1: v2}
                _ = fi12
                _ = v13
                _ = v21
            }
            var v22 uint8 = 171
            _ = fi10
            _ = v11
            _ = v22
        }
    }
    v23 := GenS3{F0: -665, F1: -1, F2: v4}
    if false {
        v24 := GenS1{F0: -9223372036854775808, F1: 143}
        switch 509 {
        case 32768:
            var v25 int16 = v2
            _ = v25
        case 707:
            _ = len(v3)
        case 32767:
            v4 = (v4 << ((4294967295) & 63))
        }
        fmt.Println(v5)
        _ = v24
    } else {
        for fi26 := 0; fi26 < 18; fi26++ {
            _ = len(v6)
            _ = fi26
        }
        var v27 uint32 = 4294967295
        _ = v27
    }
    v2 = genFn0()
    var v28 int8 = -1
    if true {
        var v29 uint16 = 4
        var v30 uint32 = 0
        _ = v29
        _ = v30
    } else {
        v1 -= -2.25
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (!true) {
                var v32 uint64 = 4294967295
                var v33 int16 = v2
                switch (fi31 / (((fi31 + -14)) | 1)) {
                case -2147483648:
                    v34 := GenS2{F0: v2, F1: v33}
                    _ = v34
                case 1:
                    if false {
                        v1 -= v1
                        v33 = (v33 >> ((18446744073709551615) & 63))
                    }
                }
                _ = v32
                _ = v33
            } else {
                v2 -= (-(v2))
            }
        } else {
            v35 := GenS2{F0: 42, F1: (v2 | 32767)}
            _ = v35
        }
        fmt.Println(v2, fi31)
        _ = fi31
    }
    fmt.Println(v5)
    var v36 int16 = (^(v2))
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 float32 = ((v1 + float32((1e10 / 0.1))) - v1)
            v40 := GenS2{F0: v36, F1: 42}
            v28 = (v28 + 0)
            v36 += 32767
            if (false && true) {
                var v41 float32 = v1
                fmt.Println(v4, fi37)
                var v42 int16 = 32767
                _ = v41
                _ = v42
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi37
    }
    fmt.Println(v36, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
