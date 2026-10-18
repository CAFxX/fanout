package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int32; F2 int32 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0() int32 {
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 -= (v4 * v4)
        _ = fi5
    }
    if (102 != 2) {
        var v6 uint16 = v4
        if (!true) {
            switch 0 {
            case -439:
                switch -195 {
                case 1:
                    var v7 float64 = 1e10
                    _ = v7
                case 2147483647:
                    var v8 uint8 = 127
                    _ = v8
                }
            case -1:
                v4 += 395
            default:
                v9 := []int{128}
                _ = v9
            }
            v10 := []int{128, -349}
            var v11 int16 = 820
            _ = v10
            _ = v11
        }
        v12 := []int{127}
        _ = v12
        _ = v6
    } else {
        v13 := []int{316}
        var v14 uint32 = 2147483647
        if (!(42 >= 2147483647)) {
            switch ("a" + "a") {
            case "\n\t":
                if true {
                    fmt.Println(v4)
                } else {
                    v4 = (v4 * (v4 - v4))
                }
            }
            var v15 int32 = -2147483648
            v16 := GenS1{F0: int32(v14), F1: 1000, F2: 255}
            _ = v15
            _ = v16
        } else {
            _ = len(v13)
            switch -752 {
            case 353:
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 string = "hello"
                    var v19 int = fi17
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 4294967295:
                v20 := GenS2{F0: 0}
                _ = v20
            case 2147483647:
                for fi21 := 0; fi21 < 6; fi21++ {
                    v13 = append(v13, fi21)
                    _ = fi21
                }
            default:
                var v22 float32 = 1.5
                _ = v22
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        _ = v13
        _ = v14
    }
    _ = v4
    return -817
}
func genFn1(p0 uint64, p1 int8, p2 uint8) uint {
    v23 := GenS1{F0: -554, F1: 1000, F2: (p2 | (p2 - 2))}
    var v24 int8 = 42
    p1 = p1
    v25 := GenS3{F0: 604, F1: 882, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    _ = v25
    return 0
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v26 := []int{-9223372036854775808, 4294967295, 128, -9223372036854775808}
    if false {
        v1 += v1
    }
    if true {
        v27 := GenS4{F0: genFn1(2, -1, 32), F1: -2147483648}
        _ = v27
    }
    var v28 int64 = 42
    v26 = append(v26, 128)
    var v29 uint64 = 2
    v1 = 1e10
    var v30 bool = (false || (!(0 <= 65535)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
