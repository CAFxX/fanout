package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 int, p2 uint64) int {
    v5 := []int{}
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            switch ((("foo bar" + p0) + p0) + p0) {
            case "\n\t":
                v7 := GenS1{F0: 2147483647}
                _ = v7
            default:
                if (2 <= (p2 | 2)) {
                    var v8 int16 = 920
                    _ = v8
                }
            }
            switch p0 {
            case "a":
                for fi9 := 0; fi9 < 6; fi9++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(p2, fi9, p0)
                    _ = fi9
                }
            default:
                v10 := GenS2{F0: 966, F1: -2147483648, F2: fi6}
                _ = v10
            }
        } else {
            _ = len(v5)
            if ((835 < 1000) || ((1e10 / (-2.25 * 1.5)) <= 1.5)) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := GenS2{F0: -9223372036854775808, F1: 42, F2: fi11}
                    switch p0 {
                    case "a":
                        var v13 float32 = 0.1
                        fmt.Println(p2)
                        if true {
                            var v14 int16 = -1
                            var v15 float64 = 0.0
                            _ = v14
                            _ = v15
                        } else {
                            _ = len(v5)
                        }
                        _ = v13
                    case "\n\t":
                        v16 := GenS1{F0: -957}
                        var v17 uint64 = 971
                        _ = v16
                        _ = v17
                    case "":
                        var v18 float32 = 1.5
                        _ = v18
                    default:
                        p0 = p0
                    }
                    _ = fi11
                    _ = v12
                }
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                switch (fi19 << ((2) & 63)) {
                case 128:
                    fmt.Println(p1, fi19, p0)
                case 2147483647:
                    v5 = append(v5, fi6)
                case 82:
                    var v20 int32 = -751
                    var v21 int32 = v20
                    var v22 uint8 = 42
                    _ = v20
                    _ = v21
                    _ = v22
                }
                var v23 int64 = -1
                v5 = append(v5, p1)
                v24 := GenS1{F0: -1}
                var v25 float32 = ((1.5 / -2.25) / 1.0)
                v26 := GenS1{F0: 2147483647}
                _ = fi19
                _ = v23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = fi6
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        if (-9223372036854775808 < 42) {
            if true {
                var v28 bool = false
                _ = len(v5)
                _ = v28
            }
            var v29 uint8 = 42
            _ = v29
        } else {
            var v30 uint64 = (p2 >> ((42) & 63))
            var v31 uint = 18446744073709551615
            var v32 uint64 = (^(v30))
            _ = v30
            _ = v31
            _ = v32
        }
        fmt.Println(p0, p1, fi27)
        p1 -= (-(p1))
        _ = fi27
    }
    if (0 != 2147483648) {
        var v33 uint64 = (p2 ^ p2)
        var v34 string = "hello"
        v35 := GenS1{F0: -2147483648}
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (-(p1))
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]bool)
    v3 := []uint32{42, 357, 0}
    v4 := make(map[int]bool)
    var v36 int8 = 42
    var v37 string = "hello"
    v1 -= (v1 % (((v1 ^ v1)) | 1))
    var v38 bool = true
    v4[-939] = true
    for fi39 := 0; fi39 < 20; fi39++ {
        if len(v3) > 0 {
            v3[0] = 493
        }
        switch v37 {
        case "\n\t":
            var v40 int = fi39
            _ = v40
        case "hello":
            for fi41 := 0; fi41 < 9; fi41++ {
                switch (fi41 / ((fi39) | 1)) {
                case -39:
                    var v42 float32 = (3.14159 - 0.0)
                    _ = v42
                case 0:
                    fmt.Println(v38, v1)
                }
                v43 := GenS3{F0: 255}
                _ = fi41
                _ = v43
            }
        case "a":
            v1 += (v1 - v1)
        }
        _ = fi39
    }
    fmt.Println(v1, v36, v38)
    v37 = v37
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
