package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 bool = (42 == 100)
    var v2 int8 = -1
    var v3 bool = false
    var v4 float64 = 1.5
    v5 := make(map[string]bool)
    var v6 int8 = (-(v2))
    var v7 int64 = 42
    switch 2147483647 {
    case -506:
        fmt.Println(v7, v1)
    case -360:
        if false {
            if (384 >= 1) {
                switch -1 {
                case 128:
                    v4 = 3.14159
                case 506:
                    v7 -= v7
                case -737:
                    v8 := []int{-2147483648, -251, 559, 818}
                    _ = v8
                default:
                    v4 = v4
                }
                v4 -= (((v4 * 1.5) + v4) - 0.0)
            } else {
                if (-2147483648 != 481) {
                    v5[("foo bar" + "hello")] = v1
                    if v1 {
                        v9 := GenS1{F0: 2, F1: 854, F2: 351}
                        v6 -= (v2 * 1)
                        v10 := []int{}
                        _ = v10
                        _ = v9
                    }
                }
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 uint = 255
                    _ = fi11
                    _ = v12
                }
            }
            var v13 uint = 2
            _ = v13
        } else {
            switch "" {
            case "hello":
                var v14 int = 32767
                _ = v14
            }
        }
    default:
        if v1 {
            var v15 string = "a"
            v16 := []int{}
            v7 = -1
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: "\n\t", F1: 785}
    v2 -= (v6 * (^(v6)))
    v18 := GenS3{F0: (!(v3 || v1)), F1: 716, F2: v7}
    if (!(127 >= 42)) {
        v2 -= v6
    } else {
        var v19 uint32 = 2147483648
        switch ((("foo bar" + "foo bar") + ("\n\t" + "hello")) + "") {
        case "":
            v6 = (v6 - 42)
        }
        _ = v19
    }
    v20 := GenS4{F0: 2, F1: 1, F2: 913}
    v21 := []int{9223372036854775807}
    v22 := GenS4{F0: 1, F1: 0, F2: 65535}
    var v23 uint8 = 2
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 string = ""
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
