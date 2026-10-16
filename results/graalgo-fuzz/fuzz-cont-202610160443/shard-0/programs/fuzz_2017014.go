package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 uint }

func main() {
    v1 := []string{("hello" + "foo bar")}
    var v2 int16 = 725
    var v3 int16 = (v2 ^ (v2 / ((-41) | 1)))
    v4 := make(map[uint32]int)
    fmt.Println(v3, v3)
    v2 = (-(v3))
    v2 = 0
    fmt.Println(v2, v2)
    var v5 int16 = (-(v2))
    fmt.Println(v3, v5, v5)
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(v2, fi6, v3)
            _ = fi7
        }
        var v8 int16 = (v2 & v2)
        v5 = v5
        _ = fi6
        _ = v8
    }
    var v9 float64 = 0.0
    if true {
        if (!(!(!false))) {
            if true {
                for fi10 := 0; fi10 < 6; fi10++ {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        fmt.Println(fi11)
                        if ((v9 + 1e10) <= v9) {
                            var v12 uint32 = 1
                            switch "" {
                            case "hello":
                                v13 := GenS1{F0: 127}
                                _ = v13
                            }
                            _ = v12
                        }
                        v14 := GenS1{F0: 42}
                        v3 += v2
                        _ = fi11
                        _ = v14
                    }
                    _ = fi10
                }
                if false {
                    v5 = (v2 % ((v3) | 1))
                    v15 := GenS2{F0: -928, F1: -74, F2: 0}
                    v3 += (v2 & v3)
                    _ = v15
                }
            } else {
                switch 4294967295 {
                case -329:
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v3 = 288
                        _ = fi16
                    }
                case 579:
                    v17 := GenS1{F0: 255}
                    _ = v17
                case 360:
                    v9 -= v9
                default:
                    var v18 int = 2147483647
                    _ = v18
                }
                var v19 int32 = -659
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 float32 = 0.0
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        } else {
            var v22 float64 = ((1.5 + (v9 / v9)) + v9)
            if false {
                for fi23 := 0; fi23 < 19; fi23++ {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if (!true) {
                            v25 := GenS2{F0: 127, F1: 697, F2: 255}
                            _ = v25
                        }
                        delete(v4, 2)
                        if true {
                            v3 = (v5 - (v2 / (((-(v3))) | 1)))
                        } else {
                            v26 := GenS1{F0: 255}
                            v27 := GenS1{F0: 127}
                            v9 = 0.1
                            _ = v26
                            _ = v27
                        }
                        var v28 float32 = 3.14159
                        _ = fi24
                        _ = v28
                    }
                    _ = fi23
                }
                switch 463 {
                case 0:
                    v22 = 1.5
                }
                v9 -= 1.5
            } else {
                var v29 int32 = 0
                _ = v29
            }
            _ = v22
        }
        v2 -= v2
    }
    var v30 int16 = v5
    if (-2.25 > float32(1e10)) {
        v2 = v2
        v5 = (v3 * (^(v30)))
    } else {
        if true {
            for fi31 := 0; fi31 < 8; fi31++ {
                _ = len(v4)
                v9 = v9
                _ = fi31
            }
            for fi32 := 0; fi32 < 8; fi32++ {
                delete(v4, 292)
                _ = fi32
            }
        } else {
            v33 := GenS1{F0: 42}
            var v34 int8 = 100
            _ = v33
            _ = v34
        }
        var v35 int32 = -2147483648
        var v36 uint64 = 0
        _ = v35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
