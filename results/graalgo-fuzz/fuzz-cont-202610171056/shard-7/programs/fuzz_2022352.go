package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0(p0 uint, p1 uint32, p2 bool) uint64 {
    v4 := GenS1{F0: "", F1: "\n\t", F2: -882}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = 1
        switch fi5 {
        case -9223372036854775808:
            var v7 uint = 168
            _ = v7
        case 128:
            var v8 string = ((("\n\t" + ("hello" + "a")) + "") + "a")
            _ = v8
        case 32767:
            fmt.Println(p2, v6, fi5)
        }
        var v9 int8 = -63
        _ = fi5
        _ = v6
        _ = v9
    }
    if ((p2 && (0.1 != -2.25)) || (1 < -489)) {
        for fi10 := 0; fi10 < 18; fi10++ {
            p0 = p0
            fmt.Println(p0)
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := []int{fi10, (^(fi10))}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    if p2 {
        fmt.Println(p1, p0)
    } else {
        if ((-1 > -1) || p2) {
            var v13 int = 2147483647
            _ = v13
        } else {
            fmt.Println(p2, p0, p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint64(p0)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 uint = 42
    v14 := GenS2{F0: ("foo bar" + "a"), F1: -1, F2: 0}
    v15 := []int{1, 127, -872, -2147483648}
    var v16 int8 = 0
    _ = len(v1)
    v17 := v1["\n\t"]
    var v18 int16 = 74
    for fi19 := 0; fi19 < 15; fi19++ {
        if (1 > 0) {
            switch fi19 {
            case -9223372036854775808:
                if true {
                    var v20 int = (fi19 * fi19)
                    _ = v20
                } else {
                    if len(v15) > 0 {
                        v15[0] = int(v16)
                    }
                    fmt.Println(v17, v2, fi19)
                    var v21 uint16 = 571
                    _ = v21
                }
            case -885:
                v17 = (v17 && v17)
            }
            var v22 uint = (v3 - 647)
            _ = v22
        }
        v23 := GenS1{F0: "foo bar", F1: "", F2: 642}
        var v24 int8 = v16
        _ = fi19
        _ = v23
        _ = v24
    }
    v2 -= (-(v2))
    v25 := GenS3{F0: 128, F1: 4294967295, F2: 97}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
