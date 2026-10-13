package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int32 }
func genFn0(p0 int32) uint {
    if (841 <= 663) {
        var v4 int32 = p0
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 string = (("hello" + ("hello" + "")) + ("" + ""))
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 0, F1: 1, F2: -9223372036854775808}
        _ = v4
        _ = v7
    } else {
        v8 := GenS1{F0: -1, F1: 335, F2: 919}
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        p0 = (p0 >> ((255) & 63))
    case "a":
        var v9 int32 = p0
        _ = v9
    }
    _ = p0
    return 255
}
func genFn1(p0 string) uint16 {
    fmt.Println(p0)
    var v10 uint16 = 42
    _ = p0
    _ = v10
    return v10
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 0
    v3 := make(map[uint32]string)
    for fi11 := 0; fi11 < 1; fi11++ {
        switch "foo bar" {
        case "":
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint16 = genFn1(("hello" + ("a" + ("a" + ""))))
                fmt.Println(fi12, fi12)
                _ = fi12
                _ = v13
            }
        case "hello":
            var v14 uint32 = 4294967295
            _ = v14
        case "a":
            v2 = 524
        default:
            var v15 int64 = 0
            _ = v15
        }
        var v16 uint8 = 42
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int64 = 0
            v19 := GenS2{F0: 86, F1: 127}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi11
        _ = v16
    }
    if (v1 == (v1 ^ (-(v1)))) {
        v20 := GenS3{F0: 0}
        var v21 uint64 = (v2 * 4294967295)
        _ = v20
        _ = v21
    }
    v22 := []int{}
    var v23 uint8 = 126
    v24 := GenS1{F0: -53, F1: 1, F2: 0}
    fmt.Println(v1)
    if (uint16(v23) >= 2) {
        if true {
            var v25 int64 = 381
            v23 -= v23
            _ = v25
        }
        fmt.Println(v2, v1, v2)
        fmt.Println(v23, v1, v23)
    } else {
        var v26 uint8 = 255
        fmt.Println(v26)
        var v27 uint16 = 2
        _ = v26
        _ = v27
    }
    var v28 string = "\n\t"
    if false {
        var v29 uint = 330
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 bool = (0 == 1)
            var v32 int8 = -1
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            if false {
                v34 := GenS4{F0: ((fi33 & fi33) >= fi33), F1: 1}
                _ = v34
            }
            if false {
                fmt.Println(v2, v28, v28)
                fmt.Println(v29, v28)
                fmt.Println(v28, v23)
            } else {
                fmt.Println(v1, v2)
                v29 = (v29 | genFn0(551))
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 int8 = 1
                var v37 uint8 = 2
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi33
        }
        _ = v29
    }
    if (false || (!((!(true && true)) || (65535 != genFn1("hello"))))) {
        fmt.Println(v28, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
}
