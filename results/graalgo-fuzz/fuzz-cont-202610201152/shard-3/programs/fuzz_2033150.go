package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int16; F2 int16 }

func main() {
    v1 := []int{32768, 128, 128, -9223372036854775808}
    var v2 uint8 = 1
    var v3 uint32 = 2147483647
    var v4 string = "a"
    var v5 uint16 = 275
    var v6 int16 = 42
    if (true || false) {
        fmt.Println(v2)
    } else {
        v7 := GenS1{F0: 32767, F1: 9223372036854775807}
        if false {
            v8 := GenS1{F0: -155, F1: -1}
            v6 += (v6 ^ (v6 ^ v6))
            _ = v8
        } else {
            switch "" {
            case "\n\t":
                _ = len(v1)
            case "foo bar":
                for fi9 := 0; fi9 < 9; fi9++ {
                    v6 = (v6 >> ((4294967295) & 63))
                    _ = fi9
                }
            }
        }
        _ = v7
    }
    if (589 == 0) {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 bool = ((v3 % (((v3 % ((896) | 1))) | 1)) < v3)
            v1 = append(v1, 9223372036854775807)
            _ = fi10
            _ = v11
        }
        fmt.Println(v4, v3, v4)
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            if (uint32(v2) != (v3 + 2147483648)) {
                v2 -= (v2 << ((1) & 63))
                v13 := GenS2{F0: 42}
                _ = v13
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            _ = fi12
        }
        switch -613 {
        case 4294967295:
            v14 := GenS1{F0: 0, F1: 42}
            _ = v14
        case 145:
            if len(v1) > 0 {
                v1[0] = -2147483648
            }
        case -479:
            var v15 uint64 = 0
            _ = v15
        default:
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 string = (("a" + "foo bar") + (((v4 + "hello") + v4) + v4))
                _ = fi16
                _ = v17
            }
        }
        if false {
            v18 := GenS1{F0: 982, F1: 0}
            for fi19 := 0; fi19 < 12; fi19++ {
                v3 = 149
                v2 -= (v2 & uint8(v5))
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := GenS3{F0: v4}
                    var v22 float32 = 1e10
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi19
            }
            _ = v18
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v1)
        v6 = (v6 | (v6 % ((19) | 1)))
        var v24 string = (v4 + v4)
        _ = fi23
        _ = v24
    }
    var v25 string = v4
    if len(v1) > 0 {
        v1[0] = 173
    }
    v5 -= v5
    if ((v5 > v5) || false) {
        v26 := GenS4{F0: (v3 / ((v3) | 1)), F1: int16(v2), F2: -32768}
        _ = v26
    }
    var v27 string = "foo bar"
    v27 = v27
    if true {
        v2 += (v2 + v2)
        var v28 int64 = 640
        _ = v28
    } else {
        v1 = append(v1, 32767)
        switch int(v6) {
        case 0:
            fmt.Println(v4, v4)
        }
    }
    v1 = append(v1, -334)
    for fi29 := 0; fi29 < 17; fi29++ {
        v1 = append(v1, fi29)
        v25 = "\n\t"
        if len(v1) > 0 {
            v1[0] = (^(fi29))
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
