package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 0
    v3 := make(map[string]string)
    v4 := []int{-1, 128, -1}
    fmt.Println(v2)
    fmt.Println(v2)
    if (false || false) {
        var v5 uint8 = 0
        v2 += v2
        _ = v5
    } else {
        var v6 float64 = 1.5
        _ = v6
    }
    v2 += (v2 - v2)
    switch -746 {
    case 9223372036854775807:
        switch 2147483647 {
        case 127:
            var v7 uint8 = 168
            _ = v7
        case 2147483647:
            var v8 uint16 = 889
            _ = v8
        case 9223372036854775807:
            v9 := GenS1{F0: ("\n\t" + ("\n\t" + "\n\t"))}
            _ = v9
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float64 = 0.0
                for fi12 := 0; fi12 < 9; fi12++ {
                    v2 -= 1000
                    v13 := GenS2{F0: 413, F1: -128, F2: false}
                    if (!true) {
                        var v14 int = (fi10 * 42)
                        if (!true) {
                            v15 := GenS1{F0: "\n\t"}
                            v16 := GenS3{F0: 42}
                            _ = v15
                            _ = v16
                        }
                        if len(v4) > 0 {
                            v4[0] = 172
                        }
                        v2 += 0
                        var v17 string = "a"
                        _ = v14
                        _ = v17
                    }
                    if (!(!true)) {
                        if true {
                            if true {
                                if (true || false) {
                                    if len(v4) > 0 {
                                        v4[0] = 0
                                    }
                                    fmt.Println(v11)
                                    fmt.Println(v11, fi12)
                                } else {
                                    var v18 uint = 2
                                    var v19 int8 = 0
                                    _ = len(v1)
                                    _ = v18
                                    _ = v19
                                }
                                v20 := GenS3{F0: 108}
                                v2 = v2
                                _ = v20
                            }
                            v2 = (v2 << ((711) & 63))
                            var v21 float64 = float64((3.14159 * (0.1 - float32(v11))))
                            _ = v21
                        } else {
                            var v22 int64 = 42
                            v23 := GenS4{F0: 1}
                            _ = v22
                            _ = v23
                        }
                    }
                    if true {
                        v24 := GenS4{F0: (fi10 ^ (fi10 ^ fi10))}
                        v2 += (v2 + -641)
                        var v25 float32 = 1.5
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = fi10
                                }
                                switch fi12 {
                                case 4294967295:
                                    var v26 bool = false
                                    _ = v26
                                case -1:
                                    v27 := v1[2147483647]
                                    _ = v27
                                }
                            }
                        }
                        _ = v24
                        _ = v25
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        }
    }
    v3[""] = ("\n\t" + "hello")
    var v28 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
