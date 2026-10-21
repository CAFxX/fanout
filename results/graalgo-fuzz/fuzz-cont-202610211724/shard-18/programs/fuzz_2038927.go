package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int32; F1 int16; F2 string }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint = 290
    var v2 int64 = -1
    v3 := make(map[string]uint32)
    var v4 string = "foo bar"
    var v5 uint16 = 1
    var v6 uint = (^(v1))
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            if false {
                var v8 float64 = -2.25
                v9 := []int{}
                for fi10 := 0; fi10 < 8; fi10++ {
                    if (128 >= fi7) {
                        fmt.Println(v1, v5, v2)
                        if (1 > 42) {
                            if ((v5 | uint16(v1)) <= (v5 % ((65535) | 1))) {
                                var v11 float64 = (v8 - float64(float32((-2.25 - 0.1))))
                                if false {
                                    _ = len(v9)
                                } else {
                                    if ((v5 << ((810) & 63)) == (^(v5))) {
                                        fmt.Println(v6, v5, v2)
                                        v12 := GenS1{F0: 32767, F1: 2147483647}
                                        v2 = v2
                                        _ = v12
                                    }
                                    switch v4 {
                                    case "\n\t":
                                        if len(v9) > 0 {
                                            v9[0] = (-(fi7))
                                        }
                                    case "":
                                        fmt.Println(v5, v4)
                                    case "a":
                                        var v13 int32 = -747
                                        _ = v13
                                    }
                                    v14 := GenS2{F0: 255, F1: 787}
                                    _ = v14
                                }
                                _ = v11
                            } else {
                                delete(v3, "\n\t")
                                fmt.Println(v8, v8)
                                fmt.Println(v8, v4)
                            }
                            fmt.Println(v4, v5)
                            v3[v4] = 42
                            switch fi10 {
                            case 1:
                                v15 := GenS2{F0: 81, F1: -607}
                                _ = v15
                            }
                            var v16 uint64 = 42
                            var v17 uint32 = 2147483648
                            switch v4 {
                            case "":
                                var v18 uint32 = (v17 >> ((v1) & 63))
                                _ = v18
                            case "\n\t":
                                v19 := GenS3{F0: -967, F1: 1000, F2: "a"}
                                _ = v19
                            }
                            _ = v16
                            _ = v17
                        } else {
                            v20 := GenS3{F0: 417, F1: 261, F2: "foo bar"}
                            v6 -= v6
                            v21 := GenS2{F0: 42, F1: 32767}
                            _ = v20
                            _ = v21
                        }
                    } else {
                        v8 -= v8
                    }
                    fmt.Println(v1)
                    v2 = (v2 * -29)
                    _ = fi10
                }
                _ = v8
                _ = v9
            } else {
                v22 := GenS3{F0: -390, F1: -1, F2: v4}
                _ = v22
            }
            switch (fi7 * 4294967295) {
            case 4294967295:
                var v23 uint = (v6 / ((v6) | 1))
                _ = v23
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                v4 = v4
                _ = fi24
            }
        }
        _ = fi7
    }
    fmt.Println(v2, v5, v4)
    if (!true) {
        v25 := GenS4{F0: 42}
        for fi26 := 0; fi26 < 18; fi26++ {
            if (true || ((v1 <= (v6 + (v6 / (((^(v6))) | 1)))) && (18446744073709551615 <= (v1 & v1)))) {
                var v27 int8 = 1
                var v28 uint = v6
                _ = v27
                _ = v28
            } else {
                fmt.Println(v1, v2, fi26)
                var v29 float32 = ((1.5 - 1e10) + float32(1e10))
                _ = v29
            }
            if true {
                fmt.Println(v5)
            }
            _ = fi26
        }
        if false {
            v6 -= v6
            var v30 bool = (!(false && true))
            _ = v30
        } else {
            fmt.Println(v4)
        }
        _ = v25
    }
    var v31 int8 = -88
    var v32 int32 = -399
    for fi33 := 0; fi33 < 4; fi33++ {
        switch (fi33 % ((fi33) | 1)) {
        case 127:
            var v34 float32 = ((1.5 * -2.25) / (1e10 + 1.5))
            _ = v34
        default:
            delete(v3, "\n\t")
        }
        _ = fi33
    }
    _ = len(v3)
    v6 = (v1 / ((1) | 1))
    fmt.Println(v2)
    _ = len(v3)
    if false {
        var v35 string = ""
        _ = v35
    }
    v2 = 42
    v32 = (v32 >> (((^(v6))) & 63))
    switch v4 {
    case "\n\t":
        v36 := GenS1{F0: 32768, F1: 2147483647}
        _ = v36
    }
    var v37 int32 = v32
    v37 = (v37 & -2147483648)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
