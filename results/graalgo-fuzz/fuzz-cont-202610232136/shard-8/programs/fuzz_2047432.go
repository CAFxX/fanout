package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16) int32 {
    p0 = p0
    var v6 int16 = -1
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 2147483647, F1: 2147483647}
            for fi9 := 0; fi9 < 8; fi9++ {
                p0 += p0
                var v10 int64 = 0
                var v11 float64 = (1e10 - 0.0)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    }
    switch "foo bar" {
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint32 = 42
                fmt.Println(fi13, v14, fi13)
                switch (-(fi13)) {
                case 581:
                    v15 := GenS1{F0: 616, F1: -1}
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case "foo bar":
        if (-128 < 127) {
            p0 += (p0 << ((18446744073709551615) & 63))
            var v16 string = "\n\t"
            _ = v16
        } else {
            v17 := []int{32768, 32767, 2147483647, 2147483647}
            _ = v17
        }
    default:
        fmt.Println(v6, p0)
    }
    v18 := GenS2{F0: 845}
    _ = p0
    _ = v18
    _ = v6
    return 2147483647
}

func main() {
    v1 := []uint32{0, 2}
    v2 := make(map[int]int)
    var v3 float64 = (0.0 - 3.14159)
    var v4 float64 = float64(((0.1 + 0.1) + 0.0))
    var v5 int64 = 42
    if (255 > 0) {
        var v19 bool = true
        switch -1 {
        case 32767:
            if v19 {
                v20 := GenS1{F0: 0, F1: -2147483648}
                if v19 {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v5 = (v5 | 42)
                        var v22 float64 = v3
                        v5 -= (v5 & 9223372036854775807)
                        _ = fi21
                        _ = v22
                    }
                    _ = len(v2)
                    var v23 int = -9223372036854775808
                    _ = v23
                }
                _ = v20
            } else {
                v24 := GenS2{F0: 2}
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 float32 = 1e10
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    switch "a" {
                    case "a":
                        fmt.Println(fi25)
                    default:
                        v3 = v4
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        case 1:
            fmt.Println(v5)
        case 79:
            var v27 float64 = v3
            _ = v27
        }
        _ = v19
    } else {
        var v28 uint32 = 404
        v28 = (v28 + v28)
        if true {
            v4 = 1.5
        } else {
            v29 := GenS1{F0: genFn0(42), F1: 42}
            _ = v29
        }
        _ = v28
    }
    v30 := GenS1{F0: genFn0(32767), F1: -1}
    var v31 int8 = 0
    var v32 int64 = v5
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint64 = 18446744073709551615
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 uint8 = 243
            v2[(fi35 & 32768)] = fi33
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    var v37 int32 = -1
    for fi38 := 0; fi38 < 19; fi38++ {
        fmt.Println(v3)
        if (9223372036854775807 > (v5 << ((425) & 63))) {
            v31 = v31
        }
        var v39 int16 = -558
        _ = fi38
        _ = v39
    }
    v37 += genFn0(-32768)
    if true {
        v2[4294967295] = 876
        for fi40 := 0; fi40 < 8; fi40++ {
            _ = len(v2)
            var v41 int8 = int8(fi40)
            _ = fi40
            _ = v41
        }
    } else {
        fmt.Println(v32, v3, v3)
        for fi42 := 0; fi42 < 13; fi42++ {
            for fi43 := 0; fi43 < 12; fi43++ {
                v44 := GenS2{F0: 42}
                var v45 uint32 = 284
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi42
        }
    }
    v2[788] = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
