package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0() uint8 {
    var v6 int16 = 1000
    var v7 int = 42
    switch (v7 ^ 9223372036854775807) {
    case -897:
        if true {
            var v8 uint = 2
            v8 = (^(v8))
            _ = v8
        }
    case 32767:
        v9 := []int{v7}
        _ = v9
    case -1:
        v6 = v6
    }
    fmt.Println(v7, v6)
    v10 := GenS1{F0: 42, F1: "hello"}
    _ = v10
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    v3 := []string{}
    var v4 int = -1
    v5 := []string{"foo bar", "hello", "\n\t"}
    switch (v4 + (v4 + v4)) {
    case 32767:
        v11 := GenS1{F0: 0, F1: "hello"}
        _ = v11
    default:
        var v12 uint16 = 404
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: 2147483647, F1: "a"}
        v15 := GenS1{F0: 0, F1: "hello"}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = "a"
    for fi17 := 0; fi17 < 8; fi17++ {
        switch v4 {
        case -9223372036854775808:
            var v18 bool = v2
            _ = v18
        }
        _ = fi17
    }
    if (42 != 65535) {
        var v19 int = (v4 / ((v4) | 1))
        _ = v19
    }
    v1[263] = 2147483648
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v2, fi20)
        var v21 uint8 = 0
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 0
    for fi23 := 0; fi23 < 18; fi23++ {
        if (!(0.0 == -2.25)) {
            var v24 float32 = 0.0
            _ = v24
        } else {
            switch "hello" {
            case "a":
                v22 -= (^(v22))
            }
        }
        var v25 float64 = 1.5
        var v26 int32 = 42
        _ = fi23
        _ = v25
        _ = v26
    }
    v16 = v16
    v16 = "foo bar"
    var v27 int16 = -924
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
