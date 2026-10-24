package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 int32, p1 uint32) string {
    v7 := GenS1{F0: true}
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            v9 := GenS1{F0: (!true)}
            _ = v9
        }
        _ = fi8
    }
    switch (((("" + ("\n\t" + "foo bar")) + "") + "foo bar") + "foo bar") {
    case "foo bar":
        v10 := GenS1{F0: (-1 < -128)}
        _ = v10
    case "\n\t":
        v11 := []int{int(p1)}
        _ = v11
    case "":
        v12 := GenS2{F0: 958, F1: 0}
        _ = v12
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                p1 = (p1 & p1)
                if true {
                    p1 = (p1 - 0)
                } else {
                    v15 := GenS3{F0: 32767}
                    _ = v15
                }
                var v16 string = "foo bar"
                _ = fi14
                _ = v16
            }
            _ = fi13
        }
    }
    var v17 int = 32767
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v17
    _ = v7
    return ("\n\t" + "a")
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    var v6 uint32 = 2
    v18 := GenS4{F0: -151, F1: -1}
    v6 = (^(v6))
    var v19 float32 = 0.0
    v19 = v2
    for fi20 := 0; fi20 < 9; fi20++ {
        switch fi20 {
        case 32768:
            var v21 int32 = 2147483647
            _ = v21
        }
        _ = fi20
    }
    var v22 uint = v1
    delete(v3, ("a" + "\n\t"))
    var v23 int16 = -495
    var v24 uint64 = 473
    v1 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
