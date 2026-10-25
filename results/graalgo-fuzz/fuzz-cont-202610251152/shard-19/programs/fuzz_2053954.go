package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 string = ""
    var v2 string = ((v1 + v1) + (((v1 + (v1 + (v1 + v1))) + (v1 + (v1 + "a"))) + v1))
    var v3 float32 = 0.0
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    var v6 uint8 = 2
    switch -1 {
    case 42:
        v1 = "a"
    }
    v4[-9223372036854775808] = -2147483648
    switch v1 {
    case "":
        v7 := GenS1{F0: 4294967295}
        _ = v7
    case "a":
        if (350 != 0) {
            v8 := []int{-9223372036854775808}
            _ = v8
        } else {
            v9 := v4[9223372036854775807]
            v10 := GenS1{F0: 2}
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float64 = 1e10
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
    default:
        if ((float64((v3 + -2.25)) > (1.5 / 1.5)) && (9223372036854775807 >= 0)) {
            for fi13 := 0; fi13 < 16; fi13++ {
                fmt.Println(v3)
                v1 = ("hello" + "\n\t")
                for fi14 := 0; fi14 < 11; fi14++ {
                    v1 = v1
                    _ = fi14
                }
                _ = fi13
            }
        }
    }
    if true {
        v1 = (v1 + ("a" + (v2 + "")))
    } else {
        switch v1 {
        case "":
            v2 = v2
        case "a":
            if false {
                v15 := GenS2{F0: 128}
                _ = v15
            }
        }
    }
    v1 = "\n\t"
    v16 := GenS2{F0: 18446744073709551615}
    v2 = (v1 + v2)
    v1 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
