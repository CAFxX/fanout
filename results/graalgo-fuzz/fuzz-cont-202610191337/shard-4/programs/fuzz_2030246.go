package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
func genFn0(p0 int16, p1 uint16, p2 uint16) uint32 {
    var v6 int64 = 42
    var v7 uint8 = 255
    v8 := []int{-891, 32768, -1}
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 -= -32768
        _ = fi9
    }
    var v10 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 uint) int32 {
    p0 = 42
    fmt.Println(p0)
    switch "hello" {
    case "hello":
        fmt.Println(p0)
    }
    p0 = (^(p0))
    _ = p0
    return -422
}

func main() {
    var v1 int64 = 1
    var v2 string = (("\n\t" + ("foo bar" + "hello")) + ("" + "a"))
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v5 int = -737
    for fi11 := 0; fi11 < 10; fi11++ {
        _ = len(v3)
        v1 -= (-(v1))
        _ = fi11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: -32768, F1: 1, F2: 1}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(v5, v5)
            _ = fi15
        }
        if len(v3) > 0 {
            v3[0] = 65535
        }
        _ = fi14
    }
    v16 := v4[v5]
    var v17 float32 = -2.25
    v1 = v1
    fmt.Println(v17, v16, v2)
    switch (-(v5)) {
    case -9223372036854775808:
        switch v5 {
        case 32768:
            fmt.Println(v17, v17, v5)
        case 266:
            var v18 uint16 = 255
            _ = v18
        case -382:
            v2 = v2
        }
    case -1:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v5 -= fi19
                var v21 bool = false
                var v22 int8 = 127
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    default:
        _ = len(v4)
    }
    var v23 bool = v16
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v16, fi24, v2)
        fmt.Println(v16, v17)
        _ = fi24
    }
    var v25 float64 = float64(0.0)
    switch v5 {
    case 32767:
        v26 := GenS1{F0: 32767, F1: -671, F2: 255}
        _ = v26
    default:
        var v27 string = ""
        _ = v27
    }
    if v23 {
        switch 42 {
        case 2147483647:
            var v28 int = (v5 ^ v5)
            _ = v28
        }
        if (!v16) {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(v17)
                _ = fi29
            }
        } else {
            var v30 string = "hello"
            if v16 {
                switch ("" + v2) {
                case "a":
                    if v16 {
                        fmt.Println(v17)
                        v31 := GenS1{F0: -32768, F1: (v1 << ((891) & 63)), F2: 0}
                        _ = len(v4)
                        _ = v31
                    }
                default:
                    v5 += (v5 << ((2) & 63))
                }
                v1 -= (v1 + (v1 + 9223372036854775807))
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 bool = (1 >= -1)
                    switch (v5 + 9223372036854775807) {
                    case -507:
                        v17 -= 0.1
                    case 4294967295:
                        fmt.Println(v30)
                    case 2147483647:
                        for fi34 := 0; fi34 < 13; fi34++ {
                            var v35 uint16 = 551
                            _ = fi34
                            _ = v35
                        }
                    default:
                        v16 = (!true)
                    }
                    switch v2 {
                    case "foo bar":
                        switch (fi32 << ((684) & 63)) {
                        case 127:
                            switch v5 {
                            case -2147483648:
                                var v36 float32 = (v17 - v17)
                                _ = v36
                            }
                        }
                    case "":
                        v25 += (v25 * (-2.25 / float64(v17)))
                    case "hello":
                        v37 := GenS2{F0: 0, F1: true, F2: 100}
                        _ = v37
                    default:
                        var v38 float32 = 1e10
                        _ = v38
                    }
                    _ = fi32
                    _ = v33
                }
                var v39 uint32 = 122
                _ = v39
            }
            if (!v23) {
                var v40 int = (v5 - (v5 - v5))
                _ = v40
            } else {
                if (255 != 42) {
                    if v16 {
                        var v41 int64 = v1
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v25 -= 0.0
                        v23 = v16
                        _ = fi42
                    }
                    _ = len(v4)
                }
            }
            _ = v30
        }
        v3 = append(v3, 2)
    }
    switch v5 {
    case 32768:
        _ = len(v4)
    case -9223372036854775808:
        delete(v4, v5)
    case -793:
        v43 := GenS2{F0: genFn1(65535), F1: true, F2: 42}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
