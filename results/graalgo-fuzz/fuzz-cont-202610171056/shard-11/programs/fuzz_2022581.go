package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32 }
func genFn0(p0 uint, p1 bool, p2 string) uint {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 1}
        var v8 uint32 = 636
        p1 = (((fi6 - -2147483648) > (fi6 % ((fi6) | 1))) && (4294967295 <= 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}
func genFn1(p0 uint) int32 {
    p0 = (p0 ^ p0)
    var v10 int = 42
    var v11 uint64 = 0
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint16 = 1
    var v2 bool = true
    var v3 float64 = float64(1.5)
    var v4 int32 = -1
    v5 := make(map[uint32]int)
    v12 := v5[1]
    for fi13 := 0; fi13 < 15; fi13++ {
        v3 -= -2.25
        _ = fi13
    }
    switch (v12 / ((int(v4)) | 1)) {
    case 4294967295:
        v14 := v5[2]
        _ = v14
    case 9223372036854775807:
        var v15 string = ("a" + ("" + "a"))
        _ = v15
    default:
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: 770}
            fmt.Println(fi16, v4)
            _ = fi16
            _ = v17
        }
    }
    v2 = true
    fmt.Println(v3)
    fmt.Println(v2)
    var v18 uint64 = 783
    switch v12 {
    case 9223372036854775807:
        switch (("" + "hello") + "hello") {
        case "":
            switch (("a" + ("foo bar" + "")) + "\n\t") {
            case "\n\t":
                var v19 uint16 = (v1 % ((v1) | 1))
                _ = v19
            case "":
                var v20 uint64 = v18
                _ = v20
            case "foo bar":
                if ("hello" > "foo bar") {
                    switch "a" {
                    case "\n\t":
                        v21 := GenS2{F0: int(v4)}
                        _ = v21
                    case "foo bar":
                        v22 := []int{v12}
                        _ = v22
                    default:
                        v12 += (v12 + v12)
                    }
                    fmt.Println(v3)
                    fmt.Println(v18)
                } else {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 int = 9223372036854775807
                        _ = fi23
                        _ = v24
                    }
                    if (v2 && v2) {
                        var v25 int64 = 0
                        _ = v25
                    } else {
                        v12 += v12
                        v26 := GenS3{F0: (v12 << ((755) & 63)), F1: genFn1(4294967295)}
                        _ = v26
                    }
                }
            default:
                fmt.Println(v4, v18, v18)
            }
        case "\n\t":
            var v27 bool = (1e10 == 3.14159)
            _ = v27
        case "a":
            if (!v2) {
                v28 := []int{(v12 << ((2) & 63))}
                for fi29 := 0; fi29 < 5; fi29++ {
                    switch 4294967295 {
                    case 1:
                        var v30 uint8 = 237
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 7; fi31++ {
                            var v32 uint8 = 1
                            if len(v28) > 0 {
                                v28[0] = 127
                            }
                            _ = fi31
                            _ = v32
                        }
                    }
                    switch "" {
                    case "hello":
                        var v33 string = "a"
                        _ = v33
                    }
                    fmt.Println(v4)
                    _ = fi29
                }
                _ = len(v28)
                _ = v28
            } else {
                switch "foo bar" {
                case "":
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v12 = (v12 * fi34)
                        _ = fi34
                    }
                case "hello":
                    var v35 float32 = -2.25
                    _ = v35
                }
                v36 := v5[744]
                v18 -= (^(v18))
                _ = v36
            }
        default:
            v2 = (171 > 659)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
