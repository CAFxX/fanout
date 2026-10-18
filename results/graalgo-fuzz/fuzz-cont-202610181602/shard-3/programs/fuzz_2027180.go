package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 int, p1 uint64, p2 uint32) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 1}
        p2 = (p2 << ((0) & 63))
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: 1}
            var v8 string = "\n\t"
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    p1 = (p1 | 18446744073709551615)
    _ = p0
    _ = p1
    _ = p2
    return (323 >= 456)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 uint32 = 1
    v3 = 4294967295
    v3 -= (v3 >> ((371) & 63))
    delete(v2, 2147483647)
    if (true || true) {
        if (!(255 != 66)) {
            if false {
                if (false && false) {
                    var v9 int8 = 42
                    switch "hello" {
                    case "\n\t":
                        switch 4294967295 {
                        case 0:
                            fmt.Println(v3, v3)
                        case -2147483648:
                            var v10 int64 = -345
                            _ = v10
                        default:
                            var v11 float64 = (1e10 + float64(1.5))
                            _ = v11
                        }
                    case "hello":
                        if (430 <= 2) {
                            if false {
                                var v12 uint16 = 0
                                switch 32768 {
                                case 32767:
                                    v13 := []int{}
                                    _ = v13
                                case 4294967295:
                                    switch ((("a" + ("" + ("" + "\n\t"))) + ("a" + ("hello" + "foo bar"))) + "") {
                                    case "foo bar":
                                        var v14 float32 = (((float32(1e10) * 3.14159) / 1.0) * 0.1)
                                        _ = v14
                                    case "hello":
                                        v3 -= v3
                                    }
                                }
                                _ = v12
                            }
                            switch "hello" {
                            case "hello":
                                v15 := []int{619, -718, 32768, -785}
                                _ = v15
                            }
                            var v16 uint16 = 0
                            _ = v16
                        } else {
                            v3 -= v3
                        }
                    default:
                        fmt.Println(v9, v9)
                    }
                    _ = v9
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    fmt.Println(v3)
                    v3 = 2147483647
                    var v18 float32 = ((0.1 / ((1e10 - 1.5) - (-2.25 + 3.14159))) - 0.1)
                    _ = fi17
                    _ = v18
                }
                var v19 string = (("\n\t" + "hello") + ("" + ("a" + "\n\t")))
                _ = v19
            } else {
                v20 := v1[(v3 / ((v3) | 1))]
                _ = v20
            }
            v21 := GenS1{F0: -1}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 string = "foo bar"
                _ = fi22
                _ = v23
            }
            delete(v1, v3)
        }
        var v24 uint16 = 42
        v24 -= v24
        _ = v24
    } else {
        fmt.Println(v3)
        v25 := GenS2{F0: 4294967295}
        v3 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        fmt.Println(v3, fi26)
        var v27 float64 = (1e10 / (-2.25 - 0.1))
        var v28 bool = (0 < 1)
        _ = fi26
        _ = v27
        _ = v28
    }
    switch "" {
    case "hello":
        v3 -= v3
    case "foo bar":
        v3 = v3
    }
    v29 := GenS3{F0: ((217 < 321) && (true && (4294967295 != 4294967295))), F1: 381}
    var v30 int64 = -477
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
