package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint; F1 int32; F2 uint64 }

func main() {
    var v1 int32 = 0
    v2 := []uint32{2147483648, 65535, 4294967295}
    var v3 uint16 = 141
    v4 := make(map[string]uint32)
    fmt.Println(v3, v3)
    var v5 string = (((("foo bar" + ("" + "\n\t")) + "") + "foo bar") + "\n\t")
    if (true || (true && true)) {
        var v6 int = 127
        v3 = 701
        _ = v6
    } else {
        v7 := GenS1{F0: 1000, F1: -1, F2: v5}
        if (v3 > (v3 * (^(v3)))) {
            _ = len(v2)
        } else {
            delete(v4, "hello")
            v5 = "a"
            _ = len(v4)
        }
        v3 -= v3
        _ = v7
    }
    v1 += 42
    v4[v5] = 4294967295
    if ("hello" == v5) {
        v1 += (-(v1))
    }
    switch (((("hello" + v5) + "foo bar") + v5) + v5) {
    case "":
        if false {
            v3 = (^(v3))
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 1; fi9++ {
                    fmt.Println(v1)
                    fmt.Println(v5)
                    if true {
                        var v10 uint16 = v3
                        var v11 uint8 = 64
                        var v12 uint64 = 0
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                if false {
                    v2 = append(v2, 460)
                }
                v1 = (v1 + v1)
                _ = fi8
            }
        } else {
            if (162 <= 128) {
                switch -569 {
                case -568:
                    switch 4294967295 {
                    case 589:
                        v1 = (v1 >> ((18446744073709551615) & 63))
                    default:
                        var v13 int16 = -619
                        _ = v13
                    }
                default:
                    var v14 uint16 = (^(v3))
                    _ = v14
                }
                _ = len(v4)
                var v15 bool = true
                _ = v15
            }
            v5 = ((v5 + ("a" + "")) + (v5 + v5))
            if true {
                var v16 float32 = 1.5
                if (false && (-177 <= 520)) {
                    switch 842 {
                    case -1:
                        var v17 bool = (2147483647 < 0)
                        _ = v17
                    case 32767:
                        v5 = v5
                    case 127:
                        v3 += v3
                    }
                } else {
                    if true {
                        if true {
                            for fi18 := 0; fi18 < 20; fi18++ {
                                if (false || (!((v1 << ((725) & 63)) >= v1))) {
                                    v4[v5] = 65535
                                }
                                _ = fi18
                            }
                            if (float32((0.1 / 1.0)) > 0.0) {
                                v19 := GenS1{F0: 1000, F1: 1, F2: v5}
                                v20 := GenS2{F0: 2147483648}
                                _ = v19
                                _ = v20
                            }
                            var v21 int16 = 1000
                            _ = v21
                        }
                    }
                }
                if ((v16 - v16) >= float32(((float64(v16) + -2.25) - 0.0))) {
                    var v22 int16 = 1000
                    var v23 uint = 255
                    _ = v22
                    _ = v23
                }
                _ = v16
            } else {
                var v24 string = "\n\t"
                _ = len(v2)
                _ = v24
            }
        }
    default:
        var v25 string = v5
        _ = v25
    }
    var v26 int8 = -1
    var v27 uint64 = 42
    switch 34 {
    case -9223372036854775808:
        v28 := GenS3{F0: 0, F1: v27, F2: -778}
        _ = v28
    case 879:
        v29 := GenS4{F0: 2, F1: v1, F2: v27}
        _ = v29
    }
    v3 = (^(v3))
    var v30 uint32 = 588
    var v31 int64 = -807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
