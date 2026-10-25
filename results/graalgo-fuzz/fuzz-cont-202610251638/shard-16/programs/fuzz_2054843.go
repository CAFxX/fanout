package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
type GenS4 struct { F0 uint64 }
func genFn0() string {
    var v7 int8 = 42
    if (((9223372036854775807 == -2147483648) && (!true)) && false) {
        v8 := GenS1{F0: 1}
        v9 := []int{32767, 0, -401, 4294967295}
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    if (65535 <= 65535) {
                        v12 := GenS1{F0: 1}
                        v9 = append(v9, (-(fi11)))
                        _ = v12
                    } else {
                        var v13 int8 = 127
                        _ = v13
                    }
                    _ = len(v9)
                    _ = fi11
                }
                _ = fi10
            }
            if (42 == 0) {
                var v14 bool = (-1 == 42)
                v15 := GenS2{F0: 1}
                v14 = (v14 || (!(!true)))
                _ = v14
                _ = v15
            }
            fmt.Println(v7)
        }
        _ = v8
        _ = v9
    }
    v7 = 1
    v7 -= v7
    v16 := GenS1{F0: 42}
    _ = v16
    _ = v7
    return "a"
}
func genFn1() int8 {
    v17 := []int{32767, 128, -658}
    for fi18 := 0; fi18 < 14; fi18++ {
        _ = len(v17)
        var v19 int8 = 0
        _ = fi18
        _ = v19
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = v17
    return 127
}

func main() {
    var v1 int64 = 353
    v2 := make(map[int]int)
    var v3 float32 = (1e10 - 1e10)
    var v4 uint16 = 592
    v5 := []int{}
    var v6 uint16 = v4
    var v20 int32 = 1
    if (!true) {
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS3{F0: -622, F1: "a", F2: 42}
            v3 = v3
            v23 := GenS3{F0: -805, F1: ("\n\t" + "\n\t"), F2: 0}
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS4{F0: 2}
        _ = v24
    }
    v4 = (v4 / ((v6) | 1))
    fmt.Println(v20, v6)
    v6 = 65535
    var v25 uint8 = 99
    _ = len(v2)
    v4 = (v4 / ((374) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
