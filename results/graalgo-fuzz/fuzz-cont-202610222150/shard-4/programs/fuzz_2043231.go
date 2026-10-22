package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 int8; F1 bool }
func genFn0(p0 uint16) int {
    for fi4 := 0; fi4 < 12; fi4++ {
        if (438 >= 0) {
            fmt.Println(fi4)
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := []int{fi4, 1000, (fi5 - int(p0)), fi5}
                v7 := GenS1{F0: 16, F1: -20, F2: -570}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 int64 = 0
            _ = v8
        } else {
            v9 := []int{-48}
            _ = v9
        }
        _ = fi4
    }
    var v10 uint32 = 191
    if (!(42 == 42)) {
        for fi11 := 0; fi11 < 17; fi11++ {
            if false {
                var v12 uint32 = (v10 ^ 661)
                _ = v12
            }
            v10 -= (v10 ^ (^(v10)))
            v13 := GenS2{F0: 0}
            _ = fi11
            _ = v13
        }
    } else {
        var v14 int16 = 32767
        if (((18446744073709551615 <= 255) && (-1 <= -514)) || (9223372036854775807 == -1)) {
            v14 -= 1
            v15 := GenS3{F0: 559, F1: -9223372036854775808}
            _ = v15
        } else {
            switch "foo bar" {
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 uint64 = uint64(p0)
                        fmt.Println(fi16)
                        var v19 int16 = v14
                        v20 := []int{(fi17 ^ (fi17 + 32768)), fi16, 32767, int(v14)}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    switch fi16 {
                    case 32767:
                        v21 := GenS3{F0: 708, F1: 123}
                        _ = v21
                    case 156:
                        if true {
                            var v22 uint8 = 2
                            fmt.Println(v14)
                            _ = v22
                        } else {
                            v23 := []int{(fi16 | fi16), (fi16 ^ fi16), (fi16 / (((fi16 / (((fi16 >> ((65535) & 63))) | 1))) | 1)), fi16}
                            _ = v23
                        }
                    }
                    _ = fi16
                }
            case "foo bar":
                p0 = 42
            }
            v24 := GenS2{F0: v14}
            p0 += (p0 - p0)
            _ = v24
        }
        _ = v14
    }
    _ = p0
    _ = v10
    return 2147483647
}

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := []string{"", (v1 + (v1 + v1))}
    var v3 string = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        if false {
            v1 = v1
            var v25 uint64 = 1
            _ = v25
        }
    }
    var v26 uint16 = 255
    if len(v2) > 0 {
        v2[0] = (v1 + (v1 + ""))
    }
    if false {
        v1 = ((v1 + "\n\t") + v1)
    }
    fmt.Println(v3, v1, v1)
    for fi27 := 0; fi27 < 10; fi27++ {
        if (!true) {
            var v28 uint8 = 255
            v29 := GenS4{F0: 42, F1: 0, F2: v1}
            for fi30 := 0; fi30 < 5; fi30++ {
                fmt.Println(v3, fi30, fi27)
                switch (("" + "foo bar") + "") {
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v31 := GenS5{F0: 42, F1: true}
                    switch (((v3 + v1) + "a") + (v1 + "hello")) {
                    case "":
                        var v32 int8 = 1
                        _ = v32
                    }
                    var v33 uint = 42
                    _ = v31
                    _ = v33
                }
                _ = fi30
            }
            _ = v28
            _ = v29
        }
        var v34 int64 = -9223372036854775808
        _ = fi27
        _ = v34
    }
    var v35 uint8 = 0
    var v36 int = int(v26)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
}
