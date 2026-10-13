package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0(p0 uint) int64 {
    fmt.Println(p0)
    var v7 bool = true
    fmt.Println(v7)
    _ = p0
    _ = v7
    return -402
}

func main() {
    var v1 uint32 = 652
    var v2 int = 127
    v3 := make(map[int]string)
    var v4 string = ("\n\t" + "a")
    var v5 int = -2147483648
    v6 := make(map[uint32]int)
    var v8 uint64 = 957
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v2, v4)
        switch "foo bar" {
        case "\n\t":
            v10 := GenS1{F0: v4, F1: 42}
            _ = v10
        case "foo bar":
            if ((v2 != (fi9 / ((0) | 1))) || (!true)) {
                _ = len(v6)
            }
        case "a":
            switch (^(v5)) {
            case 128:
                v8 = v8
            case 127:
                v5 = v2
            case 1:
                v4 = (v4 + v4)
            default:
                v11 := GenS1{F0: (v4 + "\n\t"), F1: 42}
                _ = v11
            }
        default:
            var v12 uint = 42
            _ = v12
        }
        var v13 uint16 = 42
        _ = fi9
        _ = v13
    }
    var v14 float64 = 1e10
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint64 = 18446744073709551615
        v17 := GenS1{F0: "\n\t", F1: 789}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint8 = 2
    v19 := GenS2{F0: uint(v8), F1: -1}
    var v20 int8 = -71
    switch v4 {
    case "foo bar":
        var v21 int = 1
        _ = v21
    case "\n\t":
        if false {
            var v22 int32 = 2147483647
            for fi23 := 0; fi23 < 20; fi23++ {
                v20 += (v20 << ((42) & 63))
                _ = fi23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint8 = 2
                _ = fi24
                _ = v25
            }
            _ = v22
        } else {
            if (100 > v20) {
                fmt.Println(v18, v4)
                v18 += v18
                var v26 int32 = int32(v1)
                _ = v26
            } else {
                v18 -= v18
                delete(v3, v2)
            }
            v27 := []int{(v5 & int(v8)), v2, v5, v5}
            var v28 uint8 = v18
            _ = v27
            _ = v28
        }
    }
    if false {
        fmt.Println(v8, v4)
        switch (v2 * (v5 % ((2147483647) | 1))) {
        case 734:
            switch (v2 & (v2 ^ v5)) {
            case 9223372036854775807:
                var v29 int8 = -128
                _ = v29
            case 128:
                v30 := GenS1{F0: v4, F1: 32767}
                _ = v30
            case 42:
                v1 = 2
            }
        case -2147483648:
            fmt.Println(v8)
        }
        var v31 uint8 = 127
        _ = v31
    } else {
        fmt.Println(v2, v1)
        v32 := GenS1{F0: (v4 + "hello"), F1: 255}
        switch v2 {
        case -715:
            fmt.Println(v5, v18)
        case -1:
            v6[v1] = (^(v2))
        default:
            v33 := []int{0, 2147483647}
            _ = v33
        }
        _ = v32
    }
    v34 := GenS3{F0: uint8(v1), F1: 9223372036854775807, F2: 65535}
    v2 -= v5
    fmt.Println(v4)
    var v35 float64 = v14
    switch (v5 * (v5 << ((255) & 63))) {
    case 32768:
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 uint16 = 190
            if (!((1e10 * 0.0) > 0.0)) {
                if true {
                    v4 = v4
                    v2 = fi36
                }
            }
            _ = fi36
            _ = v37
        }
    case 1:
        switch (v5 | v2) {
        case 32768:
            for fi38 := 0; fi38 < 17; fi38++ {
                var v39 int = v5
                if false {
                    v35 -= v35
                } else {
                    var v40 uint8 = v18
                    v20 -= 99
                    switch 2147483647 {
                    case -245:
                        fmt.Println(v18, v8)
                    }
                    _ = v40
                }
                _ = fi38
                _ = v39
            }
        }
    default:
        v41 := GenS2{F0: 18446744073709551615, F1: -2147483648}
        _ = v41
    }
    v42 := []int{v5}
    v2 += -896
    v18 -= 127
    if ((v4 > v4) && true) {
        var v43 uint32 = (v1 & v1)
        v44 := v3[v2]
        _ = v43
        _ = v44
    } else {
        v2 -= v2
        var v45 uint32 = v1
        var v46 uint16 = 152
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
