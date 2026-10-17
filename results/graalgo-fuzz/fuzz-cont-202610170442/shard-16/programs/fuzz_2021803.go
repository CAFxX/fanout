package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint32) uint16 {
    switch 175 {
    case 32768:
        p0 -= 2147483648
    case -406:
        p0 = p0
    case 9223372036854775807:
        fmt.Println(p0)
    }
    switch ("a" + "a") {
    case "\n\t":
        var v7 float64 = float64(0.0)
        _ = v7
    }
    p0 += p0
    if (32767 >= 2) {
        v8 := []int{4294967295, -300, 981, 530}
        _ = v8
    } else {
        var v9 float32 = float32(1e10)
        v10 := []int{32768, 9223372036854775807, 2147483647, 2147483647}
        _ = v10
        _ = v9
    }
    var v11 int8 = 1
    _ = p0
    _ = v11
    return 255
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 uint = 65535
    v5 := []uint32{2}
    v6 := make(map[string]int)
    var v12 string = (((((v1 + ("a" + ((v1 + v1) + v1))) + "") + ("" + "foo bar")) + v1) + v1)
    for fi13 := 0; fi13 < 11; fi13++ {
        v1 = ((v1 + (("\n\t" + ("hello" + v1)) + (v12 + ""))) + ((v12 + v12) + ("\n\t" + "a")))
        _ = fi13
    }
    switch ("hello" + (v1 + v12)) {
    case "":
        v14 := GenS1{F0: (true && (9223372036854775807 > 9223372036854775807))}
        _ = v14
    case "\n\t":
        v12 = (((("a" + v12) + "") + "a") + v1)
    }
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 8; fi16++ {
        v4 -= 0
        v17 := GenS2{F0: 4294967295, F1: 42}
        _ = fi16
        _ = v17
    }
    v4 = 65535
    for fi18 := 0; fi18 < 8; fi18++ {
        if true {
            v5 = append(v5, 65535)
            if false {
                v19 := GenS3{F0: -776, F1: (v12 + v1)}
                _ = v19
            } else {
                fmt.Println(v4, fi18)
            }
            var v20 uint16 = genFn0(0)
            _ = v20
        }
        var v21 float64 = ((0.1 / 1e10) / 1.5)
        _ = fi18
        _ = v21
    }
    if (((-1 >= 921) || (false || true)) || (1000 != 32767)) {
        v12 = ((v12 + v1) + v12)
        if (!((-947 <= 1) || false)) {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS2{F0: 809, F1: -2147483648}
                v1 = "\n\t"
                _ = fi22
                _ = v23
            }
            if (-2147483648 >= 0) {
                fmt.Println(v15, v12, v12)
                delete(v2, uint32(v4))
                var v24 string = v12
                _ = v24
            }
        }
    }
    v25 := GenS1{F0: (!(1 <= 127))}
    for fi26 := 0; fi26 < 11; fi26++ {
        v1 = v1
        _ = fi26
    }
    v27 := GenS1{F0: false}
    if true {
        switch 9223372036854775807 {
        case 49:
            v1 = "hello"
        }
        v6[(((v1 + "") + ("" + (v1 + v12))) + "hello")] = 32767
    } else {
        var v28 int64 = 1
        if true {
            var v29 uint32 = 0
            _ = v29
        } else {
            fmt.Println(v1)
            _ = len(v3)
            v30 := GenS4{F0: 1}
            _ = v30
        }
        var v31 bool = (-121 == -115)
        _ = v28
        _ = v31
    }
    v32 := GenS2{F0: 18446744073709551615, F1: -2147483648}
    v33 := GenS5{F0: -128, F1: 2, F2: 42}
    v1 = ("a" + (v1 + v1))
    v4 += 556
    switch "hello" {
    case "a":
        if true {
            fmt.Println(v1)
            if true {
                for fi34 := 0; fi34 < 17; fi34++ {
                    v35 := GenS6{F0: 4294967295, F1: 127, F2: -128}
                    var v36 uint64 = 496
                    var v37 float32 = float32(0.0)
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            var v38 int8 = 100
            _ = v38
        }
    case "":
        var v39 int32 = 2147483647
        _ = v39
    case "hello":
        v15 -= ((((0.1 / ((v15 - v15) * 1.5)) / v15) * v15) * (1e10 + -2.25))
    default:
        var v40 float64 = (float64(1e10) - 0.0)
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        v42 := GenS3{F0: -9223372036854775808, F1: (v1 + v12)}
        v43 := GenS7{F0: (v4 / ((v4) | 1)), F1: 42}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
