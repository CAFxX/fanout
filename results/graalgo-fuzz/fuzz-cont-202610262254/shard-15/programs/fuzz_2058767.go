package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
func genFn0() int {
    if false {
        if true {
            var v7 string = ("\n\t" + "foo bar")
            _ = v7
        } else {
            var v8 float32 = 1.5
            _ = v8
        }
    } else {
        var v9 uint8 = 2
        if true {
            for fi10 := 0; fi10 < 5; fi10++ {
                for fi11 := 0; fi11 < 9; fi11++ {
                    if true {
                        v9 = 2
                        var v12 uint16 = 807
                        if (17 > 33) {
                            v13 := []int{(fi11 + fi11)}
                            switch fi11 {
                            case 965:
                                var v14 float64 = 0.1
                                _ = v14
                            case -1:
                                var v15 float32 = 3.14159
                                _ = v15
                            case -438:
                                fmt.Println(v12, fi10)
                            }
                            var v16 uint8 = 2
                            _ = v13
                            _ = v16
                        }
                        _ = v12
                    }
                    _ = fi11
                }
                v17 := []int{}
                for fi18 := 0; fi18 < 18; fi18++ {
                    if len(v17) > 0 {
                        v17[0] = 1
                    }
                    v9 -= (v9 - v9)
                    var v19 string = ""
                    var v20 uint16 = 248
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi10
                _ = v17
            }
        } else {
            if true {
                fmt.Println(v9)
                var v21 uint8 = 221
                _ = v21
            } else {
                fmt.Println(v9)
                var v22 int8 = 42
                _ = v22
            }
        }
        _ = v9
    }
    if (1 > 0) {
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS1{F0: -2147483648, F1: -2147483648, F2: 1}
            fmt.Println(fi23)
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS1{F0: -128, F1: 2147483647, F2: -987}
    _ = v25
    return 460
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    var v3 uint = 0
    var v4 uint8 = 1
    var v5 int64 = -239
    var v6 int = 130
    v4 = v4
    fmt.Println(v1, v3)
    v26 := []int{v6, int(v3)}
    v3 += v3
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            v6 = genFn0()
            _ = fi28
        }
        fmt.Println(fi27, v5)
        _ = fi27
    }
    fmt.Println(v3)
    var v29 string = "hello"
    v26 = append(v26, (v6 - genFn0()))
    v6 = (v6 << ((v3) & 63))
    v5 = (v5 << ((649) & 63))
    if (9223372036854775807 > (v5 + (v5 - -1))) {
        var v30 string = v29
        _ = len(v26)
        _ = v30
    } else {
        var v31 string = "\n\t"
        var v32 bool = (2 > 703)
        switch -2147483648 {
        case -558:
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 uint16 = 702
                _ = fi33
                _ = v34
            }
        default:
            fmt.Println(v29)
        }
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
