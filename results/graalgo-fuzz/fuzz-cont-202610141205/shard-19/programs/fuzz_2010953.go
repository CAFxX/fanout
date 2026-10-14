package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 int32, p2 int32) uint {
    switch ("foo bar" + "") {
    case "":
        if (0 > 0) {
            var v5 uint8 = 42
            _ = v5
        }
    }
    switch "" {
    case "foo bar":
        var v6 uint32 = 42
        _ = v6
    case "hello":
        switch 9223372036854775807 {
        case 906:
            p0 += (p0 << ((552) & 63))
        default:
            var v7 int32 = 29
            _ = v7
        }
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            if (p0 <= p0) {
                p0 -= 65535
            } else {
                fmt.Println(p0, fi8, p0)
                if (-1 <= 1) {
                    var v9 string = "hello"
                    p0 -= 255
                    v10 := GenS1{F0: 0, F1: 571}
                    _ = v10
                    _ = v9
                } else {
                    var v11 float32 = 1.5
                    switch 128 {
                    case 127:
                        if ("hello" > "hello") {
                            p2 = 2147483647
                            v11 = 0.1
                        } else {
                            fmt.Println(p0)
                            v12 := []int{}
                            _ = v12
                        }
                    default:
                        fmt.Println(p2, v11, p2)
                    }
                    _ = v11
                }
            }
            _ = fi8
        }
    default:
        v13 := GenS1{F0: 0, F1: 2}
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = "a"
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS2{F0: 0, F1: p1}
        if (!false) {
            p2 -= int32(p0)
            var v18 string = ((("hello" + "foo bar") + "hello") + "\n\t")
            var v19 uint16 = 591
            _ = v18
            _ = v19
        } else {
            fmt.Println(p1, p2, p0)
            p1 = p1
            v20 := GenS2{F0: 18446744073709551615, F1: p1}
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            if (true && (32767 >= -512)) {
                var v22 float64 = ((float64((3.14159 / -2.25)) * 0.0) + 1.5)
                if (0 > 166) {
                    v23 := []int{fi16, 0, fi16}
                    if false {
                        fmt.Println(fi16, v22)
                    }
                    _ = v23
                }
                v24 := []int{fi16, fi21}
                _ = v22
                _ = v24
            } else {
                v25 := []int{4294967295, fi21}
                v26 := GenS3{F0: 42}
                _ = v25
                _ = v26
            }
            p1 = -385
            fmt.Println(p2, fi21, p0)
            _ = fi21
        }
        _ = fi16
        _ = v17
    }
    if false {
        p2 += (p2 >> ((34) & 63))
        if ((3.14159 * 1.5) > ((0.0 + 1.5) - 1e10)) {
            var v27 int = 4294967295
            var v28 uint64 = uint64(p0)
            _ = v27
            _ = v28
        }
        var v29 uint = p0
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1(p0 uint8, p1 bool, p2 int32) int8 {
    v30 := []int{}
    var v31 int32 = (p2 | p2)
    v32 := GenS3{F0: 9223372036854775807}
    _ = len(v30)
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v31
    _ = v32
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 float64 = (-2.25 / ((3.14159 * 0.1) + 0.1))
    var v4 uint16 = 42
    switch "hello" {
    case "a":
        for fi33 := 0; fi33 < 12; fi33++ {
            v2[145] = 251
            var v34 string = (("foo bar" + (("\n\t" + "a") + "hello")) + ((("foo bar" + "hello") + ("hello" + "hello")) + "a"))
            _ = fi33
            _ = v34
        }
    }
    var v35 uint = 18446744073709551615
    v35 -= v35
    v36 := []int{766, -2147483648, 127, 445}
    var v37 int16 = 1000
    var v38 float32 = (1e10 * 0.1)
    fmt.Println(v35, v37, v3)
    fmt.Println(v3, v4, v38)
    fmt.Println(v35, v35, v3)
    switch "foo bar" {
    case "":
        v39 := GenS2{F0: 42, F1: 2147483647}
        _ = v39
    case "\n\t":
        v36 = append(v36, 673)
    default:
        _ = len(v36)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
