package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 uint8 = 127
    var v5 int32 = -1
    var v6 uint64 = 18446744073709551615
    switch "hello" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    if true {
        if true {
            fmt.Println(v5, v6, v6)
            v1 = v1
        } else {
            var v7 int32 = (v5 * (v5 ^ v5))
            _ = v7
        }
        if len(v3) > 0 {
            v3[0] = 0
        }
        if (!false) {
            if (true || (42 > 32767)) {
                if (!(1e10 <= (v1 / (v1 - (1e10 / (0.1 - v1)))))) {
                    v6 += v6
                    v6 = (v6 + (v6 | (v6 * uint64(v4))))
                }
            }
            fmt.Println(v6, v1, v4)
            var v8 int32 = 1
            _ = v8
        } else {
            v4 += v4
        }
    } else {
        v2[9223372036854775807] = 1
        v9 := GenS1{F0: (v6 >> ((947) & 63))}
        var v10 uint = 0
        _ = v10
        _ = v9
    }
    var v11 string = (("hello" + "foo bar") + "foo bar")
    var v12 int64 = 1
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int = 0
        switch v11 {
        case "":
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    v16 := v2[(v14 * v14)]
                    switch "\n\t" {
                    case "hello":
                        v17 := GenS1{F0: (^(v6))}
                        _ = v17
                    }
                    if (true || (((3.14159 - 0.0) / 1.0) > 1e10)) {
                        var v18 int64 = v12
                        var v19 string = (v11 + v11)
                        _ = v18
                        _ = v19
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (fi15 | (fi13 * (^(fi13))))
                        }
                        v12 += (v12 ^ (v12 ^ (^(v12))))
                        switch v11 {
                        case "foo bar":
                            fmt.Println(v14)
                        default:
                            var v20 string = v11
                            _ = v20
                        }
                        _ = len(v3)
                        var v21 int = 32768
                        _ = v21
                    }
                    v3 = append(v3, 128)
                    v1 += (v1 + 0.1)
                    _ = v16
                }
                v4 = (v4 - (v4 * v4))
                _ = fi15
            }
        case "a":
            switch (v11 + v11) {
            case "hello":
                for fi22 := 0; fi22 < 13; fi22++ {
                    v5 = v5
                    _ = fi22
                }
            case "\n\t":
                if (389 <= 519) {
                    var v23 uint = 65535
                    switch (v14 << ((v23) & 63)) {
                    case 32768:
                        v24 := GenS1{F0: (v6 & (v6 + (v6 & v6)))}
                        _ = v24
                    case 9223372036854775807:
                        if (v1 > ((3.14159 - float32(0.1)) + v1)) {
                            v11 = "foo bar"
                            v23 += (^(v23))
                        } else {
                            v14 -= (v14 + fi13)
                            v25 := GenS2{F0: 0}
                            fmt.Println(v6, v23)
                            _ = v25
                        }
                    }
                    _ = v23
                } else {
                    fmt.Println(v6)
                    switch fi13 {
                    case -43:
                        var v26 int64 = (^(v12))
                        _ = v26
                    }
                    if (!(1e10 < ((-2.25 - 1e10) / (-2.25 - 0.0)))) {
                        var v27 int = -1
                        _ = v27
                    }
                }
            case "foo bar":
                fmt.Println(v4, v1, v6)
            }
        case "hello":
            v4 = (v4 - 2)
        }
        _ = fi13
        _ = v14
    }
    v11 = ("" + v11)
    if len(v3) > 0 {
        v3[0] = -606
    }
    var v28 int64 = -420
    v28 = 0
    fmt.Println(v11, v5, v1)
    switch 32767 {
    case 1:
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int64 = (v28 & v28)
            _ = fi29
            _ = v30
        }
    case 32767:
        v1 += v1
    default:
        v2[244] = 2
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 += (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
