package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint16 }
func genFn0() bool {
    var v7 float32 = (0.1 - float32(((((3.14159 - 1e10) + -2.25) * 1.5) - ((1e10 / 0.1) - 3.14159))))
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            if (!(42 <= 714)) {
                if false {
                    v11 := GenS1{F0: 65535, F1: -9223372036854775808}
                    var v12 uint8 = v8
                    var v13 int64 = 795
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: fi9, F1: uint16(v8)}
                v15 := []int{}
                _ = v14
                _ = v15
            }
            if true {
                var v16 uint64 = 879
                switch fi10 {
                case 425:
                    v17 := []int{fi9, -373, (fi10 << ((97) & 63))}
                    _ = v17
                case 765:
                    v18 := GenS1{F0: 42, F1: 1}
                    _ = v18
                case -783:
                    v7 += (-2.25 - v7)
                }
                var v19 uint32 = 891
                v19 = 2147483647
                _ = v16
                _ = v19
            } else {
                v20 := []int{fi9, fi9}
                _ = len(v20)
                _ = v20
            }
            _ = fi10
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v7 -= 0.1
            _ = fi21
        }
        switch fi9 {
        case -172:
            v22 := GenS3{F0: 408, F1: 1}
            _ = v22
        case -188:
            v7 = v7
        case -2147483648:
            v23 := GenS2{F0: (fi9 ^ 935), F1: 1}
            _ = v23
        }
        _ = fi9
    }
    var v24 int = 128
    fmt.Println(v24, v8)
    _ = v24
    _ = v7
    _ = v8
    return (!(false || (!true)))
}
func genFn1(p0 int64, p1 string) uint16 {
    var v25 uint = 255
    p1 = "hello"
    p0 = (p0 & -1)
    _ = p0
    _ = p1
    _ = v25
    return 1
}

func main() {
    var v1 float64 = -2.25
    var v2 string = "foo bar"
    var v3 int32 = 2147483647
    var v4 int64 = -76
    var v5 uint16 = 65535
    v6 := []int{-320}
    v4 -= v4
    fmt.Println(v3, v4)
    var v26 uint64 = 2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v5, v3, v4)
    v27 := GenS3{F0: 65535, F1: 255}
    var v28 float64 = ((v1 / (v1 * v1)) + 0.0)
    v2 = (v2 + v2)
    v1 -= (v28 + 1e10)
    switch "a" {
    case "":
        fmt.Println(v26, v1)
    case "a":
        v29 := GenS4{F0: ("foo bar" + (v2 + v2)), F1: (^(v26))}
        _ = v29
    case "\n\t":
        var v30 int16 = 1000
        _ = v30
    }
    v31 := GenS5{F0: -1, F1: (v5 << ((609) & 63))}
    fmt.Println(v28, v3, v5)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
