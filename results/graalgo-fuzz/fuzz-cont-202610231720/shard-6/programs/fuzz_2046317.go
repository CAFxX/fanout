package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 uint16) uint16 {
    var v4 uint8 = 1
    var v5 uint = p0
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 = v5
        v7 := GenS1{F0: 204}
        v8 := GenS1{F0: (fi6 % (((fi6 >> (((p0 * p0)) & 63))) | 1))}
        _ = fi6
        _ = v7
        _ = v8
    }
    v4 = uint8(p1)
    fmt.Println(p1, v5, v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 957
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: 225}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 32767}
    var v12 string = ("a" + "")
    var v13 int64 = -9223372036854775808
    if false {
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                switch fi14 {
                case -9223372036854775808:
                    var v16 float64 = float64(1.5)
                    _ = v16
                default:
                    v17 := []int{}
                    _ = v17
                }
                _ = fi15
            }
            if true {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int32 = -1
                    if false {
                        var v20 uint8 = 195
                        var v21 uint64 = 569
                        v13 += v13
                        v22 := GenS3{F0: uint8(v21)}
                        if true {
                            v13 = v13
                        }
                        v23 := []int{2147483647, (fi18 + (-(fi18))), 127}
                        var v24 int32 = (-(v19))
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v12, fi18, fi18)
                        v25 := GenS4{F0: (v19 - (v19 | (v19 | (^(v19)))))}
                        fmt.Println(fi18)
                        _ = v25
                    }
                    var v26 float32 = 0.0
                    _ = fi18
                    _ = v19
                    _ = v26
                }
            }
            fmt.Println(v12)
            _ = fi14
        }
    } else {
        v27 := GenS4{F0: 0}
        v13 = 948
        v28 := GenS2{F0: 0}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := []int{}
        if false {
            var v31 int64 = 42
            fmt.Println(v12, v13)
            _ = v31
        } else {
            var v32 int32 = 1
            switch (^(fi29)) {
            case -680:
                switch -666 {
                case 613:
                    if true {
                        var v33 uint = 36
                        fmt.Println(v32, v13)
                        var v34 int8 = -7
                        _ = v33
                        _ = v34
                    }
                case 1:
                    var v35 float32 = 0.1
                    _ = v35
                }
            default:
                switch fi29 {
                case 32768:
                    v13 = v13
                case 127:
                    v12 = "hello"
                }
            }
            _ = v32
        }
        var v36 int64 = v13
        _ = fi29
        _ = v30
        _ = v36
    }
    v3[1] = (0.1 >= 0.0)
    if (int8(v13) < 127) {
        if false {
            v37 := GenS1{F0: 548}
            v12 = "foo bar"
            _ = len(v1)
            _ = v37
        }
        v12 = v12
    }
    v38 := GenS4{F0: -1}
    var v39 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
}
