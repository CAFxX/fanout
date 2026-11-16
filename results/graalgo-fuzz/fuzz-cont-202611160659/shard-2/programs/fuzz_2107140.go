package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint64; F2 int }
func genFn0(p0 int16, p1 int8) bool {
    fmt.Println(p0)
    v7 := GenS1{F0: 42, F1: 65535}
    fmt.Println(p1, p0)
    p1 += (p1 + (^(p1)))
    _ = p0
    _ = p1
    _ = v7
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    var v3 string = ("a" + "")
    v4 := []string{"", "hello", (v3 + "foo bar"), "a"}
    var v5 uint8 = 0
    var v6 float32 = (1e10 - 1e10)
    var v8 int8 = -128
    var v9 uint32 = 2
    if (!false) {
        v2 = v2
    }
    if false {
        switch 879 {
        case -305:
            for fi10 := 0; fi10 < 1; fi10++ {
                if true {
                    if false {
                        v3 = (v3 + "\n\t")
                        var v11 uint32 = v9
                        if false {
                            delete(v1, 1)
                        }
                        _ = v11
                    } else {
                        v12 := GenS2{F0: 636}
                        _ = v12
                    }
                    fmt.Println(v9, v3)
                } else {
                    fmt.Println(v5, fi10)
                    v13 := GenS2{F0: 1000}
                    v6 -= (v6 / 1.5)
                    _ = v13
                }
                fmt.Println(v6)
                _ = fi10
            }
        case -2147483648:
            var v14 int32 = v2
            _ = v14
        default:
            fmt.Println(v3, v2, v5)
        }
        v15 := GenS3{F0: (v2 >> ((2) & 63)), F1: (v8 - (v8 & (v8 | (v8 & v8)))), F2: 1000}
        if ((v5 * v5) >= 206) {
            for fi16 := 0; fi16 < 13; fi16++ {
                _ = len(v1)
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v17 := GenS4{F0: (v3 + v3)}
                _ = fi16
                _ = v17
            }
            v18 := v1[161]
            _ = v18
        }
        _ = v15
    } else {
        var v19 uint32 = (v9 / (((v9 + v9)) | 1))
        switch "" {
        case "hello":
            v20 := v1[1]
            _ = v20
        case "foo bar":
            var v21 int32 = 462
            _ = v21
        case "":
            if ((255 == 4294967295) && true) {
                v6 += float32((0.0 * 0.1))
                var v22 int32 = v2
                _ = v22
            }
        }
        _ = v19
    }
    fmt.Println(v3)
    v5 += v5
    v23 := GenS2{F0: -1}
    var v24 int32 = v2
    var v25 float32 = 1.5
    if (v24 != (v2 << ((646) & 63))) {
        switch 207 {
        case 551:
            v26 := GenS5{F0: 429, F1: 0, F2: 9223372036854775807}
            _ = v26
        case 2147483647:
            v6 -= v6
        }
        v8 = -1
        fmt.Println(v3, v3)
    }
    var v27 int8 = v8
    _ = len(v4)
    if true {
        v2 += (v24 & -2147483648)
    }
    fmt.Println(v9)
    for fi28 := 0; fi28 < 20; fi28++ {
        if true {
            var v29 float64 = (float64((v6 + 1.5)) + (3.14159 / 1.0))
            _ = v29
        } else {
            v30 := GenS3{F0: 1, F1: (v27 ^ (v8 - v27)), F2: 1}
            _ = v30
        }
        v27 -= (v27 >> ((65535) & 63))
        _ = fi28
    }
    fmt.Println(v6)
    switch "hello" {
    case "hello":
        switch -492 {
        case 2147483647:
            for fi31 := 0; fi31 < 20; fi31++ {
                v24 = -2147483648
                v32 := GenS3{F0: (^(v24)), F1: (v8 / ((-91) | 1)), F2: 925}
                switch (fi31 * (fi31 ^ fi31)) {
                case 9223372036854775807:
                    var v33 int8 = v27
                    _ = v33
                case 233:
                    var v34 int8 = v8
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
        case 32767:
            switch 2147483647 {
            case 128:
                v3 = "a"
            case 628:
                var v35 int = 823
                _ = v35
            case -1:
                switch v3 {
                case "\n\t":
                    v5 += v5
                case "":
                    var v36 int8 = v8
                    _ = v36
                default:
                    delete(v1, 1)
                }
            default:
                v27 += (v27 * -128)
            }
        }
    case "\n\t":
        var v37 int8 = v27
        _ = v37
    default:
        if (true || false) {
            fmt.Println(v25, v8, v24)
            fmt.Println(v6, v2)
        } else {
            var v38 uint = 2
            for fi39 := 0; fi39 < 4; fi39++ {
                for fi40 := 0; fi40 < 6; fi40++ {
                    var v41 uint32 = (v9 % ((v9) | 1))
                    fmt.Println(v27, v9)
                    _ = fi40
                    _ = v41
                }
                _ = fi39
            }
            v5 -= v5
            _ = v38
        }
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        v1[(fi42 | fi42)] = (fi42 & fi42)
        v43 := v1[fi42]
        var v44 int16 = 32767
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
