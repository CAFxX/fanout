package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0() int16 {
    switch 4294967295 {
    case 42:
        v5 := []int{}
        _ = v5
    case 32767:
        var v6 int32 = -268
        _ = v6
    default:
        if (4294967295 <= -497) {
            var v7 uint = 1
            v8 := []int{-2147483648, 127}
            fmt.Println(v7)
            _ = v7
            _ = v8
        }
    }
    v9 := []int{}
    if (32768 <= -2147483648) {
        var v10 int16 = 42
        v10 = 1000
        _ = v10
    } else {
        switch "hello" {
        case "":
            for fi11 := 0; fi11 < 5; fi11++ {
                if false {
                    var v12 bool = true
                    _ = v12
                }
                fmt.Println(fi11)
                v13 := GenS1{F0: ""}
                _ = fi11
                _ = v13
            }
        case "foo bar":
            var v14 uint = 4294967295
            _ = v14
        case "a":
            var v15 int32 = 1
            _ = v15
        default:
            v16 := GenS1{F0: ""}
            _ = v16
        }
        if len(v9) > 0 {
            v9[0] = 1
        }
    }
    _ = v9
    return -32768
}

func main() {
    var v1 uint64 = 42
    var v2 int = 4294967295
    v3 := make(map[int]bool)
    var v4 bool = false
    var v17 int32 = 1
    _ = len(v3)
    v18 := []int{}
    if v4 {
        var v19 uint32 = 2147483647
        _ = v19
    } else {
        if v4 {
            switch "hello" {
            case "a":
                var v20 int = (v2 * v2)
                _ = v20
            case "foo bar":
                _ = len(v18)
            }
        } else {
            var v21 bool = ((1 < 255) && true)
            if (100 > 100) {
                v22 := v3[(v2 << ((715) & 63))]
                var v23 int32 = (v17 % (((v17 ^ (v17 / ((v17) | 1)))) | 1))
                if v21 {
                    var v24 int = (v2 << ((255) & 63))
                    _ = v24
                }
                _ = v22
                _ = v23
            } else {
                switch "foo bar" {
                case "hello":
                    v25 := v3[(v2 - (v2 + -120))]
                    _ = v25
                default:
                    if (!v4) {
                        var v26 int = (v2 + v2)
                        v21 = true
                        _ = v26
                    }
                }
                var v27 int64 = 0
                v2 -= 15
                _ = v27
            }
            _ = v21
        }
        if (!(3.14159 != (0.1 + 3.14159))) {
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 int16 = 686
                v4 = false
                v30 := v3[fi28]
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = len(v3)
        }
    }
    var v31 int = 42
    var v32 int = -2147483648
    v33 := GenS1{F0: "foo bar"}
    switch "" {
    case "hello":
        v2 -= v32
    }
    var v34 uint32 = 2147483648
    var v35 uint16 = 501
    v1 += 4294967295
    v34 -= (v34 * v34)
    v36 := GenS1{F0: ((("hello" + "a") + "foo bar") + "foo bar")}
    if ((92 == 69) || (!(3.14159 < 0.1))) {
        v35 = uint16(v34)
    } else {
        for fi37 := 0; fi37 < 18; fi37++ {
            v38 := GenS2{F0: v1, F1: ((v35 - v35) > 42)}
            fmt.Println(v1, fi37, fi37)
            v39 := GenS3{F0: (false || (0 <= 9223372036854775807))}
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    v40 := GenS4{F0: 749, F1: 127, F2: 42}
    fmt.Println(v2, v1, v17)
    switch -9223372036854775808 {
    case 0:
        for fi41 := 0; fi41 < 19; fi41++ {
            for fi42 := 0; fi42 < 10; fi42++ {
                v43 := GenS4{F0: 0, F1: 127, F2: 191}
                var v44 int64 = 0
                _ = fi42
                _ = v43
                _ = v44
            }
            if ((v1 & (v1 * (^(v1)))) <= (v1 >> ((255) & 63))) {
                v1 += 42
                var v45 int8 = int8(fi41)
                v46 := GenS4{F0: 42, F1: (v45 / (((v45 | (v45 & int8(v35)))) | 1)), F2: 89}
                _ = v45
                _ = v46
            }
            _ = fi41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
}
