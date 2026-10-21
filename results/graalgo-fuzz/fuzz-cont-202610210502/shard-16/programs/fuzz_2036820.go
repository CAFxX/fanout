package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int }

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    v4 := GenS1{F0: 1, F1: v3}
    v5 := GenS2{F0: 1, F1: 2}
    fmt.Println(v3)
    v1 = (1e10 * 1e10)
    for fi6 := 0; fi6 < 20; fi6++ {
        v1 = v1
        v1 -= (0.0 / 1.0)
        _ = fi6
    }
    v3 = 0
    switch "\n\t" {
    case "foo bar":
        var v7 uint64 = 0
        _ = v7
    default:
        fmt.Println(v1, v1)
    }
    v8 := GenS3{F0: 18446744073709551615}
    v1 -= v1
    v1 += (v1 - 0.0)
    v9 := GenS1{F0: -32768, F1: (-(v3))}
    if (65535 > 0) {
        v2[127] = 2147483647
        if false {
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v3)
                var v11 uint = 4294967295
                v11 += (v11 << ((v11) & 63))
                _ = fi10
                _ = v11
            }
        }
        if false {
            if ((true && (false && (true || (65535 >= 0)))) && true) {
                if (true && true) {
                    switch (("hello" + "") + "a") {
                    case "foo bar":
                        var v12 uint8 = 1
                        _ = v12
                    case "":
                        for fi13 := 0; fi13 < 10; fi13++ {
                            var v14 bool = (0.0 == 1.5)
                            _ = fi13
                            _ = v14
                        }
                    case "hello":
                        fmt.Println(v1, v3)
                    default:
                        v1 -= 0.0
                    }
                }
            }
            fmt.Println(v3, v3)
        }
    }
    if (((false && true) || false) || true) {
        var v15 uint64 = 18446744073709551615
        v16 := []int{589, 42, 0}
        for fi17 := 0; fi17 < 12; fi17++ {
            v3 = v3
            v18 := GenS2{F0: 2147483648, F1: 65535}
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    v1 = (v1 + 1e10)
    var v19 uint8 = 184
    var v20 string = ("a" + ("foo bar" + "\n\t"))
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 float64 = 0.1
        var v23 int = -135
        if true {
            var v24 bool = false
            v3 += 2147483647
            _ = v24
        } else {
            var v25 float64 = (v22 * float64(-2.25))
            v25 += 0.1
            switch (fi21 >> ((42) & 63)) {
            case 128:
                fmt.Println(v25, v3, v22)
            case 48:
                switch ("foo bar" + "hello") {
                case "hello":
                    if false {
                        v26 := []int{146}
                        _ = v26
                    }
                case "\n\t":
                    v25 = v1
                case "a":
                    v27 := GenS1{F0: -32768, F1: -394}
                    _ = v27
                default:
                    v28 := GenS4{F0: 0}
                    _ = v28
                }
            }
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    v29 := GenS5{F0: (-6 >= 54), F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
