package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool) uint8 {
    var v4 int16 = 42
    p0 = p0
    switch "\n\t" {
    case "foo bar":
        var v5 int64 = -739
        _ = v5
    case "":
        switch "foo bar" {
        case "hello":
            p0 = true
        }
    case "hello":
        var v6 string = ((("\n\t" + "") + "") + "")
        _ = v6
    default:
        v7 := []int{2147483647, 156}
        _ = v7
    }
    fmt.Println(v4, v4)
    fmt.Println(v4)
    _ = p0
    _ = v4
    return 237
}
func genFn1() uint32 {
    v8 := GenS1{F0: 128}
    var v9 bool = ("foo bar" <= ("" + ""))
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            v9 = v9
            v12 := GenS2{F0: 42}
            fmt.Println(v9, fi10)
            _ = fi11
            _ = v12
        }
        fmt.Println(fi10)
        _ = fi10
    }
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 int16 = -1
    var v13 uint32 = 549
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 string = ("\n\t" + "foo bar")
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 0
    _ = len(v2)
    v13 = v13
    var v17 uint8 = v16
    var v18 int64 = 69
    for fi19 := 0; fi19 < 17; fi19++ {
        v3 += 1000
        v16 = v17
        _ = fi19
    }
    var v20 float32 = (-2.25 / 3.14159)
    v21 := GenS1{F0: -825}
    if true {
        fmt.Println(v17, v3)
    } else {
        var v22 float64 = 1.5
        for fi23 := 0; fi23 < 19; fi23++ {
            v1["hello"] = ("\n\t" + "foo bar")
            var v24 int8 = -66
            _ = fi23
            _ = v24
        }
        var v25 uint = 65535
        _ = v22
        _ = v25
    }
    fmt.Println(v16, v13, v17)
    switch (("foo bar" + "hello") + "\n\t") {
    case "\n\t":
        v18 -= v18
    case "hello":
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v3)
            var v27 uint = 453
            _ = fi26
            _ = v27
        }
    case "":
        if true {
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS1{F0: fi28}
                var v30 uint8 = 1
                for fi31 := 0; fi31 < 9; fi31++ {
                    var v32 uint = 42
                    _ = fi31
                    _ = v32
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            var v33 bool = true
            if false {
                fmt.Println(v18, v13)
            }
            _ = v33
        }
    default:
        if true {
            var v34 uint = 4294967295
            _ = v34
        }
    }
    fmt.Println(v16, v3)
    var v35 float64 = (-2.25 / 0.1)
    v16 = uint8(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
}
