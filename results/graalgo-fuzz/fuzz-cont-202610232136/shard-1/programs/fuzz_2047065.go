package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    var v1 uint64 = 408
    var v2 int = 450
    var v3 string = "hello"
    v4 := make(map[string]int)
    v5 := []uint32{42, 6}
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 = "a"
        _ = len(v6)
        v8 := GenS1{F0: -9223372036854775808, F1: 28}
        _ = fi7
        _ = v8
    }
    if true {
        var v9 int64 = -9223372036854775808
        _ = v9
    } else {
        var v10 float32 = 0.1
        for fi11 := 0; fi11 < 19; fi11++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v1 -= v1
            fmt.Println(v10)
            _ = fi11
        }
        _ = len(v5)
        _ = v10
    }
    if ((v3 + (v3 + (v3 + ""))) < v3) {
        _ = len(v5)
        v1 -= v1
        if (false || (v1 > (v1 >> ((42) & 63)))) {
            for fi12 := 0; fi12 < 14; fi12++ {
                if (!false) {
                    var v13 float32 = ((0.1 + ((3.14159 - 0.1) * 0.0)) * (-2.25 * float32((-2.25 + -2.25))))
                    _ = len(v4)
                    _ = v13
                } else {
                    switch v2 {
                    case 128:
                        v2 -= v2
                    default:
                        var v14 string = v3
                        _ = v14
                    }
                }
                var v15 uint64 = (v1 * v1)
                _ = fi12
                _ = v15
            }
            var v16 uint32 = 65535
            var v17 int8 = 1
            _ = v16
            _ = v17
        }
    }
    v18 := GenS1{F0: 9223372036854775807, F1: -27}
    for fi19 := 0; fi19 < 8; fi19++ {
        v4[(v3 + (("\n\t" + "") + v3))] = (-(v2))
        v20 := GenS1{F0: -9223372036854775808, F1: 42}
        _ = fi19
        _ = v20
    }
    v2 -= v2
    var v21 float64 = 1e10
    switch (v2 - (^(v2))) {
    case 4294967295:
        var v22 int32 = -415
        _ = v22
    case -9223372036854775808:
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                switch (fi23 + (fi23 ^ (v2 - (fi23 ^ v2)))) {
                case -1:
                    v2 = (fi23 & 832)
                case -9223372036854775808:
                    var v24 int = v2
                    _ = v24
                case 1:
                    v25 := v4[v3]
                    _ = v25
                }
                _ = len(v5)
            }
            if false {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int16 = 1000
                    if (!((255 != 4294967295) && true)) {
                        v28 := GenS1{F0: -9223372036854775808, F1: 120}
                        _ = v28
                    } else {
                        fmt.Println(v3)
                        var v29 string = v3
                        fmt.Println(fi26, fi23)
                        _ = v29
                    }
                    _ = len(v6)
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v3, fi23, v21)
            }
            v1 = (v1 % ((v1) | 1))
            _ = fi23
        }
    case -2147483648:
        v1 += 18446744073709551615
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        delete(v6, "foo bar")
        _ = fi30
    }
    v3 = v3
    var v31 float32 = 1.5
    v1 = v1
    v32 := GenS1{F0: 0, F1: 1}
    fmt.Println(v3)
    if (1 == 0) {
        var v33 float64 = v21
        _ = v33
    } else {
        for fi34 := 0; fi34 < 11; fi34++ {
            fmt.Println(fi34)
            _ = fi34
        }
        if (!(1e10 <= v31)) {
            var v35 float64 = 0.0
            if false {
                _ = len(v4)
            }
            _ = v35
        } else {
            var v36 int = v2
            _ = v36
        }
        v37 := GenS1{F0: 0, F1: 127}
        _ = v37
    }
    var v38 int16 = 1
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = (v1 & v1)
        var v39 int64 = 9223372036854775807
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
