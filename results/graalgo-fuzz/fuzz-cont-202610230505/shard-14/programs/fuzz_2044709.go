package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint32 = 65535
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    if (v2 >= v2) {
        fmt.Println(v2, v3, v2)
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        _ = len(v4)
        _ = fi7
    }
    if false {
        var v8 string = ""
        _ = v8
    } else {
        v9 := GenS1{F0: 0, F1: 255}
        var v10 uint = 18446744073709551615
        _ = len(v4)
        _ = v10
        _ = v9
    }
    var v11 uint32 = (^(v2))
    if (v11 != v2) {
        for fi12 := 0; fi12 < 19; fi12++ {
            if true {
                delete(v5, fi12)
                v11 -= (^(v2))
                for fi13 := 0; fi13 < 5; fi13++ {
                    if (1.5 > -2.25) {
                        var v14 uint8 = 128
                        switch 555 {
                        case 1:
                            var v15 uint64 = 4294967295
                            _ = v15
                        }
                        _ = v14
                    } else {
                        v16 := GenS1{F0: -30, F1: 2}
                        v17 := GenS1{F0: 100, F1: 287}
                        v1 = (fi13 << ((0) & 63))
                        var v18 bool = true
                        var v19 uint = 18446744073709551615
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    if (((42 < 42) && (true || (true && false))) || (0 == 42)) {
                        switch -751 {
                        case 9223372036854775807:
                            v3 -= v3
                            v2 = 0
                            v20 := GenS1{F0: 42, F1: 155}
                            _ = v20
                        case 42:
                            v11 += 65535
                        default:
                            v21 := GenS2{F0: -9223372036854775808, F1: 0}
                            _ = v21
                        }
                        delete(v6, "\n\t")
                    } else {
                        v2 = v11
                        var v22 uint32 = 330
                        _ = v22
                    }
                    var v23 int16 = 42
                    v11 = (v11 % ((v11) | 1))
                    _ = fi13
                    _ = v23
                }
            }
            _ = fi12
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            v2 = v11
            v3 = -2.25
            _ = fi24
        }
    } else {
        if (316 <= 9223372036854775807) {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v1, v1)
            case "hello":
                for fi25 := 0; fi25 < 20; fi25++ {
                    if true {
                        delete(v4, 2147483647)
                    }
                    _ = fi25
                }
            }
            v2 = v2
        } else {
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int32 = 1
                var v28 uint32 = (v2 * v2)
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    if (false && (!true)) {
        if true {
            for fi29 := 0; fi29 < 8; fi29++ {
                var v30 int16 = 1000
                var v31 int64 = -335
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 bool = true
            v11 -= v2
            _ = v32
        } else {
            switch v1 {
            case -218:
                if true {
                    v1 += (v1 << ((4294967295) & 63))
                    if ("hello" > "hello") {
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            delete(v6, "a")
                        case "foo bar":
                            for fi33 := 0; fi33 < 13; fi33++ {
                                v1 = fi33
                                _ = fi33
                            }
                        default:
                            switch v1 {
                            case 9223372036854775807:
                                v2 = v11
                            }
                        }
                    }
                    v2 = v11
                }
            case 0:
                if (61 == 288) {
                    _ = len(v6)
                }
            }
        }
        v2 += (v11 | 65535)
        switch (((("foo bar" + ("" + "")) + "foo bar") + "a") + "foo bar") {
        case "a":
            var v34 uint16 = 42
            _ = v34
        }
    } else {
        fmt.Println(v3, v2, v11)
        v35 := v5[(v1 - (v1 / ((v1) | 1)))]
        _ = v35
    }
    switch "\n\t" {
    case "":
        var v36 bool = false
        _ = v36
    default:
        v2 = 2
    }
    var v37 int32 = 1
    for fi38 := 0; fi38 < 11; fi38++ {
        v39 := v6["foo bar"]
        v37 -= v37
        _ = fi38
        _ = v39
    }
    v40 := []int{(^(v1))}
    v41 := GenS2{F0: -9223372036854775808, F1: -128}
    var v42 int8 = 1
    v2 -= (v11 / ((v11) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
