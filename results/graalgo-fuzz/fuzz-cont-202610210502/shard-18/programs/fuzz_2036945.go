package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 string) int16 {
    switch ("" + ("" + ((p1 + (((p1 + ("\n\t" + "hello")) + (("hello" + p1) + p1)) + "")) + (p1 + p1)))) {
    case "foo bar":
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := []int{fi5, p0}
            var v7 int32 = -1
            var v8 uint64 = 130
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 1000
}
func genFn1() int8 {
    var v9 string = ("foo bar" + ("\n\t" + "hello"))
    v10 := []int{}
    if ((v9 + (v9 + "a")) <= ((v9 + ("foo bar" + "\n\t")) + ("foo bar" + v9))) {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int16 = -1
            v10 = append(v10, 692)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v10 = append(v10, fi13)
            _ = fi13
        }
    }
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []string{"hello"}
    var v2 uint64 = 2
    var v3 float64 = (float64(1.5) + (1.5 + (float64(-2.25) + (1e10 - 1.5))))
    var v4 int8 = -8
    fmt.Println(v3, v2, v2)
    var v14 int32 = -1
    var v15 bool = (!true)
    var v16 uint32 = 482
    switch "hello" {
    case "foo bar":
        v1 = append(v1, "a")
    case "":
        if true {
            v3 += v3
            var v17 int16 = 32767
            _ = v17
        }
    }
    switch "" {
    case "hello":
        var v18 uint64 = v2
        _ = v18
    case "foo bar":
        fmt.Println(v2)
    }
    var v19 uint16 = 65535
    var v20 bool = v15
    switch "foo bar" {
    case "\n\t":
        v15 = ((!(3.14159 <= 0.1)) && false)
    case "hello":
        fmt.Println(v3, v14)
    default:
        v2 += v2
    }
    v16 = (v16 << ((65535) & 63))
    fmt.Println(v3, v16, v2)
    var v21 float32 = 0.1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int16 = 1
    var v23 uint8 = 157
    var v24 string = ("\n\t" + "")
    v25 := GenS1{F0: v24, F1: ((!v15) || v15)}
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 string = v24
        var v28 int16 = -32768
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
