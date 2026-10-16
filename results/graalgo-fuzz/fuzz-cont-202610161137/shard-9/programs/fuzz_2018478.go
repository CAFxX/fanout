package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int32) uint {
    v4 := GenS1{F0: -1}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 610
    var v3 uint16 = 289
    for fi5 := 0; fi5 < 3; fi5++ {
        if true {
            v6 := GenS2{F0: v1}
            for fi7 := 0; fi7 < 6; fi7++ {
                switch fi7 {
                case 1:
                    v2 += (^(v2))
                case 504:
                    if (!false) {
                        v8 := GenS3{F0: 221}
                        v2 = v2
                        _ = v8
                    } else {
                        v9 := []int{4294967295, (fi5 & 4294967295), int(v1), fi5}
                        var v10 int8 = 0
                        var v11 int64 = 9223372036854775807
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                case 154:
                    var v12 uint64 = v2
                    _ = v12
                default:
                    var v13 uint = genFn0(-2147483648)
                    _ = v13
                }
                v3 -= v3
                var v14 uint8 = uint8(v2)
                _ = fi7
                _ = v14
            }
            _ = v6
        } else {
            if (!true) {
                if (false || (-321 == 42)) {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        fmt.Println(fi5, v2)
                        _ = fi15
                    }
                }
            } else {
                v1 = v1
            }
            switch ("foo bar" + ("a" + "hello")) {
            case "a":
                v3 -= (v3 * v3)
            case "\n\t":
                var v16 float32 = 1e10
                _ = v16
            case "hello":
                v17 := GenS2{F0: 4294967295}
                _ = v17
            }
            v3 += 0
        }
        _ = fi5
    }
    fmt.Println(v1)
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := []int{127, fi18}
            fmt.Println(v1)
            _ = fi19
            _ = v20
        }
        v21 := GenS1{F0: 0}
        _ = fi18
        _ = v21
    }
    var v22 uint64 = v2
    v1 = 249
    for fi23 := 0; fi23 < 8; fi23++ {
        v2 -= v2
        _ = fi23
    }
    var v24 int8 = 1
    for fi25 := 0; fi25 < 8; fi25++ {
        if (2 == uint(v2)) {
            var v26 string = ("\n\t" + "a")
            var v27 uint = 712
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println(v3, v22, v3)
    v28 := GenS2{F0: 4294967295}
    v29 := GenS4{F0: v22, F1: (v22 - (^(v2)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
