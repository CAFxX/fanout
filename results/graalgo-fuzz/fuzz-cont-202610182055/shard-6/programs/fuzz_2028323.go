package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 int16; F1 uint; F2 string }

func main() {
    var v1 uint = 4294967295
    var v2 int64 = -1
    var v3 int = 701
    for fi4 := 0; fi4 < 12; fi4++ {
        if ((0 >= 0) || (("\n\t" + "foo bar") >= "a")) {
            switch "a" {
            case "":
                v3 = fi4
            case "hello":
                fmt.Println(fi4, v2, v3)
            default:
                for fi5 := 0; fi5 < 5; fi5++ {
                    v1 = 65535
                    var v6 int8 = 1
                    _ = fi5
                    _ = v6
                }
            }
            if false {
                v1 -= 42
                v1 = v1
            } else {
                var v7 int32 = 1
                var v8 float32 = 1.5
                v9 := GenS1{F0: ("hello" + "a"), F1: fi4, F2: v1}
                _ = v7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int = -683
                v12 := []int{(fi10 + -522), (v11 % ((127) | 1)), (v3 & 972), (v11 / ((42) | 1))}
                var v13 int8 = 67
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        v1 -= (v1 + v1)
        if ("" <= "foo bar") {
            if true {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v1 = v1
                    _ = fi14
                }
                if true {
                    v1 += v1
                    v15 := GenS2{F0: -1, F1: (v1 << (((v1 / ((v1) | 1))) & 63)), F2: ("a" + "\n\t")}
                    _ = v15
                }
                v16 := []int{}
                _ = v16
            }
            fmt.Println(fi4, v2)
            if (82 == 2) {
                v17 := []int{(-(fi4)), (fi4 | fi4), 32768, -1}
                _ = v17
            }
        }
        _ = fi4
    }
    v18 := []int{v3, (^(v3)), (v3 << ((v1) & 63)), 42}
    v3 = (v3 & -832)
    for fi19 := 0; fi19 < 19; fi19++ {
        if ((v2 / (((v2 / (((-(v2))) | 1))) | 1)) >= 105) {
            var v20 float32 = 3.14159
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 float32 = float32(3.14159)
                v20 -= v20
                v2 = (v2 / ((v2) | 1))
                fmt.Println(fi21)
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        _ = fi19
    }
    if len(v18) > 0 {
        v18[0] = 104
    }
    fmt.Println(v2)
    var v23 int8 = -1
    v24 := GenS1{F0: "\n\t", F1: (v3 << ((v1) & 63)), F2: (v1 + 221)}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
