package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
func genFn0(p0 int32, p1 int64, p2 uint64) uint {
    var v7 uint8 = 42
    v8 := []int{927, -2147483648, -881}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = 2147483647
    var v3 bool = ((false || false) || (1e10 >= (1e10 - ((float64(1.5) + 1e10) - 0.1))))
    v4 := []string{"\n\t"}
    var v5 uint = 1
    v6 := make(map[uint32]string)
    fmt.Println(v5, v2)
    switch (("" + (("foo bar" + ((("a" + "a") + (("\n\t" + ("\n\t" + "\n\t")) + "hello")) + "foo bar")) + "\n\t")) + "a") {
    case "\n\t":
        switch 801 {
        case 490:
            v9 := GenS1{F0: 42, F1: "\n\t"}
            _ = v9
        }
    case "":
        var v10 int8 = 100
        _ = v10
    }
    v11 := GenS1{F0: uint8(v1), F1: "foo bar"}
    v1 = v1
    var v12 int64 = 1
    if v3 {
        _ = len(v6)
        if v3 {
            if (true || ((-2.25 <= -2.25) && v3)) {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                var v13 int8 = 1
                _ = v13
            } else {
                var v14 uint16 = 246
                var v15 bool = false
                _ = v14
                _ = v15
            }
        }
    }
    v3 = v3
    v1 = v1
    for fi16 := 0; fi16 < 5; fi16++ {
        v3 = (v3 && v3)
        var v17 uint8 = 128
        _ = fi16
        _ = v17
    }
    var v18 int8 = -96
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v19 int16 = 1000
    switch ("foo bar" + "foo bar") {
    case "a":
        if (!v3) {
            _ = len(v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
