package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int, p2 uint64) int16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int16 = 1
        _ = fi7
        _ = v8
    }
    if ((true || (3.14159 >= ((float64(float32((0.0 * -2.25))) / 1.5) / ((0.0 + 1.5) + (float64(-2.25) * 1e10))))) && true) {
        fmt.Println(p2, p2)
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (fi9 << ((2) & 63)) {
            case 42:
                var v10 int8 = (p0 - p0)
                _ = v10
            case -1:
                p0 = (p0 % ((int8(p2)) | 1))
            default:
                var v11 bool = (!false)
                _ = v11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                if false {
                    v13 := GenS1{F0: "\n\t", F1: -940}
                    if false {
                        if false {
                            fmt.Println(fi12, p0, p0)
                            var v14 int64 = -663
                            _ = v14
                        }
                        var v15 int16 = 0
                        switch "foo bar" {
                        case "a":
                            if false {
                                fmt.Println(v15, p2)
                            } else {
                                v16 := []int{}
                                _ = v16
                            }
                        case "":
                            if (p2 == 2) {
                                p2 = p2
                                p1 += (^(p1))
                            }
                        }
                        v15 -= (-(v15))
                        fmt.Println(p2)
                        v17 := GenS2{F0: uint(p2), F1: 42, F2: (-(v15))}
                        v18 := []int{fi12, (fi9 + fi9), p1, (^(p1))}
                        _ = v15
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint32 = 2147483648
                        v20 := GenS2{F0: 65535, F1: 1, F2: 876}
                        _ = v19
                        _ = v20
                    }
                    v21 := []int{42, p1}
                    fmt.Println(p1, p1, fi9)
                    _ = v13
                    _ = v21
                } else {
                    var v22 uint64 = p2
                    if (127 >= 255) {
                        p0 += (p0 - (p0 * (-(p0))))
                        switch "foo bar" {
                        case "a":
                            v23 := GenS1{F0: "a", F1: -2147483648}
                            _ = v23
                        default:
                            var v24 int = (p1 << ((2) & 63))
                            _ = v24
                        }
                    }
                    _ = v22
                }
                _ = fi12
            }
            _ = fi9
        }
    } else {
        var v25 int16 = -670
        p2 -= (^(p2))
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 int16 = 42
    var v2 bool = false
    var v3 int32 = -2147483648
    var v4 string = "foo bar"
    var v5 int64 = 488
    var v6 int16 = 855
    v26 := GenS2{F0: 698, F1: (v3 / ((2147483647) | 1)), F2: -637}
    var v27 float32 = (0.1 / float32((1.5 * float64(1.5))))
    v28 := GenS2{F0: 2, F1: (v3 >> ((653) & 63)), F2: 0}
    var v29 uint32 = 2147483647
    var v30 uint16 = 2
    v27 = (v27 - (v27 + v27))
    v31 := GenS3{F0: v5}
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 float64 = float64(0.0)
        _ = fi32
        _ = v33
    }
    var v34 string = v4
    fmt.Println(v27, v1)
    var v35 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
