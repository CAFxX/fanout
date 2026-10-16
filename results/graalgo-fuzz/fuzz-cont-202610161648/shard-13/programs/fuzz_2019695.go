package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0(p0 int32, p1 uint) uint16 {
    var v6 int64 = 290
    var v7 bool = ((255 != 1) && true)
    p0 += (p0 << ((67) & 63))
    v6 -= v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 37
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!(48 <= -878))
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 uint16 = 255
    if v2 {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint32 = 0
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: v2}
        for fi11 := 0; fi11 < 11; fi11++ {
            v5 = (v5 ^ v5)
            v2 = v2
            v2 = v2
            _ = fi11
        }
        _ = v10
    } else {
        var v12 bool = (((1e10 + -2.25) * (-2.25 - 1e10)) == (0.1 - (0.0 / ((float32(1.5) - ((float32(-2.25) + 1.5) / float32(-2.25))) + 1.5))))
        v2 = v12
        v13 := []int{}
        _ = v12
        _ = v13
    }
    fmt.Println(v5, v5)
    switch -9223372036854775808 {
    case 295:
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS2{F0: 127, F1: (false || v2)}
            _ = fi14
            _ = v15
        }
    }
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            if (1.5 > float64(float32(3.14159))) {
                v2 = ((^(v5)) > v5)
                if v2 {
                    var v17 int32 = -2147483648
                    if (v2 && true) {
                        fmt.Println(v5, fi16)
                        v4[("\n\t" + ((("\n\t" + "\n\t") + "\n\t") + ""))] = v2
                        var v18 int32 = v17
                        _ = v18
                    } else {
                        var v19 int = (fi16 % ((fi16) | 1))
                        fmt.Println(v2, v19, v5)
                        var v20 uint8 = 210
                        _ = v19
                        _ = v20
                    }
                    _ = v17
                } else {
                    var v21 int64 = -9223372036854775808
                    if (17 < 0) {
                        var v22 bool = (42 != 505)
                        var v23 uint = 255
                        if v2 {
                            v24 := v4["foo bar"]
                            v25 := []int{fi16, fi16, (fi16 + 500), fi16}
                            _ = v24
                            _ = v25
                        } else {
                            v21 = (v21 ^ 118)
                            var v26 uint = v23
                            v23 -= (v23 << (((v26 / ((0) | 1))) & 63))
                            _ = v26
                        }
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
            }
            v27 := []int{fi16, -834}
            _ = fi16
            _ = v27
        }
    } else {
        v28 := GenS3{F0: 128, F1: v2}
        _ = v28
    }
    v29 := []int{1, 2147483647}
    var v30 uint64 = 0
    if (v2 && (1 >= 255)) {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 bool = v2
            v30 = v30
            if (!((-1 < 9223372036854775807) || v32)) {
                var v33 string = "hello"
                _ = v33
            } else {
                v34 := GenS1{F0: (v32 || (v5 < (^(v5))))}
                switch "foo bar" {
                case "foo bar":
                    delete(v1, 2)
                case "hello":
                    var v35 int32 = 965
                    _ = v35
                case "a":
                    var v36 uint32 = 65535
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 17; fi37++ {
                        var v38 uint64 = 4294967295
                        _ = fi37
                        _ = v38
                    }
                }
                fmt.Println(v5, v5, v32)
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    }
    if len(v29) > 0 {
        v29[0] = 42
    }
    var v39 int8 = 127
    if v2 {
        var v40 uint16 = v5
        _ = v40
    }
    v41 := GenS2{F0: -1, F1: v2}
    _ = len(v29)
    v39 = v39
    var v42 float32 = 1e10
    if v2 {
        if true {
            v39 = (v39 | (v39 * -128))
        } else {
            delete(v1, 572)
            v5 -= v5
        }
        v43 := v4["\n\t"]
        _ = v43
    }
    fmt.Println(v5, v2, v39)
    var v44 bool = ((-1 != 42) || (v2 && v2))
    var v45 int8 = v39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v5
}
