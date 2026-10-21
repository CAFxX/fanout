package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }
func genFn0(p0 int64, p1 int32) int8 {
    v6 := []int{160, int(p0), -1, -687}
    var v7 int32 = (p1 ^ -2147483648)
    switch 0 {
    case 128:
        var v8 int16 = 1000
        _ = v8
    case 32768:
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 string = "a"
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    v7 = 1
    for fi12 := 0; fi12 < 6; fi12++ {
        _ = len(v6)
        if (true || (true || true)) {
            p0 = (p0 ^ (p0 % ((p0) | 1)))
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 uint32) uint16 {
    fmt.Println(p0, p0)
    var v13 uint8 = 1
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    v3 := []int{}
    v4 := make(map[int]string)
    var v5 int = -792
    switch "" {
    case "\n\t":
        var v14 uint32 = 972
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            v5 += (-(fi15))
            var v16 int32 = 0
            _ = fi15
            _ = v16
        }
    default:
        var v17 uint = 65535
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v2 = v2
        _ = fi18
    }
    if (genFn1(346, 2147483647) <= genFn1(255, 2147483648)) {
        switch (^(v5)) {
        case 2147483647:
            if true {
                var v19 int8 = (v2 >> ((18446744073709551615) & 63))
                var v20 uint = 42
                _ = v19
                _ = v20
            }
        case 453:
            var v21 uint16 = 255
            _ = v21
        }
        switch -1 {
        case -917:
            v22 := GenS1{F0: 1}
            _ = v22
        case -2147483648:
            fmt.Println(v5)
        case 20:
            for fi23 := 0; fi23 < 7; fi23++ {
                for fi24 := 0; fi24 < 2; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v25 int16 = 836
                    v25 = (v25 ^ v25)
                    v26 := GenS2{F0: 42, F1: 194}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        default:
            v5 += v5
        }
    }
    v27 := GenS2{F0: 1, F1: 32767}
    if (-9223372036854775808 >= 1) {
        v5 = -947
    } else {
        fmt.Println(v5, v2)
    }
    v28 := GenS3{F0: -455, F1: ((1000 > -408) && false)}
    v2 += -128
    v29 := GenS1{F0: 128}
    if (false && false) {
        if false {
            if false {
                var v30 float32 = 0.0
                _ = v30
            }
            var v31 int16 = 99
            var v32 int8 = -1
            _ = v31
            _ = v32
        } else {
            v2 += genFn0(372, 42)
            fmt.Println(v2, v5)
            if false {
                var v33 uint = 1
                var v34 uint32 = 65535
                _ = v33
                _ = v34
            }
        }
        var v35 float32 = 1e10
        v36 := GenS4{F0: genFn1(574, 2147483647), F1: 2147483647, F2: false}
        _ = v35
        _ = v36
    } else {
        fmt.Println(v2, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
