package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 1000
    var v3 uint64 = 2
    var v4 string = "\n\t"
    v4 = v4
    switch -39 {
    case 9223372036854775807:
        if (v3 <= (v3 >> ((0) & 63))) {
            fmt.Println(v3)
            var v5 string = v4
            _ = v5
        }
    default:
        var v6 uint8 = 196
        _ = v6
    }
    switch "\n\t" {
    case "a":
        var v7 uint64 = v3
        _ = v7
    case "\n\t":
        v8 := []int{32768, 127, 128}
        _ = v8
    default:
        fmt.Println(v1, v2, v3)
    }
    var v9 int64 = -9223372036854775808
    var v10 int32 = 1
    switch (v4 + (v4 + v4)) {
    case "a":
        v9 -= 1
    }
    if false {
        v11 := GenS1{F0: v3}
        v12 := []int{}
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS1{F0: v3}
                var v16 uint64 = 2
                var v17 int32 = v10
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi13
        }
        _ = v11
        _ = v12
    } else {
        if false {
            switch -1 {
            case 561:
                v10 = -1
            default:
                var v18 uint16 = v1
                _ = v18
            }
        }
        fmt.Println(v1, v10)
        v4 = ("a" + v4)
    }
    var v19 int64 = v9
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(v9, v1, v2)
            fmt.Println(v9, v1)
            v2 = (v2 * (-(v2)))
            _ = fi21
        }
        var v22 uint64 = 2
        _ = fi20
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        switch v4 {
        case "a":
            fmt.Println(fi23, v19, v9)
        case "foo bar":
            switch int(v19) {
            case 882:
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 int32 = (v10 & v10)
                    switch "\n\t" {
                    case "":
                        fmt.Println(v4, v10)
                    }
                    _ = fi24
                    _ = v25
                }
            case 32767:
                var v26 int64 = (^(v9))
                _ = v26
            }
        default:
            switch (fi23 * 723) {
            case 1:
                if false {
                    var v27 uint = 1
                    if true {
                        v19 += v19
                        if true {
                            v28 := GenS1{F0: v3}
                            _ = v28
                        }
                        v29 := GenS1{F0: (^(v3))}
                        _ = v29
                    }
                    _ = v27
                } else {
                    if (!false) {
                        v30 := GenS1{F0: v3}
                        _ = v30
                    }
                    v9 += v9
                    var v31 uint8 = 255
                    _ = v31
                }
            case 127:
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 int16 = -32768
                    _ = fi32
                    _ = v33
                }
            default:
                v9 = -1
            }
        }
        v34 := GenS1{F0: v3}
        var v35 uint32 = 2147483648
        _ = fi23
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
