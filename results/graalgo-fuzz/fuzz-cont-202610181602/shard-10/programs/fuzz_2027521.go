package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 uint32, p1 int16, p2 uint) uint64 {
    if true {
        v4 := GenS1{F0: -1, F1: p2, F2: p1}
        fmt.Println(p1, p1)
        p0 -= 4294967295
        _ = v4
    } else {
        p1 = 690
        v5 := []int{32768, -9223372036854775808, 1, 26}
        fmt.Println(p2, p0)
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 1, F2: -32768}
    p0 = (p0 & p0)
    switch "a" {
    case "foo bar":
        p1 -= (p1 / ((p1) | 1))
    case "":
        p2 = p2
    default:
        var v7 bool = (!true)
        _ = v7
    }
    switch "" {
    case "hello":
        v8 := []int{0, 128}
        _ = v8
    default:
        v9 := []int{0, 9223372036854775807, 42}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32) uint {
    fmt.Println(p0)
    v10 := GenS2{F0: 42}
    var v11 uint16 = 72
    _ = p0
    _ = v10
    _ = v11
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32767
    v3 := make(map[string]string)
    if (("a" + ("" + "")) <= ("" + ("foo bar" + ("a" + "a")))) {
        v2 += (-(v2))
    }
    v2 -= (v2 | 128)
    v2 = v2
    switch ("a" + ("foo bar" + "")) {
    case "foo bar":
        v2 += 32768
    case "hello":
        if (1 == 0) {
            var v12 float32 = 0.0
            fmt.Println(v12, v2)
            _ = v12
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                switch ((("hello" + ("hello" + "\n\t")) + "hello") + "hello") {
                case "hello":
                    for fi14 := 0; fi14 < 12; fi14++ {
                        fmt.Println(fi13, fi14, fi13)
                        v15 := GenS3{F0: 32767, F1: 4294967295, F2: 0}
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    var v16 float32 = 0.0
                    _ = v16
                case "":
                    delete(v3, "\n\t")
                }
                var v17 uint = genFn1(1)
                _ = fi13
                _ = v17
            }
            var v18 float32 = 3.14159
            v1["hello"] = true
            _ = v18
        }
    }
    v2 = v2
    fmt.Println(v2)
    var v19 int8 = -55
    var v20 float64 = -2.25
    v21 := GenS3{F0: 2, F1: 18446744073709551615, F2: 0}
    if (true && (0 < 0)) {
        var v22 float64 = (v20 + (-2.25 / 0.1))
        _ = v22
    } else {
        var v23 float32 = (0.0 * (1e10 * ((float32(1e10) * (-2.25 - 3.14159)) / (3.14159 + -2.25))))
        _ = v23
    }
    var v24 uint64 = 18446744073709551615
    switch ("\n\t" + "hello") {
    case "a":
        for fi25 := 0; fi25 < 1; fi25++ {
            switch "hello" {
            case "":
                v24 += 4294967295
            case "a":
                for fi26 := 0; fi26 < 15; fi26++ {
                    if false {
                        _ = len(v1)
                    } else {
                        fmt.Println(fi26, v2)
                        var v27 int64 = 1
                        v24 = 2
                        _ = v27
                    }
                    var v28 int64 = 766
                    v2 -= fi26
                    _ = fi26
                    _ = v28
                }
            }
            v29 := GenS2{F0: 103}
            var v30 int64 = 42
            _ = fi25
            _ = v29
            _ = v30
        }
    case "\n\t":
        v31 := GenS4{F0: int16(v2)}
        _ = v31
    case "hello":
        v20 -= 3.14159
    default:
        if (0 < -1) {
            v32 := []int{}
            _ = v32
        } else {
            v24 = (v24 | v24)
            switch ((("hello" + "a") + "a") + "a") {
            case "a":
                v24 -= genFn0(1, 42, 18446744073709551615)
            case "foo bar":
                var v33 int8 = v19
                _ = v33
            default:
                v34 := GenS5{F0: 233, F1: 42, F2: (v20 < v20)}
                _ = v34
            }
            v20 = float64((0.1 / 1.0))
        }
    }
    if ("hello" >= "a") {
        fmt.Println(v19, v20)
    } else {
        v24 -= (^(v24))
        for fi35 := 0; fi35 < 8; fi35++ {
            v2 -= fi35
            _ = fi35
        }
        var v36 int32 = -907
        _ = v36
    }
    v24 = v24
    v19 -= (v19 << ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
}
