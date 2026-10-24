package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 uint = 18446744073709551615
    var v4 bool = (!(0 >= -128))
    v5 := make(map[string]uint32)
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 5; fi6++ {
            v4 = false
            _ = fi6
        }
    case "":
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(v2, v4, v3)
            v3 = v3
            _ = fi7
        }
    case "hello":
        v4 = (2147483648 <= 2)
    default:
        v8 := []int{}
        _ = v8
    }
    v3 -= 255
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: fi9, F1: (981 >= 255), F2: 809}
        var v11 float64 = float64(((0.0 / 1e10) - 0.1))
        v2 = (!((128 < 1) || (!(107 != 42))))
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int32 = 1
            if true {
                fmt.Println(v3)
            }
            var v15 int8 = -1
            _ = fi13
            _ = v14
            _ = v15
        }
        switch fi12 {
        case 42:
            if false {
                var v16 uint = (v3 ^ (v3 & 18446744073709551615))
                if (!(v16 >= 2)) {
                    if (128 <= 42) {
                        var v17 uint8 = 255
                        var v18 uint16 = 255
                        v5[("foo bar" + "foo bar")] = 1
                        _ = v17
                        _ = v18
                    } else {
                        var v19 int8 = 42
                        _ = v19
                    }
                    switch "\n\t" {
                    case "\n\t":
                        for fi20 := 0; fi20 < 10; fi20++ {
                            v21 := GenS2{F0: 100, F1: 42}
                            v4 = ("hello" > "a")
                            fmt.Println(v3)
                            _ = fi20
                            _ = v21
                        }
                    }
                }
                _ = v16
            }
        }
        if (!v2) {
            fmt.Println(fi12, v3)
        }
        _ = fi12
    }
    v4 = (!(1 == -651))
    var v22 uint16 = 243
    v3 = v3
    if v4 {
        switch "\n\t" {
        case "foo bar":
            var v23 uint64 = 42
            _ = v23
        }
    }
    if (!v4) {
        v24 := []int{}
        _ = v24
    } else {
        v25 := GenS1{F0: 128, F1: (0.1 != 3.14159), F2: 416}
        v26 := GenS2{F0: -1, F1: 534}
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint64 = 18446744073709551615
            var v29 int = (fi27 << (((v3 % ((1) | 1))) & 63))
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v25
        _ = v26
    }
    var v30 string = ("a" + "\n\t")
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
