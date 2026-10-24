package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v2 = v2
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = (v5 == v5)
        v5 = (v5 | (v5 >> ((491) & 63)))
        v8 := v1[-2147483648]
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float64 = 3.14159
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: "foo bar"}
        _ = fi9
        _ = v12
    }
    var v13 float32 = 3.14159
    if false {
        if ((!true) || (!true)) {
            if (!false) {
                v5 -= v5
            }
            v2 -= v2
        } else {
            var v14 uint = 877
            fmt.Println(v5)
            _ = v14
        }
    }
    v5 = (v5 / ((2147483647) | 1))
    var v15 uint32 = 670
    v16 := GenS1{F0: (("foo bar" + "") + ("\n\t" + "foo bar"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
