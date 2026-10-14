package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint; F1 string }
func genFn0() uint32 {
    v6 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: 56}
    var v7 uint = 42
    var v8 bool = false
    switch "foo bar" {
    case "hello":
        v9 := []int{}
        _ = v9
    case "":
        v10 := GenS2{F0: 255, F1: 42}
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 900
}
func genFn1() uint8 {
    var v11 uint = 4294967295
    if true {
        var v12 uint8 = 255
        switch ("hello" + "") {
        case "":
            var v13 uint32 = 0
            _ = v13
        }
        _ = v12
    } else {
        v11 = v11
        var v14 int8 = 48
        _ = v14
    }
    _ = v11
    return 39
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 128
    var v3 int32 = 2147483647
    v4 := make(map[uint32]string)
    var v5 bool = (!(!(82 <= 18446744073709551615)))
    delete(v4, 65535)
    var v15 string = "hello"
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 255, F1: v3}
        var v18 uint32 = 2147483648
        v3 = v3
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint = 1
    if v5 {
        if v5 {
            var v20 float32 = -2.25
            v21 := v4[42]
            _ = v20
            _ = v21
        }
        if (321 <= -1) {
            v19 -= v19
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 float32 = 1.5
                _ = fi22
                _ = v23
            }
            v2 -= v2
        } else {
            var v24 bool = v5
            var v25 uint16 = 253
            _ = v24
            _ = v25
        }
    }
    v26 := GenS3{F0: (v19 * (v19 >> ((v19) & 63))), F1: v15}
    v2 = v2
    switch "" {
    case "foo bar":
        v5 = v5
    case "hello":
        v27 := []int{-2147483648}
        _ = v27
    default:
        for fi28 := 0; fi28 < 19; fi28++ {
            v5 = ((((v15 + "a") + ("a" + v15)) + "") == "foo bar")
            v29 := []int{758, fi28, (fi28 >> ((364) & 63)), fi28}
            _ = len(v29)
            _ = fi28
            _ = v29
        }
    }
    var v30 uint16 = 42
    var v31 uint8 = (v2 * 127)
    v32 := GenS2{F0: (v30 - 65535), F1: (v3 % (((^(v3))) | 1))}
    v30 = (v30 | v30)
    var v33 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
