package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int16, p1 int, p2 uint16) uint16 {
    var v4 int64 = 42
    v5 := GenS1{F0: 65535, F1: 2}
    var v6 int8 = 42
    v7 := GenS1{F0: p2, F1: 42}
    var v8 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (p2 >> ((400) & 63))
}

func main() {
    var v1 int = 127
    var v2 string = ("hello" + "foo bar")
    var v3 uint8 = 0
    var v9 uint8 = v3
    if (2147483647 < 1) {
        if false {
            switch "\n\t" {
            case "":
                fmt.Println(v3, v1)
            case "hello":
                var v10 uint = 0
                _ = v10
            default:
                var v11 int64 = -171
                _ = v11
            }
            var v12 uint16 = 0
            _ = v12
        }
    } else {
        var v13 int8 = 65
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := []int{v1, 4294967295, fi14, fi14}
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS2{F0: 32767}
                fmt.Println(v2, v9, v3)
                v3 += (v9 << ((0) & 63))
                var v18 float64 = 3.14159
                _ = len(v15)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if true {
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 float64 = 3.14159
                var v22 float64 = 1e10
                switch 1 {
                case 821:
                    v23 := []int{fi20, 32767, v1, (fi20 + (fi20 ^ (fi20 | fi19)))}
                    _ = v23
                case -850:
                    v24 := GenS3{F0: 32767, F1: 283, F2: 2}
                    _ = v24
                case 128:
                    v25 := []int{}
                    _ = v25
                default:
                    v26 := GenS4{F0: -1}
                    _ = v26
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            var v27 uint16 = 42
            switch fi19 {
            case 505:
                for fi28 := 0; fi28 < 3; fi28++ {
                    v3 = 0
                    if true {
                        fmt.Println(fi19, fi19, v9)
                        v9 = (v3 & v3)
                        v29 := GenS1{F0: (v27 * v27), F1: 1}
                        switch v2 {
                        case "a":
                            v9 = (v3 % ((v3) | 1))
                        default:
                            var v30 string = v2
                            _ = v30
                        }
                        _ = v29
                    }
                    v31 := GenS3{F0: 32767, F1: 42, F2: 576}
                    _ = fi28
                    _ = v31
                }
            case -27:
                var v32 string = ("a" + v2)
                _ = v32
            default:
                v27 += (v27 + (v27 << ((4294967295) & 63)))
            }
            _ = fi19
            _ = v27
        }
        v3 -= 67
        v1 += 647
    } else {
        var v33 uint = 0
        _ = v33
    }
    v34 := []int{9223372036854775807, (v1 | 1), v1}
    v34 = append(v34, 42)
    if (724 < -1) {
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        switch v2 {
        case "foo bar":
            var v35 float32 = -2.25
            _ = v35
        case "":
            var v36 uint64 = 790
            _ = v36
        case "hello":
            v37 := GenS5{F0: 695, F1: (v3 | v9)}
            _ = v37
        }
    }
    v38 := GenS2{F0: 42}
    if len(v34) > 0 {
        v34[0] = v1
    }
    var v39 uint8 = v9
    var v40 uint64 = 670
    if true {
        var v41 int16 = int16(v40)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v40
    _ = v9
}
