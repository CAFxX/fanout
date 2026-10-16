package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int16 = -608
    v2 := make(map[uint32]int)
    var v3 int64 = -9
    var v4 int32 = -2147483648
    if true {
        fmt.Println(v3, v1, v3)
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint16 = 700
            v7 := GenS1{F0: 255}
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 uint32 = 42
        _ = v8
    }
    v9 := []int{170, -2147483648}
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint = 255
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: int(v1)}
    switch 4294967295 {
    case 9223372036854775807:
        v3 = v3
    }
    switch ("hello" + "foo bar") {
    case "":
        switch 128 {
        case -591:
            var v13 string = "\n\t"
            _ = v13
        case -908:
            var v14 uint32 = 0
            _ = v14
        }
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS2{F0: fi15}
        v17 := GenS3{F0: -120}
        fmt.Println(v3)
        _ = fi15
        _ = v16
        _ = v17
    }
    switch "a" {
    case "":
        if true {
            v4 -= 0
            var v18 int16 = (v1 * 32767)
            for fi19 := 0; fi19 < 2; fi19++ {
                switch (fi19 << ((2) & 63)) {
                case 2147483647:
                    v4 += -923
                case 1:
                    v4 -= 869
                default:
                    v9 = append(v9, 42)
                }
                fmt.Println(v4)
                _ = fi19
            }
            _ = v18
        } else {
            v20 := GenS1{F0: 128}
            _ = v20
        }
    case "a":
        if (!((255 == 2) || (true && (true || false)))) {
            if true {
                v3 = (-(v3))
            } else {
                v2[568] = -923
                delete(v2, 534)
                fmt.Println(v3, v4)
            }
        }
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v3, v4)
        _ = fi21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v3 += (v3 & (^(v3)))
        v3 -= (v3 * v3)
        _ = fi22
    }
    v1 = v1
    var v23 int = 2147483647
    v3 -= (v3 << ((54) & 63))
    delete(v2, 42)
    v3 += -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
