package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 bool; F2 int }
type GenS4 struct { F0 int32; F1 string; F2 string }
func genFn0(p0 int16, p1 uint, p2 int64) int8 {
    var v5 string = "foo bar"
    switch 32768 {
    case -1:
        switch 0 {
        case 32767:
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 int64 = p2
                fmt.Println(fi6, v5)
                _ = fi6
                _ = v7
            }
        case -575:
            v5 = ""
        case 127:
            var v8 bool = ("hello" == "hello")
            _ = v8
        }
    case 42:
        var v9 float64 = 1.5
        _ = v9
    case 1:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 42}
            var v12 int64 = p2
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: -1}
    var v14 uint16 = 42
    p0 = (p0 << ((p1) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v5
    return 100
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    v3 := []int{9223372036854775807}
    v4 := []uint32{42, 2147483648, 2}
    if len(v4) > 0 {
        v4[0] = 666
    }
    switch "a" {
    case "hello":
        var v15 uint8 = 127
        _ = v15
    case "a":
        v16 := GenS1{F0: 874}
        _ = v16
    }
    fmt.Println(v1)
    switch "\n\t" {
    case "hello":
        v1 += 445
    case "foo bar":
        fmt.Println(v1)
    case "\n\t":
        if true {
            var v17 int16 = 1000
            _ = v17
        } else {
            if false {
                for fi18 := 0; fi18 < 19; fi18++ {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        fmt.Println(v1, fi19, fi18)
                        if (!false) {
                            var v20 uint8 = 1
                            v20 += (v20 ^ 128)
                            v21 := GenS3{F0: fi18, F1: (-672 >= -1), F2: (^(fi18))}
                            v22 := GenS4{F0: v1, F1: ("\n\t" + "\n\t"), F2: "hello"}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                    _ = fi18
                }
                var v23 int64 = -1
                v23 -= v23
                _ = v23
            }
            delete(v2, 580)
        }
    default:
        v1 = (-(v1))
    }
    var v24 int16 = 1000
    v1 = (v1 % (((^(v1))) | 1))
    for fi25 := 0; fi25 < 15; fi25++ {
        if true {
            v26 := GenS1{F0: 4294967295}
            var v27 string = ("hello" + "")
            _ = len(v4)
            _ = v26
            _ = v27
        } else {
            v1 = v1
        }
        _ = len(v2)
        _ = fi25
    }
    var v28 float64 = (0.0 + 1.5)
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(v24)
        fmt.Println(fi29, fi29)
        for fi30 := 0; fi30 < 17; fi30++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                var v31 uint16 = 770
                _ = v31
            case "hello":
                v1 += (v1 ^ (v1 & (v1 ^ v1)))
            case "foo bar":
                var v32 int64 = 0
                _ = v32
            default:
                v33 := GenS3{F0: (^(fi29)), F1: true, F2: -158}
                v24 = v24
                fmt.Println(fi29, v1)
                _ = v33
            }
            var v34 float32 = float32(v28)
            _ = fi30
            _ = v34
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
