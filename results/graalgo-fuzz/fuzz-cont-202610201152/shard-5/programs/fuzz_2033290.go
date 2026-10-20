package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint16 = 0
    v2 := []string{("\n\t" + "foo bar"), "", ""}
    v3 := make(map[int]bool)
    var v4 int32 = 1
    var v5 int64 = 9223372036854775807
    v6 := make(map[uint32]bool)
    var v7 bool = (42 != 0)
    if v7 {
        for fi8 := 0; fi8 < 5; fi8++ {
            v3[fi8] = (!(42 >= v1))
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint64 = 524
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    } else {
        fmt.Println(v7, v5, v7)
        var v11 int32 = v4
        v12 := GenS1{F0: (1 < 156), F1: 18446744073709551615, F2: v5}
        _ = v11
        _ = v12
    }
    v7 = (v7 && true)
    v4 = (v4 << ((42) & 63))
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v5 = (v5 | v5)
    v1 -= (^(v1))
    switch "foo bar" {
    case "a":
        v1 = (v1 - (v1 & 255))
    case "\n\t":
        for fi13 := 0; fi13 < 11; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int16 = 32767
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v16 := GenS2{F0: 1, F1: uint64(v1)}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    }
    if true {
        var v17 float32 = 1.5
        v7 = (42 < 2147483647)
        _ = v17
    } else {
        switch "a" {
        case "":
            switch 42 {
            case -437:
                v1 = (v1 * 1)
            case 32767:
                delete(v6, 65535)
            case 236:
                for fi18 := 0; fi18 < 13; fi18++ {
                    v5 += (v5 * (v5 + v5))
                    v19 := GenS1{F0: true, F1: 991, F2: v5}
                    _ = fi18
                    _ = v19
                }
            }
        }
        var v20 int16 = -792
        v21 := GenS2{F0: v1, F1: 0}
        _ = v20
        _ = v21
    }
    switch ("" + "foo bar") {
    case "":
        v7 = v7
    case "foo bar":
        v7 = (127 < -128)
    case "\n\t":
        v3[32768] = (128 >= 14)
    default:
        fmt.Println(v4, v4, v1)
    }
    switch (("foo bar" + "hello") + "a") {
    case "hello":
        v1 = v1
    case "":
        v1 = (v1 ^ 65535)
    default:
        var v22 float64 = 1.5
        _ = v22
    }
    v4 += (v4 + v4)
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint = 764
        _ = fi23
        _ = v24
    }
    if v7 {
        var v25 int8 = 21
        var v26 uint8 = 1
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int = fi27
            _ = fi27
            _ = v28
        }
        var v29 int32 = (v4 | (v4 / ((718) | 1)))
        v30 := v3[1]
        _ = v29
        _ = v30
    }
    var v31 uint = 658
    v32 := GenS1{F0: (1e10 < 0.1), F1: (v31 * uint(v1)), F2: -1}
    fmt.Println(v4)
    if v7 {
        switch 32768 {
        case 4294967295:
            if (v7 || v7) {
                var v33 int32 = 1
                fmt.Println(v4, v33, v5)
                v4 = v4
                _ = v33
            } else {
                if ((v1 | 2) != v1) {
                    v31 -= (v31 + (v31 * (^(v31))))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v34 uint8 = 248
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v7 = (true && (v7 && true))
                        _ = fi35
                    }
                    v36 := GenS2{F0: 2, F1: 42}
                    var v37 uint32 = 2
                    _ = v36
                    _ = v37
                }
            }
        }
    } else {
        fmt.Println(v5, v5)
        v31 += v31
        v38 := GenS2{F0: (^(v1)), F1: 1}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
