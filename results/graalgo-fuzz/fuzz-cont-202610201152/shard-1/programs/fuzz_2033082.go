package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int32 = -1
    var v2 uint = 2
    v3 := make(map[string]int)
    var v4 uint = v2
    var v5 float32 = (-2.25 / float32(3.14159))
    var v6 uint64 = 1
    var v7 float32 = v5
    _ = len(v3)
    if true {
        var v8 uint64 = 18446744073709551615
        _ = v8
    } else {
        v9 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 32767}
        _ = v9
    }
    delete(v3, ("foo bar" + "a"))
    switch "hello" {
    case "":
        if false {
            v10 := []int{128, 128, 2147483647, 4294967295}
            var v11 int = -9223372036854775808
            _ = v10
            _ = v11
        }
    default:
        if false {
            var v12 int64 = -9223372036854775808
            switch "\n\t" {
            case "":
                fmt.Println(v6)
            case "a":
                v4 -= (v4 / ((255) | 1))
            case "foo bar":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int8 = int8(fi13)
                    delete(v3, "")
                    _ = fi13
                    _ = v14
                }
            default:
                v15 := []int{}
                _ = v15
            }
            _ = v12
        }
    }
    v4 = v4
    var v16 int = -603
    for fi17 := 0; fi17 < 15; fi17++ {
        if true {
            switch (^(fi17)) {
            case 1:
                if (!(v6 >= v6)) {
                    fmt.Println(v6, v4)
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 uint16 = 42
                        var v20 int8 = int8(v19)
                        if ((^(v19)) <= (v19 + v19)) {
                            var v21 float32 = v7
                            _ = v21
                        } else {
                            _ = len(v3)
                            var v22 int16 = int16(v4)
                            _ = v22
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
            default:
                var v23 int = (v16 | fi17)
                _ = v23
            }
            var v24 uint = (^(v2))
            switch (("foo bar" + "\n\t") + ("hello" + "")) {
            case "hello":
                if (42 < -9223372036854775808) {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v26 := GenS2{F0: -960}
                        v6 += (v6 << ((65535) & 63))
                        _ = fi25
                        _ = v26
                    }
                }
            case "":
                v27 := GenS1{F0: 1, F1: v16, F2: 255}
                _ = v27
            }
            _ = v24
        }
        _ = fi17
    }
    switch "foo bar" {
    case "":
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint16 = 1
            if ((float64(((v5 / (((v7 / ((0.1 / v7) - (v5 * (1e10 * (3.14159 / ((0.1 * (0.1 / 1.5)) + 1e10)))))) - (float32((0.0 * 1.5)) * -2.25)) * v7)) * (1.5 * v5))) + 0.0) == 1.5) {
                var v30 uint64 = v6
                if false {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v3["foo bar"] = (fi28 >> (((v4 ^ 255)) & 63))
                        v32 := GenS1{F0: 2147483647, F1: fi28, F2: 255}
                        v33 := GenS2{F0: 42}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v6 -= v6
                }
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 float32 = 1.5
                    v5 = ((1e10 * ((1e10 + v7) - v7)) + v5)
                    _ = fi34
                    _ = v35
                }
                _ = v30
            }
            if (1 == int8(v1)) {
                _ = len(v3)
                var v36 uint8 = 105
                _ = v36
            }
            _ = fi28
            _ = v29
        }
    }
    if true {
        fmt.Println(v6)
    }
    var v37 uint16 = 567
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := v3[""]
        var v40 float32 = float32(1.5)
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := []int{v16, v16, v16, (v16 * (-(v16)))}
    fmt.Println(v37, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
