package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    v1 := []uint32{4294967295, 2, 0, 0}
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]uint32)
    var v5 string = (("\n\t" + (("a" + "foo bar") + "")) + ("a" + ("hello" + "a")))
    v6 := GenS1{F0: 42, F1: ""}
    if false {
        if false {
            if (2147483647 == 468) {
                if (631 >= 255) {
                    if true {
                        if false {
                            delete(v2, "")
                            _ = len(v2)
                        }
                        fmt.Println(v5, v5)
                    } else {
                        var v7 string = "a"
                        fmt.Println(v3, v5, v3)
                        for fi8 := 0; fi8 < 7; fi8++ {
                            fmt.Println(v3)
                            v7 = "\n\t"
                            _ = fi8
                        }
                        _ = v7
                    }
                    v9 := GenS1{F0: 0, F1: v5}
                    _ = v9
                }
            }
            if (50 >= -128) {
                v3 -= 1
                for fi10 := 0; fi10 < 5; fi10++ {
                    v5 = v5
                    for fi11 := 0; fi11 < 13; fi11++ {
                        var v12 bool = false
                        _ = fi11
                        _ = v12
                    }
                    var v13 int32 = 1
                    _ = fi10
                    _ = v13
                }
                if (!(true && (-2.25 >= ((0.1 * 1e10) * 1.5)))) {
                    fmt.Println(v5, v3)
                }
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 int8 = -118
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v15, fi16, fi16)
                        v5 = v5
                        _ = fi16
                    }
                    _ = fi14
                    _ = v15
                }
            }
            if (!(false || ((-2.25 * 1e10) == 3.14159))) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v3 = v3
                    _ = fi17
                }
                v18 := GenS1{F0: -1, F1: v5}
                _ = v18
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        fmt.Println(v5)
        v5 = "hello"
    }
    v19 := GenS1{F0: -1, F1: "hello"}
    fmt.Println(v5)
    switch -55 {
    case 2147483647:
        _ = len(v2)
    case 127:
        switch 32767 {
        case 42:
            v20 := GenS2{F0: 599, F1: 383}
            _ = v20
        case 4294967295:
            var v21 bool = true
            _ = v21
        case 2147483647:
            v5 = v5
        }
    case 1:
        v5 = "a"
    }
    delete(v2, ((v5 + "") + ((("hello" + "a") + v5) + (((((v5 + ("" + v5)) + (((((v5 + v5) + (("hello" + "\n\t") + ((v5 + v5) + v5))) + "") + ((v5 + v5) + "hello")) + v5)) + "") + ("foo bar" + v5)) + v5))))
    if (false && ((0 >= 255) || (141 != 185))) {
        v22 := GenS2{F0: 42, F1: 255}
        v23 := v2[v5]
        if (false && ((false && (-2.25 > (float32(1.5) * 1.5))) || (912 >= 0))) {
            if len(v1) > 0 {
                v1[0] = 42
            }
            v24 := GenS1{F0: 103, F1: v5}
            _ = v24
        } else {
            if true {
                v5 = ((("a" + ((v5 + v5) + ("a" + ""))) + "") + "")
                fmt.Println(v23, v5, v5)
            }
            v5 = ""
        }
        _ = v22
        _ = v23
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        switch (v5 + ("\n\t" + v5)) {
        case "":
            var v26 float64 = 1.5
            _ = v26
        case "a":
            fmt.Println(v5)
        }
        switch fi25 {
        case 4294967295:
            var v27 int = fi25
            _ = v27
        case 143:
            var v28 uint = 2
            _ = v28
        default:
            v3 += v3
        }
        _ = fi25
    }
    if (!(32767 >= 847)) {
        var v29 uint8 = 2
        _ = v29
    } else {
        v5 = (v5 + v5)
        var v30 uint8 = 42
        _ = v30
    }
    v31 := GenS3{F0: 495, F1: 127, F2: 1}
    v3 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
