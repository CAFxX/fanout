package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int) uint32 {
    var v6 uint32 = 392
    v7 := GenS1{F0: uint64(v6), F1: (!(!(-579 != 9223372036854775807))), F2: 255}
    var v8 int8 = 1
    v9 := []int{p0}
    v8 = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v6 * v6)
}
func genFn1(p0 string, p1 uint16, p2 int32) string {
    var v10 uint64 = 1
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int16 = 1000
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS1{F0: 2, F1: false, F2: 127}
                switch 884 {
                case 127:
                    var v15 int32 = -2147483648
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        if false {
            var v16 float32 = 3.14159
            for fi17 := 0; fi17 < 13; fi17++ {
                switch (p0 + p0) {
                case "hello":
                    v18 := GenS1{F0: 42, F1: false, F2: 255}
                    _ = v18
                }
                v19 := []int{int(p1), -1, -415, fi17}
                _ = fi17
                _ = v19
            }
            _ = v16
        }
        fmt.Println(p2, p0)
    }
    p0 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return ("" + p0)
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{268}
    var v3 uint64 = 0
    var v4 int = 478
    var v5 int32 = v1
    if true {
        fmt.Println(v1, v1)
    } else {
        if (!(670 != 260)) {
            _ = len(v2)
            var v20 float64 = (0.0 * (0.1 / -2.25))
            _ = v20
        } else {
            var v21 int64 = -477
            _ = v21
        }
        v22 := GenS1{F0: (v3 % ((1) | 1)), F1: (!true), F2: 2}
        var v23 int = 1
        _ = v22
        _ = v23
    }
    var v24 uint8 = 127
    if true {
        v24 -= v24
    }
    v2 = append(v2, 42)
    v25 := GenS1{F0: (v3 % ((v3) | 1)), F1: true, F2: 0}
    switch v4 {
    case 952:
        if (false && false) {
            fmt.Println(v3)
            v26 := GenS2{F0: -1, F1: 19}
            v27 := GenS2{F0: -1, F1: -117}
            _ = v26
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 1; fi28++ {
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = len(v2)
            _ = fi28
        }
    }
    var v29 uint = 1
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if (0 != 2147483647) {
        if len(v2) > 0 {
            v2[0] = 626
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (-810 >= 1) {
            if len(v2) > 0 {
                v2[0] = uint32(v3)
            }
            v30 := GenS1{F0: (v3 / ((v3) | 1)), F1: (1e10 >= 0.0), F2: uint8(v3)}
            _ = v30
        }
    } else {
        v2 = append(v2, 65535)
        v31 := GenS2{F0: 1, F1: -27}
        _ = v31
    }
    fmt.Println(v29)
    var v32 int64 = 9223372036854775807
    v2 = append(v2, genFn0((v4 * v4)))
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v3)
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 = 2147483647
            v2 = append(v2, uint32(v29))
            _ = fi34
        }
        var v35 uint16 = 1
        _ = fi33
        _ = v35
    }
    var v36 float32 = (((float32((0.1 + 0.1)) / float32((-2.25 / 1e10))) / (-2.25 * 3.14159)) - -2.25)
    v32 -= v32
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
