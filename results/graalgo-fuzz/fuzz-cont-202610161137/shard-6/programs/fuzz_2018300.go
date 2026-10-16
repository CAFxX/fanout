package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 int16 = 1000
    var v4 int8 = -117
    var v5 bool = true
    var v6 uint = 66
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = append(v2, 459)
        v3 = int16(v6)
        if v5 {
            for fi8 := 0; fi8 < 5; fi8++ {
                fmt.Println(fi8)
                fmt.Println(v6)
                _ = fi8
            }
        } else {
            v2 = append(v2, 65535)
            v3 = v3
        }
        _ = fi7
    }
    if (((("foo bar" + "hello") + "a") + "hello") <= "a") {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v3 -= v3
    }
    if true {
        var v9 int32 = 42
        _ = v9
    } else {
        v3 -= 42
        v3 -= v3
        for fi10 := 0; fi10 < 3; fi10++ {
            v2 = append(v2, 2)
            var v11 int = 32767
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        switch "foo bar" {
        case "hello":
            var v13 int32 = -2147483648
            _ = v13
        case "a":
            var v14 uint32 = 448
            _ = v14
        case "foo bar":
            var v15 int8 = v4
            _ = v15
        }
        v3 = (v3 << ((v6) & 63))
        var v16 string = ""
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi17
    }
    if v5 {
        v18 := GenS1{F0: 1, F1: 1}
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS1{F0: 4294967295, F1: 4294967295}
        for fi21 := 0; fi21 < 13; fi21++ {
            v3 -= v3
            _ = fi21
        }
        fmt.Println(v5, v3, fi19)
        _ = fi19
        _ = v20
    }
    v22 := GenS1{F0: 4294967295, F1: 1}
    v23 := GenS1{F0: 18446744073709551615, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
