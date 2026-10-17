package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 string }

func main() {
    var v1 string = ""
    var v2 uint32 = 1
    var v3 string = "\n\t"
    var v4 float32 = -2.25
    var v5 int64 = 42
    v6 := []string{v3, v3, "\n\t", "\n\t"}
    if (675 != 0) {
        switch 42 {
        case 127:
            if false {
                var v7 int16 = 1
                var v8 uint64 = 4294967295
                var v9 int16 = (v7 << ((897) & 63))
                _ = v7
                _ = v8
                _ = v9
            }
        case 2147483647:
            var v10 int = 32767
            _ = v10
        case 32767:
            var v11 float64 = (-2.25 - 0.0)
            _ = v11
        default:
            var v12 uint = 2
            _ = v12
        }
        var v13 uint64 = 4294967295
        var v14 int = -9223372036854775808
        _ = v13
        _ = v14
    } else {
        if ((-2147483648 < -415) && (!(100 <= -102))) {
            if true {
                fmt.Println(v1, v2)
            } else {
                var v15 uint16 = 255
                _ = v15
            }
        } else {
            _ = len(v6)
            if (false || false) {
                v2 += (v2 - (v2 & (v2 | v2)))
                var v16 bool = (float64((float32(0.1) * float32(0.1))) == 1.5)
                v6 = append(v6, v3)
                _ = v16
            }
            var v17 int = 4294967295
            _ = v17
        }
        if len(v6) > 0 {
            v6[0] = v1
        }
    }
    fmt.Println(v1)
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    if (v2 != 42) {
        fmt.Println(v4, v1, v3)
    }
    fmt.Println(v3, v2)
    if true {
        fmt.Println(v3, v3)
        for fi18 := 0; fi18 < 15; fi18++ {
            if false {
                var v19 int64 = (v5 ^ -925)
                _ = v19
            } else {
                if ((!true) && (37 == 255)) {
                    fmt.Println(v2, v1, v4)
                    v20 := GenS1{F0: -1, F1: 2, F2: -7}
                    _ = v20
                }
                var v21 uint32 = 42
                _ = v21
            }
            fmt.Println(v4, v1, v1)
            _ = fi18
        }
    }
    v2 = v2
    fmt.Println(v3)
    var v22 uint64 = 441
    switch -9223372036854775808 {
    case 1:
        v22 = (v22 - v22)
    default:
        v3 = v1
    }
    if false {
        var v23 uint32 = 2
        v23 = (v23 % (((v23 >> ((1) & 63))) | 1))
        _ = v23
    } else {
        if (true && true) {
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS2{F0: v5, F1: -91, F2: 795}
                _ = fi24
                _ = v25
            }
        } else {
            v22 -= (v22 % ((513) | 1))
            switch (("foo bar" + ("a" + ((v3 + v1) + v1))) + v1) {
            case "foo bar":
                fmt.Println(v5, v5)
            case "hello":
                var v26 int = -2147483648
                _ = v26
            case "a":
                v27 := GenS2{F0: v5, F1: 100, F2: 1000}
                _ = v27
            }
        }
    }
    switch 0 {
    case 105:
        fmt.Println(v1, v4, v2)
    case -890:
        if true {
            var v28 uint32 = 655
            _ = v28
        }
    default:
        if (false && true) {
            if true {
                if ((v2 - (v2 * 631)) <= 2) {
                    if (0 >= -32768) {
                        v29 := GenS3{F0: 1}
                        var v30 int16 = int16(v5)
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 int32 = 2147483647
                            v33 := GenS4{F0: 255, F1: v1}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(v22, v4)
                    v34 := GenS1{F0: -348, F1: 2, F2: -128}
                    _ = v34
                }
                var v35 float32 = v4
                v22 = v22
                _ = v35
            } else {
                v36 := GenS2{F0: 502, F1: 20, F2: 42}
                var v37 float64 = (1.5 * -2.25)
                _ = v36
                _ = v37
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
