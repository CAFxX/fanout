package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64; F1 string; F2 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint16 = 585
    var v4 string = "a"
    fmt.Println(v4, v4)
    v5 := GenS1{F0: 563}
    switch v4 {
    case "\n\t":
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint8 = 211
            _ = fi6
            _ = v7
        }
    case "":
        v3 = 654
    case "a":
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        v9 := []int{180, 537, -2147483648}
        _ = v9
    }
    fmt.Println(v3)
    var v10 uint64 = 18446744073709551615
    v3 = (v3 ^ v3)
    fmt.Println(v4, v4)
    v4 = (v4 + v4)
    for fi11 := 0; fi11 < 5; fi11++ {
        switch (-(fi11)) {
        case -426:
            for fi12 := 0; fi12 < 20; fi12++ {
                switch (("\n\t" + v4) + ((((v4 + v4) + v4) + "") + ("a" + "hello"))) {
                case "a":
                    if (42 != 1) {
                        switch fi12 {
                        case 2147483647:
                            v10 += v10
                        case 32767:
                            var v13 int8 = 104
                            _ = v13
                        default:
                            v14 := []int{fi11, -524, 127}
                            _ = v14
                        }
                        v3 += (v3 ^ 32767)
                        v15 := GenS2{F0: -185, F1: ("a" + "\n\t")}
                        if (!(true && (v3 <= (^(v3))))) {
                            if true {
                                v16 := GenS2{F0: 0, F1: v4}
                                _ = v16
                            } else {
                                if (0.0 == 0.0) {
                                    var v17 uint16 = v3
                                    v18 := GenS2{F0: -1, F1: (v4 + v4)}
                                    v19 := GenS3{F0: -1, F1: v4, F2: 42}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    v20 := GenS4{F0: (2147483647 >= 1), F1: int16(fi11), F2: (v10 - (v10 / ((988) | 1)))}
                                    _ = v20
                                }
                                var v21 int = fi11
                                var v22 bool = false
                                _ = v21
                                _ = v22
                            }
                            v4 = v4
                            v23 := v2["hello"]
                            _ = v23
                        } else {
                            v24 := GenS5{F0: 42, F1: -113}
                            var v25 uint32 = 0
                            _ = v24
                            _ = v25
                        }
                        _ = v15
                    } else {
                        var v26 uint16 = 5
                        _ = len(v2)
                        v26 += v3
                        v27 := v2[""]
                        v28 := []int{fi11, (fi11 | fi11), (fi12 | (fi11 + (fi11 | fi12))), (^(fi11))}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                case "":
                    if true {
                        switch v4 {
                        case "":
                            fmt.Println(v10, v3, fi12)
                        case "foo bar":
                            var v29 int8 = 1
                            v30 := GenS2{F0: 9223372036854775807, F1: "\n\t"}
                            fmt.Println(v10, v29)
                            _ = v29
                            _ = v30
                        }
                        v3 += (v3 & v3)
                        v4 = v4
                    }
                case "foo bar":
                    switch fi11 {
                    case 816:
                        v3 = v3
                    default:
                        var v31 int = (fi12 / ((fi12) | 1))
                        _ = v31
                    }
                    var v32 uint64 = v10
                    if true {
                        if true {
                            var v33 string = ("a" + ((((v4 + v4) + "foo bar") + v4) + ("foo bar" + (v4 + v4))))
                            delete(v1, v33)
                            _ = v33
                        } else {
                            var v34 float64 = (1e10 - float64(1.5))
                            var v35 bool = (!false)
                            v36 := []int{(fi12 ^ (fi11 * 2147483647))}
                            var v37 float32 = 1.5
                            var v38 string = "a"
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        v3 += (v3 & 968)
                        v32 = 901
                    }
                    _ = v32
                default:
                    if true {
                        switch fi12 {
                        case -2147483648:
                            if false {
                                var v39 int64 = 141
                                if false {
                                    fmt.Println(fi11, v4)
                                    if false {
                                        v1[v4] = 2
                                    } else {
                                        fmt.Println(fi11, v10, v4)
                                        fmt.Println(v4, v4)
                                    }
                                }
                                _ = v39
                            }
                        case 293:
                            if (379 != 42) {
                                delete(v2, v4)
                                var v40 string = v4
                                var v41 uint64 = v10
                                _ = v40
                                _ = v41
                            }
                        case 127:
                            v1[(v4 + ("\n\t" + v4))] = 2147483647
                        }
                        v42 := v2[""]
                        if ((v10 >= (v10 | (v10 + v10))) && false) {
                            var v43 string = "\n\t"
                            _ = v43
                        }
                        var v44 bool = (!v42)
                        var v45 float64 = (0.0 + 0.1)
                        _ = v42
                        _ = v44
                        _ = v45
                    } else {
                        fmt.Println(fi12)
                        v46 := GenS6{F0: (0 >= 2), F1: (fi11 & (fi11 * 498))}
                        _ = v46
                    }
                }
                fmt.Println(v4, fi12)
                v47 := GenS3{F0: 704, F1: v4, F2: 0}
                v10 = (^(v10))
                v48 := GenS7{F0: 0}
                _ = fi12
                _ = v47
                _ = v48
            }
        case 32768:
            v49 := []int{961, (fi11 - (^(fi11))), fi11, fi11}
            _ = v49
        case -9223372036854775808:
            switch (fi11 << ((18446744073709551615) & 63)) {
            case 9223372036854775807:
                v50 := []int{fi11, (fi11 % ((fi11) | 1)), fi11, (fi11 + fi11)}
                _ = v50
            default:
                var v51 uint = 0
                _ = v51
            }
        default:
            fmt.Println(v3, v3, fi11)
        }
        _ = fi11
    }
    var v52 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
