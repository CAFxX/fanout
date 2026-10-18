package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int16 }

func main() {
    var v1 float64 = (1.5 * ((0.0 - 0.1) - 3.14159))
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 2147483647
    var v6 string = "\n\t"
    v2 = (0.1 - v2)
    if (-9223372036854775808 < 42) {
        var v7 int16 = 690
        _ = v7
    } else {
        if (!(((((1e10 / (0.0 + 1e10)) / 0.1) + 1e10) - (1e10 - 1e10)) <= 1.5)) {
            if (false || (("a" + "hello") >= (v6 + ("hello" + ("" + ("" + v6)))))) {
                if false {
                    v8 := GenS1{F0: -2147483648, F1: (v6 + v6)}
                    for fi9 := 0; fi9 < 19; fi9++ {
                        for fi10 := 0; fi10 < 7; fi10++ {
                            fmt.Println(v6)
                            _ = fi10
                        }
                        _ = fi9
                    }
                    _ = v8
                }
                delete(v4, (v5 + 2147483648))
            }
            v6 = v6
        }
        if false {
            var v11 bool = true
            v12 := GenS1{F0: 9223372036854775807, F1: v6}
            _ = v11
            _ = v12
        }
        if (-32768 <= 271) {
            v13 := GenS1{F0: 32768, F1: "foo bar"}
            _ = v13
        } else {
            switch 128 {
            case 4294967295:
                v14 := []int{32767, 42, -762}
                _ = v14
            case 415:
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 int32 = -1
                    _ = fi15
                    _ = v16
                }
            case 852:
                switch 0 {
                case -2147483648:
                    v17 := []int{42, -202}
                    _ = v17
                case 704:
                    switch 32767 {
                    case 2147483647:
                        v18 := GenS1{F0: -710, F1: v6}
                        _ = v18
                    }
                case 0:
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v2 += (1e10 - 1e10)
                        _ = fi19
                    }
                default:
                    var v20 float64 = v2
                    _ = v20
                }
            }
            if false {
                v21 := GenS1{F0: 4294967295, F1: (v6 + "foo bar")}
                _ = v21
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                fmt.Println(v1)
                _ = fi22
            }
        }
    }
    v23 := []int{}
    switch 491 {
    case 0:
        v2 += v1
    default:
        fmt.Println(v5)
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            if len(v23) > 0 {
                v23[0] = int(v5)
            }
            v26 := GenS2{F0: 4294967295, F1: -2147483648}
            v27 := GenS3{F0: 967}
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            v2 = ((v1 * v2) * v2)
            v2 -= float64(1e10)
            v2 += 1.5
            _ = fi28
        }
        if true {
            var v29 uint64 = 243
            _ = v29
        } else {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int8 = 1
                if true {
                    v5 = v5
                    v32 := GenS4{F0: 2}
                    v31 = v31
                    _ = v32
                } else {
                    delete(v4, (^(v5)))
                    var v33 uint16 = 852
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        }
        _ = fi24
    }
    switch v6 {
    case "\n\t":
        v34 := GenS5{F0: 42}
        _ = v34
    case "a":
        v35 := GenS6{F0: int(v5)}
        _ = v35
    }
    v6 = (v6 + (((v6 + "") + ("" + v6)) + "hello"))
    v36 := GenS7{F0: 100, F1: 42}
    v3[-774] = (2 >= 255)
    fmt.Println(v6, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
