package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint; F2 bool }
type GenS4 struct { F0 int8; F1 int16; F2 string }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16) uint64 {
    fmt.Println(p0)
    switch "a" {
    case "a":
        p0 -= p0
    case "foo bar":
        p0 -= p0
    }
    var v4 string = "hello"
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 uint, p1 uint64) uint32 {
    v5 := GenS1{F0: 1000, F1: 979, F2: ""}
    var v6 uint8 = 1
    v7 := GenS2{F0: true}
    p0 -= p0
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            if (!(false || (127 <= fi8))) {
                var v9 uint = p0
                for fi10 := 0; fi10 < 2; fi10++ {
                    if (!false) {
                        p1 += 287
                        v6 += 238
                    } else {
                        v11 := GenS1{F0: 32767, F1: 725, F2: "hello"}
                        v12 := GenS3{F0: ("foo bar" + "hello"), F1: (p0 - 768), F2: (false || false)}
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
                _ = v9
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v13 float32 = float32(0.1)
    fmt.Println(v1, v13)
    if (!false) {
        switch ("\n\t" + (("a" + ("" + ("hello" + "foo bar"))) + "a")) {
        case "hello":
            v1 += (v1 ^ v1)
        case "\n\t":
            v13 -= ((v13 / float32(-2.25)) * 1e10)
        case "foo bar":
            var v14 uint = (v1 & 18446744073709551615)
            _ = v14
        }
        v15 := []int{1, -321}
        var v16 float64 = 3.14159
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := GenS4{F0: 61, F1: int16(v1), F2: "foo bar"}
        _ = fi17
        _ = v18
    }
    switch (("foo bar" + "\n\t") + ("a" + "")) {
    case "hello":
        v13 = (v13 - v13)
    case "a":
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 int32 = -1
            v21 := GenS5{F0: (v1 - 1)}
            fmt.Println(fi19)
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 int = 32768
    if false {
        v23 := GenS1{F0: -1, F1: 255, F2: ""}
        v1 += 42
        _ = v23
    }
    var v24 string = "a"
    v22 += (v22 / ((v22) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v24
    _ = v3
}
