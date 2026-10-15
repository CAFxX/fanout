package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 bool, p2 bool) string {
    switch (((("foo bar" + (p0 + p0)) + p0) + "") + "") {
    case "a":
        v6 := []int{-45}
        _ = v6
    }
    fmt.Println(p1, p2, p2)
    v7 := GenS1{F0: -1, F1: -478}
    var v8 int = 132
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS2{F0: 2147483647, F1: (v8 >> ((375) & 63))}
        fmt.Println(p2, fi9)
        if p2 {
            var v11 int32 = 1
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int64 = 0
                v14 := []int{}
                var v15 uint64 = 1
                var v16 uint64 = (^(v15))
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 uint = 782
            _ = v11
            _ = v17
        } else {
            switch -658 {
            case -9223372036854775808:
                v18 := GenS2{F0: 2, F1: 42}
                _ = v18
            default:
                v19 := []int{411, v8, (v8 + v8)}
                _ = v19
            }
            p1 = (0 > 1)
            p1 = ("a" >= (p0 + p0))
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 bool) int8 {
    p0 = p0
    var v20 int64 = -611
    v21 := GenS2{F0: 4294967295, F1: 128}
    _ = p0
    _ = v20
    _ = v21
    return -65
}

func main() {
    var v1 int32 = 42
    var v2 float64 = float64(1.5)
    var v3 int64 = 764
    var v4 string = ("foo bar" + ((("\n\t" + "\n\t") + "foo bar") + "a"))
    v5 := []string{"hello"}
    v3 = (v3 % ((v3) | 1))
    if (-32768 <= -1) {
        v3 = v3
        var v22 int = 2147483647
        v5 = append(v5, v4)
        _ = v22
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = v4
    v23 := GenS3{F0: 1, F1: 733, F2: 1}
    var v24 uint64 = 1
    var v25 float32 = ((1.5 + 0.1) + 0.0)
    for fi26 := 0; fi26 < 20; fi26++ {
        v5 = append(v5, v4)
        v27 := GenS4{F0: 42}
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 int8 = 42
            if (!((v24 << ((42) & 63)) >= (v24 / ((v24) | 1)))) {
                v2 += v2
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v30 := GenS3{F0: v24, F1: 65535, F2: 128}
                _ = v30
            }
            v29 = (v29 * (v29 - v29))
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    if true {
        v5 = append(v5, (v4 + v4))
        switch v4 {
        case "\n\t":
            var v31 uint64 = (v24 & (v24 | v24))
            _ = v31
        }
        var v32 uint64 = v24
        _ = v32
    } else {
        var v33 int64 = int64(v1)
        switch 0 {
        case 0:
            v5 = append(v5, v4)
        case 1:
            switch "hello" {
            case "foo bar":
                switch 9223372036854775807 {
                case -486:
                    var v34 int64 = v33
                    _ = v34
                case 9223372036854775807:
                    v1 = v1
                case -75:
                    v24 = (v24 << ((48) & 63))
                default:
                    v5 = append(v5, ("foo bar" + ((v4 + "a") + "")))
                }
            default:
                v35 := GenS3{F0: v24, F1: 190, F2: 253}
                _ = v35
            }
        case 42:
            fmt.Println(v33, v3, v4)
        default:
            fmt.Println(v3, v33, v2)
        }
        _ = v33
    }
    var v36 uint16 = 255
    for fi37 := 0; fi37 < 3; fi37++ {
        for fi38 := 0; fi38 < 17; fi38++ {
            v3 = v3
            _ = fi38
        }
        _ = fi37
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        if false {
            v40 := GenS1{F0: v3, F1: 2147483647}
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 uint = 0
                v4 = v4
                _ = fi41
                _ = v42
            }
            for fi43 := 0; fi43 < 18; fi43++ {
                _ = len(v5)
                fmt.Println(v36, v1, fi39)
                if (!(!(((42 >= 1000) || false) || true))) {
                    if len(v5) > 0 {
                        v5[0] = ("" + (v4 + v4))
                    }
                    if ((v24 % ((2) | 1)) < (v24 | v24)) {
                        v24 += v24
                    } else {
                        if (v36 >= v36) {
                            var v44 int8 = 1
                            v45 := GenS1{F0: v3, F1: (-(fi39))}
                            v2 = float64(((0.0 - (v25 + 1.5)) * v25))
                            v2 -= v2
                            var v46 int32 = v1
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                    }
                }
                _ = fi43
            }
            _ = v40
        }
        var v47 int = 0
        if false {
            v5 = append(v5, "")
            var v48 int64 = v3
            _ = v48
        } else {
            var v49 int = (fi39 | fi39)
            _ = v49
        }
        _ = fi39
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
