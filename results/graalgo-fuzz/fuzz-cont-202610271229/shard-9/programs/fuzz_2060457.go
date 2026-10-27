package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int; F1 uint16; F2 string }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 int) bool {
    if (73 < 2) {
        var v5 uint = 255
        fmt.Println(v5)
        _ = v5
    } else {
        var v6 int32 = -926
        _ = v6
    }
    var v7 int64 = -897
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 bool) uint8 {
    v8 := GenS1{F0: 128}
    v9 := []int{}
    _ = p0
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    var v3 uint64 = 620
    var v4 uint32 = 594
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: fi10}
        switch ("foo bar" + "hello") {
        case "hello":
            if true {
                for fi12 := 0; fi12 < 18; fi12++ {
                    v1 = v1
                    var v13 int32 = 0
                    var v14 string = "hello"
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 uint = 2
                var v16 string = ("a" + "hello")
                _ = v15
                _ = v16
            } else {
                v17 := []int{(fi10 ^ fi10), int(v4), (fi10 * 32767)}
                v1 += v1
                v18 := GenS2{F0: 32767, F1: 0, F2: 21}
                _ = v17
                _ = v18
            }
        default:
            var v19 int = (fi10 ^ fi10)
            _ = v19
        }
        var v20 int8 = 100
        _ = fi10
        _ = v11
        _ = v20
    }
    var v21 string = ""
    var v22 int32 = 2147483647
    if false {
        delete(v2, 4294967295)
    } else {
        fmt.Println(v1, v4, v21)
        var v23 int32 = 771
        _ = v23
    }
    if (!(!(!(uint(v3) < uint(v3))))) {
        var v24 uint64 = v3
        _ = v24
    }
    v25 := GenS1{F0: -197}
    var v26 float64 = 1e10
    var v27 uint = (v1 ^ (v1 % ((969) | 1)))
    if (1000 <= 32767) {
        v28 := GenS3{F0: 2, F1: v3, F2: 0}
        v29 := []int{32767}
        v1 -= (v27 ^ (v27 ^ (v1 << (((v27 >> (((^(v1))) & 63))) & 63))))
        _ = v28
        _ = v29
    } else {
        var v30 int = -42
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 bool = false
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    fmt.Println(v26, v21, v1)
    var v33 bool = (((^(v4)) != (v4 % (((v4 | v4)) | 1))) || ((^(v3)) == (v3 >> ((v27) & 63))))
    v34 := []int{9223372036854775807}
    switch 42 {
    case 127:
        fmt.Println(v27, v1)
    case 128:
        if (-2.25 >= 3.14159) {
            var v35 bool = (!(v3 < (v3 % ((v3) | 1))))
            v1 += v27
            for fi36 := 0; fi36 < 3; fi36++ {
                switch (fi36 >> (((v27 * (v1 - 65535))) & 63)) {
                case -565:
                    var v37 string = (v21 + v21)
                    _ = v37
                case 0:
                    if v35 {
                        v1 = (v27 ^ v1)
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 uint = (v1 - (v1 - 451))
                            v40 := GenS4{F0: 396, F1: uint64(v1)}
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        v26 = 1.5
                    } else {
                        switch "\n\t" {
                        case "hello":
                            if (1 != 42) {
                                v41 := GenS5{F0: fi36, F1: 255, F2: "hello"}
                                v27 = v27
                                _ = v41
                            } else {
                                v27 += 18446744073709551615
                            }
                        case "\n\t":
                            var v42 int64 = 1
                            _ = v42
                        default:
                            if v33 {
                                switch (fi36 * fi36) {
                                case 1:
                                    switch fi36 {
                                    case -300:
                                        var v43 float32 = (-2.25 / 1.0)
                                        _ = v43
                                    default:
                                        for fi44 := 0; fi44 < 12; fi44++ {
                                            if false {
                                                v35 = v33
                                            } else {
                                                v1 -= v1
                                            }
                                            if v35 {
                                                v45 := GenS6{F0: 0}
                                                v22 += 2147483647
                                                _ = v45
                                            }
                                            switch (fi44 * 127) {
                                            case 128:
                                                fmt.Println(v1, v1)
                                            case -9223372036854775808:
                                                v35 = (v33 || (!(-1 >= -9223372036854775808)))
                                            default:
                                                var v46 int64 = 1
                                                if (32767 > 65535) {
                                                    var v47 uint = v27
                                                    _ = v47
                                                }
                                                _ = v46
                                            }
                                            _ = fi44
                                        }
                                    }
                                case -9223372036854775808:
                                    if false {
                                        if len(v34) > 0 {
                                            v34[0] = fi36
                                        }
                                    }
                                case -661:
                                    v26 += (-2.25 * 0.1)
                                default:
                                    var v48 int16 = -1
                                    _ = v48
                                }
                                var v49 uint64 = 0
                                _ = v49
                            } else {
                                var v50 uint64 = v3
                                var v51 float64 = v26
                                fmt.Println(fi36, v1, v4)
                                _ = v50
                                _ = v51
                            }
                        }
                    }
                case -318:
                    var v52 string = v21
                    _ = v52
                }
                _ = fi36
            }
            _ = v35
        } else {
            fmt.Println(v33, v22, v27)
        }
    case 32768:
        if len(v34) > 0 {
            v34[0] = 42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
