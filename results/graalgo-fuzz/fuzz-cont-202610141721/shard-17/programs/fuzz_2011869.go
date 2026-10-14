package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 int16) string {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 string = "a"
        p0 = 221
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 74}
    _ = p0
    _ = p1
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint16, p1 bool, p2 uint8) int16 {
    if p1 {
        p2 -= p2
        p0 += p0
        fmt.Println(p0, p2, p2)
    }
    var v9 uint8 = (p2 * uint8(p0))
    var v10 uint64 = uint64(p0)
    fmt.Println(p2, v10, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 0
    var v2 uint32 = (^(v1))
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 int64 = 42
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int8 = 42
        _ = fi11
        _ = v12
    }
    v5 = (v5 / ((v5) | 1))
    var v13 float32 = 3.14159
    for fi14 := 0; fi14 < 14; fi14++ {
        v13 += v13
        _ = fi14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int = 2147483647
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v17 := GenS1{F0: -9223372036854775808}
            _ = v17
        case "":
            v18 := GenS1{F0: (v5 << ((2) & 63))}
            _ = v18
        case "hello":
            var v19 int = -246
            _ = v19
        default:
            var v20 string = "foo bar"
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    v21 := []int{-575, 128, -2147483648}
    if len(v21) > 0 {
        v21[0] = 32768
    }
    switch 4294967295 {
    case -573:
        var v22 int8 = 127
        _ = v22
    case 32767:
        if (!(v2 >= (v1 - (^(v1))))) {
            var v23 uint64 = 18446744073709551615
            var v24 int = 32767
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 float32 = float32(1.5)
                _ = fi25
                _ = v26
            }
            _ = v23
            _ = v24
        } else {
            var v27 uint8 = 255
            fmt.Println(v1, v2)
            v28 := GenS1{F0: 9223372036854775807}
            _ = v27
            _ = v28
        }
    case -9223372036854775808:
        switch "a" {
        case "hello":
            var v29 int8 = -1
            _ = v29
        default:
            if true {
                switch "foo bar" {
                case "foo bar":
                    v30 := GenS1{F0: v5}
                    _ = v30
                case "":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                v2 += (v2 << ((221) & 63))
                for fi31 := 0; fi31 < 5; fi31++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    var v32 int32 = 987
                    _ = fi31
                    _ = v32
                }
            }
        }
    }
    v1 = 18
    _ = len(v4)
    switch (("" + "") + ("foo bar" + ("" + "foo bar"))) {
    case "foo bar":
        var v33 int8 = -119
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        v2 -= (v2 << ((909) & 63))
        v5 = v5
        _ = fi34
    }
    var v35 uint32 = v2
    v36 := GenS1{F0: v5}
    if len(v21) > 0 {
        v21[0] = int(v1)
    }
    v1 -= 4294967295
    v37 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
