package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
func genFn0(p0 bool, p1 uint64, p2 uint) int64 {
    v6 := GenS1{F0: p0}
    v7 := []int{0, -962, 81, 0}
    if (((!(-561 == -1)) || (32767 >= 1)) || p0) {
        p2 -= (^(p2))
        var v8 int = int(p2)
        _ = v8
    } else {
        var v9 uint8 = 2
        var v10 string = "foo bar"
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint16 = 0
            if false {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -8
}

func main() {
    var v1 uint = 42
    v2 := []string{""}
    var v3 bool = true
    var v4 float64 = 1.5
    var v5 int32 = 236
    v13 := GenS1{F0: true}
    var v14 uint64 = 4294967295
    v5 = 42
    _ = len(v2)
    v1 = v1
    fmt.Println(v3)
    v14 = v14
    for fi15 := 0; fi15 < 19; fi15++ {
        if false {
            if true {
                fmt.Println(v5, v4)
                v2 = append(v2, "foo bar")
            } else {
                switch (fi15 & (fi15 - fi15)) {
                case 1:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 int64 = -9223372036854775808
                        var v18 uint = 0
                        if v3 {
                            v19 := GenS2{F0: 9223372036854775807, F1: 127, F2: 60}
                            _ = v19
                        }
                        var v20 uint64 = v14
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v20
                    }
                case 9223372036854775807:
                    v21 := GenS1{F0: v3}
                    _ = v21
                }
                var v22 int8 = 127
                _ = v22
            }
            var v23 int64 = -9223372036854775808
            v2 = append(v2, "a")
            _ = v23
        } else {
            var v24 float64 = v4
            var v25 float32 = 0.1
            var v26 int64 = 1
            _ = v24
            _ = v25
            _ = v26
        }
        v3 = v3
        v5 = (v5 & (v5 + 0))
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
