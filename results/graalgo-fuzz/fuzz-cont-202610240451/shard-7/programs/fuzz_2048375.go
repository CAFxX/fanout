package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 bool; F1 bool }
func genFn0(p0 uint) uint {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint16 = 0
        if (false || ((true && (562 >= 1)) || false)) {
            v6 := []int{int(p0), fi4}
            v5 = 0
            _ = v6
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                if (!((("\n\t" + ("foo bar" + "hello")) + "hello") > "foo bar")) {
                    fmt.Println(fi7, fi7, v5)
                    if true {
                        v8 := []int{fi7, (fi4 | (-(fi4))), int(v5), fi4}
                        switch (fi4 * (fi4 - (-(fi7)))) {
                        case 2147483647:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case -9223372036854775808:
                            var v9 int16 = -583
                            _ = v9
                        }
                        _ = v8
                    }
                }
                if (147 >= 2147483647) {
                    v5 -= v5
                    if (-1 > -128) {
                        v10 := GenS1{F0: false, F1: 1, F2: 100}
                        v11 := []int{}
                        _ = v10
                        _ = v11
                    }
                }
                fmt.Println(p0)
                v5 -= 42
                v12 := GenS2{F0: 250, F1: 42, F2: 0}
                v13 := []int{(fi4 << ((42) & 63)), (fi4 % ((128) | 1))}
                v5 -= v5
                var v14 uint64 = 4294967295
                _ = fi7
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: 1, F1: -9223372036854775808, F2: (v5 & (^(v5)))}
            if ("hello" <= "") {
                var v16 int8 = -128
                var v17 uint64 = 4294967295
                _ = v16
                _ = v17
            } else {
                v18 := []int{}
                v18 = append(v18, (^(fi4)))
                p0 = (p0 % ((p0) | 1))
                _ = v18
            }
            _ = v15
        }
        _ = fi4
        _ = v5
    }
    v19 := []int{0, 127}
    p0 = 603
    _ = p0
    _ = v19
    return p0
}

func main() {
    v1 := []int{4294967295}
    var v2 float64 = ((1.5 - ((-2.25 / (1.5 - 1e10)) - 1e10)) - 0.0)
    var v3 int16 = -32768
    for fi20 := 0; fi20 < 10; fi20++ {
        v3 += (v3 / ((v3) | 1))
        _ = fi20
    }
    v2 = v2
    var v21 float64 = v2
    fmt.Println(v21, v3)
    switch 228 {
    case 9223372036854775807:
        var v22 uint64 = 791
        _ = v22
    case 127:
        var v23 uint = 42
        _ = v23
    case 42:
        switch 51 {
        case 128:
            fmt.Println(v21, v3, v2)
        }
    }
    v24 := GenS3{F0: 81, F1: 1}
    if len(v1) > 0 {
        v1[0] = 1
    }
    switch "\n\t" {
    case "hello":
        v2 -= v2
    }
    v25 := GenS4{F0: (0 >= 0), F1: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
}
