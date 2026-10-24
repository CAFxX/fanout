package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint64; F1 int; F2 int32 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int; F2 uint32 }
type GenS10 struct { F0 bool; F1 bool; F2 uint }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{2147483648, 0, 1, 604}
    v3 := make(map[int]string)
    var v4 uint64 = 1
    var v5 int8 = 127
    var v6 bool = false
    switch 128 {
    case -1:
        v7 := GenS1{F0: 40}
        _ = v7
    case -245:
        v6 = v6
    case 57:
        var v8 int = 1
        _ = v8
    }
    var v9 bool = ((244 > 255) && (!("" <= "")))
    v5 = 42
    var v10 uint64 = 0
    _ = len(v2)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS2{F0: v1}
            if false {
                v14 := GenS3{F0: 42}
                var v15 uint16 = 0
                var v16 uint = 501
                _ = v14
                _ = v15
                _ = v16
            }
            switch fi12 {
            case -9223372036854775808:
                var v17 uint16 = 255
                _ = v17
            case 127:
                switch fi11 {
                case 32768:
                    v2 = append(v2, 187)
                case 0:
                    switch 0 {
                    case 9223372036854775807:
                        v4 = v10
                    }
                case -2147483648:
                    v3[fi12] = "\n\t"
                default:
                    switch (fi11 % (((fi11 % ((fi11) | 1))) | 1)) {
                    case -1:
                        v18 := GenS4{F0: 255}
                        _ = v18
                    default:
                        v19 := GenS5{F0: 4294967295, F1: 0, F2: 2147483647}
                        _ = v19
                    }
                }
            case 1:
                fmt.Println(v10, v5, fi12)
            default:
                v5 += -128
            }
            _ = fi12
            _ = v13
        }
        if (42 <= 2) {
            v6 = v9
            var v20 int8 = 100
            _ = v20
        } else {
            v1 += v1
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS6{F0: 0, F1: (fi11 / (((fi11 & fi21)) | 1)), F2: -2147483648}
                var v23 uint8 = 255
                var v24 int32 = 1
                if v9 {
                    if (false || (v24 != (v24 / ((1) | 1)))) {
                        v25 := GenS7{F0: 648, F1: v23}
                        var v26 int16 = -32768
                        _ = v25
                        _ = v26
                    }
                    v23 += (v23 - (v23 << ((0) & 63)))
                }
                if (!v6) {
                    v27 := v3[fi21]
                    var v28 uint = 2
                    _ = v27
                    _ = v28
                } else {
                    var v29 int16 = -341
                    switch "\n\t" {
                    case "\n\t":
                        if v6 {
                            v2 = append(v2, 2147483647)
                        } else {
                            var v30 bool = (!((v24 * 0) <= -2147483648))
                            if v30 {
                                var v31 bool = (!((((float32(0.1) / 1.0) * ((3.14159 * 0.1) * -2.25)) >= 1.5) || (("" + "hello") >= "")))
                                _ = v31
                            }
                            v32 := GenS8{F0: 65535}
                            _ = v30
                            _ = v32
                        }
                    case "foo bar":
                        v33 := GenS7{F0: 42, F1: 138}
                        _ = v33
                    }
                    _ = v29
                }
                v24 -= v24
                var v34 bool = v9
                v35 := GenS9{F0: 4294967295, F1: (fi11 >> ((0) & 63)), F2: 56}
                v23 = 0
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
                _ = v34
                _ = v35
            }
        }
        _ = fi11
    }
    v4 = v10
    v5 = v5
    v3[1] = (("hello" + "\n\t") + (("a" + "hello") + ""))
    v36 := GenS10{F0: v6, F1: (0.0 == 1e10), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
