package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 127
    var v3 uint8 = 128
    var v4 uint64 = 596
    fmt.Println(v3, v2, v4)
    if (false && false) {
        fmt.Println(v3)
        var v5 bool = false
        var v6 uint16 = 443
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: uint(v4)}
    switch "hello" {
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            if (!true) {
                v9 := []int{}
                v10 := GenS2{F0: 2, F1: 1, F2: "\n\t"}
                _ = v10
                _ = v9
            } else {
                if false {
                    fmt.Println(v4)
                } else {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v12 := GenS1{F0: 579}
                        if ((255 >= v3) && (true && false)) {
                            var v13 uint32 = 617
                            _ = v13
                        } else {
                            var v14 float32 = (float32(3.14159) - 1.5)
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                    delete(v1, 249)
                    v15 := GenS1{F0: 65535}
                    _ = v15
                }
            }
            v16 := []int{(fi8 >> ((509) & 63))}
            delete(v1, 2147483647)
            _ = fi8
            _ = v16
        }
    }
    v17 := GenS3{F0: v4, F1: false, F2: false}
    v18 := GenS4{F0: uint16(v3)}
    switch "a" {
    case "a":
        var v19 int = (v2 & v2)
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := []int{(v2 << ((0) & 63)), (v2 / (((-(v2))) | 1))}
        for fi22 := 0; fi22 < 6; fi22++ {
            switch ("a" + "foo bar") {
            case "hello":
                v23 := GenS4{F0: 1}
                _ = v23
            case "":
                var v24 bool = (1.5 > 3.14159)
                _ = v24
            case "\n\t":
                var v25 float64 = -2.25
                _ = v25
            }
            _ = fi22
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(v4, v4, v4)
            v27 := GenS5{F0: -1}
            _ = fi26
            _ = v27
        }
        _ = fi20
        _ = v21
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(fi28, fi28)
        _ = fi28
    }
    v29 := GenS6{F0: 885, F1: (v2 | (v2 % ((v2) | 1))), F2: 1}
    switch (v2 * v2) {
    case 32767:
        for fi30 := 0; fi30 < 9; fi30++ {
            v31 := []int{(^(fi30)), (v2 & v2)}
            v2 -= 128
            _ = fi30
            _ = v31
        }
    case 42:
        fmt.Println(v4, v4)
    case -1:
        v32 := v1[2147483647]
        _ = v32
    }
    v33 := GenS3{F0: (v4 * v4), F1: true, F2: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v7
}
