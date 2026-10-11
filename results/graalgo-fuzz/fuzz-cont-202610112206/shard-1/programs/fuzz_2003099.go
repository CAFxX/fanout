package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint32 }
func genFn0(p0 uint16) int64 {
    v6 := []int{2147483647}
    var v7 bool = ((!false) || true)
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 int32 = 42
    var v3 uint32 = 2147483648
    v4 := make(map[string]bool)
    var v5 bool = (1 >= 237)
    fmt.Println(v5, v3)
    fmt.Println(v5, v3, v3)
    var v8 uint32 = (v3 + (v3 + 908))
    var v9 bool = true
    v3 -= v8
    var v10 int64 = 1
    var v11 int64 = -245
    switch "a" {
    case "foo bar":
        v5 = v5
    case "\n\t":
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: true, F1: 42}
            var v14 int8 = 42
            var v15 int32 = int32(v11)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    case "hello":
        for fi16 := 0; fi16 < 15; fi16++ {
            if (float32((-2.25 - (0.0 * (float64(0.0) + 0.1)))) == -2.25) {
                v17 := GenS1{F0: v5, F1: 0}
                _ = v17
            } else {
                v18 := GenS2{F0: 255}
                if v5 {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint = 131
                        _ = len(v4)
                        switch v1 {
                        case "foo bar":
                            v2 += (v2 / (((v2 ^ v2)) | 1))
                        default:
                            if false {
                                var v21 uint = v20
                                var v22 uint8 = 255
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    v1 = v1
                    var v23 int8 = 100
                    _ = v23
                }
                _ = v18
            }
            v11 = v10
            var v24 int8 = 1
            _ = fi16
            _ = v24
        }
    default:
        fmt.Println(v2, v3, v8)
    }
    var v25 uint32 = 2
    v3 = v3
    switch v1 {
    case "a":
        v4[v1] = (32767 <= -1)
    }
    var v26 uint8 = 42
    v3 += v3
    v27 := []int{32768}
    for fi28 := 0; fi28 < 4; fi28++ {
        if v9 {
            var v29 int16 = 1000
            v30 := GenS1{F0: (v9 || true), F1: 838}
            v4[(v1 + ("foo bar" + v1))] = true
            _ = v29
            _ = v30
        } else {
            if v9 {
                switch "" {
                case "foo bar":
                    for fi31 := 0; fi31 < 15; fi31++ {
                        switch -1 {
                        case 0:
                            fmt.Println(v3, v2, v11)
                        }
                        fmt.Println(v8, v8)
                        _ = fi31
                    }
                case "hello":
                    v2 = int32(fi28)
                case "a":
                    if (0 <= (^(v2))) {
                        if (genFn0(2) > (v11 >> ((18446744073709551615) & 63))) {
                            var v32 float64 = float64(-2.25)
                            if v9 {
                                var v33 int64 = 1
                                var v34 int32 = int32(v3)
                                for fi35 := 0; fi35 < 3; fi35++ {
                                    var v36 int = fi28
                                    v26 -= 128
                                    var v37 int8 = -128
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v33
                                _ = v34
                            } else {
                                v2 = (v2 << ((1) & 63))
                                v8 = (v25 * (v3 >> ((42) & 63)))
                                fmt.Println(fi28, v26, v9)
                            }
                            var v38 uint32 = v8
                            _ = v32
                            _ = v38
                        } else {
                            v39 := GenS2{F0: 65535}
                            _ = len(v27)
                            _ = v39
                        }
                    }
                default:
                    switch (v1 + (v1 + (v1 + "foo bar"))) {
                    case "a":
                        var v40 uint = 42
                        _ = v40
                    default:
                        var v41 int64 = v10
                        _ = v41
                    }
                }
            } else {
                v11 -= (v10 << ((704) & 63))
                var v42 uint16 = uint16(v26)
                _ = v42
            }
            v43 := GenS3{F0: 186, F1: (v25 % ((854) | 1)), F2: v8}
            _ = v43
        }
        _ = fi28
    }
    if false {
        fmt.Println(v2, v2, v9)
        v44 := GenS3{F0: 2, F1: 65535, F2: (v25 + v25)}
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int = (fi45 / ((9223372036854775807) | 1))
            switch fi45 {
            case -801:
                v10 = 1
            case 42:
                fmt.Println(fi45, v1, v46)
            case -1:
                v25 = (^(v3))
            }
            var v47 uint8 = (v26 * v26)
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = v44
    } else {
        v26 += (v26 & 123)
    }
    fmt.Println(v2, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
