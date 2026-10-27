package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint16) int32 {
    var v5 int8 = 42
    v5 += (v5 % ((44) | 1))
    v6 := []int{-9223372036854775808, 416}
    p0 = (p0 + (p0 ^ p0))
    _ = p0
    _ = v5
    _ = v6
    return -192
}
func genFn1(p0 uint16, p1 int8, p2 uint32) int32 {
    p1 = (p1 << ((255) & 63))
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 32767
    }
    switch "foo bar" {
    case "hello":
        var v8 float32 = 1e10
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := []string{("a" + ("hello" + ("hello" + "a"))), "\n\t", "\n\t", "hello"}
    if (v1 || true) {
        delete(v3, 42)
        var v9 uint8 = 114
        _ = v9
    }
    var v10 uint32 = 139
    var v11 int8 = -1
    v12 := GenS1{F0: (v11 - -10), F1: 72}
    v10 = (v10 ^ (v10 * (v10 << ((255) & 63))))
    v11 += (v11 & 100)
    for fi13 := 0; fi13 < 5; fi13++ {
        switch ("" + "foo bar") {
        case "hello":
            v14 := GenS2{F0: -951}
            _ = v14
        default:
            v1 = ((-1 < genFn0(0)) || true)
        }
        var v15 bool = (!v1)
        _ = fi13
        _ = v15
    }
    var v16 int32 = -770
    v17 := GenS1{F0: (v11 ^ 127), F1: 1}
    v18 := GenS2{F0: 32768}
    v1 = (true || v1)
    if (0 > genFn0(0)) {
        for fi19 := 0; fi19 < 6; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 float32 = (1e10 * 1e10)
                _ = fi20
                _ = v21
            }
            fmt.Println(fi19)
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 uint16 = 255
                v1 = true
                _ = fi22
                _ = v23
            }
            _ = fi19
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 string = ("" + ("a" + "foo bar"))
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        switch ("hello" + ("\n\t" + "a")) {
        case "a":
            v16 -= (v16 & genFn1(1, v11, (v10 / (((^(v10))) | 1))))
        case "\n\t":
            if (!(v11 != v11)) {
                var v27 uint64 = 18446744073709551615
                v11 = (v11 + v11)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v27
            } else {
                if (true && (934 <= 593)) {
                    fmt.Println(v10)
                    switch "a" {
                    case "foo bar":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            fmt.Println(v10, fi28)
                            switch (fi28 & (fi28 % ((-1) | 1))) {
                            case 4294967295:
                                var v29 uint8 = 22
                                _ = v29
                            default:
                                fmt.Println(v16)
                            }
                            _ = fi28
                        }
                    case "a":
                        v10 += (v10 & v10)
                    }
                } else {
                    fmt.Println(v1, fi26)
                    v30 := GenS2{F0: fi26}
                    _ = v30
                }
            }
        }
        fmt.Println(v11, v10)
        var v31 uint64 = 1
        _ = fi26
        _ = v31
    }
    var v32 int = -394
    var v33 int32 = 344
    if len(v4) > 0 {
        v4[0] = (("a" + "") + ((("" + ("\n\t" + "\n\t")) + "\n\t") + ((("a" + ("hello" + "hello")) + ("hello" + "")) + "a")))
    }
    v33 = int32(v10)
    if (!true) {
        var v34 uint16 = 639
        if (1e10 == 3.14159) {
            if (65535 > 560) {
                var v35 int16 = -32768
                _ = v35
            }
            for fi36 := 0; fi36 < 9; fi36++ {
                switch (v32 + v32) {
                case 127:
                    var v37 int16 = -1
                    _ = v37
                default:
                    var v38 float32 = 1e10
                    _ = v38
                }
                for fi39 := 0; fi39 < 18; fi39++ {
                    fmt.Println(v11, fi39, v10)
                    v40 := v3[v10]
                    _ = len(v2)
                    _ = fi39
                    _ = v40
                }
                fmt.Println(v11)
                _ = fi36
            }
        }
        v33 -= v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
