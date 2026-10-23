package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 288
        v8 += 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint8 = 127
        var v11 int32 = 281
        switch (("hello" + ((("\n\t" + ("foo bar" + ((("foo bar" + "foo bar") + (("hello" + "") + (("hello" + "hello") + "\n\t"))) + ("hello" + (("hello" + "a") + ""))))) + ("a" + "")) + (("" + "hello") + ("foo bar" + "a")))) + ("foo bar" + ("" + "\n\t"))) {
        case "foo bar":
            switch -1 {
            case -1:
                v12 := GenS1{F0: 216, F1: 2, F2: -128}
                _ = v12
            default:
                if false {
                    var v13 int16 = -1
                    _ = v13
                } else {
                    v14 := GenS1{F0: 861, F1: 4294967295, F2: 127}
                    fmt.Println(v11)
                    _ = v14
                }
            }
        case "hello":
            var v15 uint8 = 127
            _ = v15
        case "a":
            v16 := GenS2{F0: 797}
            _ = v16
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 524 {
    case -646:
        var v17 float32 = (0.0 - 0.0)
        _ = v17
    case 128:
        v18 := []int{650}
        _ = v18
    default:
        var v19 int64 = -668
        _ = v19
    }
    return -69
}

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 91
    var v3 bool = true
    var v4 string = ""
    var v5 uint64 = 21
    var v6 int32 = -2147483648
    var v20 int8 = int8(v6)
    switch v4 {
    case "a":
        v5 = (v5 - 2)
    case "foo bar":
        var v21 int16 = 907
        _ = v21
    case "":
        v5 -= 18446744073709551615
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 uint32 = 65535
            v1 -= v1
            _ = fi22
            _ = v23
        }
    }
    v2 = v20
    if (float32(1e10) >= v1) {
        var v24 string = v4
        var v25 uint = 42
        _ = v24
        _ = v25
    }
    if false {
        fmt.Println(v2, v20)
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 int = (fi26 >> ((65535) & 63))
            var v28 int = (fi26 >> ((526) & 63))
            v28 = 32767
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v29 uint = 255
    if ((!(0.0 <= (((-2.25 * 1e10) - (1e10 - 0.0)) * 3.14159))) || (!v3)) {
        var v30 int8 = (v2 >> ((1) & 63))
        _ = v30
    }
    if true {
        var v31 uint32 = 65535
        if v3 {
            var v32 int8 = v2
            _ = v32
        } else {
            var v33 uint = 4294967295
            _ = v33
        }
        _ = v31
    } else {
        v2 = v2
        var v34 int = -748
        _ = v34
    }
    var v35 int64 = genFn0()
    v36 := []int{-1, -9223372036854775808, 32768, -574}
    var v37 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
