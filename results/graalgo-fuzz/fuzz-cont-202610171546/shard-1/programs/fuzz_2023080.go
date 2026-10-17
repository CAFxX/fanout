package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 float32 = 3.14159
    var v3 int64 = 9223372036854775807
    v1 += 1
    v4 := GenS1{F0: -881}
    if (1e10 != v2) {
        if (!false) {
            v5 := GenS1{F0: -1}
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := GenS2{F0: -128}
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        v1 -= (v1 << ((373) & 63))
        v8 := []int{}
        _ = v8
    } else {
        var v9 float64 = (1.5 / -2.25)
        switch ("a" + "foo bar") {
        case "\n\t":
            if (42 < 268) {
                if (v9 != v9) {
                    v9 += (1.5 / 3.14159)
                    var v10 uint8 = 156
                    _ = v10
                } else {
                    v11 := GenS3{F0: 42, F1: 1}
                    v12 := GenS2{F0: 28}
                    v1 -= (^(v1))
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 42}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 uint64 = 1
                    fmt.Println(v9, v3, fi14)
                    _ = fi14
                    _ = v15
                }
                v9 += float64(v2)
                v3 -= v3
            }
        case "":
            var v16 bool = true
            _ = v16
        case "a":
            if true {
                var v17 uint8 = 0
                _ = v17
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := []int{fi19}
                        var v21 float64 = ((((v9 * 3.14159) / -2.25) - float64(v2)) - ((1e10 + float64(v2)) - v9))
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v9 += v9
                    _ = fi18
                }
                fmt.Println(v1, v3, v2)
                var v22 int16 = 32767
                _ = v22
            }
        default:
            v23 := []int{357, 0, 127, 42}
            _ = v23
        }
        _ = v9
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 int32 = -697
        fmt.Println(v1, v25)
        fmt.Println(v2)
        _ = fi24
        _ = v25
    }
    if (!true) {
        switch int(v3) {
        case -9223372036854775808:
            v26 := GenS2{F0: -78}
            _ = v26
        case 127:
            var v27 int32 = -832
            _ = v27
        case 32767:
            for fi28 := 0; fi28 < 4; fi28++ {
                v3 += v3
                _ = fi28
            }
        default:
            var v29 float32 = ((0.1 / (0.0 + 1.5)) * (v2 / (-2.25 * (v2 - (v2 * v2)))))
            _ = v29
        }
        var v30 int64 = (v3 ^ 1)
        _ = v30
    }
    if (v3 <= v3) {
        if true {
            v1 = v1
        } else {
            v3 -= (v3 % ((v3) | 1))
            v31 := []int{}
            var v32 int64 = v3
            _ = v31
            _ = v32
        }
        v33 := GenS2{F0: -128}
        _ = v33
    } else {
        v1 += v1
        var v34 bool = (true || true)
        _ = v34
    }
    var v35 uint16 = v1
    for fi36 := 0; fi36 < 14; fi36++ {
        v3 += (v3 % (((^(v3))) | 1))
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
