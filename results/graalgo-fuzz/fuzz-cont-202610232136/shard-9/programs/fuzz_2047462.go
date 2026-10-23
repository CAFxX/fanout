package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0() int64 {
    var v4 string = "foo bar"
    v4 = "\n\t"
    _ = v4
    return -218
}
func genFn1(p0 int32) uint {
    v5 := GenS1{F0: 0, F1: 2, F2: 135}
    var v6 int64 = 653
    for fi7 := 0; fi7 < 3; fi7++ {
        if (1.5 != (3.14159 + ((0.1 / float32((float64(1.5) * float64(0.1)))) / 3.14159))) {
            v8 := []int{(fi7 + (fi7 >> ((0) & 63))), (fi7 >> ((412) & 63))}
            v6 = 0
            var v9 uint8 = 2
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{422}
    _ = p0
    _ = v10
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint8 = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 float64 = 1.5
        _ = fi11
        _ = v12
    }
    v3 -= 116
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: (fi13 + (-(fi13))), F1: 128, F2: 2}
        v15 := GenS1{F0: 32767, F1: (v3 >> ((997) & 63)), F2: v3}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint = 799
    var v17 float32 = (0.0 - (1.5 * 3.14159))
    if ((0 != 896) || true) {
        v1[((("hello" + "\n\t") + "") + "\n\t")] = -58
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(v3)
                var v20 float32 = v17
                _ = fi19
                _ = v20
            }
            switch -728 {
            case 128:
                v21 := GenS1{F0: 127, F1: 42, F2: 119}
                _ = v21
            case 127:
                v22 := GenS1{F0: (fi18 ^ fi18), F1: (v3 + v3), F2: (^(v3))}
                _ = v22
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint = (^(v16))
                _ = fi23
                _ = v24
            }
            _ = fi18
        }
    }
    var v25 string = "a"
    switch v25 {
    case "a":
        if false {
            v16 += v16
            for fi26 := 0; fi26 < 20; fi26++ {
                if (false && false) {
                    var v27 int32 = 507
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        fmt.Println(v25)
                        fmt.Println(fi26)
                        v29 := v2["foo bar"]
                        _ = fi28
                        _ = v29
                    }
                }
                var v30 int8 = 127
                v3 = v3
                _ = fi26
                _ = v30
            }
        }
    case "hello":
        v31 := []int{}
        _ = v31
    }
    var v32 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v32
}
