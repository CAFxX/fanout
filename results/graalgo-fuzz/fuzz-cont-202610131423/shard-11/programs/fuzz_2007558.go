package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(fi4)
        var v5 uint = 2
        v6 := []int{32767, fi4, fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint = 107
    if (-419 != -9223372036854775808) {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(v7)
            fmt.Println(fi8, v7)
            v7 += (^(v7))
            _ = fi8
        }
    }
    if false {
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(fi9, fi9)
            if (-869 < 409) {
                var v10 string = ((("a" + ("foo bar" + "a")) + "") + "a")
                _ = v10
            }
            _ = fi9
        }
        var v11 uint32 = 25
        _ = v11
    }
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("foo bar" + ("hello" + "hello"))
    v3 := []int{1, 32768}
    v12 := GenS1{F0: 42, F1: genFn0()}
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint64 = 4294967295
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 1
        _ = fi13
        _ = v16
    }
    v2 = (v2 + "\n\t")
    var v17 float64 = 1e10
    v17 += (v17 * 3.14159)
    v17 += ((float64(float32(v17)) - ((v17 - float64(0.0)) * v17)) + (0.0 * v17))
    _ = len(v3)
    switch "foo bar" {
    case "foo bar":
        var v18 int16 = 32767
        _ = v18
    case "a":
        fmt.Println(v2, v17)
    case "":
        var v19 string = v2
        _ = v19
    default:
        v17 += v17
    }
    v17 = v17
    var v20 float64 = (0.0 * -2.25)
    var v21 int32 = 547
    v22 := GenS1{F0: -2147483648, F1: 1}
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: 127, F1: 42}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 0, F1: 1}
    switch ("foo bar" + ((v2 + (((v2 + v2) + "hello") + v2)) + "a")) {
    case "":
        var v26 string = v2
        _ = v26
    case "hello":
        fmt.Println(v21, v2)
    }
    switch 127 {
    case 673:
        fmt.Println(v2)
    case 42:
        v27 := GenS1{F0: 9223372036854775807, F1: 119}
        _ = v27
    default:
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 6; fi29++ {
                v2 = "foo bar"
                _ = len(v1)
                var v30 bool = true
                _ = fi29
                _ = v30
            }
            switch (^(fi28)) {
            case 4294967295:
                v31 := GenS2{F0: 4294967295, F1: 1}
                _ = v31
            case -627:
                var v32 bool = (true && (42 == 2147483647))
                _ = v32
            case 999:
                var v33 int32 = (-(v21))
                _ = v33
            }
            _ = fi28
        }
    }
    var v34 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v34
}
