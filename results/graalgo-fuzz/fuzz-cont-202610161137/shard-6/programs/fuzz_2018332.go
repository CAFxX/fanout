package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0() uint16 {
    var v5 int16 = -763
    if true {
        v5 += (v5 << ((42) & 63))
        v6 := []int{32767, 231, 1}
        if false {
            v5 = (v5 << ((34) & 63))
            var v7 int8 = 100
            _ = v7
        } else {
            fmt.Println(v5)
        }
        _ = v6
    }
    v8 := []int{}
    fmt.Println(v5)
    _ = v5
    _ = v8
    return 883
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]string)
    var v3 bool = (42 >= 65535)
    v4 := []string{("" + "a"), "a", "a"}
    v3 = v3
    v3 = v3
    v3 = false
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint = 65535
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 2147483647
    if (2 > 0) {
        var v12 string = "\n\t"
        v11 = v11
        for fi13 := 0; fi13 < 16; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (fi13 & -9223372036854775808) {
            case 340:
                var v14 int8 = 1
                _ = v14
            }
            _ = fi13
        }
        _ = v12
    } else {
        v11 -= (v11 - v11)
        v3 = v3
        switch "\n\t" {
        case "":
            for fi15 := 0; fi15 < 10; fi15++ {
                v11 = (v11 >> ((0) & 63))
                var v16 uint64 = 1
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: -2147483648, F1: 32767}
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
    }
    switch "\n\t" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        switch 32768 {
        case 42:
            var v19 float64 = 3.14159
            _ = v19
        case -2147483648:
            if (v11 != 232) {
                v20 := GenS1{F0: -78, F1: 0}
                _ = v20
            } else {
                var v21 uint = 42
                fmt.Println(v21)
                _ = v21
            }
        }
    case "a":
        var v22 uint64 = 2
        _ = v22
    }
    switch (("a" + "hello") + "a") {
    case "":
        var v23 int16 = 32767
        _ = v23
    }
    if v3 {
        v24 := GenS1{F0: 0, F1: -1}
        fmt.Println(v11, v3)
        v11 = v11
        _ = v24
    }
    var v25 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
