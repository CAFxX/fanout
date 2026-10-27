package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 2
    var v3 string = (("hello" + ((("hello" + "a") + ((("" + ("\n\t" + (("\n\t" + ("foo bar" + "a")) + "foo bar"))) + (("a" + "") + "hello")) + "\n\t")) + "\n\t")) + ("foo bar" + "hello"))
    var v4 int64 = 1
    v5 := make(map[uint32]bool)
    switch -2147483648 {
    case -9223372036854775808:
        v3 = v3
    case -512:
        delete(v1, -197)
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v2)
        _ = fi6
    }
    var v7 int = 128
    v8 := []int{int(v4), v7}
    if true {
        if (true || true) {
            if true {
                fmt.Println(v3)
                v9 := GenS1{F0: 2}
                _ = v9
            } else {
                v2 -= 1
                v2 = (^(v2))
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                v7 += (fi10 << ((19) & 63))
                v2 -= v2
                _ = fi10
            }
            var v11 uint32 = 42
            _ = v11
        }
        switch "a" {
        case "hello":
            v4 += 42
        case "foo bar":
            v7 -= v7
        case "a":
            var v12 float32 = float32(1.5)
            _ = v12
        default:
            for fi13 := 0; fi13 < 14; fi13++ {
                v7 = v7
                if false {
                    if (uint32(v2) != 852) {
                        if ((!((v7 / (((fi13 & 9223372036854775807)) | 1)) >= v7)) || (int(v2) >= (fi13 >> ((0) & 63)))) {
                            v2 = 1
                            v7 = v7
                            fmt.Println(v7)
                        } else {
                            var v14 int32 = 2147483647
                            var v15 uint64 = 0
                            _ = v14
                            _ = v15
                        }
                        v3 = "a"
                        v16 := GenS1{F0: 2}
                        _ = v16
                    }
                    var v17 uint8 = 127
                    _ = v17
                } else {
                    var v18 uint32 = 42
                    _ = v18
                }
                _ = fi13
            }
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            switch (v3 + v3) {
            case "foo bar":
                if len(v8) > 0 {
                    v8[0] = (v7 & v7)
                }
            case "a":
                v2 -= 255
            }
            _ = fi19
        }
    }
    if true {
        var v20 int64 = (-(v4))
        v20 = v4
        var v21 int64 = -9223372036854775808
        _ = v20
        _ = v21
    }
    switch (v7 + v7) {
    case 9223372036854775807:
        v7 += (v7 ^ 42)
    case 996:
        var v22 int16 = 42
        _ = v22
    default:
        v23 := GenS2{F0: (v2 ^ 886), F1: 100}
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if ((int8(v2) < -128) || (-380 != v4)) {
            var v25 uint64 = 0
            _ = v25
        }
        _ = fi24
    }
    var v26 float32 = (3.14159 * 0.0)
    if true {
        v5[2] = false
    } else {
        v4 -= (v4 << ((4294967295) & 63))
        var v27 int8 = -1
        v28 := GenS1{F0: 0}
        _ = v27
        _ = v28
    }
    fmt.Println(v26)
    v7 = (^(v7))
    var v29 uint32 = 2
    for fi30 := 0; fi30 < 7; fi30++ {
        if ((v3 + (v3 + v3)) != v3) {
            if true {
                var v31 int32 = -2147483648
                for fi32 := 0; fi32 < 4; fi32++ {
                    fmt.Println(v29)
                    v33 := GenS1{F0: 255}
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
