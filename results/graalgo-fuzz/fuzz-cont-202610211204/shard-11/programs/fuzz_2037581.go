package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64) bool {
    v6 := GenS1{F0: -2147483648, F1: 1, F2: ("" + "a")}
    v7 := []int{0, 9223372036854775807, 32767, int(p0)}
    var v8 float64 = 0.0
    var v9 string = ("\n\t" + "\n\t")
    for fi10 := 0; fi10 < 17; fi10++ {
        switch ("" + v9) {
        case "":
            for fi11 := 0; fi11 < 5; fi11++ {
                v9 = ((v9 + v9) + v9)
                _ = fi11
            }
        case "hello":
            if ((1.5 - (0.1 + 1.5)) == -2.25) {
                var v12 float64 = ((v8 * v8) * v8)
                _ = v12
            }
        }
        v13 := GenS2{F0: 0, F1: 635}
        var v14 uint64 = (p0 + p0)
        _ = fi10
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (109 >= -75)
}
func genFn1(p0 int32) int32 {
    if (270 > 32768) {
        p0 += (-(p0))
    }
    var v15 float32 = -2.25
    v16 := []int{9223372036854775807}
    _ = p0
    _ = v15
    _ = v16
    return (p0 * p0)
}

func main() {
    var v1 int16 = 0
    var v2 int64 = -1
    v3 := []uint32{1, 1, 0}
    v4 := make(map[int]bool)
    v5 := []uint32{4294967295, 0}
    switch 9223372036854775807 {
    case 4294967295:
        v2 += v2
    }
    if true {
        if false {
            v1 = -362
            for fi17 := 0; fi17 < 12; fi17++ {
                switch (fi17 % (((fi17 + 32768)) | 1)) {
                case 9223372036854775807:
                    v1 += (v1 + v1)
                case 32768:
                    fmt.Println(v1)
                case 32767:
                    if (37 != 42) {
                        v1 = 222
                    }
                }
                var v18 int64 = -1
                _ = fi17
                _ = v18
            }
        } else {
            v19 := GenS3{F0: 2, F1: (v2 ^ v2), F2: true}
            switch 863 {
            case 524:
                for fi20 := 0; fi20 < 18; fi20++ {
                    v1 += v1
                    _ = fi20
                }
            case 0:
                v1 -= (v1 << ((609) & 63))
            case -2147483648:
                v2 = v2
            default:
                if false {
                    v21 := GenS4{F0: 945}
                    v1 -= -1
                    v2 = (v2 + (v2 & v2))
                    _ = v21
                } else {
                    if false {
                        v5 = append(v5, 65535)
                        v22 := GenS1{F0: 764, F1: genFn1(genFn1(0)), F2: "hello"}
                        var v23 string = "hello"
                        _ = v22
                        _ = v23
                    }
                }
            }
            _ = v19
        }
    }
    v24 := GenS3{F0: 42, F1: (^(v2)), F2: true}
    var v25 bool = false
    var v26 bool = (v25 && (!(true && true)))
    v4[32767] = true
    fmt.Println(v26)
    switch 2147483647 {
    case 0:
        var v27 bool = true
        _ = v27
    case -250:
        v25 = (!v26)
    default:
        v28 := GenS5{F0: 2, F1: 42}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
