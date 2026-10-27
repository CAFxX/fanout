package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535, 16, 2}
    v3 := []int{983, 15}
    v4 := make(map[int]uint32)
    var v5 bool = (((1 > 298) || false) || (2147483647 > -608))
    v6 := make(map[uint32]bool)
    v5 = v5
    delete(v1, ("hello" + "foo bar"))
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            if (((1.5 - float64((1.5 * 1.5))) * 0.0) == 3.14159) {
                var v9 int32 = -2147483648
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v9
            } else {
                v5 = v5
                fmt.Println(fi7, fi7)
                v10 := GenS1{F0: 42, F1: -32768}
                _ = v10
            }
            _ = fi8
        }
        fmt.Println(fi7, fi7)
        switch (fi7 >> ((1) & 63)) {
        case 32768:
            fmt.Println(fi7)
        case -1:
            if v5 {
                v5 = true
                delete(v1, "")
            }
        }
        _ = fi7
    }
    if ("hello" >= "hello") {
        v11 := GenS1{F0: 255, F1: 32767}
        var v12 int64 = 1
        switch 128 {
        case -1:
            v13 := GenS1{F0: 255, F1: 0}
            _ = v13
        case 2147483647:
            delete(v4, 42)
        default:
            switch -843 {
            case 32768:
                fmt.Println(v5, v5)
            default:
                fmt.Println(v5, v12)
            }
        }
        _ = v11
        _ = v12
    }
    var v14 float64 = -2.25
    var v15 uint16 = 2
    var v16 int64 = -323
    v15 = (v15 << ((953) & 63))
    var v17 int = -1
    v18 := GenS1{F0: 544, F1: 537}
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!v5) {
            v20 := GenS2{F0: v17, F1: 2147483647}
            if false {
                _ = len(v1)
                v2 = append(v2, 2)
                fmt.Println(fi19, v17)
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    switch "a" {
                    case "hello":
                        var v22 uint64 = uint64(v15)
                        _ = v22
                    case "":
                        fmt.Println(fi19, v14)
                    }
                    v23 := GenS3{F0: -124, F1: ((-(v17)) == 128)}
                    _ = fi21
                    _ = v23
                }
                fmt.Println(v14, v17, v14)
                v14 = float64(-2.25)
            }
            _ = v20
        } else {
            var v24 int16 = 1000
            if (!((false && (true || (!((fi19 + (v17 / ((fi19) | 1))) >= v17)))) && (0 <= 4294967295))) {
                v25 := GenS2{F0: v17, F1: -6}
                fmt.Println(v16, v5, v24)
                if (true && (false && v5)) {
                    v24 -= v24
                    v14 -= (v14 / v14)
                } else {
                    var v26 uint16 = 2
                    var v27 int8 = 1
                    _ = v26
                    _ = v27
                }
                _ = v25
            }
            v14 = v14
            _ = v24
        }
        _ = fi19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 bool = (95 <= 128)
    fmt.Println(v15, v17, v28)
    v29 := GenS2{F0: 1, F1: 2147483647}
    fmt.Println(v14, v15)
    v30 := GenS1{F0: 0, F1: -32768}
    v17 += v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
