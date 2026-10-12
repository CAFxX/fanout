package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint32, p1 uint64) uint32 {
    v6 := []int{}
    if (true && (-2147483648 >= -332)) {
        v7 := GenS1{F0: 356, F1: 0}
        if true {
            p0 -= p0
            if true {
                for fi8 := 0; fi8 < 17; fi8++ {
                    p0 = p0
                    _ = fi8
                }
                var v9 uint16 = 255
                fmt.Println(v9, p1)
                _ = v9
            }
        } else {
            _ = len(v6)
        }
        var v10 int64 = 0
        _ = v10
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 427
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]uint32)
    var v11 uint = 42
    v12 := []int{128, 1}
    var v13 float64 = 0.0
    if true {
        v14 := GenS1{F0: (v11 * 215), F1: 509}
        _ = v14
    }
    if true {
        var v15 string = "a"
        _ = v15
    }
    v16 := GenS1{F0: (v11 - 2), F1: 1}
    for fi17 := 0; fi17 < 11; fi17++ {
        v11 = v11
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: v11, F1: 967}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    if true {
        var v20 bool = true
        var v21 uint32 = 42
        switch (("hello" + "a") + "") {
        case "a":
            v12 = append(v12, -2147483648)
        }
        _ = v20
        _ = v21
    } else {
        if false {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            var v22 int8 = 1
            var v23 int = int(v11)
            _ = v22
            _ = v23
        }
        switch 454 {
        case 22:
            if (42 != 0) {
                var v24 uint = (v11 << ((v11) & 63))
                v25 := GenS1{F0: v24, F1: 2}
                _ = v24
                _ = v25
            } else {
                v11 += (v11 % (((^(v11))) | 1))
                v26 := v3[65535]
                if true {
                    var v27 int64 = -30
                    _ = v27
                } else {
                    var v28 uint8 = uint8(v11)
                    _ = v28
                }
                _ = v26
            }
        default:
            var v29 int = 42
            _ = v29
        }
    }
    var v30 int = 9223372036854775807
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
