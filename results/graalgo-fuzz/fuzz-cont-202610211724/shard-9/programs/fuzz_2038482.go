package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 uint16, p1 string) int32 {
    if true {
        v7 := []int{570}
        var v8 int = -922
        p1 = p1
        _ = v7
        _ = v8
    } else {
        p0 = p0
    }
    switch -796 {
    case -2147483648:
        v9 := []int{42, 851}
        _ = v9
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: true}
            switch "\n\t" {
            case "a":
                v12 := []int{fi10, (fi10 * fi10), fi10, (-(fi10))}
                _ = v12
            default:
                fmt.Println(p1, p1)
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS2{F0: 103, F1: (0.1 > ((((1e10 * -2.25) - -2.25) - (((1.5 / -2.25) * 1.5) - 0.1)) + (3.14159 + 1e10))), F2: p0}
                var v15 float64 = ((0.0 / -2.25) - (0.1 / -2.25))
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint64, p1 bool, p2 uint16) string {
    var v16 int = 1
    var v17 float64 = -2.25
    if (!(float64(1.5) != ((v17 + (0.0 * v17)) / v17))) {
        var v18 uint64 = p0
        v19 := GenS3{F0: (p0 / ((v18) | 1)), F1: 42}
        _ = v18
        _ = v19
    }
    v20 := []int{263, v16, 0, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v20
    return "a"
}

func main() {
    var v1 string = ("a" + "")
    var v2 bool = false
    var v3 uint32 = 2
    v4 := []uint32{437}
    v5 := []uint32{4294967295, 0}
    var v6 int8 = 127
    v6 -= 0
    var v21 string = v1
    if (322 <= 0) {
        v6 = (v6 % (((v6 ^ (v6 + 100))) | 1))
    } else {
        for fi22 := 0; fi22 < 2; fi22++ {
            v2 = (!(2 == 42))
            v1 = "a"
            _ = fi22
        }
        if ((((1e10 * -2.25) / ((0.1 / ((float32(1e10) / float32((1.5 + 0.0))) - 0.0)) / (1e10 * 3.14159))) - -2.25) < 1.5) {
            var v23 int32 = 0
            _ = v23
        } else {
            if (v2 || true) {
                v21 = ""
                fmt.Println(v1, v3, v3)
                var v24 uint = 255
                _ = v24
            }
            var v25 int16 = -835
            _ = v25
        }
        v1 = v1
    }
    if v2 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v26 bool = false
        _ = v26
    }
    var v27 bool = (!(146 > 42))
    v6 = (v6 % ((v6) | 1))
    fmt.Println(v21, v2)
    v2 = (3.14159 < (float64(3.14159) * (((1e10 - 0.1) / (1e10 + 0.1)) - (3.14159 + -2.25))))
    v2 = (!true)
    v28 := GenS4{F0: 42, F1: ((v21 + v21) + (v1 + v21))}
    v1 = (v1 + ((v21 + v1) + v1))
    v6 -= v6
    v29 := GenS5{F0: -93, F1: -357, F2: ((v1 + "\n\t") + "a")}
    v6 -= v6
    _ = len(v4)
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 string = "\n\t"
        _ = fi30
        _ = v31
    }
    if false {
        if v27 {
            v3 += 386
        }
    } else {
        var v32 float64 = 1e10
        if (!false) {
            v21 = ""
        }
        if len(v5) > 0 {
            v5[0] = 323
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
