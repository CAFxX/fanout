package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 uint16) uint16 {
    p0 = (p0 / ((p0) | 1))
    v5 := GenS1{F0: true, F1: (p0 ^ (p0 / ((716) | 1)))}
    var v6 uint8 = 127
    _ = p0
    _ = v5
    _ = v6
    return (p0 / (((p0 | 0)) | 1))
}
func genFn1(p0 string, p1 string) uint8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(p1, fi7, fi7)
            var v9 int32 = 333
            if false {
                switch "a" {
                case "foo bar":
                    var v10 float32 = 1.5
                    _ = v10
                default:
                    var v11 int8 = 42
                    _ = v11
                }
            } else {
                fmt.Println(v9, p0, fi8)
                var v12 uint64 = 18446744073709551615
                _ = v12
            }
            switch (fi8 % ((32767) | 1)) {
            case 9223372036854775807:
                v13 := []int{fi7, fi7, (fi7 / (((fi7 - -498)) | 1)), fi8}
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        v14 := []int{}
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint8 = 0
            var v17 float64 = 1.5
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi7
        _ = v14
    }
    fmt.Println(p0, p0)
    var v18 int = -213
    v19 := []int{(v18 / (((-(v18))) | 1)), 4294967295, 9223372036854775807, (^(v18))}
    var v20 string = p0
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return 128
}

func main() {
    v1 := []uint32{1, 2147483647, 42}
    v2 := make(map[int]uint32)
    v3 := []int{32768, -978}
    v4 := make(map[int]bool)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -245 {
    case 128:
        var v21 float64 = float64((1.5 + float32(-2.25)))
        _ = v21
    case -9223372036854775808:
        if true {
            if (2 >= 2) {
                var v22 bool = (0 > -394)
                _ = v22
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(fi23)
                    var v24 int8 = -128
                    _ = fi23
                    _ = v24
                }
                v25 := GenS2{F0: 64, F1: 4294967295, F2: -1}
                _ = v25
            }
        }
    }
    switch ((("hello" + "hello") + ("foo bar" + "foo bar")) + ("" + "hello")) {
    case "foo bar":
        var v26 uint64 = 51
        _ = v26
    case "hello":
        var v27 uint = 710
        _ = v27
    case "":
        if (412 < 1) {
            if false {
                var v28 int8 = -26
                for fi29 := 0; fi29 < 8; fi29++ {
                    _ = len(v1)
                    v30 := GenS3{F0: 989, F1: true}
                    var v31 uint8 = 128
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                var v32 uint = 1
                _ = v28
                _ = v32
            } else {
                var v33 uint64 = 1
                var v34 float64 = (float64(1.5) / -2.25)
                switch ("" + ("\n\t" + ("hello" + "hello"))) {
                case "":
                    fmt.Println(v34, v33)
                case "foo bar":
                    var v35 float32 = (3.14159 / 1.5)
                    _ = v35
                }
                _ = v33
                _ = v34
            }
        }
    }
    var v36 int64 = 42
    fmt.Println(v36)
    fmt.Println(v36)
    var v37 int = 948
    for fi38 := 0; fi38 < 14; fi38++ {
        fmt.Println(v36, v36)
        _ = fi38
    }
    fmt.Println(v36, v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
