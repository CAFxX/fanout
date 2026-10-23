package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 int16 = 77
    v4 := []uint32{}
    if true {
        var v5 int64 = -1
        var v6 uint = 95
        _ = v5
        _ = v6
    }
    var v7 float64 = 0.0
    var v8 int8 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        switch (fi9 | -1) {
        case -194:
            v1 = append(v1, "a")
        case -699:
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
        _ = fi9
    }
    var v10 float32 = -2.25
    v11 := v2[118]
    var v12 uint64 = 564
    if (!true) {
        if false {
            var v13 bool = (!(1 <= 42))
            _ = v13
        }
        v8 = v8
    }
    v7 += v7
    v10 = 1e10
    switch ((("a" + v11) + (v11 + v11)) + v11) {
    case "":
        fmt.Println(v12, v11, v12)
    case "a":
        var v14 int8 = v8
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 6; fi15++ {
            v11 = ""
            v16 := GenS1{F0: uint16(v12)}
            delete(v2, fi15)
            _ = fi15
            _ = v16
        }
    }
    if false {
        v12 = (v12 * v12)
    }
    if true {
        v10 += 1.5
        v11 = v11
    } else {
        v3 = (v3 << ((4294967295) & 63))
    }
    var v17 bool = (int32(v3) < 1)
    v18 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
