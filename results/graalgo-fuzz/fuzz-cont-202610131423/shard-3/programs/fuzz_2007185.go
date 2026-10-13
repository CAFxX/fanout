package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint16 {
    v4 := GenS1{F0: -733}
    switch "a" {
    case "\n\t":
        var v5 uint64 = 717
        _ = v5
    case "hello":
        v6 := GenS1{F0: 1}
        _ = v6
    }
    _ = v4
    return 576
}
func genFn1(p0 int, p1 uint) uint8 {
    fmt.Println(p0)
    var v7 bool = false
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            switch fi9 {
            case 164:
                if v7 {
                    v10 := []int{42, fi9, p0, p0}
                    _ = v10
                } else {
                    v11 := []int{}
                    var v12 float32 = (0.1 - float32(1.5))
                    _ = v11
                    _ = v12
                }
            case 1:
                v7 = v7
            }
            var v13 int = p0
            _ = fi9
            _ = v13
        }
        var v14 int8 = 100
        _ = fi8
        _ = v14
    }
    fmt.Println(p0, v7)
    switch ("" + "hello") {
    case "a":
        v7 = (("foo bar" + "a") <= "\n\t")
    case "foo bar":
        v15 := GenS1{F0: 9223372036854775807}
        _ = v15
    case "\n\t":
        var v16 uint16 = 0
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 1
    v3 := make(map[int]bool)
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 = (v1 >> ((v2) & 63))
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: 0}
            if true {
                v1 = 128
                switch (fi17 >> (((v2 / (((v2 - (v2 / ((65535) | 1)))) | 1))) & 63)) {
                case 128:
                    fmt.Println(v2, fi18)
                    var v20 int8 = 1
                    fmt.Println(fi17, v1)
                    _ = v20
                case 9223372036854775807:
                    switch "\n\t" {
                    case "hello":
                        v1 = fi17
                    default:
                        if (!true) {
                            var v21 float32 = 0.0
                            _ = v21
                        } else {
                            fmt.Println(v1, v1)
                        }
                    }
                }
                if true {
                    var v22 int64 = 9223372036854775807
                    _ = v22
                }
            }
            _ = fi18
            _ = v19
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            switch "a" {
            case "\n\t":
                v1 = fi23
            }
            v24 := GenS1{F0: 0}
            _ = fi23
            _ = v24
        }
        _ = fi17
    }
    _ = len(v3)
    fmt.Println(v1, v2)
    var v25 int64 = -9223372036854775808
    v3[(v1 + -1)] = (((-(v1)) < v1) || (!(!(4294967295 >= 303))))
    var v26 int8 = int8(v25)
    for fi27 := 0; fi27 < 5; fi27++ {
        v2 = (^(v2))
        _ = fi27
    }
    v25 -= (v25 + -9223372036854775808)
    v25 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
