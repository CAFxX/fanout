package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
func genFn0(p0 string, p1 uint16) uint64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        p1 = 32767
        p0 = ""
        _ = fi6
    }
    var v7 int8 = -1
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = 1e10
        _ = fi8
        _ = v9
    }
    v7 = v7
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 18446744073709551615
    v3 := make(map[int]int)
    v4 := []uint32{2147483647}
    var v5 uint = 18446744073709551615
    switch "\n\t" {
    case "a":
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                v1 += 1.5
                switch "a" {
                case "":
                    _ = len(v4)
                case "hello":
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if (2 != 2) {
                            var v12 int32 = 42
                            _ = v12
                        }
                        v13 := GenS1{F0: -2147483648}
                        _ = fi11
                        _ = v13
                    }
                }
                _ = fi10
            }
            switch ((("foo bar" + "hello") + ("hello" + "hello")) + "a") {
            case "foo bar":
                fmt.Println(v1, v2, v1)
            case "":
                v2 -= (v5 + 837)
            }
            fmt.Println(v1, v2, v1)
        } else {
            var v14 int = 2147483647
            v14 = v14
            v2 += (v5 | 368)
            _ = v14
        }
    case "":
        v15 := GenS1{F0: 1}
        _ = v15
    }
    var v16 uint8 = 84
    fmt.Println(v2, v2)
    var v17 uint8 = (v16 << (((v2 << ((v2) & 63))) & 63))
    switch 4294967295 {
    case 1:
        fmt.Println(v17, v17, v1)
    case 42:
        v2 = v5
    default:
        _ = len(v3)
    }
    v18 := GenS2{F0: v2, F1: 0}
    if ("a" <= "foo bar") {
        var v19 uint8 = 42
        _ = v19
    }
    var v20 uint = (v2 | (v2 / ((v5) | 1)))
    var v21 uint = v20
    v22 := GenS1{F0: -871}
    if (!(0 == 107)) {
        fmt.Println(v5, v5)
        if len(v4) > 0 {
            v4[0] = 800
        }
        v23 := GenS1{F0: -2147483648}
        _ = v23
    } else {
        v5 = v20
    }
    if true {
        var v24 string = ("a" + "foo bar")
        var v25 int64 = 392
        _ = v24
        _ = v25
    } else {
        switch 42 {
        case 422:
            if (!false) {
                var v26 int64 = 42
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int32 = -88
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        default:
            if (32767 < -1) {
                v1 = v1
            }
        }
        _ = len(v4)
    }
    fmt.Println(v17, v2, v5)
    var v29 int16 = -818
    fmt.Println(v21)
    var v30 int16 = 32767
    if (!true) {
        var v31 int8 = 1
        _ = v31
    } else {
        var v32 string = ""
        if false {
            fmt.Println(v30, v30, v16)
            v33 := GenS3{F0: 42, F1: true, F2: 267}
            _ = v33
        }
        _ = v32
    }
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
