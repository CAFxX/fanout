package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32) uint16 {
    p0 = p0
    p0 = p0
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 9223372036854775807}
        switch (fi7 << ((42) & 63)) {
        case 2147483647:
            var v9 uint16 = 32767
            _ = v9
        case -259:
            if (42 != 247) {
                var v10 int8 = 42
                fmt.Println(p0, fi7)
                _ = v10
            }
        case 32767:
            v11 := GenS1{F0: 9223372036854775807}
            _ = v11
        }
        v12 := []int{87, (-(fi7)), (fi7 / ((fi7) | 1))}
        _ = fi7
        _ = v12
        _ = v8
    }
    _ = p0
    return 968
}
func genFn1(p0 string, p1 uint8) bool {
    p1 = (p1 * (p1 & 255))
    var v13 uint = 255
    v14 := GenS1{F0: 0}
    v15 := []int{}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return (true || true)
}

func main() {
    var v1 int32 = -795
    var v2 int16 = 1
    v3 := make(map[int]bool)
    var v4 uint64 = 1
    var v5 int64 = 126
    v6 := make(map[int]int)
    v4 -= 262
    v16 := []int{9223372036854775807, -315, -9223372036854775808, -1}
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint32 = 65535
        _ = fi17
        _ = v18
    }
    delete(v6, -2147483648)
    v19 := v3[1]
    fmt.Println(v5)
    if false {
        var v20 uint16 = genFn0(207)
        _ = v20
    } else {
        _ = len(v16)
        var v21 string = ("hello" + ("hello" + ("foo bar" + "foo bar")))
        v5 = int64(v2)
        _ = v21
    }
    if len(v16) > 0 {
        v16[0] = 1
    }
    v22 := GenS2{F0: v19}
    if (!(!(-2.25 > 0.1))) {
        switch 256 {
        case 128:
            if false {
                var v23 float32 = float32(1.5)
                v24 := GenS2{F0: (v5 > (v5 & -1))}
                _ = v23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v1 += (v1 << ((42) & 63))
            var v26 float32 = 1e10
            for fi27 := 0; fi27 < 7; fi27++ {
                v5 += (v5 % ((v5) | 1))
                var v28 int16 = (v2 | (v2 / (((v2 * (v2 << ((65535) & 63)))) | 1)))
                fmt.Println(v1, fi27, v1)
                v28 = -111
                fmt.Println(v28, v26, v2)
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    }
    var v29 int64 = (v5 >> ((241) & 63))
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
