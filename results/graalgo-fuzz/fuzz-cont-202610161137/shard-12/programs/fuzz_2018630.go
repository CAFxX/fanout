package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 uint64, p2 int8) bool {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(fi7 - 32768), 9, int(p1)}
        p0 = p0
        var v9 uint32 = uint32(p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (^(p0))
    v10 := GenS1{F0: ((^(p2)) == -128), F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (!(1e10 <= 0.1))
}
func genFn1() uint {
    v11 := []int{}
    for fi12 := 0; fi12 < 10; fi12++ {
        if len(v11) > 0 {
            v11[0] = fi12
        }
        _ = fi12
    }
    _ = v11
    return 2
}

func main() {
    var v1 int64 = -1
    var v2 uint8 = 120
    var v3 bool = true
    var v4 uint32 = 710
    v5 := make(map[string]string)
    var v6 int32 = 711
    fmt.Println(v1, v6, v3)
    v13 := v5[""]
    var v14 uint16 = 132
    v15 := GenS2{F0: 4294967295, F1: 4294967295}
    var v16 uint = 646
    v1 += (v1 ^ (v1 & 403))
    if (!((0.0 - (3.14159 * float32(float64(float32(1e10))))) != (0.1 - float32(0.0)))) {
        _ = len(v5)
        v1 += v1
        v14 -= (^(v14))
    } else {
        v5[((v13 + v13) + v13)] = v13
        v2 += 255
        if false {
            switch "hello" {
            case "":
                fmt.Println(v13)
            case "foo bar":
                if ((v16 + v16) >= (v16 & uint(v4))) {
                    v17 := []int{1}
                    for fi18 := 0; fi18 < 12; fi18++ {
                        v13 = "hello"
                        switch -549 {
                        case -378:
                            v6 += v6
                        }
                        for fi19 := 0; fi19 < 20; fi19++ {
                            v1 -= v1
                            _ = len(v17)
                            v3 = v3
                            _ = fi19
                        }
                        _ = fi18
                    }
                    v17 = append(v17, 966)
                    _ = v17
                } else {
                    v5[v13] = v13
                    v20 := GenS2{F0: 1, F1: genFn1()}
                    _ = v20
                }
            }
            if false {
                var v21 uint = (v16 / (((^(v16))) | 1))
                _ = v21
            } else {
                v16 += genFn1()
            }
        }
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        switch v13 {
        case "\n\t":
            fmt.Println(v3)
        case "foo bar":
            var v23 int8 = -31
            _ = v23
        case "":
            fmt.Println(v14)
        }
        _ = fi22
    }
    var v24 int8 = 0
    v5["\n\t"] = v13
    v25 := GenS2{F0: 42, F1: v16}
    var v26 uint = (^(v16))
    v5[(v13 + "hello")] = v13
    var v27 float64 = 3.14159
    v28 := GenS3{F0: 32768, F1: v24}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
