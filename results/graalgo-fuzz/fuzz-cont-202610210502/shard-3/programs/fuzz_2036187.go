package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int }

func main() {
    v1 := []int{0, 127}
    var v2 int16 = 42
    var v3 int = 4294967295
    v3 -= (v3 + v3)
    _ = len(v1)
    var v4 int = (v3 - v3)
    v5 := GenS1{F0: -126}
    var v6 int16 = 1000
    var v7 int = (v4 << ((2) & 63))
    for fi8 := 0; fi8 < 11; fi8++ {
        v6 -= v2
        if ((-2.25 * 0.0) != 1.5) {
            v7 += (^(v7))
        }
        var v9 uint32 = 393
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 42}
    var v11 uint = 0
    v11 = v11
    var v12 uint = v11
    if false {
        if len(v1) > 0 {
            v1[0] = (v7 + v3)
        }
        if ((42 > 1) && (false || (true && (!(false && false))))) {
            switch 9223372036854775807 {
            case 4294967295:
                var v13 uint64 = 0
                _ = v13
            case 128:
                v14 := GenS1{F0: 42}
                _ = v14
            }
            v15 := GenS2{F0: 65535}
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v2 = v2
            var v17 uint = v11
            _ = fi16
            _ = v17
        }
    } else {
        if (0.1 != 0.0) {
            v1 = append(v1, (v7 + v4))
            if (("\n\t" + (("a" + "\n\t") + "a")) > ("foo bar" + "")) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if true {
                    v18 := GenS2{F0: 89}
                    if (true && (609 < v6)) {
                        v19 := GenS1{F0: 127}
                        v20 := GenS3{F0: v11, F1: (v3 & (v4 & 128))}
                        _ = v19
                        _ = v20
                    } else {
                        v12 += (v12 & 4294967295)
                    }
                    _ = v18
                } else {
                    v6 += 702
                    fmt.Println(v2, v3, v6)
                }
            }
        } else {
            switch (v3 & 90) {
            case 135:
                v3 = 32768
            case 2147483647:
                v1 = append(v1, v3)
            }
            if false {
                if true {
                    var v21 int = v4
                    var v22 uint32 = 0
                    switch "a" {
                    case "\n\t":
                        v4 = 2147483647
                    default:
                        var v23 bool = false
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                }
                v11 = v11
            }
            v1 = append(v1, v7)
        }
        fmt.Println(v6, v2)
        v1 = append(v1, 32767)
    }
    var v24 float64 = (((3.14159 * 0.1) - float64(0.1)) / -2.25)
    switch (("" + "foo bar") + "a") {
    case "":
        v2 = v2
    case "a":
        v12 = v12
    case "\n\t":
        var v25 int = 0
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
