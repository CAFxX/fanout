package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (50 >= 0)
    var v3 uint16 = 954
    var v4 int32 = 1
    var v5 uint = 1
    if v2 {
        switch -1 {
        case 9223372036854775807:
            if v2 {
                for fi6 := 0; fi6 < 3; fi6++ {
                    v2 = true
                    _ = fi6
                }
                v3 -= (v3 / (((^(v3))) | 1))
                v7 := GenS1{F0: 912, F1: 32767, F2: 255}
                _ = v7
            }
        default:
            var v8 bool = false
            _ = v8
        }
    }
    v4 = (v4 ^ v4)
    v9 := GenS1{F0: v5, F1: -233, F2: 412}
    var v10 uint64 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        if (v2 && (v2 || false)) {
            switch ("foo bar" + "foo bar") {
            case "":
                switch "" {
                case "foo bar":
                    if v2 {
                        v4 = v4
                        var v12 float32 = 3.14159
                        if v2 {
                            var v13 int8 = 0
                            if false {
                                v12 = (v12 + v12)
                            } else {
                                v4 -= (v4 + 0)
                                var v14 int32 = (v4 | (v4 << ((v5) & 63)))
                                if (v2 || (v5 <= (^(v5)))) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v15 := GenS2{F0: 9223372036854775807, F1: fi11}
                                        _ = v15
                                    case "foo bar":
                                        var v16 int = fi11
                                        _ = v16
                                    default:
                                        v5 += (v5 % ((1) | 1))
                                    }
                                    var v17 int16 = 42
                                    for fi18 := 0; fi18 < 15; fi18++ {
                                        var v19 float32 = (0.0 - v12)
                                        var v20 int64 = -1
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v17
                                }
                                _ = v14
                            }
                            _ = v13
                        } else {
                            v5 = uint(v10)
                            for fi21 := 0; fi21 < 7; fi21++ {
                                switch (fi21 - (fi11 & 827)) {
                                case -2147483648:
                                    var v22 int64 = -1
                                    _ = v22
                                case 502:
                                    _ = len(v1)
                                }
                                _ = fi21
                            }
                            if v2 {
                                var v23 uint16 = (v3 & (v3 | v3))
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    var v25 int64 = 1
                                    v5 += v5
                                    var v26 int64 = (v25 * int64(v4))
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                _ = v23
                            }
                        }
                        _ = v12
                    }
                case "\n\t":
                    var v27 int16 = int16(v10)
                    _ = v27
                case "hello":
                    _ = len(v1)
                default:
                    v28 := []int{(fi11 * (fi11 - 2147483647)), -9223372036854775808, (-(fi11)), (fi11 | fi11)}
                    _ = v28
                }
            case "foo bar":
                if ((-480 < -1) && (3.14159 < 3.14159)) {
                    v1[0] = fi11
                    var v29 int8 = -1
                    _ = v29
                }
            }
        }
        _ = fi11
    }
    v30 := GenS2{F0: 42, F1: 929}
    switch 4294967295 {
    case 128:
        switch "a" {
        case "foo bar":
            v31 := GenS2{F0: 42, F1: -194}
            _ = v31
        default:
            v32 := []int{}
            _ = v32
        }
    case -49:
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := GenS2{F0: 1, F1: (fi33 / ((int(v3)) | 1))}
            _ = fi33
            _ = v34
        }
    }
    var v35 uint16 = 2
    v36 := GenS1{F0: (v5 & 42), F1: 42, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
