package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int, p1 int64, p2 string) uint16 {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: 255}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 uint64) int16 {
    v9 := []int{170}
    if len(v9) > 0 {
        v9[0] = 32768
    }
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS2{F0: -550}
                var v13 uint = 42
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                p0 += p0
                var v15 int8 = 78
                _ = fi14
                _ = v15
            }
            p0 = 380
            _ = fi10
        }
        var v16 string = "hello"
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int = (-(fi17))
            for fi19 := 0; fi19 < 13; fi19++ {
                v18 = fi19
                fmt.Println(v16, v16, v18)
                _ = fi19
            }
            var v20 int16 = 163
            _ = fi17
            _ = v18
            _ = v20
        }
        _ = v16
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 string = "hello"
    var v5 uint16 = 946
    var v6 uint64 = 42
    var v21 int8 = 104
    fmt.Println(v21)
    v1 += v1
    switch (v1 ^ (v1 ^ (v1 - (^(v1))))) {
    case 32768:
        var v22 uint32 = 4294967295
        _ = v22
    case -629:
        v23 := GenS3{F0: 1000, F1: v1}
        _ = v23
    }
    fmt.Println(v21, v6)
    switch int(v5) {
    case -1:
        v21 = (v21 * 42)
    }
    v5 = (v5 - 242)
    switch (v4 + "") {
    case "hello":
        v24 := GenS3{F0: 1000, F1: (v1 / (((v1 & v1)) | 1))}
        _ = v24
    case "":
        if (!(v21 != v21)) {
            var v25 uint64 = v6
            _ = v25
        } else {
            v21 += (-(v21))
        }
    case "foo bar":
        var v26 int = v1
        _ = v26
    default:
        var v27 uint64 = (v6 - (v6 - v6))
        _ = v27
    }
    v28 := GenS4{F0: v6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
