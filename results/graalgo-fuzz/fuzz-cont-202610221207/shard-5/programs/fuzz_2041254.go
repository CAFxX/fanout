package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
func genFn0(p0 uint, p1 int64, p2 bool) bool {
    var v7 uint64 = 614
    if (65535 < 42) {
        v8 := []int{-944, 358, 32767, -2147483648}
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint64 = (v7 << ((p0) & 63))
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v7)
        switch (fi11 & fi11) {
        case 32767:
            if ((0.1 - 0.0) != float32((float64(((1.5 + 0.1) - ((float32(float64(3.14159)) + (float32((0.1 / 1e10)) * 3.14159)) / (1e10 + 0.0)))) - 1e10))) {
                fmt.Println(p0, v7, p0)
            }
        case -1:
            if (!(p0 < (^(p0)))) {
                p2 = p2
            } else {
                v7 = 1
                v12 := GenS1{F0: 2, F1: 128}
                var v13 float64 = 0.1
                _ = v12
                _ = v13
            }
        case 738:
            v14 := []int{fi11, fi11, (-(fi11)), fi11}
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            if (65535 > 0) {
                v16 := []int{}
                if false {
                    v17 := GenS2{F0: 597, F1: p2, F2: p1}
                    var v18 uint = uint(v7)
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
            _ = fi15
        }
        _ = fi11
    }
    if (!true) {
        if (p1 != p1) {
            p0 -= (p0 ^ (^(p0)))
            var v19 int8 = 42
            _ = v19
        } else {
            fmt.Println(p2, p1, p1)
            v20 := []int{2147483647, 2147483647, 4294967295}
            fmt.Println(p0, p0)
            _ = v20
        }
    } else {
        v21 := GenS2{F0: 4294967295, F1: (42 < (p1 ^ (p1 ^ 525))), F2: p1}
        p0 -= (p0 - (^(p0)))
        v22 := GenS1{F0: 15, F1: 32767}
        _ = v21
        _ = v22
    }
    var v23 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v7
    return (4294967295 == 4294967295)
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -373
    var v3 int8 = 99
    var v4 uint = 255
    var v5 int = -2147483648
    v6 := []int{v5, 32768, v5}
    v5 = -2147483648
    for fi24 := 0; fi24 < 6; fi24++ {
        v3 = (v3 / ((42) | 1))
        if len(v6) > 0 {
            v6[0] = -1
        }
        _ = fi24
    }
    switch (((("foo bar" + "\n\t") + ("" + "a")) + "hello") + ("a" + "foo bar")) {
    case "":
        v25 := v1["hello"]
        _ = v25
    }
    fmt.Println(v5)
    v26 := GenS1{F0: 0, F1: v5}
    if (498 < 566) {
        v5 = (v5 << ((v4) & 63))
    } else {
        if false {
            fmt.Println(v3)
            var v27 uint64 = 0
            _ = v27
        } else {
            for fi28 := 0; fi28 < 1; fi28++ {
                switch (fi28 % ((-887) | 1)) {
                case 4294967295:
                    v1["a"] = ("hello" + "a")
                default:
                    v5 += 283
                }
                if (0 != 128) {
                    for fi29 := 0; fi29 < 1; fi29++ {
                        _ = len(v6)
                        _ = fi29
                    }
                    _ = len(v6)
                }
                var v30 uint64 = 0
                _ = fi28
                _ = v30
            }
            v3 += -1
            if (false && false) {
                v6 = append(v6, 42)
                if len(v6) > 0 {
                    v6[0] = v5
                }
            } else {
                for fi31 := 0; fi31 < 10; fi31++ {
                    fmt.Println(v3, v3, v2)
                    _ = fi31
                }
            }
        }
    }
    v2 -= (v2 % ((v2) | 1))
    var v32 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
