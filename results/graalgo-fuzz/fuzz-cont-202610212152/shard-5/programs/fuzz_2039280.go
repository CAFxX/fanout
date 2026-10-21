package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int16, p1 int64) uint {
    p0 -= p0
    var v6 uint8 = 2
    _ = p0
    _ = p1
    _ = v6
    return 216
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{614, 2, 42, 993}
    v3 := make(map[uint32]uint32)
    var v4 uint = 197
    var v5 string = ("a" + "")
    var v7 int = 328
    v4 -= (v4 - v4)
    var v8 int64 = -1
    if true {
        v4 -= genFn0(42, (v8 * 42))
    } else {
        v2 = append(v2, 67)
    }
    v9 := GenS1{F0: false, F1: 577}
    if (-1 >= -888) {
        if false {
            fmt.Println(v4, v7)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v8 += v8
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(v7, v5, v4)
                v11 := GenS1{F0: (144 != 0), F1: 0}
                var v12 int16 = 1
                _ = fi10
                _ = v11
                _ = v12
            }
            fmt.Println(v4, v8, v5)
            v13 := GenS2{F0: -1}
            _ = v13
        }
    }
    fmt.Println(v5, v1, v5)
    if false {
        v4 -= (v4 >> ((0) & 63))
        v3[2147483647] = 2147483648
        delete(v3, 170)
    } else {
        if false {
            var v14 uint = (v4 + 1)
            _ = v14
        }
        v15 := GenS3{F0: 1, F1: 36}
        if false {
            v7 -= -1
            switch v1 {
            case "":
                for fi16 := 0; fi16 < 17; fi16++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi16
                }
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            default:
                v17 := GenS3{F0: 850, F1: 0}
                _ = v17
            }
        } else {
            v7 += (-(v7))
            var v18 int = v7
            switch (v18 | (v7 ^ (v18 ^ v18))) {
            case 9223372036854775807:
                v19 := GenS3{F0: 0, F1: 242}
                _ = v19
            default:
                for fi20 := 0; fi20 < 6; fi20++ {
                    if false {
                        v21 := GenS3{F0: 255, F1: 42}
                        var v22 uint64 = 619
                        _ = v21
                        _ = v22
                    }
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        switch ("hello" + "hello") {
                        case "\n\t":
                            v1 = v5
                        case "":
                            v23 := v3[0]
                            _ = v23
                        case "foo bar":
                            if ((v18 & v18) > 32768) {
                                var v24 float32 = 1.5
                                _ = v24
                            }
                        }
                    case "":
                        v4 -= 16
                    }
                    fmt.Println(v7, v18)
                    _ = fi20
                }
            }
            _ = v18
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
