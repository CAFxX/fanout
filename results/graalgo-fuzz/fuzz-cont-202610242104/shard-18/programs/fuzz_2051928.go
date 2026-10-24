package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 int8, p1 int8) int {
    v6 := []int{32768, 42}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (false || (false || true)) {
            p0 = (p1 << ((4294967295) & 63))
            for fi8 := 0; fi8 < 12; fi8++ {
                p0 = 42
                _ = fi8
            }
        }
        v6 = append(v6, fi7)
        _ = fi7
    }
    _ = len(v6)
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}
func genFn1() int64 {
    var v9 int64 = 42
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 bool = (434 == -755)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v9 = (v9 / ((1) | 1))
    if (true && true) {
        var v13 uint8 = 185
        _ = v13
    }
    var v14 uint8 = 0
    _ = v14
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int8 = -50
    var v4 int32 = 1
    var v5 uint8 = 42
    var v15 uint32 = 939
    v16 := GenS1{F0: 255, F1: false, F2: 139}
    v1["a"] = ((2 == 42) || false)
    switch ("hello" + "") {
    case "":
        v5 = 127
    case "a":
        fmt.Println(v3, v3, v3)
    default:
        if false {
            switch 128 {
            case 128:
                if (!false) {
                    v17 := []int{0, 32768, 42}
                    _ = v17
                } else {
                    var v18 int32 = (-(v4))
                    for fi19 := 0; fi19 < 15; fi19++ {
                        delete(v2, (fi19 ^ fi19))
                        var v20 float32 = 3.14159
                        var v21 int16 = 42
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                }
            default:
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 bool = true
                    _ = fi22
                    _ = v23
                }
            }
            var v24 int = 0
            v15 = v15
            _ = v24
        }
    }
    v25 := v1[((("" + "hello") + "a") + "hello")]
    v5 = (v5 >> ((658) & 63))
    v5 += 0
    v26 := v2[-353]
    var v27 int = -9223372036854775808
    v28 := []int{v27, v27, (v27 >> ((0) & 63))}
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            fmt.Println(v5, v26, v25)
            _ = len(v1)
            _ = fi30
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 uint = 234
        _ = fi31
        _ = v32
    }
    v33 := GenS1{F0: 1, F1: (!((^(v27)) > (v27 % (((-(v27))) | 1)))), F2: 391}
    for fi34 := 0; fi34 < 19; fi34++ {
        v35 := GenS2{F0: 4294967295, F1: (^(v3))}
        var v36 int = (v27 % ((-127) | 1))
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println(v3, v25, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
