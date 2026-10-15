package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 int8) int8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        switch ("foo bar" + "") {
        case "":
            switch fi7 {
            case -1:
                var v8 int = -830
                _ = v8
            case 42:
                if true {
                    v9 := GenS1{F0: fi7}
                    _ = v9
                } else {
                    switch ("" + "a") {
                    case "a":
                        if (fi7 > fi7) {
                            fmt.Println(p1, p0)
                            v10 := []int{fi7, fi7, fi7, fi7}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v10
                        }
                    }
                    p1 = p1
                }
            case 127:
                var v11 uint16 = 42
                _ = v11
            default:
                fmt.Println(fi7, p0)
            }
        }
        p0 -= (p0 + p0)
        var v12 int16 = 1000
        _ = fi7
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(fi13)
            _ = fi14
        }
        v15 := GenS2{F0: -2147483648}
        _ = fi13
        _ = v15
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{4294967295}
    v3 := []uint32{4294967295, 2147483648, 2147483648, 65535}
    var v4 string = ("" + "a")
    var v5 int64 = 1
    var v6 int64 = (v5 - 42)
    fmt.Println(v6)
    var v16 int32 = 1
    v16 = v16
    fmt.Println(v6, v16)
    var v17 int8 = 100
    fmt.Println(v4, v6, v6)
    v6 = 1
    v18 := GenS1{F0: 2147483647}
    delete(v1, 2)
    var v19 uint64 = 42
    var v20 uint = 255
    _ = len(v1)
    var v21 string = v4
    v3 = append(v3, 622)
    v2 = append(v2, 42)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 string = ((v4 + "hello") + "foo bar")
        var v24 uint64 = (v19 / ((v19) | 1))
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
