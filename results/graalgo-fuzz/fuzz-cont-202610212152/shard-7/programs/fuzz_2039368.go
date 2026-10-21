package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 258
    v3 := []string{"a", ("a" + (("" + "hello") + "")), ((("foo bar" + ("" + (("a" + "hello") + "a"))) + "a") + ((("a" + (("" + "hello") + "a")) + "a") + "hello"))}
    var v4 int = 42
    v2 -= v2
    var v5 float64 = -2.25
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5, v4, v2)
    case "a":
        _ = len(v1)
    case "":
        if ((v5 >= v5) && false) {
            if false {
                if (true || false) {
                    _ = len(v3)
                } else {
                    fmt.Println(v4)
                    fmt.Println(v5)
                    var v6 int16 = -32768
                    _ = v6
                }
            } else {
                fmt.Println(v5, v2)
                v7 := GenS1{F0: 2, F1: 447, F2: 127}
                delete(v1, int(v2))
                _ = v7
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = v5
    var v8 int64 = 1000
    var v9 int64 = v8
    var v10 int8 = 1
    v5 = 0.0
    if true {
        v8 = (^(v9))
        if (false && true) {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS2{F0: v2}
                v13 := GenS3{F0: 1, F1: 179}
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            var v14 bool = true
            var v15 uint = 314
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v2 += (v2 << ((968) & 63))
                _ = fi17
            }
            if false {
                var v18 uint8 = 2
                v2 += (^(v2))
                _ = v18
            }
            v10 -= v10
            _ = fi16
        }
    }
    if (true && true) {
        fmt.Println(v8, v5)
        v4 = (v4 | 32768)
        fmt.Println(v10, v4, v5)
    }
    var v19 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
