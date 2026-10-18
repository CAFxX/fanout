package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 string, p1 uint16) bool {
    v5 := GenS1{F0: p1}
    var v6 int8 = 0
    var v7 int16 = 42
    switch (p0 + p0) {
    case "hello":
        var v8 uint32 = 65535
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 143
    var v3 float64 = (0.1 * -2.25)
    var v4 int = 9223372036854775807
    v9 := []int{}
    v3 -= -2.25
    var v10 string = "hello"
    fmt.Println(v2, v2)
    _ = len(v9)
    v2 = v2
    _ = len(v9)
    v2 = v2
    var v11 string = "hello"
    v12 := GenS2{F0: v10, F1: 665, F2: v2}
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (((v10 + "") + v11) + v10) {
        case "hello":
            for fi14 := 0; fi14 < 9; fi14++ {
                if len(v9) > 0 {
                    v9[0] = (fi14 >> ((v2) & 63))
                }
                if true {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v3 = v3
                }
                _ = fi14
            }
        }
        _ = fi13
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
