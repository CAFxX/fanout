package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0() uint32 {
    var v7 uint8 = 42
    if (0 < 13) {
        var v8 string = "hello"
        _ = v8
    } else {
        switch -2147483648 {
        case 2147483647:
            v7 -= v7
        default:
            if (!false) {
                v9 := GenS1{F0: (!false), F1: -1}
                switch 32768 {
                case -1:
                    v7 = v7
                case 128:
                    v10 := []int{1, 4294967295, -996}
                    _ = v10
                }
                _ = v9
            } else {
                var v11 uint = 262
                _ = v11
            }
        }
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        switch (("hello" + ("foo bar" + ("\n\t" + "a"))) + "\n\t") {
        case "\n\t":
            switch fi12 {
            case 270:
                fmt.Println(v7, fi12)
            case -60:
                switch "a" {
                case "hello":
                    fmt.Println(fi12, fi12)
                }
            }
        }
        _ = fi12
    }
    _ = v7
    return 1
}

func main() {
    v1 := []int{127, 4294967295}
    var v2 uint = 523
    var v3 int32 = 73
    var v4 uint = 255
    var v5 uint16 = 2
    var v6 string = "a"
    var v13 int16 = 42
    var v14 int32 = (-(v3))
    var v15 int64 = -79
    v16 := GenS1{F0: ((v13 % ((int16(v15)) | 1)) == v13), F1: v14}
    v14 = v3
    var v17 string = (v6 + "hello")
    var v18 float32 = ((0.0 + ((-2.25 * (-2.25 + 1.5)) * 1e10)) * 1e10)
    v19 := GenS1{F0: false, F1: v3}
    _ = len(v1)
    if (true || (true || (248 == v2))) {
        v1 = append(v1, -1)
    }
    if true {
        v20 := GenS1{F0: false, F1: int32(v2)}
        v4 += (v4 * v2)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS1{F0: (true && false), F1: -89}
            switch fi21 {
            case 563:
                v4 = v4
            case 128:
                if true {
                    if (fi21 <= fi21) {
                        var v23 float64 = 0.1
                        v18 = v18
                        _ = v23
                    }
                    v6 = ("foo bar" + v6)
                    fmt.Println(v13, v15)
                }
            }
            _ = fi21
            _ = v22
        }
        var v24 uint = v4
        v4 = (v24 % (((^(v4))) | 1))
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
