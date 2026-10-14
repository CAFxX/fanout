package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1e10
    v3 := make(map[uint32]uint32)
    var v4 uint = 59
    var v5 bool = ("" < "hello")
    var v6 int = 953
    v7 := GenS1{F0: 18446744073709551615, F1: 63}
    switch v6 {
    case -1:
        _ = len(v3)
    case -896:
        switch ((("\n\t" + (("\n\t" + "") + "hello")) + "hello") + ("hello" + "\n\t")) {
        case "":
            var v8 int8 = -65
            _ = v8
        case "a":
            fmt.Println(v2)
        }
    }
    switch (v6 << ((v4) & 63)) {
    case -1:
        var v9 uint64 = 0
        _ = v9
    case -5:
        v2 -= float64((0.1 + 3.14159))
    }
    if ((0.1 * -2.25) <= 3.14159) {
        v6 += (v6 / ((9223372036854775807) | 1))
        v4 += 255
        for fi10 := 0; fi10 < 2; fi10++ {
            if (!true) {
                fmt.Println(fi10, fi10)
                v6 = (fi10 << ((42) & 63))
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 int16 = -833
                    v4 += v4
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v5, v6, v5)
            }
            delete(v3, 765)
            _ = fi10
        }
    } else {
        v13 := GenS2{F0: v5}
        v4 = (v4 & (v4 + v4))
        _ = v13
    }
    fmt.Println(v2, v2)
    v2 = 0.1
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 1; fi15++ {
        v6 += v6
        if (!(1 == 2147483647)) {
            var v16 int = (fi15 | v6)
            fmt.Println(v5)
            _ = v16
        }
        v17 := []int{}
        _ = fi15
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int = 128
        fmt.Println(v5, v6)
        _ = fi18
        _ = v19
    }
    var v20 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
