package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", ("a" + "hello"), ""}
    var v5 uint16 = 255
    var v6 uint = 65535
    v7 := GenS1{F0: 1}
    switch ((("a" + "hello") + ("hello" + "foo bar")) + "") {
    case "foo bar":
        var v8 int = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 127}
        v6 = (v6 | (^(v6)))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 46}
    fmt.Println(v5, v5, v1)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: -122}
        _ = fi12
        _ = v13
    }
    var v14 int64 = -1
    if true {
        fmt.Println(v5, v1)
        v5 -= 305
        v5 += (^(v5))
    } else {
        fmt.Println(v5)
    }
    v15 := GenS2{F0: "hello", F1: -842}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v16 uint32 = 4294967295
    switch "hello" {
    case "":
        if false {
            if (731 != 32767) {
                v14 -= 1
            } else {
                if (!(true && true)) {
                    v17 := GenS1{F0: -44}
                    _ = v17
                } else {
                    switch -688 {
                    case 2147483647:
                        v6 = v6
                    default:
                        _ = len(v4)
                    }
                    if (true && (false || true)) {
                        var v18 uint = 2
                        var v19 int16 = 1000
                        _ = v18
                        _ = v19
                    }
                }
                v20 := GenS3{F0: true, F1: (v5 - v5)}
                _ = v20
            }
            v14 = -1
            if true {
                v1 -= (^(v1))
            } else {
                var v21 uint8 = v1
                _ = v21
            }
        } else {
            fmt.Println(v16)
        }
    }
    var v22 uint32 = v16
    v16 = 46
    v6 += 1
    for fi23 := 0; fi23 < 12; fi23++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
