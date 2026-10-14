package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]bool)
    var v3 uint = (v1 & (v1 - v1))
    v4 := []int{}
    var v5 uint = (v3 >> ((255) & 63))
    var v6 string = "\n\t"
    var v7 uint64 = 889
    if false {
        var v8 int32 = 0
        _ = len(v4)
        _ = v8
    }
    var v9 int = 127
    v7 = (v7 - v7)
    var v10 int = 4294967295
    delete(v2, 4294967295)
    var v11 uint16 = 255
    var v12 uint16 = 852
    var v13 uint16 = (^(v11))
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v1 -= v3
            _ = fi15
        }
        fmt.Println(v6)
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            if true {
                if true {
                    var v16 float32 = 1e10
                    _ = v16
                }
                v10 += int(v1)
                if (-64 > -119) {
                    var v17 string = (v6 + "hello")
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v13 = (v12 / ((v12) | 1))
                        _ = len(v4)
                        if true {
                            fmt.Println(v10, fi18, v5)
                            var v19 uint = 2
                            var v20 int16 = -149
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float32 = 0.0
                            var v22 bool = true
                            var v23 int = v10
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        var v24 uint16 = v13
                        v25 := v2[42]
                        _ = fi18
                        _ = v24
                        _ = v25
                    }
                    _ = v17
                }
            } else {
                v26 := GenS1{F0: v6}
                v9 -= -976
                switch fi14 {
                case 0:
                    _ = len(v4)
                case 128:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    var v27 uint64 = v7
                    _ = v27
                }
                _ = v26
            }
        case "foo bar":
            v6 = ((v6 + "a") + (v6 + ("hello" + (v6 + (((v6 + (v6 + "hello")) + (v6 + "\n\t")) + v6)))))
        default:
            var v28 int64 = -685
            _ = v28
        }
        _ = fi14
    }
    v12 = (v12 * 0)
    v5 = 18446744073709551615
    var v29 uint16 = (v13 ^ (v12 * (^(v11))))
    fmt.Println(v11, v1)
    for fi30 := 0; fi30 < 18; fi30++ {
        v9 = v10
        v5 = (v5 * v1)
        var v31 uint64 = (v7 % ((v7) | 1))
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
