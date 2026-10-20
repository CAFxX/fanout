package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 uint64, p1 string, p2 uint8) uint64 {
    if (!((!false) && false)) {
        if false {
            var v7 uint64 = (^(p0))
            if (!(369 >= 826)) {
                v8 := []int{1}
                p1 = ""
                _ = v8
            }
            _ = v7
        }
        if (940 <= 65535) {
            switch ("hello" + p1) {
            case "\n\t":
                v9 := GenS1{F0: 1}
                _ = v9
            case "foo bar":
                v10 := []int{}
                _ = v10
            case "hello":
                var v11 uint64 = (p0 + (p0 | 18446744073709551615))
                _ = v11
            }
            p1 = p1
        } else {
            v12 := []int{127}
            _ = v12
        }
    } else {
        if true {
            for fi13 := 0; fi13 < 18; fi13++ {
                p0 -= p0
                _ = fi13
            }
            if true {
                fmt.Println(p2)
                var v14 float32 = float32((((-2.25 + float64(1.5)) + (float64(0.0) - 0.1)) + 3.14159))
                _ = v14
            }
            v15 := GenS2{F0: 0, F1: 255, F2: 42}
            _ = v15
        }
    }
    switch -314 {
    case -585:
        var v16 int32 = -45
        _ = v16
    case 360:
        for fi17 := 0; fi17 < 3; fi17++ {
            switch (^(fi17)) {
            case 765:
                switch p1 {
                case "\n\t":
                    var v18 uint8 = p2
                    _ = v18
                case "foo bar":
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v20 := []int{fi19, int(p2), 162, (^(fi19))}
                        v21 := GenS3{F0: 1, F1: uint32(p0)}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            case -896:
                v22 := []int{(fi17 % ((4294967295) | 1))}
                _ = v22
            case -979:
                if false {
                    var v23 uint64 = (p0 % (((p0 % ((p0) | 1))) | 1))
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS2{F0: 164, F1: 845, F2: 49}
                        _ = fi24
                        _ = v25
                    }
                }
            default:
                p2 -= p2
            }
            fmt.Println(p1, fi17, fi17)
            _ = fi17
        }
    default:
        if (1 != 4294967295) {
            p2 += 128
        }
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return (p0 ^ p0)
}
func genFn1(p0 uint32) int64 {
    p0 = (^(p0))
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int8 = -111
        v28 := GenS1{F0: 86}
        _ = fi26
        _ = v27
        _ = v28
    }
    p0 -= (p0 ^ p0)
    _ = p0
    return -121
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = 1.5
    var v3 uint8 = 6
    var v4 uint = 255
    v5 := []uint32{42}
    v6 := []uint32{42, 2147483648}
    var v29 float32 = 0.1
    switch "a" {
    case "foo bar":
        var v30 string = "a"
        _ = v30
    default:
        v31 := GenS4{F0: 297, F1: -299}
        _ = v31
    }
    fmt.Println(v4, v4, v1)
    fmt.Println(v1, v2, v2)
    var v32 float32 = ((v2 + float32(1e10)) + (float32(1e10) - float32(((float64((v29 / ((v1 * v1) + 1.5))) - (((1.5 + 3.14159) - 1e10) / 1.0)) - float64(((0.0 * float32(3.14159)) - 0.0))))))
    if len(v5) > 0 {
        v5[0] = 65535
    }
    var v33 int8 = 1
    var v34 bool = false
    v35 := GenS4{F0: 53, F1: -600}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
