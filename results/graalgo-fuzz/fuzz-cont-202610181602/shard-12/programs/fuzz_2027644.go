package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{"a"}
    var v2 int = 502
    var v3 uint16 = 255
    for fi4 := 0; fi4 < 19; fi4++ {
        if true {
            for fi5 := 0; fi5 < 11; fi5++ {
                if true {
                    v6 := GenS1{F0: 2, F1: 1000}
                    v7 := GenS1{F0: 2, F1: 565}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v8 uint8 = 0
                    var v9 bool = ((v3 >> ((2) & 63)) == 42)
                    v9 = v9
                    v2 = (fi5 >> ((18446744073709551615) & 63))
                    _ = v6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi5
            }
            switch (("\n\t" + "\n\t") + "\n\t") {
            case "hello":
                switch "a" {
                case "":
                    fmt.Println(v2)
                }
            }
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        }
        v10 := GenS2{F0: 1, F1: 1}
        _ = fi4
        _ = v10
    }
    var v11 int = v2
    for fi12 := 0; fi12 < 3; fi12++ {
        v11 = fi12
        switch "foo bar" {
        case "":
            var v13 float32 = (0.1 / 3.14159)
            _ = v13
        case "a":
            v2 -= (v2 / ((v11) | 1))
        case "hello":
            v2 -= (v2 * v2)
        }
        _ = fi12
    }
    _ = len(v1)
    var v14 string = ""
    if (127 > 2) {
        v1 = append(v1, (v14 + "a"))
    } else {
        v15 := GenS2{F0: 588, F1: -1}
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v11 = v11
        for fi17 := 0; fi17 < 11; fi17++ {
            switch v11 {
            case -9223372036854775808:
                fmt.Println(v3)
            case 157:
                if false {
                    v11 = (-(v2))
                    fmt.Println(v11, fi16)
                } else {
                    var v18 int32 = -624
                    switch (v11 / ((v11) | 1)) {
                    case -214:
                        var v19 int16 = -1
                        _ = v19
                    case 748:
                        var v20 uint16 = v3
                        _ = v20
                    default:
                        var v21 uint32 = 2147483647
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v1)
                        var v22 string = v14
                        _ = v21
                        _ = v22
                    }
                    v18 = -303
                    _ = v18
                }
            }
            v23 := GenS1{F0: 335, F1: -32768}
            if (!("hello" >= "foo bar")) {
                if false {
                    var v24 float64 = 0.1
                    _ = v24
                } else {
                    switch 9223372036854775807 {
                    case 4294967295:
                        var v25 uint16 = v3
                        _ = v25
                    default:
                        if true {
                            v2 = v11
                            var v26 uint64 = 612
                            _ = v26
                        } else {
                            v1 = append(v1, v14)
                        }
                    }
                }
            } else {
                v2 = fi17
                switch fi17 {
                case 890:
                    v11 = (fi17 | v11)
                case 2147483647:
                    if len(v1) > 0 {
                        v1[0] = v14
                    }
                default:
                    v3 = v3
                }
                fmt.Println(v11, v14, v11)
            }
            _ = fi17
            _ = v23
        }
        fmt.Println(v14, v3)
        _ = fi16
    }
    var v27 uint = 65535
    v28 := GenS3{F0: 244, F1: (-(v2))}
    var v29 uint16 = 0
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 uint32 = 888
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println(v2, v29, v29)
    var v33 int32 = -1
    fmt.Println(v14, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
}
