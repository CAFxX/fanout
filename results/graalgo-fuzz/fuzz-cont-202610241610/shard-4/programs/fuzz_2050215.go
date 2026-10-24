package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }
func genFn0() int {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 string = "hello"
        if ((0 == 803) || (255 <= 202)) {
            var v9 int8 = 1
            v10 := []int{(fi7 % ((fi7) | 1))}
            _ = v10
            _ = v9
        } else {
            if true {
                switch (fi7 - fi7) {
                case 9223372036854775807:
                    var v11 uint64 = 857
                    _ = v11
                case 127:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 uint = 18446744073709551615
                        if true {
                            if true {
                                v8 = ("\n\t" + ((v8 + ((v8 + "") + v8)) + v8))
                                v8 = ((v8 + v8) + "a")
                            } else {
                                v14 := GenS1{F0: -905}
                                fmt.Println(v13, fi12, v13)
                                _ = v14
                            }
                            v15 := GenS2{F0: 0, F1: 4294967295, F2: 877}
                            v8 = "hello"
                            _ = v15
                        }
                        fmt.Println(v8, v13)
                        _ = fi12
                        _ = v13
                    }
                case 639:
                    v8 = (v8 + "a")
                }
                v8 = ("hello" + (v8 + v8))
            } else {
                var v16 uint = 1
                if false {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v18 := GenS3{F0: fi17}
                        _ = fi17
                        _ = v18
                    }
                    v16 = 0
                    v8 = (v8 + "\n\t")
                }
                v16 = v16
                _ = v16
            }
        }
        _ = fi7
        _ = v8
    }
    var v19 string = "a"
    switch "\n\t" {
    case "a":
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := []int{(-(fi20)), (fi20 ^ fi20), 780, 32767}
            v22 := GenS4{F0: false, F1: v19, F2: 255}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 uint32 = 1
    var v24 uint32 = v23
    _ = v19
    _ = v23
    _ = v24
    return 128
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    var v5 int32 = 0
    var v6 float32 = v1
    fmt.Println(v1, v5, v6)
    v1 = (0.1 * v1)
    if ((!true) && (((384 <= 4294967295) && false) || false)) {
        fmt.Println(v6, v1, v6)
    }
    var v25 bool = (false && ((255 <= 255) || (-79 >= 0)))
    v26 := []int{566, 4294967295}
    var v27 string = "a"
    if (true && (127 != -12)) {
        var v28 uint64 = 117
        _ = v28
    } else {
        if (v25 && (v6 >= -2.25)) {
            v1 -= 0.1
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
        v5 = (v5 >> ((329) & 63))
        var v29 uint16 = 65535
        _ = v29
    }
    var v30 uint64 = 430
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
