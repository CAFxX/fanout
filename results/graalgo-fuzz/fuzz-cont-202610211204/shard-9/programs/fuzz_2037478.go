package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int8; F2 int8 }
type GenS6 struct { F0 uint32; F1 int8 }
func genFn0() int16 {
    var v4 uint32 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        v4 = v4
        switch "hello" {
        case "\n\t":
            switch (fi5 << ((109) & 63)) {
            case 0:
                var v6 uint64 = 2
                _ = v6
            case -1:
                fmt.Println(fi5, fi5)
            case 613:
                v4 = (v4 * 1)
            }
        case "foo bar":
            v4 = (v4 * (v4 / (((v4 << ((0) & 63))) | 1)))
        case "a":
            v4 = (^(v4))
        default:
            var v7 uint8 = 71
            _ = v7
        }
        _ = fi5
    }
    var v8 uint = 1
    if (((-9223372036854775808 < -443) || false) && (!(!(1 >= 9223372036854775807)))) {
        for fi9 := 0; fi9 < 15; fi9++ {
            v4 = 2147483647
            v8 += (v8 << ((v8) & 63))
            var v10 uint64 = 882
            _ = fi9
            _ = v10
        }
        v8 += 651
    } else {
        v8 += (v8 - v8)
    }
    _ = v4
    _ = v8
    return -32768
}

func main() {
    var v1 int64 = 1
    var v2 float32 = (1e10 / 3.14159)
    v3 := make(map[uint32]bool)
    fmt.Println(v2, v2)
    v11 := []int{128, -9223372036854775808}
    if (!true) {
        v12 := v3[2147483648]
        _ = v12
    } else {
        var v13 uint64 = 42
        switch 636 {
        case -271:
            var v14 float32 = v2
            _ = v14
        case 42:
            v15 := GenS1{F0: 4294967295}
            _ = v15
        }
        _ = v13
    }
    v1 = v1
    v16 := GenS2{F0: 4294967295, F1: (v1 % (((v1 & 929)) | 1))}
    var v17 uint = 4294967295
    if true {
        v17 += 255
    } else {
        if false {
            v18 := GenS3{F0: 1}
            switch 1 {
            case -69:
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS3{F0: int8(fi19)}
                    for fi21 := 0; fi21 < 15; fi21++ {
                        if true {
                            fmt.Println(v2, fi19)
                        }
                        if (true || (65535 != v17)) {
                            if (false && ((fi21 ^ (fi21 >> ((v17) & 63))) == (fi19 % (((fi21 & fi21)) | 1)))) {
                                var v22 int16 = 32767
                                v1 += v1
                                _ = v22
                            }
                            if (128 > 128) {
                                if ((true || (false || false)) && false) {
                                    var v23 int32 = 480
                                    _ = v23
                                }
                                v2 += ((0.0 + 0.1) * 1e10)
                            } else {
                                var v24 int = fi19
                                _ = v24
                            }
                        } else {
                            v2 = v2
                            v17 -= 65535
                            switch ("hello" + "\n\t") {
                            case "a":
                                v3[0] = ((!false) || false)
                            }
                        }
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
            case -816:
                var v25 float32 = v2
                _ = v25
            }
            fmt.Println(v17)
            _ = v18
        } else {
            v26 := GenS4{F0: 709}
            _ = v26
        }
    }
    fmt.Println(v17, v2)
    if (!true) {
        if (true || (!(!true))) {
            delete(v3, 42)
            fmt.Println(v2, v1, v2)
        } else {
            var v27 bool = false
            for fi28 := 0; fi28 < 13; fi28++ {
                v2 += -2.25
                _ = fi28
            }
            _ = v27
        }
        v29 := GenS5{F0: int32(v17), F1: 1, F2: 1}
        _ = v29
    }
    switch 221 {
    case -697:
        v2 += v2
    default:
        var v30 uint = v17
        _ = v30
    }
    v11 = append(v11, 127)
    switch 1 {
    case 42:
        v31 := GenS6{F0: 2, F1: 0}
        _ = v31
    case 0:
        v2 += v2
    case -928:
        if (42 >= 0) {
            switch 676 {
            case 32767:
                v1 = (v1 << ((578) & 63))
            }
            _ = len(v11)
            fmt.Println(v1)
        }
    default:
        var v32 int64 = (v1 & (v1 % (((v1 - v1)) | 1)))
        _ = v32
    }
    v11 = append(v11, 32767)
    v33 := GenS4{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
}
