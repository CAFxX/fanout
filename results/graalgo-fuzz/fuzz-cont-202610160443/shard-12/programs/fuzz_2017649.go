package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    v1 := []int{}
    var v2 uint32 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 243
    v5 := []string{"a", "a", "hello", "hello"}
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int16 = -211
        if ((v2 << ((255) & 63)) == v2) {
            var v9 int32 = 1
            _ = v9
        }
        switch "a" {
        case "foo bar":
            fmt.Println(fi7, v2, v4)
        case "a":
            fmt.Println(v4, fi7, v8)
        case "":
            v10 := GenS1{F0: (("a" + "foo bar") < "a")}
            _ = v10
        default:
            var v11 bool = ((3.14159 + 0.1) >= 3.14159)
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    switch 422 {
    case 1:
        v2 += 42
    case 4294967295:
        switch -2147483648 {
        case 9223372036854775807:
            v4 -= 255
        }
    default:
        v2 += 2
    }
    v12 := GenS1{F0: (0.1 > 3.14159)}
    v4 = v4
    fmt.Println(v4, v4)
    v13 := GenS1{F0: true}
    var v14 float64 = 0.1
    for fi15 := 0; fi15 < 2; fi15++ {
        switch 32768 {
        case 32767:
            v16 := v3[v2]
            _ = v16
        case 362:
            var v17 uint8 = v4
            _ = v17
        case 62:
            v3[905] = v2
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int16 = 846
            v19 = (v19 / ((v19) | 1))
            _ = fi18
            _ = v19
        }
        _ = fi15
    }
    v20 := GenS2{F0: 229}
    v14 -= v14
    var v21 uint8 = 29
    if (114 != 35) {
        v4 -= v21
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            v4 = (v21 & 127)
            for fi23 := 0; fi23 < 11; fi23++ {
                v2 -= (v2 ^ (v2 & v2))
                var v24 int64 = -1
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            if false {
                for fi26 := 0; fi26 < 11; fi26++ {
                    _ = len(v1)
                    var v27 int32 = 2147483647
                    _ = fi26
                    _ = v27
                }
                v28 := GenS3{F0: (fi25 - 128), F1: 695}
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                if true {
                    fmt.Println(v21, v4)
                    fmt.Println(v14)
                } else {
                    var v30 int16 = 42
                    var v31 bool = (!(!true))
                    var v32 bool = false
                    _ = v30
                    _ = v31
                    _ = v32
                }
                fmt.Println(v21)
                _ = fi29
            }
            _ = fi25
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
