package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }

func main() {
    v1 := []int{1}
    v2 := []string{"a", "a", "hello"}
    var v3 int = 42
    v3 = (-(v3))
    switch "hello" {
    case "a":
        fmt.Println(v3)
    }
    _ = len(v1)
    var v4 float64 = 1e10
    v4 = float64(-2.25)
    var v5 float64 = 1.5
    var v6 int16 = 860
    if true {
        fmt.Println(v6, v5)
        fmt.Println(v6, v5, v4)
    } else {
        fmt.Println(v5)
        v3 = (-(v3))
        var v7 bool = ((!(!(60 != 42))) && (-1 <= 1))
        _ = v7
    }
    _ = len(v2)
    var v8 uint8 = 118
    if true {
        fmt.Println(v6, v4, v3)
    } else {
        switch ("\n\t" + "\n\t") {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            switch (v3 ^ (v3 + (v3 ^ (-(v3))))) {
            case 302:
                v9 := GenS1{F0: 65535, F1: "foo bar", F2: v3}
                _ = v9
            case 2147483647:
                fmt.Println(v4, v4)
            }
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (-(v3)) {
            case -190:
                if (true && false) {
                    fmt.Println(v3, v3)
                }
            }
            switch "hello" {
            case "a":
                fmt.Println(v6, v6, fi10)
            case "":
                v11 := GenS2{F0: 18446744073709551615}
                _ = v11
            }
            _ = fi10
        }
        if (1 == 837) {
            var v12 uint8 = v8
            fmt.Println(v6)
            _ = v12
        } else {
            switch "a" {
            case "foo bar":
                for fi13 := 0; fi13 < 7; fi13++ {
                    v5 = v5
                    _ = fi13
                }
            case "\n\t":
                switch (^(v3)) {
                case 371:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 uint64 = 2
                        v5 = -2.25
                        v3 = v3
                        _ = fi14
                        _ = v15
                    }
                case 2147483647:
                    v8 -= v8
                default:
                    v16 := GenS2{F0: 4294967295}
                    _ = v16
                }
            default:
                var v17 uint = 0
                _ = v17
            }
            v18 := GenS2{F0: 719}
            _ = v18
        }
    }
    switch v3 {
    case -2147483648:
        var v19 int8 = 100
        _ = v19
    case 4294967295:
        var v20 int64 = -1
        _ = v20
    case 1:
        if (51 >= 621) {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint8 = v8
                var v23 uint8 = (v22 << ((65535) & 63))
                if true {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v25 := GenS1{F0: 255, F1: "", F2: fi21}
                        _ = fi24
                        _ = v25
                    }
                    var v26 uint32 = 635
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        fmt.Println(v23, v5)
                        var v28 uint8 = 42
                        _ = len(v2)
                        _ = fi27
                        _ = v28
                    }
                }
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    v29 := GenS3{F0: -852}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += 1e10
    v4 += (v4 + (3.14159 * float64((1e10 - 0.1))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
