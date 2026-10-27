package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{-451, 867, 32767}
    var v2 uint = 0
    var v3 uint = v2
    v4 := make(map[string]string)
    var v5 int16 = -32768
    v6 := []uint32{}
    fmt.Println(v3, v3, v3)
    v3 += (^(v3))
    var v7 float64 = -2.25
    if (!false) {
        if (true && false) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v2 -= 755
                _ = fi8
            }
            var v9 uint64 = 42
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint64 = uint64(v3)
                delete(v4, "a")
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS1{F0: 1, F1: 255}
                _ = fi12
                _ = v13
            }
        }
        v3 += v2
        _ = len(v4)
    } else {
        v14 := v4["foo bar"]
        delete(v4, (v14 + v14))
        _ = v14
    }
    v15 := GenS2{F0: 246}
    if false {
        v2 = (v3 & v2)
    }
    var v16 uint = (v2 << (((v3 >> (((^(v2))) & 63))) & 63))
    delete(v4, (("hello" + "hello") + "a"))
    switch (("hello" + "a") + ("foo bar" + ("a" + "\n\t"))) {
    case "":
        v3 -= v3
    case "a":
        v5 = (v5 + 32767)
    default:
        var v17 int = 2147483647
        _ = v17
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v4[("\n\t" + ("" + "foo bar"))] = ("foo bar" + "foo bar")
        _ = fi18
    }
    if true {
        fmt.Println(v7, v3, v7)
    }
    if ((-751 <= 210) && (0 >= -1)) {
        for fi19 := 0; fi19 < 12; fi19++ {
            if len(v6) > 0 {
                v6[0] = 635
            }
            var v20 uint64 = 42
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v7)
    fmt.Println(v3, v16)
    var v21 string = (("foo bar" + "hello") + "a")
    v22 := GenS1{F0: 601, F1: 570}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
