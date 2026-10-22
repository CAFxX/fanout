package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }

func main() {
    var v1 float32 = (0.0 * 1.5)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    var v6 uint16 = 1
    switch ("a" + "a") {
    case "":
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 bool = true
            v3[128] = (!("\n\t" < ""))
            if true {
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: 2}
                    _ = fi9
                    _ = v10
                }
                var v11 int32 = -2147483648
                v6 = (v6 % ((v6) | 1))
                _ = v11
            } else {
                var v12 uint64 = 246
                var v13 uint8 = 42
                _ = v12
                _ = v13
            }
            _ = fi7
            _ = v8
        }
    case "a":
        v6 += 65535
    case "hello":
        v14 := GenS1{F0: 42}
        _ = v14
    default:
        v15 := GenS1{F0: 0}
        _ = v15
    }
    if ((42 >= 1) && false) {
        v16 := GenS1{F0: 2}
        _ = v16
    }
    switch "hello" {
    case "foo bar":
        if (uint32(v6) != 2147483648) {
            var v17 uint64 = 1
            v18 := v2[337]
            _ = len(v3)
            _ = v17
            _ = v18
        }
    case "a":
        if (-1 <= 1) {
            if (253 <= 2) {
                switch "" {
                case "a":
                    var v19 int32 = 484
                    _ = v19
                }
                var v20 float64 = 0.1
                _ = len(v3)
                _ = v20
            }
            var v21 float32 = v1
            var v22 int64 = 9223372036854775807
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(v6, fi23)
                var v24 uint32 = 0
                _ = fi23
                _ = v24
            }
        }
    case "":
        v6 = v6
    }
    switch 32767 {
    case -1:
        if false {
            v4[-1] = 42
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := []int{-39, (fi25 - fi25), (-(fi25)), (fi25 / ((42) | 1))}
                var v27 int8 = 127
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            if ((!false) || (65535 != 0)) {
                if true {
                    switch ("hello" + (("hello" + ("foo bar" + "foo bar")) + ("a" + "hello"))) {
                    case "hello":
                        v28 := GenS2{F0: -1, F1: 128, F2: 231}
                        _ = v28
                    }
                } else {
                    v29 := []int{}
                    delete(v4, -9223372036854775808)
                    _ = v29
                }
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 uint64 = 1
                    var v32 uint32 = 4294967295
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 int64 = -212
                _ = v33
            }
        }
    case 42:
        fmt.Println(v1, v6)
    default:
        v34 := GenS1{F0: 49}
        _ = v34
    }
    v35 := GenS2{F0: -2147483648, F1: 128, F2: 1}
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            if false {
                switch fi36 {
                case -247:
                    var v38 float64 = 1.5
                    _ = v38
                default:
                    switch (fi37 << ((1) & 63)) {
                    case 128:
                        v5[32767] = 2147483648
                    default:
                        var v39 uint = 571
                        _ = v39
                    }
                }
            } else {
                v1 = 1.5
            }
            var v40 uint32 = 1
            v40 = v40
            _ = fi37
            _ = v40
        }
        _ = fi36
    }
    v41 := GenS2{F0: -1, F1: 0, F2: int64(v6)}
    switch "hello" {
    case "hello":
        for fi42 := 0; fi42 < 9; fi42++ {
            switch "foo bar" {
            case "foo bar":
                v1 = (float32(0.1) * 1.5)
            case "a":
                fmt.Println(v6)
            }
            for fi43 := 0; fi43 < 8; fi43++ {
                fmt.Println(fi42, v1)
                var v44 int32 = 727
                v45 := []int{fi42, 4294967295, (fi42 | fi42)}
                fmt.Println(v6, v44, v1)
                var v46 int16 = 1
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            fmt.Println(fi42)
            _ = fi42
        }
    case "foo bar":
        v47 := GenS1{F0: uint64(v6)}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
