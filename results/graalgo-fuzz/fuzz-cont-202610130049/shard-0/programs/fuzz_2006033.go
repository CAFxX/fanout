package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string }

func main() {
    var v1 int64 = -1
    var v2 int16 = -357
    var v3 uint8 = 128
    if (v3 == 26) {
        v4 := []int{int(v1), 708, 127}
        for fi5 := 0; fi5 < 3; fi5++ {
            _ = len(v4)
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            _ = fi5
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
        _ = v4
    }
    var v6 uint32 = 4294967295
    v7 := GenS1{F0: uint8(v6), F1: 4294967295}
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{}
        v10 := GenS2{F0: (((("\n\t" + "a") + "hello") + ("a" + "hello")) + "foo bar")}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: (v3 >> ((1) & 63)), F1: 18446744073709551615}
            switch "foo bar" {
            case "\n\t":
                var v14 uint64 = 28
                _ = v14
            case "a":
                if true {
                    v3 += 13
                }
            case "":
                switch "\n\t" {
                case "foo bar":
                    var v15 string = ("hello" + ("" + "foo bar"))
                    _ = v15
                case "a":
                    v6 = v6
                }
            }
            v6 = (v6 | v6)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v16 := GenS2{F0: "a"}
    if false {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v3, v3)
            var v18 uint8 = (^(v3))
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint = 400
            var v21 bool = (0.1 > -2.25)
            fmt.Println(v1, v6)
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 uint16 = 255
    if ((0.0 + -2.25) < (3.14159 - 1e10)) {
        fmt.Println(v22)
        v23 := []int{}
        _ = v23
    }
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v6
    _ = v7
}
