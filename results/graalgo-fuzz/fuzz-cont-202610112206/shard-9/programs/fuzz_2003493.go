package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint64 = 843
    var v2 int = -1
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 float64 = 1e10
    var v6 int8 = 3
    v4[987] = ((("\n\t" + "\n\t") + "foo bar") + "hello")
    fmt.Println(v6)
    v5 = 1e10
    var v7 int32 = 0
    switch "hello" {
    case "foo bar":
        var v8 float64 = ((((0.0 * v5) * v5) / v5) * -2.25)
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float64 = (v5 / ((v5 / v5) - (-2.25 + 1.5)))
            fmt.Println(v7, v6, v7)
            if ((fi9 | v2) == 128) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    v7 = 1
                    _ = fi11
                }
            } else {
                if true {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        fmt.Println(v5)
                        v13 := []int{fi12, (fi9 & -237), fi9, fi9}
                        fmt.Println(v10, fi12, v5)
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 int32 = -2147483648
                        v15 += (v7 | (v7 >> ((42) & 63)))
                        v5 = v5
                        var v16 uint16 = 42
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := GenS1{F0: v7}
                    _ = v17
                }
                var v18 uint = 377
                _ = v18
            }
            _ = fi9
            _ = v10
        }
    case "":
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 int64 = -1
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v2 = fi21
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := []int{}
            _ = fi22
            _ = v23
        }
        if (true || false) {
            if (v1 != 2) {
                fmt.Println(v6, v1, v1)
                v4[2] = "hello"
                v24 := GenS1{F0: 0}
                _ = v24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := []int{0}
                v2 -= fi21
                var v27 uint8 = 160
                v3["\n\t"] = (v7 > (v7 / (((v7 & 0)) | 1)))
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        _ = fi21
    }
    var v28 uint16 = 434
    var v29 int8 = v6
    v30 := []int{(v2 | -9223372036854775808), v2, v2}
    if true {
        var v31 int = int(v7)
        v32 := GenS2{F0: true, F1: 290}
        v30 = append(v30, -2147483648)
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(fi33, fi33)
            _ = fi33
        }
    }
    v34 := GenS3{F0: 2, F1: 0}
    v29 -= v6
    for fi35 := 0; fi35 < 19; fi35++ {
        v5 = v5
        _ = fi35
    }
    var v36 uint64 = 2
    v6 += (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
