package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8) int64 {
    var v5 int16 = -364
    p0 = p0
    v6 := GenS1{F0: -9223372036854775808}
    switch -113 {
    case 1:
        v7 := GenS1{F0: 2147483647}
        _ = v7
    default:
        var v8 uint8 = 127
        _ = v8
    }
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    _ = v6
    return -49
}

func main() {
    var v1 int32 = -240
    var v2 float32 = (float32(1.5) + 0.0)
    v3 := []uint32{65535, 4294967295, 86, 4294967295}
    var v4 int = -237
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int = fi9
        for fi11 := 0; fi11 < 20; fi11++ {
            v1 += v1
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS1{F0: (v4 >> ((291) & 63))}
    v3 = append(v3, 2147483647)
    if false {
        fmt.Println(v4, v1, v1)
        v13 := GenS1{F0: 9223372036854775807}
        _ = v13
    } else {
        switch "" {
        case "\n\t":
            v4 += (v4 % ((v4) | 1))
        case "a":
            var v14 int16 = -1
            _ = v14
        case "hello":
            v1 -= v1
        }
        var v15 uint = 42
        _ = v15
    }
    var v16 uint = 255
    _ = len(v3)
    v2 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 uint = (v16 + (v16 / ((v16) | 1)))
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            if false {
                v20 := GenS2{F0: v17}
                v21 := GenS3{F0: 2, F1: 4294967295}
                switch "hello" {
                case "":
                    var v22 string = (("foo bar" + "hello") + ("" + ("hello" + "a")))
                    _ = v22
                case "a":
                    var v23 bool = ("foo bar" < "\n\t")
                    _ = v23
                }
                _ = v20
                _ = v21
            }
            var v24 string = "hello"
            v24 = "foo bar"
            _ = fi19
            _ = v24
        }
        fmt.Println(v1, v16)
        _ = fi18
    }
    _ = len(v3)
    var v25 bool = true
    v26 := GenS3{F0: 230, F1: 2147483647}
    v27 := GenS1{F0: (v4 & 32768)}
    for fi28 := 0; fi28 < 14; fi28++ {
        fmt.Println(v2)
        var v29 int64 = -9223372036854775808
        for fi30 := 0; fi30 < 18; fi30++ {
            v3 = append(v3, 42)
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v25 = v25
        if v25 {
            v32 := GenS3{F0: 32767, F1: 939}
            var v33 string = (("a" + "foo bar") + ("\n\t" + (("\n\t" + "\n\t") + "a")))
            var v34 uint16 = 927
            _ = v32
            _ = v33
            _ = v34
        } else {
            var v35 uint64 = 4294967295
            v1 += v1
            _ = v35
        }
        v25 = (1000 < 1)
        _ = fi31
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v36 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
}
