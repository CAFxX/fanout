package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0(p0 uint32, p1 uint8, p2 uint64) int16 {
    v6 := GenS1{F0: -1, F1: 42, F2: 42}
    p2 -= 314
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    if false {
        v9 := []int{4294967295, -577, 42, 32768}
        if len(v9) > 0 {
            v9[0] = 32767
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -617
}

func main() {
    var v1 uint = 65535
    var v2 int32 = -696
    v3 := make(map[int]int)
    var v4 bool = (4294967295 == (v1 % (((v1 >> ((v1) & 63))) | 1)))
    var v5 bool = true
    var v10 uint8 = 128
    if true {
        for fi11 := 0; fi11 < 7; fi11++ {
            switch "" {
            case "hello":
                v1 = v1
            case "a":
                var v12 int16 = -221
                _ = v12
            case "":
                for fi13 := 0; fi13 < 18; fi13++ {
                    switch ("a" + "a") {
                    case "a":
                        if true {
                            fmt.Println(v1, fi13)
                            var v14 int64 = 9223372036854775807
                            _ = v14
                        }
                    case "foo bar":
                        v15 := GenS1{F0: genFn0(2147483648, 1, 18446744073709551615), F1: (v10 * v10), F2: 42}
                        _ = v15
                    default:
                        fmt.Println(v2)
                    }
                    _ = fi13
                }
            }
            _ = fi11
        }
        var v16 uint = uint(v10)
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 int64 = 504
        _ = fi17
        _ = v18
    }
    if ((v1 & v1) == v1) {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS1{F0: 1000, F1: (v10 / ((v10) | 1)), F2: 0}
            switch ("\n\t" + "\n\t") {
            case "":
                v2 = v2
            default:
                v21 := []int{(fi19 | (fi19 & (fi19 * fi19))), 127, fi19, (fi19 ^ (fi19 * fi19))}
                _ = v21
            }
            if (((125 != 65535) || ((^(v10)) < v10)) || true) {
                fmt.Println(v5, v4, v1)
            }
            _ = fi19
            _ = v20
        }
        fmt.Println(v2)
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v4 = (0.1 <= -2.25)
            v24 := v3[fi22]
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v25 := v3[128]
    v26 := GenS1{F0: 1000, F1: v10, F2: 247}
    var v27 float64 = 3.14159
    for fi28 := 0; fi28 < 16; fi28++ {
        v1 = v1
        v5 = false
        _ = fi28
    }
    v29 := GenS1{F0: 32767, F1: (v10 << ((v1) & 63)), F2: 0}
    var v30 uint8 = 127
    v31 := v3[(v25 % (((^(v25))) | 1))]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
