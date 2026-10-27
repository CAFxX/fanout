package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int8 = 42
    v2 := []string{}
    var v3 uint16 = 42
    var v4 string = "foo bar"
    v5 := []uint32{}
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            switch v4 {
            case "a":
                switch "a" {
                case "":
                    switch (fi6 * (fi6 >> ((365) & 63))) {
                    case 2147483647:
                        _ = len(v5)
                    }
                }
            case "\n\t":
                var v7 bool = (!(1e10 > 0.0))
                _ = v7
            }
        }
        v5 = append(v5, 0)
        _ = fi6
    }
    fmt.Println(v1, v3, v1)
    if (-1 < 42) {
        v4 = v4
        var v8 uint8 = 1
        if true {
            var v9 uint8 = (v8 >> ((435) & 63))
            switch (v4 + v4) {
            case "\n\t":
                v1 += (^(v1))
            }
            _ = v9
        } else {
            v10 := GenS1{F0: 0}
            _ = v10
        }
        _ = v8
    }
    fmt.Println(v4, v4)
    switch v4 {
    case "":
        switch v4 {
        case "\n\t":
            var v11 string = (v4 + "foo bar")
            _ = v11
        case "hello":
            if true {
                v3 = (v3 ^ 0)
                if true {
                    if (v4 <= "a") {
                        var v12 int64 = 42
                        _ = v12
                    } else {
                        var v13 int = 42
                        _ = v13
                    }
                    switch "hello" {
                    case "\n\t":
                        var v14 string = v4
                        _ = v14
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "hello":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            for fi16 := 0; fi16 < 19; fi16++ {
                                switch (fi15 | (-(fi16))) {
                                case -389:
                                    var v17 int32 = -2147483648
                                    _ = v17
                                case 0:
                                    v4 = (v4 + (v4 + v4))
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v4 = ""
                                _ = fi16
                            }
                            _ = len(v2)
                            _ = fi15
                        }
                    default:
                        var v18 string = v4
                        _ = v18
                    }
                }
                v3 = v3
            }
        case "foo bar":
            var v19 uint32 = 65535
            _ = v19
        default:
            v5 = append(v5, 4294967295)
        }
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        v1 = (v1 % ((v1) | 1))
    }
    v20 := GenS1{F0: 32767}
    v1 += (v1 & v1)
    v3 -= 118
    if true {
        v3 -= v3
        v21 := GenS2{F0: 2147483647, F1: 2, F2: -341}
        _ = v21
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 = append(v5, 636)
    v4 = (((("" + ("hello" + "a")) + "foo bar") + v4) + ("foo bar" + v4))
    if true {
        var v22 bool = true
        switch ("\n\t" + v4) {
        case "hello":
            v22 = v22
        case "foo bar":
            v4 = "foo bar"
        case "a":
            v23 := GenS3{F0: 1}
            _ = v23
        default:
            var v24 uint16 = v3
            _ = v24
        }
        _ = v22
    }
    var v25 bool = (!true)
    var v26 int16 = 420
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
