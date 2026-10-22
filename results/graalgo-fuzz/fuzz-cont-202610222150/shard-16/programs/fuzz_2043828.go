package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0(p0 int16) int8 {
    v6 := []int{128}
    fmt.Println(p0)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint32 = 166
            fmt.Println(v9, fi7)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return -79
}

func main() {
    var v1 float32 = ((0.0 / (-2.25 * 0.1)) - (1e10 + 1e10))
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 int16 = -1
    var v5 float64 = (((1e10 * float64(v1)) / 1.0) / (1.5 / 0.1))
    v10 := []int{}
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            if false {
                switch -9223372036854775808 {
                case -540:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 bool = false
                        _ = fi12
                        _ = v13
                    }
                }
                v4 -= 1000
            }
            v14 := GenS1{F0: "\n\t", F1: 42}
            v15 := v3[2147483648]
            _ = v14
            _ = v15
        } else {
            if false {
                v2 = v2
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v16 float32 = ((v1 * 3.14159) + 1e10)
                fmt.Println(fi11, fi11)
                _ = v16
            }
            var v17 float32 = v1
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v4 = (v4 & (v4 + (v4 ^ (v4 * v4))))
            fmt.Println(v5, v2, v5)
            var v19 uint16 = 873
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: ((("" + "hello") + (("" + "a") + "")) + "hello"), F1: 42}
        _ = fi11
        _ = v20
    }
    delete(v3, (v2 << ((471) & 63)))
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS1{F0: "a", F1: 196}
            switch "foo bar" {
            case "":
                var v24 uint32 = (v2 - (v2 >> ((255) & 63)))
                _ = v24
            case "a":
                v4 = v4
                _ = len(v10)
            case "hello":
                v10 = append(v10, fi22)
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v25 int = 4294967295
    var v26 uint8 = 232
    fmt.Println(v4, v2)
    var v27 bool = true
    if len(v10) > 0 {
        v10[0] = 1
    }
    delete(v3, v2)
    v28 := GenS1{F0: "hello", F1: v26}
    var v29 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
