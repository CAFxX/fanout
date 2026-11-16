package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8) int {
    switch 127 {
    case 9223372036854775807:
        v4 := []int{9223372036854775807, 32768, 32768}
        _ = v4
    }
    v5 := []int{965, -348, 32768}
    v6 := GenS1{F0: 42}
    var v7 uint16 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{42, 725, 2147483647}
    var v2 int16 = 1000
    var v3 string = "\n\t"
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int = fi8
        switch (("foo bar" + "a") + v3) {
        case "foo bar":
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(v9)
                switch (fi10 & v9) {
                case 32767:
                    fmt.Println(fi8)
                case -1:
                    v11 := GenS1{F0: v2}
                    _ = v11
                case -578:
                    v3 = (v3 + (v3 + ((v3 + v3) + (v3 + ("\n\t" + v3)))))
                default:
                    if true {
                        fmt.Println(fi8)
                        fmt.Println(v3, v2, v2)
                    } else {
                        if false {
                            switch "a" {
                            case "foo bar":
                                switch v3 {
                                case "hello":
                                    var v12 float32 = (((-2.25 * 0.0) / (1.5 + 0.0)) * 1e10)
                                    _ = v12
                                case "":
                                    var v13 int64 = 42
                                    if true {
                                        v3 = "foo bar"
                                        v13 += (-(v13))
                                        switch "" {
                                        case "a":
                                            v2 = (v2 * v2)
                                        }
                                        v14 := GenS1{F0: v2}
                                        _ = v14
                                    } else {
                                        v15 := GenS1{F0: v2}
                                        _ = v15
                                    }
                                    _ = v13
                                case "a":
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                    } else {
                                        v9 += genFn0(26)
                                    }
                                }
                            default:
                                v16 := GenS2{F0: 1, F1: true}
                                _ = v16
                            }
                            v3 = v3
                        }
                        v9 += (fi8 & v9)
                    }
                }
                _ = fi10
            }
        }
        v3 = (((v3 + "a") + v3) + "foo bar")
        _ = fi8
        _ = v9
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS1{F0: v2}
        fmt.Println(v2, v3, v2)
        v3 = v3
        _ = fi17
        _ = v18
    }
    fmt.Println(v3)
    v2 -= v2
    fmt.Println(v2, v2)
    var v19 float64 = 0.1
    if (9223372036854775807 != -1) {
        _ = len(v1)
    }
    var v20 uint8 = 1
    var v21 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
