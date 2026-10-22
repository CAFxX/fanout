package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 128
    v3 := []string{"\n\t", v1, "hello"}
    var v4 int64 = 9223372036854775807
    v5 := []int{346, -2147483648}
    var v6 uint16 = 42
    _ = len(v5)
    var v7 uint64 = 458
    switch 0 {
    case 0:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    default:
        if true {
            v8 := GenS1{F0: v2, F1: (false || (-109 <= -99))}
            v9 := GenS1{F0: v2, F1: (!(1 > 32767))}
            _ = v8
            _ = v9
        }
    }
    if (true || (!(false || (42 == 0)))) {
        var v10 int8 = -1
        var v11 uint64 = (v7 | (v7 * (v7 << ((696) & 63))))
        for fi12 := 0; fi12 < 8; fi12++ {
            if len(v5) > 0 {
                v5[0] = (fi12 % ((fi12) | 1))
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        _ = len(v5)
        _ = fi13
    }
    var v14 float64 = (3.14159 / 1e10)
    fmt.Println(v2)
    v7 += 4294967295
    v2 -= v2
    var v15 uint8 = 128
    v15 = v2
    fmt.Println(v7, v7, v6)
    switch ("\n\t" + ("foo bar" + (((v1 + (v1 + v1)) + "foo bar") + "foo bar"))) {
    case "":
        if false {
            switch ((v1 + v1) + v1) {
            case "":
                for fi16 := 0; fi16 < 11; fi16++ {
                    _ = len(v5)
                    _ = fi16
                }
            }
            _ = len(v3)
        }
    }
    if (true || ((v4 ^ v4) == -1)) {
        v7 += 4294967295
    }
    if (false || (!(true && true))) {
        _ = len(v3)
    } else {
        if (176 > 0) {
            var v17 uint = 4294967295
            _ = v17
        }
        v6 += (v6 + (v6 - v6))
    }
    v18 := GenS1{F0: 2, F1: true}
    v5 = append(v5, -917)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
