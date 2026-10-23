package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 2147483647
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 uint32 = 375
    v2 += 0
    var v6 float32 = ((-2.25 + 1e10) + (1.5 - 1e10))
    v5 += v5
    v5 = (^(v5))
    var v7 uint8 = 191
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 -= (fi8 >> ((255) & 63))
            _ = fi9
        }
        _ = fi8
    }
    var v10 string = (("foo bar" + "a") + ("\n\t" + (("a" + ("foo bar" + "")) + "\n\t")))
    var v11 uint64 = 751
    fmt.Println(v7, v2)
    var v12 float64 = -2.25
    var v13 uint16 = 42
    v14 := GenS1{F0: (v5 - v5)}
    if (true && true) {
        var v15 uint8 = uint8(v13)
        if false {
            v16 := GenS2{F0: uint64(v5), F1: (v11 + (v11 << ((42) & 63)))}
            fmt.Println(v13, v15)
            var v17 int = 32768
            _ = v16
            _ = v17
        }
        _ = v15
    } else {
        v11 -= v11
        if (true && (89 == 0)) {
            var v18 int64 = 42
            _ = v18
        } else {
            if (42 < -32768) {
                v19 := GenS3{F0: 0, F1: "a", F2: 42}
                v20 := GenS3{F0: -104, F1: v10, F2: v11}
                _ = v19
                _ = v20
            } else {
                switch 32767 {
                case 1:
                    v7 = v7
                }
                var v21 float32 = 1e10
                v22 := GenS3{F0: -1, F1: (v10 + ((((v10 + "") + v10) + v10) + "a")), F2: v11}
                _ = v21
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint8 = v7
        _ = fi23
        _ = v24
    }
    if (false && ((v2 - -1) > int(v13))) {
        switch v10 {
        case "a":
            var v25 int32 = 2147483647
            _ = v25
        case "hello":
            var v26 int16 = 1
            _ = v26
        }
    } else {
        v27 := GenS1{F0: 4294967295}
        var v28 int8 = 100
        _ = v27
        _ = v28
    }
    if false {
        v29 := GenS4{F0: 42}
        v11 = 1
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
