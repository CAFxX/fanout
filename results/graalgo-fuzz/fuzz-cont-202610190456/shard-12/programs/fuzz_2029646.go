package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
func genFn0(p0 int16) uint32 {
    switch -1 {
    case 1:
        v5 := GenS1{F0: -9223372036854775808}
        _ = v5
    default:
        fmt.Println(p0)
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        switch (fi6 * fi6) {
        case 127:
            var v7 uint = 396
            _ = v7
        }
        v8 := GenS2{F0: 959, F1: "a"}
        _ = fi6
        _ = v8
    }
    p0 = p0
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int = fi9
        fmt.Println(v10)
        if false {
            v11 := GenS1{F0: -9223372036854775808}
            var v12 uint64 = 0
            switch (fi9 ^ 127) {
            case 128:
                switch "hello" {
                case "hello":
                    var v13 uint16 = 65535
                    _ = v13
                case "":
                    v14 := []int{}
                    _ = v14
                default:
                    v12 += 42
                }
            case 127:
                v15 := []int{(fi9 ^ v10), (v10 - -238)}
                _ = v15
            default:
                v16 := []int{}
                _ = v16
            }
            _ = v11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{(fi17 - 128), (fi17 >> ((372) & 63)), fi17}
        p0 = (p0 / ((p0) | 1))
        var v19 int8 = 46
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    return 976
}
func genFn1(p0 int8, p1 uint16) string {
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    fmt.Println(p0, p0)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int8 = (p0 << ((65535) & 63))
        _ = fi22
        _ = v23
    }
    v24 := []int{102, 840}
    _ = p0
    _ = p1
    _ = v24
    return ("" + "foo bar")
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{0, 177, 112}
    var v3 int32 = -1
    var v4 uint8 = 45
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 float32 = v1
            _ = fi26
            _ = v27
        }
        fmt.Println(v1, v3, v1)
        var v28 uint32 = 65535
        _ = fi25
        _ = v28
    }
    _ = len(v2)
    v29 := GenS3{F0: 570, F1: 115, F2: genFn0(1000)}
    fmt.Println(v4)
    for fi30 := 0; fi30 < 19; fi30++ {
        v3 += int32(v4)
        _ = fi30
    }
    v31 := GenS2{F0: 411, F1: "\n\t"}
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 int16 = -100
        _ = fi32
        _ = v33
    }
    _ = len(v2)
    v34 := GenS2{F0: 18446744073709551615, F1: "foo bar"}
    v35 := GenS3{F0: 4294967295, F1: 1000, F2: 4294967295}
    v36 := GenS2{F0: 18446744073709551615, F1: ("a" + "")}
    var v37 int16 = 1
    if false {
        if true {
            var v38 uint8 = v4
            _ = v38
        }
        var v39 int32 = (v3 >> ((18446744073709551615) & 63))
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
