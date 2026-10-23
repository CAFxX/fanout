package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []int{1, 368}
    v2 := make(map[uint32]int)
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 0
    v5 := make(map[int]bool)
    fmt.Println(v4, v4)
    var v6 int64 = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            if false {
                v8 := GenS1{F0: (-2.25 >= 3.14159)}
                _ = v8
            } else {
                fmt.Println(v6)
            }
            switch (fi7 * fi7) {
            case 32767:
                v3 = 2
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(v3)
                v10 := GenS1{F0: (!false)}
                v11 := v5[(^(fi9))]
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi7
    }
    switch 32768 {
    case -1:
        var v12 int32 = -2147483648
        _ = v12
    case 179:
        var v13 string = ("foo bar" + ("a" + (("hello" + "\n\t") + "foo bar")))
        _ = v13
    case -2147483648:
        var v14 int16 = 32767
        _ = v14
    default:
        var v15 int8 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if len(v1) > 0 {
            v1[0] = fi16
        }
        _ = fi16
    }
    var v17 uint32 = 65535
    v6 = v6
    switch ("" + "") {
    case "foo bar":
        v18 := GenS2{F0: 1, F1: -1}
        _ = v18
    }
    v6 -= (^(v6))
    v3 = 0
    v3 = (^(v3))
    var v19 int8 = -1
    if false {
        fmt.Println(v4)
        if (("foo bar" + ("\n\t" + "")) >= "hello") {
            v3 -= 4294967295
            var v20 uint16 = uint16(v4)
            v6 += v6
            _ = v20
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            if false {
                v17 += (v17 * v17)
                switch "hello" {
                case "\n\t":
                    if true {
                        switch (fi21 << ((4294967295) & 63)) {
                        case -1:
                            fmt.Println(fi21, v19, v19)
                        default:
                            switch (fi21 % (((fi21 + (fi21 * 127))) | 1)) {
                            case 703:
                                var v22 bool = false
                                _ = v22
                            case 4294967295:
                                if false {
                                    v6 -= v6
                                    var v23 uint8 = 37
                                    var v24 float64 = 1e10
                                    _ = v23
                                    _ = v24
                                } else {
                                    fmt.Println(v6, fi21)
                                    v3 = v3
                                }
                            }
                        }
                    }
                case "":
                    v3 -= v3
                default:
                    switch "\n\t" {
                    case "\n\t":
                        var v25 uint32 = (v17 - v17)
                        _ = v25
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            } else {
                v17 = uint32(v4)
                v6 = (v6 | 643)
            }
            if (394 >= 44) {
                delete(v2, uint32(v3))
                v17 = (v17 + v17)
                fmt.Println(v4, v17)
            }
            for fi26 := 0; fi26 < 17; fi26++ {
                fmt.Println(v17, v6, fi21)
                var v27 int64 = int64(v3)
                if false {
                    v6 += v27
                    switch ("" + "\n\t") {
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = (-(fi21))
                        }
                    case "\n\t":
                        v28 := GenS2{F0: int16(v3), F1: (fi26 ^ (fi21 & -19))}
                        _ = v28
                    }
                }
                _ = fi26
                _ = v27
            }
            _ = fi21
        }
    } else {
        if false {
            v29 := GenS3{F0: 42, F1: (v4 * v4)}
            v30 := GenS1{F0: true}
            for fi31 := 0; fi31 < 13; fi31++ {
                v1 = append(v1, 128)
                _ = fi31
            }
            _ = v29
            _ = v30
        }
        var v32 uint16 = 393
        _ = v32
    }
    v4 += (v4 - 65)
    v17 += 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
