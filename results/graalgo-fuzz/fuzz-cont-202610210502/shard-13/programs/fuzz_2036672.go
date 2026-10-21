package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    var v3 int16 = -729
    v4 := make(map[uint32]string)
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 -= (v5 + (((v5 + 1.5) * v5) - v5))
        _ = fi6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        v1 = (fi7 + (v1 & (fi7 | (v1 | 42))))
        _ = fi7
    }
    var v8 uint64 = 966
    var v9 uint32 = 0
    if true {
        fmt.Println(v8, v1)
        var v10 bool = true
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 4294967295}
            switch v1 {
            case 680:
                v8 = (v8 % ((v8) | 1))
            case -1:
                v13 := v2[2147483647]
                _ = v13
            case 4294967295:
                v8 = v8
            default:
                switch (^(v1)) {
                case 42:
                    v3 = 0
                case 9223372036854775807:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        if (v10 && (false || false)) {
                            v15 := GenS2{F0: 18446744073709551615}
                            _ = v15
                        } else {
                            fmt.Println(v1, fi11, v9)
                        }
                        fmt.Println(v10)
                        _ = fi14
                    }
                default:
                    v8 = v8
                }
            }
            if (!(false || v10)) {
                switch "" {
                case "a":
                    if v10 {
                        v16 := GenS1{F0: 2}
                        _ = v16
                    }
                case "hello":
                    var v17 float32 = 0.0
                    _ = v17
                case "foo bar":
                    if (!(int(v8) <= v1)) {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            fmt.Println(v9)
                            v19 := GenS2{F0: 255}
                            v20 := GenS1{F0: v8}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        v10 = (42 > 419)
                        v3 = v3
                    } else {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            fmt.Println(v5, fi11)
                            _ = fi21
                        }
                        for fi22 := 0; fi22 < 20; fi22++ {
                            var v23 float64 = float64(0.1)
                            _ = fi22
                            _ = v23
                        }
                    }
                default:
                    v24 := []int{-383, fi11}
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 15; fi25++ {
                    switch (fi11 | fi25) {
                    case 4294967295:
                        v10 = false
                    default:
                        v9 += (v9 << ((0) & 63))
                    }
                    switch 308 {
                    case 127:
                        fmt.Println(fi25)
                    }
                    _ = fi25
                }
                v26 := []int{(v1 << ((1) & 63)), v1, fi11}
                _ = v26
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v27 float64 = 0.1
    var v28 int64 = 660
    v29 := GenS1{F0: v8}
    switch (v1 & int(v3)) {
    case 9223372036854775807:
        fmt.Println(v27)
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        _ = len(v2)
        if (v9 <= 0) {
            v8 = (^(v8))
            v3 = v3
            if true {
                fmt.Println(fi30)
                v31 := GenS1{F0: 645}
                v32 := GenS1{F0: v8}
                _ = v31
                _ = v32
            } else {
                v33 := GenS1{F0: (v8 % ((0) | 1))}
                _ = v33
            }
        }
        _ = fi30
    }
    v27 += ((v27 - (v27 + -2.25)) * v27)
    var v34 uint8 = 127
    v35 := GenS1{F0: v8}
    var v36 float32 = float32((v27 / (v27 + v27)))
    v37 := []int{(v1 * v1), (v1 / ((681) | 1))}
    fmt.Println(v1, v36, v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
