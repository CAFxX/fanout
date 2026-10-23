package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 bool; F1 int8; F2 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 uint16) int64 {
    switch "hello" {
    case "\n\t":
        v7 := GenS1{F0: 32767, F1: "hello", F2: true}
        _ = v7
    case "foo bar":
        p0 = (p0 >> ((4294967295) & 63))
    case "hello":
        v8 := []int{-269, 685, 4294967295}
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = -47
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 42
    var v5 int = 32767
    var v6 string = "hello"
    if (false && true) {
        _ = len(v3)
        var v11 bool = (("\n\t" + ((v6 + "a") + v6)) > ((v6 + v6) + ((v6 + (((v6 + v6) + ((((v6 + v6) + (v6 + ("foo bar" + "foo bar"))) + "") + v6)) + v6)) + ("\n\t" + "hello"))))
        _ = v11
    }
    var v12 int8 = -31
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: (fi13 + 176), F1: "\n\t", F2: true}
        switch (v5 & (v5 | v2)) {
        case 4294967295:
            var v15 float64 = 0.1
            _ = v15
        case 127:
            v16 := []int{(fi13 >> ((2) & 63)), 127, -7, 0}
            _ = v16
        case 200:
            var v17 int64 = 9223372036854775807
            _ = v17
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 float32 = ((float32(0.1) * (3.14159 + 3.14159)) * 1.5)
                _ = fi18
                _ = v19
            }
        }
        _ = fi13
        _ = v14
    }
    var v20 string = "\n\t"
    switch (v2 & 2147483647) {
    case 42:
        switch 128 {
        case 9223372036854775807:
            var v21 int16 = -1
            _ = v21
        case 127:
            v5 = (v5 | v2)
        }
    }
    fmt.Println(v4, v4)
    v3[789] = 65535
    if true {
        if (128 != 1) {
            fmt.Println(v1)
        }
        if true {
            v12 = -55
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                _ = len(v3)
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int8 = -1
                    v25 := GenS1{F0: v5, F1: v20, F2: (!true)}
                    v26 := GenS2{F0: (42 > 1), F1: v24, F2: fi23}
                    var v27 float32 = (1e10 - 1.5)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v4 = v4
                _ = fi22
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                switch (v20 + (v20 + (v20 + v20))) {
                case "foo bar":
                    v29 := GenS3{F0: 1}
                    _ = v29
                case "\n\t":
                    fmt.Println(v2, v6)
                case "hello":
                    var v30 int = v5
                    _ = v30
                }
                var v31 int64 = 0
                switch (v5 | fi28) {
                case -1:
                    v5 -= (-(v5))
                default:
                    v32 := []int{}
                    _ = v32
                }
                _ = fi28
                _ = v31
            }
        }
    }
    var v33 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
