package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool }

func main() {
    var v1 bool = true
    v2 := []uint32{2, 4294967295}
    v3 := []uint32{2, 65535}
    v4 := []int{897, -990}
    var v5 uint8 = 0
    v6 := GenS1{F0: -128}
    v1 = false
    v7 := GenS2{F0: 18446744073709551615}
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 int8 = -1
    var v10 uint = 99
    var v11 int32 = -596
    var v12 uint16 = 42
    switch (("foo bar" + ("a" + "hello")) + ("" + ("" + "\n\t"))) {
    case "\n\t":
        v13 := GenS1{F0: (v9 * v9)}
        _ = v13
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "a":
        _ = len(v4)
    }
    var v14 uint8 = v5
    v5 = (v5 << (((v10 ^ v10)) & 63))
    v15 := GenS2{F0: 984}
    var v16 string = ("hello" + "a")
    if (1 != v12) {
        for fi17 := 0; fi17 < 5; fi17++ {
            if ((!((v11 ^ 42) != 42)) || false) {
                if v1 {
                    if (v1 && false) {
                        v12 += (v12 / ((v12) | 1))
                        switch "a" {
                        case "":
                            fmt.Println(v10, v16, fi17)
                        case "a":
                            if (!(v1 || v1)) {
                                var v18 uint32 = 42
                                if (v1 && true) {
                                    if true {
                                        if (95 == 0) {
                                            v19 := GenS2{F0: 0}
                                            v14 = (v14 / ((uint8(v10)) | 1))
                                            _ = v19
                                        }
                                        fmt.Println(v12, v18)
                                        var v20 int16 = -1
                                        _ = v20
                                    }
                                    if (v1 && true) {
                                        var v21 uint16 = (v12 << ((v10) & 63))
                                        _ = v21
                                    } else {
                                        v1 = v1
                                        v1 = v1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    v22 := GenS3{F0: (-(v9)), F1: v16}
                                    _ = v22
                                } else {
                                    var v23 int = (fi17 * int(v18))
                                    _ = v23
                                }
                                var v24 bool = true
                                _ = v18
                                _ = v24
                            } else {
                                switch (v16 + "foo bar") {
                                case "":
                                    if ((fi17 + (fi17 - fi17)) != 708) {
                                        v25 := GenS4{F0: v11, F1: v9}
                                        _ = v25
                                    } else {
                                        fmt.Println(fi17, v5, v9)
                                    }
                                case "a":
                                    fmt.Println(v5, v10)
                                }
                                switch v16 {
                                case "\n\t":
                                    if (false && false) {
                                        v10 -= uint(v5)
                                        v4 = append(v4, (fi17 / ((32768) | 1)))
                                        var v26 uint64 = 584
                                        _ = v26
                                    }
                                case "foo bar":
                                    var v27 int = (fi17 >> (((v10 * v10)) & 63))
                                    _ = v27
                                case "":
                                    var v28 uint16 = (v12 >> (((^(v10))) & 63))
                                    _ = v28
                                }
                            }
                        case "hello":
                            if (1e10 != ((-2.25 / (float32(3.14159) - 1e10)) - 0.0)) {
                                fmt.Println(fi17, v1)
                                switch v16 {
                                case "a":
                                    v29 := GenS3{F0: (v9 >> ((2) & 63)), F1: v16}
                                    _ = v29
                                default:
                                    v1 = v1
                                }
                                fmt.Println(v10, fi17)
                            }
                        default:
                            v16 = ((("foo bar" + (v16 + "hello")) + "\n\t") + (("a" + v16) + "a"))
                        }
                    }
                    v10 -= 0
                    if v1 {
                        _ = len(v3)
                        var v30 string = (v16 + "foo bar")
                        v31 := GenS2{F0: 4294967295}
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 float32 = 3.14159
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v32
                }
                v33 := GenS5{F0: v1}
                _ = v33
            }
            var v34 float32 = 0.0
            switch (fi17 / ((127) | 1)) {
            case 42:
                v10 = 4294967295
            case 0:
                var v35 uint64 = 4294967295
                _ = v35
            case 32768:
                var v36 string = "hello"
                _ = v36
            }
            _ = fi17
            _ = v34
        }
    }
    v9 += v9
    v5 += (v14 >> (((v10 ^ v10)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
