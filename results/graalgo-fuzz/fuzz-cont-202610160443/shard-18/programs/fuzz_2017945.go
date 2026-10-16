package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint64 }

func main() {
    var v1 int16 = 1
    var v2 float64 = float64(0.1)
    var v3 bool = true
    var v4 int8 = 42
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(v2, fi5)
        if (!true) {
            var v6 uint = 255
            v7 := GenS1{F0: 1, F1: 0, F2: v3}
            var v8 uint64 = 4294967295
            _ = v6
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 128, F1: 18446744073709551615, F2: true}
            v3 = (v3 && v3)
            switch 128 {
            case 42:
                v10 := []int{(fi5 ^ fi5)}
                _ = v10
            case -9223372036854775808:
                var v11 float64 = 1.5
                _ = v11
            default:
                fmt.Println(v2, fi5)
            }
            _ = v9
        }
        var v12 uint8 = 2
        _ = fi5
        _ = v12
    }
    var v13 int64 = 1
    var v14 int16 = (v1 >> ((958) & 63))
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS2{F0: 2147483647, F1: 1}
            var v18 int8 = -128
            _ = fi16
            _ = v17
            _ = v18
        }
        v14 = (v1 / ((1000) | 1))
        var v19 uint64 = 18446744073709551615
        _ = fi15
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 int32 = 42
        _ = fi20
        _ = v21
    }
    v3 = ((v14 & 0) == v1)
    v4 += (v4 | v4)
    var v22 int8 = (v4 + 42)
    if (!true) {
        if (false || false) {
            fmt.Println(v1, v22)
            for fi23 := 0; fi23 < 11; fi23++ {
                switch fi23 {
                case 4294967295:
                    v24 := GenS3{F0: 255, F1: fi23, F2: 65535}
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v26 := GenS4{F0: 73, F1: 315, F2: 1}
                        var v27 string = "a"
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 int = fi23
                var v29 int = fi23
                _ = fi23
                _ = v28
                _ = v29
            }
            fmt.Println(v2)
        }
        switch -758 {
        case 127:
            if (v3 && (!true)) {
                v2 = float64(0.0)
                var v30 int = -769
                _ = v30
            }
        case -109:
            v31 := []int{0, -2147483648}
            _ = v31
        case 128:
            switch -454 {
            case 9223372036854775807:
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 string = "hello"
                    v34 := GenS3{F0: 1, F1: -9223372036854775808, F2: 2}
                    if (-149 <= (^(v13))) {
                        if true {
                            for fi35 := 0; fi35 < 14; fi35++ {
                                v4 = 0
                                _ = fi35
                            }
                            fmt.Println(v14)
                            var v36 int = (fi32 << ((65535) & 63))
                            _ = v36
                        } else {
                            var v37 float32 = (float32(3.14159) - 1e10)
                            fmt.Println(v13)
                            _ = v37
                        }
                        v22 -= int8(fi32)
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                switch "foo bar" {
                                case "\n\t":
                                    v38 := []int{fi32, (fi32 | 9223372036854775807), (fi32 ^ fi32)}
                                    _ = v38
                                case "a":
                                    v1 = v1
                                case "foo bar":
                                    v39 := []int{fi32, (fi32 << ((1) & 63)), -9223372036854775808}
                                    _ = v39
                                }
                            case "":
                                var v40 int32 = 1
                                _ = v40
                            default:
                                if (42 <= (-(v4))) {
                                    v41 := GenS3{F0: 1, F1: fi32, F2: 935}
                                    fmt.Println(v4)
                                    v42 := []int{fi32}
                                    _ = v41
                                    _ = v42
                                } else {
                                    v1 += v14
                                    v14 = (^(v1))
                                }
                            }
                            v43 := GenS5{F0: 199, F1: 128, F2: 1}
                            v4 -= (v22 ^ (-(v22)))
                            _ = v43
                        }
                    } else {
                        if v3 {
                            v44 := []int{(fi32 & (fi32 - 886)), fi32}
                            _ = v44
                        } else {
                            v1 += (v1 * (v1 - int16(v4)))
                        }
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            default:
                v3 = (9223372036854775807 <= -2147483648)
            }
        }
        fmt.Println(v1, v2, v13)
    }
    v14 += -515
    var v45 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
}
