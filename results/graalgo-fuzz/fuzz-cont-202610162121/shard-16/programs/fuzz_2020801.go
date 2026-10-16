package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 string = "foo bar"
    var v3 float64 = float64((3.14159 + 1.5))
    var v4 int = 1
    v5 := []int{2147483647}
    switch v2 {
    case "\n\t":
        var v6 uint64 = 4294967295
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 17; fi7++ {
            if true {
                if true {
                    if false {
                        v5 = append(v5, v4)
                    } else {
                        if (!false) {
                            if (-669 > -749) {
                                var v8 float32 = -2.25
                                fmt.Println(v3)
                                v8 = v8
                                _ = v8
                            } else {
                                fmt.Println(v2)
                                v9 := GenS1{F0: (v1 >> ((283) & 63)), F1: -354, F2: 1}
                                _ = v9
                            }
                        }
                    }
                    var v10 float64 = v3
                    fmt.Println(v10)
                    _ = v10
                } else {
                    v11 := GenS2{F0: ("\n\t" + "\n\t")}
                    var v12 float64 = 3.14159
                    var v13 bool = true
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := GenS3{F0: -28, F1: (-(v4)), F2: (615 < v1)}
                switch v2 {
                case "hello":
                    var v15 uint32 = 14
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 uint16 = 88
                        fmt.Println(fi7, v17, fi16)
                        _ = fi16
                        _ = v17
                    }
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v14
            }
            _ = fi7
        }
    default:
        v1 = v1
    }
    v18 := GenS2{F0: ((v2 + v2) + "\n\t")}
    fmt.Println(v3, v2, v2)
    if false {
        for fi19 := 0; fi19 < 19; fi19++ {
            switch v2 {
            case "a":
                v1 = (^(v1))
            case "hello":
                switch (fi19 << ((464) & 63)) {
                case 0:
                    var v20 uint8 = uint8(v1)
                    _ = v20
                case 2147483647:
                    var v21 int64 = 9223372036854775807
                    _ = v21
                case -1:
                    v1 = v1
                }
            }
            v4 = 563
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 uint32 = (v1 >> ((18446744073709551615) & 63))
                v4 = 0
                _ = fi22
                _ = v23
            }
            _ = fi19
        }
        var v24 int32 = -1
        var v25 string = v2
        _ = v24
        _ = v25
    }
    switch ((("a" + "") + "hello") + "") {
    case "a":
        if true {
            var v26 uint16 = 0
            _ = v26
        } else {
            var v27 uint32 = 4294967295
            _ = len(v5)
            switch v4 {
            case 128:
                for fi28 := 0; fi28 < 9; fi28++ {
                    if len(v5) > 0 {
                        v5[0] = fi28
                    }
                    if true {
                        var v29 uint32 = v1
                        _ = v29
                    }
                    _ = fi28
                }
            }
            _ = v27
        }
    case "":
        v4 = v4
    case "\n\t":
        v30 := GenS4{F0: 732}
        _ = v30
    default:
        v31 := GenS4{F0: 0}
        _ = v31
    }
    switch (("a" + v2) + v2) {
    case "\n\t":
        _ = len(v5)
    case "":
        v32 := GenS5{F0: (v4 ^ v4)}
        _ = v32
    }
    var v33 uint16 = 65535
    v34 := GenS6{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
