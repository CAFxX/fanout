package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool }
func genFn0() uint16 {
    var v5 int32 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{391, fi6, (fi6 + 331)}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: v5, F1: 42, F2: 833}
    fmt.Println(v5)
    _ = v5
    _ = v8
    return 822
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 783
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    var v9 uint = 2
    v10 := GenS1{F0: 1, F1: 1, F2: (^(v9))}
    v2 = 0
    v11 := []int{}
    if ((v9 | (^(v9))) >= (v9 ^ 0)) {
        fmt.Println(v9, v9)
        v12 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: 908}
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v9 += (v9 * (v9 << ((v9) & 63)))
        _ = len(v4)
        var v14 bool = ((!(2147483647 > -1)) || (true || true))
        _ = fi13
        _ = v14
    }
    v1[971] = (("a" + ("" + "\n\t")) + "")
    for fi15 := 0; fi15 < 10; fi15++ {
        if ("a" != "hello") {
            switch "\n\t" {
            case "":
                fmt.Println(v2, v2, v2)
            case "foo bar":
                fmt.Println(fi15, v2)
            case "\n\t":
                switch "a" {
                case "":
                    if (45 < 255) {
                        v9 -= 394
                        v2 = (^(v2))
                        switch v2 {
                        case -236:
                            switch 1 {
                            case -426:
                                v16 := GenS2{F0: -368, F1: -961}
                                _ = v16
                            case 0:
                                var v17 int = 2147483647
                                _ = v17
                            case 32768:
                                v2 += fi15
                            }
                        case 9223372036854775807:
                            fmt.Println(v2, fi15)
                        case -1:
                            fmt.Println(v9)
                        }
                    } else {
                        v9 = v9
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v19 := v4[(fi18 & 32767)]
                            var v20 int8 = 42
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                }
            }
            v2 += 42
            var v21 int8 = 42
            _ = v21
        } else {
            var v22 float64 = float64(1.5)
            var v23 uint16 = genFn0()
            _ = v22
            _ = v23
        }
        _ = fi15
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        switch 636 {
        case 9223372036854775807:
            for fi24 := 0; fi24 < 15; fi24++ {
                v9 = v9
                v25 := GenS1{F0: 42, F1: 42, F2: v9}
                var v26 int = (fi24 << ((v9) & 63))
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            v9 = 65535
        }
    case 4294967295:
        v27 := GenS2{F0: 864, F1: -8}
        _ = v27
    case 463:
        switch v2 {
        case 42:
            switch (v2 * (v2 - 360)) {
            case -62:
                if (!(!true)) {
                    switch "" {
                    case "a":
                        if (!true) {
                            var v28 string = "hello"
                            fmt.Println(v2)
                            _ = v28
                        } else {
                            v9 += (v9 - (v9 * v9))
                        }
                    }
                    switch 9223372036854775807 {
                    case 669:
                        fmt.Println(v2, v2)
                    default:
                        _ = len(v11)
                    }
                } else {
                    switch 9223372036854775807 {
                    case 42:
                        v29 := GenS3{F0: true}
                        _ = v29
                    case 190:
                        v2 += v2
                    case 32767:
                        var v30 uint = 1
                        _ = v30
                    }
                }
            }
        case 342:
            var v31 string = ("foo bar" + "a")
            _ = v31
        case 101:
            var v32 float32 = (1.5 * float32((-2.25 * 3.14159)))
            _ = v32
        }
    }
    fmt.Println(v9, v2)
    v9 = v9
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
