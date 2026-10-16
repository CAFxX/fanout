package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
func genFn0() bool {
    var v4 float64 = -2.25
    v5 := GenS1{F0: 682, F1: 0, F2: 349}
    _ = v4
    _ = v5
    return (2147483647 <= 2147483647)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.0
    v3 := make(map[string]int)
    v2 = 0.1
    v6 := GenS1{F0: 42, F1: 42, F2: 352}
    var v7 uint32 = 889
    if (true || false) {
        var v8 int64 = 42
        if false {
            fmt.Println(v7, v7)
            v2 -= ((v2 / float32(float64(1.5))) + float32(0.1))
            if false {
                delete(v3, "foo bar")
                var v9 uint16 = 65535
                fmt.Println(v2)
                _ = v9
            } else {
                var v10 uint32 = 4294967295
                if ((false || (32767 > 65535)) && false) {
                    var v11 int64 = v8
                    if true {
                        if false {
                            var v12 bool = false
                            _ = v12
                        } else {
                            fmt.Println(v11, v2)
                            switch "hello" {
                            case "\n\t":
                                _ = len(v1)
                            }
                        }
                        for fi13 := 0; fi13 < 13; fi13++ {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                fmt.Println(v10, v8)
                                var v15 float64 = (float64(v2) / (float64((3.14159 - v2)) - -2.25))
                                _ = fi14
                                _ = v15
                            }
                            fmt.Println(v10, v2)
                            v16 := GenS2{F0: 1, F1: 32767, F2: 255}
                            _ = fi13
                            _ = v16
                        }
                        var v17 int64 = v11
                        _ = v17
                    }
                    _ = v11
                }
                _ = v10
            }
        }
        _ = v8
    }
    var v18 int = 2147483647
    if true {
        v2 = v2
    } else {
        var v19 uint8 = 255
        if false {
            var v20 float32 = 1e10
            var v21 uint8 = uint8(v7)
            v18 = 32768
            _ = v20
            _ = v21
        } else {
            if false {
                var v22 int8 = 0
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := []int{(fi24 & (v18 + int(v7))), fi23, (fi24 >> ((114) & 63))}
                    v26 := GenS2{F0: 339, F1: 0, F2: v19}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
            var v27 uint = 2
            _ = v27
        }
        v7 = uint32(v19)
        _ = v19
    }
    v2 = (1e10 - -2.25)
    var v28 int16 = 13
    v18 += -609
    fmt.Println(v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v6
    _ = v7
}
