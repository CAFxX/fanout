package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string, p1 int32, p2 uint) int64 {
    var v7 int = -60
    if true {
        p0 = (p0 + p0)
        var v8 int8 = -74
        for fi9 := 0; fi9 < 19; fi9++ {
            switch fi9 {
            case 1:
                if true {
                    var v10 uint32 = 65535
                    v10 += 319
                    _ = v10
                }
            default:
                if (4294967295 < 0) {
                    if (-1 != p1) {
                        v8 = 1
                    }
                    var v11 uint16 = 65535
                    _ = v11
                } else {
                    fmt.Println(p0, fi9)
                    p1 = p1
                }
            }
            switch v7 {
            case -1:
                v7 = (v7 + v7)
            case -344:
                p2 -= (p2 & p2)
            }
            p0 = p0
            _ = fi9
        }
        _ = v8
    } else {
        v12 := []int{v7, (v7 - (v7 - v7))}
        var v13 int16 = 833
        v14 := GenS1{F0: 1}
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: p2}
    switch p0 {
    case "\n\t":
        switch p0 {
        case "\n\t":
            if true {
                fmt.Println(p0, p1)
                v16 := GenS1{F0: 1}
                _ = v16
            } else {
                var v17 int16 = -1
                if (882 >= 65535) {
                    fmt.Println(p0, p2, v17)
                    if false {
                        v18 := []int{9223372036854775807, v7}
                        v19 := GenS1{F0: (p2 - 42)}
                        _ = v18
                        _ = v19
                    }
                }
                _ = v17
            }
        case "foo bar":
            if ((((0.1 / 1.5) * -2.25) + 1e10) == (0.0 + (1.5 * 1.5))) {
                var v20 uint64 = 2
                for fi21 := 0; fi21 < 8; fi21++ {
                    p2 += (p2 / (((p2 + (p2 | (p2 & p2)))) | 1))
                    if (false && ("hello" < (p0 + p0))) {
                        var v22 uint64 = 4294967295
                        _ = v22
                    }
                    _ = fi21
                }
                switch "foo bar" {
                case "foo bar":
                    var v23 int = 2147483647
                    _ = v23
                default:
                    fmt.Println(p2)
                }
                _ = v20
            } else {
                v24 := []int{(v7 / ((v7) | 1))}
                fmt.Println(v7)
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                _ = v24
            }
        case "a":
            var v25 float32 = -2.25
            _ = v25
        default:
            fmt.Println(p0, p1)
        }
    case "foo bar":
        for fi26 := 0; fi26 < 7; fi26++ {
            p0 = p0
            _ = fi26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 1
}
func genFn1(p0 int16) int64 {
    var v27 uint = 65535
    var v28 bool = true
    var v29 int16 = p0
    v28 = false
    var v30 int32 = 578
    _ = p0
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    return 42
}

func main() {
    var v1 uint64 = 805
    v2 := make(map[int]int)
    var v3 int32 = 200
    var v4 int16 = 42
    v5 := []string{}
    v6 := []string{"", "hello", ""}
    if (9223372036854775807 != -505) {
        _ = len(v2)
        v31 := GenS1{F0: 0}
        fmt.Println(v4)
        _ = v31
    }
    if false {
        if (!(!false)) {
            v32 := GenS1{F0: 949}
            _ = v32
        }
        v33 := GenS1{F0: 612}
        var v34 float64 = 1.5
        _ = v33
        _ = v34
    } else {
        fmt.Println(v1, v1)
        v6 = append(v6, "")
    }
    v3 += (^(v3))
    v35 := GenS1{F0: 65535}
    var v36 int32 = (v3 << ((42) & 63))
    v36 += v36
    var v37 int32 = (v36 * v3)
    var v38 uint64 = v1
    fmt.Println(v37)
    fmt.Println(v1, v36, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v39 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
