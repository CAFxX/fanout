package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = (!(true || (true || ((((((((("hello" + "foo bar") + "\n\t") + "\n\t") + ("hello" + "a")) + ("hello" + ("" + "hello"))) + "\n\t") + "foo bar") < "") && (0 != 42)))))
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 2
    v5 := GenS1{F0: 0}
    v4 = 1
    switch 32767 {
    case 699:
        v1 = (-1 == 0)
    case -1:
        var v6 string = "a"
        _ = v6
    default:
        fmt.Println(v1, v2)
    }
    fmt.Println(v4, v2, v2)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int8 = 100
        _ = fi7
        _ = v8
    }
    v1 = true
    v9 := GenS2{F0: 65535, F1: "", F2: -128}
    if (!false) {
        v1 = false
        fmt.Println(v1, v4)
    } else {
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: 346}
            _ = v10
        case "a":
            var v11 int8 = 42
            _ = v11
        case "\n\t":
            if (!false) {
                v3[2147483648] = 681
                v12 := []int{4294967295, 32767}
                fmt.Println(v1)
                _ = v12
            }
        default:
            v4 = (v4 - v4)
        }
        v13 := GenS1{F0: 2}
        _ = v13
    }
    switch 128 {
    case 42:
        switch "hello" {
        case "\n\t":
            var v14 uint32 = 626
            _ = v14
        case "hello":
            if v1 {
                v4 = v4
                var v15 float32 = float32((3.14159 / 1.5))
                _ = v15
            } else {
                v4 -= v4
            }
        default:
            v1 = (v2 <= "foo bar")
        }
    case -9223372036854775808:
        switch v2 {
        case "foo bar":
            switch v2 {
            case "foo bar":
                for fi16 := 0; fi16 < 12; fi16++ {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v1 = v1
                        v2 = "\n\t"
                        _ = fi17
                    }
                    _ = fi16
                }
            }
        case "hello":
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS1{F0: 2147483648}
                _ = fi18
                _ = v19
            }
        }
    default:
        var v20 uint16 = 0
        _ = v20
    }
    v21 := GenS3{F0: 106}
    if (!(0 <= -1)) {
        switch 128 {
        case 42:
            var v22 int16 = 32767
            _ = v22
        case 0:
            v23 := []int{32768, 2147483647, 127, -509}
            _ = v23
        default:
            v1 = ((128 >= 127) || v1)
        }
    }
    var v24 uint8 = 0
    if false {
        v2 = v2
        var v25 string = v2
        switch 32767 {
        case 2147483647:
            var v26 int32 = 0
            _ = v26
        }
        _ = v25
    } else {
        v2 = v2
        if (!v1) {
            v4 -= (v4 + 0)
            v24 = 127
        }
        if v1 {
            v24 += (v24 | (^(v24)))
            fmt.Println(v2, v2, v4)
        }
    }
    if ((v24 < (^(v24))) || ((!v1) && v1)) {
        v24 = v24
        if true {
            v27 := []int{-131}
            var v28 int = 9223372036854775807
            if ((v28 ^ v28) < v28) {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS3{F0: 0}
                    _ = fi29
                    _ = v30
                }
            }
            _ = v27
            _ = v28
        }
    }
    v4 = 651
    v4 += v4
    switch 9223372036854775807 {
    case -80:
        v4 = 1000
    default:
        var v31 int16 = -32768
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
