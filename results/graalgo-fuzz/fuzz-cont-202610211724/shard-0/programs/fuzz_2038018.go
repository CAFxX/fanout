package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) uint {
    if true {
        switch 128 {
        case 128:
            if true {
                var v4 int = -2147483648
                _ = v4
            } else {
                if false {
                    p0 -= -123
                    switch "\n\t" {
                    case "hello":
                        v5 := GenS1{F0: -486}
                        _ = v5
                    }
                }
            }
        case 127:
            v6 := []int{128, -1, -1, 0}
            _ = v6
        case 2147483647:
            v7 := []int{127, 32768}
            _ = v7
        default:
            p0 -= p0
        }
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int = (^(fi8))
            _ = fi8
            _ = v9
        }
    }
    var v10 int = 42
    fmt.Println(p0, v10)
    _ = p0
    _ = v10
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]uint32)
    var v3 string = "a"
    switch 460 {
    case -9223372036854775808:
        v11 := GenS1{F0: -1}
        _ = v11
    case 493:
        v12 := GenS1{F0: 667}
        _ = v12
    }
    v3 = ""
    v13 := GenS1{F0: -32768}
    v3 = "a"
    if (1e10 <= float32(0.0)) {
        switch 38 {
        case 32767:
            delete(v2, (v3 + v3))
        case 0:
            v3 = ""
        }
    }
    v3 = ("a" + v3)
    var v14 int64 = 1
    switch 582 {
    case 2147483647:
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 = v1
            _ = fi15
        }
    case 833:
        for fi16 := 0; fi16 < 4; fi16++ {
            v1 = v1
            for fi17 := 0; fi17 < 12; fi17++ {
                if true {
                    var v18 uint32 = 4294967295
                    if true {
                        v19 := []int{}
                        if (true && (false && true)) {
                            var v20 int8 = 100
                            _ = v20
                        }
                        _ = v19
                    } else {
                        v3 = "foo bar"
                    }
                    v21 := GenS1{F0: -32768}
                    v22 := []int{-2147483648}
                    _ = v18
                    _ = v21
                    _ = v22
                }
                _ = fi17
            }
            v23 := GenS2{F0: (v1 | (^(v1)))}
            _ = fi16
            _ = v23
        }
    case 1:
        v3 = "a"
    }
    _ = len(v2)
    v14 = (^(v14))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
