package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool, p1 int, p2 int8) uint {
    switch (p1 / (((p1 * p1)) | 1)) {
    case 488:
        var v7 uint32 = 1
        _ = v7
    case -517:
        switch "a" {
        case "\n\t":
            var v8 string = "a"
            _ = v8
        case "foo bar":
            var v9 float32 = float32(0.0)
            _ = v9
        }
    case 42:
        var v10 uint64 = 4294967295
        _ = v10
    default:
        if p0 {
            v11 := GenS1{F0: 0, F1: 42}
            v12 := GenS2{F0: 42, F1: -9223372036854775808, F2: 851}
            p1 = 0
            _ = v11
            _ = v12
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 float32 = (-2.25 / 0.1)
    var v2 int32 = 42
    var v3 int8 = 100
    var v4 int = 32767
    v5 := []int{978, v4}
    v6 := []int{(v4 << ((493) & 63)), v4, (v4 & int(v3))}
    switch (v4 % ((-2147483648) | 1)) {
    case 32767:
        var v13 int16 = -1
        _ = v13
    case 127:
        var v14 string = "\n\t"
        _ = v14
    case 32768:
        v3 += 127
    }
    fmt.Println(v2, v1)
    v15 := GenS2{F0: 2, F1: -685, F2: 4294967295}
    var v16 int64 = 42
    v1 = -2.25
    var v17 uint = 65535
    v18 := GenS2{F0: (v17 | v17), F1: (v16 & v16), F2: 65535}
    var v19 int16 = 0
    if len(v6) > 0 {
        v6[0] = (v4 % (((v4 | (v4 - v4))) | 1))
    }
    _ = len(v5)
    if (!false) {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int64 = v16
            _ = fi20
            _ = v21
        }
        switch -164 {
        case -2147483648:
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 string = "\n\t"
                    var v25 int = v4
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        case 0:
            _ = len(v5)
        case 4294967295:
            if false {
                fmt.Println(v3, v2, v2)
                for fi26 := 0; fi26 < 12; fi26++ {
                    v1 -= v1
                    v27 := GenS1{F0: (^(v2)), F1: 128}
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: v2, F1: 2}
                _ = v28
            }
        default:
            if false {
                v5 = append(v5, -410)
                v29 := GenS3{F0: 675, F1: 0, F2: true}
                if len(v5) > 0 {
                    v5[0] = (v4 & (v4 & -2147483648))
                }
                _ = v29
            } else {
                var v30 uint8 = 128
                if false {
                    v30 += v30
                } else {
                    var v31 string = "hello"
                    v17 += (v17 << ((v17) & 63))
                    _ = v31
                }
                _ = v30
            }
        }
    }
    v16 = v16
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            v34 := GenS4{F0: v16}
            _ = fi33
            _ = v34
        }
        v35 := GenS4{F0: v16}
        _ = fi32
        _ = v35
    }
    var v36 uint16 = 1
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
