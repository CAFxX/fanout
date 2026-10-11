package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int32 }
func genFn0(p0 string, p1 string, p2 bool) uint16 {
    p1 = ((("foo bar" + p0) + "\n\t") + p1)
    fmt.Println(p1, p2, p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint64) int32 {
    var v4 int32 = -206
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 802, F1: 0, F2: (v4 % ((-274) | 1))}
            var v8 string = ""
            var v9 int = (fi5 - (fi5 / ((fi5) | 1)))
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if true {
            var v10 string = "\n\t"
            switch "foo bar" {
            case "\n\t":
                var v11 int32 = -801
                _ = v11
            default:
                v12 := []int{(fi5 >> ((0) & 63))}
                _ = v12
            }
            _ = v10
        } else {
            var v13 int = fi5
            for fi14 := 0; fi14 < 5; fi14++ {
                v4 -= v4
                var v15 float64 = 0.1
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            if (p0 <= p0) {
                v17 := []int{}
                p0 -= (p0 / ((p0) | 1))
                _ = v17
            }
            v18 := GenS1{F0: 269, F1: 127, F2: (v4 % ((-713) | 1))}
            v4 = (v4 % (((v4 | v4)) | 1))
            _ = fi16
            _ = v18
        }
        _ = fi5
    }
    var v19 bool = ((v4 ^ v4) > (v4 - (v4 % ((v4) | 1))))
    _ = p0
    _ = v19
    _ = v4
    return (^(v4))
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 511
    var v3 uint = 255
    fmt.Println(v3, v2, v3)
    fmt.Println(v2, v2, v3)
    var v20 uint16 = 0
    var v21 uint32 = 65535
    v22 := GenS1{F0: v21, F1: 100, F2: 1}
    var v23 int8 = 100
    var v24 float64 = (0.0 + ((1.5 + 0.1) - (1.5 / 3.14159)))
    v23 -= v23
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
