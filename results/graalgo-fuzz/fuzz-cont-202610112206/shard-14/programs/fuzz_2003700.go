package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
func genFn0() uint8 {
    v5 := GenS1{F0: false, F1: "foo bar"}
    v6 := []int{-2147483648, 102, -261, 42}
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint = 255
    var v2 int16 = -32768
    v3 := []uint32{87}
    var v4 int32 = 772
    var v7 int64 = -1
    if ((v1 - (^(v1))) > v1) {
        v2 = v2
        if len(v3) > 0 {
            v3[0] = 648
        }
    }
    var v8 float64 = (0.0 - 1.5)
    v7 += (v7 | -9223372036854775808)
    var v9 bool = true
    if ("hello" == "foo bar") {
        if ((v7 | -541) < (v7 - v7)) {
            v2 = 32767
        }
    } else {
        fmt.Println(v8, v1, v9)
    }
    var v10 int16 = 30
    v11 := GenS1{F0: (-479 > 1), F1: "hello"}
    for fi12 := 0; fi12 < 20; fi12++ {
        if v9 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if false {
            for fi13 := 0; fi13 < 4; fi13++ {
                v2 += v10
                v8 = v8
                _ = fi13
            }
            v14 := GenS2{F0: (v2 ^ (-(v10)))}
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS2{F0: (v2 * 0)}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            if v9 {
                switch "foo bar" {
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        v18 := GenS1{F0: ((v7 > 42) || (630 < 333)), F1: ("hello" + ("hello" + "hello"))}
                        _ = v18
                    case "\n\t":
                        v19 := GenS3{F0: (!v9)}
                        _ = v19
                    }
                case "a":
                    var v20 uint8 = 42
                    _ = v20
                case "foo bar":
                    v9 = (746 >= 971)
                }
                fmt.Println(v4, fi12, v9)
                v21 := GenS2{F0: v2}
                _ = v21
            }
            _ = fi17
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
