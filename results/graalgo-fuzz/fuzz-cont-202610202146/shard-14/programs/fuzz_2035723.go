package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 int8) int32 {
    fmt.Println(p1, p1)
    if p0 {
        fmt.Println(p0, p0)
    }
    p0 = false
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]string)
    var v3 string = ("a" + "\n\t")
    v4 := make(map[int]bool)
    fmt.Println(v3)
    switch 642 {
    case -884:
        var v5 uint = 2
        _ = v5
    }
    var v6 bool = (!true)
    var v7 int = 4294967295
    if (1 <= 222) {
        fmt.Println(v7, v3, v3)
        v1 = v1
        v6 = (42 <= 127)
    } else {
        v8 := []int{(v7 * 128)}
        v6 = (100 < 0)
        if v6 {
            fmt.Println(v1, v7)
            fmt.Println(v1, v3, v3)
            var v9 uint32 = 4294967295
            _ = v9
        } else {
            v10 := GenS1{F0: 462}
            fmt.Println(v6, v7)
            _ = v10
        }
        _ = v8
    }
    v7 = v7
    v7 = (v7 << ((0) & 63))
    var v11 uint32 = 65535
    v6 = false
    switch v3 {
    case "foo bar":
        for fi12 := 0; fi12 < 19; fi12++ {
            v6 = (0 <= 0)
            v6 = (0.1 == v1)
            v13 := GenS2{F0: 2147483647}
            _ = fi12
            _ = v13
        }
    case "hello":
        delete(v2, v7)
    case "":
        v14 := v4[v7]
        _ = v14
    default:
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint64 = 906
            var v17 uint16 = 18
            for fi18 := 0; fi18 < 10; fi18++ {
                switch v3 {
                case "foo bar":
                    v19 := GenS1{F0: 0}
                    v16 += 18446744073709551615
                    switch "\n\t" {
                    case "hello":
                        var v20 int64 = 982
                        _ = v20
                    case "a":
                        if v6 {
                            var v21 uint8 = 109
                            var v22 int8 = 0
                            _ = v21
                            _ = v22
                        } else {
                            switch (("foo bar" + ("" + v3)) + (((v3 + (v3 + ((v3 + v3) + ""))) + v3) + ("" + v3))) {
                            case "":
                                v16 = v16
                            case "a":
                                var v23 uint = 0
                                _ = v23
                            case "\n\t":
                                if false {
                                    var v24 int8 = -111
                                    var v25 uint = 255
                                    var v26 uint16 = (^(v17))
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        }
                    }
                    _ = v19
                default:
                    v3 = v3
                    v16 += (v16 - 4294967295)
                    if true {
                        v27 := GenS1{F0: 42}
                        _ = v27
                    }
                }
                var v28 int32 = 2147483647
                fmt.Println(fi15, v7, v6)
                var v29 uint16 = (v17 | 265)
                v30 := GenS2{F0: 0}
                _ = fi18
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
