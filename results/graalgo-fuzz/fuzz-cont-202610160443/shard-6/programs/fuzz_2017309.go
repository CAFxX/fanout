package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int8; F2 int }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0(p0 int16, p1 bool, p2 uint) int32 {
    var v5 string = "foo bar"
    var v6 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 int32) bool {
    var v7 bool = (42 != 4294967295)
    var v8 int64 = 0
    var v9 int64 = (v8 ^ (v8 >> ((42) & 63)))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (!(v7 || ((-2.25 <= float32((0.0 + float64(float32(((1.5 * 1e10) - (3.14159 * 0.0))))))) || (((false || true) && v7) || (165 >= 255)))))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (((true && ((39 != 75) || (false && true))) && true) || (-2147483648 < -9223372036854775808))
    v3 := []uint32{1}
    var v4 float32 = -2.25
    if (((!false) || (0 == -649)) && v2) {
        var v10 uint16 = 390
        fmt.Println(v2)
        _ = v10
    } else {
        if true {
            if v2 {
                v1[576] = 115
                var v11 string = "hello"
                var v12 float32 = v4
                _ = v11
                _ = v12
            }
            switch ("\n\t" + (("" + "hello") + "\n\t")) {
            case "\n\t":
                v2 = (9223372036854775807 > 209)
            case "hello":
                v4 = (0.1 + v4)
            }
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: 42, F1: 42, F2: 2}
                _ = fi13
                _ = v14
            }
            var v15 int64 = 1
            _ = v15
        }
    }
    if (!(!((v2 || (2147483648 == 2147483648)) || ((61 >= 0) || false)))) {
        var v16 uint16 = 0
        v16 += (v16 << ((18446744073709551615) & 63))
        var v17 uint32 = 1
        _ = v16
        _ = v17
    }
    switch -54 {
    case -9223372036854775808:
        v18 := GenS2{F0: v2, F1: 42, F2: -653}
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v2, v2, fi19)
        var v20 int16 = -1
        _ = fi19
        _ = v20
    }
    switch "hello" {
    case "foo bar":
        var v21 int8 = -1
        _ = v21
    case "\n\t":
        v22 := GenS3{F0: -2147483648, F1: 32767}
        _ = v22
    case "a":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: 1, F1: 65535, F2: 65535}
            if (fi23 == -865) {
                var v25 int64 = 0
                var v26 uint32 = 0
                var v27 uint8 = 127
                _ = v25
                _ = v26
                _ = v27
            }
            v28 := v1[1]
            _ = fi23
            _ = v24
            _ = v28
        }
    }
    v2 = v2
    var v29 int16 = -32768
    var v30 int16 = 32767
    fmt.Println(v30, v30, v30)
    var v31 int = 4294967295
    var v32 int64 = -227
    var v33 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
