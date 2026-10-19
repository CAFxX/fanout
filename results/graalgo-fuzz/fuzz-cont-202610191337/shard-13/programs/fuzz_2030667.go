package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 string = (((("foo bar" + ("" + (("foo bar" + "foo bar") + "hello"))) + "foo bar") + "hello") + "hello")
    var v2 uint8 = 239
    var v3 uint16 = 2
    var v4 uint8 = (v2 & v2)
    if (1000 <= 796) {
        v5 := GenS1{F0: 654, F1: 0, F2: false}
        _ = v5
    }
    v6 := GenS1{F0: 32767, F1: 0, F2: true}
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "a" {
            case "foo bar":
                if true {
                    var v9 uint16 = 65535
                    var v10 uint8 = 95
                    _ = v10
                    _ = v9
                } else {
                    var v11 float32 = 1e10
                    if false {
                        if (uint16(v4) != (^(v3))) {
                            v2 -= (v2 + v2)
                        } else {
                            v12 := GenS2{F0: 2, F1: 2147483647}
                            _ = v12
                        }
                        v13 := []int{(fi7 - 0), 590}
                        _ = v13
                    }
                    v14 := GenS3{F0: 1, F1: true}
                    _ = v11
                    _ = v14
                }
            case "":
                fmt.Println(v1)
                var v15 bool = (true && (true || false))
                _ = v15
            }
            v16 := []int{(^(fi7)), 32767, (fi8 ^ (fi8 >> ((746) & 63)))}
            if (42 > 0) {
                var v17 int64 = -1
                v2 -= (v2 % (((v2 << ((255) & 63))) | 1))
                switch 985 {
                case -2147483648:
                    v17 = (v17 / ((v17) | 1))
                }
                _ = v17
            } else {
                var v18 uint = 4294967295
                fmt.Println(v4)
                v19 := GenS2{F0: 42, F1: 491}
                _ = v18
                _ = v19
            }
            _ = fi8
            _ = v16
        }
        _ = fi7
    }
    var v20 uint = 1
    var v21 int = 128
    var v22 int32 = 669
    v23 := []int{v21, v21, (v21 ^ v21), v21}
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
