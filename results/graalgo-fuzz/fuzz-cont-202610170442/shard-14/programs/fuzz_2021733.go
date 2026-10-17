package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 int32 = -1
    var v2 int64 = -1
    v3 := make(map[int]int)
    var v4 int8 = 0
    v4 -= v4
    var v5 int64 = (-(v2))
    var v6 uint8 = 219
    switch ("foo bar" + (("hello" + "a") + "\n\t")) {
    case "\n\t":
        v7 := []int{127, -752}
        _ = v7
    case "hello":
        v4 += 119
    }
    v8 := GenS1{F0: v6}
    fmt.Println(v2)
    if (-1 > v4) {
        v3[127] = -802
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := v3[(fi9 % ((fi10) | 1))]
                fmt.Println(v1, v6, v5)
                _ = fi10
                _ = v11
            }
            var v12 uint = 255
            _ = fi9
            _ = v12
        }
    } else {
        v13 := GenS2{F0: true}
        v2 = (v2 ^ (^(v5)))
        v5 = int64(v4)
        _ = v13
    }
    var v14 bool = true
    var v15 uint32 = 2
    v16 := GenS2{F0: v14}
    if (v14 || false) {
        if ((v15 << ((1) & 63)) < v15) {
            if v14 {
                if true {
                    switch -469 {
                    case 128:
                        switch -742 {
                        case -1:
                            var v17 string = (("a" + "a") + "a")
                            _ = v17
                        }
                    case 1:
                        fmt.Println(v6, v1, v6)
                    }
                    if ((1.5 + 1.5) != 0.1) {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            if (!false) {
                                v3[(fi18 % ((fi18) | 1))] = (-(fi18))
                            } else {
                                fmt.Println(v5, v4, v5)
                            }
                            var v19 uint16 = 255
                            switch "\n\t" {
                            case "hello":
                                var v20 uint = 65535
                                _ = v20
                            case "":
                                var v21 float64 = (((1.5 - 1.5) + (1.5 - -2.25)) / (-2.25 / 1e10))
                                _ = v21
                            default:
                                var v22 uint16 = 1
                                _ = v22
                            }
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        switch 608 {
                        case 2147483647:
                            switch 32767 {
                            case 9223372036854775807:
                                var v23 uint = 488
                                _ = v23
                            case 826:
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    v6 = v6
                                    var v25 float32 = 0.1
                                    _ = fi24
                                    _ = v25
                                }
                            case 2147483647:
                                var v26 uint = 0
                                _ = v26
                            default:
                                var v27 string = "hello"
                                _ = v27
                            }
                        case 42:
                            for fi28 := 0; fi28 < 15; fi28++ {
                                v29 := GenS3{F0: 128, F1: (^(v5))}
                                v4 = v4
                                v15 = v15
                                _ = fi28
                                _ = v29
                            }
                        default:
                            v30 := []int{-9223372036854775808}
                            _ = v30
                        }
                    }
                    v31 := GenS1{F0: 1}
                    _ = v31
                } else {
                    var v32 uint64 = 351
                    v15 -= (v15 >> ((42) & 63))
                    _ = v32
                }
                var v33 int64 = v2
                _ = v33
            } else {
                var v34 int64 = (v5 / ((v2) | 1))
                _ = v34
            }
            for fi35 := 0; fi35 < 6; fi35++ {
                _ = len(v3)
                _ = fi35
            }
            if v14 {
                var v36 int8 = v4
                _ = v36
            }
        }
    } else {
        v15 = 42
        v4 -= -98
    }
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
