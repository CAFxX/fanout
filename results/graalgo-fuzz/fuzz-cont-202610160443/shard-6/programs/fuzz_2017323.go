package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 1
    v3 := make(map[uint32]string)
    var v4 int64 = (v2 + v2)
    v5 := make(map[string]uint32)
    var v6 int16 = -1
    v7 := []int{2147483647, 32767, 9223372036854775807}
    if true {
        v2 = v2
        var v8 string = (((("\n\t" + "hello") + "\n\t") + (("hello" + (("foo bar" + "") + "hello")) + ("a" + "a"))) + ("\n\t" + "foo bar"))
        _ = v8
    }
    switch "" {
    case "":
        fmt.Println(v4, v4, v6)
    default:
        if false {
            delete(v5, "hello")
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                v6 -= 1000
                _ = fi9
            }
            var v10 uint32 = 0
            v2 = v4
            _ = v10
        }
    }
    var v11 uint8 = 15
    if (!((false && (!true)) && (false || true))) {
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(v11, fi12, v11)
            v5[""] = 793
            if (127 >= 127) {
                delete(v1, 342)
            } else {
                if len(v7) > 0 {
                    v7[0] = 0
                }
                v11 -= (v11 / ((v11) | 1))
                fmt.Println(v4, v2, v2)
            }
            _ = fi12
        }
        var v13 uint = 255
        v14 := GenS1{F0: 2147483648, F1: 42}
        _ = v13
        _ = v14
    } else {
        var v15 uint64 = 1
        var v16 uint8 = (^(v11))
        v17 := GenS2{F0: 1}
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 bool = (0.0 >= -2.25)
            var v21 uint = 2
            v5[""] = 872
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := GenS2{F0: 4294967295}
        for fi23 := 0; fi23 < 19; fi23++ {
            v7 = append(v7, (-(fi23)))
            v11 -= (v11 * v11)
            v24 := GenS3{F0: v11, F1: 255, F2: 65535}
            v11 += (v11 & (^(v11)))
            var v25 uint = 1
            v4 = (v2 << ((684) & 63))
            switch -212 {
            case 2147483647:
                if false {
                    var v26 uint64 = 1
                    fmt.Println(v11, v4, v26)
                    var v27 uint64 = (^(v26))
                    var v28 float64 = 1.5
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v29 uint32 = 726
                        _ = v29
                    case "foo bar":
                        _ = len(v7)
                    }
                    var v30 int = fi18
                    switch "" {
                    case "hello":
                        v31 := GenS3{F0: v11, F1: 2, F2: 2147483647}
                        _ = v31
                    default:
                        var v32 uint8 = 42
                        _ = v32
                    }
                    v33 := GenS3{F0: 116, F1: 42, F2: 817}
                    v3[0] = ("hello" + "foo bar")
                    _ = v30
                    _ = v33
                }
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi18
        _ = v22
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 float32 = ((-2.25 * 1.5) - float32(3.14159))
        _ = fi34
        _ = v35
    }
    v11 += (v11 ^ 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
