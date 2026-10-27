package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 int = 4294967295
    var v4 int32 = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        v1 = (v1 / 1.5)
        _ = fi5
    }
    var v6 int8 = 1
    v3 = (v3 ^ v3)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{v3, fi7, v3}
        v4 += v4
        _ = fi7
        _ = v8
    }
    var v9 int = (v3 << ((880) & 63))
    if true {
        if true {
            v10 := []int{(v9 % ((v3) | 1))}
            v4 += 1
            v4 = (v4 ^ (v4 & 1))
            _ = v10
        } else {
            v4 -= v4
            fmt.Println(v6, v3, v3)
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint64 = 880
                _ = fi11
                _ = v12
            }
        }
        fmt.Println(v3, v4)
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := []int{32768, 32768}
            delete(v2, (("foo bar" + "\n\t") + (("foo bar" + (("hello" + (("foo bar" + "hello") + ("foo bar" + "\n\t"))) + "\n\t")) + ("a" + ("" + "a")))))
            switch (("a" + ("a" + ("" + "hello"))) + ("a" + "\n\t")) {
            case "foo bar":
                v16 := GenS1{F0: 1, F1: 622}
                _ = v16
            case "\n\t":
                _ = len(v2)
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v17 bool = false
            v15 = append(v15, fi14)
            _ = fi14
            _ = v15
            _ = v17
        }
        _ = fi13
    }
    v1 = 3.14159
    v18 := []int{v9, 1, 127}
    fmt.Println(v9, v3, v6)
    v1 = v1
    v9 += int(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
