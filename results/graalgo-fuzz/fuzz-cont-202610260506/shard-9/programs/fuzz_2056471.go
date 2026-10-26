package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0(p0 bool) int64 {
    var v5 int8 = 49
    v6 := []int{263, 926, 9223372036854775807, 1}
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint32 {
    var v7 uint8 = 31
    var v8 int64 = 1
    var v9 int32 = -969
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -555
    var v3 uint32 = 2147483647
    var v4 int16 = 32767
    var v10 int8 = 115
    delete(v1, ("\n\t" + (("foo bar" + (("hello" + ("\n\t" + (((("a" + "\n\t") + (("foo bar" + "") + "hello")) + ("\n\t" + "a")) + ""))) + "\n\t")) + ((((("" + "") + (("" + (("foo bar" + ("\n\t" + "foo bar")) + ("" + ""))) + ("hello" + "foo bar"))) + "") + ("foo bar" + "a")) + "hello"))))
    if (0 >= 42) {
        var v11 uint64 = 978
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = (v3 % ((v3) | 1))
            var v14 int8 = (v10 << ((1) & 63))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    switch 746 {
    case 9223372036854775807:
        v15 := GenS1{F0: 135, F1: (((("foo bar" + "") + "\n\t") + "") + "foo bar"), F2: false}
        _ = v15
    case -1:
        switch "a" {
        case "hello":
            switch ("a" + "\n\t") {
            case "a":
                v16 := GenS1{F0: 32767, F1: "foo bar", F2: false}
                _ = v16
            case "":
                v10 = (^(v10))
            }
        default:
            v2 = -2147483648
        }
    }
    var v17 int16 = v4
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint32 = 16
        _ = fi18
        _ = v19
    }
    switch "hello" {
    case "hello":
        if (false || (-1 >= (v17 / ((32767) | 1)))) {
            v20 := GenS1{F0: 255, F1: ("foo bar" + "hello"), F2: false}
            _ = v20
        }
    case "a":
        v4 += (v17 << ((2) & 63))
    case "":
        delete(v1, "a")
    }
    var v21 float32 = 0.1
    var v22 int64 = -9223372036854775808
    var v23 int8 = (v10 - v10)
    if (!true) {
        v21 += 1.5
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 += genFn1()
            var v25 string = "hello"
            v3 += 0
            _ = fi24
            _ = v25
        }
    } else {
        v23 += (-(v10))
        var v26 int8 = v10
        _ = v26
    }
    v27 := []int{}
    v10 = (v10 + v10)
    v28 := GenS2{F0: v10, F1: v10, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
