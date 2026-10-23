package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    var v5 uint8 = 83
    var v6 string = "foo bar"
    v6 = v6
    v7 := []int{-9223372036854775808, 9223372036854775807}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() int {
    if true {
        var v8 int8 = 14
        _ = v8
    } else {
        var v9 int64 = 1
        fmt.Println(v9)
        _ = v9
    }
    switch (("\n\t" + "a") + "") {
    case "foo bar":
        if (18446744073709551615 == 591) {
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(fi10)
                switch "a" {
                case "a":
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 int64 = -9223372036854775808
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi10
            }
            switch "foo bar" {
            case "foo bar":
                if (-120 >= 42) {
                    v13 := GenS1{F0: -1, F1: ("hello" + "foo bar"), F2: 418}
                    _ = v13
                }
            default:
                switch ("hello" + ("a" + ("a" + "foo bar"))) {
                case "":
                    var v14 int8 = 100
                    _ = v14
                }
            }
        }
    case "hello":
        v15 := []int{32768, 1, 2147483647, 126}
        _ = v15
    case "":
        var v16 uint16 = 255
        _ = v16
    }
    return 9223372036854775807
}

func main() {
    var v1 bool = (-1 < 32768)
    var v2 int = -1
    var v3 uint16 = 65535
    var v4 int64 = 0
    v3 -= v3
    var v17 float64 = 0.0
    v18 := GenS2{F0: 35, F1: 7, F2: 1}
    if (!v1) {
        var v19 int16 = 32767
        var v20 string = "\n\t"
        v21 := GenS1{F0: v4, F1: (v20 + (v20 + ((("\n\t" + v20) + "foo bar") + ""))), F2: (v19 / ((v19) | 1))}
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println(v17, v4, v2)
    var v22 float32 = 1e10
    v2 += -9223372036854775808
    for fi23 := 0; fi23 < 2; fi23++ {
        switch (fi23 / ((-2147483648) | 1)) {
        case 90:
            v24 := []int{}
            _ = v24
        case 1:
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := []int{}
                _ = fi25
                _ = v26
            }
        case 656:
            if false {
                v4 = -1
                for fi27 := 0; fi27 < 6; fi27++ {
                    v4 = v4
                    v28 := GenS2{F0: 57, F1: -20, F2: 65535}
                    _ = fi27
                    _ = v28
                }
            }
        }
        v29 := []int{847, (fi23 + (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))), (fi23 >> ((4294967295) & 63)), (v2 / (((v2 | -501)) | 1))}
        _ = fi23
        _ = v29
    }
    v22 -= v22
    for fi30 := 0; fi30 < 13; fi30++ {
        v2 = (fi30 >> ((42) & 63))
        var v31 string = ("a" + "")
        _ = fi30
        _ = v31
    }
    var v32 int64 = v4
    fmt.Println(v3)
    v33 := GenS1{F0: 42, F1: "\n\t", F2: 32767}
    v34 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
