package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int16, p1 int32) uint {
    var v7 float32 = ((1e10 + 0.0) / 1e10)
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 913
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 int = 127
    v5 := make(map[int]bool)
    var v6 float32 = ((-2.25 / 0.1) / (0.0 + 0.1))
    v1 = (v1 << ((genFn0(-192, 2147483647)) & 63))
    v6 = v6
    v9 := GenS1{F0: genFn0(-32768, -2147483648)}
    var v10 uint8 = 185
    var v11 uint32 = 884
    fmt.Println(v11, v4, v1)
    var v12 int32 = 0
    for fi13 := 0; fi13 < 19; fi13++ {
        if ((0 <= 9223372036854775807) && (!true)) {
            var v14 float32 = 1e10
            if true {
                if false {
                    v15 := GenS1{F0: 340}
                    v6 -= v6
                    v16 := []int{}
                    _ = v15
                    _ = v16
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v10)
                    case "":
                        if ((!true) && false) {
                            var v17 uint16 = uint16(v10)
                            v18 := GenS1{F0: 0}
                            _ = v17
                            _ = v18
                        }
                    default:
                        v11 -= (v11 - (v11 | (^(v11))))
                    }
                    var v19 float64 = (3.14159 + 1e10)
                    v12 = v12
                    _ = v19
                }
                fmt.Println(v10)
            } else {
                var v20 float32 = float32(-2.25)
                v21 := []int{(v4 | -1), 268, -1, v4}
                var v22 uint16 = 255
                _ = v20
                _ = v21
                _ = v22
            }
            v12 -= (v12 << ((255) & 63))
            _ = v14
        }
        _ = fi13
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            _ = len(v2)
            _ = fi24
        }
        _ = fi23
    }
    v25 := GenS1{F0: 1}
    if true {
        if false {
            var v26 bool = (false || (!((v11 | 1) == v11)))
            for fi27 := 0; fi27 < 12; fi27++ {
                if (true && (!true)) {
                    var v28 uint16 = 255
                    v11 = 5
                    _ = v28
                } else {
                    if true {
                        var v29 int16 = -921
                        _ = v29
                    } else {
                        var v30 int32 = (v12 + 41)
                        v1 = (v10 + 2)
                        if v26 {
                            v31 := GenS1{F0: 255}
                            switch 32768 {
                            case 9223372036854775807:
                                v32 := v5[fi27]
                                _ = v32
                            case 100:
                                _ = len(v5)
                            }
                            var v33 uint = 1
                            _ = v31
                            _ = v33
                        }
                        _ = v30
                    }
                }
                var v34 float64 = 3.14159
                _ = fi27
                _ = v34
            }
            if v26 {
                for fi35 := 0; fi35 < 10; fi35++ {
                    fmt.Println(v4, v6, v26)
                    _ = fi35
                }
                var v36 uint8 = 0
                _ = v36
            }
            _ = v26
        }
    }
    v6 -= v6
    fmt.Println(v11, v10, v6)
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 uint8 = v10
        _ = fi37
        _ = v38
    }
    if true {
        var v39 int64 = 340
        v40 := GenS2{F0: 693, F1: int64(v12)}
        _ = v39
        _ = v40
    } else {
        v12 = (^(v12))
    }
    var v41 uint64 = 862
    fmt.Println(v12, v1)
    delete(v5, (v4 & v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
