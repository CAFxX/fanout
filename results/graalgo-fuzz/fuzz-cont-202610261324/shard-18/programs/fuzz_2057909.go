package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
func genFn0(p0 uint, p1 string) bool {
    v7 := []int{9223372036854775807, -9223372036854775808, -2147483648, 32768}
    fmt.Println(p1, p1)
    v8 := GenS1{F0: -111, F1: (^(p0)), F2: 9223372036854775807}
    _ = len(v7)
    p1 = "hello"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 uint32 = 2
    var v5 int64 = -1
    var v6 int8 = -11
    v9 := GenS1{F0: 1, F1: (^(v1)), F2: (v5 - v5)}
    switch (("\n\t" + "") + "hello") {
    case "foo bar":
        v10 := []int{937}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            v6 = v6
            v5 += v5
            _ = fi11
        }
    }
    v5 = v5
    fmt.Println(v1, v5, v1)
    if (("foo bar" <= "\n\t") && true) {
        var v12 uint8 = 6
        if (1e10 > 1e10) {
            var v13 float64 = 0.0
            var v14 int8 = (v6 / ((v6) | 1))
            v15 := GenS1{F0: v6, F1: 4294967295, F2: 0}
            _ = v13
            _ = v14
            _ = v15
        } else {
            v4 = v4
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := []int{fi16, 128, (fi16 * (fi16 / ((fi16) | 1))), (fi16 >> (((v1 << ((42) & 63))) & 63))}
            var v18 int = fi16
            var v19 int8 = 100
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v12
    }
    v1 -= 0
    v20 := v3[""]
    if true {
        fmt.Println(v4)
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (fi21 << ((255) & 63)) {
        case 0:
            v1 += 737
        case 127:
            if true {
                var v22 float64 = (3.14159 - 0.0)
                v20 = v20
                v23 := GenS2{F0: v4, F1: -516}
                _ = v22
                _ = v23
            } else {
                var v24 uint16 = 722
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            if ((fi25 ^ (^(fi21))) != 32767) {
                var v26 string = (((v20 + v20) + v20) + v20)
                var v27 int64 = v5
                v28 := GenS2{F0: v4, F1: -1}
                _ = v26
                _ = v27
                _ = v28
            }
            v5 = (v5 % ((v5) | 1))
            v5 = v5
            _ = fi25
        }
        v29 := []int{128, fi21}
        _ = fi21
        _ = v29
    }
    var v30 bool = (-391 != -337)
    v31 := []int{127, 128, 1}
    v32 := GenS3{F0: (^(v5)), F1: v1, F2: 277}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
