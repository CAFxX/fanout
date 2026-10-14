package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int32 }

func main() {
    v1 := []int{32767, -562, 127}
    v2 := make(map[string]uint32)
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    var v5 float32 = -2.25
    var v6 int64 = 9223372036854775807
    _ = len(v1)
    switch "a" {
    case "\n\t":
        v6 -= v6
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    _ = len(v1)
    var v7 bool = (!(128 >= 9223372036854775807))
    switch 4294967295 {
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = 128
        }
    case 32768:
        v8 := v4["foo bar"]
        _ = v8
    case 32767:
        fmt.Println(v5)
    default:
        switch 127 {
        case 128:
            var v9 int32 = -1
            _ = v9
        default:
            v10 := GenS1{F0: -9223372036854775808, F1: -1, F2: 2147483648}
            _ = v10
        }
    }
    var v11 int64 = (v6 | v6)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 string = ("a" + "a")
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    var v15 float64 = ((0.1 + 1e10) + (3.14159 * 0.1))
    v16 := GenS1{F0: 279, F1: 32767, F2: 0}
    v11 = 0
    var v17 uint32 = 0
    var v18 int8 = 127
    if true {
        if (65535 != 247) {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int16 = -440
                if false {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        switch "hello" {
                        case "":
                            v22 := GenS2{F0: 32767}
                            _ = v22
                        default:
                            v23 := GenS3{F0: 613}
                            _ = v23
                        }
                        v4["hello"] = (v17 << ((4294967295) & 63))
                        v24 := GenS1{F0: (fi21 << ((18446744073709551615) & 63)), F1: fi21, F2: v17}
                        _ = fi21
                        _ = v24
                    }
                    fmt.Println(v20, v3, v14)
                } else {
                    var v25 float64 = v15
                    var v26 uint16 = 607
                    _ = v25
                    _ = v26
                }
                v27 := v4[""]
                _ = fi19
                _ = v20
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v5 += float32(v15)
            if false {
                if false {
                    v29 := GenS4{F0: "hello"}
                    _ = v29
                }
                if true {
                    v30 := GenS5{F0: v17}
                    v31 := GenS4{F0: "hello"}
                    _ = v30
                    _ = v31
                } else {
                    v7 = v7
                    v14 = (v14 - 2)
                }
            } else {
                v2["a"] = 4294967295
                var v32 int32 = -2147483648
                _ = v32
            }
            _ = fi28
        }
    } else {
        var v33 int16 = (v3 >> ((458) & 63))
        _ = v33
    }
    v34 := GenS6{F0: (false || false), F1: 111}
    if true {
        if (v18 != v18) {
            if v7 {
                v15 += v15
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
