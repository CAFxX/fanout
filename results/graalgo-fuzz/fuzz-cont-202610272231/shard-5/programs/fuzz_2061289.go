package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int8) uint64 {
    p0 -= (p0 - p0)
    v7 := GenS1{F0: 255}
    var v8 float32 = (3.14159 - float32(0.1))
    if ((!(true && true)) && false) {
        v9 := GenS1{F0: 127}
        var v10 float64 = 1e10
        _ = v10
        _ = v9
    } else {
        v11 := []int{4294967295}
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(p0)
        v13 := []int{-436, fi12, (fi12 << ((42) & 63))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint, p1 bool, p2 int16) int8 {
    var v14 int = 127
    var v15 int8 = -1
    var v16 bool = (true || (!(2147483648 > 0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return v15
}

func main() {
    var v1 uint = 241
    v2 := make(map[int]bool)
    var v3 string = (("\n\t" + "") + "a")
    var v4 uint = (v1 + (^(v1)))
    var v5 uint = (v1 >> (((v4 * v4)) & 63))
    v6 := make(map[int]bool)
    for fi17 := 0; fi17 < 14; fi17++ {
        if ((1 > 42) && false) {
            v5 = 1
            var v18 int16 = -1
            _ = v18
        } else {
            switch fi17 {
            case -950:
                v19 := GenS1{F0: 81}
                _ = v19
            }
        }
        _ = fi17
    }
    delete(v2, int(v5))
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v4)
        _ = fi20
    }
    var v21 float64 = (((1.5 + 3.14159) / 3.14159) + (0.0 * (0.1 / 1e10)))
    switch -759 {
    case 128:
        var v22 float64 = (0.0 + v21)
        _ = v22
    case 32767:
        var v23 uint32 = 2
        _ = v23
    default:
        if true {
            v24 := v6[127]
            var v25 int = 274
            _ = v24
            _ = v25
        }
    }
    if false {
        v5 += v1
        for fi26 := 0; fi26 < 18; fi26++ {
            if ((v5 | (v4 >> (((v5 * (v4 + 1))) & 63))) == v1) {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 float64 = ((0.0 - float64((0.1 * 0.1))) / v21)
                    _ = fi27
                    _ = v28
                }
                switch v3 {
                case "a":
                    var v29 uint32 = 2
                    _ = v29
                default:
                    v30 := []int{0, (fi26 ^ fi26), (fi26 % (((-(fi26))) | 1))}
                    _ = v30
                }
                var v31 uint = (v1 - v1)
                _ = v31
            }
            _ = fi26
        }
        v6[2147483647] = true
    }
    var v32 uint32 = 902
    v33 := GenS2{F0: uint16(v4)}
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 2; fi35++ {
            v21 = (v21 + (v21 - v21))
            v36 := GenS2{F0: 1}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 uint = (v5 * (v4 | v5))
    v3 = v3
    v4 += 65535
    if (713 == -9223372036854775808) {
        v38 := []int{-771}
        v32 -= v32
        v5 = v4
        _ = v38
    } else {
        v39 := GenS3{F0: v32}
        fmt.Println(v1, v32)
        var v40 float32 = (1.5 * ((3.14159 / 1e10) - ((-2.25 * (0.1 + ((3.14159 - 1e10) * 1.5))) + (0.1 - 0.0))))
        _ = v39
        _ = v40
    }
    v41 := GenS4{F0: 42}
    v42 := []int{32768, 1}
    v21 += ((v21 / 1.0) - 0.0)
    v21 = v21
    v43 := GenS5{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
