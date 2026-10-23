package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 uint16 = v2
    var v4 bool = false
    v4 = (!false)
    if v4 {
        fmt.Println(v2)
    } else {
        fmt.Println(v3, v4, v4)
        v5 := []int{1, 0}
        _ = v5
    }
    v6 := v1[0]
    var v7 int16 = 32767
    switch 2147483647 {
    case -2147483648:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
    }
    delete(v1, 4294967295)
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{fi10, int(v3), fi10, int(v2)}
            if v6 {
                if (v7 > v7) {
                    fmt.Println(v2, v3)
                    v1[2147483648] = ("foo bar" > "foo bar")
                    fmt.Println(fi11)
                }
                var v13 float64 = 0.1
                _ = v13
            } else {
                if (1.5 == -2.25) {
                    v14 := GenS1{F0: (fi10 - (fi11 ^ fi10)), F1: "hello"}
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
        v15 := v1[702]
        _ = fi10
        _ = v15
    }
    var v16 uint = 1
    v17 := GenS1{F0: 2147483647, F1: ("foo bar" + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
