package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        switch "foo bar" {
        case "hello":
            v5 := []int{}
            _ = v5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    v8 := []int{1, 950, 0}
    _ = v8
    return 65535
}
func genFn1(p0 int8) uint8 {
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 32767}
        var v11 int32 = -2147483648
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "a":
        p0 = -1
    case "hello":
        v12 := []int{}
        _ = v12
    default:
        var v13 string = (("" + "hello") + ("" + (("" + "a") + "\n\t")))
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{551, (-(fi14)), fi14}
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint32 = 1
            _ = fi16
            _ = v17
        }
        v1 += v1
        _ = fi14
        _ = v15
    }
    var v18 uint8 = 39
    if false {
        v1 = (v1 * (-(v1)))
        var v19 uint8 = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            delete(v3, 4294967295)
            _ = fi20
        }
        _ = len(v3)
        var v21 uint16 = 255
        _ = v21
    }
    v18 = 0
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{1, fi22, 4294967295, fi22}
        var v24 int32 = 2147483647
        switch "" {
        case "foo bar":
            v24 = 0
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(fi25)
        _ = fi25
    }
    switch 128 {
    case 0:
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS2{F0: 42}
            _ = fi26
            _ = v27
        }
    case 4294967295:
        v28 := v2[""]
        _ = v28
    case -886:
        v29 := GenS1{F0: 65535}
        _ = v29
    }
    v30 := []int{-573}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
}
