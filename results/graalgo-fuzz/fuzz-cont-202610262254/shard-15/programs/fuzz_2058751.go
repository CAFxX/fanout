package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 uint8, p2 uint8) uint16 {
    if true {
        p2 = p1
    }
    v7 := GenS1{F0: 127, F1: 948}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 int16 = (^(v1))
    v3 := []int{0, 214, -171, -172}
    v4 := make(map[uint32]int)
    v5 := []string{"", "", ""}
    var v6 int32 = 1
    if true {
        var v8 uint8 = 42
        fmt.Println(v8, v6, v1)
        var v9 int8 = 0
        _ = v8
        _ = v9
    } else {
        v4[42] = -900
        v6 = (^(v6))
        var v10 int32 = v6
        _ = v10
    }
    fmt.Println(v2, v2, v6)
    var v11 int8 = 42
    var v12 string = "a"
    if (!true) {
        v5 = append(v5, (v12 + ((((((v12 + v12) + "") + v12) + "foo bar") + v12) + "a")))
        v13 := GenS1{F0: (-(v11)), F1: 161}
        var v14 string = ("a" + "")
        _ = v13
        _ = v14
    }
    v1 = v1
    var v15 uint16 = 32767
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS2{F0: 2}
        switch (v12 + v12) {
        case "hello":
            switch "foo bar" {
            case "hello":
                if true {
                    fmt.Println(v12, v12)
                    v18 := GenS3{F0: (fi16 ^ 128)}
                    v15 = v15
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 int64 = 0
                        var v21 uint = 42
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            }
        case "\n\t":
            switch v12 {
            case "":
                for fi22 := 0; fi22 < 6; fi22++ {
                    fmt.Println(fi16, v6, fi22)
                    _ = fi22
                }
            case "\n\t":
                var v23 uint16 = (v15 & 1)
                _ = v23
            case "a":
                _ = len(v4)
            }
        default:
            v24 := GenS3{F0: (fi16 * (fi16 | int(v1)))}
            _ = v24
        }
        v4[42] = 0
        _ = fi16
        _ = v17
    }
    v2 += (v1 << ((265) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
