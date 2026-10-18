package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 int32 }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 int, p1 int8) uint {
    v7 := []int{p0, 32768}
    if (65535 > 728) {
        var v8 string = ("a" + (("foo bar" + "") + ("hello" + "foo bar")))
        v7 = append(v7, p0)
        p1 = (^(p1))
        _ = v8
    }
    fmt.Println(p0, p1)
    var v9 float32 = (((0.1 - 0.0) * (1e10 * (float32(1e10) - 1.5))) * -2.25)
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 65535
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 uint32 = 4294967295
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    v4 += 0
    var v10 uint64 = 42
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v3, v10, fi11)
        _ = fi11
    }
    v3 = (false || (true && (-996 != -1)))
    fmt.Println(v1, v3, v10)
    var v12 uint = uint(v10)
    delete(v6, 128)
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            delete(v6, fi14)
            fmt.Println(fi13, v12)
            var v15 float32 = 1e10
            _ = fi14
            _ = v15
        }
        switch -1 {
        case 4294967295:
            if v3 {
                v10 += 320
            }
        }
        _ = fi13
    }
    if (!(!(v3 || false))) {
        if false {
            v3 = v3
            v16 := GenS1{F0: 866, F1: 271}
            _ = v16
        }
    }
    v17 := GenS2{F0: 32767}
    switch 38 {
    case 127:
        switch int(v4) {
        case -2147483648:
            if ((!v3) && false) {
                v18 := GenS3{F0: (v12 | v12), F1: 720, F2: 2147483647}
                var v19 float64 = (0.0 + 1e10)
                _ = v18
                _ = v19
            } else {
                v4 -= 2
            }
        case 136:
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS1{F0: 869, F1: 668}
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int64 = -426
                    var v24 int16 = -32768
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                var v25 string = "foo bar"
                _ = fi20
                _ = v21
                _ = v25
            }
        case 619:
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 uint64 = (^(v10))
                delete(v6, 9223372036854775807)
                _ = fi26
                _ = v27
            }
        }
    case 2147483647:
        var v28 float64 = 1.5
        _ = v28
    default:
        v29 := GenS4{F0: (^(v10)), F1: 42, F2: -543}
        _ = v29
    }
    v30 := []int{783, 127}
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
