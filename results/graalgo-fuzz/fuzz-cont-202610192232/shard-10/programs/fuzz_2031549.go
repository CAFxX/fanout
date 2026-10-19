package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8, p1 uint16) uint8 {
    var v6 int8 = p0
    p1 -= (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v6
    return 75
}
func genFn1() uint16 {
    var v7 int32 = 2147483647
    var v8 int64 = -230
    _ = v7
    _ = v8
    return 406
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 uint64 = v1
    var v5 string = (("a" + "a") + ("a" + "a"))
    var v9 int32 = -1
    fmt.Println(v1, v4)
    v10 := []int{}
    v11 := GenS1{F0: (v5 + v5)}
    v12 := GenS1{F0: v5}
    switch 4294967295 {
    case 42:
        delete(v3, 122)
    case 9223372036854775807:
        var v13 uint64 = v4
        _ = v13
    case 32767:
        if (!(-9223372036854775808 != 9223372036854775807)) {
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint8 = 56
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS2{F0: v4}
                v5 = v5
                _ = fi16
                _ = v17
            }
        }
    }
    delete(v3, 2147483647)
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            for fi19 := 0; fi19 < 4; fi19++ {
                switch fi18 {
                case 128:
                    var v20 int16 = -466
                    _ = v20
                case 32768:
                    v1 -= 338
                default:
                    var v21 int64 = 42
                    _ = v21
                }
                _ = fi19
            }
        }
        switch v5 {
        case "\n\t":
            var v22 float32 = 0.0
            _ = v22
        case "a":
            for fi23 := 0; fi23 < 19; fi23++ {
                v1 -= (v1 * 4294967295)
                var v24 int16 = 1000
                var v25 int32 = 44
                _ = fi23
                _ = v24
                _ = v25
            }
        case "":
            v26 := GenS2{F0: 109}
            _ = v26
        }
        fmt.Println(v4, v1)
        _ = fi18
    }
    if len(v10) > 0 {
        v10[0] = 128
    }
    if ((255 != genFn0(-128, 255)) || true) {
        var v27 float32 = 1.5
        _ = v27
    } else {
        v28 := GenS3{F0: 32767}
        v5 = "a"
        v5 = v5
        _ = v28
    }
    var v29 int8 = 25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
