package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 int, p2 uint32) uint32 {
    fmt.Println(p2)
    if true {
        p0 -= (^(p0))
        if ("hello" >= "a") {
            fmt.Println(p0)
            if true {
                switch ("a" + "foo bar") {
                case "\n\t":
                    var v7 int64 = 581
                    _ = v7
                case "":
                    if true {
                        v8 := []int{p1}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v8
                    }
                default:
                    p2 -= 2
                }
                switch (p1 >> ((662) & 63)) {
                case -442:
                    switch ("a" + "a") {
                    case "foo bar":
                        var v9 uint8 = 128
                        _ = v9
                    default:
                        fmt.Println(p0)
                    }
                case 128:
                    var v10 int = (-(p1))
                    _ = v10
                }
                v11 := GenS1{F0: -1}
                _ = v11
            }
            v12 := []int{870, 585}
            _ = v12
        }
        var v13 uint8 = 2
        _ = v13
    }
    v14 := []int{(p1 / ((-1) | 1)), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 1
    var v4 int32 = -236
    var v5 uint = 18446744073709551615
    var v6 uint64 = 42
    if (((v5 ^ v5) < v5) || ("hello" == "")) {
        switch -446 {
        case 127:
            for fi15 := 0; fi15 < 3; fi15++ {
                v5 = 662
                v16 := []int{fi15, -144}
                _ = fi15
                _ = v16
            }
        case 0:
            v17 := []int{127, 2147483647, 2147483647}
            _ = v17
        }
        v18 := GenS1{F0: -9223372036854775808}
        _ = v18
    } else {
        v19 := GenS1{F0: 9223372036854775807}
        v4 -= (^(v4))
        for fi20 := 0; fi20 < 20; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                v6 = v6
                _ = fi21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS1{F0: 483}
                fmt.Println(v6)
                var v24 uint64 = v6
                if false {
                    v24 = 42
                    v3 -= 116
                } else {
                    v25 := GenS1{F0: 9223372036854775807}
                    var v26 uint32 = genFn0((v3 / (((v3 | 42)) | 1)), fi22, 2)
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v4 += int32(v5)
                var v28 uint16 = 65535
                _ = fi27
                _ = v28
            }
            _ = fi20
        }
        _ = v19
    }
    var v29 uint64 = (v6 << ((361) & 63))
    fmt.Println(v29)
    v30 := GenS1{F0: 42}
    var v31 int32 = (-(v4))
    var v32 uint8 = 150
    v6 = (v6 - v6)
    for fi33 := 0; fi33 < 3; fi33++ {
        var v34 int32 = v4
        var v35 int32 = v31
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 int16 = 32767
    switch 128 {
    case -204:
        v37 := GenS2{F0: (v4 / (((^(v4))) | 1)), F1: v3}
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        v29 -= (^(v6))
        _ = fi38
    }
    v39 := GenS3{F0: 586}
    var v40 bool = ("a" <= "")
    v41 := []int{-134, 268, int(v29)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
