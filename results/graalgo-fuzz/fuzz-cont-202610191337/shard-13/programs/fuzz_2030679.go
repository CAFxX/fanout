package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0(p0 int64) string {
    v6 := []int{1, -631}
    switch 0 {
    case 0:
        var v7 int16 = -119
        _ = v7
    case 9223372036854775807:
        p0 -= -9223372036854775808
    case -1:
        var v8 uint = 18446744073709551615
        _ = v8
    default:
        if false {
            p0 = p0
            if true {
                p0 += (^(p0))
            } else {
                var v9 int32 = 272
                v10 := GenS1{F0: 128, F1: "\n\t", F2: 965}
                v11 := GenS2{F0: -973, F1: "\n\t"}
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 string = "\n\t"
            _ = v12
        }
    }
    _ = p0
    _ = v6
    return "hello"
}
func genFn1() int16 {
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS2{F0: 1000, F1: "hello"}
        _ = fi13
        _ = v14
    }
    var v15 int = 361
    var v16 uint8 = 128
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = v15
        fmt.Println(fi17)
        _ = fi17
        _ = v18
    }
    var v19 int64 = 0
    _ = v15
    _ = v16
    _ = v19
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -9223372036854775808
    var v3 string = "a"
    var v4 string = ((v3 + (v3 + v3)) + ("foo bar" + "a"))
    v5 := make(map[int]uint32)
    v2 = 42
    var v20 float64 = -2.25
    var v21 float64 = (0.1 - -2.25)
    var v22 int8 = 100
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v21)
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := []int{(fi23 % ((fi24) | 1)), (fi24 / (((fi24 ^ (fi23 ^ -2147483648))) | 1)), -9223372036854775808, fi23}
            var v26 int = 32768
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 string = (("a" + "hello") + "hello")
        _ = fi23
        _ = v27
    }
    var v28 float64 = float64((float32(0.0) * (float32(((v20 / v20) + 0.0)) - 0.1)))
    switch v4 {
    case "a":
        var v29 uint = 502
        _ = v29
    case "":
        var v30 int8 = v22
        _ = v30
    default:
        switch v3 {
        case "hello":
            var v31 string = v3
            _ = v31
        case "":
            v3 = v4
        case "\n\t":
            v28 = v28
        default:
            switch (v4 + (v4 + ("" + "\n\t"))) {
            case "foo bar":
                v32 := GenS2{F0: -36, F1: (v4 + (v4 + v3))}
                _ = v32
            case "":
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := []int{}
                    v20 = (float64(0.1) + v20)
                    _ = fi33
                    _ = v34
                }
            }
        }
    }
    v35 := GenS3{F0: 478, F1: -2147483648}
    v4 = v4
    var v36 uint = 997
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint = v36
        v39 := GenS1{F0: 1, F1: "a", F2: (v36 / (((v38 >> ((42) & 63))) | 1))}
        _ = fi37
        _ = v38
        _ = v39
    }
    if ((1e10 * v20) >= (v20 + v28)) {
        var v40 float32 = 0.1
        v28 += float64((((-2.25 / (v40 / v40)) - v40) - 1.5))
        _ = v40
    } else {
        switch v4 {
        case "foo bar":
            v41 := []int{811, 2147483647, 127}
            _ = v41
        case "\n\t":
            v36 += v36
        case "hello":
            var v42 uint16 = 2
            _ = v42
        }
    }
    v36 -= v36
    v22 -= (^(v22))
    switch 127 {
    case 3:
        v3 = (("foo bar" + (((v3 + v4) + ("foo bar" + v3)) + "")) + v3)
    case -849:
        if ((v22 | v22) <= 1) {
            v43 := GenS2{F0: 0, F1: (v4 + v4)}
            var v44 uint32 = 752
            v45 := []int{4294967295}
            _ = v43
            _ = v44
            _ = v45
        }
    }
    if false {
        switch 128 {
        case 42:
            if (819 >= 1000) {
                _ = len(v5)
                var v46 int = 2147483647
                _ = v46
            }
        case 127:
            v2 = (v2 ^ (v2 & v2))
        default:
            v22 = 100
        }
    } else {
        switch (v4 + v3) {
        case "hello":
            var v47 uint = v36
            _ = v47
        case "":
            if true {
                var v48 uint32 = 42
                _ = v48
            }
        }
        var v49 string = (v3 + v3)
        _ = len(v1)
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
