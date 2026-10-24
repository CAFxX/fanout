package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint, p2 int64) int32 {
    v5 := GenS1{F0: 1, F1: 4294967295, F2: -649}
    var v6 bool = false
    var v7 uint8 = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        if (1 >= 224) {
            v7 = (v7 / ((v7) | 1))
            v9 := []int{(fi8 - fi8), (fi8 | fi8)}
            var v10 int16 = (^(p0))
            _ = v10
            _ = v9
        }
        if true {
            var v11 uint32 = 42
            _ = v11
        }
        if v6 {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch (((("foo bar" + "\n\t") + "\n\t") + (("\n\t" + "hello") + "foo bar")) + "\n\t") {
                case "hello":
                    fmt.Println(p2)
                default:
                    v13 := []int{-1, (fi8 / ((2147483647) | 1))}
                    _ = v13
                }
                fmt.Println(p0, fi8)
                v14 := GenS1{F0: -839, F1: 353, F2: -1}
                _ = fi12
                _ = v14
            }
            var v15 float32 = 1e10
            switch fi8 {
            case -253:
                for fi16 := 0; fi16 < 10; fi16++ {
                    fmt.Println(v15, fi8, fi16)
                    p1 = (p1 | (^(p1)))
                    v17 := []int{fi8, (fi16 - (-(fi16))), fi16, (fi16 % (((fi16 - 9223372036854775807)) | 1))}
                    _ = fi16
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS2{F0: 1}
                    v7 += v7
                    p1 += p1
                    _ = fi18
                    _ = v19
                }
            }
            _ = v15
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 bool = false
    v2 := []int{542}
    var v3 float32 = float32(-2.25)
    v4 := make(map[string]string)
    var v20 uint32 = 0
    if v1 {
        fmt.Println(v20, v1)
        v3 = v3
    } else {
        var v21 float32 = (1e10 - v3)
        var v22 int32 = 0
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 string = "hello"
            v4["foo bar"] = (v24 + "a")
            var v25 uint64 = 359
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v21
        _ = v22
    }
    switch (((("" + "") + "\n\t") + ("" + ("a" + ""))) + ("a" + "\n\t")) {
    case "foo bar":
        var v26 int8 = 0
        _ = v26
    }
    v27 := GenS2{F0: 255}
    switch 0 {
    case -306:
        switch -677 {
        case 32767:
            var v28 uint64 = 0
            _ = v28
        case -928:
            v29 := GenS2{F0: 65535}
            _ = v29
        }
    }
    var v30 uint = 2
    var v31 float64 = ((0.1 / float64(v3)) * (0.1 * (float64(float32(((0.1 + 0.0) + (1e10 * (0.1 + 0.0))))) + float64(-2.25))))
    if false {
        switch 128 {
        case 736:
            v31 = 1e10
        case 0:
            fmt.Println(v20, v3)
        default:
            var v32 uint32 = (v20 + 531)
            _ = v32
        }
    }
    if true {
        fmt.Println(v3)
    }
    var v33 int8 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
