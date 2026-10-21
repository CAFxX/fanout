package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int8, p1 int) uint16 {
    switch (p1 + 2147483647) {
    case 128:
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int64 = 510
            var v8 int = (fi6 >> ((924) & 63))
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        fmt.Println(p1, p1)
    }
    v9 := []int{(p1 - -759), (p1 >> ((42) & 63))}
    v10 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}
func genFn1() bool {
    v11 := GenS1{F0: false}
    var v12 int32 = 2147483647
    v13 := GenS1{F0: (!(152 == 236))}
    v14 := []int{}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return ((!(1 != 2147483648)) && (464 >= 2))
}

func main() {
    var v1 string = ""
    v2 := []uint32{2147483647, 569}
    var v3 int = -185
    v4 := make(map[uint32]string)
    v5 := []string{"hello", v1}
    v3 -= v3
    if ((!(!(-308 > -502))) && (0 == 9223372036854775807)) {
        v2 = append(v2, 2147483647)
    }
    v15 := GenS2{F0: v3}
    for fi16 := 0; fi16 < 13; fi16++ {
        v5 = append(v5, "\n\t")
        v3 = -9223372036854775808
        if true {
            fmt.Println(v1, v3)
            var v17 uint32 = 186
            v18 := GenS3{F0: v1, F1: 32767, F2: 945}
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    if (!true) {
        switch (v3 ^ v3) {
        case 9223372036854775807:
            switch v3 {
            case 32767:
                v19 := GenS2{F0: (v3 + v3)}
                _ = v19
            default:
                fmt.Println(v1, v3)
            }
        default:
            fmt.Println(v3, v3)
        }
        var v20 int8 = 1
        _ = v20
    } else {
        _ = len(v2)
        v1 = (v1 + v1)
    }
    var v21 int32 = 1
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v1, fi22, v1)
        for fi23 := 0; fi23 < 15; fi23++ {
            v3 = (v3 & 915)
            v24 := GenS4{F0: 1, F1: (0 >= 0)}
            if (((3.14159 - 1.5) + -2.25) == ((-2.25 * 1e10) * float32(0.0))) {
                var v25 int16 = 0
                v21 = 0
                _ = v25
            } else {
                var v26 int16 = 32767
                v27 := GenS3{F0: v1, F1: (v26 * int16(v21)), F2: 42}
                _ = v26
                _ = v27
            }
            fmt.Println(v21, v3, v3)
            switch -611 {
            case -854:
                switch v1 {
                case "hello":
                    if (-32768 == 1) {
                        var v28 uint64 = 42
                        _ = v28
                    }
                case "":
                    switch v1 {
                    case "\n\t":
                        var v29 float32 = (((1.5 * (1e10 * 1e10)) - 1e10) / ((((0.1 * (1.5 / -2.25)) / (0.1 * (1e10 + 1.5))) + 3.14159) / -2.25))
                        v21 += v21
                        _ = v29
                    }
                }
            case 2147483647:
                var v30 float32 = (3.14159 + 1e10)
                _ = v30
            case -1:
                var v31 bool = true
                _ = v31
            }
            _ = fi23
            _ = v24
        }
        var v32 bool = (!(!true))
        _ = fi22
        _ = v32
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
