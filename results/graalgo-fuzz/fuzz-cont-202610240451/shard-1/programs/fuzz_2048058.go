package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int16 }
func genFn0() int64 {
    v4 := []int{-2147483648, 128, 42}
    var v5 uint64 = 1
    var v6 int64 = 166
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    v2 := []string{""}
    var v3 string = ""
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v1, v3, v1)
        switch 42 {
        case 32768:
            v1 -= (v1 << ((2) & 63))
        case 144:
            var v8 int32 = (v1 & (v1 ^ (v1 >> ((42) & 63))))
            _ = v8
        }
        if true {
            var v9 bool = (!("hello" == v3))
            _ = v9
        } else {
            v3 = ""
            fmt.Println(v3, v1, v3)
        }
        _ = fi7
    }
    var v10 int16 = 1
    switch 464 {
    case 9223372036854775807:
        var v11 int16 = v10
        _ = v11
    case 2147483647:
        fmt.Println(v3, v1)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint = 2
            v15 := GenS1{F0: 180}
            _ = fi13
            _ = v14
            _ = v15
        }
        if false {
            v16 := GenS1{F0: 42}
            _ = v16
        }
        v10 += -32768
        _ = fi12
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        _ = len(v2)
        fmt.Println(v3, v10, v3)
        _ = fi17
    }
    if true {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS1{F0: 255}
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 float64 = 1e10
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: 42}
            _ = fi18
            _ = v19
            _ = v22
        }
    } else {
        v1 = 676
        v2 = append(v2, v3)
        if (!((v10 % ((v10) | 1)) < 42)) {
            v3 = v3
        }
    }
    var v23 string = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = 0
    v23 = v23
    v23 = v3
    v3 = "foo bar"
    switch 128 {
    case -990:
        fmt.Println(v10)
    case 124:
        _ = len(v2)
    }
    fmt.Println(v10, v23, v23)
    if true {
        v24 := GenS2{F0: 127, F1: int(v1)}
        switch 4294967295 {
        case -92:
            v25 := GenS2{F0: 1, F1: -419}
            _ = v25
        case -1:
            if (!((128 >= 35) || (false || (128 > 0)))) {
                switch (("a" + "foo bar") + ("a" + "a")) {
                case "\n\t":
                    v26 := GenS3{F0: v10}
                    _ = v26
                }
                var v27 int16 = 1
                var v28 uint32 = 42
                _ = v27
                _ = v28
            }
        }
        _ = v24
    } else {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
}
