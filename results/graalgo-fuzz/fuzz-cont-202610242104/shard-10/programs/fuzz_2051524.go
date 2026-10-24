package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]uint32)
    var v3 float32 = (((1e10 / -2.25) * 1.5) + 1e10)
    var v4 uint8 = 173
    fmt.Println(v4, v3, v4)
    var v5 uint64 = 2
    v6 := []int{-2147483648}
    v7 := v2[1]
    v3 += (v3 / 1.0)
    var v8 int32 = 1
    v9 := GenS1{F0: (("hello" + "foo bar") + "foo bar"), F1: -1, F2: 94}
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v6) > 0 {
            v6[0] = 1
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                var v12 int32 = (-(v8))
                _ = v12
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v3)
    for fi13 := 0; fi13 < 16; fi13++ {
        switch (fi13 ^ fi13) {
        case 128:
            fmt.Println(fi13, v8)
        case 42:
            v14 := GenS1{F0: "foo bar", F1: -1, F2: 56}
            _ = v14
        }
        fmt.Println(fi13)
        var v15 string = "a"
        _ = fi13
        _ = v15
    }
    if (("a" + ("\n\t" + "")) != "\n\t") {
        if false {
            for fi16 := 0; fi16 < 4; fi16++ {
                if ((((!(32767 == -32768)) && true) || false) || true) {
                    v7 -= v7
                }
                if true {
                    if true {
                        var v17 int8 = 1
                        var v18 uint8 = v4
                        _ = v17
                        _ = v18
                    }
                }
                _ = fi16
            }
            v19 := GenS1{F0: "", F1: 71, F2: -1}
            v8 += v8
            _ = v19
        }
    }
    var v20 bool = true
    v4 -= (v4 >> ((255) & 63))
    var v21 uint = 255
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 string = (("" + "foo bar") + ("foo bar" + ("\n\t" + "foo bar")))
        v3 -= v3
        if v20 {
            var v24 uint = 611
            var v25 bool = true
            _ = v24
            _ = v25
        } else {
            fmt.Println(v20, v7, v4)
            if v20 {
                var v26 int64 = 42
                _ = len(v2)
                v26 = v26
                _ = v26
            }
        }
        _ = fi22
        _ = v23
    }
    var v27 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
