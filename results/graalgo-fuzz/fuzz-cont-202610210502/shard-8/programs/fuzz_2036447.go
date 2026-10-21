package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int) bool {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint8 = 173
        _ = fi5
        _ = v6
    }
    v7 := []int{-9223372036854775808, p0, p0}
    var v8 int = p0
    _ = p0
    _ = v7
    _ = v8
    return (42 != 100)
}
func genFn1() int16 {
    v9 := GenS1{F0: 33}
    var v10 bool = true
    fmt.Println(v10)
    _ = v10
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!true)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v2 = (("foo bar" + ("\n\t" + "foo bar")) > ("\n\t" + "foo bar"))
    if false {
        var v11 uint16 = 1
        _ = v11
    } else {
        switch 32768 {
        case 128:
            if (1.5 != 0.0) {
                var v12 uint32 = 42
                _ = v12
            }
        case 369:
            if false {
                if true {
                    v2 = (!true)
                    v2 = v2
                } else {
                    v13 := []int{-432}
                    v2 = ((769 == 4294967295) && false)
                    _ = len(v3)
                    _ = v13
                }
                v14 := GenS1{F0: -1}
                _ = v14
            }
        case 454:
            _ = len(v3)
        default:
            fmt.Println(v2)
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint64 = 441
            _ = fi15
            _ = v16
        }
    }
    if true {
        var v17 uint8 = 2
        if v2 {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 int16 = 42
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    }
    switch 0 {
    case 9223372036854775807:
        var v20 uint = 65535
        _ = v20
    case 2147483647:
        if true {
            v21 := GenS1{F0: 127}
            for fi22 := 0; fi22 < 17; fi22++ {
                if ((fi22 * fi22) > fi22) {
                    fmt.Println(v2, fi22)
                    delete(v1, 65535)
                }
                var v23 uint32 = 65535
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            if (32768 == 0) {
                _ = len(v1)
            } else {
                v2 = v2
                fmt.Println(v2)
                var v24 uint16 = 2
                _ = v24
            }
            var v25 float64 = float64(-2.25)
            _ = v25
        }
    case 32767:
        v2 = true
    default:
        var v26 uint64 = 954
        _ = v26
    }
    fmt.Println(v2)
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(v2, fi27)
        v28 := GenS1{F0: 1}
        _ = fi27
        _ = v28
    }
    var v29 int64 = -1
    v30 := []int{-1}
    if false {
        if v2 {
            fmt.Println(v2, v2)
            var v31 uint = 81
            _ = v31
        }
        v29 += (v29 / ((v29) | 1))
    } else {
        v32 := GenS2{F0: v29}
        _ = v32
    }
    switch 32768 {
    case -9223372036854775808:
        switch "hello" {
        case "hello":
            v33 := GenS1{F0: 64}
            _ = v33
        default:
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int = fi34
                switch 128 {
                case 9223372036854775807:
                    var v36 int = (-(v35))
                    _ = v36
                case 4294967295:
                    var v37 int8 = 1
                    _ = v37
                case -67:
                    var v38 float64 = ((1.5 - 1e10) * 0.1)
                    _ = v38
                default:
                    delete(v1, 4294967295)
                }
                _ = fi34
                _ = v35
            }
        }
    case 128:
        switch "" {
        case "hello":
            var v39 int = 0
            _ = v39
        }
    default:
        for fi40 := 0; fi40 < 19; fi40++ {
            if (!v2) {
                fmt.Println(v2, v29)
                switch fi40 {
                case 1:
                    var v41 string = "\n\t"
                    _ = v41
                case -9223372036854775808:
                    var v42 int16 = -130
                    _ = v42
                }
                var v43 uint64 = 18446744073709551615
                _ = v43
            }
            v2 = true
            var v44 string = "hello"
            _ = fi40
            _ = v44
        }
    }
    if v2 {
        v29 += -411
        _ = len(v30)
    } else {
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int64 = v29
            _ = fi45
            _ = v46
        }
        fmt.Println(v2, v2)
        if v2 {
            var v47 float32 = (1.5 - (-2.25 + 0.0))
            v2 = true
            fmt.Println(v2)
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
