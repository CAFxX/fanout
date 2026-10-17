package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 int8) int64 {
    p0 += (p0 << ((255) & 63))
    fmt.Println(p0)
    _ = p0
    return 0
}
func genFn1(p0 uint, p1 uint32, p2 uint) uint32 {
    switch -1 {
    case 550:
        if ((-2.25 > 0.1) && (775 <= int16(p2))) {
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := []int{(-(fi7)), (fi7 / ((-91) | 1))}
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: -699}
            p1 = p1
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 20; fi10++ {
            p0 = 207
            _ = fi10
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (p1 >> (((p2 / (((p2 ^ (^(p2)))) | 1))) & 63))
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 65535
    v3 := []uint32{631, 494}
    var v4 int16 = 932
    v5 := make(map[uint32]string)
    var v6 int16 = (^(v4))
    v4 = (v6 / (((v4 ^ v4)) | 1))
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := v5[65535]
        v4 = (v6 + (v6 ^ -1))
        v13 := GenS2{F0: v2, F1: 0, F2: (v4 ^ v6)}
        _ = fi11
        _ = v12
        _ = v13
    }
    v4 = -791
    switch ("hello" + "a") {
    case "hello":
        if (4294967295 == (v2 + 0)) {
            var v14 uint32 = 2147483647
            v15 := GenS1{F0: 519}
            v16 := GenS1{F0: genFn0(42)}
            _ = v14
            _ = v15
            _ = v16
        } else {
            v4 = v4
            fmt.Println(v1, v6, v2)
        }
    case "":
        if true {
            switch -442 {
            case -930:
                var v17 float32 = (0.1 - 0.0)
                _ = v17
            case -9223372036854775808:
                v6 += v4
            default:
                v3 = append(v3, 501)
            }
            v2 += v2
            var v18 float32 = 1.5
            _ = v18
        }
    case "foo bar":
        v6 = (v6 | v4)
    default:
        v19 := GenS2{F0: (v2 & (^(v2))), F1: 100, F2: 42}
        _ = v19
    }
    v6 = v4
    for fi20 := 0; fi20 < 19; fi20++ {
        switch (fi20 + (fi20 - -1)) {
        case 4294967295:
            if true {
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 int8 = -10
                    var v23 int64 = -351
                    _ = len(v5)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                _ = len(v3)
            }
        default:
            v1 += 42
        }
        _ = fi20
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int8 = 0
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
