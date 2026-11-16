package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint }
type GenS3 struct { F0 bool; F1 int8 }

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 419
    var v3 uint = 4294967295
    var v4 string = "\n\t"
    var v5 float64 = 1e10
    var v6 int8 = -119
    switch 2147483647 {
    case 32767:
        fmt.Println(v3, v3, v3)
    case 4294967295:
        var v7 bool = (198 == 153)
        _ = v7
    case -33:
        switch "a" {
        case "\n\t":
            v8 := []int{-1, 9223372036854775807}
            _ = v8
        case "a":
            v3 = (v3 * v3)
        default:
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v1, fi9, v4)
                v10 := []int{fi9, (fi9 / ((fi9) | 1)), (fi9 ^ int(v2)), (fi9 * 1)}
                _ = fi9
                _ = v10
            }
        }
    }
    fmt.Println(v6)
    switch v4 {
    case "":
        v4 = "hello"
    default:
        v11 := []int{2147483647, -499, 128}
        _ = v11
    }
    v12 := GenS1{F0: 991, F1: v6, F2: (v6 * (v6 % ((v6) | 1)))}
    var v13 string = v4
    v14 := GenS2{F0: (v1 << ((v3) & 63)), F1: 127, F2: v3}
    v13 = v4
    v3 = v3
    switch 127 {
    case -2147483648:
        switch v13 {
        case "":
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 1, F1: (v6 - 127), F2: (v6 & 100)}
                v17 := GenS3{F0: (((1.5 - (0.0 * -2.25)) > 1e10) && false), F1: v6}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case -150:
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint64 = v2
            var v20 int = (-(fi18))
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int32 = -2147483648
            if false {
                var v23 bool = true
                if v23 {
                    v24 := []int{42, (fi21 * (fi21 - (^(fi21)))), 1, fi21}
                    v6 += v6
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        var v26 float64 = (v5 / 1.5)
                        v27 := []int{}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v28 := []int{int(v2)}
                    _ = v28
                }
                if v23 {
                    if (!(42 < 9223372036854775807)) {
                        v13 = ("hello" + v13)
                    }
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v30 := []int{4294967295, (fi21 | fi29), -1}
                        _ = fi29
                        _ = v30
                    }
                    var v31 uint64 = (v2 >> ((1) & 63))
                    _ = v31
                } else {
                    switch int(v3) {
                    case 4294967295:
                        v3 += v3
                    case -534:
                        switch (-(fi21)) {
                        case 0:
                            v3 = v3
                        case -1:
                            v32 := []int{}
                            _ = v32
                        }
                    default:
                        v33 := []int{}
                        _ = v33
                    }
                    v6 = v6
                    v23 = v23
                }
                _ = v23
            } else {
                fmt.Println(fi21, v1, v5)
                var v34 string = "hello"
                _ = v34
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
