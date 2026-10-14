package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 0
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 uint32 = 766
    v6 := []string{("" + "foo bar")}
    v5 = v5
    var v7 int = 58
    var v8 float32 = -2.25
    if true {
        switch ("" + "foo bar") {
        case "hello":
            fmt.Println(v5, v2, v5)
        }
    } else {
        v2 -= (v2 % ((v2) | 1))
        v2 = (v2 & (^(v2)))
    }
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint16 = 32767
            switch (("hello" + "hello") + "a") {
            case "":
                for fi11 := 0; fi11 < 15; fi11++ {
                    v7 += (fi11 << ((597) & 63))
                    _ = fi11
                }
            case "foo bar":
                var v12 bool = (true || true)
                _ = v12
            case "\n\t":
                v5 -= (^(v5))
            }
            _ = fi9
            _ = v10
        }
        v5 = 1
        var v13 uint16 = 32767
        _ = v13
    } else {
        v14 := GenS1{F0: 42, F1: 1, F2: (v7 * (v7 & v7))}
        v15 := GenS1{F0: -128, F1: 42, F2: v7}
        _ = v14
        _ = v15
    }
    v5 -= (v5 * (^(v5)))
    _ = len(v6)
    var v16 bool = (true || true)
    var v17 bool = v16
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(v16)
        v16 = v17
        _ = fi18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            if len(v6) > 0 {
                v6[0] = "foo bar"
            }
            _ = fi20
        }
        switch "hello" {
        case "hello":
            v21 := GenS1{F0: 86, F1: 2, F2: fi19}
            _ = v21
        case "":
            fmt.Println(v2, v16)
        case "\n\t":
            v8 = 1.5
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        if (("foo bar" != "hello") || v16) {
            var v23 int16 = 1
            switch (v7 % ((fi22) | 1)) {
            case -9223372036854775808:
                v24 := GenS2{F0: 94}
                _ = v24
            case 0:
                v2 = (v2 * v2)
            }
            v7 -= v7
            _ = v23
        }
        v25 := GenS3{F0: 277, F1: -10}
        v26 := GenS4{F0: (v2 ^ 14)}
        _ = fi22
        _ = v25
        _ = v26
    }
    v2 = (v2 - (v2 % ((v2) | 1)))
    if (418 > uint16(v5)) {
        if (0 <= 18446744073709551615) {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int8 = 1
                v29 := GenS4{F0: (v2 - (^(v2)))}
                v8 -= (1e10 + (((v8 / v8) / v8) / float32(((0.1 + 1.5) + (((1e10 + 3.14159) + float64(v8)) - (float64(1e10) / 1.5))))))
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := GenS1{F0: -1, F1: 2, F2: fi31}
                fmt.Println(v2, fi31)
                var v33 uint8 = 0
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: (v5 >> (((v2 - (v2 ^ 599))) & 63))}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi30
            _ = v34
        }
        var v35 float32 = float32(-2.25)
        _ = v35
    } else {
        if true {
            for fi36 := 0; fi36 < 2; fi36++ {
                if v17 {
                    v17 = true
                }
                switch int(v2) {
                case 450:
                    for fi37 := 0; fi37 < 12; fi37++ {
                        var v38 int32 = 0
                        v39 := GenS4{F0: uint(v5)}
                        v17 = false
                        v7 = fi36
                        var v40 uint = v2
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                case 42:
                    var v41 int = 2147483647
                    _ = v41
                case 128:
                    v17 = (100 > 42)
                default:
                    var v42 uint8 = 174
                    _ = v42
                }
                var v43 float32 = (v8 + 1e10)
                _ = fi36
                _ = v43
            }
            v44 := GenS2{F0: v5}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v44
        } else {
            for fi45 := 0; fi45 < 11; fi45++ {
                for fi46 := 0; fi46 < 10; fi46++ {
                    v47 := GenS5{F0: 65535, F1: (v5 ^ 65535), F2: fi46}
                    fmt.Println(v7)
                    _ = fi46
                    _ = v47
                }
                _ = fi45
            }
        }
        switch ((("a" + "") + ("" + "hello")) + "foo bar") {
        case "\n\t":
            fmt.Println(v17, v5, v5)
        default:
            var v48 int = (^(v7))
            _ = v48
        }
        fmt.Println(v8, v17)
    }
    fmt.Println(v8, v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
