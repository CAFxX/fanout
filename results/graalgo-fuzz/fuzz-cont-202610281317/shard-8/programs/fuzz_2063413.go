package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{248, -464, 2147483647, -134}
    var v3 int8 = -73
    v4 := make(map[uint32]uint32)
    v3 = 1
    v4[0] = 2147483648
    v3 += -118
    v3 += (v3 * -57)
    v3 += v3
    var v5 uint16 = 814
    switch 33 {
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 127:
        switch (((("a" + ("\n\t" + ("hello" + "a"))) + "hello") + (("\n\t" + ((("hello" + ("\n\t" + "hello")) + ("" + ("hello" + "hello"))) + ("foo bar" + ""))) + "")) + "foo bar") {
        case "a":
            switch "hello" {
            case "a":
                v5 = v5
            }
        case "foo bar":
            for fi6 := 0; fi6 < 16; fi6++ {
                v4[2147483648] = 473
                v3 = -128
                _ = fi6
            }
        }
    case 0:
        switch (("a" + "a") + ("" + ("foo bar" + ""))) {
        case "a":
            v3 += v3
        case "foo bar":
            if true {
                fmt.Println(v3)
                v7 := GenS1{F0: 9223372036854775807}
                v8 := GenS1{F0: 0}
                _ = v7
                _ = v8
            } else {
                var v9 float64 = 1e10
                v10 := v1[496]
                v11 := GenS1{F0: -337}
                _ = v10
                _ = v11
                _ = v9
            }
        default:
            fmt.Println(v3, v5)
        }
    }
    v5 -= 1
    if false {
        for fi12 := 0; fi12 < 20; fi12++ {
            v3 -= -128
            _ = fi12
        }
        v2 = append(v2, 1)
    }
    var v13 int64 = 0
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
