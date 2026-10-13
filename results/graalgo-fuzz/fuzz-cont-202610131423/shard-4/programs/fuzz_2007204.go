package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 uint64) int16 {
    v5 := []int{4294967295, 1, 42, -894}
    switch 2147483647 {
    case 998:
        var v6 uint8 = 128
        _ = v6
    case -986:
        p1 -= p1
    case 1:
        switch "a" {
        case "foo bar":
            switch 191 {
            case 4294967295:
                p0 = 1
            default:
                v7 := GenS1{F0: 42}
                _ = v7
            }
        case "a":
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS2{F0: 168}
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    p1 += p1
                    fmt.Println(fi8)
                }
                _ = fi8
                _ = v9
            }
        case "":
            v10 := GenS2{F0: 1}
            _ = v10
        default:
            p0 = -1
        }
    default:
        var v11 uint = 42
        _ = v11
    }
    var v12 float32 = 1e10
    if false {
        _ = len(v5)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v13 int16 = 0
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return 627
}

func main() {
    var v1 int16 = 1000
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 string = ""
    v1 = -32768
    switch "a" {
    case "a":
        var v14 int64 = -1
        _ = v14
    default:
        if (!true) {
            switch 32768 {
            case -9223372036854775808:
                v15 := GenS1{F0: -9223372036854775808}
                _ = v15
            }
        } else {
            v16 := v3[9223372036854775807]
            switch (v16 / (((v16 + v16)) | 1)) {
            case -1:
                var v17 bool = (false && false)
                _ = v17
            case 9223372036854775807:
                for fi18 := 0; fi18 < 3; fi18++ {
                    v2 = append(v2, 4294967295)
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v1 -= (v1 - -32768)
                        _ = fi19
                    }
                    _ = fi18
                }
            }
            _ = v16
        }
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            v4 = v4
            _ = fi21
        }
        _ = fi20
    }
    v1 += 1000
    v4 = (v4 + v4)
    if (true || (-229 < 495)) {
        v1 = -32768
    }
    if (65535 == 0) {
        _ = len(v2)
        v1 = genFn0(1, 957)
    } else {
        var v22 int = -1
        v22 = (v22 * (v22 % ((v22) | 1)))
        _ = v22
    }
    var v23 uint16 = 0
    v24 := GenS2{F0: 65535}
    var v25 uint8 = 14
    switch 1 {
    case -2147483648:
        var v26 int16 = 1000
        _ = v26
    case -9223372036854775808:
        for fi27 := 0; fi27 < 11; fi27++ {
            v28 := GenS2{F0: 709}
            var v29 int16 = (^(v1))
            _ = fi27
            _ = v28
            _ = v29
        }
    case 32767:
        var v30 int = 9223372036854775807
        _ = v30
    default:
        var v31 uint32 = 65535
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
