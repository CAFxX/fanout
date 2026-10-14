package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 uint32 = 4294967295
            v8 := []int{fi5}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int32 = (p0 / ((p0) | 1))
        _ = fi5
        _ = v9
    }
    var v10 int32 = (p0 + p0)
    v11 := []int{2147483647}
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    _ = v11
    return 105
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint64 = 758
    var v4 uint32 = 65535
    var v12 int64 = -714
    if (617 <= 795) {
        v13 := GenS1{F0: 1, F1: "a", F2: -1}
        _ = v13
    } else {
        if (1 >= -1) {
            fmt.Println(v3, v12)
            var v14 uint64 = v3
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := []int{(fi15 >> ((0) & 63)), (fi15 + (fi15 + (fi15 / ((fi15) | 1)))), (fi15 | fi15)}
                switch (("\n\t" + "foo bar") + (("\n\t" + "a") + ("a" + "a"))) {
                case "foo bar":
                    fmt.Println(v1, v3, v12)
                default:
                    var v17 string = "hello"
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            if ((127 >= 100) && (v12 > 1)) {
                var v18 string = ((("hello" + ((("foo bar" + "a") + ((("hello" + "foo bar") + ("a" + "a")) + "\n\t")) + ("foo bar" + "\n\t"))) + "\n\t") + "\n\t")
                var v19 float32 = -2.25
                var v20 float32 = (-2.25 + 3.14159)
                _ = v18
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 int16 = 1000
                    switch "a" {
                    case "":
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v24 := GenS2{F0: v12}
                            _ = fi23
                            _ = v24
                        }
                    default:
                        var v25 uint64 = 2
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
            }
        }
        v1 = (!((1.5 >= (((1e10 * 1.5) * -2.25) / float64(float32((-2.25 * 1e10))))) && v1))
    }
    var v26 int = 1
    var v27 uint16 = 2
    if (v1 || true) {
        var v28 int16 = -32768
        _ = v28
    } else {
        switch v26 {
        case 1:
            v29 := []int{4294967295, v26}
            _ = v29
        case 2147483647:
            _ = len(v2)
        }
    }
    v27 = v27
    v30 := GenS2{F0: 1}
    if v1 {
        v27 -= 0
    } else {
        if (v1 || (-789 > -860)) {
            if true {
                var v31 uint = 2
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := []int{v26, -9223372036854775808, v26, (fi32 ^ fi32)}
                    v1 = v1
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v31, v1)
                _ = v31
            }
        }
        fmt.Println(v12, v1)
    }
    switch "a" {
    case "":
        v26 = int(v27)
    case "a":
        fmt.Println(v12, v3)
    }
    v34 := GenS3{F0: v1}
    fmt.Println(v12, v12)
    v35 := []int{(v26 - v26), 128}
    _ = len(v35)
    var v36 uint32 = v4
    v35 = append(v35, v26)
    switch (v26 >> ((65535) & 63)) {
    case -1:
        var v37 uint32 = 637
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
