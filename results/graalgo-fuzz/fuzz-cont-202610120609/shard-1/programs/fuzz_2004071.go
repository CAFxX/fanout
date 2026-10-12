package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 int32, p1 string) int64 {
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 9; fi5++ {
        switch v4 {
        case "\n\t":
            if false {
                var v6 uint8 = 48
                for fi7 := 0; fi7 < 16; fi7++ {
                    p1 = (("foo bar" + v4) + ("foo bar" + (p1 + (v4 + "a"))))
                    _ = fi7
                }
                fmt.Println(v4, p1)
                _ = v6
            } else {
                if (!(-838 <= 0)) {
                    v4 = "foo bar"
                    if false {
                        for fi8 := 0; fi8 < 19; fi8++ {
                            v9 := GenS1{F0: fi5}
                            _ = fi8
                            _ = v9
                        }
                        if true {
                            p1 = (v4 + (p1 + "hello"))
                            if true {
                                fmt.Println(v4, p1, p1)
                            }
                        } else {
                            var v10 int32 = p0
                            _ = v10
                        }
                        v11 := GenS2{F0: -23}
                        _ = v11
                    } else {
                        p0 = (p0 + p0)
                    }
                    if false {
                        for fi12 := 0; fi12 < 1; fi12++ {
                            var v13 bool = (42 < 42)
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        var v14 int = (fi5 / ((fi5) | 1))
                        for fi15 := 0; fi15 < 19; fi15++ {
                            if (!false) {
                                v4 = p1
                            }
                            v4 = (v4 + "foo bar")
                            _ = fi15
                        }
                        _ = v14
                    }
                }
                v4 = "foo bar"
            }
        default:
            v4 = (v4 + (v4 + p1))
        }
        _ = fi5
    }
    p0 = p0
    if false {
        var v16 uint64 = 2
        switch 9223372036854775807 {
        case -1:
            v17 := []int{556, 127, -2147483648, 32767}
            _ = v17
        default:
            fmt.Println(v4)
        }
        _ = v16
    } else {
        var v18 uint8 = 255
        var v19 uint8 = (^(v18))
        v20 := []int{}
        _ = v18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1() string {
    v21 := []int{}
    v22 := GenS1{F0: 127}
    var v23 int = 1
    if (3.14159 == 0.0) {
        var v24 uint = 18446744073709551615
        switch "foo bar" {
        case "a":
            var v25 uint16 = 42
            _ = v25
        case "\n\t":
            fmt.Println(v23, v24)
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        _ = v24
    }
    _ = v21
    _ = v22
    _ = v23
    return ("" + "hello")
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    v1 -= v1
    var v26 uint64 = 770
    v26 = v26
    var v27 uint32 = 2147483648
    var v28 float32 = 0.0
    var v29 uint8 = 127
    if false {
        v26 += (v26 >> ((603) & 63))
    } else {
        switch "hello" {
        case "\n\t":
            var v30 uint = 615
            _ = v30
        default:
            v31 := GenS1{F0: v1}
            _ = v31
        }
    }
    var v32 uint8 = (v29 ^ (v29 + 245))
    v32 = 42
    v33 := GenS3{F0: 0, F1: 0, F2: v27}
    var v34 int = (v1 >> ((861) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
