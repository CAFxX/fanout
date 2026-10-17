package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint16) int64 {
    switch ("" + "hello") {
    case "a":
        v7 := []int{-2147483648}
        _ = v7
    case "foo bar":
        if ("hello" >= ("\n\t" + "")) {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 float32 = 0.0
                var v10 bool = true
                var v11 float64 = 1.5
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 float32 = float32(1e10)
                _ = fi12
                _ = v13
            }
        } else {
            fmt.Println(p0)
        }
    case "hello":
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: (^(p0)), F1: -122, F2: 176}
            _ = fi14
            _ = v15
        }
    }
    if true {
        v16 := GenS1{F0: (p0 / ((p0) | 1)), F1: 32768, F2: 82}
        v17 := []int{2147483647, -9223372036854775808}
        var v18 uint = 2
        _ = v16
        _ = v17
        _ = v18
    } else {
        var v19 uint16 = (^(p0))
        _ = v19
    }
    _ = p0
    return -784
}

func main() {
    var v1 bool = false
    var v2 uint8 = 128
    v3 := make(map[uint32]string)
    var v4 bool = v1
    var v5 float32 = (float32((0.0 * (((float64(((((((((float32(3.14159) / 1.5) - ((0.0 - (3.14159 + 0.0)) - ((1e10 - (0.0 - float32(1.5))) * (-2.25 / float32(0.1))))) * -2.25) - 1.5) / float32((0.1 * 1.5))) / -2.25) / 1.5) - -2.25)) * 0.0) - -2.25) * 1.5))) / 3.14159)
    v6 := make(map[string]uint32)
    var v20 uint = 1
    var v21 uint64 = uint64(v20)
    fmt.Println(v2, v21)
    var v22 uint16 = 20
    v23 := GenS2{F0: v20}
    if v1 {
        var v24 uint32 = 0
        _ = v24
    } else {
        fmt.Println(v5)
        v20 -= v20
        v5 = v5
    }
    if (float64(0.0) <= 1.5) {
        switch -574 {
        case 128:
            v21 = (v21 ^ v21)
        case -2147483648:
            var v25 string = "hello"
            _ = v25
        case 894:
            v20 += (v20 * v20)
        }
        v20 -= 2
        var v26 int = 801
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        if v1 {
            fmt.Println(v2, v5)
        } else {
            var v28 int16 = 42
            _ = v28
        }
        switch (fi27 / ((fi27) | 1)) {
        case 4294967295:
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := GenS2{F0: (v20 & 968)}
                var v31 float64 = 0.1
                v20 -= (v20 - v20)
                _ = fi29
                _ = v30
                _ = v31
            }
        case 2147483647:
            if (-2.25 <= 0.0) {
                v2 = (v2 >> (((v20 | 18446744073709551615)) & 63))
            }
        case 127:
            v32 := GenS2{F0: v20}
            _ = v32
        }
        _ = fi27
    }
    var v33 uint64 = (v21 + (v21 - v21))
    switch 32767 {
    case 4294967295:
        v22 = v22
    case 42:
        fmt.Println(v21, v21, v20)
    case 9223372036854775807:
        var v34 int8 = 97
        _ = v34
    default:
        v35 := GenS3{F0: v20, F1: -32768}
        _ = v35
    }
    v36 := []int{int(v33), 1}
    v37 := GenS1{F0: (v22 | (^(v22))), F1: -1, F2: (v2 | 127)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
