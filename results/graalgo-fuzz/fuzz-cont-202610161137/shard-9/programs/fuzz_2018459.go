package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint; F2 int }
func genFn0() int64 {
    var v6 int8 = 0
    fmt.Println(v6)
    v6 = (v6 | (-(v6)))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: "a", F1: ("" + "")}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint8 = 2
            var v11 uint32 = 65535
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 510
}
func genFn1(p0 uint16, p1 uint16) string {
    var v12 uint32 = uint32(p1)
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(p0, p1, p1)
            _ = fi14
        }
        _ = fi13
    }
    v12 = (v12 ^ 1)
    var v15 float32 = 1.5
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return "\n\t"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    var v3 string = ""
    v4 := make(map[int]int)
    v5 := []string{v3}
    var v16 uint16 = 65535
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v3)
        _ = fi17
    }
    v16 += (v16 * v16)
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            if (true && (v16 == v16)) {
                fmt.Println(v16, v3, v16)
                if (-5 <= 42) {
                    fmt.Println(v16, v3)
                    var v20 uint32 = 0
                    _ = v20
                } else {
                    v3 = ((v3 + "hello") + v3)
                    fmt.Println(v3, fi18, v16)
                    v2 = 2147483647
                }
            } else {
                v21 := GenS1{F0: v3, F1: v3}
                _ = v21
            }
            if (0 < -32768) {
                v22 := GenS2{F0: ("a" + v3)}
                _ = v22
            } else {
                var v23 int32 = v2
                _ = v23
            }
            _ = len(v5)
            v16 += v16
            _ = fi19
        }
        v16 = 2
        _ = fi18
    }
    switch 250 {
    case -920:
        v24 := GenS2{F0: ((v3 + "\n\t") + v3)}
        _ = v24
    case 1:
        v2 += v2
    case 2147483647:
        if true {
            var v25 uint16 = v16
            v25 -= (v25 % ((580) | 1))
            _ = v25
        } else {
            v26 := GenS1{F0: "a", F1: "\n\t"}
            var v27 int32 = v2
            v28 := GenS2{F0: v3}
            _ = v26
            _ = v27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        v16 = 0
        v3 = (v3 + (v3 + v3))
        _ = fi29
    }
    var v30 int16 = 1000
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS3{F0: (v2 % (((v2 - 2147483647)) | 1)), F1: 65535, F2: (fi31 * (-(fi31)))}
        _ = fi31
        _ = v32
    }
    _ = len(v1)
    v30 += v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
