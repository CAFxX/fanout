package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 uint64, p1 bool, p2 int16) uint16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        switch (fi4 - fi4) {
        case 946:
            for fi5 := 0; fi5 < 3; fi5++ {
                v6 := []int{}
                if p1 {
                    fmt.Println(fi4, p0)
                    fmt.Println(fi4, p2)
                }
                _ = fi5
                _ = v6
            }
        default:
            p0 += p0
        }
        v7 := []int{fi4}
        v8 := GenS1{F0: (p0 + (p0 | p0)), F1: 65535, F2: -25}
        _ = fi4
        _ = v7
        _ = v8
    }
    v9 := []int{9223372036854775807, 32768, 42, 901}
    for fi10 := 0; fi10 < 10; fi10++ {
        p1 = true
        var v11 string = ("foo bar" + ("a" + "\n\t"))
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (^(p0)), F1: 560, F2: 1}
    switch "hello" {
    case "hello":
        v13 := GenS1{F0: 649, F1: 609, F2: 1}
        _ = v13
    case "a":
        switch int(p2) {
        case 2147483647:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int8 = -123
                _ = fi14
                _ = v15
            }
        default:
            v16 := GenS1{F0: 811, F1: 2, F2: 589}
            _ = v16
        }
    default:
        var v17 float64 = (0.1 - -2.25)
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 65535
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = 0
    v3 := []uint32{2147483648}
    var v18 bool = ((0.0 - -2.25) < 1e10)
    if len(v3) > 0 {
        v3[0] = 449
    }
    fmt.Println(v18, v18, v2)
    v18 = false
    var v19 int8 = 96
    fmt.Println(v2)
    fmt.Println(v1, v19)
    v18 = v18
    if (0.1 == 0.0) {
        v2 = (v1 % (((^(v2))) | 1))
        var v20 int8 = 1
        v21 := GenS1{F0: 42, F1: 691, F2: -2147483648}
        _ = v20
        _ = v21
    } else {
        var v22 int64 = 1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v22
    }
    _ = len(v3)
    var v23 int64 = -740
    v2 = 18446744073709551615
    switch ((((("\n\t" + "foo bar") + "foo bar") + ("a" + "hello")) + "foo bar") + ("\n\t" + "hello")) {
    case "\n\t":
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS1{F0: v2, F1: 215, F2: -510}
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 int32 = -88
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    case "foo bar":
        var v28 string = "foo bar"
        _ = v28
    case "":
        var v29 uint = 42
        _ = v29
    }
    var v30 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v30
}
