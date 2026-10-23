package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 int32, p1 bool, p2 int32) string {
    v5 := GenS1{F0: 2147483647, F1: 0, F2: 0}
    switch 32767 {
    case 32767:
        v6 := []int{77, 4294967295}
        _ = v6
    case 0:
        var v7 string = "\n\t"
        _ = v7
    case 9223372036854775807:
        if ((-2.25 < 3.14159) || (true || p1)) {
            var v8 bool = false
            _ = v8
        } else {
            var v9 int64 = 42
            if p1 {
                switch 734 {
                case -1:
                    v10 := GenS2{F0: 127}
                    _ = v10
                case -9223372036854775808:
                    v11 := GenS1{F0: 107, F1: 128, F2: 32767}
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch 42 {
                        case 1:
                            for fi13 := 0; fi13 < 14; fi13++ {
                                if (185 > 2) {
                                    p2 = 1
                                    p2 = (p0 * (p0 / ((p2) | 1)))
                                }
                                var v14 float64 = 0.1
                                fmt.Println(v14, p1, p0)
                                _ = fi13
                                _ = v14
                            }
                        case 4294967295:
                            var v15 int64 = v9
                            _ = v15
                        default:
                            var v16 int = int(v9)
                            _ = v16
                        }
                        _ = fi12
                    }
                }
                p0 += (p0 / ((-1) | 1))
            }
            fmt.Println(v9, v9)
            _ = v9
        }
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            if p1 {
                fmt.Println(p1, p1, p1)
            }
            if (127 != 0) {
                v18 := GenS2{F0: (-(fi17))}
                var v19 uint16 = 533
                v20 := GenS3{F0: 62, F1: 1, F2: (p0 ^ p0)}
                _ = v18
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint16 = 32767
                    v23 := []int{fi21, (fi17 * (^(fi17))), fi21, (fi17 + (fi21 + (fi17 << ((977) & 63))))}
                    p1 = p1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi17
        }
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := []int{}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "foo bar"
}
func genFn1() int {
    v26 := GenS2{F0: -2147483648}
    switch "hello" {
    case "hello":
        var v27 uint = 18446744073709551615
        _ = v27
    case "\n\t":
        var v28 uint = 65535
        _ = v28
    }
    _ = v26
    return 32768
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint32 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v29 int8 = 1
    v2 = 1
    var v30 int64 = 9223372036854775807
    var v31 int64 = (v30 | v30)
    var v32 uint64 = 0
    fmt.Println(v2, v29, v29)
    var v33 int32 = 2147483647
    fmt.Println(v31)
    v34 := GenS2{F0: (-(v1))}
    if (!true) {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint32 = v2
            v32 = v32
            _ = fi35
            _ = v36
        }
        v30 += (v31 << ((575) & 63))
    }
    var v37 float64 = 3.14159
    v38 := GenS1{F0: (v2 * (v2 + (v2 >> ((4294967295) & 63)))), F1: 167, F2: 0}
    var v39 int = v1
    var v40 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
