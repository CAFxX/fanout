package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 65535
    var v5 float64 = 1e10
    switch 410 {
    case 32768:
        if false {
            var v6 float64 = float64(0.1)
            var v7 uint16 = 407
            _ = v6
            _ = v7
        }
    case 32767:
        if false {
            var v8 uint64 = 4294967295
            _ = v8
        }
    case -735:
        v9 := GenS1{F0: (125 < -98), F1: 2}
        _ = v9
    default:
        fmt.Println(v5, v5)
    }
    v10 := GenS1{F0: (false || true), F1: 42}
    if false {
        var v11 string = ("a" + "")
        v12 := GenS2{F0: (false || (false && false)), F1: 24}
        var v13 float32 = 1e10
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := v2[("foo bar" + "\n\t")]
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: false, F1: (v4 << ((90) & 63))}
            var v17 int32 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v5 = -2.25
        _ = fi18
    }
    v19 := []int{-1}
    var v20 uint = 255
    v5 += ((v5 - float64(0.1)) + ((v5 - 0.1) * v5))
    var v21 int64 = 42
    if (!true) {
        if (225 < 42) {
            switch -813 {
            case 127:
                for fi22 := 0; fi22 < 13; fi22++ {
                    v5 += 3.14159
                    _ = fi22
                }
            case 2147483647:
                fmt.Println(v4, v21, v21)
            case 326:
                var v23 float32 = -2.25
                _ = v23
            default:
                v5 = float64(3.14159)
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                if (true || (!(!(v5 < v5)))) {
                    var v25 string = ""
                    var v26 bool = true
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        _ = len(v19)
        for fi28 := 0; fi28 < 15; fi28++ {
            switch "\n\t" {
            case "hello":
                v29 := GenS3{F0: 0, F1: "", F2: v21}
                _ = v29
            case "\n\t":
                var v30 int8 = int8(v4)
                _ = v30
            case "":
                if len(v19) > 0 {
                    v19[0] = fi28
                }
            }
            var v31 float64 = v5
            v32 := GenS1{F0: (669 != 529), F1: 128}
            _ = fi28
            _ = v31
            _ = v32
        }
        fmt.Println(v4)
        _ = fi27
    }
    fmt.Println(v4, v21, v21)
    v33 := GenS4{F0: (!true), F1: 1, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
