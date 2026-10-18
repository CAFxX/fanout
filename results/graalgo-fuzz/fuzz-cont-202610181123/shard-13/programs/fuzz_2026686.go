package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 uint16, p1 uint64, p2 string) uint16 {
    var v7 uint8 = 1
    fmt.Println(v7, v7)
    switch "" {
    case "a":
        if false {
            v8 := []int{32767, -9223372036854775808}
            p1 -= p1
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 uint16 = p0
                var v11 uint8 = v7
                var v12 int16 = 0
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
        } else {
            var v13 int = 0
            var v14 int32 = 2147483647
            _ = v13
            _ = v14
        }
    case "":
        v15 := []int{}
        _ = v15
    case "\n\t":
        v16 := GenS1{F0: "a", F1: 0}
        _ = v16
    default:
        var v17 int = 4294967295
        _ = v17
    }
    if (false && ((^(v7)) <= v7)) {
        var v18 int8 = 0
        p1 = p1
        _ = v18
    } else {
        fmt.Println(v7, v7, v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    v1 := []int{}
    v2 := make(map[string]int)
    var v3 uint64 = 4294967295
    var v4 int32 = int32(v3)
    var v5 uint64 = v3
    v6 := make(map[int]uint32)
    var v19 uint16 = 1
    switch 2147483647 {
    case 127:
        if (true && (-155 > -1)) {
            v19 -= (^(v19))
            var v20 uint64 = (v5 - (v5 % (((^(v5))) | 1)))
            _ = len(v2)
            _ = v20
        } else {
            if (true && false) {
                for fi21 := 0; fi21 < 5; fi21++ {
                    fmt.Println(v19, v19, v19)
                    if (!(0 != 42)) {
                        v5 = 42
                    } else {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            if false {
                                v5 += (v5 % (((^(v5))) | 1))
                                if false {
                                    if (!true) {
                                        switch "foo bar" {
                                        case "a":
                                            v6[fi21] = 65535
                                        case "foo bar":
                                            v23 := GenS1{F0: "", F1: 672}
                                            _ = v23
                                        default:
                                            switch fi21 {
                                            case 32768:
                                                v4 += (v4 * 42)
                                            case 9223372036854775807:
                                                v4 = v4
                                            case 2147483647:
                                                var v24 int32 = v4
                                                _ = v24
                                            default:
                                                v4 += 965
                                                v3 = (v3 + (v3 | 18446744073709551615))
                                            }
                                        }
                                        switch "a" {
                                        case "":
                                            v25 := v2["foo bar"]
                                            _ = v25
                                        case "foo bar":
                                            switch fi22 {
                                            case 32767:
                                                v3 = 18446744073709551615
                                            case 32768:
                                                fmt.Println(fi22, fi21, v5)
                                            }
                                        default:
                                            v26 := GenS1{F0: "\n\t", F1: 42}
                                            _ = v26
                                        }
                                        var v27 int16 = 0
                                        _ = v27
                                    } else {
                                        v4 = (v4 & 492)
                                        v28 := v2["\n\t"]
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: 1, F1: "a"}
                                    switch fi21 {
                                    case -343:
                                        fmt.Println(v4, v19)
                                    case 2147483647:
                                        var v30 bool = (true && true)
                                        _ = v30
                                    default:
                                        v6[fi21] = 2147483647
                                    }
                                    _ = v29
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v3)
                                }
                            } else {
                                switch "" {
                                case "\n\t":
                                    v31 := GenS3{F0: 255}
                                    _ = v31
                                case "hello":
                                    fmt.Println(fi22, v4)
                                default:
                                    fmt.Println(fi21)
                                }
                            }
                            v32 := GenS4{F0: (v4 / ((42) | 1)), F1: 65535, F2: 2147483648}
                            _ = fi22
                            _ = v32
                        }
                    }
                    _ = fi21
                }
            }
        }
    case 231:
        switch ("a" + "foo bar") {
        case "hello":
            if false {
                var v33 int16 = 32767
                _ = v33
            }
        default:
            v6[-9223372036854775808] = 640
        }
    }
    if false {
        v5 -= (v3 ^ 4294967295)
    } else {
        _ = len(v6)
        if (4294967295 < 9223372036854775807) {
            v6[-2147483648] = 2
        } else {
            fmt.Println(v3, v19, v4)
        }
        var v34 uint16 = (v19 & genFn0((v19 << ((69) & 63)), v3, "hello"))
        _ = v34
    }
    switch 810 {
    case 143:
        v35 := v2["\n\t"]
        _ = v35
    case 1:
        fmt.Println(v5, v5, v5)
    }
    var v36 int16 = 464
    var v37 int16 = (v36 + (v36 % ((v36) | 1)))
    v36 = (v37 + (v36 * v37))
    if (!(!true)) {
        var v38 uint16 = v19
        v5 = v3
        _ = v38
    }
    var v39 uint64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
