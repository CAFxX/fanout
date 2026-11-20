package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 bool = ((!(!false)) && (!true))
    v4 := make(map[string]int)
    var v5 string = "hello"
    v1["foo bar"] = 58
    var v6 int8 = 42
    var v7 int32 = -2147483648
    var v8 string = v5
    switch "" {
    case "foo bar":
        fmt.Println(v7, v8, v8)
    case "":
        var v9 uint32 = 0
        _ = v9
    }
    var v10 uint8 = 127
    fmt.Println(v8)
    fmt.Println(v5, v10, v3)
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            if v3 {
                v13 := []int{}
                if (v3 && v3) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                v3 = v3
                _ = v13
            } else {
                fmt.Println(v5, fi11, v8)
                v14 := GenS1{F0: (!(uint64(v10) <= 18)), F1: -1, F2: v10}
                _ = v14
            }
            switch fi11 {
            case 884:
                if (0.0 == -2.25) {
                    var v15 int16 = 204
                    _ = v15
                }
            default:
                switch v8 {
                case "foo bar":
                    switch ("" + ("foo bar" + v8)) {
                    case "\n\t":
                        if (((0.1 * 0.1) != float32(1.5)) && true) {
                            fmt.Println(v6, v10, v10)
                            if v3 {
                                switch "a" {
                                case "a":
                                    if (2 != 18446744073709551615) {
                                        v4[(v5 + v5)] = (fi11 | fi11)
                                    } else {
                                        var v16 int32 = -1
                                        var v17 float64 = -2.25
                                        _ = v16
                                        _ = v17
                                    }
                                case "foo bar":
                                    v5 = "a"
                                case "":
                                    if (v3 || (((-(v7)) >= v7) && ("" == ("hello" + "a")))) {
                                        var v18 uint32 = 831
                                        _ = v18
                                    } else {
                                        if (2147483647 < 837) {
                                            v6 = (v6 ^ v6)
                                            v19 := GenS1{F0: (false && (v10 <= v10)), F1: 42, F2: 129}
                                            _ = v19
                                        }
                                        var v20 int8 = v6
                                        _ = v20
                                    }
                                }
                                var v21 uint64 = 18446744073709551615
                                fmt.Println(v10)
                                v22 := []int{fi11, fi11}
                                switch fi11 {
                                case 0:
                                    var v23 int32 = -377
                                    _ = v23
                                case 738:
                                    var v24 int8 = (-(v6))
                                    _ = v24
                                }
                                _ = v21
                                _ = v22
                            } else {
                                v6 -= v6
                                switch v8 {
                                case "hello":
                                    v25 := GenS1{F0: (127 <= 255), F1: -813, F2: 42}
                                    _ = v25
                                default:
                                    if false {
                                        var v26 float32 = 0.1
                                        fmt.Println(v10, fi12)
                                        v5 = "foo bar"
                                        _ = v26
                                    } else {
                                        fmt.Println(v8, v8, v10)
                                        v5 = v8
                                    }
                                }
                            }
                            v5 = v8
                            fmt.Println(v7, v6)
                            switch fi11 {
                            case 9223372036854775807:
                                v27 := GenS1{F0: false, F1: 625, F2: 1}
                                _ = v27
                            case 128:
                                var v28 float32 = (0.1 * 1e10)
                                _ = v28
                            }
                            v5 = (v5 + ("" + (v8 + "foo bar")))
                            var v29 uint64 = 18446744073709551615
                            if v3 {
                                _ = len(v2)
                                v3 = (v3 && v3)
                            } else {
                                switch ("foo bar" + "a") {
                                case "hello":
                                    if false {
                                        fmt.Println(v10, v10)
                                        v30 := GenS2{F0: 2147483648}
                                        var v31 int8 = 30
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v32 := GenS3{F0: 868, F1: 2147483647, F2: (0 == 770)}
                                        v29 = 18446744073709551615
                                        fmt.Println(v3, v10, v8)
                                        _ = v32
                                    }
                                case "":
                                    fmt.Println(v8)
                                case "a":
                                    v29 += v29
                                }
                                v3 = false
                                v4[v8] = -1
                            }
                            _ = v29
                        } else {
                            if v3 {
                                var v33 int8 = (^(v6))
                                switch v8 {
                                case "foo bar":
                                    fmt.Println(v6)
                                case "a":
                                    v7 -= 953
                                default:
                                    var v34 float64 = 1e10
                                    _ = v34
                                }
                                _ = v33
                            }
                            v35 := GenS2{F0: 0}
                            _ = v35
                        }
                    case "hello":
                        v6 -= (v6 & v6)
                    case "a":
                        var v36 uint32 = 2147483648
                        _ = v36
                    }
                case "":
                    v37 := []int{32767, 4294967295, (fi12 ^ fi11), (fi11 ^ fi11)}
                    _ = v37
                }
            }
            v6 -= v6
            _ = fi12
        }
        var v38 uint16 = 870
        _ = fi11
        _ = v38
    }
    v8 = (((v8 + v5) + "") + (v8 + v8))
    switch 705 {
    case 42:
        v39 := GenS1{F0: true, F1: 245, F2: 13}
        _ = v39
    default:
        var v40 int8 = (v6 % (((v6 * (v6 / (((v6 & 1)) | 1)))) | 1))
        _ = v40
    }
    v7 += (v7 * 42)
    v41 := []int{-347}
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
