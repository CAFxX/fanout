package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint, p1 int64) bool {
    var v5 float64 = (0.0 / 1.0)
    var v6 uint = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int, p1 uint8, p2 uint) int64 {
    p0 += 127
    p0 += 9223372036854775807
    for fi7 := 0; fi7 < 20; fi7++ {
        p2 = 186
        _ = fi7
    }
    v8 := GenS1{F0: "a", F1: 1, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 uint8 = 255
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int64 = -408
        switch (("foo bar" + "a") + "a") {
        case "a":
            v11 := []int{fi9, -594}
            _ = v11
        case "foo bar":
            switch (^(fi9)) {
            case 128:
                var v12 float64 = ((0.1 * 1.5) + 0.1)
                _ = v12
            case 0:
                v4 += v4
            case -2147483648:
                if (1 == 432) {
                    v13 := []int{fi9, fi9}
                    _ = v13
                }
            }
        default:
            v1 = v1
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v4)
    v1 = false
    v14 := GenS1{F0: "foo bar", F1: 1, F2: 42}
    v15 := GenS2{F0: 42, F1: v4}
    if (1e10 == 1.5) {
        if v1 {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := []int{}
                _ = fi16
                _ = v17
            }
        }
    } else {
        var v18 uint = 2
        _ = v18
    }
    var v19 uint64 = 1
    if v1 {
        for fi20 := 0; fi20 < 7; fi20++ {
            if v1 {
                switch fi20 {
                case -663:
                    for fi21 := 0; fi21 < 17; fi21++ {
                        switch 127 {
                        case 42:
                            fmt.Println(fi20, fi20, v4)
                        default:
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    var v22 float64 = 1e10
                                    _ = v22
                                case "":
                                    v23 := GenS1{F0: "\n\t", F1: 42, F2: 2147483647}
                                    _ = v23
                                case "a":
                                    v24 := GenS2{F0: 275, F1: (v4 | 231)}
                                    _ = v24
                                default:
                                    var v25 float64 = 1e10
                                    _ = v25
                                }
                            }
                        }
                        if true {
                            var v26 int8 = 42
                            if false {
                                if (!(97 != 1)) {
                                    var v27 bool = (2147483648 <= 576)
                                    _ = v27
                                } else {
                                    fmt.Println(v26, fi21, v4)
                                    var v28 float64 = (float64(float32((1e10 / (1.5 - float64(0.0))))) - 3.14159)
                                    _ = v28
                                }
                            } else {
                                v19 += (v19 & 42)
                                v19 = 1
                                v29 := v3[(fi20 << ((255) & 63))]
                                _ = v29
                            }
                            _ = v26
                        } else {
                            if false {
                                var v30 float64 = 1.5
                                _ = v30
                            }
                        }
                        _ = fi21
                    }
                case 5:
                    fmt.Println(v19)
                default:
                    fmt.Println(v19, v19, v1)
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    v19 = (v19 / (((^(v19))) | 1))
                    _ = fi31
                }
                _ = len(v3)
            }
            _ = len(v3)
            _ = fi20
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 int64 = genFn1((^(fi33)), v4, 4294967295)
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
