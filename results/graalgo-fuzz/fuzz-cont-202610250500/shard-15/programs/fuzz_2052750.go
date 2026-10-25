package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 int32 = -623
    var v2 uint32 = 0
    v3 := make(map[string]uint32)
    var v4 uint16 = 0
    switch "foo bar" {
    case "hello":
        fmt.Println(v4, v1)
    case "a":
        for fi5 := 0; fi5 < 3; fi5++ {
            if ("foo bar" <= ("foo bar" + "")) {
                for fi6 := 0; fi6 < 6; fi6++ {
                    v4 += (v4 * v4)
                    v2 += (v2 / ((858) | 1))
                    v1 = v1
                    _ = fi6
                }
                v2 += (v2 << ((42) & 63))
                fmt.Println(v1)
            } else {
                v7 := GenS1{F0: 0, F1: 128, F2: 127}
                v4 += v4
                v2 = (v2 / ((v2) | 1))
                _ = v7
            }
            if (((v2 != 2) && (32767 <= (^(v4)))) || (!false)) {
                var v8 uint = 18446744073709551615
                var v9 float64 = float64((0.0 + 1e10))
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint32 = v2
                _ = fi10
                _ = v11
            }
            _ = fi5
        }
    case "foo bar":
        for fi12 := 0; fi12 < 10; fi12++ {
            if ((1 != -9223372036854775808) && ("\n\t" < "a")) {
                var v13 float64 = 3.14159
                switch (fi12 >> ((0) & 63)) {
                case -752:
                    v14 := []int{fi12}
                    _ = v14
                case 32768:
                    fmt.Println(v2, v4, v1)
                default:
                    switch 127 {
                    case 484:
                        var v15 uint32 = v2
                        _ = v15
                    case -1:
                        for fi16 := 0; fi16 < 20; fi16++ {
                            if true {
                                var v17 uint16 = (v4 / ((v4) | 1))
                                if (!false) {
                                    v18 := GenS1{F0: 794, F1: 128, F2: -1}
                                    if false {
                                        fmt.Println(v4)
                                        v2 -= v2
                                    }
                                    if ((42 == 0) && (v2 > 827)) {
                                        v19 := []int{fi16}
                                        delete(v3, "")
                                        var v20 int64 = 0
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v18
                                } else {
                                    if true {
                                        v1 = v1
                                        v21 := []int{503, (fi16 + fi16), (-(fi12)), fi12}
                                        _ = v21
                                    }
                                    var v22 int = (fi16 ^ fi16)
                                    _ = v22
                                }
                                if (4294967295 != 4294967295) {
                                    var v23 bool = (0 > -1)
                                    _ = v23
                                } else {
                                    v24 := GenS1{F0: 2, F1: 69, F2: -128}
                                    fmt.Println(v4, v2, fi16)
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v25 uint = 262
                                        _ = v25
                                    case "hello":
                                        v1 += (v1 - v1)
                                    case "foo bar":
                                        v26 := GenS1{F0: 2, F1: 62, F2: int8(v4)}
                                        _ = v26
                                    }
                                    v27 := []int{(-(fi12)), -1, (fi12 ^ 42), (-(fi12))}
                                    _ = v24
                                    _ = v27
                                }
                                v28 := []int{(fi12 % ((fi12) | 1))}
                                fmt.Println(v4)
                                _ = v17
                                _ = v28
                            }
                            fmt.Println(v1)
                            v13 = (0.0 + float64((1e10 - ((0.0 + 1e10) * (3.14159 * 1e10)))))
                            _ = fi16
                        }
                    case 42:
                        fmt.Println(v4)
                    default:
                        var v29 int8 = 110
                        _ = v29
                    }
                }
                _ = v13
            }
            var v30 float64 = 1e10
            _ = fi12
            _ = v30
        }
    default:
        v31 := GenS2{F0: 114, F1: "", F2: 127}
        _ = v31
    }
    fmt.Println(v4, v2, v4)
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := []int{}
            v35 := GenS1{F0: 128, F1: 255, F2: int8(v1)}
            v36 := GenS3{F0: (-(fi33))}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi32
    }
    v2 -= v2
    v37 := GenS4{F0: 9223372036854775807, F1: (^(v4))}
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS5{F0: 9223372036854775807, F1: v1}
        var v40 int8 = 100
        fmt.Println(v4)
        _ = fi38
        _ = v39
        _ = v40
    }
    if (!true) {
        if false {
            v2 += (v2 % ((65535) | 1))
            v1 += -710
        }
        var v41 float64 = 3.14159
        var v42 string = "hello"
        _ = v41
        _ = v42
    }
    if (((0.1 - float32(float64(-2.25))) <= -2.25) || false) {
        var v43 uint16 = 255
        v2 += (v2 ^ v2)
        if (false || false) {
            v2 = (v2 % ((v2) | 1))
        }
        _ = v43
    } else {
        var v44 bool = (1.5 <= -2.25)
        _ = v44
    }
    v45 := GenS1{F0: 255, F1: 186, F2: 100}
    v46 := GenS6{F0: -954}
    v2 += (v2 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v45
    _ = v46
}
