package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 bool, p1 uint16, p2 uint) int32 {
    var v6 float64 = -2.25
    v6 = v6
    if (p0 || (((^(p2)) >= (p2 % ((p2) | 1))) && (p0 && true))) {
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                switch fi7 {
                case 127:
                    v8 := GenS1{F0: 170}
                    _ = v8
                case 42:
                    v9 := GenS1{F0: 1}
                    _ = v9
                }
                v10 := []int{fi7}
                if len(v10) > 0 {
                    v10[0] = (fi7 % (((fi7 | 676)) | 1))
                }
                _ = fi7
                _ = v10
            }
            for fi11 := 0; fi11 < 3; fi11++ {
                p0 = (1.5 >= 3.14159)
                p2 += (p2 % (((^(p2))) | 1))
                _ = fi11
            }
        } else {
            var v12 uint = (p2 + 182)
            v13 := GenS2{F0: 32767}
            _ = v12
            _ = v13
        }
    }
    if true {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS3{F0: "foo bar", F1: (p2 | 65535)}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS4{F0: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    return 593
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    var v4 int16 = 42
    var v5 uint64 = uint64(v1)
    var v17 uint8 = 255
    v18 := []int{42, 9223372036854775807}
    var v19 int = 32767
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint16 = 0
        var v22 int8 = int8(v2)
        _ = fi20
        _ = v21
        _ = v22
    }
    if false {
        v4 = (-(v4))
        v23 := GenS2{F0: v19}
        if len(v18) > 0 {
            v18[0] = v19
        }
        _ = v23
    } else {
        fmt.Println(v1, v19)
    }
    v1 = (v17 + (v17 + v1))
    v24 := GenS5{F0: 65535, F1: (986 != 2)}
    var v25 uint8 = (v17 ^ (v1 >> ((255) & 63)))
    if false {
        fmt.Println(v4, v25)
        v26 := GenS6{F0: 1, F1: (-(v4)), F2: 1}
        _ = v26
    } else {
        v27 := GenS4{F0: v19}
        _ = v27
    }
    var v28 uint = 4294967295
    if true {
        _ = len(v18)
    } else {
        if false {
            var v29 int64 = int64(v19)
            v30 := v3[1]
            switch "a" {
            case "\n\t":
                _ = len(v3)
            case "foo bar":
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 int8 = -24
                    _ = fi31
                    _ = v32
                }
            default:
                v19 += (v19 % ((v19) | 1))
            }
            _ = v29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
