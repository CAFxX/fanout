package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 int8, p1 uint32, p2 int8) int16 {
    fmt.Println(p2, p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = p0
        v6 := []int{fi5, fi5}
        fmt.Println(p1)
        _ = fi5
        _ = v6
    }
    switch -2147483648 {
    case -2147483648:
        var v7 int32 = 119
        _ = v7
    case 127:
        p1 = 1
    }
    switch "a" {
    case "":
        p2 -= (p2 >> ((18446744073709551615) & 63))
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            p1 -= 2
            fmt.Println(p1, fi8, p2)
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint16 {
    var v9 int16 = 32767
    fmt.Println(v9)
    if false {
        switch 42 {
        case 575:
            v9 = genFn0(int8(v9), 2147483647, 127)
        default:
            v10 := GenS1{F0: -543, F1: 253, F2: 128}
            _ = v10
        }
    }
    _ = v9
    return 65535
}

func main() {
    var v1 bool = true
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]string)
    var v4 int8 = 100
    switch 32767 {
    case 1:
        switch ("" + "hello") {
        case "a":
            if true {
                switch "foo bar" {
                case "\n\t":
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v4 += v4
                        v1 = v1
                        _ = fi11
                    }
                }
                var v12 int32 = -1
                _ = v12
            }
        case "":
            var v13 uint16 = 827
            _ = v13
        default:
            switch "" {
            case "":
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 int = (-(fi14))
                    var v16 float64 = (1e10 + 0.0)
                    switch "" {
                    case "":
                        v15 += (fi14 - v15)
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        }
    }
    if v1 {
        var v17 string = "\n\t"
        _ = v17
    }
    var v18 uint32 = 4294967295
    var v19 int64 = (v2 * -9223372036854775808)
    v19 = 0
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint32 = (v18 * (v18 | v18))
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint16 = 905
            if true {
                v2 += v19
                switch (fi22 << ((0) & 63)) {
                case 42:
                    v24 := GenS1{F0: -2147483648, F1: uint8(v21), F2: 116}
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        var v25 int32 = 2147483647
                        _ = v25
                    }
                    fmt.Println(v23)
                    _ = v24
                case 2147483647:
                    v26 := []int{32767, 420, fi22}
                    _ = v26
                case 127:
                    _ = len(v3)
                }
            }
            _ = fi22
            _ = v23
        }
        if (((v4 + v4) < (v4 + v4)) || v1) {
            var v27 int8 = v4
            fmt.Println(v18, v27)
            v28 := GenS1{F0: 176, F1: 127, F2: 132}
            _ = v27
            _ = v28
        } else {
            switch ("\n\t" + "hello") {
            case "\n\t":
                var v29 bool = (!(1 != 0))
                _ = v29
            case "foo bar":
                fmt.Println(v19, v21)
            default:
                var v30 uint8 = 255
                _ = v30
            }
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v18)
    v31 := GenS1{F0: 1, F1: 1, F2: 1}
    fmt.Println(v18, v2, v4)
    fmt.Println(v1, v4)
    var v32 uint = 1
    var v33 float32 = 1e10
    for fi34 := 0; fi34 < 11; fi34++ {
        v1 = false
        v1 = v1
        v4 += (v4 | (v4 & 42))
        _ = fi34
    }
    v35 := GenS2{F0: -9223372036854775808, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
}
