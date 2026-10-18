package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint, p1 int, p2 uint8) uint64 {
    switch "\n\t" {
    case "":
        v4 := []int{(p1 >> (((p0 | (p0 / ((p0) | 1)))) & 63)), -2147483648}
        _ = v4
    }
    v5 := []int{4294967295, p1}
    var v6 int64 = -76
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 756
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 671
    var v3 uint16 = 2
    switch ("\n\t" + ("a" + ("\n\t" + "hello"))) {
    case "a":
        if (128 > 674) {
            v3 += v3
            var v7 float64 = 1e10
            _ = v7
        } else {
            switch ("" + "a") {
            case "foo bar":
                v8 := GenS1{F0: 42}
                _ = v8
            case "":
                v9 := []int{}
                _ = v9
            }
            if false {
                var v10 uint64 = 108
                _ = v10
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := GenS1{F0: -1}
                    v2 -= v2
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: 0}
                _ = v13
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint = 255
                v16 := v1[("" + ("a" + ("\n\t" + ((("hello" + ("" + ("foo bar" + (("" + "foo bar") + ("hello" + ""))))) + "hello") + ("" + "\n\t")))))]
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    case "hello":
        var v17 bool = (1e10 <= 3.14159)
        _ = v17
    default:
        fmt.Println(v2, v3)
    }
    var v18 uint16 = v3
    v19 := v1["\n\t"]
    var v20 float64 = (-2.25 * 3.14159)
    var v21 float32 = (0.1 / (1e10 - -2.25))
    v22 := GenS2{F0: 274, F1: -199, F2: 346}
    v23 := GenS2{F0: 255, F1: -2147483648, F2: 65535}
    v18 = uint16(v2)
    v24 := GenS3{F0: 42}
    v3 = (^(v3))
    switch "hello" {
    case "foo bar":
        v2 += (v2 * (v2 >> ((1) & 63)))
    case "a":
        var v25 uint32 = 4294967295
        _ = v25
    default:
        if (!("foo bar" <= "a")) {
            v26 := GenS3{F0: 1}
            _ = v26
        }
    }
    v18 += v3
    v27 := GenS1{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
}
