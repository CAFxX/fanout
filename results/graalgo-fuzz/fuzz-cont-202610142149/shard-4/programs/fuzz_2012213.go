package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 0
    var v3 uint64 = 2
    var v4 int64 = 0
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = -2147483648
        var v6 uint16 = v2
        if ("foo bar" > "\n\t") {
            fmt.Println(v4, v4, v4)
            var v7 int32 = v1
            var v8 float32 = ((-2.25 * -2.25) + float32(1e10))
            _ = v7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v1 = v1
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int16 = -825
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch 1 {
    case 128:
        var v12 float32 = 0.0
        _ = v12
    case 9223372036854775807:
        fmt.Println(v2)
    }
    v13 := []int{2147483647, 128, 266, 9223372036854775807}
    _ = len(v13)
    v1 = (-(v1))
    if ("\n\t" == "") {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    var v14 uint16 = 2
    var v15 int64 = 1
    var v16 int8 = 94
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 42}
            _ = fi18
            _ = v19
        }
        if (-1 <= v16) {
            if false {
                var v20 uint16 = 255
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS1{F0: 617}
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                var v23 bool = (42 > v16)
                v16 = 100
                var v24 uint32 = 97
                _ = v23
                _ = v24
            }
            var v25 bool = ((v3 / ((v3) | 1)) < v3)
            fmt.Println(v15, fi17)
            _ = v25
        }
        _ = fi17
    }
    var v26 int8 = (v16 * v16)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
