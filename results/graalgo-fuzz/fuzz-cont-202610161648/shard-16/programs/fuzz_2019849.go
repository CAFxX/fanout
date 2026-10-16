package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
func genFn0(p0 uint16, p1 uint64, p2 uint64) int16 {
    var v7 float64 = (((3.14159 / 1e10) / 3.14159) * ((float64(-2.25) * 0.1) * (3.14159 * float64(-2.25))))
    if (2147483647 <= -2147483648) {
        v8 := []int{42, 42, 0}
        _ = v8
    } else {
        var v9 uint32 = 2
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint64 = (^(p1))
            v7 = (3.14159 + 0.0)
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{915, 425, 42}
    var v3 bool = true
    v4 := make(map[int]bool)
    v5 := []int{-9223372036854775808, -201, 127}
    v6 := make(map[string]string)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            if (fi12 <= 472) {
                if (118 == 32767) {
                    v3 = (!v3)
                    var v14 int8 = -1
                    var v15 uint16 = 1
                    _ = v14
                    _ = v15
                }
                var v16 uint = 704
                if false {
                    v3 = v3
                    v16 += v16
                } else {
                    var v17 int32 = -987
                    var v18 int32 = -1
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
            _ = fi13
        }
        _ = len(v4)
        switch "" {
        case "foo bar":
            var v19 uint16 = 46
            _ = v19
        default:
            delete(v6, "foo bar")
        }
        _ = fi12
    }
    v3 = v3
    v20 := GenS1{F0: 1, F1: v3, F2: 208}
    v3 = (!v3)
    _ = len(v6)
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS1{F0: -22, F1: (-32768 <= genFn0(1, 1, 0)), F2: 2}
        switch (^(fi21)) {
        case 675:
            var v23 uint8 = 2
            _ = v23
        case 128:
            fmt.Println(v3, v3)
        case 127:
            if ((!(false && (-2.25 < 0.0))) && (0 >= 1)) {
                var v24 uint8 = 2
                _ = v24
            } else {
                v25 := GenS1{F0: -128, F1: (true && v3), F2: 32767}
                var v26 bool = false
                v26 = v3
                _ = v25
                _ = v26
            }
        }
        if (1.5 >= (0.0 - 1e10)) {
            v3 = (v3 || ((!(true && (385 >= 18446744073709551615))) || ("\n\t" != "hello")))
            v3 = (!true)
            v27 := v1[(fi21 / ((fi21) | 1))]
            _ = v27
        }
        _ = fi21
        _ = v22
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v3 = (80 != 0)
        fmt.Println(v3)
        _ = fi28
    }
    v3 = true
    var v29 uint32 = 65535
    v30 := GenS1{F0: -8, F1: true, F2: 2}
    fmt.Println(v3, v3)
    v4[32768] = v3
    fmt.Println(v29, v29)
    v29 = (v29 ^ v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
