package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    var v2 int8 = -128
    var v3 uint16 = 255
    var v4 int32 = -2147483648
    var v5 int8 = 91
    var v6 uint32 = 4294967295
    if false {
        v7 := GenS1{F0: v4}
        switch -660 {
        case 42:
            v1 += (v1 - (v1 * v1))
        }
        _ = v7
    } else {
        var v8 int64 = 42
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 string = "hello"
            _ = fi9
            _ = v10
        }
        v1 += v1
        _ = v8
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint8 = (v1 >> ((234) & 63))
        var v13 uint64 = 0
        if (-2.25 <= 0.0) {
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: (v4 | v4)}
                var v16 uint = 1
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    switch ("" + "foo bar") {
    case "a":
        var v17 string = ("\n\t" + "\n\t")
        _ = v17
    case "\n\t":
        fmt.Println(v1, v5)
    case "":
        for fi18 := 0; fi18 < 10; fi18++ {
            switch 127 {
            case 1:
                v3 = v3
            case 32768:
                var v19 bool = false
                _ = v19
            default:
                switch "hello" {
                case "hello":
                    if true {
                        var v20 float64 = 0.0
                        _ = v20
                    } else {
                        var v21 uint32 = (v6 >> ((2) & 63))
                        _ = v21
                    }
                case "foo bar":
                    switch (fi18 | fi18) {
                    case -1:
                        var v22 uint32 = v6
                        _ = v22
                    case -835:
                        var v23 int16 = 0
                        _ = v23
                    }
                case "\n\t":
                    fmt.Println(v2, v2)
                default:
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := []int{int(v4), fi24, fi18}
                        v25 = append(v25, fi18)
                        _ = fi24
                        _ = v25
                    }
                }
            }
            v3 += 527
            fmt.Println(fi18)
            _ = fi18
        }
    }
    switch "foo bar" {
    case "":
        v26 := []int{-1, 855, 42}
        _ = v26
    case "a":
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint32 = v6
            fmt.Println(v5, v28)
            _ = fi27
            _ = v28
        }
    }
    var v29 float64 = 0.0
    v29 += v29
    if true {
        fmt.Println(v6)
    }
    v29 += v29
    v30 := []int{-836}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
