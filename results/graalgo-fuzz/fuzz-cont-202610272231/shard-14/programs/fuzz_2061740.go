package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
func genFn0() uint32 {
    if (!false) {
        var v6 int32 = -2147483648
        _ = v6
    } else {
        var v7 int = 4294967295
        v7 += -9223372036854775808
        if (!true) {
            var v8 uint16 = 32767
            var v9 uint = 65535
            _ = v8
            _ = v9
        } else {
            var v10 float64 = 3.14159
            v11 := GenS1{F0: -550, F1: "hello", F2: 42}
            fmt.Println(v10, v10)
            _ = v10
            _ = v11
        }
        _ = v7
    }
    if (true || (!true)) {
        var v12 uint16 = 255
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    return 973
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (0.1 / 1e10)
    var v3 int64 = 582
    var v4 float32 = 1e10
    v5 := make(map[uint32]bool)
    if (!(!true)) {
        var v15 uint64 = 259
        _ = v15
    }
    var v16 int = 128
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            switch ("foo bar" + "\n\t") {
            case "a":
                var v19 uint64 = 0
                _ = v19
            default:
                _ = len(v5)
            }
            _ = fi18
        }
        var v20 uint8 = 2
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch (^(fi21)) {
            case -585:
                v16 = fi22
            }
            fmt.Println(fi22, v3, v16)
            switch "hello" {
            case "a":
                var v23 int = (fi21 ^ v16)
                _ = v23
            case "\n\t":
                var v24 string = ("" + "")
                _ = v24
            default:
                v16 += (-(fi21))
            }
            _ = fi22
        }
        v2 += (v2 * float64(float32(0.0)))
        _ = fi21
    }
    var v25 uint8 = 5
    v2 = 3.14159
    for fi26 := 0; fi26 < 19; fi26++ {
        switch "hello" {
        case "\n\t":
            v27 := []int{v16, -235, v16, v16}
            _ = v27
        default:
            var v28 int32 = -699
            _ = v28
        }
        fmt.Println(v2, v3, fi26)
        _ = fi26
    }
    fmt.Println(v3, v3)
    v29 := []int{}
    v30 := GenS1{F0: -650, F1: (((("hello" + "\n\t") + "a") + ("hello" + "hello")) + ("foo bar" + "foo bar")), F2: 918}
    if (-32768 < 1000) {
        if (!(v16 == v16)) {
            var v31 bool = false
            _ = v31
        } else {
            v25 += 112
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        }
        v3 = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
