package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 string }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint64 = v1
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5, fi5, v4)
        var v6 int = (-(fi5))
        var v7 string = "hello"
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 32767, F1: 0, F2: 0}
    fmt.Println(v1)
    if false {
        var v9 float32 = -2.25
        var v10 int = 517
        for fi11 := 0; fi11 < 16; fi11++ {
            v1 += v1
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint16 = 255
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    } else {
        v14 := GenS2{F0: 42, F1: 255, F2: "hello"}
        var v15 float64 = 0.1
        v16 := GenS2{F0: 127, F1: 18446744073709551615, F2: (((("a" + ("a" + "foo bar")) + ("foo bar" + "")) + (((("foo bar" + "hello") + (("a" + "hello") + "")) + "a") + "")) + "\n\t")}
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := []int{42, 127, 0, 2147483647}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint16 = 32767
        if (false && (-1 <= 42)) {
            fmt.Println(fi18, v19, v4)
            for fi20 := 0; fi20 < 11; fi20++ {
                v4 = v4
                _ = fi20
            }
        } else {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v4 = uint64(v19)
            var v21 bool = true
            _ = v21
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            switch "\n\t" {
            case "hello":
                v19 = v19
            }
            _ = fi22
        }
        _ = fi18
        _ = v19
    }
    v4 += 4294967295
    var v23 int64 = 364
    if false {
        v1 = v4
        fmt.Println(v23, v4, v23)
        v24 := v2[32767]
        _ = v24
    } else {
        fmt.Println(v1, v23, v4)
    }
    v25 := GenS2{F0: -84, F1: 0, F2: ""}
    v26 := GenS1{F0: 0, F1: 42, F2: -194}
    var v27 uint8 = 201
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
