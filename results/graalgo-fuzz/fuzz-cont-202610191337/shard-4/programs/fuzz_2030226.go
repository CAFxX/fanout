package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int64 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: "foo bar", F1: -184, F2: (^(fi6))}
        _ = fi6
        _ = v7
    }
    switch 2147483647 {
    case 242:
        var v8 int8 = -128
        _ = v8
    case 0:
        v9 := GenS2{F0: -972, F1: 1, F2: 1}
        _ = v9
    }
    return -1
}

func main() {
    var v1 uint16 = 491
    var v2 uint64 = 109
    var v3 uint = 42
    var v4 int = 32767
    var v5 uint8 = 42
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{-9223372036854775808}
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint8 = (^(v5))
            _ = fi13
            _ = v14
        }
        switch ("a" + "\n\t") {
        case "":
            fmt.Println(fi11)
        default:
            if ((4294967295 <= 65535) || (545 < int16(v5))) {
                v2 = 0
                if (!(fi11 <= v4)) {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        switch "a" {
                        case "hello":
                            var v16 int16 = 42
                            _ = v16
                        case "\n\t":
                            switch (("hello" + "hello") + ("hello" + "hello")) {
                            case "":
                                if true {
                                    if (v3 < v3) {
                                        var v17 int = 9223372036854775807
                                        _ = v17
                                    } else {
                                        v18 := GenS2{F0: (fi11 >> (((^(v3))) & 63)), F1: -1, F2: 2}
                                        _ = v18
                                    }
                                    v1 += v1
                                    if true {
                                        if true {
                                            var v19 float32 = float32((float64((1.5 - (-2.25 / 0.1))) + (v10 / 1.0)))
                                            v5 = uint8(v1)
                                            var v20 int32 = 42
                                            v21 := GenS2{F0: v4, F1: v20, F2: 41}
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            v22 := GenS2{F0: fi15, F1: 1, F2: uint32(v2)}
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    v23 := GenS1{F0: "hello", F1: (fi15 | fi15), F2: (fi11 + (fi15 << ((v3) & 63)))}
                                    var v24 int8 = 42
                                    var v25 uint32 = 65535
                                    v26 := GenS3{F0: 0}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                } else {
                                    v2 -= v2
                                    if (!(42 >= 0)) {
                                        v2 = (^(v2))
                                        v10 -= -2.25
                                        v27 := GenS4{F0: 1, F1: -118}
                                        _ = v27
                                    }
                                    fmt.Println(v5, fi11, v2)
                                    fmt.Println(v2)
                                    switch int(v1) {
                                    case 32767:
                                        if false {
                                            v10 += (v10 / 0.1)
                                            v28 := GenS5{F0: 4, F1: 1, F2: 434}
                                            _ = v28
                                        }
                                    case 29:
                                        var v29 float32 = float32((float64(0.1) - v10))
                                        _ = v29
                                    case 2147483647:
                                        v30 := GenS5{F0: 100, F1: v1, F2: 591}
                                        _ = v30
                                    default:
                                        switch (fi15 * 128) {
                                        case 4294967295:
                                            var v31 bool = true
                                            if len(v12) > 0 {
                                                v12[0] = (fi11 | int(v3))
                                            }
                                            var v32 float64 = (v10 + (v10 / v10))
                                            _ = v31
                                            _ = v32
                                        }
                                        v12 = append(v12, 42)
                                        if false {
                                            if (!(2 < 128)) {
                                                v5 -= (^(v5))
                                            }
                                            v4 = v4
                                        }
                                    }
                                }
                            case "foo bar":
                                v5 = 243
                            default:
                                var v33 float32 = (-2.25 - 0.1)
                                _ = v33
                            }
                        case "a":
                            if (1 != 0) {
                                if len(v12) > 0 {
                                    v12[0] = fi15
                                }
                            } else {
                                v3 -= (^(v3))
                                v3 = (^(v3))
                            }
                        default:
                            v1 += uint16(v2)
                            if (1e10 >= 0.1) {
                                v4 -= -561
                            }
                            if len(v12) > 0 {
                                v12[0] = fi11
                            }
                        }
                        if (false || (!true)) {
                            v1 = uint16(v3)
                        } else {
                            if false {
                                var v34 int8 = 0
                                _ = v34
                            } else {
                                v35 := GenS6{F0: int8(v2), F1: v2}
                                _ = v35
                            }
                        }
                        _ = fi15
                    }
                }
                fmt.Println(v4, v2)
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    var v36 int = 127
    var v37 uint8 = 88
    v38 := GenS5{F0: -128, F1: v1, F2: genFn0()}
    if true {
        var v39 uint64 = v2
        _ = v39
    }
    if false {
        for fi40 := 0; fi40 < 1; fi40++ {
            fmt.Println(v2, v3, v37)
            v5 += v5
            switch "foo bar" {
            case "a":
                for fi41 := 0; fi41 < 15; fi41++ {
                    v36 -= fi40
                    v42 := []int{v4}
                    _ = fi41
                    _ = v42
                }
            case "":
                v43 := GenS2{F0: v36, F1: -1, F2: 0}
                _ = v43
            case "\n\t":
                v2 += (v2 / (((v2 / (((v2 - v2)) | 1))) | 1))
            }
            _ = fi40
        }
    } else {
        var v44 uint32 = 0
        _ = v44
    }
    v45 := GenS4{F0: v3, F1: 1}
    v10 = 0.1
    switch (v4 + -9223372036854775808) {
    case 0:
        var v46 uint16 = uint16(v3)
        _ = v46
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 uint64 = 42
        _ = fi47
        _ = v48
    }
    fmt.Println(v10, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
}
