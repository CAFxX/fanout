package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 string, p1 uint8) uint64 {
    v7 := []int{32768, 32767, 42}
    p0 = ((("foo bar" + "hello") + p0) + ("foo bar" + p0))
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]int)
    v3 := []string{("hello" + "a"), "foo bar", "foo bar"}
    var v4 uint64 = 2
    v5 := make(map[string]string)
    var v6 int8 = -128
    v6 = (^(v6))
    v8 := GenS1{F0: 65535, F1: (v6 - (^(v6))), F2: "a"}
    var v9 int64 = -9223372036854775808
    fmt.Println(v4)
    v10 := GenS1{F0: 981, F1: (^(v6)), F2: (("a" + ((("a" + "foo bar") + "\n\t") + "\n\t")) + (("foo bar" + "\n\t") + ("" + "")))}
    if false {
        v9 -= v9
    } else {
        if (!true) {
            v11 := GenS1{F0: 65535, F1: v6, F2: "hello"}
            fmt.Println(v4, v1)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                v4 = (^(v4))
                var v13 uint32 = 2
                _ = fi12
                _ = v13
            }
            fmt.Println(v9, v1, v1)
            var v14 int = -669
            _ = v14
        }
    }
    if (("hello" + "a") > "a") {
        v4 = v4
        for fi15 := 0; fi15 < 9; fi15++ {
            _ = len(v3)
            v16 := GenS1{F0: 1, F1: 100, F2: (("" + "foo bar") + "")}
            _ = fi15
            _ = v16
        }
    }
    var v17 uint = 255
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint32 = 1
        if (1.5 != (-2.25 + ((0.1 * 0.1) * 1.5))) {
            v20 := GenS2{F0: 65535, F1: 643}
            var v21 int = (fi18 | fi18)
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v17)
        for fi23 := 0; fi23 < 5; fi23++ {
            if true {
                _ = len(v3)
            }
            _ = fi23
        }
        v3 = append(v3, "")
        _ = fi22
    }
    fmt.Println(v4, v9)
    if false {
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((false && true) || false) {
                var v25 float64 = 0.1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v25
            }
            _ = fi24
        }
    }
    var v26 int8 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
