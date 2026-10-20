package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32, p1 uint8, p2 int16) bool {
    p2 = p2
    var v4 uint32 = 2147483647
    var v5 string = "\n\t"
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 = p0
        v4 -= v4
        _ = fi6
    }
    var v7 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    return ((v4 | v4) < v4)
}
func genFn1() uint16 {
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 19; fi9++ {
        switch (fi9 - fi9) {
        case -2147483648:
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{42, fi9, -500, 42}
                var v12 int = 924
                _ = fi10
                _ = v11
                _ = v12
            }
        case 431:
            v8 = (v8 + (v8 ^ (v8 & 0)))
        }
        v13 := []int{fi9}
        v14 := GenS1{F0: 255}
        _ = fi9
        _ = v13
        _ = v14
    }
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := []uint32{1}
    var v15 float64 = float64(1e10)
    var v16 int16 = 0
    v16 = v16
    if false {
        var v17 uint32 = 2147483648
        _ = v17
    }
    v15 = v15
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 float32 = 1e10
            fmt.Println(v15, v15, fi18)
            if true {
                var v21 uint = 218
                v15 -= v15
                v16 = (v16 << (((v21 << ((42) & 63))) & 63))
                _ = v21
            } else {
                v22 := v2[-884]
                _ = v22
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v15 -= v15
            }
            fmt.Println(v15, v15, v20)
            var v23 uint64 = 1
            _ = fi19
            _ = v20
            _ = v23
        }
        switch ("hello" + "hello") {
        case "foo bar":
            var v24 uint16 = 215
            _ = v24
        }
        _ = fi18
    }
    if ((!false) || true) {
        switch ("" + "hello") {
        case "":
            if false {
                var v25 uint8 = 127
                if (1 <= 100) {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        fmt.Println(fi26)
                        _ = fi26
                    }
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v25 += (^(v25))
                        v28 := GenS1{F0: 65535}
                        _ = fi27
                        _ = v28
                    }
                    var v29 int16 = 598
                    _ = v29
                }
                _ = v25
            } else {
                fmt.Println(v16)
            }
        case "hello":
            var v30 uint = 4294967295
            _ = v30
        default:
            v3 = append(v3, 0)
        }
    } else {
        fmt.Println(v16, v16)
    }
    fmt.Println(v16, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
