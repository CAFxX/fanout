package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 uint) string {
    fmt.Println(p0)
    v4 := GenS1{F0: 1}
    p0 += 18446744073709551615
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint8, p1 int32) uint32 {
    var v6 int32 = (^(p1))
    v7 := []int{127, 42, 556}
    v8 := GenS2{F0: 530, F1: 2147483647, F2: true}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint32 = 160
        _ = fi9
        _ = v10
    }
    p0 += (p0 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 752
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 223
    var v3 bool = true
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 float64 = 1e10
        var v13 uint8 = 128
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 4294967295, F1: 38, F2: v3}
            v16 := []int{fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v17 := GenS1{F0: 42}
    var v18 string = ("a" + "\n\t")
    var v19 bool = ((-1 > 42) || (!(18446744073709551615 > 18446744073709551615)))
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := GenS3{F0: (true && (true && false)), F1: v1}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        if (2 >= 991) {
            if (!v3) {
                v23 := []int{-9223372036854775808, fi22, fi22}
                v23 = append(v23, fi22)
                for fi24 := 0; fi24 < 15; fi24++ {
                    fmt.Println(v18, v19)
                    fmt.Println(v19, v3)
                    v25 := GenS1{F0: 1}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                v1 += (v1 - v1)
            }
            switch ("foo bar" + ("" + v18)) {
            case "\n\t":
                v1 -= (v1 % ((100) | 1))
            case "foo bar":
                v18 = v18
            }
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 bool = v3
                var v28 uint8 = 128
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        _ = fi22
    }
    v29 := GenS2{F0: 4294967295, F1: 128, F2: (false && (-9223372036854775808 < 301))}
    v30 := []int{-393, 42, -1}
    v31 := GenS3{F0: (!v3), F1: (v1 >> ((1) & 63))}
    var v32 uint = 0
    switch v18 {
    case "a":
        v18 = (v18 + "")
    case "":
        var v33 bool = (-9223372036854775808 != int64(v2))
        _ = v33
    case "foo bar":
        var v34 int32 = -1
        _ = v34
    default:
        if true {
            v30 = append(v30, int(v2))
            if len(v30) > 0 {
                v30[0] = 42
            }
            var v35 bool = ((0.1 >= 1.5) && (v19 && ((true && ((v32 >> ((v32) & 63)) == (^(v32)))) || (((v32 ^ 0) >= v32) || ((32767 <= -783) && v19)))))
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 int16 = -173
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    if len(v30) > 0 {
        v30[0] = 128
    }
    v32 = (v32 % (((v32 / ((v32) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
