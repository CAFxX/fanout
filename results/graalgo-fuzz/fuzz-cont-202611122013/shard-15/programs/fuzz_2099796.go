package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
func genFn0(p0 uint64, p1 uint, p2 int32) int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(p0, p0)
        switch "\n\t" {
        case "a":
            fmt.Println(p1)
        case "foo bar":
            v5 := GenS1{F0: 9223372036854775807}
            _ = v5
        default:
            if false {
                var v6 uint32 = 66
                _ = v6
            } else {
                p2 = (p2 << ((139) & 63))
                p1 += 803
                p2 = p2
            }
        }
        _ = fi4
    }
    if (!false) {
        if (!("hello" <= ("" + "hello"))) {
            var v7 float32 = 3.14159
            v8 := []int{}
            v8 = append(v8, -1)
            _ = v7
            _ = v8
        } else {
            v9 := GenS2{F0: p1}
            var v10 uint64 = (p0 | (p0 >> ((0) & 63)))
            _ = v10
            _ = v9
        }
        var v11 string = "foo bar"
        _ = v11
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []string{}
    v3 := []int{992, -806, 32768, 9223372036854775807}
    v1 = 0
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int = fi12
        fmt.Println(v13)
        _ = fi12
        _ = v13
    }
    v3 = append(v3, 42)
    var v14 int64 = 790
    var v15 float32 = ((3.14159 + -2.25) + (1e10 * float32(0.1)))
    if len(v2) > 0 {
        v2[0] = (("hello" + ("a" + "")) + "foo bar")
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v15 -= (v15 * ((v15 - (v15 * v15)) / (v15 + (3.14159 - (v15 / 1.5)))))
        var v17 int32 = 2147483647
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float32 = v15
        v14 -= (v14 & (-(v14)))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int16 = 1000
        switch (fi20 & (fi20 | (fi20 << ((255) & 63)))) {
        case -9223372036854775808:
            v1 = 908
        default:
            switch fi20 {
            case 4294967295:
                v1 += (v1 - 1)
            case -9223372036854775808:
                if false {
                    v21 = v21
                } else {
                    var v22 uint64 = v1
                    _ = v22
                }
            }
        }
        if (3.14159 != (((v15 / (v15 * v15)) + float32(0.0)) / 1.5)) {
            v15 -= 3.14159
        } else {
            switch (fi20 << ((1) & 63)) {
            case 9223372036854775807:
                v14 += (v14 ^ v14)
            default:
                switch int(v21) {
                case 128:
                    var v23 float32 = v15
                    _ = v23
                case 23:
                    fmt.Println(v21)
                }
            }
        }
        _ = fi20
        _ = v21
    }
    var v24 int16 = 1
    if (42 != 1) {
        if true {
            fmt.Println(v24, v1, v14)
        } else {
            v25 := GenS3{F0: "a"}
            if len(v2) > 0 {
                v2[0] = (("hello" + (("" + (("hello" + "") + (("a" + ("a" + "")) + "\n\t"))) + "a")) + "hello")
            }
            _ = v25
        }
        var v26 int32 = -957
        _ = v26
    } else {
        var v27 string = "\n\t"
        v24 = 32767
        if false {
            var v28 float64 = (1e10 * 1e10)
            _ = v28
        }
        _ = v27
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
}
