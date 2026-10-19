package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint) int64 {
    v7 := []int{-9223372036854775808, 128}
    p0 += (p0 & p0)
    var v8 uint64 = 896
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1() int16 {
    var v9 float64 = (float64(1e10) / 1e10)
    v10 := GenS1{F0: 42}
    _ = v10
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{127, -1, -939}
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    var v5 int32 = -941
    v6 := make(map[int]int)
    for fi11 := 0; fi11 < 7; fi11++ {
        if true {
            v5 = (v5 << ((616) & 63))
        }
        _ = fi11
    }
    if ((263 == 65535) && false) {
        v5 += (v5 * (v5 | -2147483648))
        v5 += (v5 * 42)
        _ = len(v2)
    }
    switch "" {
    case "\n\t":
        fmt.Println(v5)
    case "":
        var v12 int32 = (v5 << ((4294967295) & 63))
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS1{F0: 4294967295}
            _ = fi13
            _ = v14
        }
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        if false {
            v5 = v5
            v2 = append(v2, 1)
        }
    default:
        var v15 bool = true
        _ = v15
    }
    switch "foo bar" {
    case "a":
        for fi16 := 0; fi16 < 15; fi16++ {
            v5 -= (v5 << ((18446744073709551615) & 63))
            _ = fi16
        }
    default:
        var v17 float32 = 3.14159
        _ = v17
    }
    fmt.Println(v5)
    if false {
        v18 := GenS2{F0: 32767, F1: 0, F2: true}
        _ = v18
    } else {
        v5 = (v5 << ((65535) & 63))
        if true {
            v2 = append(v2, 128)
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS3{F0: (v5 ^ v5)}
                var v21 uint64 = 728
                for fi22 := 0; fi22 < 11; fi22++ {
                    switch int(v5) {
                    case 9223372036854775807:
                        var v23 string = ("a" + (("hello" + "\n\t") + (("foo bar" + ("foo bar" + ("" + "hello"))) + ("\n\t" + ("\n\t" + "\n\t")))))
                        _ = v23
                    case 800:
                        fmt.Println(fi22)
                    case 42:
                        fmt.Println(v21, v5)
                    default:
                        v21 += v21
                    }
                    _ = fi22
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            switch ((("\n\t" + "\n\t") + "foo bar") + "\n\t") {
            case "":
                if (4294967295 < 18446744073709551615) {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        v5 = 2147483647
                        _ = fi24
                    }
                    v5 = -1
                }
            case "a":
                switch "a" {
                case "":
                    v5 += 0
                }
            case "\n\t":
                v5 = (v5 + (v5 ^ v5))
            default:
                v6[128] = -156
            }
            var v25 int = -901
            v5 += -2147483648
            _ = v25
        }
    }
    var v26 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
