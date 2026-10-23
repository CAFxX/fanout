package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint }

func main() {
    var v1 int32 = 0
    var v2 float64 = (-2.25 - 1e10)
    var v3 int32 = v1
    var v4 uint = 65535
    var v5 string = "foo bar"
    v6 := make(map[int]bool)
    if ("a" > (v5 + "")) {
        var v7 int8 = int8(v3)
        _ = v7
    }
    v8 := GenS1{F0: "foo bar", F1: 100}
    v2 = v2
    var v9 float32 = (-2.25 * -2.25)
    if true {
        v10 := GenS1{F0: "", F1: 42}
        switch v5 {
        case "\n\t":
            var v11 int16 = -32768
            _ = v11
        case "hello":
            v12 := GenS2{F0: true, F1: v3, F2: 0}
            _ = v12
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 = (-(v3))
            _ = fi13
        }
    }
    fmt.Println(v2, v4, v2)
    var v14 int8 = -1
    fmt.Println(v1)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int64 = -403
        v5 = ((("a" + (v5 + "\n\t")) + "foo bar") + "hello")
        var v17 uint16 = 2
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = len(v6)
    var v18 uint64 = 0
    v2 = v2
    switch v5 {
    case "hello":
        v19 := []int{-487, 490, -386}
        _ = v19
    case "foo bar":
        v20 := GenS2{F0: (-731 >= 0), F1: -2147483648, F2: (v18 * (v18 >> ((1) & 63)))}
        _ = v20
    }
    if (!false) {
        v21 := GenS1{F0: v5, F1: v14}
        v22 := v6[42]
        var v23 int16 = 1000
        _ = v21
        _ = v22
        _ = v23
    } else {
        fmt.Println(v14, v3)
        var v24 int64 = -895
        _ = v24
    }
    v18 -= v18
    var v25 float64 = v2
    for fi26 := 0; fi26 < 20; fi26++ {
        v9 = (3.14159 - v9)
        if true {
            v27 := []int{495, 32767, fi26, fi26}
            var v28 uint8 = 42
            _ = v27
            _ = v28
        } else {
            if true {
                v3 += (v1 << ((v4) & 63))
                if true {
                    switch (v5 + "a") {
                    case "\n\t":
                        var v29 float32 = v9
                        _ = v29
                    case "":
                        v18 = uint64(v4)
                    case "foo bar":
                        var v30 bool = true
                        _ = v30
                    }
                } else {
                    if false {
                        var v31 int64 = 0
                        v18 -= v18
                        _ = v31
                    }
                    switch fi26 {
                    case 0:
                        var v32 int64 = -9223372036854775808
                        _ = v32
                    }
                }
            }
            v33 := GenS3{F0: int16(v1), F1: (^(v4))}
            _ = v33
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
