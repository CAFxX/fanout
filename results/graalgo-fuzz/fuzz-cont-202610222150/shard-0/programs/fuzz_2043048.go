package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    var v4 string = ""
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    fmt.Println(v4)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = 0.1
        v9 := GenS1{F0: (127 < 0)}
        v4 = "\n\t"
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "hello":
        if (0 > 42) {
            for fi10 := 0; fi10 < 7; fi10++ {
                if true {
                    fmt.Println(v2)
                }
                fmt.Println(v2)
                _ = fi10
            }
            delete(v3, (v4 + v4))
        }
    case "":
        fmt.Println(v2)
    }
    var v11 uint16 = 655
    if ((^(v11)) < v11) {
        if false {
            _ = len(v5)
            v11 = (^(v11))
        } else {
            if (!true) {
                var v12 string = "a"
                _ = v12
            } else {
                v2 -= v2
                var v13 int = 382
                switch (v13 - (v13 ^ v13)) {
                case -640:
                    var v14 int32 = 0
                    _ = v14
                case 32767:
                    v2 += (v2 * (v2 + v2))
                default:
                    if ((false || (!true)) && (!false)) {
                        v4 = "\n\t"
                        var v15 uint16 = v11
                        _ = v15
                    }
                }
                _ = v13
            }
        }
        if false {
            for fi16 := 0; fi16 < 6; fi16++ {
                _ = len(v6)
                _ = fi16
            }
        } else {
            v17 := []int{}
            v18 := GenS2{F0: 128, F1: 255}
            var v19 int = 1
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 uint16 = (v11 >> ((667) & 63))
    v21 := GenS1{F0: ((false && true) && true)}
    var v22 float32 = v2
    var v23 bool = (v4 < v4)
    if true {
        _ = len(v3)
        fmt.Println(v2)
        if true {
            var v24 float64 = 1e10
            for fi25 := 0; fi25 < 2; fi25++ {
                for fi26 := 0; fi26 < 10; fi26++ {
                    _ = len(v6)
                    var v27 uint8 = 2
                    _ = fi26
                    _ = v27
                }
                v28 := []int{}
                _ = fi25
                _ = v28
            }
            _ = v24
        } else {
            v4 = (v4 + (v4 + (v4 + ("" + v4))))
        }
    }
    var v29 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
