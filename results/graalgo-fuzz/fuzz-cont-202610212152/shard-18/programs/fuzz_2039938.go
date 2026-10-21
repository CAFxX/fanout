package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int8; F2 int }

func main() {
    v1 := []int{2147483647, 422, 9223372036854775807}
    var v2 bool = true
    var v3 int16 = 32767
    v4 := make(map[uint32]int)
    var v5 int8 = -57
    var v6 bool = ((0 > -1) && (42 >= 876))
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v3)
        if (!("foo bar" != ("\n\t" + ""))) {
            v2 = (v6 || v6)
        } else {
            if v2 {
                v3 = 1
                v8 := GenS1{F0: "\n\t"}
                _ = v8
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi7
    }
    v9 := GenS1{F0: "foo bar"}
    fmt.Println(v3, v3, v5)
    if true {
        v5 = (v5 + 127)
        if v6 {
            v2 = (4294967295 > 4294967295)
            var v10 int32 = -1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int32 = 117
        v4[2] = -1
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: -2147483648}
    switch 127 {
    case 1:
        v14 := GenS3{F0: 42, F1: v5, F2: 128}
        _ = v14
    }
    switch -279 {
    case -1:
        for fi15 := 0; fi15 < 6; fi15++ {
            _ = len(v1)
            _ = fi15
        }
    }
    v16 := GenS1{F0: "\n\t"}
    var v17 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
