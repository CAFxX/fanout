package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 uint) int {
    switch "a" {
    case "hello":
        if false {
            for fi4 := 0; fi4 < 16; fi4++ {
                v5 := GenS1{F0: 448}
                var v6 uint8 = 128
                _ = fi4
                _ = v5
                _ = v6
            }
            var v7 uint64 = 42
            _ = v7
        }
    case "":
        p0 += 1
    default:
        v8 := GenS1{F0: 0}
        _ = v8
    }
    if true {
        switch "foo bar" {
        case "":
            if (!((false || (-2147483648 != 675)) || (!true))) {
                v9 := GenS2{F0: 1}
                v10 := []int{4294967295, 32768}
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 13; fi11++ {
                    p0 = 2
                    _ = fi11
                }
                var v12 float32 = (3.14159 * 3.14159)
                var v13 string = "a"
                _ = v12
                _ = v13
            }
        case "a":
            v14 := []int{-561}
            _ = v14
        default:
            p0 = p0
        }
        switch "" {
        case "":
            var v15 float64 = 0.1
            _ = v15
        }
    }
    v16 := []int{32767}
    p0 = (^(p0))
    _ = p0
    _ = v16
    return 127
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]uint32)
    v3 := []string{("hello" + ("hello" + "\n\t")), ("\n\t" + "a"), "\n\t"}
    switch "foo bar" {
    case "hello":
        var v17 int32 = -773
        _ = v17
    }
    if true {
        if true {
            fmt.Println(v1)
        }
        fmt.Println(v1)
        var v18 int32 = 36
        _ = v18
    }
    var v19 uint64 = 0
    var v20 int8 = -1
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint16 = 255
        v23 := GenS1{F0: 1000}
        var v24 int = 127
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    v1 = v1
    var v25 int8 = v20
    v2[-1] = 249
    fmt.Println(v20, v1)
    v25 = v20
    fmt.Println(v25, v20)
    delete(v2, genFn0(v1))
    var v26 float32 = 0.0
    switch ("foo bar" + (("\n\t" + "hello") + "a")) {
    case "":
        fmt.Println(v20)
    case "\n\t":
        var v27 int = 2147483647
        _ = v27
    case "foo bar":
        switch 627 {
        case 32767:
            var v28 uint8 = 128
            _ = v28
        case 2147483647:
            if false {
                var v29 float64 = 1e10
                fmt.Println(v1, v1, v1)
                var v30 string = ("\n\t" + "")
                _ = v29
                _ = v30
            } else {
                if true {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v26 += (-2.25 * float32(-2.25))
                        _ = fi31
                    }
                    var v32 int8 = -1
                    v2[-340] = 219
                    _ = v32
                } else {
                    v1 += 0
                }
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v33 string = ("foo bar" + ("foo bar" + "hello"))
    if true {
        v34 := GenS3{F0: 275, F1: 127, F2: 255}
        for fi35 := 0; fi35 < 18; fi35++ {
            fmt.Println(v20, fi35)
            v36 := GenS4{F0: 1, F1: 240, F2: fi35}
            fmt.Println(v33)
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v33
}
