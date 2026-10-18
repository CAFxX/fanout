package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0(p0 uint64, p1 int64) uint8 {
    p0 -= 890
    if (((("\n\t" + ("" + "foo bar")) + "\n\t") > "\n\t") || (-1 != 127)) {
        if false {
            var v7 uint = 675
            v8 := []int{32768}
            _ = v7
            _ = v8
        }
        fmt.Println(p0, p1)
        switch -494 {
        case -402:
            if true {
                if (!false) {
                    var v9 uint8 = 127
                    _ = v9
                }
            }
        default:
            v10 := GenS1{F0: ("hello" + "a"), F1: 42}
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return 198
}
func genFn1(p0 int32) uint8 {
    p0 += p0
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 176
    var v3 int32 = -271
    var v4 int16 = 1000
    var v5 uint32 = 2147483647
    var v6 uint64 = 2
    v12 := GenS2{F0: 946, F1: -295}
    var v13 int64 = 42
    var v14 int32 = (^(v3))
    switch "a" {
    case "":
        if (!true) {
            var v15 string = ("\n\t" + "hello")
            var v16 float32 = -2.25
            _ = v15
            _ = v16
        }
    case "foo bar":
        v2 = 42
    default:
        v14 += 898
    }
    v17 := []int{9223372036854775807, 32768, 32767}
    v1[128] = true
    v18 := GenS2{F0: 4294967295, F1: (v4 & v4)}
    for fi19 := 0; fi19 < 2; fi19++ {
        switch fi19 {
        case -9223372036854775808:
            var v20 int8 = 0
            _ = v20
        case -2147483648:
            var v21 string = "foo bar"
            _ = v21
        default:
            if true {
                v22 := v1[(fi19 ^ (fi19 + fi19))]
                v23 := GenS3{F0: 42, F1: v2, F2: 65535}
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    _ = len(v1)
                    _ = fi24
                }
            }
        }
        _ = fi19
    }
    var v25 uint16 = 125
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
