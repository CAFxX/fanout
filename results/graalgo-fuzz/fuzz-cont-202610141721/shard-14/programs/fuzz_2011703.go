package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint8 = 123
    var v4 int = 0
    v5 := make(map[uint32]int)
    var v6 int = (v4 + v4)
    var v7 int8 = -1
    for fi8 := 0; fi8 < 19; fi8++ {
        switch 206 {
        case 127:
            var v9 int8 = (^(v7))
            _ = v9
        case -715:
            for fi10 := 0; fi10 < 11; fi10++ {
                v4 -= v6
                v4 = fi10
                v7 += -1
                _ = fi10
            }
        default:
            for fi11 := 0; fi11 < 2; fi11++ {
                if true {
                    v12 := GenS1{F0: 978}
                    delete(v2, fi8)
                    _ = v12
                }
                v6 = (fi11 - fi8)
                v7 = (v7 - (v7 - (v7 + v7)))
                _ = fi11
            }
        }
        var v13 int8 = 100
        var v14 uint64 = 911
        _ = fi8
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 255}
    switch "hello" {
    case "\n\t":
        if (!(("\n\t" < "hello") && false)) {
            switch "\n\t" {
            case "":
                v6 = 127
            default:
                v16 := GenS2{F0: 32767, F1: 51}
                _ = v16
            }
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := []int{(v4 ^ 1)}
                var v19 float32 = (-2.25 / -2.25)
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int = (v4 ^ 2147483647)
            _ = v20
        } else {
            if (3.14159 == 0.0) {
                v6 = v4
                v21 := GenS3{F0: 560, F1: 18446744073709551615}
                _ = v21
            } else {
                v7 += v7
            }
        }
    case "a":
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(fi22)
            _ = fi22
        }
    }
    v6 += (v4 % ((v4) | 1))
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int8 = (v7 - v7)
        v3 -= 127
        _ = len(v1)
        _ = fi23
        _ = v24
    }
    v4 = (v6 ^ v6)
    fmt.Println(v6, v6)
    v7 += (v7 + (v7 * (v7 - (-(v7)))))
    v25 := GenS4{F0: (220 > 18446744073709551615), F1: 4294967295, F2: 32767}
    var v26 int16 = 32767
    switch v4 {
    case 625:
        var v27 bool = false
        _ = v27
    }
    if (597 < 0) {
        v6 = int(v3)
        v28 := GenS5{F0: 42}
        _ = v28
    }
    var v29 uint16 = 864
    switch (((((("\n\t" + "hello") + "foo bar") + (((((((("foo bar" + "foo bar") + ("foo bar" + ("" + "hello"))) + "\n\t") + "hello") + "foo bar") + "a") + "hello") + ("a" + ("\n\t" + "a")))) + "\n\t") + "\n\t") + ("\n\t" + "")) {
    case "":
        for fi30 := 0; fi30 < 11; fi30++ {
            v6 -= int(v3)
            v7 = v7
            _ = fi30
        }
    default:
        v7 -= 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
