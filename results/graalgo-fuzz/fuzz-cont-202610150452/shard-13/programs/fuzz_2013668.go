package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int16 }
func genFn0() int16 {
    var v6 int8 = 100
    v6 += (v6 + v6)
    v6 -= v6
    var v7 float32 = -2.25
    if false {
        v8 := []int{2147483647, 4294967295}
        v9 := GenS1{F0: 51, F1: 42, F2: 0}
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int8 = v6
            var v12 uint = 0
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint = 42
            _ = fi13
            _ = v14
        }
    }
    _ = v6
    _ = v7
    return -32768
}
func genFn1() int {
    v15 := []int{127, 4294967295, -2147483648, 15}
    v16 := GenS2{F0: 0, F1: 255}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = v15
    _ = v16
    return 2147483647
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    if false {
        v1 += 4294967295
        var v17 int16 = -527
        if false {
            if false {
                switch "" {
                case "foo bar":
                    v2 -= v2
                case "\n\t":
                    v1 -= (v1 + 2)
                case "a":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 int64 = 76
                        if true {
                            fmt.Println(v19)
                            switch 2147483647 {
                            case 0:
                                switch "hello" {
                                case "":
                                    v20 := GenS3{F0: 2}
                                    _ = v20
                                case "hello":
                                    var v21 uint32 = 2147483647
                                    _ = v21
                                case "\n\t":
                                    v22 := []int{(-(fi18)), (fi18 + (fi18 - fi18)), int(v1), (fi18 & (fi18 ^ (fi18 * 32768)))}
                                    _ = v22
                                default:
                                    v2 = (v2 & v2)
                                }
                            case 128:
                                v1 = v1
                            case 581:
                                fmt.Println(v1, v1)
                            default:
                                if true {
                                    v3[434] = true
                                } else {
                                    v23 := []int{fi18, -1, fi18}
                                    for fi24 := 0; fi24 < 13; fi24++ {
                                        v17 -= (v17 * genFn0())
                                        _ = fi24
                                    }
                                    var v25 string = "\n\t"
                                    _ = v23
                                    _ = v25
                                }
                            }
                            switch (^(fi18)) {
                            case -747:
                                var v26 float32 = 0.0
                                _ = v26
                            case 52:
                                var v27 int16 = (v17 ^ v17)
                                _ = v27
                            default:
                                v28 := GenS2{F0: v1, F1: 128}
                                _ = v28
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                var v29 float32 = (-2.25 * (float32(3.14159) - -2.25))
                v1 += 1
                _ = v29
            }
        } else {
            v30 := GenS4{F0: 2147483647}
            v31 := GenS5{F0: -641, F1: 42}
            var v32 int = 128
            _ = v30
            _ = v31
            _ = v32
        }
        _ = v17
    } else {
        fmt.Println(v1, v2)
        delete(v3, uint32(v2))
        fmt.Println(v2)
    }
    v2 += v2
    v2 += v2
    v1 = (v1 ^ 42)
    if (v2 <= (v2 & v2)) {
        var v33 int64 = -1
        _ = v33
    } else {
        for fi34 := 0; fi34 < 11; fi34++ {
            if true {
                v1 -= 1
                if false {
                    v2 -= v2
                    var v35 int8 = 1
                    _ = v35
                } else {
                    switch "\n\t" {
                    case "a":
                        v36 := GenS5{F0: (-(fi34)), F1: 42}
                        _ = v36
                    case "":
                        v2 -= (v2 ^ v2)
                    default:
                        v2 = v2
                    }
                    fmt.Println(v1, v1, fi34)
                    v37 := GenS6{F0: 42}
                    _ = v37
                }
                for fi38 := 0; fi38 < 4; fi38++ {
                    v39 := GenS3{F0: v1}
                    var v40 uint = v2
                    v1 = v1
                    v2 = 0
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            } else {
                v41 := GenS1{F0: 18446744073709551615, F1: 0, F2: 1}
                _ = v41
            }
            v42 := GenS6{F0: genFn0()}
            _ = fi34
            _ = v42
        }
        v2 += v2
    }
    v43 := GenS5{F0: 127, F1: -1}
    v44 := GenS3{F0: (v1 - 1)}
    switch 2147483647 {
    case -2147483648:
        if ("hello" != "a") {
            for fi45 := 0; fi45 < 16; fi45++ {
                v1 += (v1 % ((v1) | 1))
                var v46 int64 = 1
                _ = fi45
                _ = v46
            }
        } else {
            v47 := []int{9223372036854775807, 127}
            for fi48 := 0; fi48 < 15; fi48++ {
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
                _ = fi48
            }
            _ = v47
        }
    case 616:
        fmt.Println(v2, v1)
    }
    v2 -= v2
    v5[0] = 9223372036854775807
    for fi49 := 0; fi49 < 4; fi49++ {
        v50 := GenS6{F0: 1}
        _ = fi49
        _ = v50
    }
    var v51 uint = v2
    switch 131 {
    case -438:
        _ = len(v5)
    default:
        v2 -= (v51 ^ 266)
    }
    v2 = 397
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v51
}
