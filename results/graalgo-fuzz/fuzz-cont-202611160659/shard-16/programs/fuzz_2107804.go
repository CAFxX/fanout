package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 int16, p2 bool) uint16 {
    p1 = p1
    p2 = (p0 >= p0)
    if (18446744073709551615 >= 0) {
        if p2 {
            if p2 {
                fmt.Println(p1)
            }
        }
        var v6 uint64 = 785
        _ = v6
    } else {
        switch (p0 % (((-(p0))) | 1)) {
        case -375:
            v7 := []int{}
            _ = v7
        case -2147483648:
            p1 = p1
        case 0:
            p0 -= (p0 / (((p0 << ((2) & 63))) | 1))
        default:
            p2 = true
        }
    }
    switch 0 {
    case 127:
        if (-588 == 42) {
            var v8 string = ("a" + "foo bar")
            var v9 uint32 = 65535
            for fi10 := 0; fi10 < 12; fi10++ {
                switch v8 {
                case "\n\t":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 uint64 = 0
                        _ = fi11
                        _ = v12
                    }
                case "":
                    v8 = v8
                case "a":
                    if p2 {
                        var v13 bool = (-17 < (p0 + fi10))
                        if ((!(false || (v9 > 65535))) && p2) {
                            v14 := GenS1{F0: 42, F1: 127}
                            _ = v14
                        } else {
                            p1 -= p1
                            v15 := []int{(p0 - p0), (p0 ^ int(v9))}
                            v16 := GenS2{F0: (p0 ^ (fi10 - (fi10 | (p0 & p0)))), F1: "a"}
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    }
                default:
                    var v17 float32 = 3.14159
                    _ = v17
                }
                var v18 int8 = 0
                _ = fi10
                _ = v18
            }
            _ = v8
            _ = v9
        } else {
            switch (^(p0)) {
            case -9223372036854775808:
                p0 -= (^(p0))
            default:
                p2 = false
            }
        }
    case 144:
        v19 := GenS3{F0: 0}
        _ = v19
    case 664:
        p1 = 42
    default:
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                p1 -= p1
                _ = fi21
            }
            var v22 uint16 = 32767
            _ = fi20
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int16, p1 uint64, p2 bool) uint64 {
    fmt.Println(p0, p1, p2)
    var v23 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return (p1 + (p1 ^ (^(p1))))
}

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2147483648
    v3 := []string{("\n\t" + "hello")}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    var v24 int32 = 2147483647
    var v25 bool = (1 >= 1)
    v1 += v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v26 := GenS2{F0: 127, F1: (("a" + "a") + "foo bar")}
    v27 := GenS1{F0: -1, F1: 42}
    _ = len(v3)
    v2 -= (v2 & 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
