package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) string {
    var v4 float32 = -2.25
    var v5 int8 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            var v7 int8 = 1
            var v8 uint64 = 2
            v9 := []int{fi6, (^(fi6)), 442}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := []int{int(p0), (fi6 - (fi6 + (fi6 * 42))), 880}
            switch (fi6 % (((fi6 << ((18446744073709551615) & 63))) | 1)) {
            case -190:
                for fi11 := 0; fi11 < 10; fi11++ {
                    switch "\n\t" {
                    case "a":
                        v4 -= (-2.25 / v4)
                    default:
                        var v12 int32 = (p0 * -27)
                        _ = v12
                    }
                    fmt.Println(v5, v5)
                    _ = fi11
                }
            case 128:
                p0 -= (^(p0))
            case 9223372036854775807:
                p0 = int32(v5)
            }
            _ = v10
        }
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return (("" + "") + ((("\n\t" + "a") + "") + ("\n\t" + ("a" + "a"))))
}

func main() {
    v1 := []int{2147483647, 1, 128, 2147483647}
    var v2 float32 = 0.1
    v3 := make(map[int]bool)
    if (6 == 212) {
        switch ("hello" + "foo bar") {
        case "foo bar":
            var v13 int64 = 42
            _ = v13
        case "a":
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
        default:
            if false {
                fmt.Println(v2)
                v2 += (((v2 * v2) * v2) * 1e10)
            } else {
                v14 := GenS1{F0: 2, F1: 2147483647}
                if (-330 <= 10) {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 uint = 255
                        _ = fi15
                        _ = v16
                    }
                } else {
                    v2 = v2
                    var v17 float32 = ((v2 * v2) + v2)
                    fmt.Println(v17)
                    _ = v17
                }
                _ = v14
            }
        }
    }
    v2 += v2
    v2 = ((v2 * 0.1) - v2)
    for fi18 := 0; fi18 < 11; fi18++ {
        switch 121 {
        case -277:
            v19 := GenS1{F0: 409, F1: 2147483647}
            _ = v19
        }
        switch fi18 {
        case 32768:
            var v20 uint64 = 4294967295
            _ = v20
        default:
            v21 := GenS2{F0: 41}
            _ = v21
        }
        if len(v1) > 0 {
            v1[0] = fi18
        }
        _ = fi18
    }
    var v22 string = ("a" + "")
    for fi23 := 0; fi23 < 20; fi23++ {
        if (true && true) {
            v24 := GenS1{F0: 597, F1: 803}
            _ = v24
        }
        _ = fi23
    }
    var v25 uint64 = 4294967295
    v25 -= (v25 >> ((0) & 63))
    switch v22 {
    case "":
        var v26 float64 = -2.25
        _ = v26
    case "hello":
        v1 = append(v1, 42)
    case "\n\t":
        var v27 float32 = (v2 * float32(0.1))
        _ = v27
    }
    v25 = v25
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
