package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 uint; F2 int16 }
type GenS6 struct { F0 int64; F1 uint16 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        var v7 uint = 255
        v8 := []int{-243}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 float32 = 3.14159
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 292, F1: -9223372036854775808, F2: 449}
    if true {
        v12 := []int{-2147483648}
        var v13 string = ("foo bar" + "hello")
        _ = v12
        _ = v13
    }
    _ = v11
    return -32768
}
func genFn1() uint8 {
    var v14 int64 = -9223372036854775808
    v15 := []int{-9223372036854775808, -283, 2147483647}
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 uint32 = 42
    v5 := make(map[string]string)
    for fi16 := 0; fi16 < 20; fi16++ {
        if false {
            switch 127 {
            case -669:
                var v17 uint8 = 1
                _ = v17
            default:
                v4 = v4
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                _ = len(v5)
                v19 := []int{(fi18 ^ 128), 32768, fi16}
                _ = fi18
                _ = v19
            }
            v4 -= v4
        }
        switch "\n\t" {
        case "":
            v20 := GenS2{F0: 80, F1: 345}
            _ = v20
        case "hello":
            var v21 string = ""
            _ = v21
        default:
            _ = len(v1)
        }
        v22 := []int{2147483647, fi16, (fi16 & fi16)}
        _ = fi16
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        delete(v3, 127)
        if (!(!true)) {
            for fi24 := 0; fi24 < 13; fi24++ {
                if false {
                    v2 -= 42
                }
                var v25 uint = 300
                v26 := GenS3{F0: ("\n\t" + (("hello" + "\n\t") + ((("a" + (("" + (((("hello" + ("a" + ("hello" + "a"))) + "foo bar") + (("hello" + ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))) + "\n\t")) + "hello")) + "foo bar")) + ("\n\t" + "")) + "hello")))}
                var v27 uint16 = 32767
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 bool = true
            _ = v28
        } else {
            delete(v5, "a")
        }
        _ = fi23
    }
    if true {
        v29 := v1[163]
        _ = v29
    } else {
        var v30 int64 = -952
        v31 := []int{9223372036854775807, int(v2), 4294967295}
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: v4}
    v33 := []int{128, 42}
    v2 = v2
    fmt.Println(v2, v2)
    v2 = (^(v2))
    v2 += v2
    switch 1 {
    case 42:
        fmt.Println(v2, v2)
    }
    var v34 bool = false
    if (36 < 1) {
        v4 = (v4 % ((uint32(v2)) | 1))
        switch "a" {
        case "a":
            for fi35 := 0; fi35 < 2; fi35++ {
                var v36 uint8 = 0
                var v37 float64 = 1e10
                v4 = (^(v4))
                _ = fi35
                _ = v36
                _ = v37
            }
        case "foo bar":
            fmt.Println(v4, v34)
        default:
            var v38 int32 = 1
            _ = v38
        }
    } else {
        v39 := GenS5{F0: 8, F1: 255, F2: -32768}
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 bool = (1 <= -775)
            switch (fi40 * fi40) {
            case -1:
                v41 = v41
            }
            delete(v3, fi40)
            _ = fi40
            _ = v41
        }
        if v34 {
            v42 := GenS6{F0: -1, F1: 32767}
            _ = v42
        }
        _ = v39
    }
    v4 = v4
    if len(v33) > 0 {
        v33[0] = -415
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
