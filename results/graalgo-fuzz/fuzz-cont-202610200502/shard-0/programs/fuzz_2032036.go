package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 uint64, p2 int64) uint16 {
    fmt.Println(p1, p2, p0)
    switch "hello" {
    case "a":
        var v7 uint8 = 127
        _ = v7
    case "foo bar":
        p0 = "foo bar"
    case "":
        v8 := GenS1{F0: 723, F1: (-2147483648 >= 857), F2: 1000}
        _ = v8
    }
    if (!true) {
        p1 = (^(p1))
    }
    var v9 uint16 = 449
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (v9 >> ((0) & 63))
}
func genFn1(p0 int16, p1 int8, p2 uint8) bool {
    var v10 uint8 = p2
    p1 = (p1 + (p1 - -1))
    v11 := []int{32768}
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            if (!(true && ((1e10 + (3.14159 * 0.1)) <= (-2.25 * 3.14159)))) {
                var v14 uint = 4294967295
                v15 := GenS1{F0: 65535, F1: false, F2: 0}
                p0 -= (p0 | (p0 << (((^(v14))) & 63)))
                _ = v14
                _ = v15
            } else {
                if false {
                    v10 -= p2
                } else {
                    v10 -= p2
                    var v16 uint64 = 18446744073709551615
                    if true {
                        var v17 int64 = -1
                        p1 = 109
                        _ = v17
                    }
                    _ = v16
                }
                if (!false) {
                    _ = len(v11)
                    fmt.Println(p2, p2)
                    if (!(65535 >= 65535)) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        if len(v11) > 0 {
                            v11[0] = (fi12 << ((65535) & 63))
                        }
                    } else {
                        p0 = -521
                    }
                }
            }
            var v18 int16 = 141
            p1 = (p1 ^ (-(p1)))
            switch ("\n\t" + "") {
            case "":
                if true {
                    var v19 int = (fi13 * 128)
                    _ = v19
                }
            case "\n\t":
                v18 = int16(fi12)
            case "foo bar":
                v20 := GenS1{F0: 4294967295, F1: true, F2: v18}
                _ = v20
            }
            _ = fi13
            _ = v18
        }
        fmt.Println(p2)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (!true)
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    v6 := []uint32{619, 2}
    var v21 uint8 = (v1 * 153)
    var v22 string = ""
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v23 int8 = -1
    v24 := GenS2{F0: 809, F1: 42}
    switch -705 {
    case 791:
        var v25 int64 = 0
        _ = v25
    case 128:
        fmt.Println(v23, v22, v21)
    }
    v22 = (v22 + "a")
    v26 := GenS3{F0: 2, F1: "\n\t"}
    v23 = v23
    for fi27 := 0; fi27 < 11; fi27++ {
        v22 = ((v22 + (v22 + "hello")) + v22)
        if true {
            fmt.Println(v23, v1, v1)
            if (!(975 <= 509)) {
                v28 := GenS2{F0: 2, F1: 255}
                v29 := GenS4{F0: 4294967295}
                _ = v28
                _ = v29
            }
            var v30 float32 = float32(1e10)
            _ = v30
        } else {
            v31 := GenS1{F0: 4294967295, F1: (42 != 797), F2: 0}
            _ = v31
        }
        _ = fi27
    }
    fmt.Println(v21, v22)
    v32 := GenS5{F0: 44}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
