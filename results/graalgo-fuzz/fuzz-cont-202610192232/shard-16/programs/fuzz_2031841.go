package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 string }
func genFn0(p0 uint16) uint {
    var v4 float32 = 1.5
    v4 = (v4 - v4)
    v5 := GenS1{F0: 4294967295, F1: (p0 + (p0 >> ((1) & 63)))}
    _ = p0
    _ = v4
    _ = v5
    return 899
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 2
    var v3 uint = v2
    v2 = 537
    v3 = genFn0(42)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint8 = 49
        if true {
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(v7, v2)
                var v9 float32 = float32(1e10)
                _ = len(v1)
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 2147483647, F1: 590}
            _ = v10
        } else {
            v11 := []int{(fi6 & (fi6 >> ((genFn0(32767)) & 63)))}
            for fi12 := 0; fi12 < 14; fi12++ {
                switch "hello" {
                case "":
                    v13 := GenS1{F0: uint32(v7), F1: 42}
                    _ = v13
                }
                var v14 int32 = 2147483647
                _ = fi12
                _ = v14
            }
            v2 -= v3
            _ = v11
        }
        var v15 uint16 = 32767
        _ = fi6
        _ = v15
        _ = v7
    }
    var v16 int = 465
    v17 := GenS2{F0: 366, F1: 100}
    var v18 float64 = float64((((0.0 + (1e10 / float32(1.5))) * 3.14159) / 0.1))
    v19 := GenS2{F0: 2, F1: 42}
    fmt.Println(v2)
    if false {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch (((("a" + ("\n\t" + ("foo bar" + (("hello" + (((("foo bar" + "\n\t") + "") + "a") + (("" + "a") + "a"))) + ("\n\t" + "hello"))))) + "a") + ("" + "")) + ("a" + "a")) {
            case "foo bar":
                v16 -= v16
            case "a":
                v21 := GenS1{F0: 1, F1: 791}
                _ = v21
            case "":
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(v2)
                    _ = fi22
                }
            }
            _ = fi20
        }
        var v23 string = "\n\t"
        if false {
            var v24 uint16 = 2
            var v25 int64 = 0
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 3; fi26++ {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 int = fi26
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 14; fi29++ {
                    v23 = v23
                    _ = fi29
                }
                _ = fi26
            }
            var v30 int8 = 127
            _ = v30
        }
        _ = v23
    }
    var v31 int8 = 0
    v32 := GenS3{F0: v31, F1: 2147483647}
    if false {
        v33 := GenS4{F0: 157, F1: 4294967295, F2: 0}
        _ = v33
    }
    if (966 != v16) {
        v34 := []int{}
        if (32768 < 32767) {
            _ = len(v1)
        } else {
            var v35 uint16 = 32767
            _ = v35
        }
        _ = v34
    } else {
        v36 := GenS3{F0: (v31 - (v31 + (v31 + v31))), F1: -2147483648}
        var v37 int64 = 538
        _ = v36
        _ = v37
    }
    v38 := GenS5{F0: 127, F1: -9223372036854775808, F2: ("foo bar" + "foo bar")}
    v2 += (v3 ^ 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
}
