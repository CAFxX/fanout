package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0() int8 {
    var v7 int = -9223372036854775808
    if ((222 > 0) && (1e10 <= float32(0.1))) {
        v8 := GenS1{F0: 65535, F1: 32}
        var v9 uint8 = 10
        _ = v8
        _ = v9
    } else {
        v7 = -1
        v7 += v7
        var v10 int64 = -1
        _ = v10
    }
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 int64, p2 uint32) int32 {
    switch -9223372036854775808 {
    case 839:
        p1 += (-(p1))
    case 9223372036854775807:
        p2 = (p2 ^ 2147483647)
    case 381:
        switch 127 {
        case 2147483647:
            fmt.Println(p2, p1)
        default:
            if true {
                if false {
                    v11 := []int{42, -2147483648, 4294967295}
                    _ = v11
                }
                p0 -= p0
            }
        }
    default:
        v12 := []int{-350, 32768, 127}
        _ = v12
    }
    p0 -= 267
    v13 := []int{42, 9223372036854775807, -385}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -616
    var v3 string = ("hello" + "\n\t")
    var v4 uint8 = 0
    var v5 uint16 = 42
    var v6 uint = 1
    for fi14 := 0; fi14 < 9; fi14++ {
        v2 = (v2 / ((v2) | 1))
        if (42 > genFn1(1, 768, 2147483648)) {
            var v15 uint64 = 171
            switch v3 {
            case "a":
                fmt.Println(v15)
            case "hello":
                fmt.Println(v3, v4)
            }
            v4 = v4
            _ = v15
        } else {
            var v16 string = ("\n\t" + v3)
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS1{F0: 65535, F1: (v4 / (((v4 >> ((v6) & 63))) | 1))}
            v19 := []int{}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
    }
    v20 := GenS2{F0: 9223372036854775807}
    if true {
        fmt.Println(v3, v2)
        v3 = "hello"
    } else {
        v6 = v6
    }
    v4 = v4
    v3 = (v3 + ((v3 + ("\n\t" + (("\n\t" + v3) + "\n\t"))) + "a"))
    if true {
        for fi21 := 0; fi21 < 17; fi21++ {
            if false {
                var v22 uint64 = 2
                var v23 float32 = (1e10 / 3.14159)
                for fi24 := 0; fi24 < 6; fi24++ {
                    if true {
                        v25 := GenS1{F0: 52, F1: v4}
                        v22 += v22
                        var v26 int16 = v2
                        _ = v25
                        _ = v26
                    } else {
                        var v27 bool = true
                        switch 385 {
                        case 9223372036854775807:
                            v23 = 1e10
                        default:
                            var v28 int = int(v6)
                            _ = v28
                        }
                        _ = v27
                    }
                    v29 := GenS2{F0: -614}
                    switch fi21 {
                    case -1:
                        switch fi24 {
                        case -269:
                            fmt.Println(fi24, v5)
                        case 2147483647:
                            switch fi24 {
                            case 42:
                                v23 -= 0.0
                            case 4294967295:
                                v30 := GenS2{F0: 42}
                                fmt.Println(fi24, v5, v2)
                                switch 32768 {
                                case 2147483647:
                                    v31 := []int{0, (^(fi24)), 127, (fi24 + fi21)}
                                    _ = v31
                                case 906:
                                    v32 := GenS2{F0: 0}
                                    _ = v32
                                case -2147483648:
                                    var v33 int = fi24
                                    _ = v33
                                }
                                _ = v30
                            default:
                                var v34 int = fi24
                                _ = v34
                            }
                        default:
                            var v35 uint32 = 1
                            _ = v35
                        }
                    case 42:
                        fmt.Println(v22, v2)
                    case 32767:
                        if ((^(fi21)) != 128) {
                            if false {
                                v36 := v1[(fi21 + int(v6))]
                                v23 = v23
                                switch (fi24 / (((v36 | -902)) | 1)) {
                                case -234:
                                    fmt.Println(fi24, v6, fi21)
                                case 2147483647:
                                    fmt.Println(v5, v5, v36)
                                case -633:
                                    switch v3 {
                                    case "hello":
                                        v37 := []int{int(v4)}
                                        _ = v37
                                    default:
                                        if (!false) {
                                            var v38 uint16 = 2
                                            v39 := GenS3{F0: 2147483647, F1: (v22 * (v22 * 1))}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v40 := GenS2{F0: 42}
                                            _ = v40
                                        }
                                    }
                                }
                                _ = v36
                            }
                            v5 = 2
                        }
                    default:
                        var v41 uint = v6
                        _ = v41
                    }
                    _ = fi24
                    _ = v29
                }
                _ = v22
                _ = v23
            }
            _ = fi21
        }
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        var v43 float64 = 1e10
        _ = fi42
        _ = v43
    }
    fmt.Println(v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
