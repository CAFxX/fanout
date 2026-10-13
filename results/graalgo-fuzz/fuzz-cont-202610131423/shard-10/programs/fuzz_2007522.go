package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint, p1 int64, p2 int32) int {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: fi5}
            v7 := []int{112, (^(fi5))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: fi4}
        v9 := []int{(fi4 % ((fi4) | 1))}
        _ = fi4
        _ = v8
        _ = v9
    }
    v10 := []int{0, 772, 42, 2147483647}
    v11 := GenS2{F0: 1, F1: -1}
    var v12 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 2147483647
}
func genFn1(p0 int64, p1 uint16, p2 string) int64 {
    p1 = p1
    var v13 string = ((("" + "hello") + "a") + ("hello" + "foo bar"))
    if (true && true) {
        p0 -= (p0 - p0)
    }
    switch (((v13 + p2) + p2) + v13) {
    case "foo bar":
        switch -133 {
        case 127:
            v14 := []int{-2147483648, 32767}
            _ = v14
        case 2147483647:
            if true {
                fmt.Println(p2, p0, p1)
                v15 := []int{}
                _ = v15
            }
        }
    case "":
        v16 := []int{0, 2147483647, 128}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    var v3 int = 127
    switch (v3 / (((^(v3))) | 1)) {
    case 42:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint = 746
            _ = fi17
            _ = v18
        }
    case -573:
        v3 -= (v3 | (v3 & v3))
    }
    switch 9223372036854775807 {
    case 127:
        var v19 uint16 = 636
        _ = v19
    default:
        v20 := v2[(v3 ^ v3)]
        _ = v20
    }
    v21 := v2[v3]
    v3 += 0
    var v22 uint = 1
    var v23 int8 = 0
    if (!(-2147483648 >= -1)) {
        switch v1 {
        case "a":
            var v24 int64 = 1
            _ = v24
        case "foo bar":
            fmt.Println(v21, v1, v3)
        case "\n\t":
            var v25 uint8 = 255
            _ = v25
        }
        fmt.Println(v3, v21, v1)
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := []int{fi26, (fi26 % (((^(fi26))) | 1))}
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
