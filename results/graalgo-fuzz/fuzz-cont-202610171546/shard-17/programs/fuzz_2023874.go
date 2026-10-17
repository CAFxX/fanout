package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 int8 }
type GenS4 struct { F0 int8; F1 int; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64(1e10)
    var v3 uint8 = 2
    v4 := make(map[int]uint32)
    var v5 float32 = ((1e10 / 1.0) + (1.5 - (0.1 + 3.14159)))
    v5 += ((v5 * v5) - (v5 / v5))
    v4[402] = 2147483647
    v6 := v4[675]
    switch "foo bar" {
    case "":
        if ((710 > 0) && true) {
            var v7 float64 = ((3.14159 + float64(((float32(float64(v5)) * 0.1) - v5))) - (v2 / v2))
            var v8 uint = 18446744073709551615
            _ = v7
            _ = v8
        } else {
            delete(v1, ("hello" + "\n\t"))
            v3 -= (v3 >> ((2) & 63))
        }
    default:
        fmt.Println(v5, v6, v5)
    }
    v9 := GenS1{F0: 95, F1: -1, F2: 1}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v3 = (v3 ^ v3)
            v12 := GenS1{F0: 32767, F1: 100, F2: 4294967295}
            var v13 uint = 0
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if (true && false) {
        v2 += v2
        var v14 int64 = -1
        _ = v14
    }
    var v15 bool = (false && (0 <= 0))
    var v16 uint = 0
    v15 = v15
    fmt.Println(v16)
    fmt.Println(v3)
    fmt.Println(v15, v3)
    v15 = false
    switch 193 {
    case 0:
        var v17 int16 = -1
        _ = v17
    case 1:
        if ((v6 * 768) != v6) {
            v18 := []int{127, -109, 9223372036854775807}
            _ = v18
        } else {
            switch 32768 {
            case 127:
                v16 = (v16 / ((v16) | 1))
            case 128:
                v19 := GenS2{F0: 426, F1: (v6 >> ((v16) & 63))}
                _ = v19
            case 505:
                if v15 {
                    fmt.Println(v3)
                    v20 := []int{0, -9223372036854775808, -1, 32767}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v22 := []int{(fi21 << ((1) & 63)), fi21, (^(fi21))}
                        _ = fi21
                        _ = v22
                    }
                    var v23 uint16 = 1
                    _ = v23
                }
            default:
                var v24 int32 = -1
                _ = v24
            }
        }
    case 32768:
        v25 := v4[42]
        _ = v25
    default:
        for fi26 := 0; fi26 < 19; fi26++ {
            if v15 {
                var v27 uint64 = 2
                v15 = ((v3 / ((v3) | 1)) != (^(v3)))
                _ = v27
            } else {
                for fi28 := 0; fi28 < 11; fi28++ {
                    if v15 {
                        switch (fi26 & (^(fi26))) {
                        case 2147483647:
                            var v29 bool = (int32(v3) >= -212)
                            _ = v29
                        default:
                            if v15 {
                                var v30 uint64 = 4294967295
                                v2 += ((-2.25 + v2) * v2)
                                var v31 string = "\n\t"
                                fmt.Println(v3)
                                _ = v30
                                _ = v31
                            }
                        }
                        v1[(("" + "\n\t") + "a")] = false
                    }
                    var v32 int = fi26
                    _ = fi28
                    _ = v32
                }
                switch fi26 {
                case 4294967295:
                    v16 = v16
                case 42:
                    for fi33 := 0; fi33 < 7; fi33++ {
                        switch "" {
                        case "foo bar":
                            var v34 int32 = 2147483647
                            _ = v34
                        case "\n\t":
                            if false {
                                v35 := GenS2{F0: (v6 * (v6 & v6)), F1: 2147483647}
                                fmt.Println(fi26)
                                v36 := GenS3{F0: 1, F1: ("foo bar" + ("hello" + "hello")), F2: -1}
                                _ = v35
                                _ = v36
                            }
                        case "a":
                            var v37 uint8 = (v3 * 128)
                            _ = v37
                        }
                        _ = fi33
                    }
                default:
                    if v15 {
                        for fi38 := 0; fi38 < 10; fi38++ {
                            delete(v1, "a")
                            v39 := GenS3{F0: int64(v16), F1: "hello", F2: -5}
                            fmt.Println(v5, v3)
                            _ = fi38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 16; fi40++ {
                            switch "" {
                            case "hello":
                                fmt.Println(fi26)
                            }
                            v15 = v15
                            _ = fi40
                        }
                    } else {
                        v1[("foo bar" + ("" + "hello"))] = (!v15)
                    }
                }
                var v41 uint32 = (v6 << (((v16 - 65535)) & 63))
                _ = v41
            }
            v42 := GenS4{F0: -69, F1: fi26, F2: "a"}
            switch (fi26 % ((fi26) | 1)) {
            case 2147483647:
                var v43 uint32 = v6
                _ = v43
            case -1:
                v6 -= (v6 >> ((1) & 63))
            default:
                _ = len(v1)
            }
            _ = fi26
            _ = v42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
