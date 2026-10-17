package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 int64) uint32 {
    var v4 float32 = (3.14159 - (0.1 * 1e10))
    var v5 int8 = -128
    fmt.Println(v4, p0)
    fmt.Println(v4, v5, p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 65535
}
func genFn1() uint64 {
    var v6 bool = (!(true || (((true && (482 > genFn0(128, 720))) && (-2.25 <= 0.1)) && (0 == 255))))
    v6 = v6
    _ = v6
    return 624
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 float32 = 1.5
    fmt.Println(v3)
    if (v3 == v3) {
        v3 = v3
    }
    switch ("" + ("foo bar" + "\n\t")) {
    case "a":
        v7 := []int{-1}
        _ = v7
    case "":
        delete(v2, 1)
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{2147483647, 32768, fi8}
            v3 -= v3
            var v10 string = "hello"
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 1}
        delete(v1, 2147483647)
        fmt.Println(fi11, fi11)
        _ = fi11
        _ = v12
    }
    v3 = (((3.14159 - v3) + v3) - (0.0 + ((v3 + v3) / 1.0)))
    fmt.Println(v3)
    switch -9223372036854775808 {
    case 4294967295:
        var v13 uint16 = 2
        _ = v13
    case 9223372036854775807:
        switch ("" + (("a" + "hello") + "\n\t")) {
        case "foo bar":
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS1{F0: 65535}
                v16 := []int{(fi14 << ((476) & 63)), fi14}
                _ = fi14
                _ = v15
                _ = v16
            }
        case "":
            fmt.Println(v3)
        case "hello":
            fmt.Println(v3)
        }
    case 2147483647:
        v3 = 0.1
    }
    v17 := v1[genFn0(32767, 137)]
    var v18 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
