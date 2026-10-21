package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) uint8 {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 67
}
func genFn1() int64 {
    var v5 uint16 = 32767
    fmt.Println(v5)
    _ = v5
    return 827
}

func main() {
    v1 := []string{("foo bar" + ("\n\t" + ("hello" + "foo bar"))), ("\n\t" + "foo bar"), "\n\t", "a"}
    var v2 int8 = 1
    var v3 float32 = 0.1
    var v4 uint8 = 71
    var v6 uint8 = (^(v4))
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 = 1.5
        v1 = append(v1, "foo bar")
        _ = fi7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint16 = 27
    if (true || true) {
        v6 = v4
        switch "hello" {
        case "hello":
            v9 := GenS1{F0: 411}
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 11; fi10++ {
                if true {
                    v3 -= ((v3 * (-2.25 + 0.1)) - v3)
                    _ = len(v1)
                    fmt.Println(v3, v4)
                }
                _ = len(v1)
                _ = fi10
            }
        }
    }
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint32 = 0
            _ = fi11
            _ = v12
        }
    }
    v3 -= (float32((0.0 / 0.1)) * ((v3 - v3) / (1e10 * v3)))
    var v13 bool = (42 < 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
