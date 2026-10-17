package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int; F1 int64; F2 int32 }

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    v3 := make(map[int]bool)
    v4 := []int{}
    v5 := make(map[string]int)
    v1 -= v1
    switch ("a" + "a") {
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            switch (("foo bar" + ("a" + "hello")) + ("hello" + "")) {
            case "hello":
                if (false || true) {
                    v1 = 336
                } else {
                    v7 := GenS1{F0: 127, F1: 18446744073709551615, F2: -1}
                    for fi8 := 0; fi8 < 17; fi8++ {
                        if len(v4) > 0 {
                            v4[0] = (fi8 << ((42) & 63))
                        }
                        if true {
                            v9 := GenS1{F0: 255, F1: 231, F2: 127}
                            v2 = v1
                            _ = v9
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v10 int64 = 42
                                _ = v10
                            }
                        }
                        _ = fi8
                    }
                    v2 = (v1 * (-(v2)))
                    _ = v7
                }
            case "\n\t":
                switch fi6 {
                case 42:
                    v1 = v2
                case 195:
                    var v11 float32 = (-2.25 / 3.14159)
                    _ = v11
                case 9223372036854775807:
                    v12 := GenS1{F0: 230, F1: 42, F2: 100}
                    _ = v12
                }
            default:
                v1 = 42
            }
            _ = fi6
        }
    default:
        if true {
            var v13 bool = true
            var v14 bool = (68 >= 0)
            switch ("hello" + "") {
            case "\n\t":
                fmt.Println(v2)
            case "":
                for fi15 := 0; fi15 < 18; fi15++ {
                    switch fi15 {
                    case -105:
                        for fi16 := 0; fi16 < 9; fi16++ {
                            var v17 int = (fi16 - fi15)
                            var v18 uint8 = 50
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    case 32767:
                        v1 -= v1
                    }
                    delete(v5, (("hello" + "foo bar") + ""))
                    _ = fi15
                }
            case "foo bar":
                var v19 string = ("hello" + "")
                _ = v19
            default:
                v20 := GenS2{F0: 992, F1: -688, F2: (v2 - (v1 * v2))}
                _ = v20
            }
            _ = v13
            _ = v14
        }
    }
    _ = len(v4)
    if true {
        for fi21 := 0; fi21 < 19; fi21++ {
            _ = len(v3)
            v2 += (^(v1))
            _ = fi21
        }
    }
    var v22 uint32 = 0
    switch 128 {
    case 9223372036854775807:
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS1{F0: 255, F1: 567, F2: 1}
            _ = fi23
            _ = v24
        }
    case 32768:
        var v25 int32 = v1
        _ = v25
    case 128:
        if true {
            delete(v3, 1)
            var v26 bool = ("a" < "foo bar")
            _ = v26
        } else {
            v1 -= v1
        }
    default:
        v2 = v2
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = 255
        delete(v3, (fi27 << (((v28 % (((v28 | (v28 | v28))) | 1))) & 63)))
        var v29 int16 = 241
        _ = fi27
        _ = v28
        _ = v29
    }
    v2 = int32(v22)
    if false {
        v30 := v3[-325]
        v30 = ((v22 >> ((18446744073709551615) & 63)) > (^(v22)))
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
