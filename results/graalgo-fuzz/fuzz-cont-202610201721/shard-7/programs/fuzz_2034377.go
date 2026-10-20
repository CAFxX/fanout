package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0() uint32 {
    v4 := GenS1{F0: 65535, F1: (-550 < 9223372036854775807)}
    var v5 uint16 = 43
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []string{}
    var v2 float64 = (-2.25 - ((float64((3.14159 / 3.14159)) - 1e10) / 1.0))
    var v3 int16 = -32768
    if (!(128 != 783)) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int64 = -795
            v2 = (v2 + float64(3.14159))
            _ = fi8
            _ = v9
        }
        v2 += 0.1
        v2 = v2
    }
    v1 = append(v1, "\n\t")
    v10 := GenS2{F0: 42, F1: -999, F2: 160}
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                if (((1e10 - float32((0.1 / (v2 - 0.0)))) >= (0.0 - float32(v2))) && true) {
                    v13 := GenS1{F0: 711, F1: false}
                    if false {
                        var v14 int16 = (v3 ^ -32768)
                        v15 := GenS1{F0: 1, F1: (false || ((fi11 == (fi12 - 1)) && false))}
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        _ = v14
                        _ = v15
                    } else {
                        v3 = v3
                    }
                    _ = v13
                }
                v3 -= (v3 + (^(v3)))
                var v16 int16 = v3
                _ = v16
            }
            if ((127 > fi12) || (0.0 >= (1e10 + -2.25))) {
                var v17 string = ""
                v3 = v3
                var v18 uint32 = 65535
                _ = v17
                _ = v18
            } else {
                if (42 <= v3) {
                    var v19 uint = 42
                    _ = v19
                }
            }
            v2 -= v2
            v2 += v2
            _ = fi12
        }
        switch "hello" {
        case "hello":
            if (false || (true && (!true))) {
                v20 := GenS2{F0: 893, F1: -1, F2: 42}
                fmt.Println(v2)
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v2, v3)
                case "\n\t":
                    var v21 uint64 = 18446744073709551615
                    _ = v21
                }
                _ = v20
            } else {
                v3 = v3
                fmt.Println(fi11)
                _ = len(v1)
            }
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        if true {
            fmt.Println(v3, fi11, fi11)
        } else {
            v3 = 1000
        }
        _ = fi11
    }
    v1 = append(v1, (("hello" + "") + "a"))
    v22 := GenS1{F0: 496, F1: ((18446744073709551615 >= 42) && (42 < 527))}
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int64 = 42
        v25 := GenS2{F0: 18446744073709551615, F1: -1, F2: 18446744073709551615}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 string = "a"
    fmt.Println(v3, v2, v2)
    v27 := GenS2{F0: 42, F1: 1, F2: 331}
    var v28 int32 = -2147483648
    v29 := GenS3{F0: 42, F1: 271, F2: 1}
    var v30 string = v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
