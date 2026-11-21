package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 int32, p1 bool, p2 int16) int32 {
    v4 := []int{2147483647}
    var v5 int16 = p2
    _ = len(v4)
    if (42 >= 18446744073709551615) {
        if p1 {
            var v6 int = -896
            _ = v6
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 42, F1: fi7, F2: -1}
            _ = fi7
            _ = v8
        }
    }
    p2 -= 598
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := []string{}
    var v2 int64 = 35
    var v3 int32 = -792
    var v9 int8 = 42
    v10 := GenS1{F0: 0, F1: 839, F2: (v9 ^ v9)}
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: 2, F1: fi11, F2: (v9 << ((255) & 63))}
        _ = fi11
        _ = v12
    }
    _ = len(v1)
    var v13 int8 = -109
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v2, v2)
        _ = fi14
    }
    var v15 bool = true
    if (128 <= 0) {
        v9 += (v9 << ((65535) & 63))
        v1 = append(v1, "foo bar")
        v16 := GenS1{F0: 196, F1: -1, F2: -2}
        _ = v16
    } else {
        if (!(!(1 <= 746))) {
            var v17 uint = 1
            fmt.Println(v9, v15, v17)
            for fi18 := 0; fi18 < 16; fi18++ {
                fmt.Println(v15)
                switch ("a" + "") {
                case "\n\t":
                    v19 := GenS1{F0: 42, F1: fi18, F2: v13}
                    _ = v19
                }
                _ = fi18
            }
            _ = v17
        } else {
            var v20 int64 = (v2 << ((116) & 63))
            _ = v20
        }
        var v21 string = ("a" + "")
        _ = v21
    }
    v22 := GenS1{F0: 959, F1: 127, F2: v9}
    var v23 uint32 = 551
    if (v23 <= (v23 * v23)) {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        v23 = 2
    } else {
        v3 -= int32(v23)
        v24 := GenS1{F0: 32767, F1: 678, F2: 42}
        _ = v24
    }
    fmt.Println(v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
