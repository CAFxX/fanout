package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int, p1 int32) int {
    v4 := []int{(p0 & p0), p0}
    var v5 uint = 0
    var v6 bool = (!(int16(p1) > -1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 32768
}

func main() {
    var v1 int32 = -214
    var v2 uint64 = 1
    var v3 uint16 = 176
    if true {
        var v7 uint8 = 128
        v7 += (v7 % ((107) | 1))
        _ = v7
    }
    v1 -= 720
    if true {
        v8 := []int{1, -186, 0, genFn0(32768, (v1 | (v1 | v1)))}
        _ = v8
    } else {
        v3 -= 65535
    }
    var v9 float32 = (float32(3.14159) * float32(-2.25))
    if true {
        v10 := GenS1{F0: v3, F1: (v1 & (v1 / ((0) | 1))), F2: "hello"}
        v3 = v3
        if false {
            var v11 bool = false
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: v3, F1: v1, F2: ""}
                var v14 uint16 = (v3 ^ 1)
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: 916, F1: v1, F2: ""}
                var v17 string = ""
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v11
        } else {
            v1 = 0
            v18 := GenS1{F0: v3, F1: v1, F2: (("hello" + "") + "")}
            _ = v18
        }
        _ = v10
    } else {
        v1 += v1
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := []int{}
            if (!(v2 > v2)) {
                for fi21 := 0; fi21 < 16; fi21++ {
                    v3 += (v3 / ((v3) | 1))
                    switch "hello" {
                    case "\n\t":
                        v1 -= v1
                    case "foo bar":
                        v1 = (v1 - (v1 % ((2147483647) | 1)))
                    }
                    v22 := GenS1{F0: 65535, F1: (v1 >> ((674) & 63)), F2: "\n\t"}
                    _ = fi21
                    _ = v22
                }
            } else {
                v20 = append(v20, 4294967295)
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                v9 = (v9 * (0.1 / (v9 / 0.1)))
                _ = fi23
            }
            _ = fi19
            _ = v20
        }
        v2 += v2
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int16 = 0
        _ = fi24
        _ = v25
    }
    v26 := []int{}
    v27 := GenS1{F0: v3, F1: (v1 >> ((2) & 63)), F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
