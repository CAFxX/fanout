package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []uint32{4294967295, 2147483648, 2, 2}
    var v2 uint16 = 460
    var v3 uint64 = 4294967295
    var v4 uint64 = v3
    var v5 int8 = 127
    if ((-2.25 - ((1e10 - (1.5 * -2.25)) - 1e10)) == 1.5) {
        v2 -= v2
    }
    if false {
        if (true || false) {
            if true {
                var v6 bool = (("a" > ("foo bar" + "foo bar")) || true)
                if v6 {
                    v3 -= 2
                } else {
                    var v7 uint64 = 18446744073709551615
                    _ = v7
                }
                _ = v6
            } else {
                var v8 int64 = 0
                var v9 uint32 = 425
                _ = v8
                _ = v9
            }
        } else {
            v10 := GenS1{F0: (^(v3))}
            v4 = (v3 >> ((710) & 63))
            v5 -= 0
            _ = v10
        }
        var v11 int64 = -757
        _ = len(v1)
        _ = v11
    }
    if (true && (42 != 1)) {
        if len(v1) > 0 {
            v1[0] = 921
        }
        v3 = (v3 - (^(v3)))
    } else {
        v3 = (v3 | 0)
        if true {
            v2 = (v2 - v2)
            if (!(true || ((32767 != (v2 << ((1) & 63))) && (!((^(v2)) <= (v2 * 0)))))) {
                v1 = append(v1, 625)
                if false {
                    switch -431 {
                    case -886:
                        fmt.Println(v5)
                    case 2147483647:
                        var v12 int16 = 0
                        _ = v12
                    }
                    switch 465 {
                    case 90:
                        if (345 >= -693) {
                            v2 -= v2
                            v13 := GenS2{F0: (42 > 840), F1: 163}
                            for fi14 := 0; fi14 < 2; fi14++ {
                                for fi15 := 0; fi15 < 18; fi15++ {
                                    v5 -= (v5 << ((365) & 63))
                                    _ = fi15
                                }
                                fmt.Println(v3, v4)
                                _ = fi14
                            }
                            _ = v13
                        }
                    }
                } else {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 int64 = -9223372036854775808
                        v18 := GenS1{F0: 856}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if true {
                            v4 += (v3 << ((42) & 63))
                            v3 = 2
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v20 := GenS1{F0: 1}
                        v21 := GenS2{F0: true, F1: 18446744073709551615}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            }
        } else {
            if ("\n\t" < ((("foo bar" + ("hello" + ("\n\t" + ""))) + "a") + "\n\t")) {
                var v22 float64 = 3.14159
                switch ("foo bar" + "foo bar") {
                case "hello":
                    fmt.Println(v2, v2)
                case "a":
                    v5 = v5
                default:
                    fmt.Println(v4, v22, v22)
                }
                _ = v22
            }
        }
    }
    var v23 int32 = -629
    if (372 > 127) {
        fmt.Println(v23, v3, v3)
    }
    v24 := GenS3{F0: v5}
    var v25 bool = true
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int64 = 1
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v5 = v5
    var v29 uint = 4294967295
    if (!v25) {
        fmt.Println(v23)
        var v30 int = 32767
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
