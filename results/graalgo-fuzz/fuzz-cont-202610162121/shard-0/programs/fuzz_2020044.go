package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 string, p1 uint) uint16 {
    var v7 string = (p0 + ((("" + "") + p0) + (p0 + p0)))
    var v8 int64 = int64(p1)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float32 = 0.0
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{-2147483648, 32767, 2147483647, 128}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return 582
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]uint32)
    var v3 int8 = 36
    v4 := make(map[int]uint32)
    v5 := []int{42}
    var v6 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    switch "\n\t" {
    case "hello":
        _ = len(v4)
    case "a":
        if true {
            v15 := v4[349]
            _ = v15
        } else {
            v1 = -1
            var v16 uint8 = 127
            _ = v16
        }
    }
    var v17 bool = (!(((v3 & (v3 - (v3 / ((v3) | 1)))) <= (v3 / (((v3 & 45)) | 1))) && true))
    v18 := v4[128]
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            switch fi20 {
            case -80:
                v3 = 42
            case -295:
                fmt.Println(v18, v3)
            }
            var v21 int64 = (v6 ^ (v6 >> ((344) & 63)))
            v22 := GenS1{F0: ("\n\t" + "")}
            _ = fi20
            _ = v21
            _ = v22
        }
        v1 = v1
        _ = fi19
    }
    v23 := GenS1{F0: "hello"}
    var v24 uint32 = (v18 << ((4294967295) & 63))
    switch "foo bar" {
    case "":
        var v25 int = 42
        _ = v25
    case "foo bar":
        var v26 uint16 = 805
        _ = v26
    case "\n\t":
        v6 -= -9223372036854775808
    }
    v24 = v18
    var v27 uint16 = 255
    fmt.Println(v17, v18)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
