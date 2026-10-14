package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16) string {
    var v6 int16 = (p0 | p0)
    v6 -= v6
    _ = p0
    _ = v6
    return "a"
}
func genFn1(p0 int, p1 int32, p2 int16) int {
    v7 := []int{}
    fmt.Println(p0, p1)
    for fi8 := 0; fi8 < 1; fi8++ {
        p2 -= (p2 / (((p2 << ((0) & 63))) | 1))
        _ = len(v7)
        p1 = (p1 << ((2) & 63))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 + p0)
}

func main() {
    var v1 string = "\n\t"
    var v2 string = ("foo bar" + v1)
    v3 := make(map[string]bool)
    var v4 uint = 2
    var v5 float32 = (0.0 * float32(((1e10 + 1e10) + 1e10)))
    _ = len(v3)
    if true {
        switch "a" {
        case "a":
            var v9 uint16 = 255
            _ = v9
        default:
            v4 = 255
        }
        var v10 uint8 = 170
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v1)
        _ = fi11
    }
    v5 -= v5
    switch (v1 + ("a" + v1)) {
    case "hello":
        if (723 >= 18446744073709551615) {
            v12 := GenS1{F0: 42}
            if false {
                var v13 float64 = (0.1 - ((3.14159 / 0.1) + 3.14159))
                var v14 uint64 = 4294967295
                _ = v13
                _ = v14
            }
            var v15 uint32 = 0
            _ = v12
            _ = v15
        } else {
            for fi16 := 0; fi16 < 14; fi16++ {
                if true {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v5 = v5
                        _ = fi17
                    }
                    v1 = v2
                }
                var v18 float32 = ((v5 + v5) - 1e10)
                _ = fi16
                _ = v18
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 int64 = 9223372036854775807
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 float32 = -2.25
                    v22 = v5
                    delete(v3, v1)
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            v23 := []int{}
            _ = v23
        }
    case "":
        var v24 uint32 = 4294967295
        _ = v24
    case "a":
        v1 = (v1 + v2)
    }
    switch -2147483648 {
    case -9223372036854775808:
        v2 = v2
    case -1:
        var v25 float64 = (-2.25 + 1.5)
        _ = v25
    case 1:
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                switch genFn1((fi27 % ((101) | 1)), 2147483647, -751) {
                case -1:
                    v2 = "foo bar"
                case 1:
                    var v28 string = "a"
                    _ = v28
                case -2147483648:
                    v29 := []int{}
                    _ = v29
                default:
                    if (true && true) {
                        v30 := GenS1{F0: -1}
                        _ = v30
                    }
                }
                _ = fi27
            }
            v31 := []int{}
            _ = fi26
            _ = v31
        }
    }
    _ = len(v3)
    var v32 bool = (1 == 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
