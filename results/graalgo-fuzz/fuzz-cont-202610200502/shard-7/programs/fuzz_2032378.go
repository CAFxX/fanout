package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8) uint64 {
    v5 := []int{}
    switch ("hello" + "foo bar") {
    case "\n\t":
        _ = len(v5)
    case "a":
        p0 = p0
    }
    p0 -= 255
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 float32 = (((0.0 + ((1e10 / (1e10 * (-2.25 * 1.5))) - 0.0)) - 0.0) / (0.0 - float32((3.14159 * (3.14159 / 1e10)))))
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 uint = 1
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v1 = v1
            v2["\n\t"] = fi7
            switch 217 {
            case 32767:
                var v8 string = "a"
                _ = v8
            case 4294967295:
                switch (("hello" + "") + (("\n\t" + "\n\t") + "\n\t")) {
                case "a":
                    var v9 int8 = -128
                    _ = v9
                }
            case -2147483648:
                if (v1 > v1) {
                    var v10 float64 = float64(v1)
                    var v11 int16 = -1
                    _ = v10
                    _ = v11
                }
            }
            _ = fi7
        }
        _ = fi6
    }
    var v12 int8 = 0
    v1 = (((v1 / v1) * (1.5 - (v1 + float32(0.1)))) + float32(-2.25))
    for fi13 := 0; fi13 < 11; fi13++ {
        v12 += (v12 ^ v12)
        var v14 float64 = float64(v1)
        v1 = (v1 - ((v1 + v1) / (v1 * 1e10)))
        _ = fi13
        _ = v14
    }
    fmt.Println(v4, v4, v1)
    var v15 uint64 = uint64(v4)
    v16 := GenS1{F0: 222, F1: v12}
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
