package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
func genFn0() bool {
    var v7 int64 = 42
    if false {
        var v8 float64 = 3.14159
        _ = v8
    }
    _ = v7
    return (false && (1 == 796))
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = 0
    var v3 int = -9223372036854775808
    var v4 uint = 495
    var v5 uint64 = 18446744073709551615
    var v6 uint = v4
    if false {
        if true {
            var v9 uint = v4
            if ((116 < 65535) && (true && false)) {
                switch 692 {
                case 4294967295:
                    var v10 int8 = -110
                    _ = v10
                default:
                    var v11 uint8 = 42
                    _ = v11
                }
            } else {
                var v12 string = ""
                var v13 uint32 = 107
                switch ((v12 + v12) + "a") {
                case "a":
                    var v14 bool = false
                    _ = v14
                case "":
                    fmt.Println(v1)
                }
                _ = v12
                _ = v13
            }
            _ = v9
        }
        var v15 bool = true
        switch "" {
        case "\n\t":
            var v16 string = ("a" + (("\n\t" + "") + ""))
            _ = v16
        default:
            v3 = v3
        }
        _ = v15
    }
    switch 32767 {
    case 2147483647:
        v3 -= (v3 & (v3 / (((v3 ^ (v3 ^ v3))) | 1)))
    case 127:
        v3 = 42
    case 4294967295:
        var v17 int = v3
        _ = v17
    default:
        var v18 uint8 = 255
        _ = v18
    }
    if (false && true) {
        if false {
            var v19 int = (v3 << (((v6 << (((v4 + v6)) & 63))) & 63))
            _ = v19
        }
        fmt.Println(v4, v3)
    } else {
        var v20 float64 = (-2.25 * (1.5 * ((1.5 + 3.14159) * -2.25)))
        v1 -= 0.1
        for fi21 := 0; fi21 < 1; fi21++ {
            switch "a" {
            case "\n\t":
                var v22 float32 = (0.0 - ((float32((3.14159 / v20)) / v1) * v1))
                _ = v22
            case "foo bar":
                var v23 string = "foo bar"
                _ = v23
            case "a":
                for fi24 := 0; fi24 < 2; fi24++ {
                    v6 = v4
                    var v25 int16 = 1000
                    v26 := []int{-13, (v3 * fi21)}
                    v3 = (fi24 - (fi21 / ((fi24) | 1)))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                var v27 string = "foo bar"
                _ = v27
            }
            v28 := []int{-218, fi21}
            _ = fi21
            _ = v28
        }
        _ = v20
    }
    if true {
        var v29 int64 = -1
        if (v5 == 4294967295) {
            var v30 uint64 = 1
            var v31 uint = (v4 ^ (v4 * 548))
            _ = v30
            _ = v31
        }
        v32 := []int{4294967295, v3}
        _ = v29
        _ = v32
    } else {
        var v33 uint8 = 2
        v34 := []int{32767, 406}
        _ = v33
        _ = v34
    }
    v4 -= (^(v4))
    var v35 string = ("a" + "hello")
    var v36 float32 = -2.25
    var v37 uint16 = 65535
    v38 := []int{v3}
    var v39 int16 = 42
    v40 := GenS1{F0: v39, F1: -128, F2: uint16(v4)}
    var v41 int32 = 42
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
