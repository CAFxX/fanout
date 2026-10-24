package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 1.5
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{756, fi5, (fi5 * 4294967295)}
        _ = fi5
        _ = v6
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v7 float64 = float64(1.5)
        _ = v7
    case -1:
        if (890 >= 2) {
            var v8 float32 = 0.1
            var v9 uint = 255
            var v10 uint32 = 65535
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 int16 = 86
            _ = v11
        }
    case 332:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint64 = 2
            for fi14 := 0; fi14 < 6; fi14++ {
                delete(v3, 65535)
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    }
    v2 -= v2
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS1{F0: ("foo bar" + "hello"), F1: -572, F2: ""}
        v17 := []int{fi15}
        if false {
            v18 := GenS2{F0: ("" + (("foo bar" + "hello") + "a")), F1: -1}
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 += v2
    var v19 float32 = 1.5
    var v20 int8 = 1
    var v21 int8 = (v20 & v20)
    v3[65535] = 9223372036854775807
    fmt.Println(v19, v19)
    if true {
        switch -737 {
        case 0:
            var v22 int16 = 1
            _ = v22
        case 4294967295:
            var v23 bool = true
            _ = v23
        case -511:
            v24 := v3[2147483647]
            _ = v24
        default:
            v19 = ((-2.25 + (v19 - (v19 * -2.25))) / (v19 - float32(0.1)))
        }
        v25 := GenS2{F0: "", F1: -1}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 13; fi26++ {
            if (!false) {
                v1 = v1
                var v27 uint64 = 4294967295
                if false {
                    var v28 float64 = v2
                    _ = v28
                } else {
                    var v29 int = (fi26 << ((499) & 63))
                    var v30 uint = 219
                    _ = v29
                    _ = v30
                }
                _ = v27
            } else {
                var v31 bool = (!false)
                _ = v31
            }
            fmt.Println(fi26, v21, v19)
            _ = fi26
        }
        var v32 int8 = (^(v21))
        var v33 uint = 499
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        fmt.Println(fi34, v19, v1)
        fmt.Println(v19, fi34, v1)
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 uint8 = 0
            v2 = (0.0 * float64((0.0 * 1.5)))
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v2 -= v2
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 11; fi38++ {
            v2 += 0.0
            _ = fi38
        }
        _ = fi37
    }
    v39 := []int{4294967295, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
}
