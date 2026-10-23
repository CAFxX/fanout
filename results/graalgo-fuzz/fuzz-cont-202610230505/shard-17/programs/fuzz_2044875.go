package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int; F1 int; F2 int }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 uint64 = 2
    fmt.Println(v3, v3)
    v3 += (v3 | (v3 >> ((65535) & 63)))
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 int8 = -128
            _ = fi5
            _ = v6
        }
        v3 = (v3 - 18446744073709551615)
        _ = fi4
    }
    if (!(true && (true && (400 == 2147483647)))) {
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(fi7, fi7)
                _ = fi7
            }
        } else {
            v2[0] = false
        }
        v1 = ("a" + (v1 + "\n\t"))
    } else {
        v8 := GenS1{F0: -200, F1: -194}
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v1)
        v3 += v3
        _ = fi9
    }
    v10 := []int{127}
    var v11 uint32 = 2
    var v12 uint8 = uint8(v11)
    v12 = v12
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v13 float64 = 1.5
    v3 = (v3 >> ((uint(v11)) & 63))
    v14 := v2[882]
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 943, F1: -2147483648}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS2{F0: v12, F1: -128}
        v14 = v14
        for fi19 := 0; fi19 < 20; fi19++ {
            v1 = ("\n\t" + "a")
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v20 := GenS3{F0: -495, F1: 4294967295, F2: -871}
    if len(v10) > 0 {
        v10[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
}
