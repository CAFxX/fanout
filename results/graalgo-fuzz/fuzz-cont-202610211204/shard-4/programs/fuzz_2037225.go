package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int32; F2 int16 }
func genFn0(p0 uint32, p1 int32, p2 int64) int32 {
    p0 -= 19
    var v7 int = 2147483647
    v8 := GenS1{F0: p0}
    switch "\n\t" {
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            p1 += 1
            _ = fi9
        }
    default:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint16 = 1
            fmt.Println(v7, fi10, v7)
            switch "\n\t" {
            case "hello":
                p2 = 42
            case "":
                v12 := GenS2{F0: -572, F1: 127}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 << ((2) & 63))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    var v13 bool = false
    v14 := GenS1{F0: 42}
    if v13 {
        if (false && v13) {
            v15 := []int{32768, 127}
            v13 = false
            _ = v15
        } else {
            switch 0 {
            case 9223372036854775807:
                var v16 float64 = (0.0 - (1e10 - 0.0))
                _ = v16
            }
            var v17 float64 = float64(0.1)
            _ = v17
        }
    }
    v13 = v13
    v18 := []int{}
    switch "foo bar" {
    case "foo bar":
        var v19 uint = 987
        _ = v19
    case "hello":
        if true {
            v13 = (true && (2147483648 != 1))
            v20 := GenS3{F0: 1}
            for fi21 := 0; fi21 < 16; fi21++ {
                v13 = v13
                _ = fi21
            }
            _ = v20
        }
    }
    fmt.Println(v13)
    switch "foo bar" {
    case "foo bar":
        v13 = (32767 <= 32767)
    case "\n\t":
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS4{F0: 65535}
            v24 := GenS5{F0: 1}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    v13 = v13
    if (42 < 84) {
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint16 = 42
            _ = fi25
            _ = v26
        }
    }
    if false {
        for fi27 := 0; fi27 < 11; fi27++ {
            if (100 != 1) {
                if (100 != 1) {
                    var v28 bool = v13
                    v29 := GenS4{F0: 2}
                    _ = v28
                    _ = v29
                } else {
                    v18 = append(v18, (fi27 + fi27))
                    fmt.Println(fi27)
                }
            } else {
                if (-42 >= 0) {
                    switch "foo bar" {
                    case "hello":
                        v30 := GenS3{F0: 128}
                        _ = v30
                    case "\n\t":
                        if (!v13) {
                            v13 = (v13 && false)
                            switch (fi27 - (fi27 >> ((65535) & 63))) {
                            case 1:
                                var v31 float64 = 1e10
                                _ = v31
                            case 0:
                                fmt.Println(v13, fi27)
                            case -965:
                                fmt.Println(fi27)
                            default:
                                v32 := v2[(fi27 * (fi27 << ((548) & 63)))]
                                _ = v32
                            }
                        }
                    default:
                        v13 = v13
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v13 = v13
                        v34 := GenS3{F0: 128}
                        var v35 bool = ((v13 && true) && v13)
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
                fmt.Println(fi27, v13)
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            }
            v36 := GenS6{F0: -9223372036854775808, F1: 2147483647, F2: 568}
            _ = fi27
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
