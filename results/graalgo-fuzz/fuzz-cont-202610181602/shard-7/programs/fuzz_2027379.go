package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 uint64) int8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        switch "foo bar" {
        case "foo bar":
            v5 := []int{fi4, fi4}
            _ = v5
        case "":
            for fi6 := 0; fi6 < 18; fi6++ {
                p0 = p0
                p0 = (p0 / ((p0) | 1))
                _ = fi6
            }
        }
        if ("" >= (("hello" + "\n\t") + (("foo bar" + "") + (("foo bar" + ("a" + ("foo bar" + ""))) + ("\n\t" + "foo bar"))))) {
            v7 := GenS1{F0: 32767, F1: (fi4 ^ (^(fi4))), F2: ("a" + "a")}
            var v8 string = "foo bar"
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        p0 = (p0 << ((0) & 63))
        v10 := []int{(fi9 << ((42) & 63))}
        _ = fi9
        _ = v10
    }
    _ = p0
    return 100
}
func genFn1() uint32 {
    v11 := GenS1{F0: -60, F1: 32767, F2: "\n\t"}
    var v12 uint16 = 255
    _ = v11
    _ = v12
    return 578
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483648
    var v3 int32 = 2147483647
    var v13 int64 = 1
    delete(v1, ("hello" + "\n\t"))
    switch ("" + "") {
    case "a":
        if true {
            v1["\n\t"] = 32767
        }
    case "\n\t":
        v13 -= 1
    case "":
        fmt.Println(v13, v3, v3)
    }
    var v14 float64 = 0.1
    var v15 int32 = 1
    v15 = 0
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 += (v3 << ((1) & 63))
        v17 := GenS2{F0: (fi16 ^ 0), F1: 895, F2: (v2 | (v2 >> ((18446744073709551615) & 63)))}
        _ = fi16
        _ = v17
    }
    var v18 int64 = v13
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19, v18)
        _ = fi19
    }
    var v20 uint16 = 293
    var v21 float64 = 3.14159
    var v22 int64 = 0
    var v23 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
