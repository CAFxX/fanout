package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 bool, p1 bool) string {
    var v7 bool = p1
    p1 = true
    v8 := []int{-2147483648, 9223372036854775807}
    p0 = (255 < 123)
    var v9 string = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1() int64 {
    var v10 float32 = (1.5 + ((0.1 + (1.5 / -2.25)) - (0.1 + 1.5)))
    fmt.Println(v10)
    _ = v10
    return -1
}

func main() {
    var v1 int8 = -14
    var v2 uint64 = 4294967295
    v3 := make(map[string]bool)
    var v4 int8 = v1
    var v5 int = 715
    var v6 int32 = 883
    v11 := v3[("hello" + "foo bar")]
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            switch "foo bar" {
            case "a":
                fmt.Println(v6, v4, v6)
            case "foo bar":
                var v14 int16 = 0
                _ = v14
            }
            _ = fi13
        }
        _ = fi12
    }
    v15 := []int{}
    fmt.Println(v11, v2, v5)
    fmt.Println(v11)
    _ = len(v3)
    if (389 >= 65535) {
        if (v11 || (3.14159 != (0.1 + 1.5))) {
            if (!v11) {
                v5 = (v5 >> ((252) & 63))
                v16 := GenS1{F0: 1, F1: 2, F2: 0}
                fmt.Println(v4)
                _ = v16
            }
        }
        switch (v5 & (^(v5))) {
        case 42:
            fmt.Println(v1)
        case 0:
            var v17 int8 = (v1 - v4)
            _ = v17
        case -345:
            fmt.Println(v11, v4, v11)
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 int = (fi18 >> ((0) & 63))
                _ = fi18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint32 = 1
            if (genFn1() == 0) {
                if v11 {
                    delete(v3, "a")
                    if len(v15) > 0 {
                        v15[0] = (^(v5))
                    }
                    var v22 int8 = (^(v4))
                    _ = v22
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
            var v23 uint8 = 208
            _ = fi20
            _ = v21
            _ = v23
        }
    }
    if (((127 == v4) || v11) && true) {
        var v24 int64 = 0
        _ = v24
    } else {
        var v25 uint16 = 32767
        v25 += (v25 ^ 164)
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS1{F0: v25, F1: 32767, F2: (v6 << ((458) & 63))}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v28 uint = 1
    var v29 int32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
