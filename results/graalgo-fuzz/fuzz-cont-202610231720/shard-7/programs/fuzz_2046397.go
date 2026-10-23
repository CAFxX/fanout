package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 uint8, p1 uint16) uint32 {
    switch "foo bar" {
    case "a":
        switch 32767 {
        case 693:
            p0 -= p0
        case 128:
            fmt.Println(p1, p1)
        case 1:
            if (!false) {
                p1 += 418
                for fi5 := 0; fi5 < 17; fi5++ {
                    fmt.Println(fi5)
                    _ = fi5
                }
            }
        default:
            v6 := GenS1{F0: true}
            _ = v6
        }
    case "hello":
        var v7 uint = 446
        _ = v7
    }
    v8 := []int{9223372036854775807, 32768}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = len(v8)
    var v9 uint = 65535
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    v4 := []int{34, 9223372036854775807, -189}
    fmt.Println(v2, v1)
    var v10 uint32 = 42
    var v11 uint64 = 42
    if true {
        if true {
            v1 -= (v1 * -128)
            _ = len(v3)
            var v12 uint = 1
            _ = v12
        } else {
            v11 += (^(v11))
            var v13 string = (("a" + "hello") + "foo bar")
            switch v13 {
            case "":
                v14 := GenS1{F0: false}
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS2{F0: v13, F1: v10}
                    var v17 uint32 = uint32(v11)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v13
        }
        var v18 uint = 255
        _ = v18
    } else {
        switch "hello" {
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v1, v11)
                v11 += 2
                if true {
                    v20 := GenS3{F0: 0, F1: v1}
                    _ = v20
                }
                _ = fi19
            }
        }
        v21 := GenS3{F0: 1, F1: (v1 | v1)}
        switch -1 {
        case -570:
            switch 4294967295 {
            case 32768:
                v22 := GenS2{F0: "a", F1: (v10 + (v10 & 1))}
                _ = v22
            case 4294967295:
                fmt.Println(v1, v10)
            case 2147483647:
                fmt.Println(v1, v2)
            }
        case -2147483648:
            v10 = v10
        case 2147483647:
            v11 += v11
        default:
            fmt.Println(v1)
        }
        _ = v21
    }
    var v23 bool = (v2 > v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 float32 = ((v2 * v2) * (float32(-2.25) + (-2.25 / (float32((3.14159 / (1e10 - -2.25))) - -2.25))))
    fmt.Println(v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
