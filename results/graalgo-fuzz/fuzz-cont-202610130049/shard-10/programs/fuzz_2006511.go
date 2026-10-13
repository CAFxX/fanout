package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint; F1 int64 }
func genFn0(p0 uint32, p1 bool, p2 int64) uint64 {
    p2 = p2
    v4 := GenS1{F0: -149}
    switch -878 {
    case 4294967295:
        var v5 float64 = ((3.14159 - float64(0.0)) * (1e10 - 0.1))
        _ = v5
    case -111:
        p1 = true
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{(fi6 >> ((65535) & 63)), 1}
        switch "\n\t" {
        case "a":
            _ = len(v7)
        case "hello":
            var v8 int8 = -128
            _ = v8
        case "\n\t":
            if p1 {
                switch fi6 {
                case -528:
                    var v9 uint16 = 540
                    _ = v9
                default:
                    var v10 string = "\n\t"
                    _ = v10
                }
                var v11 bool = false
                fmt.Println(p0, fi6)
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1() bool {
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS2{F0: fi12, F1: 42, F2: 1}
        v14 := []int{-2147483648, (fi12 % ((fi12) | 1)), fi12}
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS3{F0: true, F1: 0}
            var v17 uint8 = 0
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v18 := []int{32768}
    _ = v18
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v19 uint = 1
    v20 := GenS4{F0: (("hello" + (("" + "\n\t") + ("\n\t" + ("hello" + ((((("\n\t" + ("a" + "")) + ("foo bar" + "foo bar")) + "") + "\n\t") + ""))))) + "hello"), F1: 2}
    v19 += 65535
    switch 4294967295 {
    case -425:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 int32 = -1
            var v23 uint8 = 2
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 int8 = 1
                _ = fi24
                _ = v25
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    case 42:
        fmt.Println(v19)
    case -1:
        if false {
            v26 := GenS5{F0: -1}
            v27 := GenS6{F0: 871}
            _ = v26
            _ = v27
        } else {
            v28 := GenS7{F0: 127, F1: 2}
            _ = v28
        }
    }
    fmt.Println(v19)
    v19 = (v19 & (v19 << ((464) & 63)))
    v29 := GenS4{F0: "hello", F1: 114}
    var v30 int16 = 1
    if (!(!(v19 == 18446744073709551615))) {
        fmt.Println(v19)
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int16 = -215
                v19 = v19
                v19 -= (v19 - v19)
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    } else {
        switch 4294967295 {
        case -9223372036854775808:
            delete(v1, ("hello" + "hello"))
        case -665:
            v30 = v30
        }
        fmt.Println(v30, v19)
        v34 := []int{430, 928, 127, 127}
        _ = v34
    }
    var v35 bool = (false || (false || true))
    v19 += (v19 - 58)
    v36 := GenS8{F0: 65535, F1: 991}
    var v37 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
}
