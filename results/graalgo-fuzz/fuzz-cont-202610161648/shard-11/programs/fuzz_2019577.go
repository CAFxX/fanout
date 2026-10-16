package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64) uint32 {
    if (3.14159 > ((((float64((-2.25 / 1e10)) * 3.14159) - 1.5) - ((float64(1.5) - float64(1e10)) * 3.14159)) * ((0.1 - 1e10) * 3.14159))) {
        var v4 float64 = ((0.1 / float64(float32((1e10 + 3.14159)))) + (0.0 * ((float64((-2.25 + float32(float64(3.14159)))) / -2.25) + ((float64(((0.0 / ((1.5 + -2.25) - 0.1)) / 0.1)) / 1e10) - 0.1))))
        _ = v4
    }
    switch 127 {
    case 42:
        v5 := []int{977, 9223372036854775807}
        _ = v5
    case 9223372036854775807:
        fmt.Println(p0)
    default:
        var v6 float32 = 1e10
        _ = v6
    }
    _ = p0
    return 13
}

func main() {
    var v1 int8 = -83
    var v2 bool = (0 == -577)
    var v3 int16 = 32767
    v7 := []int{690}
    v2 = (!(false || (1 >= 4294967295)))
    var v8 int16 = (v3 / ((v3) | 1))
    var v9 bool = true
    if (4294967295 <= 4294967295) {
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                v11 := GenS1{F0: 255}
                var v12 int32 = 2147483647
                var v13 uint = 106
                _ = v11
                _ = v12
                _ = v13
            }
            v2 = v2
            var v14 int64 = 242
            _ = fi10
            _ = v14
        }
        var v15 float64 = 0.0
        _ = v15
    } else {
        v1 += v1
    }
    fmt.Println(v1)
    if true {
        if (!(!v2)) {
            var v16 string = ""
            var v17 int64 = -1
            _ = v16
            _ = v17
        }
        if (42 >= 1) {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = float32((1.5 * float64(0.0)))
                var v20 string = (("" + ("\n\t" + ("foo bar" + (("a" + "foo bar") + (("hello" + "hello") + ("\n\t" + ("a" + "hello"))))))) + "hello")
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            switch fi21 {
            case 585:
                v1 -= 95
            }
            v22 := GenS2{F0: "\n\t", F1: (fi21 % ((fi21) | 1)), F2: "\n\t"}
            var v23 int = (fi21 - (fi21 ^ (fi21 % ((1) | 1))))
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 int8 = (v1 & v1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v7 = append(v7, -857)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint = 279
        if (("hello" >= "hello") && (v9 && (!v2))) {
            v2 = true
        }
        if v2 {
            v8 = -785
            var v27 float64 = (((1e10 / 1e10) / 0.1) + 1e10)
            _ = v27
        } else {
            for fi28 := 0; fi28 < 11; fi28++ {
                v29 := GenS3{F0: 0, F1: int8(v8)}
                _ = fi28
                _ = v29
            }
        }
        _ = fi25
        _ = v26
    }
    v30 := GenS2{F0: "foo bar", F1: 128, F2: ""}
    v8 = v3
    if v2 {
        if v2 {
            var v31 uint64 = 42
            v1 -= v1
            _ = len(v7)
            _ = v31
        } else {
            v7 = append(v7, 32768)
            var v32 uint8 = 128
            _ = v32
        }
        fmt.Println(v3, v3)
        _ = len(v7)
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS3{F0: 42, F1: v1}
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v7
    _ = v8
    _ = v9
}
