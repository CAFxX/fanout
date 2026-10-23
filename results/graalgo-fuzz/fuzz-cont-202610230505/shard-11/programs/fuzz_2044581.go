package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{0}
    v3 := make(map[int]bool)
    var v4 float32 = (0.0 + float32((-2.25 - (float64(float32(3.14159)) - 0.0))))
    var v5 int32 = 0
    v6 := GenS1{F0: "", F1: 159, F2: v5}
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int64 = 0
            var v9 uint32 = 714
            _ = fi7
            _ = v8
            _ = v9
        }
        if (false && false) {
            var v10 int64 = 42
            if false {
                fmt.Println(v5)
                for fi11 := 0; fi11 < 13; fi11++ {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    _ = fi11
                }
            }
            _ = v10
        } else {
            var v12 int16 = (v1 << ((0) & 63))
            _ = v12
        }
        var v13 float32 = v4
        _ = v13
    } else {
        var v14 uint = 65535
        switch 132 {
        case 4294967295:
            fmt.Println(v4, v1, v1)
        default:
            var v15 uint8 = 42
            _ = v15
        }
        _ = v14
    }
    v5 -= (v5 + v5)
    fmt.Println(v1, v1, v5)
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int8 = -1
        delete(v3, (fi16 * fi16))
        if false {
            for fi18 := 0; fi18 < 3; fi18++ {
                fmt.Println(fi16)
                v17 = v17
                _ = fi18
            }
            fmt.Println(v4, v1, v4)
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int16 = (v1 % ((v1) | 1))
        switch (fi19 ^ (fi19 | fi19)) {
        case -2147483648:
            var v21 int16 = 465
            _ = v21
        case 4294967295:
            v22 := GenS1{F0: "a", F1: 255, F2: -329}
            _ = v22
        case 157:
            switch ("" + "") {
            case "\n\t":
                fmt.Println(fi19, v1, v5)
            case "hello":
                fmt.Println(v1)
            case "foo bar":
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 int8 = 100
                    delete(v3, (fi23 * (fi23 ^ -2147483648)))
                    v1 = (v20 * 851)
                    var v25 int16 = (v20 << ((628) & 63))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS2{F0: "", F1: (false && true)}
            var v28 bool = (-1 <= 100)
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v5, v4)
    _ = len(v2)
    var v29 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
