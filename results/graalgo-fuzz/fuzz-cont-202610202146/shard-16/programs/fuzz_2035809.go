package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int16 = -32768
    var v4 float32 = (1.5 / (0.1 - (0.1 * (0.0 / -2.25))))
    v5 := []string{}
    var v6 uint = 0
    fmt.Println(v3, v6, v3)
    v1[32767] = (("" + "\n\t") + "\n\t")
    for fi7 := 0; fi7 < 13; fi7++ {
        v6 += v6
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = fi7
    }
    v8 := GenS1{F0: false, F1: (1 <= 100), F2: 60}
    fmt.Println(v6, v6, v3)
    var v9 bool = true
    switch 1 {
    case -1:
        v5 = append(v5, ("" + "\n\t"))
    default:
        v10 := GenS1{F0: true, F1: false, F2: 153}
        _ = v10
    }
    var v11 uint64 = 4294967295
    var v12 int8 = 1
    for fi13 := 0; fi13 < 7; fi13++ {
        v6 = 44
        v14 := GenS1{F0: false, F1: (42 == 290), F2: 1}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v9, fi16, v9)
            _ = fi16
        }
        v17 := v1[(fi15 % ((fi15) | 1))]
        v18 := GenS2{F0: v17}
        _ = fi15
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: (v12 & v12)}
    for fi20 := 0; fi20 < 4; fi20++ {
        _ = len(v2)
        _ = fi20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v6 = v6
        if (2 == 156) {
            v22 := GenS3{F0: -1}
            for fi23 := 0; fi23 < 19; fi23++ {
                switch (((("a" + "foo bar") + "hello") + "hello") + (("" + "") + "\n\t")) {
                case "a":
                    var v24 bool = v9
                    _ = v24
                case "foo bar":
                    v25 := GenS3{F0: v12}
                    _ = v25
                }
                var v26 uint16 = 1
                _ = fi23
                _ = v26
            }
            _ = v22
        } else {
            v1[fi21] = "hello"
        }
        _ = fi21
    }
    switch 4294967295 {
    case -1:
        var v27 uint32 = 2147483647
        _ = v27
    case 367:
        var v28 bool = (("" + "a") >= "")
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        v4 = v4
        _ = fi29
    }
    delete(v2, 311)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
