package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 int8; F2 int16 }
func genFn0(p0 int16, p1 uint32) int16 {
    p1 = p1
    v6 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]int)
    var v3 int16 = 671
    var v4 uint16 = 64
    var v5 uint32 = (^(v1))
    var v7 bool = (64 > 351)
    v3 += v3
    v8 := GenS2{F0: -974, F1: 42}
    v9 := []int{32768}
    var v10 int16 = (^(v3))
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            if false {
                switch (fi11 << ((42) & 63)) {
                case 127:
                    v5 -= 1
                default:
                    v1 = v1
                }
            }
            if v7 {
                var v12 uint8 = 127
                v13 := GenS1{F0: 1}
                _ = v12
                _ = v13
            }
        }
        var v14 uint16 = v4
        _ = fi11
        _ = v14
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v15 int64 = 9223372036854775807
    if v7 {
        var v16 int64 = v15
        switch -1 {
        case -982:
            v2[-9223372036854775808] = 128
        case 1:
            v15 = (v16 & (v16 ^ v15))
        default:
            for fi17 := 0; fi17 < 11; fi17++ {
                if false {
                    var v18 uint8 = 0
                    v5 += v5
                    if len(v9) > 0 {
                        v9[0] = fi17
                    }
                    _ = v18
                } else {
                    v19 := GenS3{F0: 255, F1: -1, F2: fi17}
                    _ = v19
                }
                if true {
                    if false {
                        var v20 string = ""
                        v21 := GenS4{F0: uint8(v1)}
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi17
            }
        }
        _ = v16
    }
    switch 0 {
    case -1:
        if (uint64(v1) == 786) {
            var v22 uint = 1
            switch "hello" {
            case "a":
                for fi23 := 0; fi23 < 5; fi23++ {
                    if ((-29 < 100) || (42 == 42)) {
                        v9 = append(v9, (-(fi23)))
                    }
                    v24 := GenS5{F0: int32(fi23)}
                    _ = fi23
                    _ = v24
                }
            }
            _ = v22
        } else {
            v4 += (v4 ^ v4)
            fmt.Println(v3, v3, v15)
        }
    default:
        var v25 int8 = 0
        _ = v25
    }
    delete(v2, 42)
    _ = len(v2)
    var v26 int64 = v15
    if (!(9223372036854775807 == (v26 * -470))) {
        fmt.Println(v26, v1)
        var v27 string = ("a" + "a")
        _ = v27
    } else {
        switch 9223372036854775807 {
        case 42:
            v28 := GenS4{F0: 226}
            _ = v28
        case -9223372036854775808:
            if len(v9) > 0 {
                v9[0] = 405
            }
        case 872:
            v29 := GenS6{F0: v3, F1: -128, F2: (-(v10))}
            _ = v29
        }
        switch ("\n\t" + "a") {
        case "":
            var v30 uint64 = 42
            _ = v30
        case "foo bar":
            var v31 int16 = -32768
            _ = v31
        case "a":
            delete(v2, 627)
        default:
            var v32 float64 = -2.25
            _ = v32
        }
    }
    switch "\n\t" {
    case "hello":
        v5 = (v5 ^ (^(v1)))
    case "":
        var v33 bool = v7
        _ = v33
    case "a":
        if len(v9) > 0 {
            v9[0] = 655
        }
    default:
        var v34 int64 = -327
        _ = v34
    }
    var v35 uint16 = (v4 >> ((564) & 63))
    var v36 uint = 1
    for fi37 := 0; fi37 < 19; fi37++ {
        v5 -= v1
        for fi38 := 0; fi38 < 14; fi38++ {
            switch "\n\t" {
            case "a":
                var v39 uint = v36
                _ = v39
            default:
                v10 = (v3 * genFn0(32767, v5))
            }
            v1 -= v1
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
