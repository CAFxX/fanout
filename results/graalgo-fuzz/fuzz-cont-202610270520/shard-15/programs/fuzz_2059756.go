package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
func genFn0(p0 bool, p1 uint16, p2 int8) int64 {
    var v4 uint16 = p1
    var v5 float64 = ((((float64(0.0) / 1.5) / 1.5) - -2.25) / 1.0)
    v6 := GenS1{F0: 42, F1: 127, F2: -9223372036854775808}
    v7 := []int{128, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 943
}

func main() {
    var v1 string = ("a" + ("\n\t" + "\n\t"))
    var v2 string = v1
    var v3 int8 = 54
    v1 = v1
    fmt.Println(v2)
    var v8 int = -1
    for fi9 := 0; fi9 < 8; fi9++ {
        if (127 <= 255) {
            if false {
                v10 := []int{}
                fmt.Println(v1, fi9)
                _ = v10
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 uint32 = 643
                    v13 := []int{(v8 + (^(fi11))), fi9, fi9}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
        } else {
            v14 := []int{-9223372036854775808, -9223372036854775808, v8, (fi9 - fi9)}
            v3 = (v3 << ((986) & 63))
            _ = v14
        }
        v15 := []int{fi9}
        fmt.Println(v8, v1)
        _ = fi9
        _ = v15
    }
    v16 := []int{(^(v8)), (v8 / ((v8) | 1))}
    var v17 int16 = 1
    v16 = append(v16, (v8 + (v8 / (((v8 % ((v8) | 1))) | 1))))
    var v18 float64 = (3.14159 / 1.5)
    switch v8 {
    case 42:
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 float32 = (((1.5 / 3.14159) - 0.0) / (((0.1 / 1.5) / ((((-2.25 * float32(float64(((((1e10 + ((-2.25 / 1e10) - 0.0)) / 1e10) * (-2.25 + -2.25)) / -2.25)))) + (1e10 * 3.14159)) + -2.25) / 1.5)) / 1.5))
            _ = fi19
            _ = v20
        }
    case 842:
        v2 = v1
    case -253:
        switch (^(v8)) {
        case -2147483648:
            var v21 string = (v2 + v2)
            _ = v21
        case 0:
            v22 := GenS1{F0: 97, F1: v8, F2: 1}
            _ = v22
        case 127:
            v17 += 0
        default:
            switch "\n\t" {
            case "hello":
                var v23 uint8 = 127
                _ = v23
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
