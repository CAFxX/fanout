package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32, p1 bool) int16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 255, F1: 255, F2: -840}
        p1 = false
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 628
    _ = p0
    _ = p1
    _ = v7
    return 43
}
func genFn1() string {
    var v8 int = -840
    var v9 bool = false
    v9 = (0.1 > 1.5)
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v1 -= (v1 >> ((4294967295) & 63))
    switch ("a" + "\n\t") {
    case "hello":
        switch ((((((("a" + "hello") + ("a" + "foo bar")) + "") + (("hello" + ((("hello" + "foo bar") + "") + ("foo bar" + (("\n\t" + "") + "")))) + ("hello" + "\n\t"))) + "hello") + "") + ("" + "")) {
        case "a":
            v1 = genFn0(42, false)
        default:
            var v10 string = "a"
            _ = v10
        }
    case "a":
        fmt.Println(v1)
    case "foo bar":
        v2[32768] = 1
    default:
        fmt.Println(v1)
    }
    v1 = genFn0(42, true)
    switch 849 {
    case 603:
        if (!(!false)) {
            delete(v4, 0)
        } else {
            v11 := GenS2{F0: 0, F1: false, F2: 295}
            _ = v11
        }
    case 9223372036854775807:
        _ = len(v3)
    case 32768:
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{(fi12 << ((1) & 63))}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if false {
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 int8 = 0
                    v1 = genFn0(-1, (!(v15 > v15)))
                    var v16 uint8 = 2
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v1 += -1
            }
            _ = fi12
            _ = v13
        }
    }
    var v17 int16 = 705
    fmt.Println(v1, v1)
    var v18 int64 = 0
    v19 := GenS3{F0: 18446744073709551615}
    switch 128 {
    case 128:
        v20 := GenS2{F0: -57, F1: (3.14159 > 1.5), F2: 31}
        _ = v20
    case 9223372036854775807:
        switch "foo bar" {
        case "hello":
            if false {
                v21 := []int{-2147483648}
                switch "foo bar" {
                case "a":
                    _ = len(v2)
                case "\n\t":
                    if len(v21) > 0 {
                        v21[0] = int(v1)
                    }
                case "":
                    v18 = v18
                }
                _ = v21
            }
        default:
            v17 += v1
        }
    }
    v22 := GenS4{F0: 42, F1: 604}
    v23 := GenS5{F0: 42, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
