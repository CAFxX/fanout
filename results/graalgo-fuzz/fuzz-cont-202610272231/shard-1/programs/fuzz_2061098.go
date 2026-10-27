package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
func genFn0() int {
    var v7 uint32 = 540
    fmt.Println(v7)
    var v8 int8 = -8
    if (489 != 841) {
        var v9 uint8 = 42
        v10 := []int{32767, -777, 787}
        switch 55 {
        case 4294967295:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
        _ = v10
        _ = v9
    }
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 171
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -120
    var v3 bool = false
    v4 := make(map[uint32]bool)
    var v5 string = ("a" + "")
    var v6 int16 = 1000
    v5 = v5
    fmt.Println(v5, v2, v5)
    var v11 int8 = -1
    v12 := []int{}
    v13 := GenS1{F0: ("\n\t" + (v5 + v5)), F1: (v2 - (v2 - int(v6))), F2: -9223372036854775808}
    v14 := GenS1{F0: "hello", F1: v2, F2: -9223372036854775808}
    var v15 int16 = v6
    var v16 int8 = v11
    if len(v12) > 0 {
        v12[0] = v2
    }
    switch ("a" + v5) {
    case "\n\t":
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint = 0
            _ = fi17
            _ = v18
        }
    case "":
        if (132 != 0) {
            for fi19 := 0; fi19 < 11; fi19++ {
                v5 = (("\n\t" + v5) + (v5 + ""))
                v5 = ""
                _ = fi19
            }
            v20 := GenS2{F0: (v2 / ((int(v6)) | 1)), F1: false, F2: 32767}
            v21 := GenS3{F0: 2147483647}
            _ = v20
            _ = v21
        }
    case "foo bar":
        fmt.Println(v5)
    default:
        var v22 uint = 0
        _ = v22
    }
    var v23 string = v5
    v24 := GenS4{F0: "\n\t"}
    var v25 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
