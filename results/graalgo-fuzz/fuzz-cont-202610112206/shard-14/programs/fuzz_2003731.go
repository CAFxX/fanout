package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    v3 := []string{"a", "", "hello"}
    var v4 int32 = 885
    v5 := GenS1{F0: 1, F1: 32}
    var v6 float32 = 0.0
    var v7 uint16 = 32767
    var v8 bool = false
    v7 += 565
    var v9 int8 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: (v7 | (v7 * v7)), F1: 65535}
        v8 = v8
        fmt.Println(v7, v2)
        _ = fi10
        _ = v11
    }
    if (360 != 2) {
        if v8 {
            switch 9223372036854775807 {
            case 127:
                switch ("" + "\n\t") {
                case "":
                    var v12 float32 = float32(1.5)
                    _ = v12
                }
            case -835:
                v6 = 1.5
            case 803:
                for fi13 := 0; fi13 < 10; fi13++ {
                    v8 = (!(!(42 == 4294967295)))
                    if ((v7 * 32767) > (^(v7))) {
                        for fi14 := 0; fi14 < 15; fi14++ {
                            if v8 {
                                if true {
                                    if (0 == 1) {
                                        v15 := GenS2{F0: v8}
                                        v9 = (v9 >> ((42) & 63))
                                        _ = v15
                                    } else {
                                        var v16 int32 = int32(v9)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v9 -= v9
                                        var v17 uint8 = 81
                                        if false {
                                            v9 -= v9
                                            v18 := GenS3{F0: v9, F1: 208, F2: (v16 ^ (v16 ^ (v16 ^ v16)))}
                                            v8 = (("a" + ("a" + ("hello" + "\n\t"))) < ("a" + ""))
                                            _ = v18
                                        } else {
                                            var v19 int = (-(fi14))
                                            _ = v19
                                        }
                                        _ = v16
                                        _ = v17
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                    var v20 uint8 = 127
                                    _ = v20
                                } else {
                                    v6 = (3.14159 / (v6 + 1.5))
                                    fmt.Println(v7)
                                    var v21 uint64 = 0
                                    _ = v21
                                }
                            }
                            _ = fi14
                        }
                        switch ("hello" + (("" + "\n\t") + "hello")) {
                        case "hello":
                            var v22 uint = 0
                            _ = v22
                        default:
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi13, v6)
                            case "hello":
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    v8 = (true && v8)
                                    fmt.Println(v6, v8)
                                    _ = fi23
                                }
                            case "a":
                                fmt.Println(v8, v4)
                            default:
                                fmt.Println(v4, v8)
                            }
                        }
                    }
                    _ = fi13
                }
            default:
                v9 += (v9 / (((^(v9))) | 1))
            }
            var v24 bool = (true || (int16(v9) > (v2 + (-(v2)))))
            var v25 int8 = (v9 + v9)
            _ = v24
            _ = v25
        } else {
            v7 += v7
        }
        v26 := GenS4{F0: 127}
        _ = v26
    } else {
        v9 += (v9 << ((18446744073709551615) & 63))
        v27 := GenS5{F0: 1}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
