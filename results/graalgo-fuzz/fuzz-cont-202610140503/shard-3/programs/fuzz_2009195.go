package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0() string {
    for fi5 := 0; fi5 < 2; fi5++ {
        if (0 <= 857) {
            v6 := []int{fi5}
            v6 = append(v6, (fi5 - fi5))
            if false {
                v7 := GenS1{F0: true, F1: 128}
                _ = v7
            }
            _ = v6
        }
        var v8 uint32 = 0
        _ = fi5
        _ = v8
    }
    var v9 int8 = 69
    fmt.Println(v9)
    _ = v9
    return "a"
}

func main() {
    var v1 int16 = -1
    var v2 int = 1
    var v3 uint16 = 0
    var v4 uint64 = 148
    var v10 uint8 = 127
    fmt.Println(v2)
    v2 = int(v10)
    var v11 uint16 = (v3 * 43)
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            v10 += v10
            var v13 uint16 = v3
            _ = v13
        } else {
            var v14 bool = (403 < 296)
            if true {
                v2 = fi12
            }
            fmt.Println(v11)
            _ = v14
        }
        _ = fi12
    }
    var v15 bool = true
    if v15 {
        var v16 string = ("hello" + "hello")
        v15 = v15
        _ = v16
    } else {
        switch "a" {
        case "\n\t":
            switch ((("\n\t" + (("a" + "a") + "hello")) + ("\n\t" + (("hello" + "") + (("a" + ("" + (("foo bar" + "a") + ("\n\t" + "a")))) + "hello")))) + "hello") {
            case "a":
                fmt.Println(v1, v15, v2)
            case "hello":
                v17 := []int{865, (v2 & (v2 | (v2 % (((v2 ^ (v2 << ((uint(v11)) & 63)))) | 1)))), v2, (v2 << ((42) & 63))}
                _ = v17
            }
        case "":
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int32 = -1
                _ = fi18
                _ = v19
            }
        default:
            v20 := []int{v2}
            _ = v20
        }
        var v21 string = "foo bar"
        _ = v21
    }
    var v22 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
