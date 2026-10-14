package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint16, p2 bool) uint8 {
    fmt.Println(p0, p2, p1)
    v4 := GenS1{F0: p1, F1: false}
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v5 := GenS2{F0: 0, F1: 323, F2: -528}
        _ = v5
    case "hello":
        var v6 int32 = 0
        _ = v6
    }
    v7 := GenS1{F0: p1, F1: ((!p2) && (698 != 0))}
    v8 := GenS2{F0: 2, F1: 416, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return 86
}
func genFn1(p0 uint, p1 uint, p2 bool) int16 {
    v9 := GenS3{F0: 18}
    fmt.Println(p1, p2)
    var v10 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -32768
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 3.14159
    var v3 float32 = (((v2 * (3.14159 * v2)) / v2) * (1.5 / (1.5 - ((((v2 + v2) - (3.14159 * 1e10)) * v2) / v2))))
    fmt.Println(v3, v3, v1)
    v1 -= v1
    var v11 int8 = 0
    fmt.Println(v3, v3)
    fmt.Println(v11, v11)
    var v12 float64 = 0.0
    fmt.Println(v11)
    switch 1 {
    case -125:
        v12 = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v2)
        v3 -= (v3 - v3)
        switch "a" {
        case "foo bar":
            switch (^(fi13)) {
            case 1:
                var v14 int16 = -209
                _ = v14
            }
        }
        _ = fi13
    }
    v3 += (0.0 + float32(((((float64(v2) / v12) + (v12 - v12)) + (v12 + v12)) / (1.5 + (v12 * v12)))))
    for fi15 := 0; fi15 < 2; fi15++ {
        if (!true) {
            v16 := []int{(fi15 ^ (fi15 + (fi15 >> ((65535) & 63)))), fi15, fi15, (fi15 * (fi15 << ((42) & 63)))}
            var v17 bool = true
            _ = v16
            _ = v17
        } else {
            var v18 float32 = v2
            v1 = 65535
            var v19 int = fi15
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
