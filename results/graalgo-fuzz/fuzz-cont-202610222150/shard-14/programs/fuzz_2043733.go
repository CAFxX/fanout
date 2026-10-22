package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
func genFn0() int {
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int = fi7
            v9 := []int{}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 float32 = 3.14159
        _ = v10
    } else {
        if (false || (!false)) {
            v11 := []int{128}
            _ = v11
        } else {
            var v12 uint16 = 2
            _ = v12
        }
        var v13 int = -693
        var v14 string = ""
        _ = v13
        _ = v14
    }
    var v15 int8 = -128
    v16 := []int{42, 2147483647, 9223372036854775807, -9223372036854775808}
    _ = v15
    _ = v16
    return 127
}

func main() {
    var v1 int = 0
    var v2 uint = 42
    var v3 int32 = 132
    v4 := []int{}
    var v5 int64 = 1
    var v6 int = (v1 % (((^(v1))) | 1))
    v17 := GenS1{F0: ("a" + "hello")}
    for fi18 := 0; fi18 < 5; fi18++ {
        v6 = 329
        v19 := GenS1{F0: "hello"}
        _ = fi18
        _ = v19
    }
    if (!false) {
        fmt.Println(v3, v6)
        var v20 int = v6
        _ = v20
    } else {
        var v21 string = "foo bar"
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS1{F0: "\n\t"}
            v6 = fi22
            v4 = append(v4, 829)
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS1{F0: v21}
            var v26 uint16 = 2
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v21
    }
    var v27 uint64 = 18446744073709551615
    v27 = 2
    v4 = append(v4, genFn0())
    for fi28 := 0; fi28 < 13; fi28++ {
        if (32767 > 40) {
            for fi29 := 0; fi29 < 14; fi29++ {
                v4 = append(v4, fi29)
                v30 := GenS1{F0: "hello"}
                var v31 uint = (v2 * (v2 >> ((2) & 63)))
                v31 -= v2
                v4 = append(v4, genFn0())
                _ = fi29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                fmt.Println(v1, fi32, v27)
                var v33 bool = (v27 >= (^(v27)))
                fmt.Println(v5, v3)
                v33 = (((v27 & (v27 * 42)) != 2) || v33)
                _ = fi32
                _ = v33
            }
        } else {
            v4 = append(v4, (v6 << ((65535) & 63)))
        }
        v34 := GenS2{F0: 0, F1: v3, F2: 2}
        if true {
            v35 := GenS1{F0: "hello"}
            _ = v35
        }
        _ = fi28
        _ = v34
    }
    fmt.Println(v6)
    if (true || (true && (551 >= 1))) {
        v27 = (v27 & v27)
        for fi36 := 0; fi36 < 18; fi36++ {
            if len(v4) > 0 {
                v4[0] = 128
            }
            _ = fi36
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 uint16 = 101
            _ = fi37
            _ = v38
        }
    }
    v2 = v2
    var v39 float32 = 0.0
    fmt.Println(v6, v6)
    var v40 int64 = v5
    v4 = append(v4, v6)
    var v41 bool = true
    var v42 uint8 = 101
    v6 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
