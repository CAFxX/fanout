package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }

func main() {
    var v1 int8 = 38
    var v2 uint8 = 2
    var v3 int8 = (v1 >> ((255) & 63))
    var v4 uint8 = (v2 >> ((42) & 63))
    v5 := GenS1{F0: 1}
    if false {
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(v3, v3)
            _ = fi6
        }
    }
    v4 += (v2 % ((120) | 1))
    if (!(false || true)) {
        v7 := GenS2{F0: ((v1 * v3) > 127), F1: (!(-32768 != 1000)), F2: -977}
        if (!(true && true)) {
            v8 := []int{0, 32768}
            v9 := GenS3{F0: 42, F1: -2147483648, F2: 42}
            switch ("hello" + "foo bar") {
            case "":
                v1 += (v3 / ((v3) | 1))
            }
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: false, F1: (127 < 4294967295), F2: -9223372036854775808}
            _ = v10
        }
        _ = v7
    } else {
        v11 := GenS2{F0: false, F1: (255 <= 860), F2: -863}
        _ = v11
    }
    v2 += v2
    v12 := []int{405, int(v1), 2147483647}
    v1 = 0
    if len(v12) > 0 {
        v12[0] = 4294967295
    }
    switch 894 {
    case -317:
        v4 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
