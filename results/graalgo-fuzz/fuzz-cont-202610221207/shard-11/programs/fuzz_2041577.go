package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]uint32)
    var v3 float32 = 3.14159
    switch -148 {
    case 4294967295:
        var v4 int32 = -2147483648
        _ = v4
    case -499:
        for fi5 := 0; fi5 < 1; fi5++ {
            for fi6 := 0; fi6 < 9; fi6++ {
                v1 = 798
                _ = fi6
            }
            _ = fi5
        }
    default:
        var v7 uint32 = (v1 ^ (v1 ^ (v1 & v1)))
        _ = v7
    }
    v8 := []int{-9223372036854775808, 422, -784}
    fmt.Println(v1)
    switch (("\n\t" + "") + ("a" + "foo bar")) {
    case "a":
        var v9 int8 = 127
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: -1}
        _ = v10
    default:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: -1}
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS2{F0: 705, F1: 42, F2: 32767}
    switch "\n\t" {
    case "":
        fmt.Println(v1, v1)
    case "\n\t":
        if false {
            var v14 uint16 = 2
            v14 += (v14 >> ((813) & 63))
            var v15 int32 = -1
            _ = v14
            _ = v15
        }
    default:
        fmt.Println(v1, v1)
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint8 = 42
        var v18 int8 = 100
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 13; fi19++ {
            if (!false) {
                v1 -= (v1 | v1)
            }
            _ = fi19
        }
        v1 = v1
        v20 := GenS3{F0: 547}
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := v2[(fi21 / ((2147483647) | 1))]
        v3 += (v3 + (v3 / 0.1))
        for fi23 := 0; fi23 < 16; fi23++ {
            v3 = float32(1.5)
            var v24 int32 = 0
            switch 2147483647 {
            case 127:
                v25 := v2[fi21]
                _ = v25
            case 577:
                fmt.Println(v3, v1)
            case 32767:
                var v26 int64 = -9223372036854775808
                _ = v26
            default:
                v8 = append(v8, 610)
            }
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    _ = len(v2)
    v1 = v1
    for fi27 := 0; fi27 < 3; fi27++ {
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                for fi28 := 0; fi28 < 2; fi28++ {
                    if (fi28 >= 4294967295) {
                        _ = len(v8)
                    }
                    if (-128 > 1) {
                        var v29 int8 = -1
                        if (true && (true && (!true))) {
                            var v30 string = "\n\t"
                            var v31 float64 = 1e10
                            _ = v30
                            _ = v31
                        }
                        var v32 float64 = (0.1 - (1.5 - (1e10 - 1.5)))
                        v2[fi28] = (^(v1))
                        _ = v29
                        _ = v32
                    } else {
                        fmt.Println(fi28, v1, v3)
                    }
                    _ = fi28
                }
            default:
                v1 = (^(v1))
            }
        default:
            fmt.Println(v3)
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi27
    }
    if false {
        fmt.Println(v1)
        v1 += (v1 * v1)
        switch "" {
        case "hello":
            var v33 float64 = 0.0
            _ = v33
        case "":
            v1 -= 642
        default:
            var v34 int16 = 1000
            _ = v34
        }
    }
    v35 := GenS3{F0: 4294967295}
    var v36 uint16 = 177
    v36 = v36
    for fi37 := 0; fi37 < 13; fi37++ {
        switch ("" + "") {
        case "hello":
            var v38 float32 = 1e10
            _ = v38
        case "foo bar":
            if false {
                v36 += v36
            }
        case "":
            v39 := GenS1{F0: 1000}
            _ = v39
        }
        var v40 float32 = (v3 / 1.0)
        _ = fi37
        _ = v40
    }
    fmt.Println(v36, v36, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v8
}
