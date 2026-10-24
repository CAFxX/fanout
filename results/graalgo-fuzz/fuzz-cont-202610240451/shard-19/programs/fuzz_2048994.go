package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 bool) int {
    for fi7 := 0; fi7 < 16; fi7++ {
        if p0 {
            fmt.Println(p0, p1)
        } else {
            p1 = p0
        }
        v8 := GenS1{F0: (fi7 + fi7)}
        _ = fi7
        _ = v8
    }
    var v9 float32 = 0.1
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{(fi10 / (((fi10 % ((0) | 1))) | 1)), (^(fi10))}
        switch "" {
        case "hello":
            fmt.Println(p1)
        case "\n\t":
            var v12 bool = ((-108 == -868) && p1)
            _ = v12
        case "a":
            var v13 int16 = -769
            _ = v13
        }
        var v14 int8 = 1
        _ = fi10
        _ = v11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return 1
}
func genFn1() uint16 {
    if false {
        v15 := []int{32767, 4294967295, genFn0(true, (-1 <= -239)), -2147483648}
        v15 = append(v15, -411)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int8 = -78
            _ = fi16
            _ = v17
        }
        var v18 int8 = 42
        fmt.Println(v18)
        _ = v18
    }
    v19 := []int{227}
    v19 = append(v19, -834)
    var v20 int16 = -1
    _ = v19
    _ = v20
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    var v3 int8 = 1
    v4 := []string{"\n\t", "hello", (("foo bar" + "") + ("foo bar" + "a")), "\n\t"}
    v5 := make(map[uint32]uint32)
    var v6 int64 = 42
    if false {
        if true {
            fmt.Println(v6, v2)
            v2 = v2
            v21 := GenS1{F0: 32768}
            _ = v21
        } else {
            fmt.Println(v3, v3)
        }
        if (128 != -9223372036854775808) {
            fmt.Println(v2, v2)
            if false {
                switch 716 {
                case 556:
                    fmt.Println(v6, v6)
                case 128:
                    fmt.Println(v3, v3)
                }
                var v22 float64 = 1.5
                _ = v22
            }
            var v23 int16 = 1000
            _ = v23
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v2 += v2
            _ = fi24
        }
    }
    fmt.Println(v3)
    for fi25 := 0; fi25 < 17; fi25++ {
        v6 = 1
        v26 := GenS1{F0: 401}
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 128}
    v3 = (^(v3))
    var v28 uint = 2
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 int8 = (v3 % ((v3) | 1))
        _ = fi29
        _ = v30
    }
    if true {
        v31 := GenS1{F0: genFn0(false, true)}
        var v32 int16 = -57
        _ = v31
        _ = v32
    }
    v33 := GenS2{F0: 1, F1: (v6 >> ((v28) & 63)), F2: -95}
    v2 -= (((v2 + 3.14159) + v2) * 0.1)
    switch 471 {
    case -952:
        if ((567 <= 42) || ((689 >= 7) || true)) {
            v1[743] = 4294967295
            var v34 int64 = v6
            _ = v34
        }
    }
    fmt.Println(v28, v6)
    v2 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
