package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0() int {
    if true {
        switch "hello" {
        case "foo bar":
            var v7 uint64 = 4294967295
            _ = v7
        default:
            v8 := []int{162}
            _ = v8
        }
        if true {
            v9 := []int{}
            _ = v9
        }
        v10 := GenS1{F0: 65535, F1: (!true)}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
        var v13 int16 = -1
        var v14 uint64 = 4294967295
        _ = v13
        _ = v14
    }
    v15 := []int{9223372036854775807, -554, 32767}
    var v16 int8 = -1
    _ = v15
    _ = v16
    return -2147483648
}

func main() {
    v1 := []string{}
    var v2 uint = 65535
    var v3 int = 4294967295
    v4 := []int{(v3 | -671)}
    v5 := make(map[uint32]string)
    v6 := []int{(^(v3)), (v3 | v3)}
    var v17 uint16 = 65535
    v5[2147483648] = ""
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS1{F0: v17, F1: (false && (0 <= -2147483648))}
        _ = fi18
        _ = v19
    }
    var v20 uint = v2
    switch (v3 % ((genFn0()) | 1)) {
    case 4294967295:
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint = (v20 ^ 65535)
            var v23 uint32 = 2
            for fi24 := 0; fi24 < 19; fi24++ {
                _ = len(v1)
                v17 = (v17 / ((137) | 1))
                v23 += (v23 + v23)
                var v25 int8 = 0
                _ = fi24
                _ = v25
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    if (1 <= 2) {
        if (!(0 == -1)) {
            v26 := GenS1{F0: (v17 >> ((65535) & 63)), F1: (v3 == -362)}
            _ = v26
        }
        v3 += (v3 / ((32767) | 1))
    } else {
        var v27 bool = (255 <= 1)
        v28 := GenS1{F0: 368, F1: (0 > 42)}
        _ = v27
        _ = v28
    }
    v1 = append(v1, "a")
    if (238 != -9223372036854775808) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v29 int = (-(v3))
        _ = v29
    }
    var v30 uint8 = 1
    var v31 string = "\n\t"
    fmt.Println(v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
