package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int32; F1 int16 }
func genFn0(p0 int32) int8 {
    p0 = p0
    var v7 uint64 = 631
    v8 := []int{}
    v7 = 2
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint16 = 412
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 float32 = float32(float64(float32(3.14159)))
    var v2 uint8 = 250
    var v3 uint32 = 804
    var v4 string = ""
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    var v11 int8 = -77
    if true {
        v1 = v1
    } else {
        fmt.Println(v2, v2, v4)
        if (0.0 == v1) {
            fmt.Println(v3, v1)
        } else {
            if true {
                for fi12 := 0; fi12 < 10; fi12++ {
                    switch (fi12 | fi12) {
                    case 1:
                        v2 = (^(v2))
                    case -2147483648:
                        for fi13 := 0; fi13 < 4; fi13++ {
                            if ((^(v2)) >= 127) {
                                var v14 uint16 = 149
                                v15 := GenS1{F0: 2147483647}
                                _ = v14
                                _ = v15
                            } else {
                                var v16 string = v4
                                v3 = (v3 | (v3 % ((v3) | 1)))
                                _ = v16
                            }
                            if false {
                                v3 += v3
                                var v17 int64 = 0
                                v18 := GenS1{F0: -630}
                                _ = v17
                                _ = v18
                            } else {
                                switch ((v4 + "a") + "a") {
                                case "a":
                                    if (!true) {
                                        v11 += v11
                                        v2 += (v2 % ((1) | 1))
                                        var v19 bool = true
                                        _ = v19
                                    } else {
                                        switch 4294967295 {
                                        case -9223372036854775808:
                                            if true {
                                                v2 -= (v2 | 50)
                                                if ((3.14159 + 1.5) <= (1e10 + 1e10)) {
                                                    v11 = -128
                                                }
                                            } else {
                                                v20 := GenS2{F0: -1, F1: 710, F2: (v1 <= v1)}
                                                _ = v20
                                            }
                                        default:
                                            var v21 uint32 = (v3 / ((0) | 1))
                                            _ = v21
                                        }
                                        fmt.Println(v2, v11)
                                    }
                                case "":
                                    fmt.Println(v4, fi13)
                                }
                            }
                            var v22 uint8 = (v2 | v2)
                            _ = fi13
                            _ = v22
                        }
                    case -482:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v4 = v4
                            _ = fi23
                        }
                    default:
                        v24 := []int{}
                        _ = v24
                    }
                    v25 := GenS1{F0: -451}
                    v26 := GenS1{F0: 42}
                    _ = fi12
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: -1}
                _ = v27
            } else {
                if true {
                    v1 = (0.1 - float32(0.0))
                } else {
                    var v28 bool = true
                    _ = v28
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 uint64 = 462
                    v31 := GenS3{F0: v4, F1: 2, F2: v30}
                    var v32 uint = 361
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                if (2 == 255) {
                    var v33 int = 127
                    v2 += v2
                    if true {
                        var v34 uint32 = 1
                        var v35 float64 = (-2.25 / 0.1)
                        _ = v34
                        _ = v35
                    } else {
                        var v36 string = ((v4 + (("foo bar" + v4) + "\n\t")) + (v4 + "hello"))
                        v37 := GenS4{F0: 295, F1: 42}
                        _ = v36
                        _ = v37
                    }
                    _ = v33
                } else {
                    var v38 uint32 = v3
                    v4 = "\n\t"
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := GenS5{F0: 446, F1: 1000}
                _ = fi39
                _ = v40
            }
            switch 468 {
            case 42:
                var v41 int64 = 42
                _ = v41
            case 949:
                for fi42 := 0; fi42 < 3; fi42++ {
                    v43 := []int{fi42}
                    v11 -= 100
                    fmt.Println(v2)
                    _ = fi42
                    _ = v43
                }
            case 128:
                var v44 int64 = 42
                _ = v44
            default:
                if (((true && (true || true)) || false) && false) {
                    v1 = v1
                    fmt.Println(v2, v2)
                }
            }
        }
    }
    switch 127 {
    case 342:
        var v45 string = "\n\t"
        _ = v45
    case -699:
        v46 := []int{1, 2147483647, 942, 766}
        _ = v46
    }
    fmt.Println(v1, v3)
    fmt.Println(v2)
    v11 -= (^(v11))
    v3 = v3
    fmt.Println(v3, v2)
    fmt.Println(v2, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
