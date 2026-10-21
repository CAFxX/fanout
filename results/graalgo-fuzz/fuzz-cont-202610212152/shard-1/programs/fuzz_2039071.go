package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 997
    v3 := make(map[string]uint32)
    v4 := []int{32768, -628, 32767}
    v5 := []uint32{860, 2147483647}
    v6 := []string{"hello", ("foo bar" + (("hello" + "a") + "foo bar"))}
    fmt.Println(v2)
    var v7 bool = true
    if v7 {
        v8 := GenS1{F0: v7}
        v7 = v7
        v9 := GenS1{F0: true}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: v7}
    v11 := GenS1{F0: v7}
    var v12 uint32 = 2147483648
    v7 = ((v12 < (v12 << ((v2) & 63))) && (297 == 0))
    if v7 {
        v7 = (!(1 != 127))
        for fi13 := 0; fi13 < 16; fi13++ {
            switch (fi13 ^ 4294967295) {
            case -100:
                for fi14 := 0; fi14 < 10; fi14++ {
                    delete(v1, 42)
                    if (592 == 115) {
                        if len(v4) > 0 {
                            v4[0] = (fi14 / ((42) | 1))
                        }
                    }
                    v15 := GenS1{F0: true}
                    _ = fi14
                    _ = v15
                }
            case 4294967295:
                var v16 uint16 = uint16(v2)
                _ = v16
            default:
                for fi17 := 0; fi17 < 8; fi17++ {
                    delete(v1, fi13)
                    var v18 float32 = 0.1
                    _ = fi17
                    _ = v18
                }
            }
            var v19 uint8 = 74
            fmt.Println(v2, v7)
            _ = fi13
            _ = v19
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS2{F0: 100, F1: "hello"}
            v22 := GenS3{F0: -2}
            v7 = (911 == 404)
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if v7 {
        v12 -= 1
        switch "\n\t" {
        case "a":
            fmt.Println(v2)
        }
    }
    var v23 uint8 = 255
    var v24 float32 = ((1.5 * ((0.0 * 0.0) - ((1.5 / 1.0) - 1e10))) - 1.5)
    fmt.Println(v23, v7)
    v23 = (^(v23))
    var v25 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
