package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = (("foo bar" + "\n\t") + ("a" + "foo bar"))
    var v2 float64 = (float64(-2.25) - (-2.25 / 1.5))
    var v3 float64 = (1.5 + -2.25)
    var v4 bool = (!(!(128 >= 128)))
    v5 := []uint32{2147483647, 471, 4294967295, 1}
    if false {
        var v6 int16 = 0
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int = (-(fi7))
            v8 = fi7
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 float64 = 0.0
            var v11 bool = v4
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 762}
        _ = v12
    }
    v1 = v1
    if ((("" + v1) + ((v1 + (v1 + (v1 + v1))) + ("foo bar" + "a"))) < v1) {
        var v13 uint = 65535
        fmt.Println(v13)
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v5 = append(v5, 531)
        _ = fi14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v15 := GenS2{F0: 2147483647, F1: 461}
    v16 := GenS3{F0: 9223372036854775807}
    var v17 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
