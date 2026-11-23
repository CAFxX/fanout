package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 string }
func genFn0() int16 {
    v6 := GenS1{F0: -148, F1: "hello", F2: -1}
    switch "foo bar" {
    case "a":
        var v7 int8 = 1
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint = 887
            v9 += (v9 >> ((v9) & 63))
            _ = fi8
            _ = v9
        }
    case "foo bar":
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint16 = 32767
            var v12 uint8 = 53
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(fi13, fi10, fi13)
                if false {
                    fmt.Println(fi13)
                } else {
                    var v14 float32 = float32((1e10 + 1e10))
                    v15 := []int{32768, (-(fi10)), 32767, fi13}
                    _ = v14
                    _ = v15
                }
                var v16 int8 = -68
                _ = fi13
                _ = v16
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        v17 := []int{128, 42}
        _ = v17
    }
    var v18 float64 = float64(0.0)
    fmt.Println(v18)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint64 = 42
        var v21 int = fi19
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = v18
    _ = v6
    return -755
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = -532
    var v3 uint64 = 551
    var v4 uint16 = 1
    var v5 uint64 = 4294967295
    v22 := []int{-411, 2147483647}
    v1 = (v1 - 2147483647)
    switch 42 {
    case 127:
        switch "foo bar" {
        case "foo bar":
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 int32 = int32(v4)
                var v25 string = ""
                _ = fi23
                _ = v24
                _ = v25
            }
        case "\n\t":
            for fi26 := 0; fi26 < 13; fi26++ {
                if true {
                    var v27 int64 = -9223372036854775808
                    _ = v27
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    if (true && true) {
                        fmt.Println(v4, fi26, v3)
                    } else {
                        v29 := GenS1{F0: -299, F1: "a", F2: int64(v4)}
                        var v30 bool = (-2.25 <= (0.1 + 1e10))
                        var v31 float64 = 1e10
                        var v32 int = fi26
                        v33 := GenS2{F0: "a"}
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = fi28
                }
                _ = fi26
            }
        case "hello":
            var v34 uint8 = 30
            _ = v34
        default:
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
    case -283:
        if false {
            v1 += v1
            fmt.Println(v1, v2, v2)
        }
    }
    v3 += v5
    if len(v22) > 0 {
        v22[0] = 128
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        if (true || (!true)) {
            v3 = (v5 / ((0) | 1))
            var v36 int16 = -109
            for fi37 := 0; fi37 < 14; fi37++ {
                if false {
                    var v38 float32 = 1.5
                    var v39 int8 = -128
                    _ = v38
                    _ = v39
                }
                var v40 int64 = int64(fi37)
                _ = fi37
                _ = v40
            }
            _ = v36
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            v42 := GenS1{F0: 42, F1: ("a" + "hello"), F2: (v2 + v2)}
            v5 += v5
            _ = fi41
            _ = v42
        }
        var v43 int16 = 981
        _ = fi35
        _ = v43
    }
    var v44 int = 4294967295
    var v45 uint8 = 123
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
