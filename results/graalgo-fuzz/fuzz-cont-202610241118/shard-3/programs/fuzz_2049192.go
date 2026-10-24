package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = true
    var v2 int = 32767
    var v3 bool = false
    var v4 float32 = float32(-2.25)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 11; fi7++ {
        if v1 {
            v1 = true
            v8 := GenS1{F0: 1}
            _ = v8
        }
        var v9 uint32 = 0
        _ = fi7
        _ = v9
    }
    if v3 {
        switch "" {
        case "foo bar":
            v10 := []int{v2, (-(v2))}
            _ = v10
        case "":
            v11 := []int{-648}
            _ = v11
        }
    }
    var v12 int8 = 100
    if v3 {
        if false {
            v4 -= 3.14159
            fmt.Println(v12, v2)
        } else {
            v4 += 0.0
        }
    } else {
        v1 = v1
    }
    fmt.Println(v3, v4, v4)
    var v13 uint64 = 4294967295
    v14 := GenS1{F0: 2}
    if (0 > -44) {
        switch "hello" {
        case "a":
            if (!v1) {
                v15 := GenS1{F0: 0}
                _ = v15
            } else {
                switch v2 {
                case -372:
                    var v16 uint = 448
                    _ = v16
                case 128:
                    var v17 uint16 = 32767
                    _ = v17
                }
                switch "foo bar" {
                case "":
                    var v18 int = (v2 % (((-(v2))) | 1))
                    _ = v18
                }
            }
        case "":
            v1 = false
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                v4 = v4
                if ((v13 - 42) < (v13 << ((736) & 63))) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if ((v13 ^ (v13 * 931)) != (^(v13))) {
                            fmt.Println(v1)
                        }
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        switch "hello" {
                        case "foo bar":
                            v22 := []int{(fi21 | 4294967295), (v2 - (-(fi19)))}
                            _ = v22
                        case "\n\t":
                            v12 = (v12 % ((46) | 1))
                        default:
                            v23 := GenS2{F0: v12}
                            _ = v23
                        }
                        var v24 uint8 = 1
                        v13 += v13
                        _ = fi21
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v12 -= (v12 | 127)
                        var v26 uint = 255
                        if false {
                            var v27 bool = true
                            fmt.Println(v3, v1, v1)
                            v28 := []int{128, 583, 128}
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS3{F0: 961}
                            _ = v29
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                _ = fi19
            }
        }
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            fmt.Println(v13, fi30)
            v32 := GenS1{F0: 255}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
