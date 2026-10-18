package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    var v3 uint64 = 98
    var v4 int64 = 42
    var v5 uint16 = 990
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v1 = ((v1 + v1) + v1)
            _ = fi8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if ((v3 == (v3 ^ (v3 / (((v3 | 42)) | 1)))) && true) {
            v10 := []int{}
            var v11 uint16 = (v5 - 42)
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: (v3 % (((^(v3))) | 1))}
        v13 := GenS2{F0: (^(v3))}
        _ = fi9
        _ = v12
        _ = v13
    }
    v5 += (v5 >> ((811) & 63))
    fmt.Println(v4, v6)
    if false {
        v4 = 1
    } else {
        var v14 int32 = 42
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int = fi15
        fmt.Println(v6, v5, fi15)
        var v17 bool = ((1e10 + -2.25) != 0.1)
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint = 0
    var v19 int64 = 9223372036854775807
    v20 := []int{}
    v5 = v5
    v21 := GenS3{F0: (^(v5))}
    if true {
        fmt.Println(v18, v19)
    }
    switch 2147483647 {
    case 42:
        var v22 int32 = int32(v3)
        _ = v22
    }
    var v23 string = v1
    if true {
        if true {
            var v24 bool = true
            if (!true) {
                switch ((("foo bar" + v23) + (v1 + (v23 + "\n\t"))) + "\n\t") {
                case "":
                    var v25 float32 = -2.25
                    _ = v25
                default:
                    v24 = (0 == v3)
                }
            }
            if v24 {
                var v26 uint64 = v3
                v5 = 32767
                _ = v26
            }
            _ = v24
        } else {
            for fi27 := 0; fi27 < 6; fi27++ {
                if (120 == 127) {
                    var v28 bool = false
                    _ = v28
                }
                _ = fi27
            }
            v20 = append(v20, int(v5))
            v2["a"] = (3.14159 > 0.0)
        }
        delete(v2, ("" + v23))
        _ = len(v2)
    } else {
        v3 = v3
    }
    var v29 bool = false
    if len(v20) > 0 {
        v20[0] = -584
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
