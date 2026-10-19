package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 uint8) int16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            for fi7 := 0; fi7 < 12; fi7++ {
                if (!true) {
                    var v8 float32 = 0.0
                    _ = v8
                } else {
                    var v9 float64 = (float64(0.1) / 1.0)
                    _ = v9
                }
                _ = fi7
            }
            p1 += (p1 ^ 42)
            v10 := GenS1{F0: 143, F1: 32767}
            _ = v10
        }
        v11 := GenS2{F0: 305, F1: 13, F2: (fi6 * fi6)}
        v12 := GenS3{F0: 42}
        _ = fi6
        _ = v11
        _ = v12
    }
    p1 += (p1 ^ (p1 - p1))
    switch "" {
    case "\n\t":
        p1 -= p1
    case "foo bar":
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint8 = (^(p1))
            _ = fi13
            _ = v14
        }
    default:
        var v15 int32 = -1
        _ = v15
    }
    var v16 uint64 = 1
    _ = p0
    _ = p1
    _ = v16
    return -950
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 int8 = 42
    var v4 uint64 = 18446744073709551615
    var v5 int16 = 35
    var v17 uint16 = 42
    v5 = v5
    if false {
        delete(v2, "")
        var v18 int8 = v3
        _ = v18
    } else {
        var v19 uint8 = 177
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := v2["hello"]
            for fi22 := 0; fi22 < 1; fi22++ {
                if v1 {
                    if true {
                        var v23 uint16 = (v17 >> ((4294967295) & 63))
                        fmt.Println(v23, fi22, v23)
                        switch (^(fi22)) {
                        case 2147483647:
                            v24 := []int{}
                            _ = v24
                        default:
                            v25 := []int{}
                            _ = v25
                        }
                        _ = v23
                    } else {
                        fmt.Println(v3)
                        v3 = v3
                    }
                } else {
                    var v26 uint = 42
                    v5 += v5
                    var v27 int64 = 1
                    _ = v26
                    _ = v27
                }
                v28 := v2[(("" + "\n\t") + "hello")]
                var v29 uint32 = 2147483648
                _ = fi22
                _ = v28
                _ = v29
            }
            fmt.Println(v21)
            _ = fi20
            _ = v21
        }
        v30 := []int{1, 32768, 4294967295}
        _ = v19
        _ = v30
    }
    v31 := GenS1{F0: 383, F1: v5}
    v1 = true
    for fi32 := 0; fi32 < 2; fi32++ {
        v5 = (^(v5))
        _ = fi32
    }
    switch 9223372036854775807 {
    case 42:
        if (true || (2 != 19)) {
            delete(v2, ("a" + "foo bar"))
        } else {
            var v33 bool = true
            var v34 float64 = 0.0
            _ = v33
            _ = v34
        }
    case 1:
        var v35 int64 = 9223372036854775807
        _ = v35
    case 127:
        v4 = v4
    }
    v36 := v2[(((("a" + "a") + ("hello" + "foo bar")) + "a") + ("a" + "\n\t"))]
    v3 = 9
    v37 := GenS2{F0: 82, F1: v3, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
