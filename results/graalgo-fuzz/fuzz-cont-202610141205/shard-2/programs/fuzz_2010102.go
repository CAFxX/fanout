package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
type GenS3 struct { F0 bool }

func main() {
    v1 := []int{88, 343, -1}
    var v2 uint8 = 0
    v3 := []string{}
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v4 string = "\n\t"
    switch "" {
    case "":
        var v5 bool = (true || true)
        _ = v5
    case "\n\t":
        _ = len(v1)
    default:
        v2 -= (^(v2))
    }
    var v6 string = "foo bar"
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 65535, F1: 390, F2: 42}
    var v9 uint = 37
    switch (v6 + "") {
    case "hello":
        if true {
            if (v2 >= v2) {
                fmt.Println(v9, v6)
                if false {
                    if false {
                        var v10 int32 = 2147483647
                        _ = v10
                    } else {
                        var v11 bool = true
                        _ = v11
                    }
                    v7 = v7
                }
            }
            var v12 uint32 = 42
            _ = v12
        } else {
            v3 = append(v3, v4)
            var v13 uint = (v9 & 18446744073709551615)
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint8 = v2
        if false {
            v16 := GenS2{F0: 529, F1: -1, F2: (!(1000 <= 1))}
            var v17 uint64 = 885
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint64 = v17
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            if false {
                var v20 int16 = 37
                var v21 string = v4
                _ = v20
                _ = v21
            }
            fmt.Println(v4, v7, fi14)
            switch (((((v4 + v4) + ("\n\t" + (v4 + ("" + v4)))) + v4) + "hello") + "a") {
            case "\n\t":
                var v22 uint = v9
                _ = v22
            case "hello":
                if true {
                    _ = len(v3)
                    v4 = ("a" + (v4 + v4))
                } else {
                    if (128 > (fi14 / (((fi14 & int(v7))) | 1))) {
                        var v23 uint32 = 204
                        for fi24 := 0; fi24 < 7; fi24++ {
                            v25 := GenS3{F0: ((v15 << (((^(v9))) & 63)) > (v15 + (v2 & (v2 << ((0) & 63)))))}
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    v1 = append(v1, fi14)
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 string = v4
                        _ = fi26
                        _ = v27
                    }
                }
            }
        }
        _ = fi14
        _ = v15
    }
    v6 = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
