package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []uint32{639, 2147483648, 983}
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 16; fi4++ {
        _ = len(v2)
        var v5 uint32 = 42
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: -9223372036854775808}
    v7 := GenS1{F0: 32768}
    switch "\n\t" {
    case "":
        var v8 int8 = 0
        _ = v8
    }
    v9 := GenS1{F0: 953}
    switch 749 {
    case 42:
        var v10 uint64 = 789
        _ = v10
    case 317:
        v11 := GenS1{F0: 362}
        _ = v11
    case 4294967295:
        if (!false) {
            if false {
                var v12 string = "a"
                _ = v12
            } else {
                var v13 uint16 = 251
                v14 := v2["foo bar"]
                _ = v13
                _ = v14
            }
            if (!(!((3.14159 + 0.1) == 0.0))) {
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 string = ("hello" + ("" + ("" + ("hello" + "hello"))))
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 int32 = 42
                v18 := GenS2{F0: 2147483648}
                if (!false) {
                    var v19 int64 = 1
                    if (!(0 > 758)) {
                        var v20 int = -9223372036854775808
                        _ = v20
                    } else {
                        var v21 uint16 = 255
                        v21 -= (v21 & (v21 >> ((uint(v21)) & 63)))
                        _ = v21
                    }
                    v22 := GenS3{F0: 32767}
                    _ = v19
                    _ = v22
                } else {
                    v17 = v17
                    var v23 uint8 = 1
                    _ = v23
                }
                _ = v17
                _ = v18
            }
        } else {
            var v24 int8 = -12
            _ = len(v3)
            var v25 bool = false
            _ = v24
            _ = v25
        }
    default:
        if true {
            var v26 int8 = -34
            var v27 uint = 18446744073709551615
            var v28 float32 = 1e10
            _ = v26
            _ = v27
            _ = v28
        }
    }
    var v29 float64 = 1.5
    v30 := GenS4{F0: -423}
    var v31 int64 = 42
    var v32 int = 128
    v31 = 1
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v33 int = 42
    if true {
        v31 = (v31 << ((18446744073709551615) & 63))
    } else {
        v33 = 127
    }
    var v34 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v6
    _ = v7
    _ = v9
}
