package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
func genFn0() uint {
    var v4 int = 42
    v5 := GenS1{F0: 22, F1: "\n\t"}
    var v6 int64 = 1
    fmt.Println(v6, v6)
    v7 := []int{42, v4, 42}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 266
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 0
    var v3 uint32 = 0
    v2 -= (^(v2))
    delete(v1, 1)
    v8 := GenS1{F0: -1, F1: ("foo bar" + "foo bar")}
    fmt.Println(v3, v3)
    if (1 != 75) {
        v9 := []int{128}
        _ = v9
    }
    var v10 int8 = 42
    v11 := GenS2{F0: 0}
    for fi12 := 0; fi12 < 4; fi12++ {
        v10 = 127
        if true {
            if ((true && (((-(v10)) == v10) || (true || (!false)))) || (32767 < 1)) {
                var v13 uint16 = 447
                _ = v13
            } else {
                v2 = 1
                var v14 uint8 = 157
                v15 := GenS2{F0: 486}
                _ = v14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 float32 = -2.25
                _ = fi16
                _ = v17
            }
            var v18 int64 = 9223372036854775807
            v19 := GenS2{F0: -1}
            _ = v18
            _ = v19
        }
        switch "foo bar" {
        case "":
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 string = ""
                var v22 float32 = (0.1 / -2.25)
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi12
    }
    var v23 int32 = 14
    v3 = (v3 - v3)
    v10 -= 100
    v1[32767] = ""
    v24 := GenS1{F0: 186, F1: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v8
}
