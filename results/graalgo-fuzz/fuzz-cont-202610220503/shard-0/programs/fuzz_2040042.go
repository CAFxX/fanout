package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint) uint32 {
    v4 := GenS1{F0: 587, F1: 42, F2: -32}
    if false {
        var v5 int8 = 100
        v6 := []int{(p0 ^ p0)}
        if false {
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 bool = true
                v9 := GenS2{F0: (-2147483648 >= -2147483648), F1: 1}
                switch 127 {
                case 543:
                    v10 := GenS2{F0: ((p1 & (p1 >> ((p1) & 63))) < 65535), F1: -74}
                    _ = v10
                default:
                    v11 := GenS3{F0: -1, F1: 724}
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p0)
    v12 := []int{127, p0, (^(p0))}
    var v13 int8 = -1
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v4
    return 949
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{}
    v3 := make(map[string]int)
    var v14 int = -9223372036854775808
    var v15 string = "a"
    var v16 uint32 = 742
    switch v14 {
    case 128:
        if true {
            if len(v2) > 0 {
                v2[0] = -948
            }
            _ = len(v2)
            _ = len(v2)
        }
    case -9223372036854775808:
        switch v14 {
        case 8:
            v17 := GenS4{F0: -324}
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int32 = -1
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    var v20 string = (v15 + v15)
                    switch v15 {
                    case "foo bar":
                        fmt.Println(v16)
                    case "hello":
                        for fi21 := 0; fi21 < 15; fi21++ {
                            if true {
                                switch (fi18 - fi18) {
                                case 337:
                                    v22 := GenS2{F0: true, F1: -1}
                                    _ = v22
                                case -382:
                                    switch (("a" + v20) + v15) {
                                    case "foo bar":
                                        v14 = -9223372036854775808
                                    case "hello":
                                        var v23 uint8 = 123
                                        _ = v23
                                    }
                                default:
                                    v16 = (v16 & (v16 * v16))
                                }
                                v24 := GenS1{F0: 703, F1: 29, F2: (v19 / ((v19) | 1))}
                                v2 = append(v2, (fi21 >> ((4294967295) & 63)))
                                _ = v24
                            }
                            v25 := GenS5{F0: (v19 % (((v19 % ((v19) | 1))) | 1))}
                            if (84 >= 1) {
                                var v26 string = ("\n\t" + v20)
                                var v27 bool = ((1e10 < 3.14159) && (4294967295 >= genFn0(fi18, 1)))
                                _ = v26
                                _ = v27
                            } else {
                                var v28 float32 = 1.5
                                _ = v28
                            }
                            v29 := GenS5{F0: 0}
                            v14 -= v14
                            _ = fi21
                            _ = v25
                            _ = v29
                        }
                    }
                    _ = v20
                }
            } else {
                v30 := GenS2{F0: false, F1: int64(v19)}
                var v31 bool = (0 != 42)
                v1[(v16 + 309)] = v31
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 uint = 749
                fmt.Println(v16, v33, v14)
                v33 = 241
                var v34 uint32 = 2
                v14 += (-(fi18))
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi18
            _ = v19
        }
    }
    switch ((v15 + ("a" + ("\n\t" + v15))) + v15) {
    case "":
        var v35 uint32 = v16
        _ = v35
    default:
        fmt.Println(v14, v15)
    }
    var v36 int = v14
    fmt.Println(v36, v14, v36)
    v37 := GenS6{F0: 42, F1: 42}
    v15 = v15
    var v38 uint32 = (^(v16))
    switch "\n\t" {
    case "foo bar":
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 uint32 = (v16 - 575)
            if false {
                v41 := GenS7{F0: 355, F1: (v16 ^ v40)}
                v15 = "\n\t"
                _ = v41
            } else {
                var v42 uint = 65535
                fmt.Println(v14)
                v40 = 65535
                _ = v42
            }
            var v43 float64 = 1e10
            _ = fi39
            _ = v40
            _ = v43
        }
    }
    fmt.Println(v38)
    v44 := GenS6{F0: -1, F1: 0}
    var v45 int8 = 75
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v44
    _ = v45
}
