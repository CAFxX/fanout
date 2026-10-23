package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 string; F1 int16; F2 string }
func genFn0() uint16 {
    var v7 int32 = 50
    switch 332 {
    case -57:
        switch 127 {
        case 0:
            v7 = -772
        case 9223372036854775807:
            v8 := GenS1{F0: ("hello" + ("" + "")), F1: 842, F2: -276}
            _ = v8
        case 127:
            v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: 255, F2: 460}
            _ = v9
        }
    case 4294967295:
        var v10 int = int(v7)
        _ = v10
    }
    v7 = 1
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := []int{0, 0, (fi11 | (fi11 & (^(fi11))))}
        v13 := GenS2{F0: fi11, F1: (v7 >= v7)}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v7
    return 65535
}

func main() {
    var v1 uint = 28
    var v2 uint64 = 42
    var v3 int16 = 497
    var v4 uint16 = 1
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    v14 := GenS1{F0: "a", F1: v4, F2: 42}
    var v15 uint = 301
    v4 -= (v4 >> ((4294967295) & 63))
    v4 = (v4 >> ((v15) & 63))
    v16 := []int{851, 42}
    fmt.Println(v3, v4, v4)
    v17 := GenS3{F0: "foo bar", F1: v3, F2: ("foo bar" + "hello")}
    var v18 string = ("" + ("" + (("a" + (("" + "a") + "")) + "a")))
    if (!(2 < 255)) {
        fmt.Println(v3, v4)
    }
    fmt.Println(v3)
    v19 := GenS2{F0: 2147483647, F1: false}
    if len(v16) > 0 {
        v16[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
