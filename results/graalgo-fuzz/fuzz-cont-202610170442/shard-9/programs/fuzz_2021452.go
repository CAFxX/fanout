package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 898
    v3 := make(map[int]bool)
    var v4 int = -9223372036854775808
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 3; fi6++ {
        switch v1 {
        case "foo bar":
            v7 := []int{-190}
            _ = v7
        }
        _ = fi6
    }
    switch (^(v4)) {
    case 1:
        var v8 uint32 = 319
        _ = v8
    case 797:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint8 = 255
            var v11 int8 = -1
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if true {
            v12 := []int{}
            var v13 string = (v1 + "hello")
            var v14 float32 = 1e10
            _ = v12
            _ = v13
            _ = v14
        }
    }
    var v15 uint32 = 42
    v5 = v5
    fmt.Println(v15)
    v5 -= (v5 >> ((0) & 63))
    if ((v2 / ((65535) | 1)) >= (v2 - (v2 - v2))) {
        v5 = v5
        v15 += v15
        v3[-950] = false
    }
    v15 += 78
    switch ("\n\t" + (("" + "") + "\n\t")) {
    case "a":
        var v16 float64 = (0.0 / 0.1)
        _ = v16
    case "":
        switch v4 {
        case -996:
            v1 = (v1 + v1)
        default:
            switch 1 {
            case 4294967295:
                var v17 int8 = 40
                _ = v17
            case 0:
                fmt.Println(v2, v15)
            }
        }
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: (v1 + v1)}
            if (1 <= -9223372036854775808) {
                v20 := GenS1{F0: v1}
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v1, v2)
    v5 = 2
    fmt.Println(v2, v2, v15)
    var v21 int64 = 42
    var v22 uint = 18446744073709551615
    v23 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
