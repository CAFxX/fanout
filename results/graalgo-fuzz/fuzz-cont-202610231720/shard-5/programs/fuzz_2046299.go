package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 127
    var v3 int32 = 2147483647
    v3 -= (v3 * 0)
    v3 = v3
    for fi4 := 0; fi4 < 20; fi4++ {
        switch "foo bar" {
        case "":
            var v5 uint64 = 4294967295
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: 1, F1: ((("foo bar" + "\n\t") + "\n\t") + "foo bar"), F2: 42}
            _ = v6
        case "hello":
            var v7 uint64 = 590
            _ = v7
        default:
            v3 += 1
        }
        v8 := v1[-2147483648]
        v9 := GenS1{F0: 255, F1: "foo bar", F2: 42}
        _ = fi4
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 2}
    for fi11 := 0; fi11 < 2; fi11++ {
        v2 += int8(fi11)
        _ = fi11
    }
    var v12 bool = (!((!(!(-264 != 9223372036854775807))) || (!true)))
    v13 := GenS1{F0: 65535, F1: "a", F2: 337}
    var v14 int64 = 0
    var v15 int16 = -32768
    fmt.Println(v3, v15)
    if v12 {
        v14 -= v14
        if v12 {
            fmt.Println(v15, v15, v15)
            var v16 uint64 = 42
            delete(v1, -9223372036854775808)
            _ = v16
        } else {
            v12 = (!false)
        }
    } else {
        var v17 uint16 = 1
        _ = v17
    }
    if v12 {
        var v18 uint32 = 2147483648
        _ = v18
    } else {
        var v19 bool = v12
        _ = v19
    }
    if (v12 || (!(2 == 494))) {
        v14 += v14
        var v20 float64 = 0.1
        if v12 {
            v12 = (!v12)
            v21 := GenS2{F0: 33}
            _ = v21
        }
        _ = v20
    } else {
        v14 = (^(v14))
        v2 = v2
        var v22 uint = 255
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS1{F0: 513, F1: "", F2: v15}
        v25 := GenS3{F0: 0, F1: v2, F2: 0}
        var v26 bool = ((1.5 * 1.5) >= (3.14159 + 3.14159))
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    delete(v1, -398)
    var v27 float64 = (1.5 * 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
}
