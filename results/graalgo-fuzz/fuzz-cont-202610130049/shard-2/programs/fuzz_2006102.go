package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }

func main() {
    var v1 string = (("\n\t" + "a") + "foo bar")
    var v2 bool = ((true || (!(127 <= 154))) || (127 >= 42))
    v3 := make(map[uint32]uint32)
    var v4 int32 = -1
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float32 = 3.14159
        _ = fi5
        _ = v6
    }
    v2 = v2
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint = 255
            switch (v1 + ("" + v1)) {
            case "hello":
                v3[65535] = 2147483648
            }
            _ = fi8
            _ = v9
        }
        switch fi7 {
        case -9223372036854775808:
            v10 := GenS1{F0: (v1 > ("" + ((v1 + "a") + v1))), F1: v1, F2: 2}
            _ = v10
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            if v2 {
                var v12 float64 = 1e10
                v1 = v1
                _ = v12
            }
            _ = fi11
        }
        _ = fi7
    }
    v1 = (v1 + "hello")
    fmt.Println(v4, v4)
    if v2 {
        var v13 float32 = 3.14159
        _ = v13
    }
    fmt.Println(v4, v2)
    if v2 {
        v4 += v4
    } else {
        v2 = (!false)
        v1 = v1
        v4 += v4
    }
    v14 := GenS2{F0: false, F1: 18446744073709551615, F2: 758}
    fmt.Println(v4, v1)
    var v15 int8 = -1
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS1{F0: (766 > 0), F1: v1, F2: 1}
            var v19 int32 = v4
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v20 := []int{-254, 9223372036854775807, 797}
    switch -2147483648 {
    case -1:
        _ = len(v20)
    default:
        v2 = v2
    }
    v2 = ((42 < 42) && false)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
