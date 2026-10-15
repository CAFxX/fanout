package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 160
    var v2 int = 42
    var v3 uint8 = 0
    v4 := []string{"a", "hello"}
    var v5 float32 = (1e10 + 1e10)
    var v6 uint16 = 2
    v1 = 148
    v4 = append(v4, "foo bar")
    if ((true || (892 < -379)) || false) {
        v3 += (v3 % ((v3) | 1))
    } else {
        v7 := GenS1{F0: 65535, F1: (^(v1)), F2: v6}
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            v5 -= (v5 / v5)
            if false {
                v2 = fi8
                var v9 uint = 681
                _ = v9
            }
            fmt.Println(v6)
            _ = fi8
        }
        _ = v7
    }
    var v10 int8 = 0
    v5 -= -2.25
    if (1 <= 1) {
        fmt.Println(v3)
    }
    switch (("a" + ("a" + "\n\t")) + "foo bar") {
    case "\n\t":
        var v11 int64 = 1
        _ = v11
    default:
        if len(v4) > 0 {
            v4[0] = (("foo bar" + "hello") + "")
        }
    }
    if true {
        switch "hello" {
        case "\n\t":
            v4 = append(v4, ("foo bar" + "foo bar"))
        }
    } else {
        var v12 float64 = ((0.1 + (1e10 - (1.5 - 3.14159))) * 1e10)
        if true {
            _ = len(v4)
            v4 = append(v4, ("" + ""))
            var v13 uint = 255
            _ = v13
        }
        _ = v12
    }
    var v14 uint = 862
    if false {
        if (!(0 == 2)) {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 float64 = (float64(float32(1.5)) / 1.0)
                v17 := GenS1{F0: v14, F1: (v1 - v1), F2: (^(v6))}
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 int16 = 0
                    _ = fi18
                    _ = v19
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(v10, v2)
            var v20 uint8 = v3
            _ = v20
        }
        var v21 int = (v2 + (v2 % (((v2 + 42)) | 1)))
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(v6, fi22, v10)
            v1 += (v1 % (((v1 + uint32(v6))) | 1))
            _ = fi22
        }
        _ = v21
    } else {
        if false {
            var v23 string = "a"
            _ = v23
        }
    }
    v10 += (-(v10))
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 4; fi25++ {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v26 uint = (^(v14))
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = fi27
    }
    fmt.Println(v2, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
