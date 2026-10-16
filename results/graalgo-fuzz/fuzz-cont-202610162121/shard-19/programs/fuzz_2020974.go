package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{725, 2, 2}
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    v4 := []string{}
    v2 -= -1
    fmt.Println(v2)
    v2 -= -222
    _ = len(v3)
    v2 = v2
    if (1 >= 230) {
        switch 32768 {
        case 548:
            for fi5 := 0; fi5 < 9; fi5++ {
                v2 = int32(fi5)
                switch "foo bar" {
                case "":
                    fmt.Println(fi5, v2)
                default:
                    fmt.Println(v2, v2)
                }
                _ = fi5
            }
        case 774:
            var v6 float32 = 0.1
            _ = v6
        case -838:
            if true {
                _ = len(v3)
                var v7 string = "a"
                _ = v7
            }
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v2, fi9, fi8)
                v2 -= v2
                _ = fi9
            }
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 int8 = -1
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 9; fi12++ {
                switch "a" {
                case "\n\t":
                    if false {
                        var v13 float64 = (((0.0 / 1.5) - (0.0 * 3.14159)) + 1.5)
                        v14 := GenS1{F0: 0}
                        _ = v13
                        _ = v14
                    }
                default:
                    var v15 int = fi8
                    _ = v15
                }
                _ = len(v3)
                _ = fi12
            }
            _ = fi8
        }
        v16 := v3[892]
        _ = v16
    } else {
        if true {
            var v17 uint16 = 2
            if false {
                switch -680 {
                case 983:
                    var v18 uint64 = 1
                    _ = v18
                case -9223372036854775808:
                    v19 := GenS1{F0: 1}
                    _ = v19
                case 0:
                    if (!true) {
                        var v20 uint8 = 127
                        v21 := GenS1{F0: 65535}
                        var v22 uint32 = 397
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                default:
                    var v23 string = ((("foo bar" + "hello") + "") + "\n\t")
                    _ = v23
                }
                var v24 int64 = 1
                var v25 int = 42
                _ = v24
                _ = v25
            } else {
                var v26 int64 = 1
                fmt.Println(v2)
                if true {
                    v27 := GenS1{F0: 765}
                    fmt.Println(v26, v26)
                    v2 -= v2
                    _ = v27
                } else {
                    v26 += 1
                }
                _ = v26
            }
            _ = v17
        } else {
            for fi28 := 0; fi28 < 13; fi28++ {
                switch 4294967295 {
                case -475:
                    var v29 int8 = 127
                    _ = v29
                }
                v2 += v2
                _ = fi28
            }
            if (!false) {
                for fi30 := 0; fi30 < 10; fi30++ {
                    fmt.Println(v2, fi30)
                    _ = fi30
                }
                v2 += v2
            }
            v31 := v3[503]
            _ = v31
        }
        if (2 <= 128) {
            v2 = v2
            var v32 bool = (9223372036854775807 < 9223372036854775807)
            _ = v32
        }
        v2 = v2
    }
    fmt.Println(v2)
    if true {
        fmt.Println(v2)
    }
    if true {
        var v33 string = "a"
        v2 = 1
        for fi34 := 0; fi34 < 18; fi34++ {
            v33 = v33
            v33 = ""
            for fi35 := 0; fi35 < 18; fi35++ {
                delete(v3, 0)
                _ = fi35
            }
            _ = fi34
        }
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
