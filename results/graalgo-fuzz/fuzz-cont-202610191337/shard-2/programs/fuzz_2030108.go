package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 int; F2 bool }
func genFn0(p0 int8, p1 uint8) uint {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, -743, fi7}
        _ = fi7
        _ = v8
    }
    var v9 int32 = 0
    v10 := GenS1{F0: (false || (4294967295 == 685)), F1: 4294967295, F2: p0}
    v9 = 0
    var v11 uint32 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 831
}
func genFn1(p0 uint8, p1 int) uint16 {
    v12 := GenS1{F0: false, F1: 0, F2: 41}
    p1 = p1
    var v13 float32 = 0.1
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 int = 32767
    var v5 uint16 = 32767
    v6 := make(map[string]bool)
    v14 := GenS1{F0: (1e10 > ((((-2.25 / -2.25) * -2.25) * 0.1) + 3.14159)), F1: 1, F2: -77}
    var v15 int16 = 32767
    fmt.Println(v4, v15, v15)
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint32 = 1
        v17 -= (v17 | 0)
        v18 := GenS2{F0: false, F1: fi16, F2: (!true)}
        _ = fi16
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "\n\t":
        _ = len(v3)
    }
    v5 -= 2
    if true {
        switch (v4 & (^(v4))) {
        case 0:
            for fi19 := 0; fi19 < 16; fi19++ {
                if true {
                    if (v5 <= (v5 >> ((genFn0(100, 1)) & 63))) {
                        v5 += genFn1(2, fi19)
                        if false {
                            for fi20 := 0; fi20 < 2; fi20++ {
                                fmt.Println(fi19)
                                v5 -= 32767
                                v15 = v15
                                _ = fi20
                            }
                            v21 := []int{fi19, (^(v4))}
                            _ = v21
                        }
                    }
                    v5 += 32767
                }
                _ = fi19
            }
        case 14:
            if false {
                v15 = (v15 / ((42) | 1))
                v5 -= (^(v5))
                var v22 float64 = (float64((0.1 - -2.25)) - 0.0)
                _ = v22
            } else {
                delete(v3, "foo bar")
                fmt.Println(v4)
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 int64 = int64(v15)
                    _ = fi23
                    _ = v24
                }
            }
        }
        v25 := v3["foo bar"]
        var v26 int32 = 89
        _ = v25
        _ = v26
    } else {
        v5 += v5
        v4 -= (v4 | (v4 & (-(v4))))
        v15 = (v15 | 1000)
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(v15)
        if (!(3.14159 <= float32((1e10 - 3.14159)))) {
            v4 = v4
        }
        var v28 int32 = 871
        _ = fi27
        _ = v28
    }
    var v29 int16 = v15
    _ = len(v1)
    v29 = v15
    v15 -= -32768
    var v30 int32 = -799
    v31 := GenS2{F0: (762 <= 2), F1: (v4 | (v4 - 42)), F2: (448 < v5)}
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 string = "foo bar"
        for fi34 := 0; fi34 < 5; fi34++ {
            v15 = v15
            var v35 int8 = 9
            fmt.Println(v33)
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v36 := []int{(v4 % ((32767) | 1)), (v4 + (v4 * -911))}
    v36 = append(v36, int(v30))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
