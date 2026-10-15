package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    var v3 int8 = -128
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 uint = 255
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v7 := GenS1{F0: (v3 - (v3 - -1)), F1: 2147483648, F2: -2147483648}
        _ = v7
    case "a":
        v8 := GenS1{F0: v3, F1: 949, F2: 42}
        _ = v8
    }
    fmt.Println(v6)
    var v9 float32 = ((float32(((0.1 - ((-2.25 * 0.0) + float64(((0.0 - -2.25) - (-2.25 * float32(-2.25)))))) - -2.25)) + float32((-2.25 + float64(-2.25)))) * 1e10)
    fmt.Println(v1, v6, v3)
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: 1, F1: fi10}
        _ = fi10
        _ = v11
    }
    v6 = (v6 / ((v6) | 1))
    if (((-(v3)) >= v3) && false) {
        v12 := []int{42, -532, -1}
        v13 := v2[-286]
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: (^(v3)), F1: 4294967295, F2: 234}
        _ = v14
    }
    var v15 uint = v6
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        switch ("hello" + "foo bar") {
        case "":
            v1 = 1
        }
    }
    v6 = (v15 + (v15 * v6))
    v9 = 1e10
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint16 = v1
        v18 := []int{}
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        if (!true) {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int8 = v3
                var v22 uint8 = 1
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi19
    }
    var v23 uint64 = 944
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
