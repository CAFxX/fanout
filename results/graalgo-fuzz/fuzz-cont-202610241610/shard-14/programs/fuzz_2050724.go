package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool) uint8 {
    p0 = (1 <= 4294967295)
    var v6 uint64 = 4294967295
    v7 := []int{197, 9223372036854775807}
    v6 = v6
    v7 = append(v7, -328)
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (42 < 150)
    v3 := make(map[string]int)
    var v4 int = 275
    var v5 string = ""
    for fi8 := 0; fi8 < 14; fi8++ {
        v2 = (255 > 255)
        var v9 int32 = 668
        var v10 string = v5
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = v2
    for fi12 := 0; fi12 < 3; fi12++ {
        if (725 < 850) {
            if (42 != 540) {
                v13 := []int{(fi12 - (^(fi12))), (v4 + v4), fi12}
                for fi14 := 0; fi14 < 12; fi14++ {
                    v1[v4] = true
                    v15 := GenS1{F0: 2, F1: 42, F2: -9223372036854775808}
                    if false {
                        var v16 int16 = 932
                        v17 := v3[("" + v5)]
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                v2 = (756 <= 18446744073709551615)
                _ = v13
            } else {
                v5 = "hello"
            }
            if (!false) {
                switch (v4 & 9223372036854775807) {
                case -2147483648:
                    v18 := GenS1{F0: 1, F1: 2, F2: 42}
                    _ = v18
                case 548:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 bool = (9223372036854775807 < -1)
                        var v21 uint16 = 614
                        v4 = v4
                        delete(v3, (v5 + (("" + (v5 + v5)) + "foo bar")))
                        v4 = 593
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case 1:
                    var v22 int64 = 1
                    _ = v22
                }
                var v23 uint32 = 2
                _ = v23
            } else {
                var v24 bool = ((((0.1 - -2.25) * 0.0) * -2.25) < float32(-2.25))
                fmt.Println(v2)
                _ = v24
            }
            if false {
                v4 -= (v4 >> ((4294967295) & 63))
                switch 436 {
                case -187:
                    v25 := []int{}
                    _ = v25
                case 792:
                    if true {
                        v4 = v4
                        v11 = (1.5 != 0.0)
                        fmt.Println(fi12, v5, fi12)
                    }
                case 128:
                    v5 = (v5 + v5)
                }
                var v26 uint64 = 4294967295
                _ = v26
            }
        } else {
            var v27 int64 = -613
            fmt.Println(v4)
            _ = v27
        }
        _ = fi12
    }
    if (6 == 2) {
        v4 = v4
    } else {
        _ = len(v3)
        var v28 int = v4
        var v29 float64 = (-2.25 * 3.14159)
        _ = v28
        _ = v29
    }
    v2 = (!(1e10 > 0.1))
    switch v4 {
    case -665:
        v30 := []int{(v4 - (v4 / ((v4) | 1))), (v4 * (^(v4))), v4}
        _ = v30
    default:
        switch v5 {
        case "hello":
            v5 = ((v5 + "\n\t") + v5)
        case "foo bar":
            v31 := []int{(v4 * 32768), -2147483648, v4, (-(v4))}
            _ = v31
        case "a":
            v32 := GenS2{F0: -1}
            _ = v32
        default:
            v33 := GenS3{F0: -890}
            _ = v33
        }
    }
    var v34 string = ((v5 + "") + ((v5 + "") + "\n\t"))
    var v35 uint8 = 128
    for fi36 := 0; fi36 < 16; fi36++ {
        v2 = v2
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
