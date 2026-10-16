package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() int16 {
    var v7 float64 = -2.25
    v8 := []int{}
    v9 := GenS1{F0: "foo bar", F1: 0, F2: 42}
    _ = v7
    _ = v8
    _ = v9
    return 968
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 float64 = 3.14159
    v4 := make(map[uint32]bool)
    var v5 uint8 = 67
    var v6 uint8 = v5
    v5 += (^(v5))
    if false {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                v5 = 127
                var v11 uint64 = 2
                _ = fi10
                _ = v11
            }
        }
        switch "foo bar" {
        case "foo bar":
            var v12 uint64 = 18446744073709551615
            _ = v12
        case "":
            v3 = v3
        }
    }
    var v13 uint32 = 4294967295
    var v14 float64 = v3
    v3 -= (v3 - v14)
    if ((!(274 > 0)) || (!true)) {
        switch "" {
        case "a":
            for fi15 := 0; fi15 < 5; fi15++ {
                if true {
                    var v16 uint16 = 394
                    for fi17 := 0; fi17 < 14; fi17++ {
                        fmt.Println(v6, fi17)
                        v4[(^(v13))] = false
                        _ = fi17
                    }
                    switch -247 {
                    case 128:
                        var v18 uint32 = v13
                        _ = v18
                    case -228:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v16 += v16
                            _ = fi19
                        }
                    case 528:
                        var v20 int32 = -1
                        _ = v20
                    default:
                        _ = len(v1)
                    }
                    _ = v16
                }
                _ = fi15
            }
        case "":
            v21 := GenS2{F0: 13}
            _ = v21
        }
        var v22 uint8 = v5
        _ = v22
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            v3 = 0.1
            var v24 float32 = (1e10 + 0.1)
            v1[fi23] = true
            _ = fi23
            _ = v24
        }
        switch 1 {
        case 9223372036854775807:
            var v25 uint16 = 65535
            _ = v25
        default:
            for fi26 := 0; fi26 < 5; fi26++ {
                v27 := GenS1{F0: "a", F1: 684, F2: 2}
                fmt.Println(v14)
                if len(v2) > 0 {
                    v2[0] = 143
                }
                _ = fi26
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 int16 = -32768
                _ = fi29
                _ = v30
            }
            var v31 uint8 = 127
            _ = fi28
            _ = v31
        }
    }
    delete(v4, 4294967295)
    v5 = (v6 >> ((65535) & 63))
    for fi32 := 0; fi32 < 17; fi32++ {
        v5 += 42
        var v33 int32 = 42
        _ = fi32
        _ = v33
    }
    v34 := GenS2{F0: 56}
    switch 32768 {
    case -869:
        var v35 uint64 = 2
        _ = v35
    }
    var v36 int32 = 2147483647
    switch 9223372036854775807 {
    case -9223372036854775808:
        v2 = append(v2, 152)
    }
    switch "hello" {
    case "hello":
        switch ("\n\t" + ("a" + "foo bar")) {
        case "foo bar":
            if false {
                var v37 int32 = (v36 + v36)
                switch 32768 {
                case 604:
                    v13 = v13
                case 2147483647:
                    for fi38 := 0; fi38 < 16; fi38++ {
                        fmt.Println(v37, v5)
                        _ = fi38
                    }
                }
                _ = v37
            }
        }
    case "":
        var v39 string = "\n\t"
        _ = v39
    case "foo bar":
        v4[(v13 >> ((171) & 63))] = (65535 != 18446744073709551615)
    }
    v5 = (v6 + v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
