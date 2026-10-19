package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64) uint16 {
    v6 := GenS1{F0: 42}
    var v7 int64 = p0
    _ = p0
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int, p1 string) uint64 {
    if false {
        switch p0 {
        case -29:
            var v8 uint16 = 65535
            _ = v8
        case 4294967295:
            v9 := []int{419, (^(p0))}
            _ = v9
        default:
            var v10 string = (("hello" + (("foo bar" + p1) + "a")) + p1)
            _ = v10
        }
        p1 = "foo bar"
        v11 := GenS1{F0: 65535}
        _ = v11
    } else {
        fmt.Println(p1)
    }
    fmt.Println(p0, p0)
    for fi12 := 0; fi12 < 4; fi12++ {
        p0 -= 9223372036854775807
        v13 := []int{(p0 + p0), (^(p0)), (fi12 - -9223372036854775808)}
        var v14 int8 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = (p0 * (p0 << ((2) & 63)))
    switch p1 {
    case "\n\t":
        var v16 string = "a"
        _ = v16
    case "":
        for fi17 := 0; fi17 < 5; fi17++ {
            v15 += -9223372036854775808
            _ = fi17
        }
    default:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v15
    return 4294967295
}

func main() {
    var v1 int32 = -142
    v2 := []uint32{42}
    var v3 uint64 = 493
    var v4 int16 = 880
    v5 := make(map[int]string)
    fmt.Println(v1, v4)
    delete(v5, 982)
    var v18 int64 = 944
    var v19 int16 = (v4 << ((255) & 63))
    if false {
        _ = len(v2)
        v19 = v4
    } else {
        fmt.Println(v19, v19)
        var v20 uint32 = 4294967295
        v21 := GenS1{F0: 42}
        _ = v20
        _ = v21
    }
    if (1 < 127) {
        var v22 int32 = v1
        _ = v22
    }
    v2 = append(v2, 42)
    var v23 int8 = -128
    for fi24 := 0; fi24 < 8; fi24++ {
        if (1 == 2) {
            fmt.Println(v23)
        }
        _ = fi24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 string = "\n\t"
        var v27 int16 = (^(v4))
        v19 = (v4 - 1)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
