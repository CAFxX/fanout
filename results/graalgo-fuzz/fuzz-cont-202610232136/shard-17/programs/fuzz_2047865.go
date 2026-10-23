package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 int32 = -1
    v4 := make(map[uint32]string)
    switch 32767 {
    case 9223372036854775807:
        v3 -= (^(v3))
    case -147:
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
    default:
        if false {
            fmt.Println(v3)
            var v7 string = ("" + (("hello" + ("foo bar" + ("a" + "hello"))) + "foo bar"))
            _ = v7
        } else {
            var v8 bool = false
            var v9 uint8 = 209
            _ = v8
            _ = v9
        }
    }
    v3 = v3
    v3 = 0
    v3 += v3
    var v10 bool = (4294967295 <= 2)
    var v11 int32 = -2147483648
    fmt.Println(v10, v3, v11)
    var v12 int64 = 42
    if v10 {
        var v13 bool = (!(1 > -128))
        var v14 uint = 42
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: ("" + "\n\t")}
        _ = v15
    }
    var v16 int32 = (v3 - v3)
    switch -332 {
    case 127:
        v16 = v16
    case 9223372036854775807:
        v1["hello"] = v10
    case -1:
        v17 := v1["a"]
        _ = v17
    default:
        v3 = -1
    }
    fmt.Println(v3, v11, v11)
    if (!true) {
        switch "\n\t" {
        case "":
            v18 := GenS1{F0: "a"}
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v11, v11, v11)
            _ = fi19
        }
        v12 = (v12 / ((v12) | 1))
    } else {
        v16 = (^(v16))
        if v10 {
            v20 := GenS2{F0: v16, F1: 18446744073709551615, F2: 642}
            v3 -= (-(v3))
            _ = v20
        } else {
            v21 := []int{32767, 1, -364}
            _ = v21
        }
        v12 -= (v12 >> ((255) & 63))
    }
    v22 := []int{892, -250, 4294967295, 1}
    v11 -= (v3 / ((1) | 1))
    switch 128 {
    case 128:
        v12 = v12
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            for fi24 := 0; fi24 < 19; fi24++ {
                switch (fi24 | fi24) {
                case -2147483648:
                    v10 = v10
                case 32767:
                    v10 = false
                    fmt.Println(v12)
                    v3 = (v11 * (-(v16)))
                case 32768:
                    v25 := GenS1{F0: ("" + ("foo bar" + "a"))}
                    _ = v25
                }
                _ = fi24
            }
            var v26 int8 = 42
            if v10 {
                delete(v2, (fi23 / ((32768) | 1)))
            }
            _ = fi23
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
