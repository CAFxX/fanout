package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 bool }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int16, p1 string, p2 uint) uint32 {
    v7 := GenS1{F0: 127, F1: "hello"}
    fmt.Println(p2, p2)
    if (529 != 65535) {
        if (1e10 >= 0.0) {
            v8 := []int{}
            _ = v8
        } else {
            var v9 string = ""
            if false {
                v10 := GenS2{F0: true, F1: 679}
                v11 := GenS3{F0: true, F1: -1, F2: false}
                _ = v10
                _ = v11
            } else {
                var v12 uint16 = 0
                _ = v12
            }
            _ = v9
        }
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 19; fi13++ {
                fmt.Println(p0)
                var v14 uint64 = 441
                _ = fi13
                _ = v14
            }
        }
    } else {
        v15 := []int{865, 2147483647}
        var v16 bool = true
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}
func genFn1(p0 uint8) bool {
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint32 = 65535
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS1{F0: 4294967295, F1: "hello"}
            v22 := []int{(fi17 | 0), -2147483648}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi17
    }
    v23 := []int{-2147483648, 32767, 9223372036854775807}
    _ = p0
    _ = v23
    return (!false)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2147483648
    var v3 string = ("\n\t" + ("a" + "a"))
    var v4 int64 = 9223372036854775807
    v5 := make(map[int]int)
    var v6 uint = 255
    switch 42 {
    case -612:
        switch "\n\t" {
        case "foo bar":
            var v24 int = int(v4)
            _ = v24
        default:
            var v25 int64 = 771
            _ = v25
        }
    case 221:
        v2 = v2
    case 128:
        v2 += v2
    default:
        fmt.Println(v4, v4, v4)
    }
    if (!true) {
        var v26 float32 = 0.0
        fmt.Println(v3, v26, v4)
        _ = v26
    } else {
        var v27 uint16 = 2
        _ = v27
    }
    if (true && true) {
        v2 = v2
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS2{F0: (42 > 0), F1: genFn0(-511, "hello", v6)}
            for fi30 := 0; fi30 < 19; fi30++ {
                if (!false) {
                    v2 += v2
                    if (488 == 495) {
                        v31 := GenS2{F0: (!false), F1: 2147483648}
                        fmt.Println(v2)
                        _ = v31
                    }
                    _ = len(v1)
                    var v32 uint16 = 2
                    switch (fi30 << ((4294967295) & 63)) {
                    case 128:
                        v4 -= (-(v4))
                    case -2147483648:
                        switch 42 {
                        case 542:
                            var v33 bool = (62 < 1)
                            _ = v33
                        default:
                            v6 += v6
                        }
                    }
                    _ = v32
                }
                _ = len(v5)
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        v2 = v2
    }
    v2 = v2
    switch (v3 + v3) {
    case "foo bar":
        v2 += (v2 + (^(v2)))
    default:
        v34 := []int{-9223372036854775808, -9223372036854775808, 162}
        _ = v34
    }
    v6 = (v6 << (((v6 >> ((v6) & 63))) & 63))
    v2 = v2
    var v35 float32 = 1.5
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 float32 = float32((1e10 / 1.5))
        v38 := GenS4{F0: (v2 << (((^(v6))) & 63)), F1: 1000}
        _ = fi36
        _ = v37
        _ = v38
    }
    v4 += int64(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
