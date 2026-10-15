package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 bool; F1 string }

func main() {
    var v1 int8 = 70
    var v2 float64 = ((0.1 - 1e10) + 0.1)
    var v3 int32 = -1
    var v4 float32 = -2.25
    v5 := []uint32{2147483647, 1}
    if (!(!true)) {
        v3 = v3
        fmt.Println(v4)
        var v6 int = -2147483648
        _ = v6
    } else {
        var v7 uint64 = 1
        v3 = (^(v3))
        if (0 > 42) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, 2147483648)
            if false {
                v8 := GenS1{F0: 18446744073709551615, F1: 32767, F2: 42}
                var v9 string = ("hello" + "hello")
                _ = v8
                _ = v9
            } else {
                var v10 int64 = -9223372036854775808
                _ = v10
            }
        }
        _ = v7
    }
    switch -9223372036854775808 {
    case 32768:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v2, v3)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 float64 = float64(((0.1 + 1e10) * (3.14159 + (((1e10 + ((v4 - 1.5) - 0.1)) + float32((v2 / (float64((v4 * v4)) * float64((0.1 / (v4 / 3.14159))))))) + ((float32(((v2 / float64(v4)) + v2)) + float32(((3.14159 - v2) + v2))) + v4)))))
    switch 128 {
    case 4294967295:
        if ("a" < "hello") {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int64 = 0
                if (873 == 255) {
                    var v14 int64 = 629
                    v15 := GenS1{F0: 615, F1: 944, F2: 1}
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := GenS2{F0: (v14 % (((v14 % ((v14) | 1))) | 1))}
                        _ = fi16
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                } else {
                    v13 += -1
                    if ((128 != 0) && false) {
                        v5 = append(v5, 4294967295)
                        fmt.Println(v13, v2)
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                    }
                }
                _ = fi12
                _ = v13
            }
            fmt.Println(v11, v11, v1)
            var v18 uint64 = 1
            _ = v18
        } else {
            v11 = float64((0.0 / ((v4 / 1e10) * 1.5)))
            v19 := GenS3{F0: 2}
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 int16 = 42
                v22 := GenS4{F0: (^(fi20)), F1: 42}
                v5 = append(v5, 0)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
    case 741:
        v23 := GenS1{F0: 854, F1: 42, F2: 65535}
        _ = v23
    case 2147483647:
        v4 = v4
    }
    var v24 bool = (269 < 0)
    var v25 int = 2147483647
    if (199 <= 42) {
        var v26 uint32 = 450
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 uint = 835
            for fi29 := 0; fi29 < 14; fi29++ {
                v11 = v11
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi29
            }
            v30 := GenS5{F0: ((!false) && v24), F1: "a"}
            _ = fi27
            _ = v28
            _ = v30
        }
        _ = v26
    } else {
        var v31 bool = (!(65535 == 1))
        _ = v31
    }
    var v32 int = 42
    if true {
        if (v11 == 3.14159) {
            if ((-(v3)) >= -2147483648) {
                for fi33 := 0; fi33 < 19; fi33++ {
                    if v24 {
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    } else {
                        v11 = -2.25
                        var v35 uint8 = 124
                        _ = v35
                    }
                    v4 += ((v4 / v4) / (v4 * v4))
                    var v36 uint8 = 2
                    _ = fi33
                    _ = v36
                }
                for fi37 := 0; fi37 < 15; fi37++ {
                    fmt.Println(v2, v24, fi37)
                    fmt.Println(v11, v24, v25)
                    v2 = 1.5
                    _ = fi37
                }
                v4 += (v4 * (3.14159 + ((v4 - v4) + v4)))
            }
            v2 = v11
            var v38 uint = 136
            _ = v38
        } else {
            var v39 uint64 = 18446744073709551615
            _ = v39
        }
        var v40 int = (v32 ^ v32)
        _ = v40
    } else {
        fmt.Println(v11, v25, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
