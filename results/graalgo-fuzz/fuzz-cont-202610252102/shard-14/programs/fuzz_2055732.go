package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 int64) uint64 {
    p0 = (p0 / ((p0) | 1))
    if (57 < 42) {
        p0 -= p0
    } else {
        var v5 int32 = -2147483648
        if true {
            v6 := []int{-731, 0}
            _ = v6
        } else {
            var v7 int64 = (p0 + (p0 * (p0 | 1)))
            v8 := []int{523, int(v5), -9223372036854775808}
            v7 -= v7
            _ = v7
            _ = v8
        }
        _ = v5
    }
    var v9 uint32 = 820
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 bool = (42 != 18446744073709551615)
    var v2 int64 = 0
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    if v1 {
        v1 = v1
    } else {
        var v10 uint64 = 1
        _ = v10
    }
    fmt.Println(v1, v1)
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 0}
        var v13 uint8 = 1
        v2 = 605
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            if v1 {
                fmt.Println(fi16)
                v14 = (v14 & 0)
            }
            var v17 string = "a"
            if v1 {
                var v18 uint = 611
                _ = v18
            } else {
                _ = len(v4)
                v19 := GenS2{F0: 1}
                fmt.Println(v1)
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        switch "hello" {
        case "foo bar":
            if v1 {
                if v1 {
                    fmt.Println(v1, fi15, v2)
                    _ = len(v4)
                    v1 = false
                } else {
                    fmt.Println(v2, v14)
                }
            } else {
                var v20 float64 = ((1e10 - (0.1 / (((float64(0.1) / 1.5) * 1e10) * float64(0.1)))) / 1.5)
                _ = v20
            }
        case "a":
            v2 += (v2 | v2)
        default:
            switch 4294967295 {
            case -547:
                var v21 float32 = 0.1
                _ = v21
            case -566:
                for fi22 := 0; fi22 < 15; fi22++ {
                    v1 = v1
                    if (!(814 == 255)) {
                        var v23 uint = uint(v14)
                        switch "" {
                        case "":
                            v2 -= -353
                        case "a":
                            v24 := GenS3{F0: 310, F1: 1000, F2: false}
                            _ = v24
                        case "\n\t":
                            v25 := []int{-149}
                            _ = v25
                        default:
                            v26 := []int{fi22, fi15}
                            _ = v26
                        }
                        _ = v23
                    } else {
                        var v27 int = 1
                        v28 := []int{fi15, fi15, (-(fi15))}
                        _ = v27
                        _ = v28
                    }
                    _ = fi22
                }
            }
        }
        _ = fi15
    }
    var v29 float32 = (1.5 / (1.5 - -2.25))
    if v1 {
        v30 := GenS3{F0: 2, F1: 1000, F2: (((!(!("hello" <= ("hello" + "\n\t")))) || false) && false)}
        _ = v30
    }
    v29 = (1.5 / (v29 + (v29 + (v29 * 0.0))))
    var v31 uint64 = 4294967295
    if v1 {
        if v1 {
            var v32 int16 = -89
            fmt.Println(v14, v32)
            _ = v32
        }
        v14 -= 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
