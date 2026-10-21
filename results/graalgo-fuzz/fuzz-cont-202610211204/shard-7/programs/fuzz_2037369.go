package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    v1 := []int{0, 42, 9223372036854775807}
    v2 := []int{651, 252, 792}
    var v3 uint16 = 65535
    v4 := make(map[int]int)
    var v5 uint = uint(v3)
    var v6 uint16 = 2
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 20; fi7++ {
            if (false || (-1 < -1)) {
                var v8 uint32 = 2147483648
                var v9 int8 = 0
                var v10 int8 = v9
                _ = v10
                _ = v8
                _ = v9
            } else {
                v6 = (v6 & v6)
                switch "foo bar" {
                case "a":
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 bool = false
                        var v13 uint16 = 255
                        if (!(-475 < 1)) {
                            v2 = append(v2, (fi11 + -2147483648))
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            }
            var v14 int = fi7
            _ = fi7
            _ = v14
        }
    case "hello":
        v15 := GenS1{F0: 1, F1: ((521 < 1000) || (!false))}
        _ = v15
    default:
        v16 := GenS1{F0: 0, F1: (1 > int64(v6))}
        _ = v16
    }
    v1 = append(v1, 1)
    v6 += 962
    v6 = v3
    var v17 uint16 = (v6 / (((^(v6))) | 1))
    v18 := v4[-2147483648]
    v17 += (v6 / ((32767) | 1))
    var v19 bool = false
    for fi20 := 0; fi20 < 9; fi20++ {
        switch v18 {
        case -2147483648:
            v21 := v4[(-(fi20))]
            _ = v21
        case 127:
            v22 := GenS1{F0: 255, F1: v19}
            _ = v22
        case -473:
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 uint32 = 959
                _ = fi23
                _ = v24
            }
        }
        v25 := v4[v18]
        _ = fi20
        _ = v25
    }
    v26 := GenS1{F0: 0, F1: (!true)}
    fmt.Println(v17, v3, v17)
    if v19 {
        switch "a" {
        case "\n\t":
            if false {
                for fi27 := 0; fi27 < 4; fi27++ {
                    v5 -= (^(v5))
                    var v28 string = ""
                    switch (v28 + (v28 + "\n\t")) {
                    case "foo bar":
                        var v29 int16 = -1
                        _ = v29
                    default:
                        var v30 uint8 = 255
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
            }
        default:
            switch ("\n\t" + "hello") {
            case "\n\t":
                v17 = (v3 << (((v5 ^ v5)) & 63))
            case "a":
                var v31 int8 = 0
                _ = v31
            }
        }
        fmt.Println(v19, v17, v5)
    } else {
        switch (v18 / (((-(v18))) | 1)) {
        case -1:
            var v32 int32 = 42
            _ = v32
        }
        fmt.Println(v19)
    }
    v33 := GenS1{F0: 255, F1: v19}
    fmt.Println(v18)
    var v34 int = (v18 / ((v18) | 1))
    var v35 int8 = -26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
