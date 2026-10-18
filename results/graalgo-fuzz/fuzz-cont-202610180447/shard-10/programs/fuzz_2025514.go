package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }

func main() {
    var v1 int8 = 1
    v2 := []string{(((("hello" + (("hello" + "") + "foo bar")) + ("" + ("hello" + ("foo bar" + ("a" + ""))))) + "") + "\n\t"), "\n\t"}
    v3 := make(map[string]bool)
    var v4 bool = false
    var v5 int = 42
    v6 := []string{"", ("" + "foo bar")}
    var v7 int8 = 127
    var v8 uint64 = 4294967295
    var v9 uint = 122
    v5 += v5
    if (650 >= 1) {
        var v10 int64 = 0
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: 2147483648, F1: v4}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    }
    var v14 uint16 = 456
    v8 += (v8 << ((v9) & 63))
    var v15 uint16 = v14
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v15, v8)
    v16 := GenS1{F0: 656, F1: v4}
    var v17 int64 = 1
    for fi18 := 0; fi18 < 1; fi18++ {
        _ = len(v6)
        if false {
            var v19 uint16 = 2
            var v20 int16 = 1000
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 42, F1: (false || false)}
            if (true && v4) {
                v22 := GenS2{F0: -32768, F1: "hello"}
                _ = v22
            }
            var v23 int16 = -821
            _ = v21
            _ = v23
        }
        _ = fi18
    }
    var v24 uint = (v9 / (((v9 - v9)) | 1))
    var v25 int = v5
    v26 := GenS3{F0: v8, F1: v1, F2: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
