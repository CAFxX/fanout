package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0() uint32 {
    v5 := GenS1{F0: 1000, F1: 42}
    var v6 string = ("\n\t" + "hello")
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = -128
    var v3 int16 = 1000
    var v4 string = "foo bar"
    v7 := []int{32768}
    var v8 string = v4
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint32 = 2147483647
        _ = fi9
        _ = v10
    }
    fmt.Println(v8)
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v4, fi11)
        if len(v7) > 0 {
            v7[0] = int(v3)
        }
        var v12 int32 = 2147483647
        _ = fi11
        _ = v12
    }
    if ((v2 >= v2) || false) {
        var v13 bool = false
        switch (((((v8 + v4) + "") + (("hello" + v4) + v4)) + "hello") + ((v8 + v4) + "a")) {
        case "hello":
            v14 := GenS1{F0: -131, F1: genFn0()}
            _ = v14
        case "a":
            if (2 < 2) {
                _ = len(v7)
                var v15 int64 = -1
                v16 := GenS2{F0: v1, F1: 255}
                _ = v15
                _ = v16
            }
        case "":
            var v17 int32 = int32(v3)
            _ = v17
        default:
            if v13 {
                v18 := GenS1{F0: 1, F1: 140}
                if (-1 != v2) {
                    v3 = v3
                    v19 := GenS1{F0: (v3 / (((v3 | v3)) | 1)), F1: 0}
                    _ = v19
                } else {
                    v4 = (v4 + v8)
                    var v20 uint16 = (v1 % (((v1 - v1)) | 1))
                    if true {
                        if (uint8(v1) > 128) {
                            var v21 int64 = 0
                            if (1.5 <= float32(0.0)) {
                                switch -2147483648 {
                                case 38:
                                    var v22 int8 = v2
                                    _ = v22
                                case 32768:
                                    var v23 bool = v13
                                    _ = v23
                                case 4294967295:
                                    v13 = ((127 >= 255) && (108 != 1))
                                }
                            }
                            _ = v21
                        }
                    }
                    _ = v20
                }
                _ = v18
            } else {
                v2 -= v2
            }
        }
        var v24 string = ((v4 + "\n\t") + "foo bar")
        _ = v13
        _ = v24
    } else {
        v1 -= (v1 >> ((2) & 63))
    }
    v25 := GenS2{F0: (v1 << ((2) & 63)), F1: 42}
    for fi26 := 0; fi26 < 3; fi26++ {
        v1 += (v1 & 919)
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int = (fi27 ^ (fi27 + 9223372036854775807))
            var v29 int32 = 2147483647
            fmt.Println(v4, v2)
            _ = fi27
            _ = v28
            _ = v29
        }
        if true {
            var v30 int8 = 100
            v31 := GenS2{F0: v1, F1: 5}
            if (!false) {
                v7 = append(v7, fi26)
                v30 = -8
                v32 := GenS3{F0: v2, F1: (v2 << ((640) & 63)), F2: 531}
                _ = v32
            }
            _ = v30
            _ = v31
        } else {
            v4 = v4
            v8 = v8
        }
        _ = fi26
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 uint = 2
        for fi35 := 0; fi35 < 3; fi35++ {
            if (false || true) {
                var v36 float32 = 0.1
                _ = v36
            }
            var v37 string = (v8 + (v4 + v4))
            _ = fi35
            _ = v37
        }
        v38 := GenS2{F0: v1, F1: 68}
        _ = fi33
        _ = v34
        _ = v38
    }
    v8 = (v8 + v4)
    var v39 bool = false
    v39 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v7
    _ = v8
}
