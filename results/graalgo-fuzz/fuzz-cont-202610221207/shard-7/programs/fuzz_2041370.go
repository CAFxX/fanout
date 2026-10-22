package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0(p0 int64, p1 uint64) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = 485
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int8 = 100
            switch ("a" + "a") {
            case "\n\t":
                v11 := GenS1{F0: p0}
                _ = v11
            default:
                if (!false) {
                    var v12 int8 = -119
                    _ = v12
                } else {
                    v13 := GenS2{F0: p0, F1: (p1 * 4294967295), F2: 32767}
                    if (!(2 == 2)) {
                        v10 = v10
                        v14 := []int{-140, fi7, fi9}
                        _ = v14
                    } else {
                        v15 := []int{}
                        v10 += int8(p0)
                        fmt.Println(p0)
                        _ = v15
                    }
                    v16 := GenS3{F0: 2147483648, F1: -2147483648}
                    _ = v13
                    _ = v16
                }
            }
            v17 := GenS3{F0: 456, F1: 42}
            _ = fi9
            _ = v10
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    var v18 int32 = 0
    if ((("a" + "foo bar") != "") || (!(!(65535 < 255)))) {
        var v19 float64 = (1.5 + (1e10 * 1e10))
        v20 := []int{}
        var v21 bool = (float64(3.14159) == ((v19 / float64((3.14159 / (0.1 + (0.0 * 3.14159))))) - float64(1e10)))
        _ = v19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v18
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 9223372036854775807
    var v3 uint = 2
    v4 := []int{0}
    var v5 int64 = v2
    var v6 float32 = -2.25
    var v22 bool = (965 < 32767)
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 float64 = 1.5
        _ = fi23
        _ = v24
    }
    v1[""] = 786
    for fi25 := 0; fi25 < 11; fi25++ {
        switch "" {
        case "":
            v3 = 138
        case "hello":
            v26 := v1["\n\t"]
            _ = v26
        case "foo bar":
            delete(v1, ("a" + "hello"))
        }
        _ = fi25
    }
    v22 = false
    v3 = v3
    var v27 int = 127
    v1[(("hello" + "foo bar") + "a")] = 1
    var v28 bool = (v3 <= v3)
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
