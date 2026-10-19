package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 1
    v3 := make(map[string]int)
    if true {
        var v4 uint32 = uint32(v1)
        v4 -= v4
        _ = v4
    }
    switch ("" + "hello") {
    case "a":
        v1 = v1
    case "\n\t":
        var v5 uint = v1
        _ = v5
    case "":
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int32 = -1
            v7 = (-(v7))
            if (int8(v2) >= 42) {
                var v8 bool = (int64(v7) != 9223372036854775807)
                _ = v8
            } else {
                v9 := GenS1{F0: 32767}
                v1 += (v1 - (^(v1)))
                var v10 uint64 = 951
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    var v11 int64 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        if (fi12 != -9223372036854775808) {
            var v13 uint8 = 59
            _ = v13
        } else {
            v2 += v2
            var v14 int32 = 822
            fmt.Println(fi12, fi12)
            _ = v14
        }
        _ = fi12
    }
    v2 += 272
    if (false && true) {
        var v15 int64 = v11
        _ = v15
    }
    v2 += v2
    v16 := []int{42, 4294967295, 9223372036854775807}
    fmt.Println(v1)
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS1{F0: 32767}
        _ = fi17
        _ = v18
    }
    if true {
        var v19 int64 = v11
        var v20 int64 = v11
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(v19, v11)
            _ = fi21
        }
        _ = v19
        _ = v20
    } else {
        if (-128 != 42) {
            v22 := GenS1{F0: 222}
            _ = v22
        } else {
            var v23 int = 727
            var v24 uint8 = 0
            _ = v23
            _ = v24
        }
        var v25 int16 = 0
        _ = v25
    }
    v26 := GenS1{F0: 32767}
    v11 -= v11
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v26
    _ = v3
}
