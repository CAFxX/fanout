package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
func genFn0() uint64 {
    if (1 != 216) {
        v7 := []int{128, 4294967295, 9223372036854775807, 32767}
        var v8 uint8 = 255
        if false {
            if true {
                var v9 uint = 18446744073709551615
                _ = v9
            } else {
                v10 := GenS1{F0: 4294967295, F1: -57, F2: "\n\t"}
                var v11 uint16 = 0
                _ = v10
                _ = v11
            }
        }
        _ = v7
        _ = v8
    }
    var v12 uint64 = 687
    var v13 int64 = -767
    _ = v12
    _ = v13
    return v12
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    var v3 int16 = -609
    var v4 float32 = 0.1
    var v5 string = ("foo bar" + ("hello" + ""))
    var v6 string = v5
    var v14 int32 = 42
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(fi15, v4, v6)
        _ = fi15
    }
    if (true && true) {
        switch 32767 {
        case 32768:
            switch ("" + "hello") {
            case "hello":
                v16 := []int{}
                _ = v16
            }
        case 735:
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 string = ("\n\t" + v6)
                _ = fi17
                _ = v18
            }
        }
    }
    v19 := GenS1{F0: 2, F1: 62, F2: v6}
    v20 := GenS1{F0: 328, F1: 6, F2: ("a" + (v5 + v6))}
    if (255 > 127) {
        if (-735 > -1) {
            v6 = ("" + v6)
        }
        v6 = v6
    }
    if true {
        v4 -= v4
        _ = len(v2)
    }
    var v21 bool = (127 >= 244)
    v22 := GenS1{F0: 255, F1: 1, F2: v5}
    v23 := GenS1{F0: 0, F1: -1, F2: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
