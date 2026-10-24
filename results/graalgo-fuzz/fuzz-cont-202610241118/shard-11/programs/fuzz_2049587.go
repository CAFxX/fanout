package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = ((0.0 - (0.1 - 1e10)) >= 1.5)
    var v3 uint8 = 238
    v4 := make(map[uint32]int)
    v5 := []int{}
    v1 += (v1 - (v1 & (v1 & 4294967295)))
    if true {
        var v6 int64 = -957
        _ = v6
    } else {
        var v7 float32 = 0.1
        v1 += v1
        _ = v7
    }
    switch "\n\t" {
    case "\n\t":
        if v2 {
            for fi8 := 0; fi8 < 18; fi8++ {
                if (!v2) {
                    fmt.Println(v2)
                } else {
                    v1 -= v1
                    v3 = 255
                    var v9 int8 = -120
                    _ = v9
                }
                var v10 int = fi8
                _ = fi8
                _ = v10
            }
        }
    case "hello":
        v11 := GenS1{F0: 2}
        _ = v11
    case "foo bar":
        v1 += v1
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v3 = (v3 + v3)
        _ = fi12
    }
    if ("a" == "hello") {
        var v13 uint8 = 124
        _ = v13
    }
    fmt.Println(v2)
    var v14 int = -1
    switch ("\n\t" + (("foo bar" + "a") + "a")) {
    case "a":
        if true {
            v15 := GenS1{F0: 32767}
            fmt.Println(v1, v14)
            v2 = (!(42 < 413))
            _ = v15
        }
    case "\n\t":
        var v16 uint64 = 0
        _ = v16
    }
    if (243 != 881) {
        v17 := GenS1{F0: 162}
        v14 -= (^(v14))
        _ = len(v4)
        _ = v17
    } else {
        var v18 uint8 = v3
        v19 := GenS1{F0: 0}
        v20 := GenS1{F0: 255}
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int8 = -128
    var v22 uint8 = 0
    fmt.Println(v14, v2)
    for fi23 := 0; fi23 < 8; fi23++ {
        switch (("hello" + "foo bar") + "hello") {
        case "foo bar":
            v3 = (v3 / (((v3 % ((255) | 1))) | 1))
        case "hello":
            var v24 string = ("" + ("" + "hello"))
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint = 255
            fmt.Println(fi23, v2)
            v3 = (v22 & uint8(v1))
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    delete(v4, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
