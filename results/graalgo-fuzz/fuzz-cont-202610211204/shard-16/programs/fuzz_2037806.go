package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int }
func genFn0() int8 {
    if (128 > 1) {
        if false {
            v7 := GenS1{F0: "\n\t", F1: 4294967295, F2: -667}
            var v8 int16 = -291
            var v9 int8 = -122
            _ = v7
            _ = v8
            _ = v9
        } else {
            switch (("a" + "a") + ("\n\t" + "foo bar")) {
            case "a":
                var v10 uint16 = 0
                _ = v10
            case "\n\t":
                var v11 string = "a"
                _ = v11
            case "hello":
                v12 := GenS1{F0: (("a" + "\n\t") + "\n\t"), F1: 18446744073709551615, F2: 9223372036854775807}
                _ = v12
            default:
                var v13 bool = (0 > -1)
                _ = v13
            }
        }
    }
    switch 42 {
    case 269:
        v14 := []int{-9223372036854775808, 0}
        _ = v14
    case 42:
        var v15 float32 = -2.25
        _ = v15
    }
    switch "hello" {
    case "\n\t":
        if (881 <= 1) {
            v16 := GenS1{F0: "", F1: 2, F2: 162}
            _ = v16
        } else {
            var v17 uint64 = 593
            v17 = v17
            _ = v17
        }
    }
    if true {
        v18 := GenS1{F0: ("a" + "hello"), F1: 235, F2: 32767}
        v19 := GenS1{F0: "a", F1: 4294967295, F2: 127}
        _ = v18
        _ = v19
    }
    var v20 float32 = 1e10
    _ = v20
    return -128
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{879, 161}
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483647
    v5 := make(map[string]int)
    var v6 uint32 = v4
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: "a", F1: 42, F2: 270}
        fmt.Println(v4)
        v1 += (v1 >> ((2) & 63))
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v1 = 2147483647
        var v24 int16 = -575
        if (true && false) {
            v24 += v24
            for fi25 := 0; fi25 < 11; fi25++ {
                if (335 < 1) {
                    var v26 uint64 = 93
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v27 int16 = 32767
                    _ = v26
                    _ = v27
                } else {
                    var v28 int8 = 127
                    v29 := GenS1{F0: "hello", F1: 319, F2: 42}
                    v24 = (v24 | (v24 % ((v24) | 1)))
                    fmt.Println(fi25, fi25, v24)
                    _ = v28
                    _ = v29
                }
                _ = len(v2)
                _ = fi25
            }
        }
        _ = fi23
        _ = v24
    }
    v6 += v4
    var v30 bool = (!false)
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 bool = (int64(fi31) < 9223372036854775807)
        var v33 uint64 = 42
        switch "hello" {
        case "hello":
            var v34 int64 = 1
            _ = v34
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v30 = false
    _ = len(v2)
    fmt.Println(v30, v1, v1)
    v35 := v3[v4]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
