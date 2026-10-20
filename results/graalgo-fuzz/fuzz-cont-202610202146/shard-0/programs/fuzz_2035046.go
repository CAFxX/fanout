package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    v4 := []int{42, -290, -837, 32768}
    v5 := GenS1{F0: 1}
    if (("hello" + "") <= "\n\t") {
        if (false || true) {
            for fi6 := 0; fi6 < 16; fi6++ {
                if (!false) {
                    var v7 uint8 = 42
                    for fi8 := 0; fi8 < 10; fi8++ {
                        var v9 string = "\n\t"
                        _ = fi8
                        _ = v9
                    }
                    v7 -= (v7 + 127)
                    _ = v7
                } else {
                    if len(v4) > 0 {
                        v4[0] = (fi6 >> ((122) & 63))
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch "foo bar" {
                case "hello":
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v3[2147483648] = false
                        _ = fi10
                    }
                default:
                    var v11 float32 = 1.5
                    _ = v11
                }
                _ = fi6
            }
        } else {
            var v12 int8 = 1
            var v13 int32 = 868
            _ = v12
            _ = v13
        }
        if true {
            var v14 int64 = 0
            _ = v14
        }
        var v15 string = ("\n\t" + (("a" + "") + "\n\t"))
        _ = v15
    }
    v16 := GenS2{F0: 4294967295}
    var v17 int8 = 42
    v17 += v17
    v17 += 119
    if (0.0 < (1e10 - 3.14159)) {
        if (v17 > (v17 | v17)) {
            var v18 uint32 = 0
            _ = v18
        }
        var v19 uint32 = 2
        _ = v19
    }
    switch 128 {
    case 42:
        if true {
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 int8 = 24
                for fi22 := 0; fi22 < 19; fi22++ {
                    if (-13 >= (v17 >> ((65535) & 63))) {
                        v23 := GenS1{F0: (v21 - 0)}
                        v17 += v17
                        fmt.Println(fi20, fi20)
                        v17 += (v21 * v17)
                        var v24 int64 = 0
                        _ = v23
                        _ = v24
                    } else {
                        var v25 bool = false
                        _ = v25
                    }
                    _ = fi22
                }
                _ = fi20
                _ = v21
            }
            if false {
                var v26 uint8 = 106
                fmt.Println(v17, v26)
                fmt.Println(v17, v26)
                _ = v26
            } else {
                for fi27 := 0; fi27 < 5; fi27++ {
                    if true {
                        v17 += 127
                    }
                    _ = fi27
                }
                fmt.Println(v17)
            }
            v28 := GenS3{F0: 1, F1: 2, F2: 2147483647}
            _ = v28
        }
    case 128:
        v4 = append(v4, 127)
    }
    fmt.Println(v17)
    v17 = v17
    v17 -= -3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
