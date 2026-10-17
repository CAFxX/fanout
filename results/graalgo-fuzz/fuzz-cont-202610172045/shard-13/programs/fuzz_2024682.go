package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 int, p2 int) uint {
    p1 -= 32768
    fmt.Println(p0, p1, p2)
    var v4 uint = 1
    v5 := GenS1{F0: 2147483647, F1: 740}
    if (1 < 346) {
        switch (p2 % ((p1) | 1)) {
        case 1:
            p0 = -1
        case 4294967295:
            for fi6 := 0; fi6 < 6; fi6++ {
                p1 = 483
                _ = fi6
            }
        case 2147483647:
            v7 := GenS1{F0: -35, F1: 4294967295}
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                if ((42 <= -2147483648) || true) {
                    switch "" {
                    case "hello":
                        v10 := GenS2{F0: (((!false) && ((fi9 / (((^(p2))) | 1)) <= fi9)) && (!(!false)))}
                        _ = v10
                    case "\n\t":
                        var v11 int64 = 991
                        _ = v11
                    default:
                        var v12 int32 = -420
                        _ = v12
                    }
                    var v13 int64 = 0
                    p0 = (^(p0))
                    _ = v13
                }
                if true {
                    v14 := GenS1{F0: 1, F1: 4294967295}
                    _ = v14
                } else {
                    p2 = (p2 ^ fi9)
                }
                _ = fi9
            }
            _ = fi8
        }
        v15 := []int{p2, 32768, p1}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v4 - (v4 >> ((v4) & 63)))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    fmt.Println(v2)
    if true {
        var v16 int32 = v2
        v17 := GenS1{F0: (v16 | (v16 / ((v2) | 1))), F1: 2147483647}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 int32 = v2
                if true {
                    v21 := GenS1{F0: (v2 | v2), F1: 801}
                    fmt.Println(fi18)
                    v20 -= (v2 << ((323) & 63))
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            v22 := GenS2{F0: ("hello" >= "\n\t")}
            _ = fi18
            _ = v22
        }
    }
    fmt.Println(v2)
    if false {
        var v23 string = (("a" + (("\n\t" + "\n\t") + "hello")) + ("a" + "foo bar"))
        _ = v23
    }
    switch (("\n\t" + "\n\t") + ("hello" + "")) {
    case "hello":
        for fi24 := 0; fi24 < 11; fi24++ {
            _ = len(v1)
            _ = fi24
        }
    case "":
        switch 128 {
        case 1:
            var v25 float32 = 1.5
            _ = v25
        case 32768:
            v26 := GenS1{F0: v2, F1: 2147483647}
            _ = v26
        case 510:
            if (1 <= 42) {
                var v27 int16 = 0
                v3[42] = 0
                _ = v27
            } else {
                var v28 uint64 = 501
                _ = v28
            }
        }
    }
    if ((v2 / ((0) | 1)) == v2) {
        v2 = -1
        v29 := []int{1, 32767}
        _ = v29
    }
    var v30 int16 = 0
    v30 = v30
    switch "hello" {
    case "hello":
        v2 = v2
    default:
        var v31 uint = 2
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        if (v30 < 280) {
            var v33 string = "hello"
            var v34 string = (v33 + v33)
            _ = v33
            _ = v34
        } else {
            for fi35 := 0; fi35 < 13; fi35++ {
                v2 = 2147483647
                v2 = v2
                _ = fi35
            }
        }
        _ = fi32
    }
    v2 += -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
