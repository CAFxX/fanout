package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 uint8; F1 int; F2 string }
func genFn0(p0 int8) uint16 {
    if (true || true) {
        p0 += 42
    }
    var v7 float32 = 0.0
    var v8 int64 = 0
    var v9 uint64 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]string)
    v3 := []uint32{581, 2147483647, 4294967295}
    var v4 uint8 = 0
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    fmt.Println(v4)
    v10 := GenS1{F0: 42, F1: 410, F2: 128}
    var v11 string = ("foo bar" + "foo bar")
    var v12 uint32 = 4294967295
    v13 := GenS1{F0: 1, F1: 1, F2: -1}
    if len(v3) > 0 {
        v3[0] = (v12 >> ((569) & 63))
    }
    v14 := GenS1{F0: 171, F1: 2, F2: 1}
    var v15 uint = 42
    v1 = ((v1 + (v1 - 3.14159)) + (float32(1.5) + (v1 + (v1 / float32((1.5 + (float64(((0.0 * (1e10 * 0.1)) * (-2.25 - 1.5))) + 1.5)))))))
    v15 -= 746
    switch "a" {
    case "\n\t":
        var v16 uint = 2
        _ = v16
    case "a":
        fmt.Println(v15)
    case "foo bar":
        if (42 > -1) {
            v4 = 244
        }
    default:
        if false {
            if (v4 >= (v4 % (((v4 & v4)) | 1))) {
                var v17 uint64 = 2
                _ = len(v3)
                v17 = (^(v17))
                _ = v17
            } else {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 int8 = int8(v12)
                    fmt.Println(v15)
                    _ = fi18
                    _ = v19
                }
            }
            v20 := GenS1{F0: 42, F1: 489, F2: 127}
            var v21 float64 = 1e10
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS1{F0: -248, F1: 1, F2: fi22}
                if len(v3) > 0 {
                    v3[0] = v12
                }
                for fi24 := 0; fi24 < 19; fi24++ {
                    if false {
                        if false {
                            fmt.Println(v12, v4, v1)
                        } else {
                            v25 := GenS1{F0: 42, F1: (v15 / (((v15 % ((v15) | 1))) | 1)), F2: 32768}
                            _ = v25
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = v12
                        }
                        v15 -= 58
                        v11 = "hello"
                    }
                    _ = fi24
                }
                _ = fi22
                _ = v23
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                switch ("a" + (v11 + (v11 + (v11 + v11)))) {
                case "a":
                    switch ("foo bar" + ((v11 + ((v11 + v11) + "foo bar")) + v11)) {
                    case "":
                        v3 = append(v3, v12)
                    default:
                        fmt.Println(fi26, v12)
                    }
                case "":
                    var v27 uint8 = (^(v4))
                    _ = v27
                case "hello":
                    v28 := GenS2{F0: uint8(v15), F1: (-(fi26)), F2: v11}
                    _ = v28
                }
                _ = fi26
            }
        }
    }
    var v29 int8 = 42
    var v30 string = v11
    v31 := GenS1{F0: -34, F1: (v15 % ((v15) | 1)), F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
