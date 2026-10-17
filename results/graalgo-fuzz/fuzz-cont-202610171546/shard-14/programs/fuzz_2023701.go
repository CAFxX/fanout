package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0() bool {
    var v6 uint8 = 86
    v6 = (v6 / ((v6) | 1))
    var v7 uint = 4294967295
    v7 = v7
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 int16 = 1
    var v5 int = -1
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int = fi8
        var v10 float32 = (3.14159 - 0.0)
        v11 := GenS1{F0: ""}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v1 -= ((float64(((-2.25 / float32((0.1 - v1))) / 1.5)) - 0.1) / (v1 - ((v1 - -2.25) - (v1 - (0.0 * (3.14159 + v1))))))
    if true {
        v5 = v5
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v4)
            switch (v5 ^ fi12) {
            case -2147483648:
                if (!(65535 > 4294967295)) {
                    v13 := []int{(fi12 - v5), v5}
                    var v14 bool = false
                    v15 := GenS1{F0: "\n\t"}
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                v4 -= (v4 & -199)
            }
            _ = len(v2)
            _ = fi12
        }
        fmt.Println(v1)
    }
    v4 -= (v4 + 0)
    _ = len(v2)
    v4 += (^(v4))
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint64 = 42
        v5 = v5
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: 4294967295, F1: 9223372036854775807}
        var v20 float64 = v1
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        if (-459 >= 42) {
            v1 -= 1e10
            fmt.Println(v4, v5, fi21)
            var v22 string = "\n\t"
            _ = v22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := []int{(^(fi23)), (fi23 * (fi23 >> ((21) & 63))), (v5 ^ fi23), v5}
        if (444 != 4294967295) {
            v25 := GenS2{F0: 65535, F1: 9223372036854775807}
            var v26 int = 4294967295
            switch "foo bar" {
            case "foo bar":
                var v27 uint16 = 32767
                _ = v27
            case "":
                switch -775 {
                case 2147483647:
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v5 = (v5 & (v5 | (v5 & (-(v26)))))
                        v29 := GenS3{F0: 1}
                        switch "a" {
                        case "a":
                            if false {
                                fmt.Println(v4, v5)
                                v30 := GenS4{F0: 65535, F1: (v4 % ((v4) | 1))}
                                _ = v30
                            }
                        case "hello":
                            fmt.Println(v5)
                        case "foo bar":
                            var v31 float32 = 1.5
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            default:
                if false {
                    v3[v5] = 65535
                    var v32 int32 = 42
                    _ = v32
                }
            }
            _ = v25
            _ = v26
        } else {
            v5 = v5
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            fmt.Println(fi23, v4)
            var v34 uint16 = 447
            v35 := GenS4{F0: 1, F1: 42}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi23
        _ = v24
    }
    v5 -= (v5 * (v5 / ((v5) | 1)))
    var v36 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
