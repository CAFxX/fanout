package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 uint8, p1 int16) uint16 {
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{fi8, (^(fi8)), (fi8 * (fi8 << ((0) & 63))), fi8}
        p0 -= 255
        v10 := GenS1{F0: 96, F1: 1, F2: "a"}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 42
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    var v12 int = -1
    _ = p0
    _ = v12
    return 32767
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 uint16 = 65535
    var v4 uint8 = v1
    var v5 uint8 = 1
    var v6 bool = (!false)
    v13 := GenS2{F0: 1}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int16 = genFn1(v6)
        _ = fi14
        _ = v15
    }
    var v16 bool = (("foo bar" == ("foo bar" + "foo bar")) || (42 <= 42))
    var v17 int8 = 97
    v5 += v5
    var v18 uint8 = 2
    switch ("a" + ("" + "")) {
    case "":
        fmt.Println(v18, v6, v6)
    default:
        var v19 int64 = 1
        _ = v19
    }
    v20 := GenS2{F0: (v17 | v17)}
    if v16 {
        var v21 uint16 = uint16(v5)
        _ = v21
    }
    var v22 uint16 = (^(v3))
    if false {
        v17 -= (v17 + v17)
    }
    switch -126 {
    case 9223372036854775807:
        var v23 int = 32767
        _ = v23
    case 32768:
        for fi24 := 0; fi24 < 18; fi24++ {
            if (fi24 != (fi24 | -9223372036854775808)) {
                switch "a" {
                case "":
                    v25 := GenS2{F0: (v17 - v17)}
                    _ = v25
                case "hello":
                    switch fi24 {
                    case -9223372036854775808:
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    case -867:
                        var v27 uint = 255
                        _ = v27
                    }
                }
            } else {
                v4 = 127
                fmt.Println(v5, v22, v4)
                delete(v2, (fi24 & fi24))
            }
            var v28 uint32 = 1
            switch (fi24 + fi24) {
            case 1:
                v29 := GenS3{F0: (v17 > -128), F1: v28, F2: -1}
                _ = v29
            }
            _ = fi24
            _ = v28
        }
    default:
        var v30 uint16 = v3
        _ = v30
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
