package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int; F1 uint16; F2 string }
func genFn0(p0 string) uint {
    v5 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8) int16 {
    var v6 int32 = 1
    fmt.Println(v6)
    switch "" {
    case "a":
        if false {
            fmt.Println(p1, p0)
            var v7 uint16 = 0
            _ = v7
        } else {
            switch ("\n\t" + "a") {
            case "":
                var v8 uint8 = (^(p1))
                _ = v8
            case "\n\t":
                if false {
                    v9 := GenS1{F0: "hello"}
                    _ = v9
                }
            default:
                v10 := GenS1{F0: "hello"}
                _ = v10
            }
            switch (("foo bar" + "\n\t") + "hello") {
            case "":
                for fi11 := 0; fi11 < 12; fi11++ {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        fmt.Println(p1)
                        switch "\n\t" {
                        case "\n\t":
                            var v13 uint32 = p0
                            _ = v13
                        case "hello":
                            fmt.Println(p1, p0, v6)
                        default:
                            v14 := GenS2{F0: 0}
                            _ = v14
                        }
                        _ = fi12
                    }
                    _ = fi11
                }
            case "hello":
                var v15 int32 = (v6 * -918)
                _ = v15
            case "foo bar":
                if true {
                    v16 := []int{}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        switch fi17 {
                        case 32768:
                            v6 += v6
                        case 9223372036854775807:
                            if true {
                                var v18 int64 = int64(fi17)
                                v19 := GenS1{F0: "a"}
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = fi17
                    }
                    v6 -= (v6 | -113)
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 313
    var v3 uint16 = 210
    var v4 uint32 = 2
    var v20 int64 = -9223372036854775808
    v21 := GenS1{F0: "\n\t"}
    v22 := GenS3{F0: 128, F1: 145, F2: v3}
    v23 := []int{-9223372036854775808, 127, 0, 800}
    fmt.Println(v2, v2, v20)
    var v24 uint32 = (v4 & (^(v4)))
    if false {
        switch "foo bar" {
        case "foo bar":
            v25 := GenS4{F0: -1, F1: (v3 << ((42) & 63)), F2: ("hello" + "hello")}
            _ = v25
        case "\n\t":
            v2 -= -671
        }
        if len(v23) > 0 {
            v23[0] = -22
        }
    } else {
        fmt.Println(v2, v20, v3)
    }
    v24 += (v4 | (v4 + (^(v24))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
