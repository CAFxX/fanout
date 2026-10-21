package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
func genFn0(p0 int8) int32 {
    fmt.Println(p0)
    if (865 <= 1) {
        if (4294967295 == 22) {
            v6 := GenS1{F0: true, F1: 913, F2: 160}
            p0 = p0
            var v7 int16 = -815
            _ = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(fi9, fi9)
                    _ = fi9
                }
            }
            _ = fi8
        }
        p0 += -1
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 uint, p1 uint8, p2 string) uint64 {
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{127, (fi10 % (((fi10 - fi10)) | 1)), -1, fi10}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: (float32(((float64(0.1) + ((0.1 * 1.5) / (0.1 / 1.0))) - -2.25)) == ((1.5 + (float32(0.0) / (1e10 + 0.0))) / (-2.25 - (-2.25 / -2.25)))), F1: 2147483647, F2: 682}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint32 = 579
    var v2 int32 = 0
    var v3 uint16 = 562
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    v1 = 42
    var v14 float64 = 3.14159
    var v15 int64 = 52
    var v16 uint64 = 42
    if true {
        v2 = genFn0(100)
        v17 := []int{934}
        _ = v17
    } else {
        var v18 int16 = 50
        v19 := v5["foo bar"]
        v18 -= (v18 >> ((0) & 63))
        _ = v18
        _ = v19
    }
    _ = len(v5)
    if false {
        if (127 == 32768) {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 int8 = 1
                var v22 int16 = -760
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS1{F0: true, F1: (v1 & (^(v1))), F2: v16}
            var v24 uint64 = v16
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                if true {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 int = 32768
                        _ = fi26
                        _ = v27
                    }
                }
                v16 -= (v16 ^ (^(v16)))
                delete(v4, fi25)
                _ = fi25
            }
            switch "foo bar" {
            case "\n\t":
                v28 := []int{}
                _ = v28
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                v2 -= -1
                _ = fi29
            }
        }
        v30 := []int{4294967295, 128, -864, -1}
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 6; fi31++ {
            for fi32 := 0; fi32 < 13; fi32++ {
                v16 = (v16 & (v16 / (((^(v16))) | 1)))
                var v33 uint = 42
                _ = fi32
                _ = v33
            }
            var v34 uint64 = (v16 ^ (^(v16)))
            switch "a" {
            case "":
                for fi35 := 0; fi35 < 14; fi35++ {
                    if (0 != 120) {
                        v16 += genFn1(1, 93, (("" + "hello") + ""))
                        v2 += (v2 << ((0) & 63))
                    }
                    var v36 uint32 = v1
                    _ = fi35
                    _ = v36
                }
            case "hello":
                v37 := []int{552, (fi31 + fi31)}
                _ = v37
            case "foo bar":
                v2 = v2
            }
            _ = fi31
            _ = v34
        }
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
