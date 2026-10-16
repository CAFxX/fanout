package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 uint16, p1 int64) int32 {
    var v5 int64 = (^(p1))
    var v6 uint = 4294967295
    v7 := GenS1{F0: 9223372036854775807}
    var v8 int16 = 1
    var v9 uint8 = 255
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint = 622
    v2 := make(map[uint32]string)
    var v3 bool = (128 <= 111)
    var v4 uint32 = 902
    var v10 int16 = 116
    switch ((("foo bar" + "\n\t") + "\n\t") + ("a" + "\n\t")) {
    case "foo bar":
        switch ("hello" + ((("a" + "") + (("\n\t" + ("" + "\n\t")) + ("a" + ""))) + ("" + "a"))) {
        case "hello":
            v11 := GenS2{F0: -1, F1: -2, F2: ""}
            _ = v11
        }
    }
    v12 := GenS2{F0: 21, F1: -2147483648, F2: "foo bar"}
    if (uint8(v1) <= 255) {
        delete(v2, (v4 - v4))
        v3 = ((96 > 45) && (!(true || v3)))
    }
    if v3 {
        v1 -= 2
    } else {
        v13 := GenS1{F0: -9223372036854775808}
        v3 = ((!v3) || v3)
        fmt.Println(v10, v10, v10)
        _ = v13
    }
    var v14 int16 = (v10 ^ v10)
    var v15 uint8 = 231
    var v16 int64 = 405
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
