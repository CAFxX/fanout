package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 string, p1 int, p2 uint64) int32 {
    var v4 string = p0
    v4 = ""
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 string = ""
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return int32(p2)
}
func genFn1(p0 uint64) uint8 {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch ("hello" + "hello") {
        case "":
            var v8 uint16 = 615
            _ = v8
        case "\n\t":
            if (1 >= 9223372036854775807) {
                if (false || (548 <= 1)) {
                    var v9 int32 = -809
                    _ = v9
                } else {
                    fmt.Println(fi7)
                    p0 += p0
                }
                v10 := []int{127, 0}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 11; fi11++ {
                    if (!(283 == -2147483648)) {
                        v12 := []int{fi7}
                        _ = len(v12)
                        _ = v12
                    }
                    _ = fi11
                }
                var v13 string = ("hello" + "a")
                fmt.Println(v13)
                _ = v13
            }
        }
        var v14 uint = 4294967295
        _ = fi7
        _ = v14
    }
    v15 := []int{127, 32767, 720}
    if false {
        var v16 uint16 = 775
        fmt.Println(v16, p0)
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 float64 = 1.5
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v19 int64 = 9223372036854775807
    v20 := GenS1{F0: 4294967295, F1: 2}
    _ = p0
    _ = v15
    _ = v19
    _ = v20
    return 127
}

func main() {
    v1 := []int{-699}
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    v3 += genFn0("foo bar", 663, 266)
    v3 += (v3 % ((v3) | 1))
    switch -36 {
    case 128:
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 int = (fi21 * fi21)
            _ = fi21
            _ = v22
        }
    }
    var v23 int32 = (v3 | genFn0("hello", 0, 537))
    var v24 uint64 = 1
    _ = len(v1)
    var v25 uint8 = 1
    fmt.Println(v3, v23, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2[4294967295] = 4294967295
    v24 = v24
    var v26 float32 = float32(0.0)
    switch "\n\t" {
    case "\n\t":
        if true {
            switch 128 {
            case 128:
                for fi27 := 0; fi27 < 14; fi27++ {
                    if ((-54 <= 100) && true) {
                        v28 := GenS2{F0: 32767, F1: v25}
                        _ = v28
                    } else {
                        var v29 int16 = -1
                        _ = len(v1)
                        for fi30 := 0; fi30 < 9; fi30++ {
                            if true {
                                v29 += v29
                                fmt.Println(v24, v3)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = fi30
                        }
                        _ = v29
                    }
                    v31 := GenS2{F0: -615, F1: v25}
                    _ = fi27
                    _ = v31
                }
            default:
                v32 := GenS2{F0: -32768, F1: (v25 % ((128) | 1))}
                _ = v32
            }
            var v33 float64 = 0.0
            _ = v33
        }
    }
    v34 := GenS2{F0: -32768, F1: (^(v25))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
}
