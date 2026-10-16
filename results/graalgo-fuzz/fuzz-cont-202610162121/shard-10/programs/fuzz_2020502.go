package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 int64 }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0(p0 int32, p1 string) bool {
    if false {
        v7 := GenS1{F0: ((p1 + p1) + (p1 + p1))}
        v8 := []int{}
        v9 := GenS1{F0: p1}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        switch "\n\t" {
        case "foo bar":
            v11 := []int{}
            _ = v11
        case "hello":
            switch ("a" + p1) {
            case "foo bar":
                switch (fi10 | 141) {
                case 42:
                    v12 := GenS2{F0: p0}
                    _ = v12
                default:
                    v13 := []int{fi10}
                    _ = v13
                }
            case "\n\t":
                p1 = p1
            }
        case "":
            v14 := GenS1{F0: "a"}
            _ = v14
        }
        _ = fi10
    }
    fmt.Println(p0, p0)
    p0 -= (p0 | p0)
    p1 = p1
    _ = p0
    _ = p1
    return (-332 >= 62)
}
func genFn1(p0 int64, p1 uint) int64 {
    var v15 uint = (p1 | (p1 - p1))
    v16 := GenS1{F0: "hello"}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return (p0 >> ((v15) & 63))
}

func main() {
    var v1 int16 = -635
    v2 := []uint32{4294967295, 705, 573, 2147483648}
    v3 := []int{32767, -1, -363, 32768}
    var v4 string = (("\n\t" + "hello") + "hello")
    var v5 int32 = -839
    var v6 float32 = 0.0
    v17 := GenS3{F0: 42, F1: 67, F2: 770}
    var v18 int16 = (-(v1))
    if true {
        _ = len(v3)
        if false {
            fmt.Println(v6)
        } else {
            var v19 int16 = v1
            v4 = "a"
            _ = v19
        }
        var v20 int32 = -2147483648
        _ = v20
    }
    fmt.Println(v1)
    var v21 uint = 258
    var v22 int16 = 42
    v3 = append(v3, 1)
    v23 := GenS4{F0: 0, F1: (v5 + v5), F2: -97}
    fmt.Println(v1, v4, v18)
    var v24 uint8 = 0
    for fi25 := 0; fi25 < 10; fi25++ {
        switch fi25 {
        case 32767:
            for fi26 := 0; fi26 < 16; fi26++ {
                v21 += v21
                var v27 bool = (127 <= -128)
                _ = fi26
                _ = v27
            }
        case 128:
            fmt.Println(v5, v6, v18)
        }
        v24 = 0
        _ = fi25
    }
    fmt.Println(v4, v4)
    if (!true) {
        switch v4 {
        case "":
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 uint = 922
                _ = fi28
                _ = v29
            }
        case "a":
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS2{F0: v5}
                _ = len(v2)
                _ = fi30
                _ = v31
            }
        default:
            v32 := GenS3{F0: 65535, F1: v24, F2: -375}
            _ = v32
        }
    } else {
        v33 := GenS4{F0: v22, F1: v5, F2: 0}
        v34 := GenS5{F0: 1, F1: 1}
        fmt.Println(v24, v18, v18)
        _ = v33
        _ = v34
    }
    var v35 int = 1
    v21 = (v21 / ((v21) | 1))
    v6 = v6
    if false {
        switch (v4 + (v4 + "\n\t")) {
        case "a":
            fmt.Println(v1, v35, v35)
        case "hello":
            var v36 float64 = 0.1
            _ = v36
        }
        var v37 int16 = 32767
        _ = v37
    } else {
        if true {
            var v38 uint8 = v24
            _ = v38
        }
        fmt.Println(v1, v18)
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
