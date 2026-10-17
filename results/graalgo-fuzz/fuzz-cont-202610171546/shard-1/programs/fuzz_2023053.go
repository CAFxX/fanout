package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 string, p1 uint32) uint32 {
    switch "a" {
    case "hello":
        fmt.Println(p0, p1)
    case "":
        for fi4 := 0; fi4 < 18; fi4++ {
            v5 := GenS1{F0: -677, F1: 29, F2: (17 != uint(p1))}
            var v6 uint32 = p1
            _ = fi4
            _ = v5
            _ = v6
        }
    case "\n\t":
        v7 := GenS2{F0: 1000, F1: -1}
        _ = v7
    }
    p0 = ("foo bar" + "foo bar")
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    p1 -= (p1 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    return (p1 + p1)
}

func main() {
    v1 := []int{791}
    v2 := []int{32767, -458, -330, 784}
    v3 := make(map[string]uint32)
    var v10 bool = true
    v10 = false
    if ("" > "hello") {
        v11 := GenS3{F0: 42, F1: 128, F2: 96}
        v10 = (!("foo bar" >= "hello"))
        _ = v11
    }
    var v12 uint8 = 140
    if v10 {
        v12 = 1
        v2 = append(v2, 355)
    } else {
        switch -1 {
        case 1:
            v13 := GenS3{F0: 65535, F1: 912, F2: 255}
            _ = v13
        }
        var v14 int32 = 786
        if (!v10) {
            v14 += v14
            var v15 bool = v10
            switch ("foo bar" + (("a" + "hello") + "")) {
            case "hello":
                for fi16 := 0; fi16 < 6; fi16++ {
                    fmt.Println(v10, v15)
                    v10 = v15
                    _ = len(v3)
                    _ = fi16
                }
            case "a":
                fmt.Println(v14)
            case "":
                if (false && false) {
                    var v17 uint32 = 1
                    if false {
                        switch "hello" {
                        case "hello":
                            for fi18 := 0; fi18 < 9; fi18++ {
                                fmt.Println(v10)
                                var v19 int = (fi18 + (fi18 + (fi18 >> ((4294967295) & 63))))
                                _ = fi18
                                _ = v19
                            }
                        }
                    } else {
                        delete(v3, "hello")
                        if v15 {
                            var v20 string = ("hello" + ("foo bar" + "hello"))
                            _ = v20
                        } else {
                            v21 := GenS3{F0: 4294967295, F1: 0, F2: uint16(v12)}
                            _ = v21
                        }
                    }
                    _ = v17
                } else {
                    v3[((("\n\t" + ("hello" + "\n\t")) + "") + "\n\t")] = 526
                    fmt.Println(v15, v10, v12)
                }
            }
            _ = v15
        }
        _ = v14
    }
    var v22 uint = 42
    var v23 uint = v22
    delete(v3, "hello")
    v24 := GenS3{F0: v22, F1: 9223372036854775807, F2: 65535}
    v25 := GenS4{F0: "foo bar"}
    v23 -= (v23 * (v22 & v22))
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v10, v12)
        var v27 uint16 = 255
        v3["a"] = 2147483648
        _ = fi26
        _ = v27
    }
    fmt.Println(v10)
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v12)
        var v29 int = (fi28 * fi28)
        _ = fi28
        _ = v29
    }
    v23 = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
