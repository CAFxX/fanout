package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int16; F2 int8 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 int32 = 762
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint32 = 2147483648
        fmt.Println(v7, fi6, fi6)
        _ = fi6
        _ = v7
    }
    if (v5 >= v5) {
        for fi8 := 0; fi8 < 14; fi8++ {
            v5 += (v5 * v5)
            var v9 int32 = (v5 - -1)
            _ = fi8
            _ = v9
        }
    }
    if (1 >= 18446744073709551615) {
        if false {
            _ = len(v2)
            var v10 string = "a"
            var v11 int64 = -827
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: 18446744073709551615}
            v13 := GenS2{F0: 42}
            _ = v12
            _ = v13
        }
        var v14 uint32 = 0
        _ = v14
    } else {
        v5 += 376
        var v15 int = 32768
        _ = v15
    }
    switch 0 {
    case 32767:
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint32 = 1
            v17 += (v17 ^ (^(v17)))
            _ = fi16
            _ = v17
        }
    case 9223372036854775807:
        var v18 bool = false
        _ = v18
    case 238:
        v19 := GenS3{F0: 0, F1: -868, F2: 2}
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        switch "foo bar" {
        case "foo bar":
            var v21 string = "a"
            _ = v21
        case "hello":
            v22 := GenS4{F0: fi20}
            _ = v22
        case "":
            if (v5 <= (^(v5))) {
                var v23 float64 = 1e10
                for fi24 := 0; fi24 < 17; fi24++ {
                    switch (-(fi20)) {
                    case 42:
                        fmt.Println(v23, fi20)
                        var v25 float64 = v23
                        _ = v25
                    case 32768:
                        v23 = 0.1
                    default:
                        var v26 float64 = 1e10
                        _ = v26
                    }
                    v27 := v2[(fi20 << ((987) & 63))]
                    _ = fi24
                    _ = v27
                }
                v28 := GenS5{F0: 1}
                _ = v23
                _ = v28
            }
        }
        v29 := GenS6{F0: 737, F1: 32767, F2: 127}
        _ = fi20
        _ = v29
    }
    v5 = v5
    var v30 int8 = 100
    v31 := GenS7{F0: 66}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
