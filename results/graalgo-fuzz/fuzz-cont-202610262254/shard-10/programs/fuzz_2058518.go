package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float32 = (1.5 + 3.14159)
    v2 := []uint32{4294967295, 1, 1}
    v3 := []int{}
    var v4 uint8 = 128
    var v5 int8 = 1
    var v6 int8 = v5
    var v7 int16 = 2
    if false {
        if false {
            var v8 string = ("" + "\n\t")
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: (v7 % ((v7) | 1)), F1: true}
                _ = fi9
                _ = v10
            }
            if (342 != uint(v4)) {
                switch (v8 + v8) {
                case "hello":
                    v2 = append(v2, 4294967295)
                }
            } else {
                var v11 uint32 = 0
                _ = v11
            }
            _ = v8
        }
        var v12 int16 = -32768
        switch ("foo bar" + "a") {
        case "":
            if true {
                if (-32768 > v7) {
                    fmt.Println(v5)
                    if len(v3) > 0 {
                        v3[0] = 984
                    }
                    v13 := GenS2{F0: 255}
                    _ = v13
                } else {
                    fmt.Println(v5, v5)
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        _ = fi14
                    }
                }
                var v15 uint16 = 42
                _ = v15
            } else {
                var v16 bool = ((!true) || true)
                _ = v16
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v12
    }
    var v17 uint = 255
    _ = len(v2)
    for fi18 := 0; fi18 < 2; fi18++ {
        v6 = 42
        if false {
            v5 -= v5
            var v19 uint16 = 255
            v6 = v6
            _ = v19
        } else {
            var v20 uint = 255
            v4 = (^(v4))
            v21 := GenS1{F0: v7, F1: true}
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v5, v6, v6)
        _ = fi22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v7, v7, v7)
        _ = fi23
    }
    var v24 int16 = 1
    v6 = v6
    var v25 uint8 = 90
    fmt.Println(v6, v7, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
