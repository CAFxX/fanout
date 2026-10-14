package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 42
    var v3 int8 = -128
    v4 := make(map[int]bool)
    var v5 string = "\n\t"
    var v6 float64 = (float64(-2.25) + -2.25)
    var v7 bool = false
    switch v5 {
    case "a":
        v8 := GenS1{F0: -396}
        _ = v8
    }
    switch v5 {
    case "foo bar":
        v5 = (("foo bar" + "hello") + v5)
    case "a":
        for fi9 := 0; fi9 < 11; fi9++ {
            if v7 {
                var v10 uint = 65535
                _ = v10
            }
            switch (-(fi9)) {
            case 542:
                if v7 {
                    if (-1 == -9) {
                        fmt.Println(v7, v6, v5)
                        v7 = ((!(!(true && (v7 && v7)))) && ((!(!v7)) || v7))
                    }
                    if false {
                        v3 -= -1
                        v11 := []int{}
                        fmt.Println(v6)
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 7; fi12++ {
                            fmt.Println(v7, v7, fi9)
                            _ = fi12
                        }
                    }
                }
            default:
                v3 = 0
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(fi13, v2, v6)
                v14 := GenS1{F0: fi9}
                if true {
                    v15 := GenS2{F0: 1}
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            _ = fi9
        }
    case "\n\t":
        var v16 bool = false
        _ = v16
    default:
        v5 = "foo bar"
    }
    switch "a" {
    case "hello":
        var v17 uint32 = 243
        _ = v17
    case "\n\t":
        var v18 uint64 = 2
        _ = v18
    }
    var v19 string = "\n\t"
    if (3.14159 <= 0.1) {
        _ = len(v4)
    } else {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: fi20}
            var v22 bool = v7
            v2 += 42
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v23 := []int{-415, -880, 820, 551}
    if (!(!(v19 >= "\n\t"))) {
        v2 = 9223372036854775807
        v24 := GenS1{F0: -2147483648}
        _ = v24
    }
    v25 := GenS1{F0: 4294967295}
    var v26 string = v5
    v7 = v7
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 uint = 131
            v3 = (v3 >> ((255) & 63))
            v7 = true
            _ = fi28
            _ = v29
        }
        var v30 float32 = (1e10 / 1.5)
        _ = fi27
        _ = v30
    }
    var v31 uint8 = 128
    delete(v4, 793)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
