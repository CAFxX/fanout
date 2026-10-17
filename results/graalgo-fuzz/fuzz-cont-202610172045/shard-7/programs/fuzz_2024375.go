package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    var v3 bool = (-299 == 9223372036854775807)
    var v4 uint32 = 2147483647
    var v5 string = (("hello" + "\n\t") + "a")
    var v6 string = v5
    var v7 float64 = -2.25
    for fi8 := 0; fi8 < 19; fi8++ {
        if v3 {
            var v9 int32 = 0
            v10 := GenS1{F0: ("a" + v6)}
            _ = v10
            _ = v9
        } else {
            if v3 {
                fmt.Println(v1, v5)
            }
        }
        _ = fi8
    }
    var v11 float32 = (1e10 / ((((0.1 + 0.0) + 3.14159) - -2.25) + 0.1))
    for fi12 := 0; fi12 < 5; fi12++ {
        v7 += v7
        _ = fi12
    }
    if (255 < 255) {
        v13 := GenS1{F0: (((v5 + v5) + ((v5 + v6) + "\n\t")) + (v6 + v5))}
        v14 := GenS1{F0: ""}
        if (false || v3) {
            v15 := []int{}
            _ = v15
        }
        _ = v13
        _ = v14
    }
    _ = len(v2)
    v16 := []int{128, 736, 0, -1}
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 float32 = (v11 + 3.14159)
        if v3 {
            var v19 int16 = -32768
            _ = v19
        } else {
            v5 = v5
            if (!(v3 && (v18 != v18))) {
                if false {
                    v2[(v4 << ((432) & 63))] = -381
                    for fi20 := 0; fi20 < 9; fi20++ {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v11 = (v11 - v11)
                        fmt.Println(v5, fi17, v7)
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if v3 {
                            v5 = ""
                            if (670 <= -1) {
                                var v22 uint16 = 255
                                _ = v22
                            }
                            var v23 uint16 = 32767
                            _ = v23
                        } else {
                            var v24 uint32 = (v4 / (((^(v4))) | 1))
                            v24 = (v4 | 2147483648)
                            _ = v24
                        }
                        _ = fi21
                    }
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        var v26 uint64 = 2
                        _ = fi25
                        _ = v26
                    }
                    if (!(!((fi17 / ((9223372036854775807) | 1)) > fi17))) {
                        var v27 int = fi17
                        v28 := GenS1{F0: (v5 + v6)}
                        _ = v27
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 uint16 = 595
                    var v31 uint64 = 509
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        }
        _ = fi17
        _ = v18
    }
    var v32 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
