package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint) int32 {
    if true {
        var v6 int64 = -963
        _ = v6
    } else {
        fmt.Println(p0)
    }
    if true {
        v7 := GenS1{F0: 4294967295}
        _ = v7
    }
    _ = p0
    return 42
}
func genFn1(p0 bool, p1 string) int8 {
    var v8 int32 = -1
    v8 -= genFn0(42)
    _ = p0
    _ = p1
    _ = v8
    return -11
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]bool)
    var v3 int8 = -1
    var v4 float32 = float32(1.5)
    var v5 int64 = 42
    v3 = genFn1(false, "")
    v9 := []int{-1, 127}
    fmt.Println(v1)
    fmt.Println(v3, v5, v1)
    var v10 int16 = 34
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(v5, v10, v3)
        switch "hello" {
        case "a":
            switch "" {
            case "a":
                v1 = v1
            default:
                if len(v9) > 0 {
                    v9[0] = (fi11 ^ fi11)
                }
            }
        }
        switch (fi11 % (((^(fi11))) | 1)) {
        case -2147483648:
            v10 += (v10 / (((v10 % ((1) | 1))) | 1))
        }
        _ = fi11
    }
    if ((4294967295 == 42) || (true || (true || (4294967295 >= 728)))) {
        if false {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint8 = 52
                var v14 int32 = -1
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 int32 = -2147483648
            if len(v9) > 0 {
                v9[0] = 885
            }
            _ = v15
        } else {
            v16 := v2[4294967295]
            var v17 int = 127
            _ = v16
            _ = v17
        }
    }
    v3 = v3
    for fi18 := 0; fi18 < 3; fi18++ {
        if (true || false) {
            var v19 int8 = genFn1(true, ("hello" + "hello"))
            _ = v19
        }
        if (2 <= 4294967295) {
            if false {
                v3 = (v3 & v3)
            }
        } else {
            fmt.Println(v10)
        }
        switch "" {
        case "foo bar":
            if true {
                var v20 float64 = float64(v4)
                _ = len(v2)
                _ = v20
            }
        case "\n\t":
            v9 = append(v9, fi18)
        case "a":
            for fi21 := 0; fi21 < 14; fi21++ {
                if (1 <= 1) {
                    switch "a" {
                    case "hello":
                        _ = len(v2)
                    }
                } else {
                    var v22 int = fi18
                    v1 = 127
                    var v23 int32 = -2147483648
                    _ = v22
                    _ = v23
                }
                var v24 int32 = -2147483648
                _ = fi21
                _ = v24
            }
        default:
            v5 += v5
        }
        _ = fi18
    }
    var v25 float64 = float64(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
