package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint16; F1 string; F2 uint8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint64, p1 int8, p2 int) int64 {
    switch "foo bar" {
    case "":
        v7 := []int{}
        _ = v7
    case "foo bar":
        p1 += (p1 - p1)
    default:
        var v8 int32 = -1
        _ = v8
    }
    switch "hello" {
    case "a":
        v9 := GenS1{F0: 1, F1: 0, F2: 852}
        _ = v9
    }
    fmt.Println(p2, p1, p2)
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint32 = 4294967295
        if true {
            fmt.Println(p0, p0, fi10)
            if true {
                var v12 int = (fi10 / ((fi10) | 1))
                v13 := GenS2{F0: (-(p2))}
                _ = v12
                _ = v13
            }
            p1 = p1
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return int64(p1)
}
func genFn1(p0 int16) int64 {
    v14 := GenS3{F0: 2147483647}
    p0 = (p0 * 32767)
    fmt.Println(p0)
    v15 := GenS3{F0: -707}
    _ = p0
    _ = v14
    _ = v15
    return 647
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 136
    v3 := make(map[int]string)
    var v4 int32 = 26
    v5 := make(map[string]bool)
    var v6 uint64 = 915
    v6 = v6
    var v16 uint16 = 42
    var v17 int16 = 427
    v18 := GenS4{F0: ((v6 | v6) != (v6 % ((0) | 1))), F1: 324}
    v2 = 232
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 float32 = (-2.25 - 0.1)
            v22 := []int{fi20, int(v17)}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    delete(v3, 9223372036854775807)
    v23 := GenS5{F0: (v16 + (v16 % ((v16) | 1))), F1: "hello", F2: (v2 << ((18446744073709551615) & 63))}
    for fi24 := 0; fi24 < 5; fi24++ {
        if ("a" > "a") {
            fmt.Println(v2, v4, v16)
            for fi25 := 0; fi25 < 5; fi25++ {
                v6 += (^(v6))
                v16 = (^(v16))
                switch ("a" + "a") {
                case "\n\t":
                    v16 = 1
                case "a":
                    switch (fi24 / ((fi25) | 1)) {
                    case -1:
                        if (42 <= int8(fi24)) {
                            var v26 float32 = 3.14159
                            _ = v26
                        } else {
                            v2 -= 2
                            switch "\n\t" {
                            case "":
                                if (-2.25 < 0.0) {
                                    v27 := GenS6{F0: (v17 / (((v17 << ((65535) & 63))) | 1))}
                                    _ = v27
                                }
                            }
                            fmt.Println(fi25)
                        }
                    }
                case "":
                    var v28 int64 = 221
                    _ = v28
                default:
                    switch fi25 {
                    case 554:
                        v29 := v3[int(v17)]
                        _ = v29
                    default:
                        var v30 float64 = 3.14159
                        _ = v30
                    }
                }
                _ = fi25
            }
            v16 += v16
        }
        if (true || false) {
            var v31 uint64 = (^(v6))
            fmt.Println(v6, v6, v4)
            var v32 float64 = float64((0.1 + 1e10))
            _ = v31
            _ = v32
        } else {
            var v33 uint = 768
            v34 := GenS2{F0: 42}
            _ = v33
            _ = v34
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
