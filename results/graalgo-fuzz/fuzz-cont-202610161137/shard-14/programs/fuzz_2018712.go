package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }
func genFn0(p0 int, p1 string, p2 int32) uint8 {
    if true {
        var v4 float32 = 1.5
        v5 := []int{127, (p0 - -844)}
        var v6 int64 = 1
        _ = v4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 227
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 365
    var v3 float32 = 1e10
    v8 := []int{}
    switch "foo bar" {
    case "hello":
        switch 74 {
        case -2147483648:
            fmt.Println(v3, v2, v3)
        }
    }
    fmt.Println(v2)
    var v9 uint16 = 65535
    var v10 string = "\n\t"
    switch 1 {
    case 395:
        var v11 int32 = -1
        _ = v11
    case 2147483647:
        for fi12 := 0; fi12 < 12; fi12++ {
            v10 = v10
            _ = fi12
        }
    case -696:
        var v13 uint8 = genFn0(0, v10, -534)
        _ = v13
    default:
        if true {
            v14 := GenS1{F0: "foo bar"}
            v2 = 18446744073709551615
            if false {
                var v15 uint16 = v9
                if (4294967295 <= 2147483647) {
                    v9 = uint16(v2)
                    var v16 bool = (v10 != v10)
                    v17 := GenS1{F0: ((v10 + (v10 + v10)) + ((((((v10 + "") + v10) + "\n\t") + v10) + v10) + "a"))}
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        fmt.Println(v3)
                        _ = fi18
                    }
                }
                v10 = v10
                _ = v15
            }
            _ = v14
        } else {
            v8 = append(v8, int(v9))
            if (2 <= uint8(v9)) {
                for fi19 := 0; fi19 < 19; fi19++ {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v2 = 0
                        if true {
                            v21 := GenS2{F0: (fi20 ^ 32767), F1: -86, F2: 234}
                            if (-875 == 1) {
                                if (true && true) {
                                    var v22 float32 = v3
                                    _ = v22
                                }
                                fmt.Println(v1)
                            }
                            fmt.Println(v1)
                            _ = v21
                        } else {
                            v3 += (0.0 + 0.1)
                        }
                        if false {
                            if (false || (!true)) {
                                var v23 float64 = (((0.0 - float64(v3)) - (1.5 * 1.5)) - (0.0 + (1.5 + float64(v3))))
                                if (v9 < v9) {
                                    v3 = float32(v23)
                                    _ = len(v8)
                                    v24 := GenS2{F0: int(v1), F1: 42, F2: 127}
                                    var v25 uint16 = (v9 & (v9 % ((v9) | 1)))
                                    _ = v24
                                    _ = v25
                                }
                                switch (fi20 + 4294967295) {
                                case 947:
                                    v8 = append(v8, fi19)
                                }
                                _ = v23
                            } else {
                                var v26 uint = 539
                                v27 := GenS1{F0: ""}
                                switch fi19 {
                                case 9223372036854775807:
                                    var v28 uint64 = (v2 / ((v2) | 1))
                                    if len(v8) > 0 {
                                        v8[0] = -2147483648
                                    }
                                    _ = v28
                                default:
                                    var v29 bool = false
                                    _ = v29
                                }
                                v26 = v26
                                _ = v26
                                _ = v27
                            }
                            fmt.Println(fi20)
                            var v30 bool = (0.1 < 3.14159)
                            _ = v30
                        } else {
                            v9 = (v9 | 0)
                            v9 += v9
                            var v31 int8 = 100
                            if false {
                                v31 = v31
                            } else {
                                v32 := GenS3{F0: 2, F1: 255, F2: 318}
                                if false {
                                    v31 += v31
                                } else {
                                    var v33 string = ((v10 + v10) + "foo bar")
                                    _ = v33
                                }
                                _ = v32
                            }
                            var v34 uint8 = 255
                            var v35 string = ((("hello" + v10) + ((v10 + (("\n\t" + v10) + v10)) + v10)) + v10)
                            _ = v31
                            _ = v34
                            _ = v35
                        }
                        _ = fi20
                    }
                    v36 := GenS4{F0: 180, F1: v10}
                    v37 := GenS2{F0: fi19, F1: 100, F2: uint8(v2)}
                    _ = fi19
                    _ = v36
                    _ = v37
                }
                v38 := GenS5{F0: 0}
                if len(v8) > 0 {
                    v8[0] = -205
                }
                _ = v38
            } else {
                v8 = append(v8, -850)
            }
            fmt.Println(v3, v10, v10)
        }
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        switch (^(fi39)) {
        case 4294967295:
            var v40 int = (fi39 - fi39)
            _ = v40
        }
        v41 := GenS6{F0: 18446744073709551615}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi39
        _ = v41
    }
    switch v10 {
    case "":
        var v42 int64 = 241
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
