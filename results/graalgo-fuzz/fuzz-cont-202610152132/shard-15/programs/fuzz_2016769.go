package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint = 65535
    var v2 uint32 = 42
    var v3 uint64 = 840
    var v4 string = ("\n\t" + "foo bar")
    v5 := []int{127}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 bool = false
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1)
        }
        _ = fi6
        _ = v7
    }
    switch (v4 + v4) {
    case "hello":
        v8 := GenS1{F0: -1}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                fmt.Println(fi9)
                var v10 uint8 = 1
                _ = v10
            }
            _ = fi9
        }
    default:
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    }
    v11 := GenS1{F0: 42}
    v12 := GenS1{F0: 1}
    switch ("" + "") {
    case "":
        switch v4 {
        case "a":
            v4 = (("\n\t" + (("a" + ("foo bar" + v4)) + v4)) + "hello")
        case "\n\t":
            var v13 float32 = 1e10
            _ = v13
        }
    case "\n\t":
        var v14 uint16 = 2
        _ = v14
    case "a":
        v15 := GenS1{F0: 573}
        _ = v15
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS1{F0: 410}
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS1{F0: 0}
                if false {
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                } else {
                    v20 := GenS1{F0: 830}
                    _ = v20
                }
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
    }
    if (3.14159 > 3.14159) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v2, v1)
        v2 -= (v2 | (v2 & uint32(v3)))
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = (v1 * (v1 >> (((v1 & (v1 ^ 1))) & 63)))
        var v21 int64 = -9223372036854775808
        _ = v21
    }
    switch 32768 {
    case 4294967295:
        switch -119 {
        case 0:
            var v22 float64 = (0.0 + (-2.25 - 0.1))
            _ = v22
        case 9223372036854775807:
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
        }
    case 42:
        v23 := GenS1{F0: 1}
        _ = v23
    case -2147483648:
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: -1}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v1, v2)
    fmt.Println(v2, v2, v4)
    fmt.Println(v1, v1)
    v2 -= (v2 & v2)
    v26 := GenS1{F0: -1}
    switch "a" {
    case "a":
        fmt.Println(v1)
    default:
        v1 -= v1
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if false {
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int16 = 47
                switch 128 {
                case -48:
                    v29 -= 1000
                case 128:
                    switch 2147483647 {
                    case 4294967295:
                        var v30 uint32 = 1
                        _ = v30
                    case 42:
                        if false {
                            switch "hello" {
                            case "":
                                v31 := GenS1{F0: v29}
                                _ = v31
                            default:
                                v5 = append(v5, (fi28 ^ fi28))
                            }
                            v32 := GenS1{F0: (v29 / ((v29) | 1))}
                            _ = v32
                        }
                    }
                case -1:
                    fmt.Println(fi28, v4, v29)
                }
                _ = fi28
                _ = v29
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 uint64 = 921
                _ = fi33
                _ = v34
            }
        }
        fmt.Println(v2, v2)
        var v35 float64 = 1.5
        _ = fi27
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
