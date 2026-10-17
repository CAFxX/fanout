package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int32 = -1
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    var v4 int16 = 1
    if (975 < 255) {
        v2 += 42
        for fi5 := 0; fi5 < 16; fi5++ {
            if true {
                for fi6 := 0; fi6 < 13; fi6++ {
                    if (1.5 != 1e10) {
                        v2 = (v2 << ((42) & 63))
                    }
                    _ = fi6
                }
                var v7 uint = 0
                _ = v7
            } else {
                var v8 uint = 93
                fmt.Println(v1, fi5, fi5)
                _ = len(v3)
                _ = v8
            }
            _ = fi5
        }
    } else {
        v9 := GenS1{F0: 2, F1: 42}
        v1 += v1
        _ = v9
    }
    var v10 int64 = -190
    v11 := GenS1{F0: 0, F1: 95}
    if false {
        var v12 float64 = 1e10
        _ = v12
    } else {
        var v13 uint = 4294967295
        _ = v13
    }
    var v14 uint64 = 42
    for fi15 := 0; fi15 < 4; fi15++ {
        if ((((("" + ("hello" + "hello")) + "\n\t") + ("hello" + "foo bar")) <= "") || (true || true)) {
            v16 := GenS2{F0: 179, F1: (v1 ^ (-(v1))), F2: 0}
            _ = v16
        } else {
            v3[v2] = "a"
            v14 = v14
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint32 = v2
            fmt.Println(v10)
            _ = fi17
            _ = v18
        }
        var v19 int32 = 1
        _ = fi15
        _ = v19
    }
    if (false || true) {
        var v20 string = ((("" + (("" + "\n\t") + "a")) + "foo bar") + ("" + ""))
        _ = v20
    } else {
        v3[v2] = ((("\n\t" + "foo bar") + "\n\t") + ("\n\t" + "hello"))
        v4 = 32767
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        switch "foo bar" {
        case "foo bar":
            v2 = v2
        default:
            v4 += v4
        }
        v14 = 0
        v22 := GenS2{F0: 128, F1: int32(v4), F2: 143}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS3{F0: v1}
        v25 := GenS3{F0: -305}
        for fi26 := 0; fi26 < 1; fi26++ {
            if true {
                v27 := []int{}
                _ = v27
            }
            v3[uint32(v14)] = "a"
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
