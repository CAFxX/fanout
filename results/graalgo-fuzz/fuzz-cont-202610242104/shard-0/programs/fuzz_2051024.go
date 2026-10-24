package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    var v1 int16 = 1
    var v2 int64 = 698
    v3 := []uint32{2, 4294967295, 42, 510}
    var v4 string = ""
    fmt.Println(v1, v2, v4)
    var v5 int64 = -95
    if true {
        var v6 int64 = v2
        _ = v6
    }
    var v7 uint64 = 2
    var v8 int8 = -29
    for fi9 := 0; fi9 < 4; fi9++ {
        if (v1 <= (-(v1))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v10 float64 = 0.0
            for fi11 := 0; fi11 < 7; fi11++ {
                if (3.14159 < float64(0.1)) {
                    fmt.Println(v5)
                }
                var v12 int8 = (v8 >> ((18446744073709551615) & 63))
                fmt.Println(fi9)
                fmt.Println(v5)
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        _ = fi9
    }
    var v13 int64 = v2
    v14 := GenS1{F0: v4}
    v7 = (v7 + 234)
    switch -9223372036854775808 {
    case -324:
        v7 -= 2
    case 1:
        if (-1 <= 149) {
            var v15 float64 = 1e10
            _ = len(v3)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(fi16, v5, v13)
                _ = fi16
            }
        }
    }
    if false {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 string = v4
            fmt.Println(fi17, v13)
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if true {
            var v20 uint16 = 2
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int16 = 0
                _ = len(v3)
                v23 := GenS2{F0: v13, F1: (!false)}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        }
        _ = fi19
    }
    var v24 bool = ((v7 ^ (v7 / ((2) | 1))) != 808)
    switch 127 {
    case -927:
        var v25 float64 = (1.5 - float64(0.1))
        _ = v25
    case -537:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32767:
        var v26 int8 = v8
        _ = v26
    default:
        if v24 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                if true {
                    v13 -= -816
                }
                v7 = v7
                _ = fi27
            }
            v28 := GenS3{F0: (^(v8)), F1: 1}
            v2 -= 9223372036854775807
            _ = v28
        }
    }
    v4 = ""
    v3 = append(v3, 2147483648)
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 int32 = 535
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        if (128 > 1) {
            for fi32 := 0; fi32 < 10; fi32++ {
                v33 := GenS3{F0: 0, F1: 590}
                v34 := GenS3{F0: -10, F1: 18446744073709551615}
                var v35 int32 = -216
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            var v36 uint16 = 65535
            _ = v36
        } else {
            var v37 int = (fi31 | -845)
            _ = v37
        }
        if len(v3) > 0 {
            v3[0] = 492
        }
        v2 = 466
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
