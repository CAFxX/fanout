package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    v3 := make(map[int]uint32)
    v4 := []int{-886, -325}
    if ((false || (false || false)) || (547 >= 1)) {
        var v5 bool = (false && (4294967295 > 0))
        _ = v5
    } else {
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint32 = 2147483648
                if (1 > 127) {
                    var v9 int = fi7
                    var v10 int = fi7
                    v11 := GenS1{F0: -7}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                if false {
                    var v12 string = "hello"
                    _ = v12
                }
                _ = fi7
                _ = v8
            }
            v2 = 2147483647
            _ = fi6
        }
        var v13 float32 = 1.5
        v4 = append(v4, 32768)
        _ = v13
    }
    var v14 int8 = -1
    v15 := GenS2{F0: 449, F1: true}
    var v16 uint = 18446744073709551615
    var v17 uint16 = 0
    if len(v4) > 0 {
        v4[0] = int(v2)
    }
    if true {
        if (false && true) {
            switch "hello" {
            case "\n\t":
                v17 += 65535
            case "":
                if false {
                    var v18 int16 = 32767
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                if (!(!false)) {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(fi20, v14)
                        default:
                            v1["hello"] = (fi19 ^ (fi20 >> ((v16) & 63)))
                        }
                        _ = fi20
                    }
                }
                v14 = v14
                _ = fi19
            }
        }
        delete(v3, -1)
    }
    if (127 <= v14) {
        var v21 uint64 = 0
        v17 -= (v17 & 43)
        for fi22 := 0; fi22 < 2; fi22++ {
            v21 += 18446744073709551615
            _ = fi22
        }
        _ = v21
    } else {
        if ("" == "a") {
            fmt.Println(v17, v16, v16)
        }
        v23 := GenS1{F0: 42}
        _ = v23
    }
    v16 += v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
