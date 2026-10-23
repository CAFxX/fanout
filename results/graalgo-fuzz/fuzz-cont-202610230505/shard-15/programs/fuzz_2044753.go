package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 bool }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0() int16 {
    if true {
        switch -871 {
        case 0:
            v5 := GenS1{F0: "\n\t", F1: 42, F2: 185}
            _ = v5
        default:
            var v6 int = 127
            _ = v6
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            if ((0 > 1) || true) {
                if (1 <= 1) {
                    v8 := []int{fi7, (fi7 << ((65535) & 63)), fi7}
                    var v9 uint32 = 65535
                    _ = v8
                    _ = v9
                } else {
                    var v10 float64 = 0.1
                    _ = v10
                }
                v11 := GenS2{F0: 63}
                v12 := GenS2{F0: 128}
                _ = v11
                _ = v12
            }
            var v13 bool = ((788 > 9) || (true && true))
            _ = fi7
            _ = v13
        }
    }
    var v14 uint32 = 569
    if ("foo bar" <= "hello") {
        var v15 uint16 = 599
        _ = v15
    }
    switch -406 {
    case 128:
        switch "foo bar" {
        case "foo bar":
            v16 := []int{-479, 2147483647}
            _ = v16
        }
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        v17 := GenS2{F0: 45}
        _ = v17
    case "":
        if (!(true || (255 < 32767))) {
            v18 := GenS3{F0: 42, F1: "\n\t", F2: false}
            var v19 int16 = 1000
            v20 := []int{-387, 2147483647}
            _ = v18
            _ = v19
            _ = v20
        }
    case "a":
        fmt.Println(v14)
    }
    _ = v14
    return -1
}
func genFn1(p0 int32, p1 int32, p2 bool) uint {
    p0 += -2147483648
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        p2 = false
    default:
        fmt.Println(p2, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 969
}

func main() {
    var v1 string = ("\n\t" + ("" + "a"))
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 uint16 = 65535
    if len(v3) > 0 {
        v3[0] = 0
    }
    v21 := GenS2{F0: 47}
    v22 := GenS2{F0: 153}
    v1 = ("foo bar" + v1)
    v1 = ("" + (v1 + v1))
    for fi23 := 0; fi23 < 7; fi23++ {
        if true {
            for fi24 := 0; fi24 < 8; fi24++ {
                _ = len(v2)
                var v25 float32 = 1.5
                if true {
                    fmt.Println(fi23, fi24, fi23)
                }
                _ = fi24
                _ = v25
            }
        } else {
            var v26 float32 = 1.5
            var v27 uint8 = 2
            _ = v26
            _ = v27
        }
        var v28 int8 = -91
        v29 := GenS3{F0: -9223372036854775808, F1: v1, F2: (!true)}
        _ = fi23
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v1 = v1
        _ = fi30
    }
    switch -300 {
    case -2147483648:
        v1 = ("hello" + "hello")
    }
    fmt.Println(v1, v1)
    fmt.Println(v4, v4)
    v4 -= 42
    v31 := GenS4{F0: 35, F1: -481}
    v1 = "hello"
    if ((!(("foo bar" + v1) <= "\n\t")) && (false || true)) {
        for fi32 := 0; fi32 < 4; fi32++ {
            if false {
                var v33 int8 = 97
                var v34 uint32 = 0
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 17; fi35++ {
                v1 = (((v1 + v1) + "foo bar") + v1)
                _ = fi35
            }
            v1 = v1
            _ = fi32
        }
    } else {
        _ = len(v3)
    }
    var v36 uint = 1
    delete(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v36
    _ = v4
}
