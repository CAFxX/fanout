package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint64) uint32 {
    var v4 int32 = -2147483648
    if (65535 <= 4294967295) {
        for fi5 := 0; fi5 < 19; fi5++ {
            if true {
                for fi6 := 0; fi6 < 6; fi6++ {
                    v4 = (v4 << ((65535) & 63))
                    fmt.Println(fi5, v4, v4)
                    p0 -= (p0 << ((42) & 63))
                    fmt.Println(fi5, fi6)
                    _ = fi6
                }
                for fi7 := 0; fi7 < 5; fi7++ {
                    v8 := []int{fi7, (fi5 / ((32768) | 1)), fi7}
                    _ = fi7
                    _ = v8
                }
                var v9 uint16 = 57
                _ = v9
            }
            var v10 float32 = 3.14159
            _ = fi5
            _ = v10
        }
    }
    var v11 uint8 = 1
    switch "a" {
    case "\n\t":
        if false {
            p0 = 1
        }
    case "hello":
        v12 := []int{0, -801}
        _ = v12
    default:
        var v13 int = -2147483648
        _ = v13
    }
    v14 := GenS1{F0: 2, F1: -32768}
    _ = p0
    _ = v11
    _ = v14
    _ = v4
    return 4294967295
}

func main() {
    var v1 string = "a"
    var v2 bool = ((-2.25 <= (1e10 / 1e10)) || false)
    v3 := []uint32{}
    switch ((v1 + (((v1 + ("foo bar" + v1)) + "\n\t") + ("hello" + v1))) + "\n\t") {
    case "hello":
        var v15 uint64 = 377
        _ = v15
    case "":
        var v16 uint16 = 65535
        _ = v16
    }
    var v17 int16 = 32767
    var v18 int32 = 2147483647
    v18 += v18
    fmt.Println(v17)
    _ = len(v3)
    if (2 <= 42) {
        v19 := GenS1{F0: 2, F1: v17}
        if (0 != 42) {
            fmt.Println(v18, v2, v1)
            var v20 int32 = (v18 & (v18 - (^(v18))))
            if false {
                if (!v2) {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v20 = v20
                        _ = fi21
                    }
                    if (-1 == -675) {
                        fmt.Println(v2)
                        for fi22 := 0; fi22 < 2; fi22++ {
                            switch 265 {
                            case 32768:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 128:
                                v23 := GenS2{F0: 2147483648, F1: 4294967295}
                                _ = v23
                            }
                            _ = fi22
                        }
                        v24 := GenS2{F0: 2147483648, F1: 42}
                        _ = v24
                    } else {
                        var v25 int64 = -185
                        var v26 int16 = v17
                        if ((0.0 / 1e10) < 0.0) {
                            v2 = v2
                            for fi27 := 0; fi27 < 19; fi27++ {
                                v25 += 208
                                v26 = (v17 | int16(fi27))
                                _ = fi27
                            }
                            for fi28 := 0; fi28 < 18; fi28++ {
                                if v2 {
                                    var v29 string = (v1 + v1)
                                    fmt.Println(v17)
                                    var v30 int64 = 42
                                    _ = v29
                                    _ = v30
                                }
                                fmt.Println(v18, v26, v25)
                                _ = fi28
                            }
                        }
                        _ = v25
                        _ = v26
                    }
                    v18 += v20
                }
                if v2 {
                    var v31 uint32 = 4294967295
                    _ = v31
                } else {
                    if (v2 && (true && false)) {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            var v33 uint8 = 43
                            _ = fi32
                            _ = v33
                        }
                        switch 1 {
                        case 249:
                            v34 := GenS2{F0: 0, F1: 643}
                            _ = v34
                        case 1:
                            v20 += (v18 | 42)
                        case 32767:
                            if false {
                                var v35 uint64 = 0
                                v35 -= 375
                                v36 := GenS1{F0: 149, F1: v17}
                                _ = v35
                                _ = v36
                            }
                        default:
                            fmt.Println(v17)
                        }
                    }
                    v3 = append(v3, 0)
                }
                fmt.Println(v1)
            }
            _ = v20
        }
        _ = v19
    }
    if (!v2) {
        v37 := GenS1{F0: 302, F1: 784}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
