package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 uint16) bool {
    var v5 int64 = (p0 + (p0 / ((-403) | 1)))
    var v6 int32 = 42
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 int = -1
    v2 := make(map[string]string)
    var v3 float32 = (-2.25 / (((1.5 * 0.1) + (-2.25 / 0.1)) * -2.25))
    v4 := make(map[uint32]string)
    var v7 int32 = 2147483647
    var v8 uint8 = 0
    var v9 uint8 = v8
    v1 += v1
    v3 -= (1.5 * (v3 - v3))
    fmt.Println(v1)
    v10 := GenS1{F0: 12}
    switch ("" + "foo bar") {
    case "a":
        for fi11 := 0; fi11 < 16; fi11++ {
            if false {
                if false {
                    v12 := []int{278}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        if true {
                            v3 = (v3 - float32(((1.5 - float64(float32(1e10))) / 0.1)))
                            v14 := GenS1{F0: 100}
                            _ = v14
                        }
                        fmt.Println(v1, v8)
                        v15 := GenS1{F0: -35}
                        _ = fi13
                        _ = v15
                    }
                    _ = v12
                } else {
                    v3 -= 3.14159
                    fmt.Println(v9)
                }
                switch (fi11 | v1) {
                case -907:
                    var v16 bool = true
                    _ = v16
                }
            } else {
                var v17 uint16 = 42
                v18 := []int{-1}
                _ = v17
                _ = v18
            }
            if (((fi11 - (fi11 - v1)) < v1) || (false || false)) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v7 -= (^(v7))
                    v20 := GenS1{F0: -1}
                    _ = fi19
                    _ = v20
                }
            }
            if (-440 >= 0) {
                if false {
                    if false {
                        v3 = 0.0
                        v21 := []int{-2147483648, 2147483647, (v1 >> ((42) & 63))}
                        for fi22 := 0; fi22 < 3; fi22++ {
                            fmt.Println(v7, fi11)
                            _ = fi22
                        }
                        _ = v21
                    }
                } else {
                    var v23 bool = false
                    _ = v23
                }
            }
            _ = fi11
        }
    case "":
        fmt.Println(v1)
    case "foo bar":
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 14; fi25++ {
                switch "foo bar" {
                case "\n\t":
                    var v26 uint16 = 692
                    _ = v26
                case "hello":
                    v3 += float32(0.0)
                default:
                    v27 := []int{v1, fi25, 32767, 32768}
                    _ = v27
                }
                v3 = (v3 * v3)
                _ = fi25
            }
            switch (v1 << ((1) & 63)) {
            case 100:
                switch (fi24 << ((255) & 63)) {
                case 4294967295:
                    fmt.Println(v7)
                default:
                    for fi28 := 0; fi28 < 20; fi28++ {
                        v29 := []int{32768, 1, fi24}
                        v30 := GenS1{F0: -122}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
            case 1:
                var v31 float64 = (1e10 - ((0.0 / 1e10) / 1.5))
                _ = v31
            default:
                v7 = (v7 ^ (v7 | (v7 ^ (v7 & 39))))
            }
            v32 := GenS1{F0: 11}
            _ = fi24
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 10; fi33++ {
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS2{F0: -512, F1: v9, F2: 2}
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    }
    if true {
        v9 -= (v8 ^ (v8 / ((v8) | 1)))
    } else {
        switch "foo bar" {
        case "hello":
            var v36 uint = 65535
            _ = v36
        default:
            v37 := GenS3{F0: 32767, F1: 1, F2: 65535}
            _ = v37
        }
        fmt.Println(v3)
    }
    var v38 uint32 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
