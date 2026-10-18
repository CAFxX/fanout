package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 int16 = -32768
    var v3 int32 = -2147483648
    v4 := GenS1{F0: int8(v3)}
    switch 32768 {
    case 128:
        v3 = v3
    case 0:
        fmt.Println(v3)
    }
    fmt.Println(v1)
    v2 += v2
    for fi5 := 0; fi5 < 15; fi5++ {
        if (v1 < v1) {
            var v6 float32 = (3.14159 / -2.25)
            v7 := GenS1{F0: 127}
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 100}
            v3 += v3
            _ = v8
        }
        switch (("a" + (("" + "hello") + ("foo bar" + ""))) + "") {
        case "\n\t":
            switch "foo bar" {
            case "a":
                if false {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        fmt.Println(fi5)
                        _ = fi9
                    }
                    fmt.Println(fi5, fi5)
                    if true {
                        fmt.Println(fi5, v3)
                        v10 := GenS1{F0: 100}
                        fmt.Println(v1, v3)
                        _ = v10
                    }
                } else {
                    v11 := GenS1{F0: 100}
                    v3 += 0
                    _ = v11
                }
            case "\n\t":
                v1 -= (v1 % ((v1) | 1))
            }
        case "a":
            v12 := []int{(^(fi5)), (^(fi5)), int(v3), fi5}
            _ = v12
        }
        v2 = (v2 << ((0) & 63))
        _ = fi5
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := []int{fi14}
                v16 := GenS1{F0: -95}
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int32 = (v3 - v3)
            _ = v17
        }
        _ = fi13
    }
    v18 := []int{832}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
