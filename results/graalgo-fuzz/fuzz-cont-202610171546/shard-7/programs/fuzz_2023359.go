package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }

func main() {
    var v1 bool = true
    v2 := []string{"\n\t"}
    var v3 uint8 = 1
    var v4 int = 4294967295
    var v5 int32 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v1 {
        fmt.Println(v5, v3, v5)
    } else {
        v6 := GenS1{F0: 0, F1: 2}
        v7 := GenS1{F0: -32768, F1: 18446744073709551615}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -225, F1: 18446744073709551615}
    v3 = (v3 - v3)
    if v1 {
        v2 = append(v2, "")
        fmt.Println(v1, v1, v4)
    } else {
        v1 = v1
        for fi9 := 0; fi9 < 6; fi9++ {
            if v1 {
                var v10 uint32 = 42
                _ = v10
            } else {
                if v1 {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v3 = (^(v3))
                        _ = fi11
                    }
                }
                if v1 {
                    if v1 {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v12 bool = (-128 > 127)
                        _ = v12
                    } else {
                        var v13 bool = (v3 == 77)
                        v1 = v13
                        v14 := GenS2{F0: 42, F1: 4294967295}
                        _ = v13
                        _ = v14
                    }
                }
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint32 = 557
                _ = fi15
                _ = v16
            }
            _ = fi9
        }
        v2 = append(v2, ("" + ("hello" + "\n\t")))
    }
    switch ("" + ("hello" + "foo bar")) {
    case "foo bar":
        if true {
            fmt.Println(v5, v5)
            fmt.Println(v4)
        }
    case "hello":
        if true {
            v5 = (v5 / (((-(v5))) | 1))
        } else {
            var v17 int16 = 684
            if len(v2) > 0 {
                v2[0] = ("hello" + "foo bar")
            }
            _ = v17
        }
    }
    v18 := GenS2{F0: 255, F1: 2}
    fmt.Println(v4, v3, v4)
    v19 := GenS3{F0: 4294967295, F1: 2147483647, F2: 154}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
