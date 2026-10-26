package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint }
func genFn0(p0 uint64) uint16 {
    var v5 float64 = 1e10
    switch "a" {
    case "hello":
        switch "a" {
        case "foo bar":
            p0 = 4
        case "":
            v6 := GenS1{F0: 42, F1: (true && false)}
            _ = v6
        }
    case "a":
        var v7 bool = (false && (false && (65535 > 65535)))
        _ = v7
    case "":
        if false {
            var v8 int = 96
            _ = v8
        } else {
            var v9 int16 = 0
            var v10 int16 = 536
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: 37, F1: true}
        p0 = (p0 >> ((364) & 63))
        fmt.Println(p0, fi11)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v5, fi13)
        var v14 uint16 = 0
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v5
    return 503
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 bool = (!(true || (0 <= 127)))
    var v4 int8 = 127
    var v15 float64 = float64(0.1)
    v3 = (v3 || (2 < 221))
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS2{F0: 461, F1: ("hello" + "foo bar")}
            v3 = ("" != "a")
            switch fi17 {
            case 128:
                var v19 uint = 1
                _ = v19
            case 32767:
                fmt.Println(fi16, v15, fi17)
            case 1:
                v4 = (v4 + 42)
            default:
                var v20 string = (("hello" + "") + ("a" + ("" + "hello")))
                _ = v20
            }
            if true {
                var v21 uint32 = 4294967295
                v22 := GenS3{F0: 0, F1: -2147483648}
                _ = v21
                _ = v22
            }
            fmt.Println(fi17, v3)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if true {
        delete(v1, 2147483647)
        var v23 int8 = v4
        _ = v23
    }
    v3 = v3
    var v24 uint32 = 4294967295
    v25 := GenS2{F0: -992, F1: ("a" + "foo bar")}
    v24 = (v24 % (((v24 % ((v24) | 1))) | 1))
    v4 = v4
    var v26 string = "hello"
    fmt.Println(v3, v15)
    v15 -= v15
    v27 := GenS1{F0: v4, F1: (v4 >= -20)}
    for fi28 := 0; fi28 < 10; fi28++ {
        fmt.Println(v26, v24)
        v29 := GenS4{F0: v24, F1: 262, F2: 2}
        var v30 float32 = (3.14159 * 1.5)
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 bool = ((v15 + 1e10) != float64(0.1))
        var v33 int32 = 457
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
