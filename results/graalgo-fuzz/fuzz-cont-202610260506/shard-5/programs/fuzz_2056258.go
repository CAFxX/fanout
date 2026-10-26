package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    v1 := []int{}
    var v2 string = "a"
    var v3 int8 = 42
    var v4 float32 = 0.0
    v5 := []int{-815, 1, 0}
    switch 433 {
    case -782:
        if true {
            _ = len(v5)
            if false {
                switch v2 {
                case "\n\t":
                    fmt.Println(v4, v3, v4)
                case "a":
                    _ = len(v5)
                case "hello":
                    v6 := GenS1{F0: 73, F1: 255}
                    _ = v6
                default:
                    switch "foo bar" {
                    case "":
                        switch "hello" {
                        case "\n\t":
                            v5 = append(v5, -652)
                        default:
                            var v7 int16 = 869
                            _ = v7
                        }
                    case "a":
                        if (65535 >= 0) {
                            v3 += v3
                            v8 := GenS1{F0: v3, F1: 249}
                            _ = v8
                        }
                    case "hello":
                        var v9 uint64 = 0
                        _ = v9
                    default:
                        v4 = 1e10
                    }
                }
                var v10 uint16 = 2
                var v11 bool = (true && false)
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: v3, F1: 128}
            _ = v12
        } else {
            v13 := GenS1{F0: v3, F1: 151}
            _ = v13
        }
    case -891:
        var v14 int16 = 0
        _ = v14
    }
    if (1 != 32767) {
        var v15 int64 = -1
        v15 += (-(v15))
        v3 -= (v3 + v3)
        _ = v15
    }
    _ = len(v5)
    v4 -= v4
    fmt.Println(v4, v2)
    fmt.Println(v2)
    v4 += 1e10
    var v16 float32 = v4
    v3 -= (v3 >> ((2) & 63))
    var v17 int64 = 42
    switch v2 {
    case "\n\t":
        var v18 uint16 = 42
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                var v20 string = v2
                if ((0.0 * (-2.25 * 3.14159)) != float64(v16)) {
                    var v21 string = ("hello" + "")
                    _ = v21
                }
                _ = v20
            }
            _ = fi19
        }
    }
    fmt.Println(v16)
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            switch -1 {
            case 0:
                _ = len(v1)
            case 42:
                fmt.Println(v4, fi23, v4)
            case -53:
                fmt.Println(v17)
                switch "hello" {
                case "hello":
                    v4 = 0.1
                case "":
                    var v24 int = (fi22 + -1)
                    _ = v24
                case "foo bar":
                    fmt.Println(v4)
                }
                var v25 uint64 = 18446744073709551615
                _ = v25
            }
            v26 := GenS1{F0: 31, F1: 127}
            if (!true) {
                switch (fi23 / ((fi22) | 1)) {
                case 2147483647:
                    fmt.Println(v3, v2, v16)
                case -472:
                    v27 := GenS1{F0: v3, F1: 128}
                    var v28 int32 = -1
                    v29 := GenS2{F0: 2, F1: 966}
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi23
            _ = v26
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            fmt.Println(v16, v3, v2)
            var v31 int = fi30
            _ = fi30
            _ = v31
        }
        if true {
            fmt.Println(v3, v2, v4)
        }
        _ = fi22
    }
    v5 = append(v5, -716)
    switch -2147483648 {
    case 128:
        var v32 int16 = 0
        _ = v32
    case -299:
        v33 := GenS3{F0: -1, F1: 0}
        _ = v33
    }
    v34 := GenS2{F0: 811, F1: 2}
    var v35 uint8 = 255
    var v36 int8 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
