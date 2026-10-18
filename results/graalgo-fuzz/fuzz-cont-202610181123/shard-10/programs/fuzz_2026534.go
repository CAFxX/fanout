package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
type GenS4 struct { F0 int16; F1 uint; F2 uint32 }
func genFn0() int64 {
    var v4 string = ((("" + (("\n\t" + "") + "")) + "\n\t") + "\n\t")
    if (0.0 >= 0.1) {
        for fi5 := 0; fi5 < 9; fi5++ {
            v4 = v4
            if (0.1 > 1.5) {
                var v6 int8 = -128
                var v7 int16 = 340
                for fi8 := 0; fi8 < 5; fi8++ {
                    v6 -= -120
                    v9 := []int{fi8, (fi8 << ((42) & 63))}
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            }
            _ = fi5
        }
        v10 := GenS1{F0: -159, F1: 211}
        for fi11 := 0; fi11 < 11; fi11++ {
            switch fi11 {
            case -862:
                if (false || false) {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v4 = ("\n\t" + ((v4 + v4) + (("" + ((v4 + ((v4 + v4) + v4)) + "hello")) + ((v4 + (v4 + v4)) + v4))))
                        var v13 int16 = 1
                        v14 := []int{}
                        fmt.Println(v4, fi11)
                        fmt.Println(v4, v4, v13)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS2{F0: -251, F1: int16(fi15), F2: -1}
                        v4 = v4
                        _ = fi15
                        _ = v16
                    }
                    v17 := []int{(-(fi11))}
                    _ = v17
                }
            default:
                var v18 float32 = (0.0 * 0.1)
                _ = v18
            }
            v19 := GenS3{F0: 588, F1: 255, F2: 0}
            _ = fi11
            _ = v19
        }
        _ = v10
    }
    v20 := []int{1, 32768}
    v21 := GenS1{F0: -1, F1: 2}
    _ = v20
    _ = v21
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 uint = 1
    var v2 int8 = -45
    var v3 uint32 = 42
    v22 := GenS3{F0: v1, F1: 1, F2: int16(v2)}
    switch 1 {
    case -9223372036854775808:
        var v23 int64 = 721
        _ = v23
    case 98:
        v3 = v3
    case 42:
        v24 := GenS4{F0: int16(v2), F1: (v1 / ((v1) | 1)), F2: (v3 << (((v1 * v1)) & 63))}
        _ = v24
    default:
        fmt.Println(v1)
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := []int{}
        v3 -= 1
        var v27 float32 = float32(1e10)
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{127, -9223372036854775808}
    if false {
        v3 = v3
    }
    var v29 uint16 = 928
    if true {
        v28 = append(v28, -1)
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    var v30 int16 = 269
    fmt.Println(v30, v29)
    v1 = v1
    v29 -= (v29 * 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
