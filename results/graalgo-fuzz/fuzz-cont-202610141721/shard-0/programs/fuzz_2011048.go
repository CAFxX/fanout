package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int) int16 {
    var v6 uint16 = 32767
    p0 = (p0 * p0)
    var v7 float32 = 0.0
    var v8 int = 946
    if true {
        for fi9 := 0; fi9 < 9; fi9++ {
            if false {
                v6 += v6
            }
            _ = fi9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int16 = 0
            var v12 int8 = 1
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "":
                    var v13 int64 = -9223372036854775808
                    _ = v13
                case "hello":
                    var v14 uint = 64
                    _ = v14
                case "\n\t":
                    var v15 uint32 = 152
                    _ = v15
                default:
                    v6 = 2
                }
            default:
                var v16 int8 = -128
                _ = v16
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        v7 = v7
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 uint16 = 525
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 42
    fmt.Println(v5, v3)
    fmt.Println(v5, v5)
    fmt.Println(v5, v3)
    v1[949] = 2147483647
    if true {
        fmt.Println(v3, v5)
        v17 := []int{0}
        _ = v17
    } else {
        var v18 int64 = -9223372036854775808
        fmt.Println(v3)
        _ = v18
    }
    var v19 float64 = 1e10
    var v20 int8 = -1
    switch "\n\t" {
    case "foo bar":
        v5 -= v5
    case "":
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS1{F0: -9223372036854775808}
            if true {
                if ((v20 * v20) < (v20 ^ (v20 | (v20 * v20)))) {
                    v5 += v5
                    switch int(v20) {
                    case 128:
                        v19 = (v19 / v19)
                    case 42:
                        if false {
                            v23 := GenS1{F0: 42}
                            var v24 int32 = 1
                            _ = v23
                            _ = v24
                        }
                    default:
                        v20 -= (v20 + (v20 + 75))
                    }
                    _ = len(v2)
                }
                var v25 uint64 = v5
                for fi26 := 0; fi26 < 16; fi26++ {
                    fmt.Println(fi21)
                    v3 = v3
                    v5 = (v25 >> ((1) & 63))
                    if true {
                        var v27 uint = 65535
                        if (!(2147483647 <= -700)) {
                            var v28 float64 = v19
                            _ = v28
                        }
                        v29 := []int{-456, (fi26 / (((fi21 * fi26)) | 1)), int(v25)}
                        _ = v27
                        _ = v29
                    }
                    _ = fi26
                }
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    case "hello":
        switch "foo bar" {
        case "a":
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := []int{}
                _ = fi30
                _ = v31
            }
        case "":
            if (true || (true && false)) {
                v32 := GenS1{F0: -1}
                _ = v32
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
