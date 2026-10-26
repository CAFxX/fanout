package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 374
    var v3 float32 = 0.1
    var v4 int16 = 0
    fmt.Println(v4, v3)
    v2 += (v2 + 0)
    var v5 int8 = 1
    if true {
        if false {
            fmt.Println(v2, v2, v1)
            v5 = int8(v1)
        } else {
            switch -9223372036854775808 {
            case -2147483648:
                v3 += float32(-2.25)
            case 195:
                v2 = (^(v2))
            case 543:
                var v6 int = 242
                _ = v6
            default:
                v3 = 1e10
            }
        }
        v1 = (v1 ^ v1)
        var v7 string = ("" + "foo bar")
        _ = v7
    }
    v8 := GenS1{F0: "foo bar", F1: -863}
    if true {
        if false {
            var v9 int64 = 1
            if true {
                fmt.Println(v3)
            } else {
                switch 128 {
                case 128:
                    v3 = v3
                case 1:
                    switch 0 {
                    case 251:
                        for fi10 := 0; fi10 < 7; fi10++ {
                            v11 := GenS1{F0: "foo bar", F1: (-(fi10))}
                            v2 += v2
                            _ = fi10
                            _ = v11
                        }
                    case -2147483648:
                        var v12 float64 = 0.0
                        _ = v12
                    }
                }
                var v13 float32 = v3
                _ = v13
            }
            v4 = (v4 ^ -1)
            _ = v9
        } else {
            v5 += v5
            v3 = 3.14159
            v1 += v1
        }
        v14 := []int{}
        _ = len(v14)
        _ = v14
    }
    v15 := []int{}
    if false {
        if true {
            var v16 bool = true
            var v17 uint32 = v1
            var v18 float64 = 0.1
            _ = v16
            _ = v17
            _ = v18
        } else {
            if false {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v4 = 1
                    var v20 float64 = 0.0
                    if (!(1 == 2)) {
                        fmt.Println(v5)
                        var v21 uint64 = 138
                        if (!((v21 | v21) != 38)) {
                            v3 += v3
                        } else {
                            v3 += v3
                            v22 := GenS2{F0: (fi19 % ((808) | 1))}
                            _ = v22
                        }
                        _ = v21
                    } else {
                        switch "foo bar" {
                        case "":
                            v15 = append(v15, fi19)
                        case "\n\t":
                            v3 = (3.14159 - (v3 * v3))
                        }
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                if false {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        fmt.Println(v1, v1, v2)
                        _ = fi23
                    }
                } else {
                    v5 = v5
                }
                v24 := GenS3{F0: 42, F1: (v1 % (((v1 * (v1 * uint32(v2)))) | 1)), F2: v2}
                _ = len(v15)
                _ = v24
            }
        }
        if (!false) {
            if (0.0 > -2.25) {
                var v25 int16 = (v4 / ((v4) | 1))
                _ = v25
            }
            v4 = (v4 & -1)
        }
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            v5 -= 127
            _ = fi26
        }
        v1 -= v1
        v27 := GenS1{F0: (("a" + "") + ("" + (("a" + "foo bar") + "hello"))), F1: 9223372036854775807}
        _ = v27
    }
    var v28 float32 = (v3 + float32(0.1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
