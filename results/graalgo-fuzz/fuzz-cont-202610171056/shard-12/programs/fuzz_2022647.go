package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 uint16, p1 uint64, p2 int) uint32 {
    if true {
        if (false && true) {
            var v6 uint64 = p1
            var v7 float64 = -2.25
            _ = v6
            _ = v7
        }
        var v8 int64 = -668
        _ = v8
    } else {
        v9 := GenS1{F0: -32768}
        var v10 float64 = ((0.1 + (0.1 - 3.14159)) + float64(1e10))
        v11 := []int{p2, 128, (-(p2))}
        _ = v10
        _ = v11
        _ = v9
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 246
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]int)
    var v3 uint = 42
    var v4 bool = (!true)
    var v5 int8 = 77
    v12 := GenS2{F0: int16(v1), F1: 973, F2: (^(v1))}
    v1 += 2
    var v13 int16 = 42
    v4 = false
    if (0.1 > 3.14159) {
        var v14 float64 = ((3.14159 + 3.14159) / (((0.1 / 0.1) / 1e10) * (0.1 - ((0.1 * 1.5) + 3.14159))))
        if v4 {
            var v15 uint = (v3 % ((v3) | 1))
            _ = v15
        } else {
            v16 := []int{138, -527, -9223372036854775808}
            _ = v16
        }
        if (915 >= v1) {
            v2[1] = 430
        } else {
            var v17 int16 = v13
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS2{F0: (v13 >> (((^(v3))) & 63)), F1: 1, F2: v1}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v17
        }
        _ = v14
    } else {
        if (128 <= 201) {
            v5 = int8(v1)
            v4 = v4
            switch "\n\t" {
            case "a":
                v1 += (^(v1))
            default:
                var v21 uint = v3
                _ = v21
            }
        } else {
            v22 := []int{}
            var v23 uint64 = 949
            _ = v22
            _ = v23
        }
        switch 2147483647 {
        case -2147483648:
            if v4 {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    fmt.Println(v3, v13)
                case "":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v4 = true
                        var v25 float64 = 3.14159
                        v2[32768] = 9223372036854775807
                        _ = fi24
                        _ = v25
                    }
                }
            } else {
                v26 := []int{}
                switch (("a" + "foo bar") + "a") {
                case "foo bar":
                    fmt.Println(v5)
                }
                _ = v26
            }
        default:
            v27 := GenS1{F0: 1000}
            _ = v27
        }
        _ = len(v2)
    }
    var v28 int8 = v5
    var v29 float64 = 1e10
    fmt.Println(v3, v28, v29)
    v30 := []int{4294967295, 4294967295, 32767}
    v4 = (-361 != 9223372036854775807)
    _ = len(v2)
    var v31 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
