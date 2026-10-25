package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 uint }
func genFn0() int {
    var v6 int16 = 1
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = 9223372036854775807
        v6 = 32767
        var v9 int8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{}
    _ = v10
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("foo bar" + "foo bar"), "hello", ("foo bar" + "a")}
    v3 := make(map[string]uint32)
    var v4 uint64 = 606
    v5 := []uint32{42, 1}
    v11 := GenS1{F0: 1, F1: -2147483648}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v12 uint = 18446744073709551615
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            var v14 int64 = 0
            fmt.Println(fi13, v4)
            _ = v14
        }
        v12 += v12
        _ = fi13
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if (((fi15 - 2147483647) < fi15) && ((v12 << ((2) & 63)) != (v12 / ((539) | 1)))) {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := v3[""]
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: 1, F1: 311}
            v2 = append(v2, "")
            _ = v18
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS2{F0: 796, F1: 32767}
                _ = fi19
                _ = v20
            }
        }
        _ = fi15
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "hello")
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        v4 += v4
        v4 -= v4
        _ = fi22
    }
    if true {
        v23 := GenS1{F0: -1, F1: 646}
        _ = v23
    }
    if false {
        delete(v3, "")
        var v24 uint = (v12 * 255)
        _ = v24
    } else {
        var v25 float64 = (1.5 - 1.5)
        v26 := GenS1{F0: 1000, F1: -2147483648}
        var v27 int8 = 42
        _ = v25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS3{F0: (fi28 - genFn0()), F1: (v12 & 2)}
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: -2147483648}
    delete(v1, 127)
    v3["\n\t"] = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
