package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
func genFn0() uint8 {
    v4 := GenS1{F0: 1}
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint64 = 4294967295
        v8 := []int{33}
        v9 := GenS2{F0: 32767, F1: -2147483648, F2: (true && true)}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    v5 += (v5 % ((v5) | 1))
    _ = v4
    _ = v5
    return (v5 << ((65535) & 63))
}

func main() {
    var v1 uint8 = 2
    var v2 int8 = -1
    var v3 uint8 = 165
    v12 := []int{}
    fmt.Println(v1, v3)
    v3 -= (v3 << ((255) & 63))
    if true {
        var v13 int = 4294967295
        var v14 uint8 = (^(v1))
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int = -157
            _ = fi15
            _ = v16
        }
        v2 = (v2 ^ v2)
    }
    var v17 bool = true
    if ("" == ("hello" + "a")) {
        var v18 uint16 = 1
        v1 = 255
        v2 -= 42
        _ = v18
    }
    if (32767 != 346) {
        switch -261 {
        case 1:
            v17 = ((!true) || true)
        case -2147483648:
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS2{F0: -416, F1: 681, F2: true}
                _ = fi19
                _ = v20
            }
        }
    }
    if (0.0 < 1e10) {
        v17 = (3.14159 <= 1.5)
    }
    switch 275 {
    case -9223372036854775808:
        fmt.Println(v17, v3, v2)
    case 21:
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int8 = 42
            v23 := GenS1{F0: 2147483648}
            v17 = v17
            _ = fi21
            _ = v22
            _ = v23
        }
    case 2147483647:
        v24 := GenS2{F0: 1000, F1: 2147483647, F2: v17}
        _ = v24
    }
    v17 = false
    switch "" {
    case "a":
        var v25 int8 = v2
        _ = v25
    case "foo bar":
        if (-2147483648 != -2147483648) {
            v26 := GenS1{F0: 42}
            if v17 {
                if len(v12) > 0 {
                    v12[0] = 32767
                }
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(v2, fi27, fi27)
                    switch (fi27 % ((fi27) | 1)) {
                    case 522:
                        var v28 int16 = 805
                        _ = v28
                    case 9223372036854775807:
                        for fi29 := 0; fi29 < 18; fi29++ {
                            v1 = genFn0()
                            v17 = (688 >= 4294967295)
                            if (v17 || v17) {
                                fmt.Println(fi27, fi27, fi27)
                                v2 -= 127
                            }
                            _ = fi29
                        }
                    case 32768:
                        v17 = v17
                    }
                    _ = fi27
                }
                v2 += (v2 | (v2 - v2))
            }
            var v30 int64 = 1
            _ = v26
            _ = v30
        } else {
            v3 = v1
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v1 += (v1 | (v1 - 180))
        }
    default:
        fmt.Println(v2, v17)
    }
    v3 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
