package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 int64, p2 int8) string {
    var v4 float64 = float64(float32(float64(0.1)))
    v5 := []int{32768, 989}
    var v6 int8 = p2
    v7 := GenS1{F0: 114}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 int16) uint16 {
    v8 := []int{4294967295, -631, 9223372036854775807, 4294967295}
    var v9 int64 = 1
    v10 := GenS2{F0: 255, F1: 0}
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS2{F0: 510, F1: 2}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 2
    v3 := make(map[int]uint32)
    switch -75 {
    case 2147483647:
        v13 := []int{}
        _ = v13
    }
    if false {
        delete(v3, 999)
        v14 := GenS1{F0: 1}
        _ = v14
    } else {
        var v15 uint16 = (v2 << ((18446744073709551615) & 63))
        _ = v15
    }
    v1 = (v1 * v1)
    var v16 float64 = (3.14159 - (1e10 / ((0.1 / 1.0) - 1.5)))
    v16 -= v16
    v17 := v3[707]
    v1 -= (0.0 / v1)
    var v18 int = 127
    switch "\n\t" {
    case "":
        var v19 uint32 = v17
        _ = v19
    case "a":
        v1 = (v1 - 1e10)
    default:
        for fi20 := 0; fi20 < 7; fi20++ {
            v17 = (v17 % (((^(v17))) | 1))
            v21 := GenS1{F0: 4294967295}
            v22 := GenS3{F0: v18}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 int8 = 100
    var v24 bool = (0 > -584)
    for fi25 := 0; fi25 < 3; fi25++ {
        if false {
            v18 = v18
        }
        _ = fi25
    }
    v17 = 2147483648
    v18 -= 32768
    for fi26 := 0; fi26 < 12; fi26++ {
        if (1 == 510) {
            for fi27 := 0; fi27 < 5; fi27++ {
                switch "foo bar" {
                case "hello":
                    v28 := []int{(fi26 + (-(fi27))), fi27}
                    _ = v28
                case "foo bar":
                    v2 -= (v2 ^ (v2 & v2))
                case "\n\t":
                    if (!(126 != 68)) {
                        _ = len(v3)
                        v29 := []int{(v18 + fi27), (fi27 & fi26)}
                        v30 := v3[fi27]
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 4294967295}
                        var v32 int8 = (v23 & (v23 - v23))
                        if v24 {
                            v2 = (v2 << ((862) & 63))
                            var v33 int = (fi26 + (^(fi27)))
                            v1 += v1
                            v34 := GenS4{F0: -9223372036854775808}
                            var v35 int8 = (v23 ^ 0)
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = v31
                        _ = v32
                    }
                }
                v16 += (((float64(float32((float64((float32(v16) * float32(1e10))) - 1.5))) / ((v16 + (v16 + ((v16 + 3.14159) - v16))) + (float64((((v1 - v1) - v1) + (0.0 + 3.14159))) * v16))) + 1.5) * v16)
                if v24 {
                    var v36 float32 = (((v1 + float32(3.14159)) * v1) / ((0.1 - ((v1 * 1e10) / 1.5)) - (v1 + 1.5)))
                    _ = v36
                }
                _ = fi27
            }
            if (!(!false)) {
                v37 := GenS2{F0: (^(v2)), F1: 790}
                _ = v37
            } else {
                v38 := GenS1{F0: 4294967295}
                delete(v3, (-(fi26)))
                _ = v38
            }
        } else {
            var v39 uint16 = (v2 ^ v2)
            var v40 int32 = 42
            v23 = (v23 % ((-128) | 1))
            _ = v39
            _ = v40
        }
        _ = fi26
    }
    v2 -= (v2 - (^(v2)))
    v41 := GenS5{F0: 592}
    switch (-(v18)) {
    case -9223372036854775808:
        v1 = v1
    case 32768:
        v24 = v24
    case 42:
        var v42 string = ""
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v41
}
