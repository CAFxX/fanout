package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := []int{}
    v4 := []uint32{874, 65535}
    var v5 uint = 438
    var v6 uint8 = 127
    var v7 uint32 = 1
    v7 = (v7 << (((v5 / (((v5 / ((706) | 1))) | 1))) & 63))
    var v8 float64 = 0.1
    if true {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch "hello" {
            case "hello":
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(v8, v8)
                    if false {
                        var v11 float64 = v8
                        _ = v11
                    } else {
                        var v12 int8 = -1
                        var v13 float32 = (1.5 / -2.25)
                        if true {
                            if len(v3) > 0 {
                                v3[0] = (-(fi10))
                            }
                            if false {
                                var v14 uint64 = 1
                                _ = v14
                            } else {
                                v12 -= (v12 + v12)
                                fmt.Println(v5, v8, v12)
                                var v15 uint8 = v6
                                _ = v15
                            }
                        }
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                }
            default:
                v16 := GenS1{F0: v6}
                _ = v16
            }
            switch ("" + "") {
            case "":
                var v17 int16 = 32767
                _ = v17
            case "foo bar":
                switch fi9 {
                case -1:
                    var v18 uint16 = 690
                    _ = v18
                }
            case "hello":
                var v19 int = (fi9 / ((1) | 1))
                _ = v19
            }
            var v20 string = ("hello" + ("a" + ""))
            _ = fi9
            _ = v20
        }
        var v21 string = (("\n\t" + "hello") + ((("" + "\n\t") + "foo bar") + ("foo bar" + (("hello" + "a") + ""))))
        _ = v21
    }
    v5 -= (v5 % ((0) | 1))
    var v22 uint64 = 1
    v7 = (v7 | v7)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -1 {
    case -91:
        var v23 bool = false
        _ = v23
    case 2147483647:
        v6 += v6
    case 889:
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint = (v5 - v5)
            if (true && ("hello" > "foo bar")) {
                for fi26 := 0; fi26 < 19; fi26++ {
                    fmt.Println(v25, v25)
                    fmt.Println(v7, v7, fi24)
                    _ = fi26
                }
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(fi27, fi24, v5)
                    _ = fi27
                }
                v28 := GenS1{F0: v6}
                _ = v28
            }
            fmt.Println(v25, v22)
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
