package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 uint64; F1 int64; F2 int64 }

func main() {
    var v1 float32 = 1.5
    var v2 uint = 2
    var v3 string = "foo bar"
    var v4 uint64 = uint64(v2)
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            delete(v5, int(v2))
            fmt.Println(fi7)
            v8 := GenS1{F0: (^(v2)), F1: (v4 | 749)}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(v1, v4, v1)
    fmt.Println(v3, v2, v3)
    if false {
        v4 = (v4 >> ((v2) & 63))
        switch ("hello" + v3) {
        case "foo bar":
            v9 := GenS2{F0: 9223372036854775807}
            _ = v9
        case "hello":
            v4 -= 952
        case "":
            v10 := GenS2{F0: 32767}
            _ = v10
        default:
            v11 := GenS3{F0: v2, F1: 0}
            _ = v11
        }
    }
    v2 = v2
    switch "" {
    case "a":
        var v12 uint16 = 886
        _ = v12
    case "\n\t":
        _ = len(v5)
    case "hello":
        fmt.Println(v2, v1, v1)
    default:
        if false {
            if (18446744073709551615 >= v2) {
                v2 = 1
                if (!(!false)) {
                    v2 += 42
                }
            } else {
                v3 = "a"
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := v5[4294967295]
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := v5[301]
                    v16 = v14
                    if false {
                        v17 := GenS4{F0: 393, F1: -1}
                        v18 := GenS5{F0: v3, F1: -788, F2: 0}
                        _ = v17
                        _ = v18
                    } else {
                        switch v3 {
                        case "foo bar":
                            v19 := GenS6{F0: 428}
                            if false {
                                v14 = (-1 < 509)
                            } else {
                                v5[fi13] = (("a" + v3) == v3)
                                fmt.Println(v14, v2, v16)
                                v20 := GenS7{F0: (v4 % ((v4) | 1))}
                                _ = v20
                            }
                            _ = v19
                        case "a":
                            v21 := GenS8{F0: "a", F1: v4}
                            _ = v21
                        }
                        var v22 uint8 = 2
                        v5[(^(fi15))] = true
                        _ = v22
                    }
                    _ = fi15
                    _ = v16
                }
                v23 := []int{(fi13 / (((fi13 - (fi13 ^ (^(fi13))))) | 1)), -881, fi13}
                _ = fi13
                _ = v14
                _ = v23
            }
            switch "foo bar" {
            case "\n\t":
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint64 = (^(v4))
                    var v26 uint8 = 0
                    v5[4294967295] = false
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
    }
    v4 = 0
    v27 := GenS5{F0: v3, F1: 455, F2: 0}
    var v28 bool = (246 < 561)
    if true {
        if (!v28) {
            v29 := GenS5{F0: "\n\t", F1: 42, F2: 32767}
            _ = v29
        } else {
            var v30 bool = v28
            _ = v30
        }
        switch 32767 {
        case 204:
            if v28 {
                v31 := []int{4294967295}
                var v32 int8 = 1
                delete(v5, 32768)
                _ = v31
                _ = v32
            }
        default:
            var v33 int16 = 32767
            _ = v33
        }
        if v28 {
            v34 := GenS6{F0: 0}
            var v35 uint32 = uint32(v4)
            _ = v34
            _ = v35
        } else {
            var v36 float64 = -2.25
            v37 := GenS9{F0: (v4 | 2), F1: 678, F2: 136}
            var v38 uint16 = 32767
            _ = v36
            _ = v37
            _ = v38
        }
    }
    fmt.Println(v3)
    var v39 float32 = 0.1
    _ = len(v5)
    var v40 string = (("hello" + ("hello" + (v3 + ((((((v3 + (v3 + ((v3 + (v3 + ("foo bar" + "hello"))) + (("" + (((((v3 + (((v3 + (v3 + "foo bar")) + "") + v3)) + "a") + v3) + "a") + v3)) + (("\n\t" + v3) + v3))))) + v3) + ("\n\t" + v3)) + v3) + (v3 + ((v3 + v3) + v3))) + (("a" + v3) + (v3 + v3)))))) + "")
    if true {
        fmt.Println(v2, v28)
        v41 := []int{-2147483648, 100}
        for fi42 := 0; fi42 < 4; fi42++ {
            var v43 float32 = -2.25
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        if (461 >= 741) {
            for fi45 := 0; fi45 < 2; fi45++ {
                v1 -= 0.1
                var v46 uint64 = v4
                _ = fi45
                _ = v46
            }
        }
        v4 = v4
        _ = fi44
    }
    fmt.Println(v28, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
