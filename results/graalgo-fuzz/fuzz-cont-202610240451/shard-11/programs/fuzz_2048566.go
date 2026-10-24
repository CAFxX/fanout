package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 string; F2 uint8 }
func genFn0(p0 string, p1 string, p2 int) int16 {
    fmt.Println(p1, p0, p1)
    var v5 float32 = 1e10
    if true {
        v6 := []int{p2, p2}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -773
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t", ""}
    if false {
        if (188 < -853) {
            var v7 uint = uint(v1)
            for fi8 := 0; fi8 < 17; fi8++ {
                v1 -= v1
                _ = fi8
            }
            if true {
                var v9 int8 = 100
                v10 := GenS1{F0: 2147483647}
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            if (!true) {
                v11 := GenS1{F0: -64}
                switch ("hello" + "hello") {
                case "foo bar":
                    var v12 int64 = 0
                    _ = v12
                default:
                    if true {
                        v1 = 50
                        for fi13 := 0; fi13 < 20; fi13++ {
                            if (-279 >= -229) {
                                var v14 uint16 = 32767
                                v4 = append(v4, "\n\t")
                                v14 = uint16(v1)
                                _ = v14
                            } else {
                                var v15 uint8 = v1
                                _ = v15
                            }
                            fmt.Println(fi13, v1)
                            var v16 float32 = ((0.0 * (1.5 * float32(0.0))) - float32((3.14159 * 0.1)))
                            _ = fi13
                            _ = v16
                        }
                        var v17 float32 = -2.25
                        _ = v17
                    }
                }
                _ = v11
            } else {
                fmt.Println(v1)
                _ = len(v2)
            }
            switch (("" + "\n\t") + "hello") {
            case "":
                var v18 bool = (-35 == 32767)
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 6; fi19++ {
                    if false {
                        _ = len(v4)
                        var v20 uint16 = 52
                        for fi21 := 0; fi21 < 7; fi21++ {
                            if ("hello" < ((("foo bar" + ("" + ("foo bar" + "a"))) + "\n\t") + "hello")) {
                                _ = len(v4)
                                fmt.Println(v1, fi21)
                                v22 := GenS2{F0: -7}
                                v4 = append(v4, "hello")
                                _ = v22
                            }
                            v23 := GenS2{F0: 0}
                            v24 := GenS2{F0: 0}
                            _ = fi21
                            _ = v23
                            _ = v24
                        }
                        _ = v20
                    }
                    if true {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            fmt.Println(fi19, fi19)
                            var v26 int64 = -543
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        if (96 < 42) {
                            v2[fi19] = ("" + "hello")
                        }
                    }
                    _ = fi19
                }
            }
        }
        var v27 uint64 = 18446744073709551615
        _ = v27
    }
    v28 := GenS3{F0: -2147483648}
    switch 32767 {
    case -444:
        var v29 uint32 = 687
        _ = v29
    case 2147483647:
        fmt.Println(v1)
    case -411:
        var v30 int16 = -1
        _ = v30
    }
    fmt.Println(v1)
    _ = len(v4)
    v31 := GenS2{F0: -91}
    var v32 float32 = (-2.25 + 0.0)
    v33 := GenS3{F0: -585}
    for fi34 := 0; fi34 < 16; fi34++ {
        var v35 bool = (!true)
        v36 := GenS4{F0: 67, F1: ("hello" + "\n\t"), F2: v1}
        var v37 uint32 = 1
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    var v38 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
}
