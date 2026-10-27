package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
func genFn0(p0 int64, p1 uint32) uint64 {
    p1 = p1
    p0 -= (p0 % (((^(p0))) | 1))
    v7 := []int{}
    p1 = p1
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return 575
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    var v4 float64 = 3.14159
    var v5 int32 = -2147483648
    v6 := make(map[string]int)
    v8 := GenS1{F0: 42, F1: 0}
    v3 = (-535 == 1000)
    for fi9 := 0; fi9 < 15; fi9++ {
        if v3 {
            v4 = float64(v1)
        } else {
            if v3 {
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int64 = 42
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v5, v1, v1)
                v12 := GenS2{F0: 831, F1: -1, F2: 42}
                _ = v12
            } else {
                var v13 uint = 0
                if (0 > 0) {
                    v14 := GenS2{F0: v13, F1: -1, F2: -1}
                    _ = v14
                }
                _ = v13
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                if true {
                    var v16 int16 = 0
                    _ = v16
                } else {
                    _ = len(v2)
                    v17 := GenS3{F0: -419, F1: 18446744073709551615, F2: 979}
                    if (!(125 > 2)) {
                        fmt.Println(v5, fi15, v1)
                        v18 := []int{fi15, 127}
                        _ = v18
                    } else {
                        v3 = (417 == -1)
                    }
                    _ = v17
                }
                _ = fi15
            }
        }
        v1 -= 3.14159
        v19 := []int{fi9}
        _ = fi9
        _ = v19
    }
    v3 = v3
    switch 408 {
    case 127:
        var v20 int64 = -698
        _ = v20
    case 0:
        switch 0 {
        case 128:
            var v21 float32 = (v1 * (0.1 / 1e10))
            _ = v21
        }
    }
    fmt.Println(v5)
    for fi22 := 0; fi22 < 19; fi22++ {
        switch "foo bar" {
        case "\n\t":
            for fi23 := 0; fi23 < 5; fi23++ {
                v5 = (v5 * v5)
                switch (-(fi22)) {
                case 1:
                    var v24 uint64 = 4294967295
                    _ = v24
                case 106:
                    var v25 int64 = -9223372036854775808
                    _ = v25
                case -569:
                    v1 -= 3.14159
                }
                var v26 uint32 = 2147483647
                _ = fi23
                _ = v26
            }
        }
        var v27 uint64 = 1
        var v28 int64 = 799
        _ = fi22
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 9223372036854775807, F1: 4294967295}
    var v30 uint = 18446744073709551615
    fmt.Println(v5, v1, v30)
    v5 = (v5 / ((v5) | 1))
    var v31 bool = (int(v30) != 113)
    v32 := GenS1{F0: 1, F1: 18446744073709551615}
    var v33 uint32 = 4294967295
    if (((!true) && (!(v31 && (v3 || ((v30 < v30) || v3))))) && v31) {
        delete(v6, ("\n\t" + ("\n\t" + "foo bar")))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
