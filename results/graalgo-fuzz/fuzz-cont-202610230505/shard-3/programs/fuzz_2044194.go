package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 1
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 uint32 = v1
    v4[(^(v5))] = "hello"
    for fi6 := 0; fi6 < 9; fi6++ {
        v1 -= 42
        if ((-666 > -681) && (!(0 > (fi6 / ((-1) | 1))))) {
            v7 := v4[(v5 * 622)]
            var v8 bool = ("\n\t" != v7)
            _ = v7
            _ = v8
        }
        var v9 int64 = int64(v2)
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        if (32767 >= 1) {
            fmt.Println(v2)
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(fi10, v1, fi10)
                v12 := GenS1{F0: ("" + "")}
                _ = fi11
                _ = v12
            }
        } else {
            var v13 float64 = 3.14159
            _ = v13
        }
        _ = fi10
    }
    var v14 uint8 = 128
    v2 = v2
    var v15 float64 = (0.0 + (-2.25 * 1e10))
    v1 = v1
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            switch fi16 {
            case -855:
                var v17 uint16 = 65535
                _ = v17
            }
        } else {
            fmt.Println(v1, v2)
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int64 = 42
            v1 += 4294967295
            v15 = v15
            _ = fi18
            _ = v19
        }
        fmt.Println(v15, v5, v15)
        _ = fi16
    }
    v20 := []int{4294967295, 2147483647, 32767}
    if ((false && (((1.5 - (0.0 * 3.14159)) > 1e10) || (!true))) || true) {
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (fi21 * fi21) {
            case 101:
                switch ("" + "hello") {
                case "":
                    v5 = v5
                case "hello":
                    v22 := GenS1{F0: "hello"}
                    _ = v22
                default:
                    var v23 uint16 = 2
                    _ = v23
                }
            }
            _ = fi21
        }
    }
    var v24 string = ("" + "")
    fmt.Println(v24, v15, v2)
    var v25 uint32 = v5
    for fi26 := 0; fi26 < 12; fi26++ {
        fmt.Println(v14)
        v27 := GenS2{F0: 0, F1: -78}
        _ = fi26
        _ = v27
    }
    if len(v20) > 0 {
        v20[0] = 32767
    }
    var v28 int16 = 0
    if (-728 >= -1) {
        fmt.Println(v24)
        var v29 int = 127
        if true {
            fmt.Println(v28)
        } else {
            var v30 int = 9223372036854775807
            _ = v30
        }
        _ = v29
    }
    v31 := v3["a"]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
