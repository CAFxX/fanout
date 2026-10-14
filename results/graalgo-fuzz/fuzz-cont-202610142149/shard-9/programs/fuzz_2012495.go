package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    v1 := []uint32{2, 1, 960}
    v2 := []int{}
    var v3 int32 = 2147483647
    var v4 uint8 = 255
    var v5 int32 = -142
    fmt.Println(v3)
    var v6 int = 668
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: "a", F1: 42}
            var v10 bool = false
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint16 = 65535
        _ = fi7
        _ = v11
    }
    v6 = (v6 * v6)
    if true {
        v1 = append(v1, 2147483647)
    } else {
        v3 += -997
    }
    if true {
        v3 -= v5
        v12 := GenS2{F0: 42, F1: 0, F2: false}
        _ = v12
    }
    var v13 uint16 = uint16(v4)
    v6 = 32767
    fmt.Println(v13, v5)
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 string = "\n\t"
        if (true && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (fi14 ^ int(v13)) {
            case -2147483648:
                v6 = fi14
            case -9223372036854775808:
                v16 := GenS1{F0: (v15 + v15), F1: v13}
                _ = v16
            }
        } else {
            var v17 uint = 435
            v6 += (-(fi14))
            if (v3 > 1) {
                switch v15 {
                case "":
                    fmt.Println(v6, v5, v15)
                case "foo bar":
                    v13 = v13
                case "\n\t":
                    if ("a" < (v15 + v15)) {
                        v6 = fi14
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi18 := 0; fi18 < 11; fi18++ {
                            if (true || false) {
                                if len(v2) > 0 {
                                    v2[0] = (fi14 * (v6 ^ (^(v6))))
                                }
                                var v19 uint64 = 4294967295
                                fmt.Println(v19, fi18)
                                _ = v19
                            } else {
                                var v20 float64 = ((1.5 * 1.5) + float64((0.0 / (-2.25 + -2.25))))
                                fmt.Println(fi14, v3)
                                _ = v20
                            }
                            var v21 int16 = 443
                            var v22 string = v15
                            _ = fi18
                            _ = v21
                            _ = v22
                        }
                    }
                default:
                    if true {
                        v23 := GenS2{F0: 13, F1: 1, F2: ((v17 <= v17) || false)}
                        fmt.Println(v15, fi14)
                        if (1 >= 869) {
                            v24 := GenS1{F0: "a", F1: v13}
                            _ = v24
                        }
                        _ = v23
                    } else {
                        var v25 uint64 = 1
                        v3 -= v3
                        _ = v25
                    }
                }
                v13 = (v13 / ((42) | 1))
            } else {
                v4 += 1
                var v26 float32 = 3.14159
                _ = v26
            }
            _ = v17
        }
        v27 := GenS3{F0: 898}
        _ = fi14
        _ = v15
        _ = v27
    }
    v6 = (v6 * -577)
    if false {
        if len(v1) > 0 {
            v1[0] = 757
        }
        v28 := GenS4{F0: 1}
        v5 = (v5 >> ((255) & 63))
        _ = v28
    }
    if true {
        if false {
            if len(v1) > 0 {
                v1[0] = 42
            }
            fmt.Println(v5, v13)
        } else {
            fmt.Println(v4, v13)
            switch "\n\t" {
            case "foo bar":
                v29 := GenS2{F0: 2, F1: v13, F2: (v6 == v6)}
                _ = v29
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                v31 := GenS5{F0: 0, F1: 511, F2: v5}
                _ = fi30
                _ = v31
            }
        }
    } else {
        if ((0 >= 32767) && false) {
            var v32 string = ("foo bar" + ("\n\t" + ("a" + "foo bar")))
            _ = v32
        } else {
            _ = len(v2)
            v33 := GenS6{F0: 4294967295, F1: v6, F2: 174}
            _ = v33
        }
        var v34 uint16 = (v13 ^ v13)
        v3 = v3
        _ = v34
    }
    fmt.Println(v6, v4, v13)
    var v35 int64 = -339
    if false {
        var v36 float64 = (-2.25 + 1e10)
        v13 = 410
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
