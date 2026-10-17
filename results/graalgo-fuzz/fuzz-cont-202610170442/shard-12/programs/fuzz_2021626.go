package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0() uint16 {
    var v7 int64 = 0
    switch "" {
    case "hello":
        var v8 bool = (65535 >= 65535)
        _ = v8
    case "\n\t":
        v7 = -565
    case "":
        v7 = v7
    }
    fmt.Println(v7)
    _ = v7
    return 255
}
func genFn1(p0 int32) int {
    v9 := GenS1{F0: 42}
    var v10 uint64 = 147
    if true {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int16 = 871
            _ = fi11
            _ = v12
        }
        fmt.Println(v10)
        v13 := []int{-184, 9223372036854775807, 32767}
        _ = v13
    }
    v14 := GenS1{F0: 0}
    _ = p0
    _ = v10
    _ = v14
    _ = v9
    return 0
}

func main() {
    v1 := []string{"foo bar", "a", "foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := []int{-90, -2147483648, -9223372036854775808, 1}
    var v5 string = ((("a" + ("a" + "a")) + "foo bar") + "")
    var v6 uint = 513
    v2[v5] = 4294967295
    fmt.Println(v5)
    if true {
        v15 := v2[("hello" + (v5 + v5))]
        _ = v15
    } else {
        _ = len(v1)
        v16 := GenS1{F0: 0}
        _ = v16
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "\n\t" {
        case "foo bar":
            var v17 uint16 = 32767
            _ = v17
        }
        v5 = v5
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            if true {
                v19 := GenS1{F0: 0}
                if true {
                    fmt.Println(fi18, v6, v5)
                    v20 := GenS2{F0: 2147483647, F1: 1, F2: 255}
                    _ = v20
                } else {
                    v21 := GenS3{F0: true, F1: -1, F2: 2}
                    v22 := GenS4{F0: 1, F1: 2147483647, F2: 753}
                    _ = v21
                    _ = v22
                }
                _ = len(v2)
                _ = v19
            }
            _ = fi18
        }
    }
    switch (v5 + (v5 + "")) {
    case "":
        if true {
            var v23 string = v5
            _ = v23
        } else {
            for fi24 := 0; fi24 < 6; fi24++ {
                v6 = 255
                _ = fi24
            }
        }
    }
    v25 := v3[32768]
    v26 := GenS1{F0: 614}
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v25)
        v28 := GenS4{F0: 4294967295, F1: 0, F2: 9223372036854775807}
        delete(v2, (v5 + ((v5 + ("\n\t" + (v5 + "a"))) + ("foo bar" + (v5 + (v5 + ("a" + v5)))))))
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
