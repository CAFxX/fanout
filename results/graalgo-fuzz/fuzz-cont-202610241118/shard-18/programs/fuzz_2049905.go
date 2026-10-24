package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0() string {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 10
        v7 -= (^(v7))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        if (!false) {
            v9 := []int{fi8, fi8, (fi8 - -56)}
            if true {
                v10 := GenS1{F0: false, F1: 4294967295}
                _ = v10
            } else {
                switch "\n\t" {
                case "\n\t":
                    v11 := GenS1{F0: (true || true), F1: 42}
                    _ = v11
                case "":
                    var v12 int = (fi8 + -551)
                    _ = v12
                default:
                    if true {
                        var v13 string = ""
                        v13 = ""
                        _ = v13
                    }
                }
                v14 := GenS1{F0: (false && true), F1: 682}
                var v15 uint16 = 255
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint16 = 2
                _ = fi16
                _ = v17
            }
            _ = v9
        }
        var v18 int64 = 339
        v18 = v18
        _ = fi8
        _ = v18
    }
    return (("" + "\n\t") + "hello")
}

func main() {
    v1 := []int{32768, -9223372036854775808, 127}
    var v2 uint64 = 4294967295
    var v3 string = "foo bar"
    v4 := []string{}
    var v5 int64 = 42
    if (false && false) {
        if (0 != 0) {
            var v19 float64 = 1e10
            switch (("foo bar" + (("a" + v3) + (v3 + "hello"))) + "hello") {
            case "hello":
                fmt.Println(v2, v5)
            default:
                var v20 int8 = 100
                _ = v20
            }
            _ = v19
        }
        var v21 float32 = -2.25
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v2, fi22)
        _ = fi22
    }
    var v23 float64 = 1e10
    if (false && (true || ((0.0 >= v23) || ((v5 < v5) && true)))) {
        fmt.Println(v3, v2)
        v3 = v3
    }
    v2 = (v2 | (v2 | v2))
    var v24 uint8 = 0
    var v25 bool = (!true)
    for fi26 := 0; fi26 < 19; fi26++ {
        v5 += (v5 << ((1) & 63))
        v3 = (v3 + (v3 + "foo bar"))
        v27 := GenS1{F0: false, F1: 2147483647}
        _ = fi26
        _ = v27
    }
    v24 += v24
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint = 65535
        v3 = ((((v3 + "foo bar") + v3) + ("\n\t" + (v3 + "foo bar"))) + v3)
        _ = fi28
        _ = v29
    }
    fmt.Println(v3)
    v24 -= v24
    if (!v25) {
        v5 = v5
        v24 = v24
        v5 -= v5
    }
    var v30 string = (v3 + ("a" + (v3 + "foo bar")))
    if (!v25) {
        v3 = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
