package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 int64; F2 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 float32 = 0.0
    var v2 float32 = -2.25
    var v3 uint64 = 2
    var v4 int16 = 1000
    var v5 uint = 317
    v6 := GenS1{F0: 0, F1: false}
    fmt.Println(v1, v2, v2)
    fmt.Println(v4, v2)
    var v7 int16 = int16(v3)
    v8 := []int{49, -292, -2147483648, 1}
    var v9 int = -9223372036854775808
    if (v4 <= v7) {
        switch "hello" {
        case "a":
            switch (v9 / (((^(v9))) | 1)) {
            case -2147483648:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case 32768:
                v4 = (v7 * v4)
            case 4294967295:
                v2 = (1.5 - float32((0.0 * 1e10)))
            }
        case "foo bar":
            v3 = (v3 | 2)
        }
    } else {
        v5 = 65535
        if (!(114 < 255)) {
            v10 := GenS1{F0: 661, F1: (9223372036854775807 == -1)}
            var v11 string = "\n\t"
            if true {
                v3 = v3
                v3 -= v3
                var v12 int64 = -1
                _ = v12
            }
            _ = v10
            _ = v11
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            if (false || (1 >= 100)) {
                if len(v8) > 0 {
                    v8[0] = (fi13 * (^(v9)))
                }
                switch (fi13 + (fi13 - (v9 ^ v9))) {
                case 42:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(v7, v3, v2)
                        _ = fi14
                    }
                default:
                    switch 128 {
                    case 128:
                        var v15 int32 = 225
                        _ = v15
                    case 32768:
                        fmt.Println(v2, v5, fi13)
                    }
                }
            } else {
                v4 -= (v7 >> (((^(v5))) & 63))
                if len(v8) > 0 {
                    v8[0] = 128
                }
                var v16 uint8 = 0
                _ = v16
            }
            _ = fi13
        }
    }
    fmt.Println(v9)
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            if true {
                v2 -= float32(((3.14159 / (3.14159 * -2.25)) / ((0.0 - (1e10 / -2.25)) * 1.5)))
                v9 -= fi18
            } else {
                var v19 uint32 = 2147483647
                if false {
                    fmt.Println(v7)
                }
                _ = v19
            }
            v20 := GenS2{F0: 910, F1: 9223372036854775807, F2: -249}
            var v21 string = ("\n\t" + (("a" + ((("hello" + ("\n\t" + "")) + "a") + "hello")) + "hello"))
            fmt.Println(v21, fi17)
            _ = fi18
            _ = v20
            _ = v21
        }
        v3 = 42
        var v22 uint8 = 2
        _ = fi17
        _ = v22
    }
    v7 -= v4
    v3 = v3
    v5 = (v5 - 2)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if ("foo bar" == "a") {
        v2 += float32(0.0)
        v23 := GenS3{F0: true}
        fmt.Println(v5, v2)
        _ = v23
    }
    v4 -= (^(v7))
    fmt.Println(v9)
    v8 = append(v8, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
