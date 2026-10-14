package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 int; F1 int32; F2 int64 }

func main() {
    var v1 int32 = 2147483647
    var v2 string = (("" + "a") + "")
    var v3 uint64 = 113
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    fmt.Println(v3)
    if (false && true) {
        switch 32768 {
        case 1:
            for fi6 := 0; fi6 < 5; fi6++ {
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := []int{}
                    v3 = 4294967295
                    v2 = (("\n\t" + "a") + v2)
                    _ = fi7
                    _ = v8
                }
                v3 = 42
                _ = fi6
            }
        case 9223372036854775807:
            v9 := []int{1, -163, -441}
            _ = v9
        }
    }
    if true {
        v2 = "hello"
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: 127, F1: fi10, F2: 42}
            for fi12 := 0; fi12 < 3; fi12++ {
                v2 = ("" + (v2 + v2))
                v2 = v2
                _ = fi12
            }
            if (0.1 >= -2.25) {
                v13 := []int{(fi10 >> ((65535) & 63)), 42, fi10}
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        var v14 uint64 = v3
        _ = v14
    } else {
        if (1 <= 255) {
            v15 := GenS1{F0: 13, F1: -1, F2: (v1 & (v1 + v1))}
            _ = v15
        }
        var v16 bool = true
        fmt.Println(v1, v16, v3)
        _ = v16
    }
    var v17 float64 = 1.5
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int64 = 0
        for fi20 := 0; fi20 < 18; fi20++ {
            switch (fi20 << ((4294967295) & 63)) {
            case -831:
                switch v2 {
                case "\n\t":
                    v21 := v4[(fi18 * fi20)]
                    _ = v21
                case "a":
                    fmt.Println(v17, fi20, v17)
                case "foo bar":
                    v22 := []int{(^(fi18)), (^(fi18)), (fi18 / (((fi20 + fi20)) | 1))}
                    _ = v22
                default:
                    var v23 int16 = 1000
                    _ = v23
                }
            default:
                v24 := v4[9223372036854775807]
                _ = v24
            }
            switch (-(fi20)) {
            case 32768:
                fmt.Println(v17, v17, fi20)
            case -887:
                if true {
                    if true {
                        var v25 int16 = 232
                        if false {
                            if (0 >= 0) {
                                v26 := GenS2{F0: (fi20 + (fi18 * fi20)), F1: v1, F2: (v19 >> ((18446744073709551615) & 63))}
                                v17 -= 1e10
                                v27 := v5["hello"]
                                _ = v26
                                _ = v27
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        fmt.Println(v2, fi20, v2)
                        if (1 <= -45) {
                            fmt.Println(v2, v1, v1)
                            v3 += v3
                            v28 := GenS1{F0: 37, F1: fi18, F2: 188}
                            _ = v28
                        }
                        _ = v25
                    }
                    if false {
                        fmt.Println(v19, fi20, v1)
                        fmt.Println(v2)
                    }
                    var v29 int8 = 0
                    _ = v29
                }
            case 42:
                v19 += (v19 / (((v19 * (-(v19)))) | 1))
            default:
                v30 := []int{}
                _ = v30
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    v2 = "\n\t"
    var v31 int = 0
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int32 = (v1 % ((v1) | 1))
        if false {
            v2 = v2
        } else {
            switch v2 {
            case "hello":
                for fi34 := 0; fi34 < 2; fi34++ {
                    v35 := GenS2{F0: (fi34 - (^(fi34))), F1: (v33 ^ -1), F2: 9223372036854775807}
                    switch v2 {
                    case "hello":
                        v36 := v5[v2]
                        _ = v36
                    case "a":
                        v37 := []int{fi34, (fi34 >> ((255) & 63)), -74, v31}
                        _ = v37
                    default:
                        v31 = fi34
                    }
                    _ = fi34
                    _ = v35
                }
            case "a":
                fmt.Println(v1, v17, v33)
            default:
                var v38 int = (fi32 - (fi32 | fi32))
                _ = v38
            }
        }
        v1 = v33
        _ = fi32
        _ = v33
    }
    var v39 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
}
