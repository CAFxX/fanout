package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 int64) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 uint = 255
        _ = fi4
        _ = v5
    }
    fmt.Println(p1)
    for fi6 := 0; fi6 < 6; fi6++ {
        if (((^(fi6)) >= -570) && true) {
            if true {
                fmt.Println(p0, fi6, fi6)
                v7 := GenS1{F0: false, F1: 1, F2: 255}
                _ = v7
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(p0, fi6, p0)
                }
                v8 := []int{1, (^(fi6)), 127}
                _ = v8
            }
        }
        if (!true) {
            if (32767 <= 2) {
                var v9 float64 = 1e10
                _ = v9
            } else {
                for fi10 := 0; fi10 < 18; fi10++ {
                    switch "hello" {
                    case "hello":
                        var v11 int64 = 909
                        _ = v11
                    case "foo bar":
                        switch (("" + "") + ("" + "")) {
                        case "":
                            v12 := []int{}
                            _ = v12
                        case "\n\t":
                            v13 := GenS2{F0: -462, F1: 888}
                            _ = v13
                        }
                    }
                    v14 := GenS1{F0: (502 > 533), F1: 2147483647, F2: 65535}
                    _ = fi10
                    _ = v14
                }
                v15 := GenS3{F0: 0}
                v16 := []int{fi6, 201}
                _ = v15
                _ = v16
            }
        }
        _ = fi6
    }
    p0 = p0
    _ = p0
    _ = p1
    return 42
}
func genFn1() int64 {
    if true {
        var v17 uint = 0
        var v18 uint32 = uint32(v17)
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint8 = 1
            var v21 int = fi19
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 int8 = -128
    var v23 float32 = 1.5
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]string)
    var v3 uint = 806
    var v24 uint8 = 2
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 float32 = ((0.0 / 0.1) - (-2.25 - (0.0 * 1.5)))
        var v27 uint16 = 2
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{-1}
    for fi29 := 0; fi29 < 4; fi29++ {
        fmt.Println(fi29, v3, fi29)
        var v30 uint8 = v24
        var v31 float32 = 1e10
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 uint = (v3 * v3)
    v3 = (v3 | 2)
    v28 = append(v28, 837)
    if false {
        var v33 uint = v32
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 uint64 = 4294967295
            if true {
                var v36 bool = ((fi34 % ((fi34) | 1)) > (fi34 & fi34))
                v36 = (!v36)
                fmt.Println(v35)
                _ = v36
            } else {
                var v37 float64 = 0.0
                _ = v37
            }
            v35 = v35
            _ = fi34
            _ = v35
        }
        v3 += v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v32
}
