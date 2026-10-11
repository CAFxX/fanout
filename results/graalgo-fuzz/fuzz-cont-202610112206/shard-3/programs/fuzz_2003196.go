package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int = 42
    v2 := make(map[string]int)
    var v3 int = v1
    var v4 uint64 = 410
    var v5 float64 = 1e10
    delete(v2, ("a" + (("" + "") + "foo bar")))
    switch ("\n\t" + "a") {
    case "a":
        var v6 string = ("a" + "\n\t")
        _ = v6
    default:
        var v7 uint16 = 81
        _ = v7
    }
    v8 := GenS1{F0: 2}
    if true {
        v9 := []int{(v3 % (((^(v3))) | 1))}
        v10 := GenS1{F0: 2}
        _ = v10
        _ = v9
    } else {
        if (true || (true && true)) {
            v11 := GenS1{F0: 596}
            _ = v11
        } else {
            delete(v2, ("hello" + "hello"))
            var v12 string = "foo bar"
            _ = v12
        }
        fmt.Println(v5, v5)
        v2["hello"] = (^(v3))
    }
    fmt.Println(v4)
    if true {
        switch (("" + ("" + "a")) + "foo bar") {
        case "a":
            v13 := []int{v1, (v1 | 4294967295), v3}
            _ = v13
        case "hello":
            if true {
                v14 := GenS2{F0: 1, F1: 65535}
                if true {
                    v15 := GenS3{F0: -9223372036854775808, F1: 1, F2: 526}
                    var v16 int8 = -1
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
        case "":
            if false {
                var v17 float64 = v5
                for fi18 := 0; fi18 < 20; fi18++ {
                    fmt.Println(v3)
                    _ = fi18
                }
                _ = v17
            } else {
                v19 := GenS2{F0: 597, F1: 65535}
                var v20 uint32 = 2
                v3 -= v1
                _ = v19
                _ = v20
            }
        default:
            fmt.Println(v3)
        }
        _ = len(v2)
        delete(v2, "")
    } else {
        v1 -= (v1 & v3)
        v5 += 1e10
        v1 = (v1 + -900)
    }
    var v21 uint32 = 1
    v22 := GenS4{F0: 0, F1: 127}
    if (0 >= -1) {
        var v23 uint8 = 127
        _ = v23
    }
    fmt.Println(v21, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
