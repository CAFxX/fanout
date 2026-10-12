package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 bool, p1 bool, p2 int64) int {
    var v6 int16 = 42
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v7 int16 = 410
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 1
    var v3 uint32 = 2
    var v4 uint32 = (v3 / (((v1 >> ((58) & 63))) | 1))
    var v5 bool = false
    var v8 uint16 = 65535
    var v9 int16 = int16(v4)
    if (1 != 252) {
        v10 := []int{1, genFn0(((1.5 > float32(3.14159)) || v5), true, 0), 0, 127}
        _ = v10
    }
    v11 := []int{}
    var v12 int64 = 1
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 17; fi14++ {
        switch "a" {
        case "foo bar":
            v8 += (v8 % ((65535) | 1))
        case "a":
            var v15 bool = (!(42 < 30))
            _ = v15
        }
        if v5 {
            v16 := GenS1{F0: genFn0(v5, v5, v12), F1: -242}
            v11 = append(v11, (fi14 % ((genFn0((246 != 255), false, -291)) | 1)))
            v2 += (v2 & v2)
            _ = v16
        } else {
            var v17 float64 = -2.25
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int64 = v12
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = fi14
    }
    if true {
        var v20 float64 = v13
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            v1 = v3
            if len(v11) > 0 {
                v11[0] = fi21
            }
            v11 = append(v11, genFn0((0.1 == 0.1), (0 < -1), v12))
            _ = fi21
        }
        v22 := GenS1{F0: 32768, F1: (v12 - (v12 << ((616) & 63)))}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
