package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := []uint32{83, 70}
    var v2 uint = 580
    var v3 float64 = 0.1
    v4 := make(map[int]string)
    if true {
        var v5 int32 = 147
        _ = v5
    }
    var v6 uint = (v2 >> ((v2) & 63))
    v6 = v6
    v3 -= (1.5 - v3)
    if false {
        switch 32767 {
        case 32767:
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: 484, F1: 42, F2: (v2 * v6)}
                if len(v1) > 0 {
                    v1[0] = 0
                }
                _ = fi7
                _ = v8
            }
        case 32768:
            if (-87 != int8(v2)) {
                v9 := GenS2{F0: v6, F1: 0}
                _ = v9
            } else {
                fmt.Println(v3)
            }
        }
        if ((!true) && (42 >= 100)) {
            if true {
                var v10 uint64 = 0
                _ = v10
            }
        } else {
            fmt.Println(v3)
        }
    } else {
        v11 := v4[9223372036854775807]
        var v12 float32 = 1e10
        fmt.Println(v2, v6, v6)
        _ = v11
        _ = v12
    }
    v4[0] = ("\n\t" + ("hello" + ""))
    if false {
        for fi13 := 0; fi13 < 5; fi13++ {
            v2 += v2
            _ = fi13
        }
        switch ("hello" + "\n\t") {
        case "\n\t":
            v6 = (^(v2))
        }
        v1 = append(v1, 1)
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 float64 = v3
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 int = fi16
                _ = fi16
                _ = v17
            }
            if (-808 == 42) {
                v6 = v6
            } else {
                v18 := GenS3{F0: v2, F1: 1, F2: "a"}
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS1{F0: 0, F1: 0, F2: (v2 | (v2 - (v2 % ((1) | 1))))}
                    v21 := GenS4{F0: 1000, F1: -617}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                v22 := GenS5{F0: 1}
                _ = v18
                _ = v22
            }
            _ = fi14
            _ = v15
        }
        var v23 float64 = (1.5 - v3)
        v2 -= v6
        _ = v23
    }
    var v24 uint16 = 125
    v4[620] = ("" + "")
    var v25 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
