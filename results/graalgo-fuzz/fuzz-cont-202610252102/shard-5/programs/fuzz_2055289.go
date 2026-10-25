package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 uint64; F1 int64; F2 int8 }
func genFn0(p0 int64) uint32 {
    for fi5 := 0; fi5 < 1; fi5++ {
        if true {
            v6 := GenS1{F0: (1 >= 393)}
            p0 = (p0 % ((p0) | 1))
            if false {
                for fi7 := 0; fi7 < 4; fi7++ {
                    var v8 int64 = (p0 + p0)
                    v8 -= v8
                    switch "a" {
                    case "hello":
                        switch (fi7 | 2147483647) {
                        case 0:
                            v9 := []int{1}
                            _ = v9
                        default:
                            v10 := []int{}
                            _ = v10
                        }
                    case "foo bar":
                        v11 := GenS1{F0: false}
                        _ = v11
                    case "":
                        fmt.Println(v8)
                    }
                    _ = fi7
                    _ = v8
                }
                p0 -= p0
                p0 += p0
            } else {
                p0 = 1
                var v12 uint32 = 65535
                _ = v12
            }
            _ = v6
        }
        fmt.Println(fi5, p0)
        _ = fi5
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float64 = 3.14159
        v14 = (0.1 + 0.0)
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 string = (("hello" + "a") + ("foo bar" + ""))
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    return 65
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 59
    v3 := make(map[int]uint32)
    var v4 int8 = -51
    v4 = 1
    v2 -= (v2 << ((46) & 63))
    v17 := []int{391, 128, 128, 32767}
    switch "a" {
    case "\n\t":
        if (!true) {
            v2 -= v2
        } else {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            fmt.Println(v2, v2)
            v4 = (v4 * (v4 & (v4 % ((v4) | 1))))
        }
    }
    var v18 uint16 = 42
    v4 = v4
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int16 = -32768
        if false {
            var v21 uint64 = 1
            if (!false) {
                _ = len(v3)
                v18 = 42
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 uint = 4294967295
                    var v24 float32 = -2.25
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            v18 += v18
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v17 = append(v17, 0)
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(fi26, v2)
            _ = fi26
        }
        v3[(fi25 >> ((65535) & 63))] = (^(v2))
        var v27 uint8 = 78
        _ = fi25
        _ = v27
    }
    var v28 int16 = 753
    var v29 uint32 = 406
    var v30 uint8 = 127
    if (-1 < -600) {
        var v31 int = 998
        if (v2 != v2) {
            v32 := GenS2{F0: 586}
            if (!true) {
                var v33 int = (v31 << ((0) & 63))
                var v34 int16 = v28
                if true {
                    switch (("hello" + "a") + (("" + "") + "a")) {
                    case "foo bar":
                        v35 := GenS1{F0: (float32(0.0) < 0.0)}
                        _ = v35
                    case "a":
                        v31 = v33
                    default:
                        v36 := GenS3{F0: 1, F1: 1, F2: uint(v30)}
                        _ = v36
                    }
                }
                _ = v33
                _ = v34
            }
            _ = v32
        }
        _ = v31
    }
    var v37 bool = (!true)
    var v38 uint = 4294967295
    v39 := GenS4{F0: 0, F1: 771, F2: v4}
    var v40 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
