package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int32 = 42
    v2 := []uint32{631}
    v3 := []int{-2147483648, 13}
    if (false && (!true)) {
        v4 := GenS1{F0: -1, F1: 32767, F2: 116}
        var v5 int8 = -39
        for fi6 := 0; fi6 < 4; fi6++ {
            v5 -= int8(v1)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi6
        }
        _ = v4
        _ = v5
    }
    v7 := GenS1{F0: (-(v1)), F1: -32768, F2: 128}
    switch "" {
    case "hello":
        var v8 uint8 = 2
        _ = v8
    }
    v9 := GenS1{F0: (^(v1)), F1: 742, F2: 128}
    v1 = (v1 ^ (v1 & (^(v1))))
    if (v1 >= 0) {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float64 = -2.25
            fmt.Println(v11, v11)
            for fi12 := 0; fi12 < 12; fi12++ {
                _ = len(v3)
                if (4294967295 <= 486) {
                    var v13 int64 = 0
                    fmt.Println(fi12, fi10)
                    v13 += (v13 + (v13 & v13))
                    _ = v13
                } else {
                    var v14 uint = 2
                    var v15 string = "\n\t"
                    _ = v14
                    _ = v15
                }
                if true {
                    var v16 float32 = 0.0
                    var v17 int32 = 42
                    _ = v16
                    _ = v17
                } else {
                    var v18 bool = false
                    _ = v18
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        v19 := GenS2{F0: 920, F1: 0, F2: 32767}
        _ = v19
    }
    var v20 uint16 = 2
    v1 = v1
    v21 := GenS1{F0: v1, F1: -1, F2: 23}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v3) > 0 {
        v3[0] = -214
    }
    switch (("\n\t" + "a") + "hello") {
    case "a":
        if false {
            v22 := GenS2{F0: (v20 % ((v20) | 1)), F1: 2, F2: -32768}
            v1 += v1
            _ = v22
        } else {
            if (false && ((true || true) || (65535 <= 2147483648))) {
                v23 := GenS3{F0: 127, F1: -32768, F2: false}
                _ = v23
            } else {
                switch 9223372036854775807 {
                case -2147483648:
                    var v24 uint16 = 354
                    _ = v24
                case 458:
                    if (0 < 65535) {
                        var v25 uint16 = (v20 / (((v20 >> ((65535) & 63))) | 1))
                        _ = v25
                    } else {
                        var v26 int16 = -32768
                        fmt.Println(v20)
                        _ = v26
                    }
                }
            }
            for fi27 := 0; fi27 < 20; fi27++ {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS2{F0: 2, F1: 4294967295, F2: -32768}
                    var v30 bool = false
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    v1 += v1
                    if len(v3) > 0 {
                        v3[0] = (fi31 ^ (fi31 ^ fi31))
                    }
                    _ = fi31
                }
                v20 -= (v20 ^ (v20 + v20))
                _ = fi27
            }
        }
    default:
        if (167 < 42) {
            v32 := GenS1{F0: 0, F1: -986, F2: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v32
        }
    }
    if true {
        var v33 int = int(v20)
        if true {
            switch 42 {
            case 132:
                var v34 uint16 = v20
                _ = v34
            default:
                for fi35 := 0; fi35 < 16; fi35++ {
                    if len(v3) > 0 {
                        v3[0] = v33
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v36 := GenS2{F0: (v20 % (((v20 / (((v20 | v20)) | 1))) | 1)), F1: 2147483648, F2: 0}
                    _ = fi35
                    _ = v36
                }
            }
            if true {
                var v37 float64 = 0.1
                for fi38 := 0; fi38 < 2; fi38++ {
                    v39 := GenS4{F0: 0}
                    switch v33 {
                    case 9223372036854775807:
                        v40 := GenS3{F0: 0, F1: -1, F2: true}
                        _ = v40
                    default:
                        v1 += 1
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
        } else {
            var v41 uint8 = 2
            if true {
                if false {
                    switch (v33 & v33) {
                    case 2147483647:
                        var v42 int64 = 9223372036854775807
                        _ = v42
                    case 465:
                        var v43 float32 = 3.14159
                        _ = v43
                    case 128:
                        _ = len(v2)
                    default:
                        _ = len(v3)
                    }
                    v1 = (^(v1))
                    switch (("" + "a") + "foo bar") {
                    case "":
                        v20 += uint16(v41)
                    default:
                        v20 = (v20 % (((v20 << ((345) & 63))) | 1))
                    }
                }
            } else {
                if false {
                    v44 := GenS1{F0: v1, F1: -32768, F2: uint8(v20)}
                    _ = len(v2)
                    _ = v44
                }
            }
            _ = v41
        }
        var v45 int8 = 127
        _ = v33
        _ = v45
    } else {
        v20 += (^(v20))
    }
    v1 = (v1 & v1)
    var v46 uint16 = v20
    var v47 float32 = 0.0
    v1 = (v1 & (v1 / ((v1) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v46
    _ = v47
    _ = v7
    _ = v9
}
