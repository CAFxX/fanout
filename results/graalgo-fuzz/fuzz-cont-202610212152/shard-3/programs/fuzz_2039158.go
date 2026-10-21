package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint, p1 uint64, p2 int8) int32 {
    var v5 string = "\n\t"
    var v6 uint64 = uint64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := []string{("a" + "foo bar"), ((("a" + "foo bar") + "a") + "foo bar"), "hello", ((("\n\t" + ("\n\t" + "hello")) + ("" + "foo bar")) + ("hello" + ("a" + "foo bar")))}
    var v2 int64 = 627
    var v3 uint = 42
    var v4 float32 = 3.14159
    for fi7 := 0; fi7 < 7; fi7++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v8 bool = (-9223372036854775808 == -9223372036854775808)
        _ = fi7
        _ = v8
    }
    var v9 float32 = float32(((0.0 - 1.5) - (3.14159 * 3.14159)))
    switch 42 {
    case -2147483648:
        var v10 int8 = -79
        _ = v10
    default:
        var v11 float64 = (1.5 / float64((v9 + v4)))
        _ = v11
    }
    v3 -= v3
    var v12 uint16 = 449
    v13 := GenS1{F0: v2}
    var v14 float32 = v4
    fmt.Println(v14)
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: -9223372036854775808}
            _ = fi16
            _ = v17
        }
        switch (fi15 & (-(fi15))) {
        case 210:
            v12 = 32767
        case 127:
            var v18 uint32 = 65535
            _ = v18
        }
        _ = fi15
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v19 uint8 = 127
    fmt.Println(v4, v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
