package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int64 {
    var v6 int16 = -608
    fmt.Println(v6)
    v6 = (v6 % ((v6) | 1))
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "hello"
    v3 := []string{(v2 + v2), v2, (("a" + "\n\t") + ""), ("\n\t" + v2)}
    var v4 string = "a"
    var v5 uint = 4294967295
    fmt.Println(v2, v2, v4)
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = 1
        _ = fi8
        _ = v9
    }
    fmt.Println(v2, v5)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: genFn0()}
        _ = fi10
        _ = v11
    }
    v3 = append(v3, v2)
    if false {
        _ = len(v1)
        var v12 int8 = 0
        _ = v12
    } else {
        var v13 int = -71
        var v14 int32 = 608
        fmt.Println(v2, v14, v2)
        _ = v13
        _ = v14
    }
    _ = len(v3)
    v7 = (v7 ^ v7)
    v5 -= (v5 >> ((v5) & 63))
    var v15 uint32 = 0
    for fi16 := 0; fi16 < 8; fi16++ {
        switch fi16 {
        case 297:
            var v17 bool = (((1.5 - (float64(1.5) * (((0.0 / (((-2.25 * (3.14159 * 1.5)) - -2.25) * -2.25)) + 1.5) - (0.0 * 0.1)))) * 1e10) > 3.14159)
            _ = v17
        case 12:
            v18 := GenS1{F0: genFn0()}
            _ = v18
        }
        _ = fi16
    }
    if ("hello" < v4) {
        switch 128 {
        case -9223372036854775808:
            var v19 float64 = 0.0
            _ = v19
        case -1:
            if (!false) {
                fmt.Println(v2, v2)
                if true {
                    var v20 int64 = genFn0()
                    _ = v20
                }
                fmt.Println(v4)
            } else {
                var v21 string = v2
                var v22 uint64 = 900
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v21
                _ = v22
            }
        }
    } else {
        v2 = (v2 + "")
        delete(v1, 2)
        var v23 uint8 = 42
        _ = v23
    }
    v15 -= v15
    if (true || (true && (-9223372036854775808 < -467))) {
        v15 -= (v15 >> ((v5) & 63))
        var v24 int32 = -827
        _ = v24
    } else {
        fmt.Println(v15, v15)
        fmt.Println(v4)
    }
    var v25 float64 = (-2.25 / 0.1)
    for fi26 := 0; fi26 < 9; fi26++ {
        v7 += v7
        v2 = (v4 + ("a" + v4))
        switch (fi26 & (-(fi26))) {
        case 32768:
            if true {
                for fi27 := 0; fi27 < 7; fi27++ {
                    v5 += v5
                    var v28 uint = v5
                    _ = fi27
                    _ = v28
                }
            }
        case -654:
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 int = fi29
                var v31 int16 = 42
                _ = fi29
                _ = v30
                _ = v31
            }
        case -995:
            switch (fi26 | 2147483647) {
            case 854:
                if false {
                    v3 = append(v3, v2)
                }
            case 127:
                v32 := GenS1{F0: (v7 | 490)}
                _ = v32
            default:
                v4 = (v2 + "hello")
            }
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
