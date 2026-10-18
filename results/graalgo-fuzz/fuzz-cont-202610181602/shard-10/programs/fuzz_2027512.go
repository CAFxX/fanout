package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS7 struct { F0 bool }
func genFn0(p0 int) uint {
    v7 := GenS1{F0: 2}
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS2{F0: 9223372036854775807}
        for fi10 := 0; fi10 < 7; fi10++ {
            switch 32767 {
            case -774:
                v11 := GenS2{F0: 0}
                _ = v11
            case 32767:
                v12 := GenS3{F0: 1, F1: fi8, F2: 1}
                _ = v12
            case 42:
                v13 := []int{fi10, (fi10 % ((32768) | 1)), p0, 1}
                _ = v13
            default:
                var v14 int64 = 1
                _ = v14
            }
            var v15 int8 = 1
            v16 := GenS4{F0: 1, F1: 0}
            _ = fi10
            _ = v15
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    var v17 string = "foo bar"
    if (!false) {
        v18 := []int{p0, (-(p0)), (p0 << ((421) & 63))}
        _ = v18
    } else {
        var v19 int32 = 42
        v19 += v19
        _ = v19
    }
    _ = p0
    _ = v17
    _ = v7
    return 255
}
func genFn1(p0 uint, p1 uint8) string {
    p0 = 42
    v20 := []int{-9223372036854775808}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v20
    return "foo bar"
}

func main() {
    var v1 int16 = -373
    var v2 int = 87
    v3 := []int{v2, (v2 >> ((65535) & 63))}
    var v4 uint16 = 390
    v5 := make(map[string]bool)
    var v6 int = v2
    v2 = (-(v2))
    v21 := GenS5{F0: 114, F1: "foo bar"}
    if (42 == int32(v6)) {
        v22 := GenS6{F0: 255, F1: v1, F2: 255}
        var v23 uint32 = 2147483648
        var v24 uint32 = v23
        _ = v22
        _ = v23
        _ = v24
    } else {
        v5[""] = true
        v2 = 42
    }
    v1 += v1
    if (false || false) {
        fmt.Println(v6, v2)
        var v25 string = ((("a" + "a") + ("foo bar" + (("\n\t" + "hello") + "hello"))) + "foo bar")
        var v26 int32 = int32(v1)
        _ = v25
        _ = v26
    } else {
        if true {
            v27 := GenS6{F0: 42, F1: (v1 * v1), F2: 1}
            v2 = (v2 + 127)
            _ = v27
        } else {
            switch ("hello" + "") {
            case "foo bar":
                v1 += v1
            case "\n\t":
                v6 -= 32768
            case "":
                var v28 int8 = 1
                _ = v28
            default:
                v6 = 9223372036854775807
            }
            switch (v6 % ((int(v4)) | 1)) {
            case -186:
                var v29 int16 = (v1 >> ((161) & 63))
                _ = v29
            case 0:
                v1 += -423
            }
        }
        v6 -= (v2 / ((v6) | 1))
    }
    if (true || false) {
        fmt.Println(v4, v6, v2)
    }
    switch (v6 ^ (v6 >> ((genFn0(v6)) & 63))) {
    case 4294967295:
        fmt.Println(v2, v1)
    case 32768:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := GenS7{F0: false}
                v6 += v6
                v2 = (v2 + (v2 % ((v6) | 1)))
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    case "\n\t":
        v1 = (v1 & 1)
    }
    v4 = 387
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 int8 = -1
        delete(v5, "")
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
