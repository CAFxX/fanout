package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, (fi5 * fi5), fi5, fi5}
        fmt.Println(p0, fi5)
        v7 := GenS1{F0: 0}
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = (p0 / ((p0) | 1))
    var v8 bool = true
    switch ("a" + (((("" + ("foo bar" + "foo bar")) + "") + "\n\t") + "hello")) {
    case "foo bar":
        if (128 == 2) {
            v8 = v8
            var v9 int = 42
            v8 = v8
            _ = v9
        }
    case "a":
        fmt.Println(v8, v8)
    }
    _ = p0
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 264
    var v3 uint32 = 475
    v4 := make(map[string]int)
    var v10 bool = true
    fmt.Println(v2, v3, v3)
    if (v3 >= (v3 + v3)) {
        if true {
            var v11 int64 = 1
            if (-1 > 1) {
                var v12 float32 = 1.5
                v13 := GenS2{F0: 2}
                v2 = v2
                _ = v12
                _ = v13
            }
            v11 -= v11
            _ = v11
        }
    }
    fmt.Println(v2, v3, v3)
    if ((v2 - (v2 ^ 65535)) < 426) {
        v10 = (-1 >= -6)
        v2 = v2
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            if (42 >= 0) {
                fmt.Println(fi14, v3, v3)
                if v10 {
                    fmt.Println(v2)
                }
            }
            switch (fi14 << ((305) & 63)) {
            case 128:
                if (!(109 == 0)) {
                    var v15 uint8 = 180
                    _ = v15
                }
            case -9223372036854775808:
                fmt.Println(v2)
            }
            _ = fi14
        }
    }
    v10 = false
    v16 := []int{-580}
    var v17 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
