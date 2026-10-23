package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = 0.0
    v2 := []int{1, -1, -908, 32768}
    v3 := make(map[int]string)
    var v4 uint32 = 4294967295
    v5 := []int{42, 128, -82}
    switch 0 {
    case 918:
        v1 -= (v1 * v1)
    case 127:
        v6 := GenS1{F0: (true || true)}
        _ = v6
    default:
        fmt.Println(v4)
    }
    if ((-128 >= 1) && (!false)) {
        v7 := GenS2{F0: false}
        v4 -= 305
        _ = v7
    }
    v8 := GenS2{F0: false}
    v9 := GenS2{F0: (-9223372036854775808 < 0)}
    var v10 int64 = 0
    if (!true) {
        v11 := GenS1{F0: (!(true || (2 < 18446744073709551615)))}
        _ = v11
    }
    var v12 int16 = -394
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v13 bool = (!true)
        _ = v13
    }
    v4 = (^(v4))
    v14 := GenS2{F0: false}
    var v15 float64 = 1e10
    if (128 > 128) {
        if false {
            fmt.Println(v1, v12)
        }
        v16 := GenS1{F0: (((0 == 42) || true) || false)}
        _ = v16
    }
    v3[-634] = ((("hello" + "") + ("hello" + "foo bar")) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
