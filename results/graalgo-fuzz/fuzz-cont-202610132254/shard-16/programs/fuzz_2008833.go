package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 bool; F1 int; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 int16 }
func genFn0() uint32 {
    var v4 uint16 = 42
    v4 = v4
    fmt.Println(v4)
    if true {
        v4 -= (v4 % ((65535) | 1))
    }
    _ = v4
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 897
    var v3 string = "hello"
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            v3 = ("hello" + v3)
            v3 = "hello"
            fmt.Println(v3, v3, v2)
            switch 384 {
            case 0:
                if true {
                    v7 := []int{fi6, (fi6 | (fi5 % ((2147483647) | 1)))}
                    fmt.Println(v2)
                    _ = v7
                }
            case -1:
                v2 += v2
            case -438:
                if (!true) {
                    fmt.Println(fi5, v3, fi5)
                    if (false && (!false)) {
                        v2 = v2
                    } else {
                        v8 := GenS1{F0: v3}
                        var v9 int = fi6
                        v9 += v9
                        _ = v8
                        _ = v9
                    }
                } else {
                    switch (^(fi6)) {
                    case -117:
                        v10 := GenS2{F0: "foo bar", F1: 42, F2: v2}
                        _ = v10
                    case 2147483647:
                        var v11 string = "foo bar"
                        _ = v11
                    default:
                        if (!(true && true)) {
                            v1[fi6] = true
                            v2 += (v2 % (((v2 * v2)) | 1))
                            var v12 int16 = v2
                            _ = v12
                        }
                    }
                    switch (fi6 >> ((18446744073709551615) & 63)) {
                    case 1:
                        v13 := GenS3{F0: false, F1: (v3 + (v3 + v3))}
                        _ = v13
                    case 932:
                        v2 = v2
                    case 42:
                        switch (v3 + "\n\t") {
                        case "a":
                            v14 := GenS4{F0: true, F1: (fi5 | (fi6 ^ fi5)), F2: 42}
                            _ = v14
                        case "hello":
                            fmt.Println(v3)
                        }
                    default:
                        if (1 < 701) {
                            var v15 int64 = 1
                            _ = v15
                        } else {
                            var v16 int8 = -60
                            _ = v16
                        }
                    }
                    var v17 int32 = 42
                    _ = v17
                }
            }
            _ = fi6
        }
        if false {
            v18 := GenS1{F0: "hello"}
            v19 := v1[32768]
            switch 1 {
            case -9223372036854775808:
                if ((!v19) || v19) {
                    v2 += v2
                    if v19 {
                        v20 := GenS4{F0: v19, F1: (fi5 + fi5), F2: 65535}
                        fmt.Println(v3)
                        _ = v20
                    }
                    var v21 int8 = -128
                    _ = v21
                } else {
                    v2 = (v2 + (-(v2)))
                }
            }
            _ = v18
            _ = v19
        }
        _ = fi5
    }
    fmt.Println(v3, v3)
    switch 32767 {
    case 2147483647:
        v3 = "\n\t"
    case 32768:
        var v22 float64 = (0.1 + 0.1)
        _ = v22
    }
    v3 = (v3 + "hello")
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 int32 = 2147483647
        var v25 uint32 = 4294967295
        for fi26 := 0; fi26 < 3; fi26++ {
            v25 += 65535
            v3 = (v3 + v3)
            v27 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 0}
            if false {
                fmt.Println(v2, v25)
                v28 := []int{fi23, (fi23 % ((fi26) | 1)), (fi26 / ((fi26) | 1)), -1}
                if (!(!(2147483648 <= (^(v25))))) {
                    v2 -= (v2 - 1000)
                }
                _ = v28
            }
            v29 := GenS1{F0: v3}
            delete(v1, fi26)
            _ = fi26
            _ = v27
            _ = v29
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS5{F0: ("" > "foo bar")}
        _ = fi30
        _ = v31
    }
    v32 := GenS4{F0: false, F1: 841, F2: 167}
    if false {
        var v33 float64 = 3.14159
        v34 := GenS6{F0: 0, F1: v2}
        _ = v33
        _ = v34
    } else {
        v2 = v2
    }
    v35 := GenS4{F0: (true || true), F1: 128, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
}
