package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 int16; F2 int64 }
func genFn0(p0 uint) bool {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 817}
        _ = fi5
        _ = v6
    }
    var v7 float64 = 1.5
    var v8 bool = true
    _ = p0
    _ = v7
    _ = v8
    return (v8 || (!(843 != -791)))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 uint16 = 32767
    v4 := make(map[string]bool)
    if (1 == -218) {
        v1[128] = (!true)
    } else {
        if false {
            fmt.Println(v3)
        }
        v9 := GenS1{F0: 755}
        v3 = 65535
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: -1}
        v12 := []int{fi10, fi10}
        var v13 uint16 = v3
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v3 = (^(v3))
        _ = fi14
    }
    if (((57 > -1) && (false && (2 != 2))) || false) {
        v15 := GenS2{F0: 4294967295, F1: -268}
        v16 := []int{-778, 32767, 784}
        var v17 int8 = -128
        _ = v15
        _ = v16
        _ = v17
    }
    v3 = 660
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 = 807
        var v19 int = fi18
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint64 = 42
            v19 = (fi20 ^ -1)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    var v22 int16 = 32767
    var v23 uint8 = 255
    for fi24 := 0; fi24 < 10; fi24++ {
        v3 = (v3 | v3)
        for fi25 := 0; fi25 < 13; fi25++ {
            switch "hello" {
            case "":
                v26 := []int{int(v23)}
                _ = v26
            case "a":
                delete(v2, "hello")
                var v27 int = (fi24 ^ fi25)
                v28 := GenS3{F0: 32767, F1: v22, F2: 89}
                _ = v27
                _ = v28
            default:
                var v29 float64 = -2.25
                _ = v29
            }
            v30 := v4["a"]
            _ = fi25
            _ = v30
        }
        _ = fi24
    }
    switch "\n\t" {
    case "foo bar":
        v23 -= v23
    default:
        switch ("" + ("foo bar" + "foo bar")) {
        case "\n\t":
            v22 = 1
        case "hello":
            for fi31 := 0; fi31 < 15; fi31++ {
                fmt.Println(v22)
                delete(v4, (("" + "\n\t") + "a"))
                _ = fi31
            }
        }
    }
    v3 = (^(v3))
    fmt.Println(v23, v22, v22)
    if ((false || false) || (false || true)) {
        if true {
            var v32 uint = uint(v3)
            var v33 uint64 = 0
            v34 := []int{}
            _ = v32
            _ = v33
            _ = v34
        } else {
            var v35 uint8 = 42
            var v36 float32 = 0.1
            _ = v35
            _ = v36
        }
    }
    var v37 int16 = v22
    var v38 uint32 = 4294967295
    v39 := []int{-1, 32768}
    for fi40 := 0; fi40 < 3; fi40++ {
        v23 -= (v23 / ((128) | 1))
        delete(v1, (fi40 | -202))
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
