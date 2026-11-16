package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int32; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64, p1 string, p2 string) uint {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: p2, F1: 764}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(p0)
        p1 = ""
        fmt.Println(fi6, p2)
        _ = fi6
    }
    if (1 < 358) {
        var v7 int16 = 0
        p0 -= 9223372036854775807
        _ = v7
    } else {
        fmt.Println(p2, p2, p0)
    }
    v8 := GenS2{F0: 128, F1: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    v9 := GenS2{F0: 2147483647, F1: 240}
    var v10 uint8 = 128
    v11 := GenS1{F0: ("" + "\n\t"), F1: 0}
    fmt.Println(v10)
    if false {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 string = "foo bar"
            fmt.Println(fi12, v1)
            _ = fi12
            _ = v13
        }
        v1 = v1
        if false {
            if false {
                if false {
                    fmt.Println(v1, v1)
                } else {
                    var v14 uint64 = 574
                    _ = v14
                }
            } else {
                var v15 int64 = 1
                v10 = v10
                _ = v15
            }
            v10 = v10
            var v16 int64 = 1
            _ = v16
        }
    } else {
        var v17 uint64 = 42
        switch -9223372036854775808 {
        case 4294967295:
            switch (("\n\t" + ("" + (("foo bar" + (("\n\t" + "foo bar") + "")) + "hello"))) + ("hello" + "")) {
            case "foo bar":
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS2{F0: fi18, F1: uint8(v17)}
                    _ = fi18
                    _ = v19
                }
            default:
                if ((v10 + v10) != 255) {
                    switch 4294967295 {
                    case -1:
                        delete(v3, 655)
                    case 32768:
                        var v20 int = 127
                        _ = v20
                    case 42:
                        fmt.Println(v10, v1, v1)
                    }
                    var v21 int64 = 42
                    _ = v21
                } else {
                    delete(v2, 32767)
                    v1 = ((0.1 + (v1 / v1)) - float64(1e10))
                }
            }
        case -105:
            v22 := GenS1{F0: ("" + "foo bar"), F1: 0}
            _ = v22
        case 128:
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := GenS1{F0: (("foo bar" + ("\n\t" + "\n\t")) + "a"), F1: v17}
                v1 = (1e10 * (v1 - v1))
                _ = fi23
                _ = v24
            }
        default:
            v25 := []int{1, -451}
            _ = v25
        }
        v1 = (1e10 * (float64(3.14159) * (v1 - v1)))
        _ = v17
    }
    v26 := GenS2{F0: 564, F1: (v10 & v10)}
    v27 := GenS1{F0: "", F1: 4294967295}
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int64 = 1
        _ = fi28
        _ = v29
    }
    v30 := []int{-937, 24, 32768, 0}
    v31 := GenS3{F0: (v10 - (v10 - (v10 << ((2) & 63))))}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v32 := GenS4{F0: -2147483648}
    v33 := GenS3{F0: v10}
    switch "a" {
    case "foo bar":
        for fi34 := 0; fi34 < 19; fi34++ {
            v10 = 0
            _ = fi34
        }
    case "":
        fmt.Println(v1)
    case "\n\t":
        if true {
            if (4294967295 == 2) {
                v35 := GenS5{F0: 0, F1: -1, F2: 264}
                v36 := GenS6{F0: (2147483647 >= 91), F1: 65535}
                _ = v35
                _ = v36
            }
        }
    default:
        for fi37 := 0; fi37 < 8; fi37++ {
            fmt.Println(fi37, v10, v10)
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v9
}
