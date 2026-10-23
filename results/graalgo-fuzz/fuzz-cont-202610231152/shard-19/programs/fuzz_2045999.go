package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0() string {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 0
        _ = fi5
        _ = v6
    }
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 694, F1: 0}
        var v10 float32 = (1e10 + v7)
        switch "foo bar" {
        case "":
            var v11 float32 = (((float32(float64(v10)) - v7) - -2.25) / -2.25)
            _ = v11
        case "a":
            var v12 float32 = v10
            _ = v12
        case "\n\t":
            v7 = 0.1
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return "hello"
}
func genFn1() uint32 {
    var v13 int8 = -128
    v14 := []int{2147483647, 134, -1, 128}
    _ = len(v14)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = v13
    _ = v14
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 1
    v3 := make(map[int]bool)
    var v4 int8 = 93
    fmt.Println(v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint16 = 548
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int = fi17
        _ = len(v1)
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 65535, F1: 1}
    for fi20 := 0; fi20 < 5; fi20++ {
        v4 = (v4 ^ (^(v4)))
        v21 := GenS1{F0: 42, F1: v4}
        fmt.Println(v2)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if false {
            v2 += 1
            v23 := GenS2{F0: 255, F1: 4294967295}
            fmt.Println(fi22, fi22, fi22)
            _ = v23
        } else {
            v4 = -1
        }
        fmt.Println(v4, v4, fi22)
        var v24 string = "hello"
        _ = fi22
        _ = v24
    }
    v2 = v2
    v2 = (v2 >> ((1) & 63))
    for fi25 := 0; fi25 < 2; fi25++ {
        v1[("a" + ("hello" + "hello"))] = fi25
        var v26 int16 = 1000
        _ = fi25
        _ = v26
    }
    v2 -= v2
    v4 = v4
    for fi27 := 0; fi27 < 9; fi27++ {
        v2 = v2
        v2 = (v2 >> ((241) & 63))
        v4 -= 100
        _ = fi27
    }
    var v28 uint16 = 65535
    if (559 > -958) {
        switch (("a" + ("\n\t" + "foo bar")) + "\n\t") {
        case "hello":
            v4 = 0
        case "":
            var v29 uint8 = v2
            _ = v29
        }
    } else {
        v4 += (v4 | (v4 >> ((2) & 63)))
        if (!(42 >= v2)) {
            v28 = 0
            for fi30 := 0; fi30 < 14; fi30++ {
                delete(v1, "")
                v31 := GenS2{F0: 4294967295, F1: uint32(v2)}
                _ = fi30
                _ = v31
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                v32 := GenS2{F0: 2, F1: 2}
                _ = v32
            case "a":
                var v33 int16 = 1000
                _ = v33
            }
            var v34 string = ("\n\t" + "foo bar")
            switch v34 {
            case "hello":
                var v35 int32 = 0
                _ = v35
            case "foo bar":
                var v36 float64 = ((0.0 - -2.25) + 1.5)
                _ = v36
            default:
                v2 += 128
            }
            _ = v34
        }
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
