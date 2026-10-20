package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32) uint8 {
    switch "" {
    case "":
        if (-70 <= 94) {
            var v6 int16 = 32767
            v7 := GenS1{F0: 329}
            _ = v6
            _ = v7
        }
    }
    var v8 int16 = 1
    if true {
        if (true && (!false)) {
            var v9 int16 = v8
            var v10 int32 = 417
            for fi11 := 0; fi11 < 9; fi11++ {
                switch (("hello" + ("foo bar" + "foo bar")) + "hello") {
                case "":
                    switch "\n\t" {
                    case "a":
                        var v12 int64 = 533
                        _ = v12
                    default:
                        v13 := []int{int(v10)}
                        _ = v13
                    }
                case "a":
                    var v14 int = (fi11 & (fi11 ^ -793))
                    _ = v14
                }
                switch "\n\t" {
                case "\n\t":
                    var v15 float32 = 1e10
                    _ = v15
                case "foo bar":
                    fmt.Println(v9, v8)
                case "hello":
                    var v16 float64 = 0.0
                    _ = v16
                default:
                    v17 := []int{}
                    _ = v17
                }
                v18 := []int{(fi11 * 128), 9223372036854775807, fi11}
                _ = fi11
                _ = v18
            }
            _ = v10
            _ = v9
        } else {
            p0 = 287
            var v19 string = "foo bar"
            _ = v19
        }
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 0
    var v3 float32 = (3.14159 + float32(float64(0.1)))
    v4 := make(map[uint32]int)
    var v5 uint8 = 55
    var v20 int32 = -1
    v20 = 1
    v21 := []int{377, -567}
    var v22 int16 = 1000
    switch -257 {
    case 128:
        v23 := GenS1{F0: 2}
        _ = v23
    case -554:
        fmt.Println(v22, v2, v2)
    case -681:
        fmt.Println(v2, v5, v3)
    default:
        v24 := GenS1{F0: 681}
        _ = v24
    }
    if (1 <= 32767) {
        v2 = v2
        var v25 int = -1
        v22 = v22
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        fmt.Println(v20)
        _ = fi26
    }
    var v27 uint8 = genFn0(2)
    fmt.Println(v5, v2)
    if true {
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 int = (fi28 & (fi29 + 1))
                var v31 int16 = v22
                fmt.Println(v30, fi28)
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        v22 = v22
    } else {
        v32 := GenS1{F0: 358}
        var v33 int64 = -926
        v3 = 0.1
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        fmt.Println(v3, v20, v2)
        _ = fi34
    }
    v20 -= (v20 % (((v20 * -2147483648)) | 1))
    v3 = 1.5
    fmt.Println(v2, v22)
    var v35 int32 = (v20 % ((0) | 1))
    if len(v21) > 0 {
        v21[0] = 0
    }
    var v36 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
