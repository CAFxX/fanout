package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{1, 449}
    v3 := make(map[string]bool)
    var v4 bool = true
    v5 := make(map[string]string)
    fmt.Println(v4)
    var v6 int64 = 0
    if (-1 <= 930) {
        v4 = true
        if true {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 string = "hello"
                var v9 int64 = (v6 & v6)
                if (!true) {
                    var v10 uint16 = 288
                    v8 = (((((v8 + "") + "\n\t") + "foo bar") + v8) + v8)
                    _ = v10
                } else {
                    v8 = ((v8 + v8) + "foo bar")
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            if true {
                v6 = v6
            } else {
                var v11 float32 = 1.5
                v6 = v6
                if v4 {
                    if v4 {
                        v4 = v4
                    }
                    fmt.Println(v11, v11)
                }
                _ = v11
            }
        } else {
            var v12 float32 = 3.14159
            _ = v12
        }
    }
    switch ("foo bar" + ("a" + "\n\t")) {
    case "\n\t":
        if ("hello" != ("a" + "")) {
            v4 = (!(!(!(false && false))))
        } else {
            fmt.Println(v4, v4)
        }
    case "a":
        v6 = (v6 << ((255) & 63))
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        _ = len(v5)
        _ = fi13
    }
    var v14 uint64 = 508
    var v15 string = "hello"
    var v16 int32 = 1
    var v17 bool = (v6 <= v6)
    var v18 int16 = -32768
    for fi19 := 0; fi19 < 20; fi19++ {
        if true {
            switch v15 {
            case "":
                var v20 int32 = (^(v16))
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint = 42
                    delete(v3, v15)
                    v3[""] = (500 <= 305)
                    v18 += 42
                    if true {
                        var v23 uint32 = 961
                        var v24 uint64 = 42
                        fmt.Println(v14)
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
            default:
                v25 := GenS1{F0: int64(v16)}
                _ = v25
            }
            switch "a" {
            case "":
                var v26 float32 = float32(0.1)
                _ = v26
            default:
                v4 = true
            }
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
