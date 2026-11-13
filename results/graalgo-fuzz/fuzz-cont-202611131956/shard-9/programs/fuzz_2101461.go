package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 int) string {
    var v4 bool = (!true)
    p0 += (p0 >> ((65535) & 63))
    v5 := GenS1{F0: 483}
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ("a" + "\n\t")
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 int64 = -678
    if true {
        for fi7 := 0; fi7 < 4; fi7++ {
            v3 += (v3 << ((1) & 63))
            var v8 uint8 = 74
            var v9 float64 = ((0.0 / ((((3.14159 * float64(1e10)) / 0.1) - 3.14159) / 3.14159)) - (3.14159 + (-2.25 * float64(1.5))))
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v3)
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int32 = 2147483647
            v12 := GenS2{F0: v11, F1: (v3 << ((1) & 63)), F2: ("a" + (("foo bar" + "") + ((("foo bar" + "") + "\n\t") + (("" + ("hello" + "foo bar")) + "\n\t"))))}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 float32 = (0.1 + 1e10)
        _ = v13
    }
    v14 := []int{157, 344, -1, 305}
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 bool = ((42 <= 127) || (!false))
        v17 := GenS2{F0: -214, F1: -1, F2: ("foo bar" + (((("foo bar" + ("a" + ("foo bar" + (("hello" + "a") + "\n\t")))) + "") + "hello") + ("" + "")))}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch -78 {
    case 228:
        delete(v2, 2147483647)
    }
    var v18 bool = (!true)
    v19 := GenS2{F0: 0, F1: (^(v3)), F2: "foo bar"}
    v3 = (^(v3))
    var v20 uint = 18446744073709551615
    fmt.Println(v20)
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int = (fi21 ^ fi21)
        _ = fi21
        _ = v22
    }
    v20 = v20
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint = (^(v20))
        for fi25 := 0; fi25 < 7; fi25++ {
            v18 = (255 <= 2)
            v18 = (2 != 2)
            v26 := GenS1{F0: 42}
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS1{F0: 470}
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
