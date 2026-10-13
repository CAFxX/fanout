package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0() uint {
    if false {
        switch 1 {
        case 32768:
            var v6 uint = 42
            _ = v6
        case -100:
            var v7 string = ""
            _ = v7
        default:
            v8 := []int{32767, 1}
            _ = v8
        }
        if true {
            v9 := GenS1{F0: 4294967295, F1: true}
            if (-437 != 2147483647) {
                var v10 uint64 = 0
                var v11 uint32 = 907
                v11 -= 65535
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 20; fi12++ {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v14 := []int{}
                        _ = fi13
                        _ = v14
                    }
                    var v15 int32 = 0
                    _ = fi12
                    _ = v15
                }
            }
            var v16 float32 = 1e10
            _ = v16
            _ = v9
        } else {
            var v17 uint64 = 839
            var v18 float64 = 0.0
            var v19 uint16 = 65535
            _ = v17
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: 4294967295, F1: (62 < 127)}
    var v21 int32 = 42
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 uint = 550
    var v3 int32 = -1
    var v4 int8 = -1
    v5 := make(map[string]string)
    if ((2 == 42) || (!false)) {
        switch "\n\t" {
        case "a":
            v22 := v5["\n\t"]
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 3; fi23++ {
                v3 -= (v3 * 0)
                var v24 uint8 = 127
                if (65535 < 2147483648) {
                    var v25 float32 = 0.1
                    fmt.Println(v1, v3, v25)
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint32 = 413
                        v3 = (v3 & 2147483647)
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
        case "":
            v28 := []int{42, 660}
            _ = v28
        default:
            v2 += (v2 << (((v2 << ((v2) & 63))) & 63))
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS1{F0: 65535, F1: false}
        fmt.Println(v2)
        v31 := GenS1{F0: 2, F1: true}
        _ = fi29
        _ = v30
        _ = v31
    }
    v2 = v2
    v2 += (v2 * 255)
    var v32 bool = (42 > 696)
    for fi33 := 0; fi33 < 14; fi33++ {
        v1 += 1
        var v34 int16 = 32767
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v36 := GenS1{F0: uint32(v2), F1: ((fi35 >> (((v2 * v2)) & 63)) <= -1)}
        _ = fi35
        _ = v36
    }
    v1 = v4
    v4 = (v1 - v1)
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "foo bar":
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 string = ""
            _ = fi37
            _ = v38
        }
    case "hello":
        fmt.Println(v1, v1)
    }
    fmt.Println(v2, v4, v32)
    var v39 int8 = v1
    var v40 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
