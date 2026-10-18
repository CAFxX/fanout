package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string, p1 uint16) bool {
    var v4 uint64 = 4294967295
    p0 = "foo bar"
    v5 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -9223372036854775808
    var v3 uint32 = 2
    v2 = int64(v3)
    var v6 uint16 = 255
    if ((v6 >> ((2) & 63)) == (v6 | (v6 % (((v6 >> ((1000) & 63))) | 1)))) {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := v1[int(v6)]
            v2 -= (v2 << ((88) & 63))
            var v9 uint = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        if (((40 > -43) && ("\n\t" != "foo bar")) || false) {
            var v10 uint16 = (v6 * v6)
            v2 -= (v2 | (v2 + 1))
            v11 := GenS2{F0: (v3 >> ((uint(v10)) & 63))}
            _ = v10
            _ = v11
        }
        if false {
            v12 := []int{}
            _ = v12
        }
    }
    if (1 != 1) {
        v2 += v2
        fmt.Println(v6, v2, v2)
        v6 = v6
    } else {
        if (1 <= 127) {
            fmt.Println(v3)
            v3 = v3
        }
        var v13 int64 = v2
        _ = v13
    }
    v2 -= -9223372036854775808
    var v14 uint64 = 724
    if false {
        var v15 int16 = 1000
        fmt.Println(v6, v14)
        _ = v15
    } else {
        var v16 uint = uint(v14)
        v16 = v16
        fmt.Println(v3)
        _ = v16
    }
    v14 += (v14 % (((v14 / ((v14) | 1))) | 1))
    v17 := GenS2{F0: (v3 % ((v3) | 1))}
    var v18 float32 = ((0.0 - float32(1.5)) + (-2.25 / 1e10))
    fmt.Println(v3)
    fmt.Println(v2, v18, v2)
    var v19 uint64 = (^(v14))
    var v20 string = "a"
    switch "foo bar" {
    case "":
        if true {
            v14 = v19
            v3 -= (v3 - v3)
        } else {
            fmt.Println(v14, v3, v18)
            var v21 int8 = 1
            v22 := []int{1, 4294967295, 724}
            _ = v21
            _ = v22
        }
    case "hello":
        if (42 < (v2 * v2)) {
            v23 := []int{904, 0, 1}
            for fi24 := 0; fi24 < 5; fi24++ {
                fmt.Println(v20, fi24)
                var v25 bool = ((!true) || true)
                _ = fi24
                _ = v25
            }
            fmt.Println(v14, v20, v2)
            _ = v23
        }
    case "\n\t":
        v26 := []int{32768}
        _ = v26
    }
    v6 -= 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
