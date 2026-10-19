package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 int8, p1 uint32, p2 uint64) int16 {
    var v7 uint = 2
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int32 = 0
        for fi10 := 0; fi10 < 10; fi10++ {
            switch ("\n\t" + ("a" + ("a" + "foo bar"))) {
            case "":
                v11 := []int{}
                _ = v11
            case "foo bar":
                if false {
                    switch fi8 {
                    case 42:
                        var v12 uint8 = 0
                        _ = v12
                    case 1:
                        v13 := []int{(fi8 << (((^(v7))) & 63)), 2147483647, (fi10 - (fi8 * fi8)), (-(fi10))}
                        _ = v13
                    }
                    var v14 int64 = 9223372036854775807
                    fmt.Println(v7, v9)
                    _ = v14
                }
                v15 := []int{(fi10 / ((fi8) | 1)), (-(fi8)), 54}
                _ = v15
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    p2 += (p2 - p2)
    switch -609 {
    case 5:
        switch 32768 {
        case 128:
            v7 += v7
        default:
            v16 := []int{42}
            _ = v16
        }
    }
    if (814 > p2) {
        var v17 int32 = 42
        var v18 int16 = 1000
        _ = v17
        _ = v18
    } else {
        v19 := []int{334, 1, 704, 9223372036854775807}
        fmt.Println(p2)
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -32768
}
func genFn1(p0 uint8) uint {
    if true {
        var v20 int64 = 9223372036854775807
        fmt.Println(v20, p0)
        _ = v20
    } else {
        v21 := GenS1{F0: 42}
        p0 = 42
        _ = v21
    }
    p0 = p0
    p0 += 2
    _ = p0
    return 2
}

func main() {
    var v1 uint16 = 32767
    var v2 uint = 782
    var v3 int8 = 100
    var v4 uint16 = (v1 | (v1 ^ (v1 << ((v2) & 63))))
    v5 := make(map[string]bool)
    var v6 float32 = 1e10
    var v22 int = 32767
    var v23 int = (^(v22))
    delete(v5, (("hello" + "") + "foo bar"))
    var v24 bool = (255 >= 127)
    var v25 uint64 = 1
    for fi26 := 0; fi26 < 3; fi26++ {
        if ((v25 * 4294967295) <= v25) {
            v1 += 1
            fmt.Println(v25, v23)
            var v27 uint8 = 107
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            fmt.Println(fi28, v6)
            _ = fi28
        }
        _ = fi26
    }
    v29 := []int{v23, (v22 + (^(v23)))}
    var v30 bool = false
    var v31 int = (v22 >> (((v2 * v2)) & 63))
    if len(v29) > 0 {
        v29[0] = (v22 + (v23 - 32768))
    }
    v32 := v5["hello"]
    v30 = (!((!v30) || v24))
    switch "hello" {
    case "hello":
        var v33 string = ""
        _ = v33
    case "":
        fmt.Println(v6)
    case "\n\t":
        delete(v5, ("a" + ("hello" + ("foo bar" + "a"))))
    }
    v34 := GenS2{F0: (v25 | 0), F1: 1, F2: v30}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
