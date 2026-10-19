package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 uint16, p2 uint8) uint {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            p1 += (^(p1))
            var v6 float32 = 1e10
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: 884, F1: 1, F2: ("" + ("foo bar" + ""))}
    var v8 string = "foo bar"
    var v9 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 372
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-38, 127, 9223372036854775807}
    var v3 uint = 1
    v3 += v3
    if false {
        v10 := GenS2{F0: 2147483647, F1: 100}
        fmt.Println(v3)
        v11 := GenS1{F0: 18446744073709551615, F1: -128, F2: "\n\t"}
        _ = v10
        _ = v11
    } else {
        switch (("foo bar" + "hello") + "foo bar") {
        case "":
            switch 4294967295 {
            case 1:
                fmt.Println(v3)
            case -419:
                if true {
                    v3 += 163
                } else {
                    v3 -= (v3 + 783)
                }
            }
        default:
            switch ("\n\t" + "\n\t") {
            case "":
                fmt.Println(v3)
            }
        }
        var v12 int32 = -1
        _ = v12
    }
    v13 := GenS3{F0: -721}
    v3 += (v3 * (v3 / ((65535) | 1)))
    v1[42] = 1
    if (true || false) {
        v14 := GenS4{F0: ("\n\t" + "\n\t"), F1: 953}
        _ = v14
    } else {
        v15 := GenS3{F0: 0}
        fmt.Println(v3)
        var v16 uint = v3
        _ = v15
        _ = v16
    }
    if false {
        fmt.Println(v3)
    }
    var v17 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
}
