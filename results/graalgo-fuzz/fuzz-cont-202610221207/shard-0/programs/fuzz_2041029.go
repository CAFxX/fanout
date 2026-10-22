package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint32, p1 bool) uint32 {
    v4 := []int{42}
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 string = "hello"
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v6)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (p0 / (((p0 >> ((720) & 63))) | 1))}
    v8 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    _ = v8
    return (p0 | p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 127
    var v3 uint64 = 18446744073709551615
    v2 -= (v2 + (-(v2)))
    var v9 uint32 = 257
    var v10 int64 = int64(v3)
    v10 = (v10 >> ((2) & 63))
    switch 32767 {
    case 127:
        var v11 uint = 4294967295
        _ = v11
    default:
        v12 := GenS1{F0: (v9 ^ v9)}
        _ = v12
    }
    var v13 uint8 = 169
    v14 := GenS1{F0: (^(v9))}
    for fi15 := 0; fi15 < 1; fi15++ {
        delete(v1, (fi15 / ((int(v13)) | 1)))
        _ = fi15
    }
    v3 = (v3 * 643)
    fmt.Println(v3)
    v9 = (^(v9))
    for fi16 := 0; fi16 < 2; fi16++ {
        switch fi16 {
        case 32768:
            var v17 float64 = 0.1
            _ = v17
        case -23:
            v10 = v10
        case -2147483648:
            v2 = v2
        }
        var v18 bool = false
        _ = fi16
        _ = v18
    }
    if (v13 <= v13) {
        fmt.Println(v13, v10)
        var v19 string = ""
        delete(v1, -337)
        _ = v19
    }
    switch (("hello" + ("hello" + (("hello" + "") + ""))) + "foo bar") {
    case "":
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(v13, fi20, v9)
            v2 = v2
            _ = fi20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
