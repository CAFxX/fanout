package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8) uint32 {
    v7 := []int{899, 32768, 2147483647, -761}
    var v8 int8 = 102
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[int]uint32)
    var v4 int = -207
    var v5 int = v4
    v6 := make(map[uint32]int)
    fmt.Println(v5, v4)
    v9 := GenS1{F0: -60, F1: 42}
    var v10 float32 = 1.5
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint16 = 2
        var v13 int64 = 0
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = len(v3)
    if false {
        if false {
            switch "foo bar" {
            case "":
                var v14 int32 = -2147483648
                _ = v14
            case "foo bar":
                v15 := GenS2{F0: genFn0(75), F1: 2147483647, F2: 32767}
                _ = v15
            case "hello":
                v2 = append(v2, (v5 << ((42) & 63)))
            }
            if (!(-2.25 != 0.1)) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch "\n\t" {
                    case "foo bar":
                        switch 4294967295 {
                        case -2147483648:
                            v10 = (v10 * ((0.1 - (v10 / (v10 * 1.5))) - -2.25))
                        default:
                            var v17 uint32 = 4294967295
                            _ = v17
                        }
                    case "hello":
                        delete(v6, 4294967295)
                    default:
                        v10 -= (((v10 - ((1.5 + float32((float64(((((1e10 - ((v10 / v10) - v10)) / v10) + v10) * ((v10 + v10) * v10))) * (0.1 + 0.0)))) - (((0.0 + ((v10 / (((v10 + v10) - float32(0.0)) * float32(0.0))) * v10)) - v10) + v10))) * v10) - (v10 + v10))
                    }
                    switch v4 {
                    case -2147483648:
                        var v18 string = "a"
                        _ = v18
                    default:
                        v4 -= v4
                    }
                    v5 = v5
                    _ = fi16
                }
            }
            if false {
                fmt.Println(v10, v4)
            } else {
                if (!(!(1.5 < ((v10 + v10) + 1e10)))) {
                    v19 := GenS3{F0: 0}
                    if (!((!true) || false)) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int16 = -32768
                            switch (v5 * v5) {
                            case -502:
                                var v22 uint8 = 128
                                _ = v22
                            case 127:
                                v23 := GenS2{F0: 1, F1: 2147483647, F2: 928}
                                _ = v23
                            case 9223372036854775807:
                                _ = len(v2)
                            }
                            switch "foo bar" {
                            case "":
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    var v25 int16 = (v21 & v21)
                                    switch 127 {
                                    case 32767:
                                        v26 := GenS3{F0: 128}
                                        var v27 uint = 42
                                        v21 -= -1
                                        _ = v26
                                        _ = v27
                                    case -777:
                                        fmt.Println(v21, v10)
                                    default:
                                        var v28 int16 = v21
                                        _ = v28
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        var v29 uint64 = 678
                                        _ = v29
                                    default:
                                        v30 := GenS4{F0: 0, F1: 4294967295}
                                        _ = v30
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                            default:
                                fmt.Println(v4, v4)
                            }
                            _ = fi20
                            _ = v21
                        }
                        v10 = ((v10 * v10) / 1e10)
                    } else {
                        var v31 int16 = 1
                        _ = v31
                    }
                    switch (v5 >> ((18446744073709551615) & 63)) {
                    case 4294967295:
                        v32 := GenS3{F0: 1}
                        _ = v32
                    case 127:
                        v5 = (v5 - v5)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v19
                } else {
                    var v33 int = (v4 / ((v4) | 1))
                    v34 := GenS5{F0: 0, F1: 42}
                    v33 -= v4
                    _ = v33
                    _ = v34
                }
                if (1000 != 42) {
                    fmt.Println(v10)
                }
                fmt.Println(v10, v4, v10)
            }
        }
    }
    var v35 uint64 = 4294967295
    var v36 int8 = 127
    _ = len(v2)
    fmt.Println(v4, v35)
    delete(v1, (-(v5)))
    var v37 string = "hello"
    v38 := GenS6{F0: 531}
    v36 = v36
    for fi39 := 0; fi39 < 7; fi39++ {
        switch v37 {
        case "foo bar":
            fmt.Println(v5)
        default:
            var v40 bool = true
            _ = v40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v36 = 104
            var v42 int64 = -946
            _ = fi41
            _ = v42
        }
        fmt.Println(v37)
        _ = fi39
    }
    var v43 uint64 = v35
    var v44 uint64 = v43
    var v45 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
